package main

import (
	"github.com/netp/hipchat"
	"log"
)

func main() {
	c := hipchat.Client{AuthToken: "API3HNZpYZKTsKEN6odo7knb1dA7L3OuPrnfw4KA", ApiVersion: 2}

	req := hipchat.MessageRequest{
		RoomId:        "Goal",
		Message:       "This message has been sent from client github.com/netp/hipchat",
		Color:         hipchat.ColorPurple,
		MessageFormat: hipchat.FormatText,
		Notify:        true,
	}
	if err := c.PostMessage(req); err != nil {
		log.Printf("Expected no error, but got %q", err)
	}
}
