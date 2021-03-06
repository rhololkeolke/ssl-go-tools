package main

import (
	"flag"
	"github.com/RoboCup-SSL/ssl-go-tools/pkg/persistence"
	"log"
)

var addressVisionLegacy = flag.String("vision-legacy-address", "224.5.23.2:10005", "Multicast address for vision 2010 (legacy)")
var addressVision = flag.String("vision-address", "224.5.23.2:10006", "Multicast address for vision 2014")
var addressReferee = flag.String("referee-address", "224.5.23.1:10003", "Multicast address for referee 2013")

var visionLegacyEnabled = flag.Bool("vision-legacy-enabled", true, "Record legacy vision packages")
var visionEnabled = flag.Bool("vision-enabled", true, "Record vision packages")
var refereeEnabled = flag.Bool("referee-enabled", true, "Record referee packages")

var VisionLegacyType = persistence.MessageType{Id: persistence.MessageSslVision2010, Name: "vision-legacy"}
var VisionType = persistence.MessageType{Id: persistence.MessageSslVision2014, Name: "vision"}
var RefereeType = persistence.MessageType{Id: persistence.MessageSslRefbox2013, Name: "referee"}

var logFile = flag.String("file", "", "The log file to play")
var skipNonRunningStages = flag.Bool("skip", false, "Skip frames while not in a running stage")

func main() {
	flag.Parse()

	if *logFile == "" {
		log.Fatal("Missing logfile")
	}

	broadcaster := persistence.NewBroadcaster()
	broadcaster.SkipNonRunningStages = *skipNonRunningStages
	addSlots(&broadcaster)

	defer broadcaster.Stop()
	if err := broadcaster.Start(*logFile); err != nil {
		log.Fatal(err)
	}
}

func addSlots(logger *persistence.Broadcaster) {
	if *visionLegacyEnabled {
		logger.AddSlot(VisionLegacyType, *addressVisionLegacy)
	}
	if *visionEnabled {
		logger.AddSlot(VisionType, *addressVision)
	}
	if *refereeEnabled {
		logger.AddSlot(RefereeType, *addressReferee)
	}
}
