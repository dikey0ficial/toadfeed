package main

import (
	"encoding/json"
	"log"
	"os"
)

var users map[string]int64
var waitChan chan int8

func checkerr(err error) {
	if err != nil && err.Error() != "EOF" {
		log.Println("ERROR|ERROR|ERROR\n\n", err)
	}
}

func loadUsers() {
	file, err := os.Open("users.json")
	checkerr(err)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	checkerr(err)
	file.Close()
}

func main() {
	loadUsers()
	for token, id := range users {
		go StartFeed(token, id, waitChan)
	}
	log.Println("Started!")
	log.Println("...")
	for i := 0; i < len(users); i++ {
		<-waitChan
	}
}
