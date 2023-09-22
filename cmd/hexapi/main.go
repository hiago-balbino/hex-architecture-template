package main

import (
	"github.com/hiago-balbino/hex-architecture-template/internal/handlers"
)

func main() {
	handlers.NewServer().Start()
}
