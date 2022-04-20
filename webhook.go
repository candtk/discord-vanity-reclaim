package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func SendSuccess(vanity string, time string) {
	embed := []byte(fmt.Sprintf(`{
		"content": "@everyone",
		"embeds": [
		  {
			"title": "CLAIMED",
			"color": 65280,
			"fields": [
			  {
				"name": "Vanity",
				"value": "%s"
			  },
			  {
				"name": "Time Taken to attempt",
				"value": "%s"
			  }
			],
			"thumbnail": {
			  "url": "https://cdn.discordapp.com/embed/avatars/0.png"
			}
		  }
		],
		"attachments": []
	  }`, vanity, time))

	resp, err := http.Post(Config.Main.Webhook, "application/json", bytes.NewBuffer(embed))

	if err != nil {
		fmt.Println(err)
		SendSuccess(vanity, time)
	}

	defer resp.Body.Close()
}

func SendRatelimit(vanity string, time string) {
	embed := []byte(fmt.Sprintf(`{
		"content": "@everyone",
		"embeds": [
		  {
			"title": "RATELIMITED",
			"color": 16763904,
			"fields": [
			  {
				"name": "Vanity",
				"value": "%s"
			  },
			  {
				"name": "Time Taken to attempt",
				"value": "%s"
			  }
			],
			"thumbnail": {
			  "url": "https://cdn.discordapp.com/embed/avatars/0.png"
			}
		  }
		],
		"attachments": []
	  }`, vanity, time))

	resp, err := http.Post(Config.Main.Webhook, "application/json", bytes.NewBuffer(embed))

	if err != nil {
		fmt.Println(err)
		SendSuccess(vanity, time)
	}

	defer resp.Body.Close()
}

func SendFail(vanity string, time string, statuscode string) {
	embed := []byte(fmt.Sprintf(`{
		"content": "@everyone",
		"embeds": [
		  {
			"title": "CLAIM FAILED",
			"description": "Claim attempt unsuccessful",
			"color": 16763904,
			16711680
			"fields": [
			  {
				"name": "Vanity",
				"value": "%s"
			  },
			  {
				"name": "Time Taken to attempt",
				"value": "%s"
			  },
			  {
				"name": "Status Code",
				"value": "%s"
			  }
			],
			"thumbnail": {
			  "url": "https://cdn.discordapp.com/embed/avatars/0.png"
			}
		  }
		],
		"attachments": []
	  }`, vanity, time, statuscode))

	resp, err := http.Post(Config.Main.Webhook, "application/json", bytes.NewBuffer(embed))

	if err != nil {
		fmt.Println(err)
		SendSuccess(vanity, time)
	}

	defer resp.Body.Close()
}
