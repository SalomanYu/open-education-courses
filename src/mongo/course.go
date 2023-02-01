package mongo

import (
	"github.com/SalomanYu/open-education-courses/src/models"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collection *mongo.Collection
	ctx =  context.TODO()

)

func init(){
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	check_err(err)
	
	err = client.Ping(ctx, nil)
	check_err(err)

	collection = client.Database("open_education").Collection("courses")
}

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func AddCourse(course *models.Course) error{
	fmt.Println("added course:", course.Title)
	_, err := collection.InsertOne(ctx, course)
	return err
}