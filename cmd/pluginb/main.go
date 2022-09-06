package main

import (
	"log"

	_ "github.com/atsushi-kitazawa/plugin-example/pluginb"
	"github.com/atsushi-kitazawa/plugin-example/testplugin"
)

func main() {
	p, err := testplugin.GetPlugin("b")
	if err != nil {
		log.Fatalln(err)
	}

	p.Exec()
}
