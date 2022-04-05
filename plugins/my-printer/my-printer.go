package main

import (
	"fmt"

	"github.com/laurentiuNiculae/go-dynamic-lib/extensions"
)

type MyPrinterImplementation struct {
	Config extensions.Config
}

func Register(pm *extensions.PluginManager) {
	pm.RegisterImplementation("Printer", "MyPrinterImplementation",
		MyPrinterImplementation{
			Config: pm.Config(),
		})
}

func (mpi MyPrinterImplementation) Print(s string) {
	if mpi.Config.Verbose {
		fmt.Printf("GoVersion: %s", mpi.Config.GoVersion)
	}

	fmt.Printf("I am printing: %s\n", s)
}
