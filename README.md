Hipchat
=====

This is a version of andybons/hipchat with support for Hipchat API v2.

This project implements a [Go](http://golang.org) client library for 
the [Hipchat API](https://www.hipchat.com/docs/api/) 

Installing
----------
Run
```bash
go get github.com/andybons/hipchat
```

Example usage:
```go
package main

import (
	"github.com/netp/hipchat"
	"log"
)

func main() {
	c := hipchat.Client{AuthToken: "<YOUR API TOKEN>", ApiVersion: 2}

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
}
```
