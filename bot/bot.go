/*
Copyright © 2020 Kris Nova <kris@nivenly.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package bot

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/kris-nova/logger"
)

const (
	sleepTimeMinutes int = 1
	searchQuery          = "/tz0rk"
)

// Bot is the main bot service
type Bot struct {
	api *anaconda.TwitterApi
}

// New will initialize a new Bot struct
func New() *Bot {
	return &Bot{}
}

// Auth will authenticate and validate or return an error
func (b *Bot) Auth() error {
	key := os.Getenv("TZ0RK_API_KEY")
	sec := os.Getenv("TZ0RK_API_SECRET")
	tkn := os.Getenv("TZ0RK_TOKEN")
	tsc := os.Getenv("TZ0RK_TOKEN_SECRET")
	api := anaconda.NewTwitterApiWithCredentials(tkn, tsc, key, sec)
	ok, err := api.VerifyCredentials()
	if !ok {
		return fmt.Errorf("invalid credentials: %v", err.Error())
	}
	b.api = api
	return nil
}

// Run is the method that will run the bot. By design the first thing this method does is find the most recent
// tweet containing our query. It will only process tweets that happen after that so we reduce the risk of spamming
// twitter and looking like a bunch of n00bs.
//
// So basically, if this process isn't running while a tweet is tweeted with the query - it will be ignored forever.
func (b *Bot) Run(errch chan error) {
	v := url.Values{}
	v.Set("result_type", "recent")
	logger.Always("[API]  Finding recent tweets")
	latestTweets, err := b.api.GetSearch(searchQuery, nil)
	if err != nil {
		logger.Warning("unable to find latest tweet: %v", err)
		errch <- fmt.Errorf("unable to find latest tweet: %v", err)
		return
	}
	logger.Always("Found %d tweets to process", len(latestTweets.Statuses))
	for _, tweet := range latestTweets.Statuses {
		if tweet.Id > b.getLast() {
			b.cacheLast(tweet.Id)
		}
	}
	for {
		v := url.Values{}
		last := strconv.Itoa(int(b.getLast()))
		v.Set("since_id", last)
		logger.Always("[API] Searching for tweets after: %s", last)
		search, err := b.api.GetSearch(searchQuery, v)
		if err != nil {
			logger.Warning("unable to find tweets: %v", err)
			errch <- fmt.Errorf("unable to find tweets: %v", err)
			continue
		}
		logger.Always("Found %d tweets to process", len(search.Statuses))
		for _, tweet := range search.Statuses {
			// Cache first so we don't fuck up and spam twitter
			b.cacheLast(tweet.Id)
			err := b.Process(tweet)
			errch <- err
		}
		logger.Always("Hanging for %d minutes...", sleepTimeMinutes)
		time.Sleep(time.Duration(sleepTimeMinutes) * time.Minute)
	}
}

var lastID int64 = 0

func (b *Bot) cacheLast(id int64) error {
	logger.Always("Caching ID: %v", id)
	lastID = id
	return nil
}

func (b *Bot) getLast() int64 {
	return lastID
}
