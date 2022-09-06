package pluginb

import (
	"fmt"

	"github.com/atsushi-kitazawa/plugin-example/testplugin"
)

type PluginB struct{}

func init() {
	testplugin.Register("b", PluginB{})
}

func (a PluginB) Exec() {
	fmt.Println("pluginb exec()")
}
