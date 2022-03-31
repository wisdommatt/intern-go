package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/wisdommatt/intern-go/backend/graph"
	"github.com/wisdommatt/intern-go/backend/graph/generated"
	"github.com/wisdommatt/intern-go/backend/internal/accounts"
	"google.golang.org/api/option"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	firestoreClient := mustSetupFirestoreClient()
	defer firestoreClient.Close()

	accountService := accounts.NewService(firestoreClient)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		AccountService: accountService,
	}}))

	http.Handle("/graphql/playground", playground.Handler("GraphQL playground", "/graphql/query"))
	http.Handle("/graphql/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func mustSetupFirestoreClient() *firestore.Client {
	ctx := context.Background()
	sa := option.WithCredentialsFile("firebase-service-account.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
