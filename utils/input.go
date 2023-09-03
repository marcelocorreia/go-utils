package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Returns the input of the question
*/
func Question(question string) string {
	var input string
	fmt.Print(question)
	fmt.Scanln(&input)
	return input
}

func QuestionF(format string, question ...string) string {
	fmt.Printf(format, question)
	in := bufio.NewReader(os.Stdin)
	resp, _ := in.ReadString('\n')
	return strings.Replace(resp, "\n", "", -1)
}

func QuestionWithDefault(question, defaultValue string, defaultYes bool) string {

	var short, long, badge string
	if defaultYes {
		short = "y"
		long = "yes"
		badge = "[Y/n]"
	} else {
		short = "n"
		long = "no"
		badge = "[y/N]"
	}

	fmt.Printf("%s %s: ", question, badge)
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)

	empty := "\n"

	if input == empty || input == short || input == long {
		return defaultValue
	} else {
		return input
	}
}
