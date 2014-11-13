package main

import (
	"github.com/andybons/hipchat"
	"log"
)

func main() {
	c := hipchat.Client{AuthToken: "<PUT YOUR API TOKEN HERE>", ApiVersion: 2}

	req := hipchat.MessageRequest{
		RoomId:        "Goal",
		Message:       "Good news: Here's my fork github.com/netp/hipchat, which suports api (v2)",
		Color:         hipchat.ColorPurple,
		MessageFormat: hipchat.FormatText,
		Notify:        true,
	}
	if err := c.PostMessage(req); err != nil {
		log.Printf("Expected no error, but got %q", err)
	}
}
