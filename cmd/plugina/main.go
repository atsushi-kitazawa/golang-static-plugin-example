package main

import (
	"log"

	_ "github.com/atsushi-kitazawa/plugin-example/plugina"
	"github.com/atsushi-kitazawa/plugin-example/testplugin"
)

func main() {
	p, err := testplugin.GetPlugin("a")
	if err != nil {
		log.Fatalln(err)
	}
	p.Exec()
}
