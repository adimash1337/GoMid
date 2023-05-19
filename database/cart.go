package database

import (
	"context"
	"log"

	"github.com/adimash1337/GoMid/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddProductToCart(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	searchfromdb, err := prodCollection.Find(ctx, bson.M{"_id": productID})
	if err != nil {
		log.Println(err)
		return err
	}
	var productcart []models.ProductUser
	err = searchfromdb.All(ctx, &productcart)
	if err != nil {
		log.Println(err)
		return err
	}

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "usercart", Value: bson.D{{Key: "$each", Value: productcart}}}}}}
	_, err = userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func RemoveCartItem(ctx context.Context, prodCollection, userCollection *mongo.Collection, productID primitive.ObjectID, userID string) error {
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	update := bson.M{"$pull": bson.M{"usercart": bson.M{"_id": productID}}}
	_, err = userCollection.UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil

}
