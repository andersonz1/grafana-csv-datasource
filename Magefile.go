//+build mage

package main

import (
	// mage:import
	build "/home/andersonz/grafana-plugin-sdk-go-0.124.0/build"
)

// Default configures the default target.
var Default = build.BuildAll
