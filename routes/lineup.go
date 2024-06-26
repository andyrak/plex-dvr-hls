package routes

import (
	"fmt"
	"net/http"

	"github.com/andyrak/plex-dvr-hls/config"
	"github.com/gin-gonic/gin"
)

type ChannelLineup struct {
	GuideNumber string   `json:"GuideNumber"`
	GuideName   string   `json:"GuideName"`
	Tags        []string `json:"Tags"`
	URL         string   `json:"URL"`
}

func Lineup(c *gin.Context) {
	var channelLineups []ChannelLineup

	var host = c.Request.Host

	var protocol = "https"
	if !config.Cfg.Https {
		protocol = "http"
	}

	var alteredIndex = config.Cfg.ChannelStart

	for _, channel := range config.Channels {
		channelLineups = append(
			channelLineups,
			ChannelLineup{
				GuideNumber: fmt.Sprintf("%d", alteredIndex),
				GuideName:   channel.Name,
				Tags:        make([]string, 0),
				URL:         fmt.Sprintf("%s://%s/stream/%d", protocol, host, alteredIndex),
			},
		)
		alteredIndex++
	}

	c.JSON(
		http.StatusOK,
		channelLineups,
	)
}

type Status struct {
	ScanInProgress int      `json:"ScanInProgress"`
	ScanPossible   int      `json:"ScanPossible"`
	Source         string   `json:"Source"`
	SourceList     []string `json:"Cable"`
}

func LineupStatus(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		Status{
			ScanInProgress: 0,
			ScanPossible:   1,
			Source:         "Cable",
			SourceList: []string{
				"Cable",
			},
		},
	)

}
