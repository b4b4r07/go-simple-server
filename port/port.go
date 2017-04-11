// Inspired by https://github.com/phayes/freeport
package port

import (
	"fmt"
	"net"
	"strconv"
)

const ANY_PORT = 0

func listen(port int) *net.TCPListener {
	var l *net.TCPListener
	host := fmt.Sprintf("localhost:%d", port)
	addr, err := net.ResolveTCPAddr("tcp", host)
	if err != nil {
		panic(err)
	}

	l, err = net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	return l
}

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
		l = listen(n)
	}
	if len(ports) == 0 || !ok {
		l = listen(ANY_PORT)
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
