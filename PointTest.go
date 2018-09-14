package main

import (
	"fmt"
	"sync"
)

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func main() {
	var feeds [3]Feed
	feeds[0].Name = "a1"
	feeds[0].URI = "a2"
	feeds[0].Type = "a3"

	feeds[1].Name = "b1"
	feeds[1].URI = "b2"
	feeds[1].Type = "b3"

	feeds[2].Name = "c1"
	feeds[2].URI = "c2"
	feeds[2].Type = "c3"

	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	for _, feed := range feeds {

		go func(feed2 *Feed) {
			printFeed(feed2)
			waitGroup.Done()
		}(&feed)
	}

	go func() {
		waitGroup.Wait()
	}()
	//waitGroup.Wait()
}

func printFeed(feed *Feed) {
	fmt.Print(feed.Name)
}
