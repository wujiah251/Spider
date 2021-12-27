package main

import (
	"Spider/mongodb/dict"
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoCli *mongo.Client

func initEngine() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
}
func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}
	return mgoCli
}

var (
	client     = GetMgoCli()
	dictClient = client.Database("search").Collection("dict")
	record     = new(dict.DictRecord)
	// iResult    *mongo.InsertOneResult
	// id         primitive.ObjectID
	count = 0
)

func main() {
	file, err := os.Open("../idf/idf.utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	br := bufio.NewReader(file)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		temp := strings.Split(string(a), " ")
		if len(temp) != 2 {
			break
		}
		record.Word = temp[0]
		record.Idf, err = strconv.ParseFloat(temp[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		if _, err = dictClient.InsertOne(context.TODO(), record); err != nil {
			log.Fatal(err)
			return
		}
		count++
		// //_id:默认生成一个全局唯一ID
		// id = iResult.InsertedID.(primitive.ObjectID)
		// fmt.Println("自增ID", id.Hex())
	}
	fmt.Printf("count:%d\n", count)
}
