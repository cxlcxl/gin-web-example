package vars

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"os"
	"silent-cxl.top/library/config_interface"
	libredis "silent-cxl.top/library/redis"
	"strings"
)

var (
	BasePath  string
	YmlConfig config_interface.YamlConfigInterface
	Redis     *libredis.Redis
	Mysql     *gorm.DB
	HLog      *logrus.Logger
)

func init() {
	if dir, err := os.Getwd(); err != nil {
		log.Fatal("文件目录获取失败")
		return
	} else {
		// 路径进行处理，兼容单元测试程序程序启动时的奇怪路径
		if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test") {
			BasePath = strings.Replace(strings.Replace(dir, `\test`, "", 1), `/test`, "", 1)
		} else {
			BasePath = dir
		}
	}
}
