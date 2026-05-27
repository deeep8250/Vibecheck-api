package main

import (
	database "github.com/deeep8250/vibecheck-api/db/postgres"
	redisInit "github.com/deeep8250/vibecheck-api/db/redis"
)

func main() {
	database.DBinit()
	redisInit.RedisInit()
	Routes()

}
