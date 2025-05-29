package db

import (
	"blog-api/models"
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var client *mongo.Client
var collection *mongo.Collection

func Init(ctx context.Context) error {
	// Sets the version of the Stable API on the client.
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URI")).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(opts)
	if err != nil {
		return err
	}

	// Send a ping to confirm a successful connection
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// Setting the collection.
	collection = client.Database("blogsdb").Collection("blogs")

	// Creating an index to search terms in the title, content and categories of a blog.
	indexModel := mongo.IndexModel{
		Keys: bson.D{
			{Key: "title", Value: "text"},
			{Key: "content", Value: "text"},
			{Key: "category", Value: "text"},
			{Key: "tags", Value: "text"},
		},
	}
	_, err = collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		return err
	}

	return nil
}

func Disconnect() {
	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
}

func InsertBlog(ctx context.Context, blog models.Blog) error {
	_, err := collection.InsertOne(ctx, blog)
	return err
}

func UpdateBlog(ctx context.Context, id string, newBlog models.Blog) error {
	fmt.Println(id)
	_, err := collection.ReplaceOne(
		ctx,
		// Filter.
		bson.D{{Key: "id", Value: id}},
		// Updated Values.
		newBlog,
	)

	return err
}

func GetAllBlogs(ctx context.Context) ([]models.Blog, error) {
	emptyFilter := bson.M{}
	cursor, err := collection.Find(ctx, emptyFilter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var blogs []models.Blog
	if err = cursor.All(ctx, &blogs); err != nil {
		return nil, err
	}

	return blogs, nil
}

func GetBlogByID(ctx context.Context, id string) (*models.Blog, error) {
	filter := bson.D{{Key: "id", Value: id}}
	var blog models.Blog
	err := collection.FindOne(ctx, filter).Decode(&blog)
	if err != nil {
		return nil, err
	}

	return &blog, err
}

func GetBlogsBySearch(ctx context.Context, term string) ([]models.Blog, error) {
	filter := bson.M{"$text": bson.M{"$search": term}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var matchingBlogs []models.Blog
	if err = cursor.All(ctx, &matchingBlogs); err != nil {
		return nil, err
	}

	return matchingBlogs, nil
}

func DeleteBlogById(ctx context.Context, id string) error {
	filter := bson.D{{Key: "id", Value: id}}
	_, err := collection.DeleteOne(ctx, filter)
	
	return err
}