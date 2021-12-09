package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogRecord struct {
	JobName string `bson:"jobName"` // 任务名
	Command string `bson:"command"` // shell命令
	Err     string `bson:"err"`     // 脚本错误
	Content string `bson:"content"` // 脚本输出
}

func main() {
	var (
		client *mongo.Client
		err    error
		//result *mongo.InsertOneResult
	)

	// 建立mongodb连接
	clientOptions := options.Client().ApplyURI("mongodb://admin:admin123456@192.168.252.128:27017/yapi")
	if client, err = mongo.Connect(
		context.TODO(), clientOptions); err != nil {
		return
	}
	database := client.Database("PS")

	// 3, 选择表my_collection
	collection := database.Collection("processing")

	// 4, 插入记录(bson)
	/*record := &LogRecord{
		JobName: "job10",
		Command: "echo hello",
		Err: "",
		Content: "hello",
	}

	if result, err = collection.InsertOne(context.TODO(), record); err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println(result)

	//// _id: 默认生成一个全局唯一ID, ObjectID：12字节的二进制
	docId := result.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID:", docId.Hex())
	*/
	var result = &LogRecord{}
	filter := bson.D{{"jobName", "job10"}}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		fmt.Println("record does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
