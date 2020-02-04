package netutil

import (
	"io"
	"net"

	"github.com/pkg/errors"
)

// IsConnectionError check if the error is a connection error (close)
func IsConnectionError(_err interface{}) bool {
	err, ok := _err.(error)
	if !ok {
		return false
	}

	err = errors.Cause(err)
	if err == io.EOF {
		return true
	}

	neterr, ok := err.(net.Error)
	if !ok {
		return false
	}
	if neterr.Timeout() {
		return false
	}

	return true
}

// ConnectTCP connects to host:port in TCP
func ConnectTCP(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	return conn, err
}
