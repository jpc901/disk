package global

import (
	"github.com/jpc901/disk-common/conf"
	"github.com/jpc901/disk-common/db"
)

var (
	Config = conf.GetConfig()
	DB = db.GetDBInstance()
)