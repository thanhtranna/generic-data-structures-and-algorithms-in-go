package main

import (
	"fmt"

	"chapter07/set"
)

func main() {
	setOne := set.Set[int]{}
	setOne.Insert(3)
	setOne.Insert(5)
	setOne.Insert(7)
	setOne.Insert(9)
	setTwo := set.Set[int]{}
	setTwo.Insert(3)
	setTwo.Insert(6)
	setTwo.Insert(8)
	setTwo.Insert(9)
	setTwo.Insert(11)
	setTwo.Delete(11)
	fmt.Println("Items in setTwo: ", setTwo.Items())

	fmt.Println("5 in setOne: ", setOne.In(5))
	fmt.Println("5 in setTwo: ", setTwo.In(5))

	fmt.Println("Union of setOne and setTwo: ", setOne.Union(setTwo).Items())
	fmt.Println("Intersection of setOne and setTwo: ", setOne.Intersection(setTwo).Items())
	fmt.Println("Difference of setTwo with respect to setOne: ", setTwo.Difference(setOne).Items())
	fmt.Println("Size of this difference: ", setOne.Intersection(setTwo).Size())
}

/* Output
Items in setTwo:  [8 9 3 6]
5 in setOne:  true
5 in setTwo:  false
Union of setOne and setTwo:  [3 5 7 9 8 6]
Intersection of setOne and setTwo:  [3 9]
Difference of setTwo with respect to setOne:  [6 8]
Size of this difference:  2
*/
