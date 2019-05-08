package utils

import (
	"os"
	"fmt"
	"io"
	"io/ioutil"
	"errors"
)

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err == nil {
			err = os.Chmod(dest, sourceinfo.Mode())
			if err != nil {
				return nil
			}
		}
	}

	return nil
}

func ListDir(dir string) ([]os.FileInfo) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	return files
}

func ListDirWithExceptions(dir string, exceptions []string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {

		return []os.FileInfo{}, err
	}
	resp := []string{}
	for _, file := range files {
		if file.IsDir() {
			if !StringInSlice(file.Name(), exceptions) {
				resp = append(resp, file.Name())
			}
		}
	}

	return files, nil
}

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func CreateNewDirectoryIfNil(path string) (bool,error) {
	exists, err := Exists(path)
	if err != nil {
		return false,err
	}
	if !exists {
		os.MkdirAll(path, 00750)
		return true,nil
	} else {
		return false,errors.New("path already exists. please check your parameters")
	}
}

func WriteFile(path string, content string) (error) {
	err := ioutil.WriteFile(path, []byte(content), 0750)
	if err != nil {
		return err
	}
	return nil
}
