package program1

import "fmt"

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {	
		digit := number % 10		
		sum = sum + (digit * digit)
		number = number / 10	
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {	
		digit := number % 10		
		sum = sum + (digit * digit * digit)
		number = number / 10	
	}
	cubeop <- sum
}

func CalcCubeAndSquare() {
	number := 589
	squareChan := make(chan int)
	cubeChan := make(chan int)
	go calcSquares(number, squareChan)
	go calcCubes(number, cubeChan)
	squares := <- squareChan
	cubes := <-cubeChan
	// fmt.Println(squares, cubes)
	fmt.Println("Final output ", squares + cubes)
}