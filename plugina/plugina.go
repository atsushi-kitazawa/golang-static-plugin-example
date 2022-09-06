package plugina

import (
	"fmt"

	"github.com/atsushi-kitazawa/plugin-example/testplugin"
)

type PluginA struct{}

func init() {
	testplugin.Register("a", PluginA{})
}

func (a PluginA) Exec() {
	fmt.Println("plugina exec()")
}
