package main

import (
	"github.com/go-vk-api/vk"
	"time"
)

func StartFeed(token string, PeerID int64, wChan chan int8) {
	client, err := vk.NewClientWithOptions(
		vk.WithToken(token),
	)

	checkerr(err)

	for true {
		err := client.CallMethod("messages.send", vk.RequestParams{
			"peer_id":   2000000000 + PeerID,
			"message":   "Покормить жабу",
			"random_id": 0,
		}, nil)
		checkerr(err)
		time.Sleep(12*time.Hour + 1*time.Second)
	}
	wChan <- 0
}
