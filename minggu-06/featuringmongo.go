package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	//Ini bisa diganti jika mau menggunakan localhost, saya menggunakan versi atlas karena matakuliah ini bertemakan cloud service.
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://admin:root@175610064-ujktu.gcp.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sukses koneksi ke cluster 175610064")
	fmt.Println("Daftar database yang ada", databases)

	tccDatabase := client.Database("TCC")
	mahasiswaCollection := tccDatabase.Collection("mahasiswa")

	cursor, err := mahasiswaCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var mahasiswa []bson.M
	if err = cursor.All(ctx, &mahasiswa); err != nil {
		log.Fatal(err)
	}
	fmt.Println(mahasiswa)
}
