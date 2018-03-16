# Snort Log enhancer 
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fjanritter%2Fsnort-log-enhancer.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fjanritter%2Fsnort-log-enhancer?ref=badge_shield)


This Tool enhances your default snort logfile's with geoip informationen.
For more information take a look at the example section.

## Build binaries
- Cd into the blocklog folder
```
cd src/blocklog 
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
- Start BlockLog (Linux command shown) 
```
./blocklog 
```
- Follow the displayed instructions - Info : The logfile should be in the same folder as the executable
- The final csv will be outputted in the same directory and is named block_log_enhanced.csv 

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



## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fjanritter%2Fsnort-log-enhancer.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fjanritter%2Fsnort-log-enhancer?ref=badge_large)