package runtime_helper

import (
	"os/exec"
	"bytes"
	"os"
	"strings"
)

type RuntimeService interface {
	RunCommand(command string, arg ...string) (string, error)
	RunCommandLogStream(command string, arg ...string) (error)
	RunThis(fullCommand string) (string, error)
	CheckBinaryInPath(binary string) bool
}

type RuntimeHelper struct{}

func New() (*RuntimeService) {
	var rt RuntimeService
	rt = RuntimeHelper{}
	return &rt
}

func (r RuntimeHelper) RunThis(fullCommand string) (string, error) {
	c := strings.Split(fullCommand, " ")[0]
	a := strings.Join(strings.Split(fullCommand, " ")[1:], " ")
	var args []string
	if len(a) == 0 {
		args = []string{}
	} else {
		args = nil
	}
	cmd := exec.Command(c, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), err
}

func (r RuntimeHelper) RunCommand(command string, arg ...string) (string, error) {
	cmd := exec.Command(command, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), err
}

func (r RuntimeHelper) CheckBinaryInPath(binary string) bool {
	_, err := exec.LookPath(binary)
	if err != nil {
		return false
	}
	return true
}

func (r RuntimeHelper) RunCommandLogStream(command string, arg ...string) (error) {
	cmd := exec.Command(command, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
}
