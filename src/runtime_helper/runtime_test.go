package runtime

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestRunCommand(t *testing.T) {
	r := GetRuntimeService()
	args := []string{"-ls"}
	resp, e := r.RunCommand("ls", args)
	fmt.Println(resp, e)
	assert.NotEmpty(t, resp)
	e = r.RunCommandLogStream("ls", args)
}

func TestCheckBinaryInPath(t *testing.T) {
	r := GetRuntimeService()
	assert.True(t, r.CheckBinaryInPath("ls"))
	assert.False(t, r.CheckBinaryInPath("dfghdhnedtumdfychb56urth45bertaw34bt "))
}



