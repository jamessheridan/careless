package main

import (
  "fmt"
  "strconv"
  "github.com/ChimeraCoder/anaconda"
  "github.com/boltdb/bolt"
  "stathat.com/c/jconfig"
  "log"
  "time"
)

var ConsumerKey, ConsumerSecret, AccessToken, AccessTokenSecret, DatabaseFile, DatabaseBucket string

func init() {
  config := jconfig.LoadConfig("careless.conf")
  ConsumerKey       = config.GetString("ConsumerKey")
  ConsumerSecret    = config.GetString("ConsumerSecret")
  AccessToken       = config.GetString("AccessToken")
  AccessTokenSecret = config.GetString("AccessTokenSecret")
  DatabaseFile      = config.GetString("DatabaseFile")
  DatabaseBucket    = config.GetString("DatabaseBucket")
}

func delaySecond(n time.Duration) {
  time.Sleep(n * time.Second)
}

func main() {
  db, err := bolt.Open(DatabaseFile, 0644, nil)
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  err = db.Update(func(tx *bolt.Tx) error {
    bucket, err := tx.CreateBucketIfNotExists([]byte(DatabaseBucket))
    if err != nil {
      return fmt.Errorf("create bucket: %s", bucket)
    }
  return err
  })

  anaconda.SetConsumerKey(ConsumerKey)
  anaconda.SetConsumerSecret(ConsumerSecret)
  api := anaconda.NewTwitterApi(AccessToken, AccessTokenSecret)

  search_result, err := api.GetSearch("i could care less", nil)
  if err != nil {
    log.Fatal(err)
  }

  for _, tweet := range search_result.Statuses {
    err = db.Update(func(tx *bolt.Tx) error {
      bucket := tx.Bucket([]byte(DatabaseBucket))
      v := bucket.Get([]byte(strconv.FormatInt(tweet.Id, 10)))
      if v != nil {
        fmt.Printf("Tweet found in db, not tweeting abuse this time.\n")
      } else {
        fmt.Printf("%v - %v - @%v: %v\n", tweet.Id, tweet.CreatedAt, tweet.User.ScreenName, tweet.Text)
        fmt.Printf("Adding to db: Key:%v - Value:%v\n", tweet.Id, tweet.User.ScreenName)
        key := strconv.FormatInt(tweet.Id, 10)
        value := tweet.User.ScreenName
        err = bucket.Put([]byte(key), []byte(value))
        if err != nil {
          log.Fatal(err)
        }
        message := fmt.Sprintf("Hi @%v! You probably meant that you couldn't care less, not that you could care less. Have a great day!\n", tweet.User.ScreenName)
        fmt.Printf(message)
        fmt.Printf("Waiting for 10 seconds to avoid rate limit")
        delaySecond(10)
        _, err = api.PostTweet(message, nil)
        if err != nil {
          log.Fatal(err)
        }
      }
    return nil
    })
  }
}
