package ship

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

var opts Option

func It() {

	d, i := parse()

	dest, ok := findDestination(d)
	if !ok {
		fmt.Printf("Not defined destination. [%s]\n", d)
		os.Exit(1)
	}

	invoice, _ := findInvoice(i)
	if !ok {
		fmt.Printf("Not defined invoice. [%s]\n", i)
		os.Exit(1)
	}

	invoice.sendTo(dest)
}

func parse() (dest, invoice string) {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "ship-it"
	parser.Usage = "[OPTIONS] DESTINATION INVOICE"

	args, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	if len(args) < 2 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	return args[0], args[1]
}
