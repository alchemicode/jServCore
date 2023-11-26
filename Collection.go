package jServCore

import (
	"fmt"
	"os"
	"path/filepath"

	m "github.com/vmihailenco/msgpack/v5"
)

type Collection struct {
	Name string
	List []DataObject
}

// Named Constructor Constructor
// Creates an empty Collection object with just a name
// and reads its file
func (c *Collection) New(name string) {
	c.Name = name
	c.List = make([]DataObject, 0)
	c.ReadFile()
}

// Asynchronously reads a collection's Json file
func (c *Collection) ReadFile() {
	//Channel provides success data
	ch := make(chan bool)
	go c.readFile(ch)
	result := <-ch
	if !result {
		fmt.Println("Failed to read database: " + c.Name + ".dat")
	}
}

// Reads the collection from its file
func (c *Collection) readFile(ch chan bool) {
	path := filepath.Join("Collections", c.Name+".dat")
	//Reads the contents of the file as a []byte
	content, err2 := os.ReadFile(path)
	if err2 != nil {
		//Channel returns false if there is any error
		ch <- false
		fmt.Println("Error while reading file " + c.Name + ".dat")
	} else {
		//Generates map data from the file
		var dat []map[string]interface{}
		if err3 := m.Unmarshal(content, &dat); err3 != nil {
			ch <- false
			fmt.Println("Error when generating data for " + c.Name + ".dat")
		} else {
			//Reads each object in the generated data from the file
			//and populates the collection's list
			for i := 0; i < len(dat); i++ {
				obj := new(DataObject)
				obj.WithData(dat[i]["id"].(string), dat[i]["data"].(map[string]interface{}))
				c.List = append(c.List, *obj)
			}
			//Channel returns true if the read was successful
			ch <- true
		}

	}

}

// Asynchronously updates the collection's Json file
func (c *Collection) UpdateFile() {
	//Channel provides success data
	ch := make(chan bool)
	go c.updateFile(ch)
	result := <-ch
	if result {
		fmt.Println(" * Updated " + c.Name + ".dat")
	} else {
		fmt.Println(" > Failed to update " + c.Name + ".dat")
	}
}

func (c *Collection) updateFile(ch chan bool) {
	path := filepath.Join("Collections", c.Name+".dat")
	//Converts the collection's list data into Json data
	bytes, _ := m.Marshal(c.List)
	//Overwrites Json file with new data
	if err := os.WriteFile(path, bytes, 0644); err != nil {
		//Channel returns false if there is any error
		ch <- false
		fmt.Println(" > Error when opening file " + c.Name + ".dat")
	} else {
		//Channel returns true if the update was successful
		ch <- true
	}
}
