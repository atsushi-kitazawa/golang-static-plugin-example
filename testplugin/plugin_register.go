package testplugin

import "errors"

type Plugin interface {
	Exec()
}

var plugins = map[string]Plugin{}

func Register(name string, p Plugin) {
	plugins[name] = p
}

func Plugins() []Plugin {
	var r []Plugin
	for _, p := range plugins {
		r = append(r, p)
	}
	return r
}

func GetPlugin(name string) (Plugin, error) {
	if p, ok := plugins[name]; ok {
		return p, nil
	}
	return nil, errors.New("not found plugin " + name)
}
