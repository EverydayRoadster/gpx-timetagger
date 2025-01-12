package main

import (
	"os"
	"fmt"
	"time"
	"path/filepath"
	"github.com/tkrajina/gpxgo/gpx"
)

func usage() {
	fmt.Println("Add timestamps into a GPX file based on start time and interval for each track point.")
	fmt.Println("Usage: go run . example.gpx 2023-11-09T15:16:20Z [1s]")
	os.Exit(0)
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){
	if len(os.Args) < 3 {
		usage()
	}

	// GPX input file
	filename := os.Args[1]
	payload,err := os.ReadFile(filename)
	check(err)

	// start time for timestamps
	startTime,err := time.Parse("2006-01-02T15:04:05Z", os.Args[2])
	check(err)

	// interval between track points, as duration
	interval := "1s"
	if len(os.Args) > 3 {
		interval = os.Args[3]
	} 	
	deltaDuration,err := time.ParseDuration(interval)
	check(err)

	// parse input from GPX format
	gpxFile,err := gpx.ParseBytes(payload)
	check(err)

	// start start time for output GPX
	gpxFile.Time = &startTime
	fmt.Println("GPX " + gpxFile.Version + " starting " + gpxFile.Time.String() + " with points every " + deltaDuration.String())
	// current time reflects accumulated intervals
	currentTime := startTime

	// for each track, segments inside track, all points inside each of the segments
	for trackIndex, _ := range gpxFile.Tracks {
		for segIndex, _ := range gpxFile.Tracks[trackIndex].Segments {
			for pointIndex, _ := range gpxFile.Tracks[trackIndex].Segments[segIndex].Points {
				// add the timestamp into point
				gpxFile.Tracks[trackIndex].Segments[segIndex].Points[pointIndex].Timestamp = currentTime
				// increase increment
				currentTime = currentTime.Add(deltaDuration)
			}
		}
	}

	// create output stream
	xmlBytes, err := gpxFile.ToXml(gpx.ToXmlParams{Version: "1.1", Indent: true})
	check(err)
	// write GPX XML output
	filename = filename[0:len(filename)-len(filepath.Ext(filename))]
	err = os.WriteFile(filename + ".timetagged.gpx", xmlBytes, 0666)
	check(err)
}