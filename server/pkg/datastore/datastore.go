package datastore

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/omgupta1608/chatex/server/cmd/config"
)

var (
	database *redis.Client
	Hashes   = map[string]string{
		"UnDelivered Message": "UD_MSG",
		"Last Seen":           "LSEEN",
	}
)

func init() {
	database = redis.NewClient(&redis.Options{
		Addr:     config.Flags.RedisPath,
		Password: config.Flags.RedisAuth,
		DB:       0,
	})
}

func ParseRedisIntCmd(resp *redis.IntCmd) bool {
	if statusCode, _ := resp.Result(); statusCode == 1 {
		return true
	}
	return false
}

// Just for testing
func StoreUndeliveredMessages(messages []string, id string) error {
	if resp := database.HSet(context.Background(), Hashes["UnDelivered Message"], id, messages); ParseRedisIntCmd(resp) {
		return nil
	} else {
		return resp.Err()
	}
}
