package machine_io

import (
	"testing"
	"fmt"
)

type TestType struct {
	Name     string `json:"name" yaml:"name"`
	Age      int `json:"age" yaml:"age"`
	Children []TestType `json:"children" yaml:"children"`
}

func TestOutput(t *testing.T) {
	
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
	
	outJ, _ := JsonOutput(dad)
	fmt.Println(outJ)
	outY, _ := YamlOutput(dad)
	
	fmt.Println(outY)
}


