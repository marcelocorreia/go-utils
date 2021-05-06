package utils_test

import (
	"fmt"
	"github.com/marcelocorreia/go-utils/utils"
	"testing"
)

func TestQuestionWithDefault(t *testing.T) {
	name :="Jimi Hendrix"
	input := utils.QuestionWithDefault(fmt.Sprintf("Name: [%s] ",name),name,true )
	fmt.Println(input)
}
