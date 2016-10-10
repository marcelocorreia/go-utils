package vault_runner

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestMountsExists(t *testing.T) {
	vh := VaultHelper{
		VaultAddress:"http://127.0.0.1:8200",
		VaultToken: "86447a85-6389-3e94-a95e-b9c57bdde1a1",
	}
	exists, _ := vh.MountExists("sys/")
	assert.True(t, exists)
}

func TestListMounts(t *testing.T) {
	vh := VaultHelper{
		VaultAddress:"http://127.0.0.1:8200",
		VaultToken: "86447a85-6389-3e94-a95e-b9c57bdde1a1",
	}

	mounts, _ := vh.Mounts()
	size := len(mounts)
	fmt.Printf("Found %d mounts", size)
	assert.True(t, size > 0)
}