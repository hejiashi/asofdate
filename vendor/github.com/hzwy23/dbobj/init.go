package dbobj

import (
	_ "github.com/hzwy23/dbobj/mysql"
	//_ "github.com/hzwy23/dbobj/oracle"
	"github.com/hzwy23/dbobj/utils"
	"os"
	"path/filepath"
)

func init() {
	HOME := os.Getenv("HBIGDATA_HOME")
	filedir := filepath.Join(HOME, "conf", "asofdate.conf")
	conf, err := utils.GetConfig(filedir)
	if err != nil {
		panic("init database failed.")
	}
	Default, err = conf.Get("DB.type")
	if err != nil {
		panic("get default database type failed.")
	}
	InitDB(Default)
}
