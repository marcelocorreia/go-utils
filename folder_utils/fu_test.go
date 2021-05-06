package folder_utils_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/marcelocorreia/go-utils/utils"
	"os"
)

func TestRunCommand(t *testing.T) {
	fu := *New()
	tmpFolder := "/tmp/cu-da-tia"
	fu.InitialCheckAndCreate(tmpFolder)
	ck, _ := utils.Exists(tmpFolder)
	assert.True(t, ck)
	os.Remove(tmpFolder)
}
