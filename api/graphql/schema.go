package graphql

import "github.com/graphql-go/graphql"

var MessageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Message",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"payload": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
	},
})