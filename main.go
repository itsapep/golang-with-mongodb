package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MONGODB_URI = "mongodb://localhost:27017"

type Student struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Age      int
	Gender   string
	JoinDate primitive.DateTime `bson:"joinDate"`
	Senior   bool
}

func main() {
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		// AuthMechanismProperties: map[string]string{},
		// AuthSource:              "",
		Username: "yurhamafif",
		Password: "12345678",
		// PasswordSet:             false,
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI(MONGODB_URI).SetAuth(credential)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	connect, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connected ...")
	}

	defer func() {
		err := connect.Disconnect(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	// create db - connection
	// db := connect.Database("enigma")
	// collection := db.Collection("student")

	// // create InsertOne
	// newId, err := collection.InsertOne(ctx, bson.D{
	// 	{Key: "name", Value: "Jack"},
	// 	{Key: "age", Value: 22},
	// 	{Key: "gender", Value: "M"},
	// 	{Key: "senior", Value: false},
	// })

	// create InsertMany
	// newId, err := collection.InsertMany(ctx, []interface{}{
	// 	bson.D{
	// 		{Key: "name", Value: "sita"},
	// 		{Key: "age", Value: 24},
	// 		{Key: "gender", Value: "F"},
	// 		{Key: "senior", Value: true},
	// 	},
	// 	bson.D{
	// 		{Key: "name", Value: "Melani"},
	// 		{Key: "age", Value: 23},
	// 		{Key: "gender", Value: "F"},
	// 		{Key: "senior", Value: true},
	// 	},
	// 	bson.D{
	// 		{Key: "name", Value: "Suci"},
	// 		{Key: "age", Value: 21},
	// 		{Key: "gender", Value: "F"},
	// 		{Key: "senior", Value: false},
	// 	},
	// })

	// // create InsertMany
	// jd01 := parseTime("2022-07-02 15:04:05")
	// jd02 := parseTime("2022-07-03 20:59:59")
	// jd03 := parseTime("2022-07-04 15:04:05")
	// newId, err := collection.InsertMany(ctx, []interface{}{
	// 	bson.D{
	// 		{Key: "name", Value: "sita"},
	// 		{Key: "age", Value: 24},
	// 		{Key: "gender", Value: "F"},
	// 		{Key: "joinDate", Value: primitive.NewDateTimeFromTime(jd01)},
	// 		{Key: "senior", Value: true},
	// 	},
	// 	bson.D{
	// 		{Key: "name", Value: "Melani"},
	// 		{Key: "age", Value: 23},
	// 		{Key: "gender", Value: "F"},
	// 		{Key: "joinDate", Value: primitive.NewDateTimeFromTime(jd02)},
	// 		{Key: "senior", Value: true},
	// 	},
	// 	bson.D{
	// 		{Key: "name", Value: "Suci"},
	// 		{Key: "age", Value: 21},
	// 		{Key: "gender", Value: "F"},
	// 		{Key: "joinDate", Value: primitive.NewDateTimeFromTime(jd03)},
	// 		{Key: "senior", Value: false},
	// 	},
	// })

	// // using struct
	// newStudent := Student{
	// 	Id:       primitive.NewObjectID(),
	// 	Name:     "Doni",
	// 	Age:      20,
	// 	Gender:   "M",
	// 	JoinDate: primitive.NewDateTimeFromTime(parseTime("2022-07-13 13:45:04")),
	// 	Senior:   false,
	// }
	// newId, err := collection.InsertOne(ctx, newStudent)

	// deleteOne
	// filter := bson.D{{Key: "name", Value: "Suci"}}
	// opts := options.Delete().SetHint(bson.D{{Key: "_id", Value: 1}}) //Specifies the method to use the _id as the index
	// result, err := collection.DeleteOne(ctx, filter)

	// // deletoMany
	// filter := bson.D{{Key: "age", Value: 20}}
	// opts := options.Delete().SetHint(bson.D{{Key: "_id", Value: 1}}) //Specifies the method to use the _id as the index
	// result, err := collection.DeleteMany(ctx, filter, opts)

	// update by name
	// filter := bson.D{{Key: "name", Value: "sita"}}
	// update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: "Sita"}}}}
	// result, err := collection.UpdateOne(ctx, filter, update)
	// fmt.Printf("Documents matched: %v\n", result.MatchedCount)
	// fmt.Printf("Documents updated: %v\n", result.ModifiedCount)

	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Printf("inserted document with id %s\n", newId.InsertedID)
	// fmt.Printf("Number of documents deleted: %d\n", result.DeletedCount)

	// Read
	// select * from student
	// cursor, err := collection.Find(ctx, bson.D{})
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// var students []bson.D
	// err = cursor.All(ctx, &students)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// for _, student := range students {
	// 	fmt.Println(student)
	// }

	// projection
	// opt := options.FindOptions{Projection: bson.D{{Key: "name", Value: 1}, {Key: "_id", Value: 0}}}
	// cursor, err := collection.Find(ctx, bson.D{}, &opt)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// var students []bson.D
	// err = cursor.All(ctx, &students)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// for _, student := range students {
	// 	fmt.Println(student)
	// }

	// logical
	// filterGenderAndAge := bson.D{
	// 	{Key: "$and", Value: bson.A{
	// 		bson.D{
	// 			{Key: "gender", Value: "F"},
	// 			{Key: "age", Value: bson.D{
	// 				{Key: "$gte", Value: 24},
	// 			}},
	// 		},
	// 	}}}
	// projection := bson.D{
	// 	{Key: "_id", Value: 0},
	// 	{Key: "gender", Value: 1},
	// 	{Key: "age", Value: 1},
	// }
	// cursor, err := collection.Find(ctx, filterGenderAndAge, &options.FindOptions{Projection: projection})
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// var students []bson.D
	// err = cursor.All(ctx, &students)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// for _, student := range students {
	// 	fmt.Println(student)
	// }

	// mapping result query to struct
	// filterGenderAndAge := bson.D{
	// 	{Key: "$and", Value: bson.A{
	// 		bson.D{
	// 			{Key: "gender", Value: "F"},
	// 			{Key: "age", Value: bson.D{
	// 				{Key: "$gte", Value: 24},
	// 			}},
	// 		},
	// 	}}}
	// projection := bson.D{
	// 	{Key: "_id", Value: 1},
	// 	{Key: "name", Value: 1},
	// 	{Key: "gender", Value: 1},
	// 	{Key: "age", Value: 1},
	// }
	// filterResult := make([]*Student, 0)
	// cursor, err := collection.Find(ctx, filterGenderAndAge, &options.FindOptions{Projection: projection})
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// for cursor.Next(ctx) {
	// 	var student Student
	// 	err := cursor.Decode(&student)
	// 	if err != nil {
	// 		log.Println(err.Error())
	// 	}
	// 	filterResult = append(filterResult, &student)
	// }
	// for _, student := range filterResult {
	// 	fmt.Println(student)
	// }

	// aggregation
	// create db - connection
	db := connect.Database("contohDb")
	collection := db.Collection("products")
	// count, err := collection.CountDocuments(ctx, bson.D{})
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println("product total : ", count)

	// // with filter
	// count, err = collection.CountDocuments(ctx, bson.D{{Key: "category", Value: "food"}})
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println("product total with category[food] : ", count)

	// match,group,sort,dll
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "category", Value: "food"},
		}},
	}
	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$category"},
			{Key: "Total", Value: bson.D{
				{Key: "$sum", Value: 1},
			}},
		}},
	}
	cursor, err := collection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		log.Println(err.Error())
	}
	var productCount []bson.M
	err = cursor.All(ctx, &productCount)
	if err != nil {
		log.Println(err.Error())
	}
	for _, product := range productCount {
		fmt.Printf("group[%v], total[%v]\n", product["_id"], product["Total"])
	}

}

func parseTime(date string) time.Time {
	layoutFormat := "2006-01-02 15:04:05"
	parse, _ := time.Parse(layoutFormat, date)
	return parse
}

// buat koneksi ke mongodb://localhost:27017
// siapkan user auth: username & password
