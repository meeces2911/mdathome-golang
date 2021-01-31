package main

import (
	"flag"

	"github.com/meeces2911/mdathome-golang/internal/mdathome"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func main() {
	// Define arguments
	printVersion := flag.Bool("version", false, "Prints version of client")
	shrinkDatabase := flag.Bool("shrink-database", false, "Shrink cache.db (may take a long time)")
	settingsLocation := flag.String("settings", "", "Relative path to settings.json")

	// Parse arguments
	flag.Parse()

	// Shrink database if flag given, otherwise start server
	if *printVersion {
		log.Infof("MD@Home Client %s (%d) written in Golang by @lflare", mdathome.ClientVersion, mdathome.ClientSpecification)
	} else if *shrinkDatabase {
		mdathome.ShrinkDatabase(*settingsLocation)
	} else {
		mdathome.StartServer(*settingsLocation)
	}
}
