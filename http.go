package main

import (
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

var g errgroup.Group

// Run 启动 HTTP 服务.
func Run(server *http.Server) {
	// run server
	g.Go(func() error {
		return server.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
