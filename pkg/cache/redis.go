package cache

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/go-redis/redis/v7"
)

type redisClient struct {
	cl *redis.Client
}

//NewRedisClient create redis client based on url
func NewRedisClient(url string) (Cache, error) {
	options, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(options)
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &redisClient{
		cl: client,
	}, nil
}

//Close close connection
func (r *redisClient) Close() {
	log.Println("redis client closed")
	r.cl.Close()
}

//SetLocation set geo coordinates for place in cache
func (r *redisClient) SetLocation(name string, lat, long float64) (*location, error) {
	data := &location{Name: name, Latitude: lat, Longitude: long}
	json, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = r.cl.Set(prepareKey(Location, name), json, 0).Err()
	if err != nil {
		return nil, err
	}
	return data, nil
}

//GetLocation return geo coordinates for place
func (r *redisClient) GetLocation(name string) (*location, error) {
	str, err := r.cl.Get(prepareKey(Location, name)).Result()
	loc := &location{}
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		err = json.Unmarshal([]byte(str), &loc)
		if err != nil {
			return nil, err
		}
		return loc, nil
	}
}
func (r *redisClient) SetPhoto(keyName string, data string) error {
	_, err := r.cl.HSet("photos", keyName, data).Result()
	if err != nil {
		return err
	}
	return nil
}
func (r *redisClient) GetPhotos() (map[string]string, error) {
	return r.cl.HGetAll("photos").Result()
}

func (r *redisClient) SetChannelMessagesCounter(keyName string, data int) error {
	return r.cl.HSet("messages_count", keyName, strconv.Itoa(data)).Err()

}

func (r *redisClient) GetChannelMessagesCounter(keyName string) (int, error) {
	data, err := r.cl.HGet("messages_count", keyName).Result()
	if err == redis.Nil {
		data = "0"
	} else if err != nil {
		return 0, err
	}
	return strconv.Atoi(data)
}
