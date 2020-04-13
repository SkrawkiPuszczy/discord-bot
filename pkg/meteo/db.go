package meteo

import "github.com/skrawkipuszczy/discord-bot/pkg/db"

type City struct {
	db.Base
}

func (c *City) TableName() string {
	return "meteos_cities"
}
