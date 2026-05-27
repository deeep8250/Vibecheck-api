package config

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

var PostgresDB *sqlx.DB
var RedisClient *redis.Client
