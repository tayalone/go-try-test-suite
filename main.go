package main

import "go-try-test-suite/router"

func main() {
	r := router.SetUpRouter()

	r.Run(":3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
