package jServCore

import (
	"encoding/json"
	"fmt"
)

type AttributeContainer struct {
	Key   string      `json:"key" msgpack:"key" xml:"key" toml:"key" yaml:"key"`
	Value interface{} `json:"value" msgpack:"value" xml:"value" toml:"value" yaml:"value"`
}

// Default Constructor
func (ac *AttributeContainer) New(key string, value interface{}) {
	ac.Key = key
	ac.Value = value
}

// Converts the container to Json text
func (ac AttributeContainer) ToJSON() string {
	js, _ := json.Marshal(ac.ToMap())
	return string(js)
}

func (ac AttributeContainer) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["key"] = ac.Key
	m["value"] = ac.Value
	return m
}

// Converts the container to a string
func (ac AttributeContainer) String() string {
	return fmt.Sprintf("{ \" %s \" : %v  }", ac.Key, ac.Value)
}
