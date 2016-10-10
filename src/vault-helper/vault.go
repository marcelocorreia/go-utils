package vault_runner

import (
	"github.com/hashicorp/vault/api"
	"os"
	"strings"
	"github.com/op/go-logging"
)

var (
	logger *logging.Logger
)

type VaultHelper interface {
	GetClient()
	GetNewClientToken(clientId string) (string, error)
	Init()
	ReadSecret(path string) (*api.Secret, error)
	RootToken() (string)
	SecretExists(path string) bool
	Unseal()
	WriteSecret(path string, secret map[string]interface{}) (*api.Secret, error)
}

type VaultService struct {
	VaultAddress string `json:"vault_address" json:"vault_address"`
	VaultToken string `json:"vault_token" yaml:"vault_token"`
}

func (v VaultService) GetClient() api.Client {
	os.Setenv("VAULT_ADDR", v.VaultAddress)
	config := api.DefaultConfig()
	config.Address = v.VaultAddress
	client, err := api.NewClient(config)
	client.SetToken(v.VaultToken)
	if err != nil {
		logger.Fatalf("err: %s", err)
	}
	
	return *client
}

func (v VaultService) WriteSecret(path string, secret map[string]interface{}) (*api.Secret, error) {
	os.Setenv("VAULT_ADDR", v.VaultAddress)
	client := v.GetClient()

	secResp, err := client.Logical().Write(path, secret)

	if err != nil {
		logger.Error(err)
		return nil, err
	}
	
	return secResp, nil
}

func (v VaultService) ReadSecret(path string) (*api.Secret, error) {
	os.Setenv("VAULT_ADDR", v.VaultAddress)
	client := v.GetClient()
	client.SetToken(v.VaultToken)
	resp, err := client.Logical().Read(path)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v VaultService) RootToken() (string) {
	os.Setenv("VAULT_ADDR", v.VaultAddress)
	
	return v.VaultToken
}

func (v VaultService) GetNewClientToken(clientId string) (string, error) {
	os.Setenv("VAULT_ADDR", v.VaultAddress)
	
	return "", nil
}

func (v VaultService) SecretExists(path string) bool {
	_, err := v.ReadSecret(path)
	if err != nil {
		return false
	}
	
	return true
}

func (v VaultService) Mount(path string, mountType string) (error) {
	client := v.GetClient()
	client.SetToken(v.VaultToken)
	mountInfo := &api.MountInput{
		Type:mountType,
		Description:path,
	}
	err := client.Sys().Mount(path, mountInfo)

	if err != nil {
		logger.Error(err)
	}
	
	return nil
}

func (v VaultService) Mounts() (map[string]*api.MountOutput, error) {
	client := v.GetClient()
	oo, err := client.Sys().ListMounts()
	if err != nil {
		return nil, err
	}
	
	return oo, nil
}

func (v VaultService) MountExists(mount string) (bool, error) {
	mounts, err := v.Mounts()
	if err != nil {
		return false, err
	}
	
	for k, _ := range mounts {
		if k == strings.Trim(mount, " ") {
			return true, nil
		}
	}
	return false, nil
}
