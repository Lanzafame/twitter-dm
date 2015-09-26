package main

import (
	"fmt"
	"net/url"

	ana "github.com/ChimeraCoder/anaconda"
)

func main() {
	// setup twitter api
	ana.SetConsumerKey("Uub6rxnlkmCuf7WW4Xg0HDPJs")
	ana.SetConsumerSecret("5efz1p3T33kIrRIiBN6FiiabJ8DB21OhzGAefWLVPTzornlcuW")
	api := ana.NewTwitterApi("1409403655-9pFsvIJ3frI4g4jJQBqYDZ8GMflKkjDcpdSkHwZ", "0c5zM4fmtIuLVn2eg6wVxD6cYBWlG62FQNfylCrc6yZuh")

	va := url.Values{}
	va.Set("count", "100")
	friends, _ := api.GetFriendsIds(va)

	users, _ := api.GetUsersLookupByIds(friends.Ids[:100], va)
	fmt.Print(users[0].Location)
	fmt.Print(friends.Ids[:100])
}
