package init

import (
	"log"
	"os"
	v "silent-cxl.top/app/validators"
	"silent-cxl.top/app/vars"
	"silent-cxl.top/library/config"
	libmysql "silent-cxl.top/library/gorm/mysql"
	"silent-cxl.top/library/hlog"
	libredis "silent-cxl.top/library/redis"
	"time"
)

func init() {
	checkConfigFiles()

	// 初始化 WEB 配置文件
	vars.YmlConfig = config.CreateYamlFactory()
	vars.YmlConfig.ConfigFileChangeListen()

	// 检测项目需要的目录
	checkDirs()

	initDb()
	initRedis()

	v.RegisterValidators()

	// 初始日志
	hlog.NewHLog()
}

// 检查必要的配置文件
func checkConfigFiles() {
	if _, err := os.Stat(vars.BasePath + "/config/web.yaml"); err != nil {
		log.Fatal("请检查 WEB 配置文件是否存在：", err)
		return
	}
}

// 项目需要的目录
func checkDirs() {
	// 日志目录
	dirs := []string{
		vars.YmlConfig.GetString("Logs.LogDir"),
	}

	for _, dir := range dirs {
		if _, err := os.Stat(dir); err != nil {
			_ = os.MkdirAll(dir, 0755)
		}
	}
}

func initDb() {
	addr := vars.YmlConfig.GetString("Mysql.Host")
	db := vars.YmlConfig.GetString("Mysql.Database")
	username := vars.YmlConfig.GetString("Mysql.Username")
	pwd := vars.YmlConfig.GetString("Mysql.Password")
	charset := vars.YmlConfig.GetString("Mysql.Charset")
	port := vars.YmlConfig.GetInt("Mysql.Port")
	life := vars.YmlConfig.GetInt("Mysql.SetConnMaxLifetime")
	if mysql, err := libmysql.NewMysql(addr, username, pwd, db, charset, port, time.Second*time.Duration(life)); err != nil {
		log.Fatal("Redis 连接失败", err.Error())
		return
	} else {
		vars.Mysql = mysql
	}
}

func initRedis() {
	addr := vars.YmlConfig.GetString("Redis.Host")
	pwd := vars.YmlConfig.GetString("Redis.Password")
	db := vars.YmlConfig.GetInt("Redis.Db")
	prefix := vars.YmlConfig.GetString("Redis.KeyPrefix")
	if redis, err := libredis.NewRedis(addr, pwd, prefix, db); err != nil {
		log.Fatal("Redis 连接失败", err.Error())
		return
	} else {
		vars.Redis = redis
	}
}
