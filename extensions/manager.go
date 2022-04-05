package extensions

import (
	"fmt"
	"os"
	"path/filepath"
	"plugin"
)

const (
	pluginsDir = "plugins"
)

type PluginManager struct {
	implementations map[string]map[string]interface{}
	config          Config
}

var implementations map[string]map[string]interface{} = map[string]map[string]interface{}{}

func NewPluginManager(config Config) PluginManager {
	return PluginManager{
		implementations: implementations,
		config:          config,
	}
}

func (pm *PluginManager) RegisterImplementation(extPoint string, implName string, plugin interface{}) {
	// do verifications
	pm.implementations[extPoint][implName] = plugin
}

func (pm *PluginManager) Config() Config {
	return pm.config
}

func (pm *PluginManager) LoadPlugins(pluginDir string) {
	dirs, err := os.ReadDir(pluginsDir)
	if err != nil {
		panic("I cry, no pluginDir.")
	}

	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}

		plugin, err := plugin.Open(filepath.Join(pluginsDir, d.Name(), d.Name()+".so"))
		if err != nil {
			fmt.Println("Cryy!!", err)
			continue
		}

		f, err := plugin.Lookup("Register")
		if err != nil {
			fmt.Println("Cryy!! Can't find register :(")
			continue
		}
		registerFunction := f.(func(*PluginManager))
		registerFunction(pm)
	}
}
