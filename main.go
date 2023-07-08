package main

import (
	"fmt"

	"github.com/nikonor/blot/blot"
	"github.com/nikonor/blot/repo"
	"github.com/nikonor/blot/sender"
	"github.com/nikonor/blot/srv"
)

//go:generate swag init --parseDependency
//go:generate rm -f docs/swagger.json docs/docs.go
func main() {
	fmt.Printf("Start")
	defer func() {
		fmt.Printf("Finish")
	}()

	// старт ФП для рассылки
	s := sender.NewSender()

	// TODO: динамический файл
	r, err := repo.NewRepo("data/data.sqlite3")
	if err != nil {
		panic(err.Error())
	}

	b := blot.NewBlot(r, s)

	// TODO: порт динамический
	srv.NewSrv(b, ":8888")
}
