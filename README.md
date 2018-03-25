# Snort Log enhancer 
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fjanritter%2Fsnort-log-enhancer.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fjanritter%2Fsnort-log-enhancer?ref=badge_shield)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/4936e18ecd084ed2840fe73fc8ab36e2)](https://www.codacy.com/app/jan-ritter/snort-log-enhancer?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=janritter/snort-log-enhancer&amp;utm_campaign=Badge_Grade)
[![Build Status](https://travis-ci.org/janritter/snort-log-enhancer.svg?branch=master)](https://travis-ci.org/janritter/snort-log-enhancer)

This Tool enhances your default snort logfile's with geoip informationen.
For more information take a look at the example section.

## Build binaries
- Cd into the logenhancer folder
```
cd logenhancer 
```
- Resolve packages
```
go get
```
- Build the executable
```
go build
```

## How to use

- Build the binaries or download prebuilt binaries from the release page 
- Start Logenhancer (Linux command shown) 
```
logenhancer 
```
- Follow the displayed instructions - Info : The logfile should be in the same folder as the executable
- The final csv will be outputted in the same directory and is named block_log_enhanced.csv for blockfiles or alert_log_enhanced.csv for alertfiles

## Example
#### BlockLog
- The default snort IP Blocklist logfile looks like this:
```
1.2.3.4
5.6.7.8
```
- BlockLog will create an enhanced version of the logfile, which will look like this: 
```
IP,Country,Latitude,Longitude
1.2.3.4,United States,47.91300,-122.30420
5.6.7.8,Germany,51.29930,9.49100
```

#### AlertLog
- The default snort Alert logfile looks like this:
```
03/25/18-13:53:01.858091 ,1,2402000,4755,"ET DROP Dshield Block Listed Source group 1",TCP,1.2.3.4,58529,4.5.6.7,12342,39436,Misc Attack,2
03/25/18-13:53:09.261608 ,1,2402000,4755,"ET DROP Dshield Block Listed Source group 1",TCP,2.3.4.5,58597,4.5.6.7,30489,26721,Misc Attack,2
```
- AlertLog will create an enhanced version of the logfile, which will look like this:
```
Date,GeneratorID,SnortID,RevisionNumber,Description,Protocol,SourceIP,SourcePort,SourceCountry,SourceLatitude,SourceLongitude,DestIP,DestPort,DestCountry,DestLatitude,DestLongitude,UNKNOWN,Class,Priority
03/25/18-13:53:01.858091 ,1,2402000,4755,ET DROP Dshield Block Listed Source group 1,TCP,1.2.3.4,58529,United States,47.91300,-122.30420,4.5.6.7,12342,United States,37.75100,-97.82200,39436,Misc Attack,2
03/25/18-13:53:09.261608 ,1,2402000,4755,ET DROP Dshield Block Listed Source group 1,TCP,2.3.4.5,58597,France,48.85820,2.33870,4.5.6.7,30489,United States,37.75100,-97.82200,26721,Misc Attack,2
```

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fjanritter%2Fsnort-log-enhancer.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fjanritter%2Fsnort-log-enhancer?ref=badge_large)