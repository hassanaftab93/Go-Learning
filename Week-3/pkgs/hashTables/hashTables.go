package hashTables

/*
- What are hashTables?

They contain key/value pairs e.g
	- SSN/Email
	- Address/GPS-Coords
	- Name/Phone-Number

A HashTable is supposed to store these key/value pairs
	- Each value has a unique key
	- For this, A Hash Function is used
	- You don't explicitly call this "Hash" function, it goes on in the background.

Basic use is to make it easier for the programmer
	- An Array /Slice can be accessed using keys instead of just an index number e.g arr[1]
	- So we will maybe do something like arr[JOE], arr[JANE]

Advantages:
	- Faster lookup than lists							-> Constant time vs. Linear time like in lists
	- Not ints, but Keys that have a meaningful name
Disadvantages:
	- Could have collisions								-> If hash func maps JOE and JANE to same slot then you have a collision
	- Collisions can be avoided, however.
	- But collisions cause performance issues
	- Hash function is made in such a way. it is not impossible for a collision to occur, but the chances are still there.
*/

func Called() string {
	return "hashTables package called"
}
