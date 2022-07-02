package webhooks

import (
	"log"
	"testing"
)

func TestSendWebhook(t *testing.T) {
	// Create a new webhook object. Most fields are optional
	webhook := Webhook{
		Content:   "This is the webhook's content - up to 2000 characters long.",
		Username:  "go-webhooks",
		AvatarUrl: "https://golang.org/lib/godoc/images/footer-gopher.jpg",
		Embeds: []Embed{
			{
				Title:       "Embed Title with URL",
				Description: "This is the embed's description",
				Url:         "https://github.com/etaaa/go-webhooks",
				Timestamp:   GetTimestamp(),      // Returns a new timestamp matching Discords format
				Color:       GetColor("#00ff00"), // Returns the color in decimal value matching Discords format
				Footer: EmbedFooter{
					Text: "Sent via github.com/etaaa/go-webhooks",
				},
				Thumbnail: EmbedThumbnail{
					Url: "https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png",
				},
				Author: EmbedAuthor{
					Name:    "eta",
					Url:     "https://github.com/etaaa",
					IconUrl: "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png",
				},
				Fields: []EmbedFields{
					{
						Name:  "Field 1",
						Value: "Text here",
					},
				},
			},
		},
	}
	// Send the webhook
	if err := SendWebhook("https://discord.com/api/webhooks/.../...", webhook, true); err != nil {
		log.Fatal(err)
	}
}
