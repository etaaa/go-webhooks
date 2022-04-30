package webhooks

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	// CREATE A NEW WEBHOOK OBJECT - MOST FIELDS IN THE EXAMPLE ARE OPTIONAL AND THERE ARE MORE AVAILBABLE
	content := Webhook{
		Content:   "This is the webhook's content - up to 2000 characters long.",
		Username:  "go-webhooks",
		AvatarUrl: "https://golang.org/lib/godoc/images/footer-gopher.jpg",
		Embeds: []Embed{
			{
				Title:       "Embed Title with URL",
				Description: "This is the embed's description",
				Url:         "https://github.com/etaaa/go-webhooks",
				Timestamp:   GetTimestamp(),      // RETURNS NEW TIMESTAMP ACCORDING TO DISCORD'S FORMAT
				Color:       GetColor("#00ff00"), // RETURNS COLOR ACCORDING TO DISCORD'S FORMAT
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
	// SEND THE WEBHOOK
	if err := SendWebhook("https://discord.com/api/webhooks/.../...", content, true); err != nil {
		log.Fatal(err)
	}
}
