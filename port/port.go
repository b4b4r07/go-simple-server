// Inspired by https://github.com/phayes/freeport
package port

import (
	"fmt"
	"net"
)

const RANDOM = 0

func listen(port int) *net.TCPListener {
	host := fmt.Sprintf("localhost:%d", port)
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		panic(err)
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	return l
}

func Get(ports ...int) int {
	var l *net.TCPListener
	for _, port := range ports {
		if Available(port) {
			l = listen(port)
		}
	}
	if l == nil {
		l = listen(RANDOM)
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
