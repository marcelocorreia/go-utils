package utils_test

import (
	"fmt"
	"testing"
)

func TestQuestionWithDefault(t *testing.T) {
	name :="Jimi Hendrix"
	input := QuestionWithDefault(fmt.Sprintf("Name: [%s] ",name),name,true )
	fmt.Println(input)
}
