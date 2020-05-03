package models

import (
	"context"
	"time"

	"github.com/danielgospodinow/Catter/service/catter-account-service/db"
	"go.mongodb.org/mongo-driver/bson"
)

// Account is the model of an account object.
type Account struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

// CreateAccount creates an Account instance.
func CreateAccount(acc Account) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := db.GetClient().Database("catter").Collection("accounts")
	accMarshal, _ := bson.Marshal(acc)
	_, err := collection.InsertOne(ctx, accMarshal)
	return err
}

// GetAccount retrieves an Account instance by a given id.
func GetAccount(id string) (Account, error) {
	var acc Account
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := db.GetClient().Database("catter").Collection("accounts")
	err := collection.FindOne(ctx, bson.D{{"id", id}}).Decode(&acc)
	return acc, err
}

// UpdateAccount updates an Account instance by a given id.
func UpdateAccount(id string, acc Account) {

}

// DeleteAccount deletes an Account instance by a given id.
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
