package main

import "fmt"

func main() {
	printMe("Hello, World!")
	var result, remainder = intDivision(5, 2)
	fmt.Println("The result is", result)
	fmt.Println("The remainder is", remainder)
	arrays()
	slices()
	maps()

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}

funct arrays(){
	var intArr [3]int32 //can't change the length of an array once it's declared
	intArr[0] = 1
	fmt.Println(intArr[1:3])
}

funct slices(){
	//slices are like arrays but can change in size
	var intSlice []int32
	intSlice = append(intSlice, 1)
	fmt.Println("The length of intSlice is", len(intSlice))
	intSlice = append(intSlice, 2)
	intSlice = append(intSlice, 3)
	fmt.Println("The length of intSlice is", len(intSlice))
	fmt.Println(intSlice)
	var intSlice2 []int32 = make(int32[], 3, 8)
}

func maps(){
	//maps are like dictionaries in python
	var intMap map[string]int32 = make(map[string]int32)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	fmt.Println(intMap)

	var map2 = map[string]uint8{"Mickey": 15, "Minnie": 20}
	fmt.Println(map2["Minnie"])
	var age, ok = map2["Goofy"]
	if ok{
		fmt.Println("Goofy's age is", age)
	}
	else{
		fmt.Println("Goofy is not in the map")
	}
}
