package i18n

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

func UnmarshallJson(data []byte) ([]*Message, error) {
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	return recGetMessages(v, isMessage(v), true)
}

func UnmarshallToml(data []byte) ([]*Message, error) {
	var v interface{}
	err := toml.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	return recGetMessages(v, isMessage(v), true)
}

func UnmarshallYaml(data []byte) ([]*Message, error) {
	var v interface{}
	err := yaml.Unmarshal(data, &v)
	if err != nil {
		return nil, err
	}

	return recGetMessages(v, isMessage(v), true)
}

const nestedSeparator = "."

var errInvalidTranslationFile = errors.New("invalid translation file, expected key-values, got a single value")

// recGetMessages looks for translation messages inside "raw" parameter,
// scanning nested maps using recursion.
func recGetMessages(raw interface{}, isMapMessage, isInitialCall bool) ([]*Message, error) {
	var messages []*Message
	var err error

	switch data := raw.(type) {
	case string:
		if isInitialCall {
			return nil, errInvalidTranslationFile
		}
		m, err := NewMessage(data)
		return []*Message{m}, err

	case map[string]interface{}:
		if isMapMessage {
			m, err := NewMessage(data)
			return []*Message{m}, err
		}
		messages = make([]*Message, 0, len(data))
		for id, data := range data {
			// recursively scan map items
			messages, err = addChildMessages(id, data, messages)
			if err != nil {
				return nil, err
			}
		}

	case map[interface{}]interface{}:
		if isMapMessage {
			m, err := NewMessage(data)
			return []*Message{m}, err
		}
		messages = make([]*Message, 0, len(data))
		for id, data := range data {
			strid, ok := id.(string)
			if !ok {
				return nil, fmt.Errorf("expected key to be string but got %#v", id)
			}
			// recursively scan map items
			messages, err = addChildMessages(strid, data, messages)
			if err != nil {
				return nil, err
			}
		}

	case []interface{}:
		// Backward compatibility for v1 file format.
		messages = make([]*Message, 0, len(data))
		for _, data := range data {
			// recursively scan slice items
			childMessages, err := recGetMessages(data, isMessage(data), false)
			if err != nil {
				return nil, err
			}
			messages = append(messages, childMessages...)
		}

	default:
		return nil, fmt.Errorf("unsupported file format %T", raw)
	}

	return messages, nil
}

func addChildMessages(id string, data interface{}, messages []*Message) ([]*Message, error) {
	isChildMessage := isMessage(data)
	childMessages, err := recGetMessages(data, isChildMessage, false)
	if err != nil {
		return nil, err
	}
	for _, m := range childMessages {
		if isChildMessage {
			if m.ID == "" {
				m.ID = id // start with innermost key
			}
		} else {
			m.ID = id + nestedSeparator + m.ID // update ID with each nested key on the way
		}
		messages = append(messages, m)
	}
	return messages, nil
}

func parsePath(path string) (langTag, format string) {
	formatStartIdx := -1
	for i := len(path) - 1; i >= 0; i-- {
		c := path[i]
		if os.IsPathSeparator(c) {
			if formatStartIdx != -1 {
				langTag = path[i+1 : formatStartIdx]
			}
			return
		}
		if path[i] == '.' {
			if formatStartIdx != -1 {
				langTag = path[i+1 : formatStartIdx]
				return
			}
			if formatStartIdx == -1 {
				format = path[i+1:]
				formatStartIdx = i
			}
		}
	}
	if formatStartIdx != -1 {
		langTag = path[:formatStartIdx]
	}
	return
}
