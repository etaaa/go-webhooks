# go-webhooks

A easy to use Golang 

## Installation

```bash
go get github.com/etaaa/go-webhooks
```

## Example


```go
package main

import (
	"log"
	goWebhooks "github.com/etaaa/go-webhooks"
)

func main() {
	// SET YOUR WEBHOOK URL
	webookUrl := "https://discord.com/api/webhooks/.../..."
	// CREATE A NEW WEBHOOK OBJECT - MOST FIELDS IN THE EXAMPLE ARE OPTIONAL BUT THERE ARE MORE AVAILBABLE 
	newWebhook := goWebhooks.Webhook{
		Content:   "This is the webhook's content - up to 2000 characters long.",
		Username:  "go-webhooks",
		AvatarUrl: "https://golang.org/lib/godoc/images/footer-gopher.jpg",
		Embeds: []goWebhooks.Embed{
			{
				Title:       "Embed Title with URL",
				Description: "This is the embed's description",
				Url:         "https://github.com/etaaa/go-webhooks",
				Timestamp:   goWebhooks.GetTimestamp(), // RETURNS NEW TIMESTAMP ACCORDING TO DISCORD'S FORMAT
				Color:       goWebhooks.GetColor("#00ff00"), // RETURNS COLOR ACCORDING TO DISCORD'S FORMAT
				Footer: goWebhooks.EmbedFooter{
					Text: "Sent via github.com/etaaa/go-webhooks",
				},
				Thumbnail: goWebhooks.EmbedThumbnail{
					Url: "https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png",
				},
				Author: goWebhooks.EmbedAuthor{
					Name:    "eta",
					Url:     "https://github.com/etaaa",
					IconUrl: "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png",
				},
				Fields: []goWebhooks.EmbedFields{
					{
						Name:  "Field 1",
						Value: "Basic text here",
					},
				},
			},
		},
	}
	// SEND THE WEBHOOK
	err := goWebhooks.SendWebhook(webookUrl, newWebhook, true)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Questions
For any questions feel free to DM me on Discord (eta#0001).

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)