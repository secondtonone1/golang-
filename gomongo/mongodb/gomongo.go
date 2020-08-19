package mongodb

import (
	"context"
	"fmt"
	"golang-/gomongo/config"
	"golang-/gomongo/constants"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	DB *Database
)

type Database struct {
	Mongo *mongo.Client
}

//初始化
func Init() {
	DB = &Database{
		Mongo: SetConnect(),
	}
}

// 连接设置
func SetConnect() *mongo.Client {

	var retryWrites bool = false

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().SetHosts(config.GetConf().MongoConf.Hosts).
		SetMaxPoolSize(config.GetConf().MongoConf.MaxPoolSize).
		SetHeartbeatInterval(constants.HEART_BEAT_INTERVAL).
		SetConnectTimeout(constants.CONNECT_TIMEOUT).
		SetMaxConnIdleTime(constants.MAX_CONNIDLE_TIME).
		SetRetryWrites(retryWrites)

	//设置用户名和密码
	username := config.GetConf().MongoConf.Username
	password := config.GetConf().MongoConf.Password

	if len(username) > 0 && len(password) > 0 {
		clientOptions.SetAuth(options.Credential{Username: username, Password: password})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	DB.Mongo.Disconnect(ctx)
}
