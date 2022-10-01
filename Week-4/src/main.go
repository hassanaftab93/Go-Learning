package main

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
1. RFCs - Request for comments
	e.g
		HTML	1866
		URI		3986
		HTTP	2616

	There are packages that help to process RFCs
		encode/decode protocol format

		e.g Package:
			"net/http" - Web communication protocols
				http.Get("www.google.com")

		e.g	Package:
			"net" - TCP/IP socket programming
				net.Dial("tcp","uci.edu:80")


2. JSON - JavaScript Object Notation //ALL UNICODE - READABLE by HUMANS
	e.g
		RFC		7159

	Format to present structured Data

		{
			"name":"joe",
			"addr":"islamabad",
			"phone":"1234"
		}

	JSON Marshalling is converting a Go object -> JSON


3. File Access | ioutil
	e.g
		File access always has mechanical delay when HDDs are involved
		Linear access

	Basic Operations
		1. Open/Close
		2. Read/Write
		3. Seek - Move read/write head (for skipping linear approach)

	e.g Package:
		"ioutil"
			- Large files can be issues
*/

type Person struct {
	Name    string
	Address string
	Phone   string
}

func MarshalFunction(obj Person) []byte { //										Go->JSON
	barr, _ := json.Marshal(obj)
	return barr
}

func UnMarshalFunction(obj []byte) Person { //										JSON->Go
	var p2 Person
	err := json.Unmarshal(obj, &p2)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return p2
}

func ReadFile(Path string) string {
	data, _ := os.ReadFile(Path)
	return string(data)
}

func WriteFile(Path string, data []byte) { //										Write to file
	err := os.WriteFile(Path, data, 0777)
	if err == nil {
		fmt.Println("File Written")
	} else {
		fmt.Println("\nError: ", err)
	}
}

func main() {
	fmt.Println("\nWeek 4 Start")

	p1 := Person{Name: "Hassan", Address: "123ee", Phone: "923333"} //				Person Object

	jsonObject := MarshalFunction(p1) //											Go->JSON
	fmt.Println("\nGo->JSON: ", jsonObject)

	p2 := UnMarshalFunction(jsonObject) //											JSON->Go
	p2.Name = "Aftab"
	fmt.Println("\nJSON->Go: ", p2)

	fmt.Println("\nRead from File: ", ReadFile("./assets/text.txt")) //				Read File Function

	fmt.Println("\nWriting to File: ")
	data := "This is modified data"
	WriteFile("./assets/writefileFunction.txt", []byte(data))
}
