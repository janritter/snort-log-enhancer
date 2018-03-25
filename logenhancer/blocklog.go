package main

import (
	"fmt"
	"os"
	"log"
	"encoding/csv"
	"bufio"
	"io"
	"github.com/janritter/go-geo-ip/geoip"
	"github.com/cheggaaa/pb"
	"strconv"
)

type blockedIP struct {
	IP string
	Country string
	Latitude float64
	Longitude float64
}

func runBlockLog() {
	fmt.Println("Input blockfile filename: ")
	filename := ""
	fmt.Scanf("%s", &filename)

	csvFile, err := os.Open(filename)
	defer csvFile.Close()

	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	//Start new Progressbar
	count , _ := lineCounter(csvFile)
	bar := pb.StartNew(count)

	//Reset line pointer after line counting
	csvFile.Seek(0, 0) // reset/rewind back to offset and whence 0 0

	var blockedIps []blockedIP

	for {
		line, readErr := reader.Read()
		if readErr == io.EOF {
			break
		} else if readErr != nil {
			log.Fatal(readErr)
		}

		geoIpData, err := geoip.ForIP(line[0])

		//Only save successful IPs
		if err == nil {
			blockedIps = append(blockedIps, blockedIP{
				IP:        line[0],
				Country:   geoIpData.CountryName,
				Latitude:  geoIpData.Latitude,
				Longitude: geoIpData.Longitude,
			})
		}

		bar.Increment()
	}

	//Write struct array to csv
	resultFile, err := os.Create("block_log_enhanced.csv")
	defer resultFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	csvWriter := csv.NewWriter(resultFile)
	defer csvWriter.Flush()

	//Write header
	var values[] string
	values = append(values, "IP")
	values = append(values, "Country")
	values = append(values, "Latitude")
	values = append(values, "Longitude")
	if err := csvWriter.Write(values); err != nil {
		log.Fatal(err)
	}

	for i:=0; i<len(blockedIps);i++  {
		singleStruct := blockedIps[i]
		var values [] string
		values = append(values, singleStruct.IP)
		values = append(values, singleStruct.Country)
		values = append(values, strconv.FormatFloat(singleStruct.Latitude, 'f', 5, 64))
		values = append(values, strconv.FormatFloat(singleStruct.Longitude, 'f', 5, 64))
		if err := csvWriter.Write(values); err != nil {
			log.Fatal(err)
		}
	}

	bar.FinishPrint("Finished enhancing!")
}
