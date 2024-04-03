package routes

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/andyrak/plex-dvr-hls/config"
	"github.com/gin-gonic/gin"
)

type DVR struct {
	FriendlyName    string `json:"FriendlyName"`
	ModelNumber     string `json:"ModelNumber"`
	FirmwareName    string `json:"FirmwareName"`
	TunerCount      int    `json:"TunerCount"`
	FirmwareVersion string `json:"FirmwareVersion"`
	DeviceID        string `json:"DeviceID"`
	DeviceAuth      string `json:"DeviceAuth"`
	BaseURL         string `json:"BaseURL"`
	LineupURL       string `json:"LineupURL"`
	Manufacturer    string `json:"Manufacturer"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Discover(c *gin.Context) {
	var deviceID = rand.Int63n(90000000-10000000) + 10000000

	var host = c.Request.Host

	var protocol = "https"
	if !config.Cfg.Https {
		protocol = "http"
	}

	c.JSON(
		http.StatusOK,
		DVR{
			FriendlyName:    config.Cfg.Name,
			ModelNumber:     "HDTC-2US",
			FirmwareName:    "hdhomeruntc_atsc",
			TunerCount:      len(config.Channels) * 3,
			FirmwareVersion: "20150826",
			DeviceID:        fmt.Sprintf("%d", deviceID),
			DeviceAuth:      "test1234",
			BaseURL:         fmt.Sprintf("%s://%s", protocol, host),
			LineupURL:       fmt.Sprintf("%s://%s/lineup.json", protocol, host),
			Manufacturer:    "Silicondust",
		},
	)
}
