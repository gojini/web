package web_test

import (
	"fmt"

	"gojini.dev/web"
)

func ExampleTLS12Config() {
	tls12Config := web.TLS12Config()
	fmt.Println("tls12 max version:", tls12Config.MaxVersion)
	fmt.Println("tls12 min version:", tls12Config.MinVersion)
	fmt.Println("prefer cipher suites:", tls12Config.PreferServerCipherSuites)

	// Output:
	// tls12 max version: 772
	// tls12 min version: 771
	// prefer cipher suites: true
}

func ExampleMutualTLS12Config() {
	certs, err := web.LoadCACerts("./test_certs/ca.crt")
	if err != nil {
		panic(err)
	}

	tls12Config := web.MutualTLS12Config(certs)
	fmt.Println("tls12 max version:", tls12Config.MaxVersion)
	fmt.Println("tls12 min version:", tls12Config.MinVersion)
	fmt.Println("prefer cipher suites:", tls12Config.PreferServerCipherSuites)
	fmt.Println("client auth:", tls12Config.ClientAuth)

	// Output:
	// tls12 max version: 772
	// tls12 min version: 771
	// prefer cipher suites: true
	// client auth: RequireAndVerifyClientCert
}

func ExampleTLS13Config() {
	tls13Config := web.TLS13Config()
	fmt.Println("tls13 max version:", tls13Config.MaxVersion)
	fmt.Println("tls13 min version:", tls13Config.MinVersion)
	fmt.Println("prefer cipher suites:", tls13Config.PreferServerCipherSuites)

	// Output:
	// tls13 max version: 772
	// tls13 min version: 772
	// prefer cipher suites: true
}

func ExampleMutualTLS13Config() {
	certs, err := web.LoadCACerts("./test_certs/ca.crt")
	if err != nil {
		panic(err)
	}

	tls13Config := web.MutualTLS13Config(certs)
	fmt.Println("tls13 max version:", tls13Config.MaxVersion)
	fmt.Println("tls13 min version:", tls13Config.MinVersion)
	fmt.Println("prefer cipher suites:", tls13Config.PreferServerCipherSuites)
	fmt.Println("client auth:", tls13Config.ClientAuth)

	// Output:
	// tls13 max version: 772
	// tls13 min version: 772
	// prefer cipher suites: true
	// client auth: RequireAndVerifyClientCert
}

func ExampleLoadCACerts() {
	_, err := web.LoadCACerts("./unknown_file")
	if err == nil {
		panic(err)
	}

	certs, err := web.LoadCACerts("./test_certs/ca.crt")
	if err != nil || certs == nil {
		panic(err)
	}

	// Output:
}
