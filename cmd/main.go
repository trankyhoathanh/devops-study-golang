// Package main - fasthttp server (fastest HTTP framework)
package main

import (
	"fmt"
	"log"

	"devops-study-golang/structure/app"
	"devops-study-golang/structure/router"

	"github.com/valyala/fasthttp"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal("Failed to initialize app:", err)
	}
	defer app.Close()

	port := string("8080")

	server := &fasthttp.Server{
		Handler:               router.New(app),
		Concurrency:           256 * 1024,
		ReadBufferSize:        4096,
		WriteBufferSize:       4096,
		MaxRequestBodySize:    1024 * 1024,
		DisableKeepalive:      false,
		TCPKeepalive:          true,
		ReduceMemoryUsage:     true,
		NoDefaultServerHeader: true,
		NoDefaultDate:         true,
		NoDefaultContentType:  true,
	}

	fmt.Printf("🚀 fasthttp server running")
	if err := server.ListenAndServe(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}