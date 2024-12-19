package main

import (
	"github.com/st0pcha/go-fiber-example/internal/api"
	"github.com/st0pcha/go-fiber-example/internal/config"
)

func main() {
	cfg := config.Initialize()
	api.CreateAPI(cfg)
}
