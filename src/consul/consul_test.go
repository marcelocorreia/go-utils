package consul

import (
	"testing"
	"github.com/hashicorp/consul/api"
	"fmt"
	"github.com/stretchr/testify/assert"
)

var consulService ConsulService

func TestConsulService_Init(t *testing.T) {
	config := api.DefaultConfig()
	
	fmt.Println(config.Scheme, config.Address)
	
}

func TestConsulService_SaveUpdateKV(t *testing.T) {
	consulService.SaveUpdateKV("Hey", "Ho")
}

func TestKV(t *testing.T) {
	testPath := "test/hey"
	testValue := "ho lets go"
	
	consulService.SaveUpdateKV(testPath, testValue)
	exists, _ := consulService.KeyExists(testPath)
	doesNotExist, _ := consulService.KeyExists("bogus/path")
	assert.True(t, exists)
	assert.False(t, doesNotExist)
	val, _ := consulService.GetKV(testPath)
	assert.Equal(t, val, testValue)
	consulService.DeleteKV(testPath)
}

func TestConsulService_ListServices(t *testing.T) {
	services := consulService.ListServices()
	for k, v := range services {
		fmt.Println(k, "->", v)
	}
}

func TestConsulService_ListNodes(t *testing.T) {
	nodes := consulService.ListNodes()
	for _, node := range nodes {
		fmt.Println(node.Node, "->", node.Address)
	}
}

func TestConsulService_Register(t *testing.T) {
	serviceName:= "testService"
	check := make(map[string]string)
	check["HTTP"] = "http://127.0.0.1:8500"
	check["Interval"] = "10s"
	consulService.Register(serviceName, "127.0.0.1", 8500, check,[]string{serviceName})
	for k,v:=range consulService.ListServices(){
		fmt.Println(k,"->",v)
	}
	
	consulService.DeRegister(serviceName)
}

