package main

import (
	"bufio"
	"fmt"
	"github.com/gershwinlabs/gokml"
	"os"
	"strconv"
	"strings"
)

type Airport struct {
	City string
	Country string
	Lat float64
	Lon float64
}

func DMStoDD(parts []string) float64 {
	d, _ := strconv.ParseInt(parts[0], 10, 64)
	m, _ := strconv.ParseInt(parts[1], 10, 64)
	s, _ := strconv.ParseInt(parts[2], 10, 64)
	dir := parts[3]

	dd := float64(d) + ((1.0/60.0) * float64(m)) + ((1.0/(60.0 * 60.0)) * float64(s))

	if dir == "S" || dir == "U" { // U for West for some reason
		return -dd
	}

	return dd
}

func main() {
	tripsFileName := os.Args[1]
	airportsFileName := os.Args[2]

	tripsFile, err := os.Open(tripsFileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s\n", tripsFileName)
		os.Exit(1)
	}

	airportsFile, err := os.Open(airportsFileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s\n", airportsFileName)
		os.Exit(1)
	}

	defer tripsFile.Close()
	defer airportsFile.Close()

	airports := make(map[string]Airport)
	scanner := bufio.NewScanner(airportsFile)

	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, ":")

		if len(parts) < 13 {
			continue
		}

		code := parts[1]

		if code == "N/A" {
			continue
		}

		city := parts[3]
		country := parts[4]
		lat := DMStoDD(parts[5:9])
		lon := DMStoDD(parts[9:13])
		airports[code] = Airport{city, country, lat, lon}
	}

	scanner = bufio.NewScanner(tripsFile)

	for scanner.Scan() {
		s := scanner.Text()
		s = s[3:]
		parts := strings.Split(s, "#")
		s = parts[0]
		s = strings.Replace(s, "]", "", -1)
		codes := strings.Split(s, ",")

		for _, code := range codes {
			code := strings.TrimSpace(code)

			if a, found := airports[code]; found {
				fmt.Println(a.City, a.Lat, a.Lon)
			} else {
				fmt.Println(code, "not found")
			}
		}
	}

	kml := gokml.NewKML()
	fmt.Println(kml.Render())
}
