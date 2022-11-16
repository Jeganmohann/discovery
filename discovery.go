package discovery

import (
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
)

type ServiceInfo struct {
	ServiceId      string // service id for registry find
	ServiceName    string // service name
	ServiceAddress string // server name or ip address
	Port           int    // port number of the service
}

func serviceDiscoveryWithConsul(sId string) (s *ServiceInfo) {
	config := consulapi.DefaultConfig()
	consul, error := consulapi.NewClient(config)
	if error != nil {
		fmt.Println(error)
	}
	services, error := consul.Agent().Services()
	if error != nil {
		fmt.Println(error)
	}

	service := services[sId]
	//address := service.Address
	//port := service.Port
	//url = fmt.Sprintf("http://%s:%v/helloworld", address, port)
	return &ServiceInfo{
		ServiceId:   sId,
		ServiceName: service.Address,
		Port:        service.Port,
	}
}

/*
func ServiceRegistryWithConsul(entry RegisterEntry) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Println(err)
	}

	//port, _ := strconv.Atoi(getPort()[1:len(getPort())])
	//address := getHostname()
	address := "127.0.0.1"

	registration := &consulapi.AgentServiceRegistration{
		ID:      entry.ServiceId,
		Name:    entry.ServiceName,
		Port:    entry.Port,
		Address: address,
		Check: &consulapi.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/check", address, entry.Port),
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	regiErr := consul.Agent().ServiceRegister(registration)

	if regiErr != nil {
		log.Printf("Failed to register service: %s %s:%v ", entry.ServiceId, address, entry.Port)
	} else {
		log.Printf("successfully register service: %s %s:%v", entry.ServiceId, address, entry.Port)
	}
}

func getHostname() (hostname string) {
	hostname, _ = os.Hostname()
	return
}
*/
