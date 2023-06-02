package main

import (
	"log"
	"os"

	plugin "github.com/gideaworx/terraform-exporter-plugin-go"
	"github.com/gideaworx/terraform-exporter-plugin-go/example/shared"
)

var Version = plugin.Version{
	Major: 0,
	Minor: 0,
	Patch: 1,
}

func main() {
	logf, err := os.OpenFile("output.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}
	defer logf.Close()

	logger := log.New(logf, "[RAW-FILES] ", log.LstdFlags)

	plugin.ServeCommands(Version, plugin.RPCProtocol, shared.NewLocalRawFilesCommand(logger))
}
