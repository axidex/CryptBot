package main

import (
	"bot/internal/problems"
	"fmt"
)

func main() {

	_, err := problems.Des("гана", "куб")
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
