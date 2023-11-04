# gpx-timetagger
A Go program to add timestamps into a gpx file, based on a start date and time and an interval, provided as program arguments.

## Use case
Some GPS tracking devices or programs do not provide a timestamp for geo location tracks.
If the GPS recording is done in fixed intervals, like a point is recorded every second, this program allows you to add timestamps into the GPX file, according to a start date/time and a fixed time interval.

## Usage
Example:
```
go run . example.gpx "2006-01-02T15:04:05Z" 1s
```
This example above would, for each track point of your gpx file "example.gpx", add a timestamp, beginning at date and time Jan 2nd 2006 15:04:05 UTC, an each being increased by one second. The resulting gpx file will be a "example.gpx.converted.gpx" (the original source remains untouched).

The date string must be provided in this format. 

The time interval is provided as a valid "ParseDuration" string, as accepted by the Go time package, a possibly signed sequence of decimal numbers, each with optional fraction and a unit suffix, such as "300ms", "-1.5h" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h". 

## Requirements
This program needs the Go language installed on the local computer. See https://go.dev/doc/install on how to do that.
To use this program, clone the repository into your local filesystem, open a shell at theat location, and run the Go as indicated above.