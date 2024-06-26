package routes

import (
	"bytes"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/andyrak/plex-dvr-hls/config"
	"github.com/gin-gonic/gin"
)

type ChannelSimplified struct {
	ID   int
	Name string
	Logo string
}

type Programme struct {
	HourStr       string
	DateTimeStart string
	DateTimeEnd   string
}

func XMLTV(c *gin.Context) {
	var channels []ChannelSimplified

	var alteredIndex = config.Cfg.ChannelStart

	for _, channel := range config.Channels {
		channels = append(
			channels,
			ChannelSimplified{
				ID:   alteredIndex,
				Name: channel.Name,
				Logo: channel.Logo,
			},
		)
		alteredIndex++
	}

	var programmes []Programme
	var now = time.Now()

	for i := 0; i < 24; i++ {
		var start = time.Date(now.Year(), now.Month(), now.Day(), i+1, 0, 0, 0, now.Location())
		var end = time.Date(now.Year(), now.Month(), now.Day(), i+1, 59, 59, 999, now.Location())
		var dateTimeStart = start.Format("20060102150405 -0700")

		var dateTimeEnd = end.Format("20060102150405 -0700")

		var hourStr = start.Format("3PM")

		programmes = append(
			programmes,
			Programme{
				HourStr:       hourStr,
				DateTimeStart: dateTimeStart,
				DateTimeEnd:   dateTimeEnd,
			},
		)
	}

	t := template.Must(template.New("xmltv.tmpl").ParseFiles("templates/xmltv.tmpl"))

	var b bytes.Buffer
	var err = t.Execute(
		&b,
		gin.H{
			"channels":   channels,
			"programmes": programmes,
		},
	)

	if err != nil {
		log.Println(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Data(http.StatusOK, "application/xml", b.Bytes())
}
