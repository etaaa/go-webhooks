# go-webhooks

A easy to use Golang package to quickly send Discord webhooks (https://discord.com/developers/docs/resources/webhook).

## Usage

Install:
```bash
go get github.com/etaaa/go-webhooks
```

Usage:
```go
package main

import (
	"log"
	wh "github.com/etaaa/go-webhooks"
)

func main() {
	// CREATE A NEW WEBHOOK OBJECT - MOST FIELDS IN THE EXAMPLE ARE OPTIONAL AND THERE ARE MORE AVAILBABLE
	content := wh.Webhook{
		Content:   "This is the webhook's content - up to 2000 characters long.",
		Username:  "go-webhooks",
		AvatarUrl: "https://golang.org/lib/godoc/images/footer-gopher.jpg",
		Embeds: []wh.Embed{
			{
				Title:       "Embed Title with URL",
				Description: "This is the embed's description",
				Url:         "https://github.com/etaaa/go-webhooks",
				Timestamp:   wh.GetTimestamp(),      // RETURNS NEW TIMESTAMP ACCORDING TO DISCORD'S FORMAT
				Color:       wh.GetColor("#00ff00"), // RETURNS COLOR ACCORDING TO DISCORD'S FORMAT
				Footer: wh.EmbedFooter{
					Text: "Sent via github.com/etaaa/go-webhooks",
				},
				Thumbnail: wh.EmbedThumbnail{
					Url: "https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png",
				},
				Author: wh.EmbedAuthor{
					Name:    "eta",
					Url:     "https://github.com/etaaa",
					IconUrl: "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png",
				},
				Fields: []wh.EmbedFields{
					{
						Name:  "Field 1",
						Value: "Text here",
					},
				},
			},
		},
	}
	// SEND THE WEBHOOK
	if err := wh.SendWebhook("https://discord.com/api/webhooks/.../...", content, true); err != nil {
		log.Fatal(err)
	}
}
```

## Questions
For any questions feel free to add and DM me on Discord (eta#0001).

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)