package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Howie struct {
	HowieId primitive.ObjectID `bson:"_id"`
	Aa      string
	Bb      string
}

// var client *mongo.Client

func main() {
	opt := options.Client().ApplyURI("mongodb://192.168.1.89:27017")

	client, err := mongo.Connect(getContext(), opt)
	checkErr(err)

	collection := client.Database("crazywolf").Collection("test")
	cursor, _ := collection.Find(getContext(), nil)

	howieArrayEmpty := []Howie{}

	for cursor.Next(getContext()) {
		var howie Howie
		cursor.Decode(&howie)
		howieArrayEmpty = append(howieArrayEmpty, howie)
	}

	for _, h := range howieArrayEmpty {
		fmt.Println(h)
	}

}

func getContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return
}

func checkErr(err error) {
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("没有查到数据")
			os.Exit(0)
		} else {
			fmt.Println(err)
			os.Exit(0)
		}

	}
}
