package main

import (
	"flag"
	"fmt"
	"fresh/runner"
	"os"
	"path/filepath"
)

func main() {
	//configPath := flag.String("c", "", "config file path")
	//configFile := os.Getenv("GOPATH") + "\\fresh.conf"
	configFile := filepath.Join(os.Getenv("GOPATH"), "fresh.conf")
	configPath := flag.String("c", configFile, "config file path")
	flag.Parse()
	//fmt.Println(configFile)

	if *configPath != "" {
		if _, err := os.Stat(*configPath); err != nil {
			fmt.Printf("Can't find config file `%s`\n", *configPath)
			os.Exit(1)
		} else {
			os.Setenv("RUNNER_CONFIG_PATH", *configPath)
		}
	}
	runner.Start()
}
