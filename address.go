package web

import (
	"fmt"
	"net"
	"strings"
)

// Address encapsulates listening address of the http server. This
// structure implements the net.Addr interface.
type Address struct {
	network string
	address string
}

// NewAddress returns a server address by parsing the address in string form.
// The address string can be of following forms:
// * unix://var/run/server.sock
// * tcp://127.0.0.1:6666
// * 127.0.0.1:7777
// If the network is not specified in the string, it defaults to tcp and
// uses the rest of the string as the tcp address.
func NewAddress(address string) *Address {
	addr := &Address{}
	s := strings.Split(address, "://")

	if len(s) > 1 {
		addr.network = s[0]
		addr.address = s[1]
	} else {
		// Not a qualified address, default to tcp
		addr.network = "tcp"
		addr.address = address
	}

	return addr
}

// Network returns the network of this server address. For example: tcp, unix.
func (s *Address) Network() string {
	return s.network
}

// String returns address specific to the network. For example: 127.0.0.1:6666
// or /var/run/server.sock.
func (s *Address) String() string {
	return s.address
}

// Format returns a fully qualified address string as described in the NewAddress.
func (s *Address) Format() string {
	return fmt.Sprintf("%s://%s", s.network, s.address)
}

// MarshalText returns the address in text form.
func (s *Address) MarshalText() ([]byte, error) {
	return []byte(s.Format()), nil
}

// UnmarshalText parses the address in text form into an address.
func (s *Address) UnmarshalText(t []byte) error {
	n := NewAddress(string(t))
	s.network = n.network
	s.address = n.address

	return nil
}

func (s *Address) Listener() (net.Listener, error) {
	l, e := net.Listen(s.network, s.address)
	if e != nil {
		return nil, fmt.Errorf("net listen error: %w", e)
	}

	return l, nil
}
