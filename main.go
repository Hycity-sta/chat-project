package main

import (
	"github.com/spf13/viper"
	
	"ginchat/models"
	"ginchat/router"
	"ginchat/utils"

	"time"
)

func main() {
	utils.InitConfig()

	utils.InitMysql()

	utils.InitRedis()

	InitTimer()

	r := router.Router()
	r.Run(viper.GetString("port.server"))

}

// 初始化定时器
func InitTimer() {

	DelayHeartbeat := viper.GetInt("timeout.DelayHeartbeat")
	duration_a := time.Duration(DelayHeartbeat) * time.Second

	HeartbeatHz := viper.GetInt("timeout.HeartbeatHz")
	duration_b := time.Duration(HeartbeatHz) * time.Second

	utils.Timer(duration_a, duration_b, models.CleanConnection, "")
}
