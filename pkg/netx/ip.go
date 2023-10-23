package netx

import (
	"net"
)

func GetIPv4s() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	ips := make([]string, 0)
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok {
			if ipnet.IP.To4() == nil {
				continue
			}
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}
