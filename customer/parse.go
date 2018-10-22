package customer

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

// ParseXMLFile parses a given XML file into an
// APIResponse struct and prints it - for testing purposes
func ParseXMLFile(path string) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var f APIResponse
	err2 := xml.Unmarshal(dat, &f)
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("%v", f)
}
