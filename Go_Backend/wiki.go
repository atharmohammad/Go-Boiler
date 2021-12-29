package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const collectionName = "Person"
const dbname = "PersonDb"

func getPost(c *fiber.Ctx) {
	collection, err := getDbCollections(dbname, collectionName)
	if err != nil {
		c.Status(500).Send(err)
		return
	}
	var filter bson.M = bson.M{}

	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		c.Status(500).Send(err)
	}
	defer curr.Close(context.Background())

	var results []bson.M
	curr.All(context.Background(), &results)

	if results == nil {
		c.Status(404).Send(`{"message":"There is no data !"}`)
		return
	}
	json, _ := json.Marshal(results)
	c.Status(200).Send(json)
}

func createPost(c *fiber.Ctx) {
	collection, err := getDbCollections(dbname, collectionName)
	if err != nil {
		c.Status(400).Send()
		return
	}
	var newPost Post
	json.Unmarshal([]byte(c.Body()), &newPost)
	res, err := collection.InsertOne(context.Background(), newPost)
	if err != nil {
		c.Status(400).Send()
		return
	}
	response, _ := json.Marshal(res)
	c.Status(200).Send(response)
}

func deletePost(c *fiber.Ctx) {
	collection, err := getDbCollections(dbname, collectionName)
	if err != nil {
		c.Status(400).Send()
		return
	}

	var filter bson.M
	id := c.Params("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	filter = bson.M{"_id": objId}
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.Status(400).Send(`{"message":"There is error in deleting the document"}`)
		return
	}
	response, _ := json.Marshal(res)
	c.Status(201).Send(response)
}

func updatePost(c *fiber.Ctx) {
	collection, err := getDbCollections(dbname, collectionName)
	if err != nil {
		c.Status(400).Send()
		return
	}
	var updatedPost Post
	json.Unmarshal([]byte(c.Body()), &updatedPost)
	updateFilter := bson.M{
		"$set": updatedPost,
	}
	ObjId, _ := primitive.ObjectIDFromHex(c.Params("id"))
	// curr, err := collection.Find(context.Background(), bson.M{"_id": ObjId})
	// if err != nil {
	// 	c.Status(400).Send(err)
	// 	return
	// }
	// defer curr.Close(context.Background())

	// var results []bson.M
	// curr.All(context.Background(), &results)
	// jsonData, _ := json.Marshal(results)

	// fmt.Println(string(jsonData))
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": ObjId}, updateFilter)
	if err != nil {
		c.Status(400).Send(err)
		return
	}
	response, _ := json.Marshal(res)
	c.Status(200).Send(response)
}

func main() {
	app := fiber.New()
	app.Get("/", getPost)
	app.Post("/create", createPost)
	app.Delete("/delete/:id", deletePost)
	app.Patch("/update/:id", updatePost)
	log.Fatal(app.Listen(8000))
}
