package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ceckles/go-rss-scraper/internal/database"
)

func startScrapping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Scrapping on %v goroutines with a time between request of %v", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		log.Printf("Scrapping")
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error while getting feeds to scrap: %v", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("Error while marking feed as fetched: %v", err)
		return
	}
	log.Printf("Fetching %v", feed.Url)
	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Printf("Error while fetching feed: %v", err)
		return
	}
	for _, item := range rssFeed.Channel.Item {
		log.Println("Fetching Post", item.Title, "on feed", feed.Name)
	}
	log.Printf("Feed %s collected %v post found", feed.Name, len(rssFeed.Channel.Item))
}
