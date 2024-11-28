package main

import (
	"bot/internal/problems"
	"fmt"
)

func main() {

	steps := problems.BlowFish(4, 9, 1, 5, 3, 4, 5)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	fmt.Println(steps)

}
