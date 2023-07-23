package global

import (
	redis "github.com/jpc901/disk-common/cache"
	"github.com/jpc901/disk-common/conf"
	"github.com/jpc901/disk-common/db"
)

var (
	Config = conf.GetConfig()
	DB = db.GetDBInstance()
	RDB = redis.GetRDBInstance()
)