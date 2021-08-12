package goWebhooks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// RETURNS NEW TIMESTAMP ACCORDING TO DISCORD'S FORMAT (ISO8601)
func GetTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05-0700")
}

// RETURNS COLOR IN DECIMAL VALUE
func GetColor(hexColor string) int {
	hexColor = strings.Replace(hexColor, "#", "", -1)
	decimalColor, err := strconv.ParseInt(hexColor, 16, 64)
	if err != nil {
		return 0
	}
	return int(decimalColor)
}

// SEND REQUEST
func SendWebhook(webookUrl string, webhook Webhook, retryOnRateLimit bool) error {
	if webhook.Content == "" && len(webhook.Embeds) == 0 {
		return errors.New("You must attach atleast one of these: Content; Embeds")
	}
	if len(webhook.Embeds) > 10 {
		return errors.New("Maximum number of embeds per webhook is 10")
	}
	jsonData, err := json.Marshal(webhook)
	if err != nil {
		return err
	}
	for {
		res, err := http.Post(webookUrl, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}
		switch res.StatusCode {
		case 204:
			res.Body.Close()
			return nil
		case 429:
			res.Body.Close()
			if !retryOnRateLimit {
				return errors.New("Webhook ratelimited")
			}
			timeout, err := strconv.Atoi(res.Header.Get("retry-after"))
			if err != nil {
				time.Sleep(5 * time.Second)
			} else {
				time.Sleep(time.Duration(timeout) * time.Millisecond)
			}
		default:
			res.Body.Close()
			return errors.New(fmt.Sprintf("Bad request (Status %d)", res.StatusCode))
		}
	}
}

type Webhook struct {
	Content   string  `json:"content"`
	Username  string  `json:"username"`
	AvatarUrl string  `json:"avatar_url"`
	Tts       bool    `json:"tts"`
	Embeds    []Embed `json:"embeds"`
}

type Embed struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Url         string         `json:"url"`
	Timestamp   string         `json:"timestamp"`
	Color       int            `json:"color"`
	Footer      EmbedFooter    `json:"footer"`
	Image       EmbedImage     `json:"image"`
	Thumbnail   EmbedThumbnail `json:"thumbnail"`
	Video       EmbedVideo     `json:"video"`
	Provider    EmbedProvider  `json:"provider"`
	Author      EmbedAuthor    `json:"author"`
	Fields      []EmbedFields  `json:"fields"`
}

type EmbedFooter struct {
	Text         string `json:"text"`
	IconUrl      string `json:"icon_url"`
	ProxyIconUrl string `json:"proxy_icon_url"`
}

type EmbedImage struct {
	Url      string `json:"url"`
	ProxyUrl string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedThumbnail struct {
	Url      string `json:"url"`
	ProxyUrl string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedVideo struct {
	Url      string `json:"url"`
	ProxyUrl string `json:"proxy_url"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type EmbedProvider struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type EmbedAuthor struct {
	Name         string `json:"name"`
	Url          string `json:"url"`
	IconUrl      string `json:"icon_url"`
	ProxyIconUrl string `json:"proxy_icon_url"`
}

type EmbedFields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
