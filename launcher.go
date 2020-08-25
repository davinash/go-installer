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

	binclude.IncludeFromFile("./file.list")
	binclude.Include("./config.json")

	pkg.PerformAction(config, BinFS)

}
