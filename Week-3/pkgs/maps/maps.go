package maps

/*
- What are Maps?

Maps are the Implementation of HashTables in Golang

- We use make() to create a map e.g
*/

//Variables

//	var idMap map[string]int //Where string is KEY, int is VALUE
// OR
//	var idMap = make(map[string]int)

// Map Literals
var idMap = map[string]int{"Joe": 1, "Hassan": 2, "Jane": 3}

//If inside function, then:
// idMap := map[string]int {"Joe":1}

func Called() string {
	return "maps(hashtables) package called"
}

func ExportidMap() map[string]int {
	return idMap
}

func AddKeyValuePair(key string, value int) {
	idMap[key] = value
}
