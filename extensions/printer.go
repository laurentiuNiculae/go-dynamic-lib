package extensions

import "fmt"

type Printer interface {
	Print(s string)
}

var printerImplementations map[string]interface{} = map[string]interface{}{}

func init() {
	implementations["Printer"] = printerImplementations
}

func GetPrinters() []Printer {
	var printers []Printer

	for k, v := range printerImplementations {
		fmt.Println(k)
		printers = append(printers, v.(Printer))
	}

	return printers
}
