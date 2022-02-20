# bsport-exporter

An exporter for [bsport](bsport.io).

It is currently very basic and is supposed to be use to track how many bookings have been made in total:

```
➜  ~ curl -s localhost:6677/metrics | grep bsport
# HELP bsport_bookings_count Number of bookings
# TYPE bsport_bookings_count gauge
bsport_bookings_count 117
```

It uses the total number of bookings returned by the bsport API and does not parse the bookings themselves. Thus, the canceled bookings are also counted.

The gauge is updated every hour.

## Installation

Grab the latest binary from the releases, or build it yourself!

## Usage

Go on [backoffice.bsport.io](https://backoffice.bsport.io/) and inspect the network requests.

You can extract your member ID from the query string of some requests and your token from the authorization header of any request.

Then, lauch the exporter:

```sh
bsport-exporter -member xxxxxx -token xxxxxx
```

You should see something like:

```
INFO[0000] Beginning to serve on 0.0.0.0:6677
INFO[0000] Updated gauge                                 bookings=118
```
