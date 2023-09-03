package folder_utils

import (
	"github.com/marcelocorreia/go-utils/utils"
	"os"
)

type FolderUtils interface {
	InitialCheckAndCreate(path string) error
}
type FU struct {
}

func New() *FolderUtils {
	var fu FolderUtils
	fu = FU{}
	return &fu
}
func (fu FU) InitialCheckAndCreate(path string) error {
	result, err := utils.Exists(path)
	if err != nil {
		return err
	}
	if !result {
		os.MkdirAll(path, 0755)
	}
	return nil
}
