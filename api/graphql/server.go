package graphql

import (
	"net/http"
	"time"

	gql "github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func NewHandler(resolver *Resolver) http.Handler {

	// ---------------- Mutation ----------------

	mutation := gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"publish": &gql.Field{
				Type: MessageType,
				Args: gql.FieldConfigArgument{
					"payload": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.String),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {

					payload := p.Args["payload"].(string)

					msg, err := resolver.Queue.Publish([]byte(payload))
					if err != nil {
						return nil, err
					}

					return map[string]interface{}{
						"id":        msg.ID,
						"payload":   string(msg.Payload),
						"status":    string(msg.Status),
						"createdAt": msg.CreatedAt.Format(time.RFC3339),
					}, nil
				},
			},
		},
		"ack": &gql.Field{
	Type: gql.Boolean,
	Args: gql.FieldConfigArgument{
		"id": &gql.ArgumentConfig{
			Type: gql.NewNonNull(gql.String),
		},
	},
	Resolve: func(p gql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(string)

		if err := resolver.Queue.Ack(id); err != nil {
			return false, err
		}

		return true, nil
	},
},
"nack": &gql.Field{
	Type: gql.Boolean,
	Args: gql.FieldConfigArgument{
		"id": &gql.ArgumentConfig{
			Type: gql.NewNonNull(gql.String),
		},
	},
	Resolve: func(p gql.ResolveParams) (interface{}, error) {
		id := p.Args["id"].(string)

		if err := resolver.Queue.Nack(id); err != nil {
			return false, err
		}

		return true, nil
	},
},
	})

	// ---------------- Query ----------------

	query := gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"nextMessage": &gql.Field{
				Type: MessageType,
				Resolve: func(p gql.ResolveParams) (interface{}, error) {

					msg, err := resolver.Queue.Poll()
					if err != nil {
						return nil, err
					}

					return ToMessageResponse(msg), nil
				},
			},
		},
	})

	schema, err := gql.NewSchema(gql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})

	if err != nil {
		panic(err)
	}

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}