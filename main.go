package main

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

// Returns a timestamp for the embed footer according to ISO8601 format (required)
func GetTimestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05-0700")
}

// Transforms hex color code to decimal (required)
func GetColor(hexColor string) int {
	hexColor = strings.Replace(hexColor, "#", "", -1)
	decimalColor, err := strconv.ParseInt(hexColor, 16, 64)
	if err != nil {
		return 0
	}
	return int(decimalColor)
}

// Execute the webhook request
func SendWebhook(webookUrl string, content Webhook, retryOnRateLimit bool) error {
	if content.Content == "" && len(content.Embeds) == 0 {
		return errors.New("you must attach atleast one of these: content; embeds")
	}
	if len(content.Embeds) > 10 {
		return errors.New("maximum number of embeds per webhook is 10")
	}
	jsonData, err := json.Marshal(content)
	if err != nil {
		return err
	}
	for {
		res, err := http.Post(webookUrl, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			return err
		}
		defer res.Body.Close()
		switch res.StatusCode {
		case 204:
			return nil
		case 429:
			if !retryOnRateLimit {
				return errors.New("webhook rate limited")
			}
			timeout, err := strconv.Atoi(res.Header.Get("retry-after"))
			if err == nil {
				time.Sleep(time.Duration(timeout) * time.Millisecond)
			} else {
				time.Sleep(5 * time.Second)
			}
		default:
			return fmt.Errorf("bad request (status code %d)", res.StatusCode)
		}
	}
}

type Webhook struct {
	Content   string  `json:"content,omitempty"`
	Username  string  `json:"username,omitempty"`
	AvatarUrl string  `json:"avatar_url,omitempty"`
	Tts       bool    `json:"tts,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
}

type Embed struct {
	Title       string         `json:"title,omitempty"`
	Description string         `json:"description,omitempty"`
	Url         string         `json:"url,omitempty"`
	Timestamp   string         `json:"timestamp,omitempty"`
	Color       int            `json:"color,omitempty"`
	Footer      EmbedFooter    `json:"footer,omitempty"`
	Image       EmbedImage     `json:"image,omitempty"`
	Thumbnail   EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       EmbedVideo     `json:"video,omitempty"`
	Provider    EmbedProvider  `json:"provider,omitempty"`
	Author      EmbedAuthor    `json:"author,omitempty"`
	Fields      []EmbedFields  `json:"fields,omitempty"`
}

type EmbedFooter struct {
	Text         string `json:"text,omitempty"`
	IconUrl      string `json:"icon_url,omitempty"`
	ProxyIconUrl string `json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	Url      string `json:"url,omitempty"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedThumbnail struct {
	Url      string `json:"url,omitempty"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedVideo struct {
	Url      string `json:"url,omitempty"`
	ProxyUrl string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name string `json:"name,omitempty"`
	Url  string `json:"url,omitempty"`
}

type EmbedAuthor struct {
	Name         string `json:"name,omitempty"`
	Url          string `json:"url,omitempty"`
	IconUrl      string `json:"icon_url,omitempty"`
	ProxyIconUrl string `json:"proxy_icon_url,omitempty"`
}

type EmbedFields struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}
