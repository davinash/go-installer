package main

//go:generate binclude

import (
	"encoding/json"
	"flag"
	"github.com/davinash/go-installer/pkg"
	"github.com/lu4p/binclude"
	"io/ioutil"
	"log"
)

var (
	mode             = flag.String("mode", "", "Pass to mode for action [install|uninstall|status|start|stop]")
	installDirectory = flag.String("install-dir", "", "Install directory")
)

func main() {
	flag.Parse()

	var config pkg.InstallConfig
	f, err := BinFS.Open("config.json")
	if err != nil {
		log.Println("config.json should exists")
		log.Fatalln(err)
	}

	out, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(out, &config)
	if err != nil {
		log.Fatalln(err)
	}

	binclude.IncludeFromFile("./package.txt")

	switch *mode {
	case "install":
		pkg.PerformInstallation(config, BinFS)
	case "uninstall":
		pkg.PerformUnInstallation(config)
	case "status":
		pkg.PerformStatusCheck()
	default:
		log.Println("No mode provided")
	}
}
