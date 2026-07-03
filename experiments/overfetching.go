package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	// Publish TWO messages so REST and GraphQL each have one to consume.
	messages := []string{
		`{"payload":"Hello REST"}`,
		`{"payload":"Hello GraphQL"}`,
	}

	for _, msg := range messages {
		_, err := http.Post(
			"http://localhost:8080/messages",
			"application/json",
			bytes.NewBuffer([]byte(msg)),
		)
		if err != nil {
			panic(err)
		}
	}

	// ================= REST =================

	restResp, err := http.Get("http://localhost:8080/messages/next")
	if err != nil {
		panic(err)
	}
	defer restResp.Body.Close()

	restBytes, err := io.ReadAll(restResp.Body)
	if err != nil {
		panic(err)
	}

	// ================= GraphQL =================

	query := map[string]string{
		"query": `
		{
			nextMessage {
				status
			}
		}
		`,
	}

	queryBytes, _ := json.Marshal(query)

	graphqlResp, err := http.Post(
		"http://localhost:8081/graphql",
		"application/json",
		bytes.NewBuffer(queryBytes),
	)
	if err != nil {
		panic(err)
	}
	defer graphqlResp.Body.Close()

	graphqlBytes, err := io.ReadAll(graphqlResp.Body)
	if err != nil {
		panic(err)
	}

	// ================= Results =================

	fmt.Println("=========== REST ===========")
	fmt.Println(string(restBytes))
	fmt.Printf("Bytes: %d\n\n", len(restBytes))

	fmt.Println("========= GraphQL ==========")
	fmt.Println(string(graphqlBytes))
	fmt.Printf("Bytes: %d\n", len(graphqlBytes))
}