package utils

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 数据库全局共享
var (
	DB  *gorm.DB
	Red *redis.Client
)

// 初始化配置文件
func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func InitMysql() {
	// 自定义logger用于gorm日志记录
	log_writer := log.New(os.Stdout, "\r\n", log.LstdFlags)
	log_config := logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Info,
		Colorful:      true,
	}

	newlog := logger.New(log_writer, log_config)

	var (
		username = viper.GetString("mysql.username")
		password = viper.GetString("mysql.password")
		host     = viper.GetString("mysql.host")
		port     = viper.GetInt("mysql.port")
		dbname   = viper.GetString("mysql.dbname")
		dsn      = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newlog})
	DB = db

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}
}

func InitRedis() {
	redis_opt := redis.Options{
		Addr:     viper.GetString("redis.Addr"),
		Password: viper.GetString("redis.Password"),
		DB:       2,
		PoolSize: 30,
	}
	Red = redis.NewClient(&redis_opt)
}

const PublishKey = "websocket"

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error

	fmt.Println("Publish 。。。。", msg)

	err = Red.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}

	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	fmt.Println("Subscribe 。。。。", ctx)

	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Println("Subscribe 。。。。", msg.Payload)
	return msg.Payload, err
}
