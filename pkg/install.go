package pkg

import (
	"fmt"
	"github.com/lu4p/binclude"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func PerformInstallation(config InstallConfig, BinFS *binclude.FileSystem) error {

	if *installDirectory == "" {
		log.Fatal("Please specify Installation directory")
	}

	if _, err := os.Stat(*installDirectory); !os.IsNotExist(err) {
		log.Fatal("Specified Installation directory already exists")
	}

	d := filepath.Join(*installDirectory, ".uninstall")
	err := os.MkdirAll(filepath.Join(d), os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create directory %s with error %v", d, err)
	}

	for _, c := range config.Components {
		d := filepath.Join(*installDirectory, c.Name, c.TargetDir)
		log.Printf("Creating Directory %s\n", d)
		err := os.MkdirAll(filepath.Join(d), os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to crate directory %s with error %v", d, err)
		}
		for _, f := range c.Files {
			if f.Platform == "" || runtime.GOOS == f.Platform {
				t := filepath.Join(d, f.Name)
				log.Printf("Installing %s\n", t)
				err := BinFS.CopyFile(fmt.Sprintf(f.SourceRelativePath), t)
				if err != nil {
					log.Printf("Failed to copy files to %s with error %v\n", t, err)
					return err
				}
			}
		}
	}
	return nil
}
