package repository

import (
	"ccg/models"
	"ccg/storage"
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

type CCGRepository interface {
	CreateCCG(ctx echo.Context) error
	GetCGG(ctx echo.Context, id int) error
	GetAllCCG(ctx echo.Context)
}

func InitRepo() {
	db = storage.MongoClient.Database("ccg-server")
}

func CreateCCG(ccg models.CCG) (models.CCG, error) {

	_, err := db.Collection("ccg").InsertOne(context.TODO(), ccg)
	if err != nil {
		fmt.Errorf("sukaaa", err)
		return models.CCG{}, err
	}
	return ccg, nil
}

func EditCCG(ccg models.CCG) error {
	return nil
}

func GetCCG() ([]models.CCG, error) {
	var ccgs []models.CCG

	collection := db.Collection("ccg")

	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var ccg models.CCG
		if err := cursor.Decode(&ccg); err != nil {
			return nil, err
		}
		ccgs = append(ccgs, ccg)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return ccgs, nil
}
