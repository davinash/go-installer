package pkg

import (
	"flag"
	"github.com/lu4p/binclude"
	"log"
)

var (
	mode             = flag.String("mode", "", "Pass to mode for action [install|uninstall|status|start|stop]")
	installDirectory = flag.String("install-dir", "", "Install directory")
)

func PerformAction(config InstallConfig, BinFS *binclude.FileSystem) {
	switch *mode {
	case "install":
		PerformInstallation(config, BinFS)
	case "uninstall":
		PerformUnInstallation(config)
	case "status":
		PerformStatusCheck()
	default:
		log.Println("No mode provided")
	}
}
