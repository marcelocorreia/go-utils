package folder_utils

import (
	"github.com/marcelocorreia/go-utils/utils"
	"os"
)

type FolderUtils struct {
	
}

func (fu FolderUtils) InitialCheckAndCreate(path string) (error) {
	result, err := utils.Exists(path)
	if err != nil {
		return err
	}
	if !result {
		os.MkdirAll(path, 0755)
	}
	return nil
}
