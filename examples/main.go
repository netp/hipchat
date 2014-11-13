package main

import (
	"github.com/netp/hipchat"
	"log"
)

func main() {
	c := hipchat.Client{AuthToken: "<your api key>", ApiVersion: 2}

	req := hipchat.MessageRequest{
		RoomId:        "Test",
		Message:       "This message\n has been sent from \nclient\n" +
      "github.com/netp/hipchat\n" + 
      "with a lot of newlines in the form \n but hipchat will \n" +
      "ignore them anyway.\n",
		Color:         hipchat.ColorPurple,
		MessageFormat: hipchat.FormatText,
		Notify:        true,
	}
	if err := c.PostMessage(req); err != nil {
		log.Printf("Expected no error, but got %q", err)
	}
}
