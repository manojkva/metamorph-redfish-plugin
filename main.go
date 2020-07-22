package main

import (
	config "github.com/bm-metamorph/MetaMorph/pkg/config"
	"encoding/json"
	"fmt"
	"github.com/bm-metamorph/MetaMorph/pkg/db/models/node"
	"github.com/hashicorp/go-plugin"
	"github.com/manojkva/metamorph-plugin/plugins/redfish"
	driver "github.com/manojkva/metamorph-redfish-plugin/pkg/redfish"
	"os"
)

func main() {
	config.SetLoggerConfig("logger.pluginpath")
	if len(os.Args) != 2 {
		fmt.Println("Usage go-redfish-plugin <uuid>")
		os.Exit(1)
	}
	uuid := os.Args[1]


	var bmhnode driver.BMHNode

	old  :=  os.Stdout

	//discard stdout as any unfamilar output ends up with error on Server side.

	os.Stdout,_ = os.Open(os.DevNull)

	data, err := node.Describe(uuid)

	if err == nil {

		err = json.Unmarshal(data, &bmhnode)
	}
	//revert std output
	os.Stdout = old
	if err != nil {

		fmt.Printf("Failed to locate node in DB for uuid %v\n", uuid)
		os.Exit(1)
	}
	//Get node details from db
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: redfish.Handshake,
		Plugins: map[string]plugin.Plugin{
			"redfish": &redfish.RedfishPlugin{Impl: &bmhnode}},
		GRPCServer: plugin.DefaultGRPCServer})
}
