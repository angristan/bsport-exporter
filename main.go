package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var (
	gauge    prometheus.Gauge
	token    string
	memberID string
)

func main() {
	flag.StringVar(&token, "token", "", "BSport API token")
	flag.StringVar(&memberID, "member", "", "BSport member ID")

	flag.Parse()

	if token == "" {
		logrus.Fatal("missing -token flag")
	}

	if memberID == "" {
		logrus.Fatal("missing -member flag")
	}

	gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "bsport_bookings_count",
		Help: "Number of bookings",
	})
	prometheus.MustRegister(gauge)

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Hour().Do(updateGauge)
	if err != nil {
		logrus.WithError(err).Fatal("Error scheduling job")
	}
	s.StartAsync()

	http.Handle("/metrics", promhttp.Handler())
	listenAddr := fmt.Sprintf("%s:%s", "0.0.0.0", "6677")
	logrus.Infof("Beginning to serve on %s", listenAddr)
	logrus.Fatal(http.ListenAndServe(listenAddr, nil))
}

func updateGauge() {
	count, err := getBookingsCount()
	if err != nil {
		logrus.WithError(err).Error("Error getting bookings count")
		return
	}

	gauge.Set(float64(count))
	logrus.WithField("bookings", count).Info("Updated gauge")
}
