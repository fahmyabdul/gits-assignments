package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fahmyabdul/gits-assignments/test-3/config"

	"github.com/fahmyabdul/gits-assignments/test-3/internal/app"
)

const (
	serviceName string = "TEST-3-GRPC"
	appVersion  string = "v2023.01.15-1"
)

var (
	configPath string
	logPath    string
	getVersion bool
)

func init() {
	// Parse app flags
	flag.StringVar(&configPath, "config", "", "Set the Path of Configuration File (Optional)")
	flag.StringVar(&logPath, "log", "", "Set the Path & Filename of Log File (Optional)")
	flag.BoolVar(&getVersion, "version", false, "Show current app version")

	flag.Parse()
}

func main() {
	// Show app version if -version flag defined
	if getVersion {
		fmt.Println(appVersion)
		os.Exit(0)
	}

	// Load configuration
	cfg, err := config.InitConfig(configPath, serviceName)
	if err != nil {
		log.Println("| InitConfig |", err.Error())
		os.Exit(0)
	}

	// Run the app
	app.RunTest3Grpc(appVersion, cfg, logPath, serviceName)
}
