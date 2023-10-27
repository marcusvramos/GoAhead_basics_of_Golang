package main

import (
	"fmt"
	"time"
	"sync"
	"goahead/helloWorld"
	"goahead/geometria"
)


func main() {
	channel := make(chan int)
	go count(channel)

	for num := range channel {
		fmt.Println(num)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	var two, tree int
	go func(){
		defer wg.Done()
		two = return2()
	}()

	go func(){
		defer wg.Done()
		tree = return3()
	}()

	wg.Wait()
	fmt.Println(two+tree)
}

func return3() int {
	time.Sleep(2 * time.Second)
	return 3
}

func return2() int {
	return 2
}

func count(channel chan int) {
	// defer is used to execute something after end of function
	defer close(channel) 
	for i := 0; i < 5; i++ {
		channel <- i
		time.Sleep(1 * time.Second)
	}
}

func geo() {
	mySquare := geometria.NewSquare(2,4)
	myCircle := geometria.NewCircle(5)

	fmt.Println(mySquare.Area())
	fmt.Println(myCircle.Area())
}

func hello() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)

	var var1 string = "var string"
	va2 := "infer string"
	var var3 = "declare03"

	formattedPrint := fmt.Sprintf("%s, %s, %s", var1, va2, var3)
	fmt.Println(formattedPrint)

	var one, two int = 1, 2
	oneInf, twoInf := 1, 2
	formattedPrint02 := fmt.Sprintf("%d, %d, %d, %d", one, two, oneInf, twoInf)
	fmt.Println(formattedPrint02)

	x, _ := helloWorld.HelloWorld()
	fmt.Printf("%s e variavel nao utilizada", x)

	testVariable := "inutil"

	helloWorld.ContagemSlice()
	helloWorld.TransformToHelloWorld(&testVariable)
	fmt.Printf("%s", testVariable)
}
