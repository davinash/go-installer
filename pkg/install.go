package pkg

import (
	"fmt"
	"github.com/lu4p/binclude"
	"log"
	"os"
)

func createDirectories(config InstallConfig) error {
	if _, err := os.Stat(*installDirectory); !os.IsNotExist(err) {
		return fmt.Errorf(fmt.Sprintf("directory %s already exists", *installDirectory))
	}
	return nil
}

func PerformInstallation(config InstallConfig, fs *binclude.FileSystem) error {
	err := createDirectories(config)
	if err != nil {
		log.Printf("Failed to create required directories with error %v\n", err)
		return err
	}
	return nil
}
