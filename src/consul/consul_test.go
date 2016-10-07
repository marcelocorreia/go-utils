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
	//consulService.Register("zza", "127.0.0.1", 10101)
	//consulService.Register("b", 32400)
	//consulService.Register("d", 10101)
	//consulService.Register("e", 32400)
	//
	//defaultCheckTimeout := 30 * time.Second
	//fmt.Println(defaultCheckTimeout.String())
	//consulService.Blah()
	consulService.DeRegister("xxa")
	consulService.DeRegister("xxaa")
	consulService.DeRegister("xx")
	//consulService.DeRegister("zz")
	//consulService.DeRegister("zza")
	//consulService.DeRegister("d")
	//consulService.DeRegister("e")
	//x := make(map[string]string)
	//x["HTTP"] = "http://127.0.0.1:7777/health"
	//x["Interval"] = "10s"
	//consulService.Register("xxaa", "127.0.0.1", 7777,x,[]string{"x"})
}

func shit(values interface{}) {
	fmt.Println(values)
}
