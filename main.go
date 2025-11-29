package main

import (
	"fmt"
	"github.com/Tinotsu/gator/internal/config"
)
func main () {
	cfg := config.NewConfig()
	fmt.Println(*cfg)
	cfg = config.Read(cfg)
	fmt.Println(*cfg)
	cfg.SetUSer("tino")
	cfg = config.Read(cfg)
	fmt.Println(*cfg)
}
