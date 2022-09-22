//+build mage

package main

import (
	// mage:import
        build "github.com/andersonz1/grafana-plugin-sdk-go/build"
)

// Default configures the default target.
var Default = build.BuildAll
