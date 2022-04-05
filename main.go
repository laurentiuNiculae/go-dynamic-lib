package main

import "github.com/laurentiuNiculae/go-dynamic-lib/extensions"

func main() {
	pm := extensions.NewPluginManager(extensions.Config{
		Verbose:   true,
		GoVersion: "1.17.8",
	})

	pm.LoadPlugins("/plugins")
}
