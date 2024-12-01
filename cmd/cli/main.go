package main

import (
	"bot/internal/problems"
	"fmt"
)

func main() {

	steps := problems.Md5(5, 1, 9, 1, 6, 10)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	fmt.Println(steps)

}
