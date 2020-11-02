package common

import "net"

// GetIPs information
func GetIPs(IP string) (res []string) {
	if IP == "0.0.0.0" {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			return
		}
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok {
				if ipnet.IP.To4() != nil {
					res = append(res, ipnet.IP.String())
				}
			}
		}
	} else {
		res = append(res, IP)
	}

	return
}
