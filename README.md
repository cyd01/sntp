sntp
====

> Source: https://github.com/btfak/sntp

## A implementation of NTP server with Golang
### What this?
- Base on [RFC 2030](http://tools.ietf.org/html/rfc2030)
- Multiple equipments sync time from local
- Design for multiple equipments which can't connect to internet and need synchronization time
- Compatible with [ntpdate](http://www.eecis.udel.edu/~mills/ntp/html/ntpdate.html) service on the linux
- NTP client is fork from [beevik](https://github.com/beevik/ntp/)，a better client

## Manual
### Server
#### Build

    cd cmd/ntps
    go build

#### Usage

    $ ./ntps
    Init time is  1649683178
    Usage: ntp port [speed]
	    usual default port is 123

Speed parameter is used to add an increasing offset each second. `./ntps 123 2` means time passes twice as fast.

### Client
#### Build

    cd cmd/ntpc
    go build

#### Usage

    $ ./ntpc
    Usage: ADDRESS[:PORT] [ADDRESS[:PORT] ...]
	    default port is 123

#### Example

    $ ./ntpc 0.beevik-ntp.pool.ntp.org 0.fr.pool.ntp.org 0.europe.pool.ntp.org 0.uk.pool.ntp.org
    0.beevik-ntp.pool.ntp.org => 2022-04-11 15:26:47.838432902 +0200 CEST
    0.fr.pool.ntp.org => 2022-04-11 15:26:47.866064203 +0200 CEST
    0.europe.pool.ntp.org => 2022-04-11 15:26:47.902505517 +0200 CEST
    0.uk.pool.ntp.org => 2022-04-11 15:26:47.928748865 +0200 CEST
