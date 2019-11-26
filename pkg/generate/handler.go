package generate

import (
	"fmt"
	"github.com/spf13/viper"
)

type (
	// Handler - exposed struct that implementd Handler interface
	Handler struct{}
)

// Handle - the function that will be called from the CLI with viper config
// to provide access to all flags
func (g *Handler) Handle(cnf *viper.Viper) error {
	fmt.Printf("Handler for command: generate\n")
	
	
	fmt.Printf("flag name: name, value: %s\n", cnf.GetString("name"))
	
	
	return nil
}