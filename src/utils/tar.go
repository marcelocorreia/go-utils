package utils

import (
	"os"
	"runtime-helper"
	"runtime"
)

func Tar(source, target string) {
	rt := runtime.RuntimeService{}

	args := []string{"-cvzf", target, source}
	rt.RunCommand("tar", args)
}

func Untar(tarball, targetDir string) error {
	rt := runtime.RuntimeService{}
	os.Chdir(targetDir)
	args := []string{"-xvzf", tarball}
	_, error := rt.RunCommand("tar", args)
	return error
}