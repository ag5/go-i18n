package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/ag5/go-i18n/v2/i18n"
	"github.com/ag5/go-i18n/v2/internal/plural"
	"golang.org/x/text/language"
	yaml "gopkg.in/yaml.v2"
)

func writeFile(outdir, label string, langTag language.Tag, format string, messageTemplates map[string]*i18n.MessageTemplate, sourceLanguage bool) (path string, content []byte, err error) {
	var extension string
	content, extension, err = marshal(messageTemplates, sourceLanguage, format)
	if err != nil {
		return "", nil, fmt.Errorf("failed to marshal %s strings to %s: %s", langTag, format, err)
	}
	path = filepath.Join(outdir, fmt.Sprintf("%s.%s.%s", label, langTag, extension))
	return
}

func marshallDefault(messageTemplates map[string]*i18n.MessageTemplate, sourceLanguage bool) interface{} {
	v := make(map[string]interface{}, len(messageTemplates))

	for id, template := range messageTemplates {
		if other := template.PluralTemplates[plural.Other]; sourceLanguage && len(template.PluralTemplates) == 1 &&
			other != nil && template.Description == "" && template.LeftDelim == "" && template.RightDelim == "" {
			v[id] = other.Src
		} else {
			m := map[string]string{}
			if template.Description != "" {
				m["description"] = template.Description
			}
			if !sourceLanguage {
				m["hash"] = template.Hash
			}
			for pluralForm, template := range template.PluralTemplates {
				m[string(pluralForm)] = template.Src
			}
			v[id] = m
		}
	}
	return v
}

func marshallLokaliseJson(messageTemplates map[string]*i18n.MessageTemplate, sourceLanguage bool) interface{} {
	v := make(map[string]interface{}, len(messageTemplates))

	for id, template := range messageTemplates {
		m := map[string]any{}
		if template.Description != "" {
			m["notes"] = template.Description
		}
		if !sourceLanguage {
			m["hash"] = template.Hash
		}

		if len(template.PluralTemplates) == 1 {
			m["translation"] = template.PluralTemplates[plural.Other].Src
		} else {
			t := map[string]string{}
			for pluralForm, tmpl := range template.PluralTemplates {
				t[string(pluralForm)] = tmpl.Src
			}
			m["translation"] = t
		}

		v[id] = m
	}

	return v
}

func marshal(messageTemplates map[string]*i18n.MessageTemplate, sourceLanguage bool, format string) ([]byte, string, error) {
	switch format {
	case "json":
		v := marshallDefault(messageTemplates, sourceLanguage)
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", "  ")
		err := enc.Encode(v)
		return buf.Bytes(), format, err

	case "lokalise-json":
		v := marshallLokaliseJson(messageTemplates, sourceLanguage)
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", "  ")
		err := enc.Encode(v)
		return buf.Bytes(), "json", err

	case "toml":
		v := marshallDefault(messageTemplates, sourceLanguage)
		var buf bytes.Buffer
		enc := toml.NewEncoder(&buf)
		enc.Indent = ""
		err := enc.Encode(v)
		return buf.Bytes(), format, err

	case "yaml":
		v := marshallDefault(messageTemplates, sourceLanguage)
		output, err := yaml.Marshal(v)
		return output, format, err
	}

	return nil, "", fmt.Errorf("unsupported format: %s", format)
}
