package old

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	ana "github.com/ChimeraCoder/anaconda"
	"github.com/boltdb/bolt"
)

var twitterfriends = []byte("twitterfriends")

func main() {
	// Initialise boltdb
	db, err := bolt.Open("friends.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// setup twitter api
	ana.SetConsumerKey("Uub6rxnlkmCuf7WW4Xg0HDPJs")
	ana.SetConsumerSecret("5efz1p3T33kIrRIiBN6FiiabJ8DB21OhzGAefWLVPTzornlcuW")
	api := ana.NewTwitterApi("1409403655-9pFsvIJ3frI4g4jJQBqYDZ8GMflKkjDcpdSkHwZ", "0c5zM4fmtIuLVn2eg6wVxD6cYBWlG62FQNfylCrc6yZuh")

	va := url.Values{}
	va.Set("count", "5000")
	friends, _ := api.GetFriendsIds(va)

	// write data to boltdb
	for key, value := range friends.Ids {
		k := []byte(strconv.Itoa(key))
		v := []byte(strconv.FormatInt(value, 16))
		err = db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists(twitterfriends)
			if err != nil {
				return err
			}

			err = bucket.Put(k, v)
			if err != nil {
				return err
			}
			return nil

		})
	}

	if err != nil {
		log.Fatal(err)
	}

	// retrieve data from boltdb
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(twitterfriends)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", twitterfriends)
		}

		for key := range friends.Ids {
			val := bucket.Get([]byte(strconv.Itoa(key)))
			fmt.Println(string(val))
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}
