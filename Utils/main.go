package Utils

import (
	"fmt"
	"time"
)

func AssertAnswer(question string, answer interface{}, countVar *int) {

	var input interface{}

	time.Sleep(1 * time.Second)

	fmt.Println(question)

	fmt.Scan(&input)

	if input == answer {
		time.Sleep(1 * time.Second)
		fmt.Println("Correct!")
		*countVar++
	} else {
		fmt.Println("Oh close")
	}
}
