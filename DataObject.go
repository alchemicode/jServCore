package jServCore

import (
	"encoding/json"
	"encoding/xml"
	"fmt"

	toml "github.com/pelletier/go-toml/v2"
	yaml "gopkg.in/yaml.v3"
)

type DataObject struct {
	Id   string                 `json:"id" msgpack:"id" xml:"id" toml:"id" yaml:"id"`
	Data map[string]interface{} `json:"data" msgpack:"data" xml:"data" toml:"data" yaml:"data"`
}

// Default Constructor
// Creates an empty Object with only an id
func (d *DataObject) WithoutData(id string) {
	d.Id = id
	d.Data = make(map[string]interface{})
}

// Map Constructor
// Creates an Object with given id and data map
func (d *DataObject) WithData(id string, data map[string]interface{}) {
	d.Id = id
	d.Data = data
}

// Json Constructor
// Creates an Object from given Json string
func (d *DataObject) FromJSON(s string) {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(s), &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat["id"])
	d.Id = dat["id"].(string)
	d.Data = dat["data"].(map[string]interface{})
}
func (d *DataObject) FromTOML(s string) {
	var dat map[string]interface{}
	if err := toml.Unmarshal([]byte(s), &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat["id"])
	d.Id = dat["id"].(string)
	d.Data = dat["data"].(map[string]interface{})
}
func (d *DataObject) FromXML(s string) {
	var dat map[string]interface{}
	if err := xml.Unmarshal([]byte(s), &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat["id"])
	d.Id = dat["id"].(string)
	d.Data = dat["data"].(map[string]interface{})
}
func (d *DataObject) FromYAML(s string) {
	var dat map[string]interface{}
	if err := yaml.Unmarshal([]byte(s), &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat["id"])
	d.Id = dat["id"].(string)
	d.Data = dat["data"].(map[string]interface{})
}

// Returns the Object as a map
func (d DataObject) ToMap() map[string]interface{} {
	m := make(map[string]interface{})
	m["id"] = d.Id
	m["data"] = d.Data
	return m
}

// Returns the Object in Json string format
func (d DataObject) ToJSON() string {
	js, _ := json.Marshal(d)
	return string(js)
}

// Returns the Object as a string
func (d DataObject) String() string {
	return fmt.Sprintf(" \"id\" : %d\n \"data\" : %v", d.Id, d.Data)
}
