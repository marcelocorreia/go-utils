package consul

import (
	"github.com/hashicorp/consul/api"
	"fmt"
)

type ConsulService struct {
}

type Consul interface {
	SaveUpdateKV(keyPath string, value string) (error)
	GetKV(k string) (string, error)
	Register(service string, address string, port int, check *api.AgentServiceCheck, tags []string) (error)
	Deregister(service string) (error)
	DeleteKV(k string) (error)
	KeyExists(path string)
	ListServices() (map[string][]string)
	ListNodes() ([]*api.Node)
}

func (c ConsulService) SaveUpdateKV(keyPath string, value string) (error) {
	ccc, err := api.NewClient(api.DefaultConfig())
	kv := ccc.KV()
	p := &api.KVPair{Key: keyPath, Value: []byte(value)}
	_, err = kv.Put(p, nil)
	if err != nil {
		return err
	}
	
	return nil
}

func (c ConsulService) KeyExists(path string) (bool, error) {
	fmt.Println("Checking path:", path)
	
	client, err := api.NewClient(api.DefaultConfig())
	
	if err != nil {
		return false, err
	}
	
	// Get a handle to the KV API
	kv := client.KV()
	
	// Lookup the pair
	kp, _, errPair := kv.Get(path, nil)
	
	if errPair != nil {
		return false, errPair
	}
	
	if (kp != nil) {
		return true, nil
	}
	return false, nil
}

func (c ConsulService) DeleteKV(k string) (error) {
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		return err
	}
	
	kv := client.KV()
	_, err = kv.Delete(k, nil)
	
	if err != nil {
		return err
	}
	
	return nil
}

func (c ConsulService) GetKV(k string) (string, error) {
	// Get a new client
	
	client, err := api.NewClient(api.DefaultConfig())
	
	if err != nil {
		return "", err
	}
	
	// Get a handle to the KV API
	kv := client.KV()
	
	// Lookup the pair
	pair, _, err := kv.Get(k, nil)
	if err != nil {
		return "", err
	}
	
	return string(pair.Value), nil
}

func (c ConsulService) ListServices() (map[string][]string) {
	client, _ := api.NewClient(api.DefaultConfig())
	cat := client.Catalog()
	services, _, _ := cat.Services(nil)
	
	return services
}

func (c ConsulService) ListNodes() ([]*api.Node) {
	client, _ := api.NewClient(api.DefaultConfig())
	cat := client.Catalog()
	nodes, _, _ := cat.Nodes(nil)
	
	return nodes
}

func (c ConsulService) Register(service string, address string, port int, check map[string]string, tags []string) (error) {
	client, _ := api.NewClient(api.DefaultConfig())
	
	chk := &api.AgentServiceCheck{
		HTTP: check["HTTP"],
		Interval: check["Interval"],
	}
	
	reg := api.AgentServiceRegistration{
		ID: service,
		Name: service,
		Port: port,
		Address: address,
		Tags: tags,
		Check: chk,
	}
	
	if err := client.Agent().ServiceRegister(&reg); err != nil {
		return err
	}
	
	return nil
}

func (c ConsulService) DeRegister(service string) (error) {
	client, _ := api.NewClient(api.DefaultConfig())
	return client.Agent().ServiceDeregister(service)
}

