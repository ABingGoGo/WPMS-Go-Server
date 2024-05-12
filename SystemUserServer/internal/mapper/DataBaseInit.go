package mapper

import (
	"SystemUserServer/config"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

var Engine *xorm.Engine

func DataBaseMapperInit() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", config.YamlConfig.DataBase.Username, config.YamlConfig.DataBase.Password, config.YamlConfig.DataBase.Url,
		config.YamlConfig.DataBase.Port, config.YamlConfig.DataBase.Name, config.YamlConfig.DataBase.UrlAdditionalParameters)
	var err error
	Engine, err = xorm.NewEngine("mysql", dsn)
	if err != nil {
		logrus.Errorf("%s", "[XORM] - 数据库连接字符串存在问题，请检查\n")
	}
	err = Engine.Ping()
	if err != nil {
		logrus.Errorf("%s", "[XORM] - 数据库连接字符串没有问题，但无法建立数据库连接，请检查配置或网络等相关问题\n")
	} else {
		logrus.Info("[XORM] - 数据库连接成功")
	}

	Engine.ShowSQL(true)
	Engine.SetConnMaxIdleTime(5 * time.Second)
	Engine.SetMaxOpenConns(10)
	Engine.SetMaxIdleConns(2)

}
