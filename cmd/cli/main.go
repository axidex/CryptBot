package main

import (
	"bot/internal/problems"
	"fmt"
)

func main() {
	text := "047b16c2"
	keys := "67e99b3c"

	_, err := problems.Feistel(text, keys)
	if err != nil {
		fmt.Println(err)
		return
	}

}
