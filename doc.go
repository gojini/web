// Copyright 2022 Cisco Systems, Inc. All rights reserved.

/*
Package web provides structs and functions to support configuration and building
of an HTTPS web server via TLS 1.2 or TLS 1.3.

TLS

To encrypt data across network communications, this package supports the use of
Transport Layer Security (TLS, formerly SSL). Specifically, versions 1.2 and 1.3
are supported. Using types from crypto/tls and crypto/x509, this package
implements functions to configure TLS and mTLS, look up the available cipher
suites, and load CA certificates from a given file to construct a CA pool.

Address

Pertaining to network addresses, this package provides helpful types and
functions to store information (network and address strings) and manipulate
information (transform from text form of address to ServerAddress form and vice
versa).

Adapter

An adapter "adapts" an http.Handler, returning a wrapped http.Handler. This
package provides the type definition for Adapter and function definition for
Wrap() to wrap such an http.Handler given a slice of Adapter objects.
*/
package web
