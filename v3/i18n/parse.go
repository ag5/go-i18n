package i18n

import (
	"fmt"

	"golang.org/x/text/language"
)

// MessageFile represents a parsed message file.
type MessageFile struct {
	Path     string
	Tag      language.Tag
	Format   string
	Messages []*Message
}

// ParseMessageFileBytes returns the messages parsed from file.
func ParseMessageFileBytes(buf []byte, path string, unmarshalFuncs map[string]UnmarshalFunc) (*MessageFile, error) {
	lang, format := parsePath(path)
	tag := language.Make(lang)
	messageFile := &MessageFile{
		Path:   path,
		Tag:    tag,
		Format: format,
	}
	if len(buf) == 0 {
		return messageFile, nil
	}
	unmarshalFunc := unmarshalFuncs[messageFile.Format]
	if unmarshalFunc == nil {
		if messageFile.Format == "json" {
			unmarshalFunc = UnmarshallJson
		} else {
			return nil, fmt.Errorf("no unmarshaler registered for %s", messageFile.Format)
		}
	}
	var err error
	if messageFile.Messages, err = unmarshalFunc(buf); err != nil {
		return nil, err
	}

	return messageFile, nil
}
