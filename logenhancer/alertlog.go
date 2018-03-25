package main

import (
	"fmt"
	"os"
	"log"
	"encoding/csv"
	"bufio"
	"github.com/cheggaaa/pb"
	"io"
	"github.com/janritter/go-geo-ip/geoip"
	"reflect"
	"strconv"
)

type alertIP struct {
	Date string
	GeneratorID string
	SnortID string
	RevisionNumber string
	Description string
	Protocol string
	SourceIP string
	SourcePort string
	SourceCountry string
	SourceLatitude float64
	SourceLongitude float64
	DestIP string
	DestPort string
	DestCountry string
	DestLatitude float64
	DestLongitude float64
	UNKNOWN string
	Class string
	Priority string
}

func runAlertLog() {
	fmt.Println("Input alertfile filename: ")
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

	var alertIPs []alertIP

	for {
		line, readErr := reader.Read()
		if readErr == io.EOF {
			break
		} else if readErr != nil {
			log.Fatal(readErr)
		}

		sourceGeoIpData, errSrc := geoip.ForIP(line[6])
		destGeoIPData, errDest := geoip.ForIP(line[8])

		//Only save successful IPs
		if errSrc == nil && errDest == nil {
			alertIPs = append(alertIPs, alertIP{
				Date: line[0],
				GeneratorID: line[1],
				SnortID: line[2],
				RevisionNumber: line[3],
				Description: line[4],
				Protocol: line[5],
				SourceIP: line[6],
				SourcePort: line[7],
				SourceCountry: sourceGeoIpData.CountryName,
				SourceLatitude: sourceGeoIpData.Latitude,
				SourceLongitude: sourceGeoIpData.Longitude,
				DestIP: line[8],
				DestPort: line[9],
				DestCountry: destGeoIPData.CountryName,
				DestLatitude: destGeoIPData.Latitude,
				DestLongitude: destGeoIPData.Longitude,
				UNKNOWN: line[10],
				Class: line[11],
				Priority: line[12],
			})
		}

		bar.Increment()
	}

	//Write struct array to csv
	resultFile, err := os.Create("alert_log_enhanced.csv")
	defer resultFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	csvWriter := csv.NewWriter(resultFile)
	defer csvWriter.Flush()

	//Write header
	var values[] string
	values = append(values, "Date")
	values = append(values, "GeneratorID")
	values = append(values, "SnortID")
	values = append(values, "RevisionNumber")
	values = append(values, "Description")
	values = append(values, "Protocol")
	values = append(values, "SourceIP")
	values = append(values, "SourcePort")
	values = append(values, "SourceCountry")
	values = append(values, "SourceLatitude")
	values = append(values, "SourceLongitude")
	values = append(values, "DestIP")
	values = append(values, "DestPort")
	values = append(values, "DestCountry")
	values = append(values, "DestLatitude")
	values = append(values, "DestLongitude")
	values = append(values, "UNKNOWN")
	values = append(values, "Class")
	values = append(values, "Priority")
	if err := csvWriter.Write(values); err != nil {
		log.Fatal(err)
	}

	for i:=0; i<len(alertIPs);i++  {
		singleStruct := alertIPs[i]
		reflectionStruct := reflect.ValueOf(&singleStruct).Elem()
		var values [] string
		for a:=0; a<reflectionStruct.NumField(); a++ {
			f := reflectionStruct.Field(a)
			if f.Type().String() == "float64" {
				values = append(values, strconv.FormatFloat(f.Interface().(float64), 'f', 5, 64))
			} else {
				values = append(values, f.Interface().(string))
			}
		}
		if err := csvWriter.Write(values); err != nil {
			log.Fatal(err)
		}
	}

	bar.FinishPrint("Finished enhancing!")
}
