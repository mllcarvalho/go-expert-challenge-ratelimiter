package main

import (
	"github.com/mllcarvalho/go-expert-challenge-ratelimiter/config"
	"github.com/mllcarvalho/go-expert-challenge-ratelimiter/internal/pkg/dependencyinjector"
)

func main() {
	configs, err := config.Load(".")
	if err != nil {
		panic(err)
	}

	di := dependencyinjector.NewDependencyInjector(configs)

	deps, err := di.Inject()
	if err != nil {
		panic(err)
	}

	deps.WebServer.Start()
}
