package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func user_input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(prompt)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input

}

func main() {

	question := []string{
		"who am i: ",
		"who is the president of US: ",
		"What is the name of prachanda: ",
	}

	answer := []string{
		"sisir",
		"donald trump",
		"prasanna",
	}

	score := 0

	for i := 0; i < len(question); i++ {

		user_answer := user_input(question[i])

		if user_answer == answer[i] {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect")
		}

	}

	fmt.Printf("Your total score is %v", score)
}
