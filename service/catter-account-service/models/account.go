package models

import (
	"context"
	"github.com/danielgospodinow/Catter/service/catter-account-service/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Account struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

func CreateAccount(acc Account) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := db.GetClient().Database("catter").Collection("accounts")
	accMarshal, _ := bson.Marshal(acc)
	_, err := collection.InsertOne(ctx, accMarshal)
	return err
}

func GetAccount(id string) (Account, error) {
	var acc Account
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := db.GetClient().Database("catter").Collection("accounts")
	err := collection.FindOne(ctx, bson.D{{"id", id}}).Decode(&acc)
	return acc, err
}

func UpdateAccount(id string, acc Account) {

}

func DeleteAccount(id string) (Account, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := db.GetClient().Database("catter").Collection("accounts")
	acc, err := GetAccount(id)

	if err != nil {
		return Account{}, err
	} else {
		_, err := collection.DeleteOne(ctx, bson.D{{"id", id}})
		return acc, err
	}
}
