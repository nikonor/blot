package main

import (
	"fmt"

	"github.com/nikonor/blot/blot"
	"github.com/nikonor/blot/srv"
)

func main() {
	fmt.Printf("Start")
	defer func() {
		fmt.Printf("Finish")
	}()

	// TODO: старт ФП для рассылки

	// TODO: порт динамический
	b := blot.NewBlot()
	srv.NewSrv(b, ":8888")
}
