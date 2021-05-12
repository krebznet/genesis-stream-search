package mongo

import (
	"context"
	"fmt"
	"genesis-stream-search/pkg/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Handler for operating the mongo database
type Handler struct {
	client   *mongo.Client  // the connected mongo client
	settings *config.Config // server configuration
}

// NewHandler returns a new Handler for mongo database
func NewHandler(settings *config.Config) *Handler {
	return &Handler{
		settings: settings,
	}
}

// Start the handler for mongo database
func (s *Handler) Start() {
	// mongodb://localhost:27017
	uri := fmt.Sprintf("mongodb://%s:%s", s.settings.DB.Host, s.settings.DB.Port)

	options := options.Client().ApplyURI(uri).SetAuth(options.Credential{
		Username: s.settings.DB.User,
		Password: s.settings.DB.Password,
	})
	// new a mongo client
	client, err := mongo.NewClient(options)
	if err != nil {
		panic(fmt.Sprintf("new mongo client: %v", err))
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	// connect the mongo
	if err := client.Connect(ctx); err != nil {
		panic(fmt.Sprintf("connect mongo: %v", err))
	}
	// ping the mongo
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(fmt.Sprintf("ping mongo: %v", err))
	}
	s.client = client

	log.Printf("ping【%s】is done", uri)
}

// Stop the handler for mongo database
func (s *Handler) Stop() {
	if s.client != nil {
		if err := s.client.Disconnect(context.Background()); err != nil {
			log.Printf("disconnect mongo: %v", err)
		}
	}
}

// Database returns a special database
func (s *Handler) Database() *mongo.Database {
	return s.client.Database(s.settings.DB.Name)
}

// Collection returns a special collection
func (s *Handler) Collection(name string) *mongo.Collection {
	return s.Database().Collection(name)
}
