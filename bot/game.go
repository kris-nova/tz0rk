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
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
)

// Process will build the response and poll and respond to the found tweet.
//
// Sample poll json
// "entities":{
//      "hashtags":[],
//      "urls":[],
//      "user_mentions":[],
//      "symbols":[],
//      "polls":[
//         {
//            "options":[
//               {
//                  "position":1,
//                  "text":"The better answer"
//               },
//               {
//                  "position":2,
//                  "text":"The best answer"
//               }
//            ],
//            "end_datetime":"Sat Feb 04 15:33:11 +0000 2017",
//            "duration_minutes":1440
//         }
//      ]
//   },
func (b *Bot) Process(tweet anaconda.Tweet) error {
	v := url.Values{}
	v.Set("in_reply_to_status_id", strconv.Itoa(int(tweet.Id)))
	entities := map[interface{}]interface{}{
		"meeps": "meeps",
	}
	jsonBytes, err := json.Marshal(&entities)
	if err != nil {
		return err
	}
	jsonStr := string(jsonBytes)
	v.Set("entities", jsonStr)
	exampleMessage := "You enter a cave. You:"
	status := fmt.Sprintf("@%s %s", tweet.User.Name, exampleMessage)
	b.api.PostTweet(status, v)

	return nil
}
