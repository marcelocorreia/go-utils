package runtime_helper_test

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/marcelocorreia/go-utils/runtime_helper"
)

func TestRunCommand(t *testing.T) {
	r := *runtime_helper.New()
	args := []string{"-ls"}
	resp, e := r.RunCommand("ls", args)
	fmt.Println(resp, e)
	assert.NotEmpty(t, resp)
	e = r.RunCommandLogStream("ls", args)
}

func TestCheckBinaryInPath(t *testing.T) {
	r := *runtime_helper.New()
	assert.True(t, r.CheckBinaryInPath("ls"))
	assert.False(t, r.CheckBinaryInPath("dfghdhnedtumdfychb56urth45bertaw34bt "))
}



