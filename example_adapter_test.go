package web_test

import (
	"fmt"
	"net/http"

	"gojini.dev/web"
)

func ExampleWrap() {
	god := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		fmt.Println("You reached God!")
	})

	charity := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Before
			fmt.Println("Donate and volunteer blood and sweat!")

			h.ServeHTTP(w, r)

			// After
			fmt.Println("Donate more and volunteer more!")
		})
	}

	prayer := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Before
			fmt.Println("Pray in the morning!")

			h.ServeHTTP(w, r)

			// After
			fmt.Println("Pray in the evening!")
		})
	}

	faith := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Before
			fmt.Println("Believe when things are good!")

			h.ServeHTTP(w, r)

			// After
			fmt.Println("Believe more when things are bad!")
		})
	}

	web.Wrap(god, charity, prayer, faith).ServeHTTP(nil, nil)

	// Output:
	// Donate and volunteer blood and sweat!
	// Pray in the morning!
	// Believe when things are good!
	// You reached God!
	// Believe more when things are bad!
	// Pray in the evening!
	// Donate more and volunteer more!
}
