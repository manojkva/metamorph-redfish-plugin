package main

import (
	config "github.com/manojkva/metamorph-plugin/pkg/config"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-plugin"
	"github.com/manojkva/metamorph-plugin/common/bmh"
	driver "github.com/manojkva/metamorph-redfish-plugin/pkg/redfish"
	"os"
	"encoding/base64"
)

func main() {
	config.SetLoggerConfig("logger.plugins.redfishpluginpath")
	if len(os.Args) != 2 {
		fmt.Println("Usage metamorph-redfish-plugin <inputConfig>")
		os.Exit(1)
	}
	data := os.Args[1]


	var bmhnode driver.BMHNode

        inputConfig,err :=  base64.StdEncoding.DecodeString(data)

	if  err != nil {
		fmt.Printf("Failed to decode input config %v\n", data)
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}

	err = json.Unmarshal([]byte(inputConfig), &bmhnode)

	if err != nil {

		fmt.Printf("Failed to decode input config %v\n", inputConfig)
		fmt.Printf("Error %v\n", err)
		os.Exit(1)
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: bmh.Handshake,
		Plugins: map[string]plugin.Plugin{
			"metamorph-redfish-plugin": &bmh.BmhPlugin{Impl: &bmhnode}},
		GRPCServer: plugin.DefaultGRPCServer})
}
