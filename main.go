package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type SlackPostPayload struct {
	UserName    string             `json:"username"`
	IconURL     string             `json:"icon_url"`
	Channel     string             `json:"channel"`
	LinkNames   bool               `json:"link_names"`
	Attachments []*SlackAttachment `json:"attachments"`
}

type SlackAttachment struct {
	Color string `json:"color"`
	Text  string `json:"text"`
}

func main() {
	payload := &SlackPostPayload{
		UserName:  "go-async-test",
		IconURL:   "https://avatars0.githubusercontent.com/u/25784240?s=400&v=4",
		Channel:   "#channel",
		LinkNames: true,
		Attachments: []*SlackAttachment{
			&SlackAttachment{
				Color: "good",
				Text:  "hi",
			},
		},
	}

	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"{{SLACK_HOOK_URL}}",
		bytes.NewBuffer(b),
	)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}
