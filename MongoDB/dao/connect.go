package dao

import (
    "context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "log"
)

var Client *mongo.Client
//var DB *mongo.Database

func Connect() {
    //clientOptions := options.Client().ApplyURI("mongodb://120.79.29.70:27017")

    credential := options.Credential{
        Username:      "ww",
        Password:      "123111",
    }
    clientOptions := options.Client().ApplyURI("mongodb://ww:123111@120.79.29.70:27017").SetAuth(credential) // 不知道行不行
    //Client, err := mongo.Connect(
    //    context.TODO(),
    //    clientOptions,
    //    options.Client().SetAuth(awsCredential))

    // 连接到MongoDB
    c, err := mongo.Connect(context.TODO(), clientOptions)
    Client = c
    if err != nil {
        log.Fatal(err)
    }

    // 检查连接
    err = Client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB!")

}
func Close() {
    err := Client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connection to MongoDB closed.")
}
func getDB(dbName string) (DB *mongo.Database) {
    DB = Client.Database(dbName)
    return
}

func getCollection(db *mongo.Database) {
    collection := db.Collection("student")
    filter := bson.D{{"name", "小王子"}}
    singleResult := collection.FindOne(context.TODO(), filter)
    var res student
    err := singleResult.Decode(&res)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(res)
}

type student struct {
    Name string  // 字段必须大写，外部包才能在查询时修改传进去的对象字段
    Age int
}

func main() {
    Connect()
    db := getDB("jinyun")
    getCollection(db)

}

