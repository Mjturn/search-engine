package main

import (
    "fmt"
    "log"
    "os"
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "github.com/joho/godotenv"
    "github.com/Mjturn/search-engine/routes"
)

func main() {
    err := godotenv.Load("../.env")
    if err != nil {
        log.Fatal(err)
    }

    databaseUsername := os.Getenv("DATABASE_USERNAME")
    databasePassword := os.Getenv("DATABASE_PASSWORD")
    
    databaseConnectionString := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.wwuyzc5.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0", databaseUsername, databasePassword)
    
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(databaseConnectionString).SetServerAPIOptions(serverAPI)

    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
        panic(err)
    }

    defer func() {
        if err = client.Disconnect(context.TODO()); err != nil {
            panic(err)
        }
    } ()
    
    if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
        panic(err)
    }

    fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

    router := gin.Default()
    router.Static("/static", "../frontend/static")
    router.LoadHTMLGlob("../frontend/templates/*")
    routes.HandleRoutes(router)

    err = router.Run(":8080")
    if err != nil {
        log.Fatal(err)
    }
}
