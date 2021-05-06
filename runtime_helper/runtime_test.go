package runtime_helper_test

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestRunCommand(t *testing.T) {
	r := *New()
	args := []string{"-ls"}
	resp, e := r.RunCommand("ls", args...)
	fmt.Println(resp, e)
	assert.NotEmpty(t, resp)
	e = r.RunCommandLogStream("ls", args...)

	resp, err := r.RunCommand("asdakaucixaj8aysdm8-ahs[cioja8[su")
	assert.Error(t, err)
	assert.Empty(t, resp)

	err = r.RunCommandLogStream("asdakaucixaj8aysdm8-ahs[cioja8[su")
	assert.Error(t, err)
}

func TestCheckBinaryInPath(t *testing.T) {
	r := *New()
	assert.True(t, r.CheckBinaryInPath("ls"))
	assert.False(t, r.CheckBinaryInPath("dfghdhnedtumdfychb56urth45bertaw34bt "))
}

func TestRuntimeHelper_RunThis(t *testing.T) {
	r := *New()
	o, err := r.RunThis("uname")
	assert.NoError(t, err)
	assert.NotEmpty(t, o)

	o, err = r.RunThis("ls -l")
	assert.NoError(t, err)
	assert.NotEmpty(t, o)

	o, err = r.RunThis("asdakaucixaj8aysdm8-ahs[cioja8[suscds")
	assert.Error(t, err)
	assert.Empty(t, o)

	o, err = r.RunThis("asdakaucixaj8aysdm8-ahs[cioja8[sxzxzx uscds")
	assert.Error(t, err)
	assert.Empty(t, o)
}
