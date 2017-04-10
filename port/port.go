// Inspired by https://github.com/phayes/freeport
package port

import (
	"fmt"
	"net"
	"strconv"
)

// TODO
func Get(ports ...interface{}) int {
	var (
		l   *net.TCPListener
		n   int
		err error
	)
	for _, port := range ports {
		switch port.(type) {
		case int:
			n = port.(int)
		case string:
			n, err = strconv.Atoi(port.(string))
			if err != nil {
				continue
			}
		default:
			continue
		}

		host := fmt.Sprintf("localhost:%d", n)
		addr, err := net.ResolveTCPAddr("tcp", host)
		if err != nil {
			continue
		}

		l, err = net.ListenTCP("tcp", addr)
		if err != nil {
			continue
		}
		defer l.Close()
	}
	if len(ports) == 0 {
		host := fmt.Sprintf("localhost:%d", 0)
		addr, err := net.ResolveTCPAddr("tcp", host)
		if err != nil {
			panic(err)
		}

		l, err = net.ListenTCP("tcp", addr)
		if err != nil {
			panic(err)
		}
		defer l.Close()
	}
	return l.Addr().(*net.TCPAddr).Port
}
