package cmd

import (
	"github.com/vietanhduong/resume/pkg/utils/output"
	"os"
)

var (
	varInitFunctions []func()
	cmdInitFunctions []func()
)

// RegisterCommandVar is used to register with px the initialization function
// for the command variable.
func RegisterCommandVar(c func()) bool {
	varInitFunctions = append(varInitFunctions, c)
	return true
}

// RegisterCommandInit is used to register with px the initialization function
// for the command flags.
func RegisterCommandInit(c func()) bool {
	cmdInitFunctions = append(cmdInitFunctions, c)
	return true
}

func init() {
	// Setup all variables
	for _, v := range varInitFunctions {
		v()
	}
	// Call all plugin inits
	for _, f := range cmdInitFunctions {
		f()
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		output.Eprintf("%v\n", err)
		os.Exit(1)
	}
}
