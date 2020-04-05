package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

	"github.com/eduardohitek/gqlgen-todos/graph/generated"
	"github.com/eduardohitek/gqlgen-todos/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1").SetConnectTimeout(10 * time.Second)
	client, erro := mongo.NewClient(clientOptions)
	erro = client.Connect(context.TODO())
	if erro != nil {
		log.Println("Erro ao se conectar")
	}
	log.Println(client, input)
	res, erro := client.Database("graphql").Collection("users").InsertOne(context.TODO(), input)
	if erro != nil {
		log.Println("Erro ao inserir")
	}
	newUser := model.User{
		ID:    res.InsertedID.(primitive.ObjectID),
		Email: input.Email,
		Name:  input.Name,
		Pass:  input.Pass,
	}
	return &newUser, erro
}

func (r *queryResolver) User(ctx context.Context, limit *int) ([]*model.User, error) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1").SetConnectTimeout(10 * time.Second)
	client, erro := mongo.NewClient(clientOptions)
	erro = client.Connect(context.TODO())
	if erro != nil {
		log.Println("Erro ao se conectar")
	}
	log.Println(client, limit)
	findOp := options.Find()
	if limit != nil {
		limit64 := int64(*limit)
		findOp.SetLimit(limit64)
	}
	var ret []*model.User
	cursor, _ := client.Database("graphql").Collection("users").Find(context.TODO(), bson.D{}, findOp)
	for cursor.Next(context.TODO()) {
		var user model.User
		cursor.Decode(&user)
		ret = append(ret, &user)
	}
	return ret, nil
}

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	log.Println(obj)
	return obj.ID.Hex(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
