package initalize

import (
	"tgwp/configs"
	"tgwp/global"
	"tgwp/internal/pkg/mysqlx"
	myRedis "tgwp/internal/pkg/redisx"
	"tgwp/log/zlog"
)

// InitDataBase 初始化数据库
func InitDataBase(config configs.Config) {
	switch config.Mysql.Driver {
	case "mysql":
		mysqlClient := mysqlx.NewMySql()
		db, err := mysqlClient.InitDataBases(config)
		if err != nil {
			zlog.Fatalf("数据库初始化失败！: %v", err)
			return
		}
		// 将初始化成功的数据库连接保存到全局变量中
		global.DB = db
		break
	// 可以添加更多数据库驱动的处理逻辑
	default:
		zlog.Fatalf("不支持的数据库驱动：%s", config.Mysql.Driver)
		return
	}

	if config.App.Env != "pro" {
		err := global.DB.AutoMigrate()
		if err != nil {
			zlog.Fatalf("数据库迁移失败！")
		}
	}

	zlog.Infof("数据库初始化成功！")
}
func InitRedis(config configs.Config) {
	if config.Redis.Enable {
		var err error
		global.Rdb, err = myRedis.GetRedisClient(config)
		if err != nil {
			zlog.Errorf("无法初始化Redis : %v", err)
		}
	} else {
		zlog.Warnf("不使用Redis")
	}

}
