package helper

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func ParseXML(path string) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var f FBAPIWithResponse
	err2 := xml.Unmarshal(dat, &f)
	if err2 != nil {
		panic(err2)
	}
	fmt.Printf("%v", f)
}
