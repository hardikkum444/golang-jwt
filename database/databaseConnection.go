package database

import (
    "fmt"
    "os"
    "time"
    "log"
    "context"
    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() *mongo.Client {
   
    err := godotenv.Load(".env")

    if err!=nil{
        log.Fatal("Error loading the .env file")
    }

    MongoDb := os.Getenv("MONGODB_URL")

    client, err := mongo.NewClient(options.Client().ApplyURI(MongoDb))

    if err!=nil{
        log.Fatal(err)
    }

    // This will ensure release of all resources if db query takes indefinite amount of time
    ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancle()

    err = client.Connect(ctx)

    if err!=nil{
        // no need of returning as after logging it calls exit(0) which then exits from the code base
        log.Fatal(err)
    }

    fmt.Println("Connected to the mongodb database successfully")

    return client
}

var Client *mongo.Client = DBinstance() 

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {

    var collection *mongo.Collection = client.Database("cluster0").Collection(collectionName)
    return collection
}




