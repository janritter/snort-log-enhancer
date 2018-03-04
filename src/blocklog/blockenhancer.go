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
	"bytes"
)

type BlockedIP struct {
	IP string
	Country string
}

func main() {
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

	var blockedIps []BlockedIP

	for {
		line, readErr := reader.Read()
		if readErr == io.EOF {
			break
		} else if readErr != nil {
			log.Fatal(readErr)
		}

		blockedIps = append(blockedIps, BlockedIP{
			IP: line[0],
			Country: geoip.ForIp(line[0]).CountryName,
		})

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

	for i:=0; i<len(blockedIps);i++  {
		singleStruct := blockedIps[i]
		var values [] string
		values = append(values, singleStruct.IP)
		values = append(values, singleStruct.Country)
		if err := csvWriter.Write(values); err != nil {
			log.Fatal(err)
		}
	}

	bar.FinishPrint("Finished enhancing!")
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
