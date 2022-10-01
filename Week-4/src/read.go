package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and
represents it in a slice of structs. Assume that there
is a text file which contains a series of names.

Each line of the text file has a first name and a last name,
in that order, separated by a single space on the line.

Your program will define a name struct which has two fields,
fname for the first name, and
lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file
and create a struct which contains the
first and last names found in the file.

Each struct created will be added to a slice,
and after all lines have been read from the file,
your program will have a slice containing one struct
for each line in the file.

After reading all lines from the file,
your program should iterate through your slice of structs and
print the first and last names found in each struct.
*/
type Person struct {
	fname string
	lname string
}

func getFileName() string {
	fmt.Println("\nEnter File Path with Name:\te.g ./assets/names.txt ")
	userinput := bufio.NewScanner(os.Stdin)
	userinput.Scan()
	output := userinput.Text()
	return output
}

func getPersonSlice(file *os.File) []Person {
	reader := bufio.NewReader(file)
	var personSlice = make([]Person, 0)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fields := strings.Split(string(line), " ")
		person := Person{fname: fields[0], lname: fields[1]}
		personSlice = append(personSlice, person)
	}
	return personSlice
}

func getFileHandle(path string) (bool, *os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, nil, err
	}

	if fi, err := file.Stat(); err != nil || fi.IsDir() {
		return false, nil, err
	}
	return true, file, nil
}

func main() {
	fmt.Println("\nWeek4 - Assignment 2")

	exist, file, _ := getFileHandle(getFileName())
	if !exist {
		fmt.Printf("File DOES NOT exist.\n")
		os.Exit(0)
	}

	persons := getPersonSlice(file)
	for _, person := range persons {
		fmt.Printf("fname: [%-20s]; lname: [%-20s]\n", person.fname, person.lname)
	}
}
