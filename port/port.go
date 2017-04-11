// Inspired by https://github.com/phayes/freeport
package port

import (
	"fmt"
	"log"
	"net"
	"strconv"
)

// TODO
func Get(ports ...interface{}) int {
	var (
		l   *net.TCPListener
		n   int
		err error
		ok  bool
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
		ok = true
		if !Available(n) {
			ok = false
			continue
		}

		host := fmt.Sprintf("localhost:%d", n)
		addr, err := net.ResolveTCPAddr("tcp", host)
		if err != nil {
			log.Printf("skip %d ...", n)
			continue
		}

		l, err = net.ListenTCP("tcp", addr)
		if err != nil {
			log.Printf("skip %d ...", n)
			continue
		}
		defer l.Close()
	}
	if len(ports) == 0 || !ok {
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

func Available(port int) bool {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false
	}
	defer l.Close()
	return true
}
