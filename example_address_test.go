package web_test

import (
	"fmt"

	"gojini.dev/web"
)

func ExampleNewAddress() {
	addr := web.NewAddress("tcp://127.0.0.1:8888")

	fmt.Println("address format:", addr.Format())
	fmt.Println("address string:", addr.String())
	fmt.Println("network:", addr.Network())

	text, err := addr.MarshalText()
	if err != nil {
		panic(err)
	}

	if err = addr.UnmarshalText(text); err != nil {
		panic(err)
	}

	addr = web.NewAddress("127.0.0.1:8888")

	fmt.Println("address format:", addr.Format())
	fmt.Println("address string:", addr.String())
	fmt.Println("network:", addr.Network())

	text, err = addr.MarshalText()
	if err != nil {
		panic(err)
	}

	if err = addr.UnmarshalText(text); err != nil {
		panic(err)
	}

	addr = web.NewAddress("unix://var/run/server.sock")

	fmt.Println("address format:", addr.Format())
	fmt.Println("address string:", addr.String())
	fmt.Println("network:", addr.Network())

	text, err = addr.MarshalText()
	if err != nil {
		panic(err)
	}

	if err = addr.UnmarshalText(text); err != nil {
		panic(err)
	}

	// Output:
	// address format: tcp://127.0.0.1:8888
	// address string: 127.0.0.1:8888
	// network: tcp
	// address format: tcp://127.0.0.1:8888
	// address string: 127.0.0.1:8888
	// network: tcp
	// address format: unix://var/run/server.sock
	// address string: var/run/server.sock
	// network: unix
}
