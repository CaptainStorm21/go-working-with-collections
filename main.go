package main

import (
	"fmt"
)
func main() {
	// arrays

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	//output [1 2 3]
	fmt.Println(arr)

	//ouput [2]
	fmt.Println(arr[1])

	// implicit initiationzation format
	numArray := [3]int{ 5, 6, 7}
	//output [5 6 7]
	fmt.Println(numArray)

	// slices

	slicesArray := [3]int{ 10, 20, 30}
	slice:= slicesArray[:]


	//output [10 20 30] [10 20 30]
	fmt.Println(slicesArray, slice)

	slicesArray[2]=200
	slice[1] = 9000
	//output [10 9000 200] [10 9000 200]
	fmt.Println(slicesArray, slice)


	openSlice := []int{1000, 2000, 5000}
	fmt.Println(openSlice)

	appendSlice := append(openSlice, 10)
	//output [1000 2000 5000 10]
	fmt.Println(appendSlice)

	appendSlice1 := append(appendSlice, 10, 30, 1233)
	//output [1000 2000 5000 10 10 30 1233]
	fmt.Println(appendSlice1)

	sliced1 :=  appendSlice1[1:]
	sliced2 :=  appendSlice1[:2]
	sliced3 := appendSlice1[1:2]
	// outpput [2000 5000 10 10 30 1233] [1000 2000] [2000]
	fmt.Println(
				sliced1,
				sliced2,
				sliced3 )
	// maps



	// structs


}

