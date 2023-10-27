package helloWorld

import "fmt"

const (
	Public = 2
	private = 3
)

const (
	// 10, 11, 12, 13
	const1 = iota + 10
	const2 
	const3
	const4
)

// it will be public if starts with a capitalize letter, else it is private
func HelloWorld() (string, string) {
	hello := "Hello"
	var world string = "World!"
	return hello, world
}

func ContagemSlice() {
	mapInteger := make(map[string]int)
	
	sliceInt := make([]int, 0)
	sliceInt = append(sliceInt, 1)
	sliceInt = append(sliceInt, 2)

	mapInteger["tree"] = 3
	mapInteger["four"] = 4

	for index, value := range sliceInt {
		fmt.Println(index, " ",value)
	}

	for index, value := range mapInteger {
		fmt.Println(index, " ",value)
	}

	// return default value, in this case will be 0 (integer)
	fmt.Println(mapInteger["seis"])
}

func TransformToHelloWorld(str *string){
	*str = "Hello World";
}