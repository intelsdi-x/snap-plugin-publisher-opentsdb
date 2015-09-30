package main

import (
	"os"

	"github.com/intelsdi-x/pulse-plugin-publisher-opentsdb/opentsdb"
	"github.com/intelsdi-x/pulse/control/plugin"
)

func main() {
	meta := opentsdb.Meta()
	plugin.Start(meta, opentsdb.NewOpentsdbPublisher(), os.Args[1])
}
