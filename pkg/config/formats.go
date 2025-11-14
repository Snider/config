package config

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
)

// ConfigFormat defines the interface for loading and saving configuration data.
type ConfigFormat interface {
	Load(path string) (map[string]interface{}, error)
	Save(path string, data map[string]interface{}) error
}

// JSONFormat implements the ConfigFormat interface for JSON files.
type JSONFormat struct{}

// Load reads a JSON file and returns a map of key-value pairs.
func (f *JSONFormat) Load(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Save writes a map of key-value pairs to a JSON file.
func (f *JSONFormat) Save(path string, data map[string]interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, jsonData, 0644)
}

// YAMLFormat implements the ConfigFormat interface for YAML files.
type YAMLFormat struct{}

// Load reads a YAML file and returns a map of key-value pairs.
func (f *YAMLFormat) Load(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Save writes a map of key-value pairs to a YAML file.
func (f *YAMLFormat) Save(path string, data map[string]interface{}) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(path, yamlData, 0644)
}

// INIFormat implements the ConfigFormat interface for INI files.
type INIFormat struct{}

// Load reads an INI file and returns a map of key-value pairs.
func (f *INIFormat) Load(path string) (map[string]interface{}, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	for _, section := range cfg.Sections() {
		for _, key := range section.Keys() {
			result[section.Name()+"."+key.Name()] = key.Value()
		}
	}
	return result, nil
}

// Save writes a map of key-value pairs to an INI file.
func (f *INIFormat) Save(path string, data map[string]interface{}) error {
	cfg := ini.Empty()
	for key, value := range data {
		parts := strings.SplitN(key, ".", 2)
		section := ini.DefaultSection
		keyName := parts[0]
		if len(parts) > 1 {
			section = parts[0]
			keyName = parts[1]
		}
		if _, err := cfg.Section(section).NewKey(keyName, fmt.Sprintf("%v", value)); err != nil {
			return err
		}
	}
	return cfg.SaveTo(path)
}

// XMLFormat implements the ConfigFormat interface for XML files.
type XMLFormat struct{}

type xmlEntry struct {
	Key   string `xml:"key"`
	Value string `xml:"value"`
}

// Load reads an XML file and returns a map of key-value pairs.
func (f *XMLFormat) Load(path string) (map[string]interface{}, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var v struct {
		XMLName xml.Name   `xml:"config"`
		Entries []xmlEntry `xml:"entry"`
	}
	if err := xml.Unmarshal(data, &v); err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	for _, entry := range v.Entries {
		result[entry.Key] = entry.Value
	}
	return result, nil
}

// Save writes a map of key-value pairs to an XML file.
func (f *XMLFormat) Save(path string, data map[string]interface{}) error {
	var entries []xmlEntry
	for key, value := range data {
		entries = append(entries, xmlEntry{
			Key:   key,
			Value: fmt.Sprintf("%v", value),
		})
	}
	xmlData, err := xml.MarshalIndent(struct {
		XMLName xml.Name   `xml:"config"`
		Entries []xmlEntry `xml:"entry"`
	}{Entries: entries}, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, xmlData, 0644)
}

// GetConfigFormat returns a ConfigFormat implementation based on the file extension.
func GetConfigFormat(path string) (ConfigFormat, error) {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".json":
		return &JSONFormat{}, nil
	case ".yaml", ".yml":
		return &YAMLFormat{}, nil
	case ".ini":
		return &INIFormat{}, nil
	case ".xml":
		return &XMLFormat{}, nil
	default:
		return nil, fmt.Errorf("unsupported config format: %s", ext)
	}
}

// SaveKeyValues saves a map of key-value pairs to a file, using the appropriate format.
func (s *Service) SaveKeyValues(key string, data map[string]interface{}) error {
	format, err := GetConfigFormat(key)
	if err != nil {
		return err
	}
	filePath := filepath.Join(s.ConfigDir, key)
	return format.Save(filePath, data)
}

// LoadKeyValues loads a map of key-value pairs from a file, using the appropriate format.
func (s *Service) LoadKeyValues(key string) (map[string]interface{}, error) {
	format, err := GetConfigFormat(key)
	if err != nil {
		return nil, err
	}
	filePath := filepath.Join(s.ConfigDir, key)
	return format.Load(filePath)
}
