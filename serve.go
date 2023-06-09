package plugin

import (
	"log"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

// ServeCommands should be called from your main function to create an RPC/GRPC
// server and serve plugins and commands to the CLI. Only go-based plugins support
// RPC, all other languages must use GRPC
func ServeCommands(version Version, protocol PluginProtocol, commands ...ExportCommand) {
	impl := NewExportPlugin(version, commands...)

	switch protocol {
	case GRPCProtocol:
		plugin.Serve(&plugin.ServeConfig{
			HandshakeConfig: plugin.HandshakeConfig{
				ProtocolVersion:  uint(GRPCProtocol),
				MagicCookieKey:   HandshakeConfig.MagicCookieKey,
				MagicCookieValue: HandshakeConfig.MagicCookieValue,
			},
			Plugins: plugin.PluginSet{
				"plugin": &GRPCExportPlugin{Impl: impl},
			},
			GRPCServer: plugin.DefaultGRPCServer,
		})
	case RPCProtocol:
		plugin.Serve(&plugin.ServeConfig{
			HandshakeConfig: plugin.HandshakeConfig{
				ProtocolVersion:  uint(RPCProtocol),
				MagicCookieKey:   HandshakeConfig.MagicCookieKey,
				MagicCookieValue: HandshakeConfig.MagicCookieValue,
			},
			Plugins: plugin.PluginSet{
				"plugin": &RPCExportPlugin{Impl: impl},
			},
			Logger: hclog.New(&hclog.LoggerOptions{
				Name:   "plugin",
				Output: os.Stderr,
				Level:  hclog.Debug,
			}),
		})
	default:
		log.Fatalf("unrecognized protocol %d", protocol)
	}
}
