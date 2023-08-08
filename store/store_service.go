package store

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	// ctx          = context.Background()
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortURL string, originalURL string, useraId string) {
	err := storeService.redisClient.Set(shortURL, originalURL, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - ShortURL: %s - OriginalURL: %s\n", err, shortURL, originalURL))
	}
}

func RetrieveInitialUrl(shortURL string) string {
	result, err := storeService.redisClient.Get(shortURL).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortURL))
	}
	return result
}
