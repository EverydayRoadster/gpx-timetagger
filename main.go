package main

import (
	"os"
	"fmt"
	"time"
	"github.com/tkrajina/gpxgo/gpx"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main(){
	filename := "example.gpx"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	} 
	payload,err := os.ReadFile(filename)
	check(err)

	startTime := time.Now()
	if len(os.Args) > 2 {
		startTime,err = time.Parse("2006-01-02T15:04:05Z", os.Args[2])
		check(err)
	} 

	interval := "2s"
	if len(os.Args) > 3 {
		interval = os.Args[3]
	} 	
	deltaDuration,err := time.ParseDuration(interval)
	check(err)

	gpxFile,err := gpx.ParseBytes(payload)
	check(err)

	gpxFile.Time = &startTime
	fmt.Println("GPX " + gpxFile.Version + " " + gpxFile.Time.String() + " every " + deltaDuration.String())

	currentTime := startTime

	for trackIndex, _ := range gpxFile.Tracks {
		for segIndex, _ := range gpxFile.Tracks[trackIndex].Segments {
			for pointIndex, _ := range gpxFile.Tracks[trackIndex].Segments[segIndex].Points {
				gpxFile.Tracks[trackIndex].Segments[segIndex].Points[pointIndex].Timestamp = currentTime
				currentTime = currentTime.Add(deltaDuration)
			}
		}
	}

	xmlBytes, err := gpxFile.ToXml(gpx.ToXmlParams{Version: "1.1", Indent: true})
	check(err)
	err = os.WriteFile(filename + ".converted.gpx", xmlBytes, 0666)
	check(err)

}