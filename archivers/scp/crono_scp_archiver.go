package main

import (
	"github.com/cronohub/sdk"
	"github.com/hashicorp/go-plugin"
)

// Archive is a concrete implementation of the archive plugin.
type Archive struct{}

// Execute is the entry point to this plugin.
func (Archive) Execute(filename string) (bool, error) {
	// ioutil.WriteFile("test.txt", []byte(filename), 0766)
	return true, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: sdk.Handshake,
		Plugins: map[string]plugin.Plugin{
			"crono_scp_provider": &sdk.ArchiveGRPCPlugin{Impl: &Archive{}},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
