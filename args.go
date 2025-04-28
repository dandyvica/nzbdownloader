package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const CONFIG_FILE = "nzbgo.yml"

// All settings
type CliOpts struct {
	check        bool          // if set, check connexion to server
	settingsFile string        // name of the YAML file where configuration is defined
	logFile      *os.File      // handle on log file
	interActive  bool          // true if we want a kind of REPL
	settings     *YAMLSettings // all settings as returned when reading YAML file
}

// This will match the YAML configuration file where all settings are defined
type YAMLSettings struct {
	Server struct {
		Name     string
		Port     uint16
		Userid   string
		Password string
		Ssl      bool
	}
	// log file specs
	Log struct {
		Name string
	}
}

func CliArgs() *CliOpts {
	// init struct
	var opts CliOpts

	flag.StringVar(&opts.settingsFile, "settings", CONFIG_FILE, "path of the settings file")
	flag.BoolVar(&opts.check, "check", false, "if set, check connexion to server")
	flag.BoolVar(&opts.interActive, "i", false, "if set, you can enter interactive commands to send")
	flag.Parse()

	// read YAML file to get settings
	opts.settings = readYAMLConfig(opts.settingsFile)

	// establish log
	opts.logFile = createLogFile(opts.settings.Log.Name)

	return &opts
}

// Read the YAML configuration file
func readYAMLConfig(configFile string) *YAMLSettings {
	// read whole file in memory
	yamlData, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("error <%v> opening YAML configuration file: <%s>", err, configFile)
	}

	// read YAML into struct
	yamlConf := &YAMLSettings{}
	err = yaml.Unmarshal(yamlData, yamlConf)
	if err != nil {
		log.Fatalf("error <%v> reading YAML configuration file: <%s>", err, configFile)
	}

	return yamlConf
}

// create or open log file
func createLogFile(logName string) *os.File {
	// Open logfile
	logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening logfile:", err)
		os.Exit(1)
	}

	// Set log output
	log.SetOutput(logFile)

	// Optional: add timestamp in logs
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return logFile
}
