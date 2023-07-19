package global

import (
	"github.com/jpc901/disk-common/conf"
	"github.com/jpc901/disk-common/db"
	"github.com/jpc901/disk-common/redis"
)

var (
	Config = conf.GetConfig()
	DB = db.GetDBInstance()
	RDB = redis.GetRDBInstance()
)