package extensions

type PluginManager struct {
	implementations map[string]map[string]interface{}
	config          Config
}

var implementations map[string]map[string]interface{}

func NewPluginManager(config Config) PluginManager {
	return PluginManager{
		implementations: implementations,
		config:          config,
	}
}

func (pm *PluginManager) RegisterImplementation(extPoint string, implName string, plugin interface{}) {
	// do tests
	pm.implementations[extPoint][implName] = plugin
}

func (pm *PluginManager) Config() Config {
	return pm.config
}

func (pm *PluginManager) LoadPlugins(pluginDir string) {
	
}
