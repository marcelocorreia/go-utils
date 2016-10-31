package machine_io

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

type TestType struct {
	Name     string `json:"name" yaml:"name"`
	Age      int `json:"age" yaml:"age"`
	Children []TestType `json:"children" yaml:"children"`
}

func TestOutput(t *testing.T) {
	
	dad := Sample()
	
	outJ, errJ := JsonOutput(dad)
	assert.Nil(t, errJ)
	fmt.Println(outJ)
	outY, errY := YamlOutput(dad)
	assert.Nil(t, errY)
	
	fmt.Println(outY)
}

func Sample() (TestType) {
	son := TestType{
		Name: "Jim", Age:50,
	}
	
	daughter := TestType{
		Name: "Mary", Age:10,
	}
	
	dad := TestType{
		Name: "John", Age:98,
		Children:[]TestType{son, daughter},
	}
	return dad
}


