package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.

Your program should create a map and add the name and address to the map
using the keys “name” and “address”, respectively.

Your program should use Marshal() to create a JSON object from the map,
and then your program should print the JSON object.


*/

var idMap = map[string]string{}

func AddKeyValuePair(name string, address string) {
	idMap[name] = address
}

func ReturnMap() map[string]string {
	return idMap
}

func WriteFile(Path string, data []byte) { //										Write to file
	err := os.WriteFile(Path, data, 0777)
	if err == nil {
		fmt.Println("\nFile Written to Directory ./assets/")
	} else {
		fmt.Println("\nError: ", err)
	}
}

func MapToJSON(mapObj map[string]string) {
	bArray, _ := json.Marshal(mapObj) //			Conversion Go Map -> JSON
	fmt.Println("\nGo -> JSON Converted")
	fmt.Println(bArray)
	WriteFile("./assets/converted.JSON", bArray)
}

func getName() string {
	fmt.Println("\nEnter Name: ")
	userinput := bufio.NewScanner(os.Stdin)
	userinput.Scan()
	output := userinput.Text()
	return output
}

func getAddr() string {
	fmt.Println("\nEnter Address: ")
	userinput := bufio.NewScanner(os.Stdin)
	userinput.Scan()
	output := userinput.Text()
	return output
}

func main() {
	fmt.Println("\nWeek 4 - Assignment 1")

	AddKeyValuePair(getName(), getAddr()) //			Getting Name,Address and Adding the key/val to the Map
	MapToJSON(ReturnMap())                //			Returning the Map for conversion purposes

}
