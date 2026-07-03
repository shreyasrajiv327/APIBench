package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("========== Field Selection Experiment ==========\n")

	// ------------------------------------------------------------
	// 1. Publish via REST
	// ------------------------------------------------------------

	restPublishBody := []byte(`{"payload":"REST Research"}`)

	resp, err := http.Post(
		"http://localhost:8080/messages",
		"application/json",
		bytes.NewBuffer(restPublishBody),
	)
	if err != nil {
		panic(err)
	}

	var restPublished struct {
		ID string `json:"id"`
	}

	json.NewDecoder(resp.Body).Decode(&restPublished)
	resp.Body.Close()

	// ------------------------------------------------------------
	// 2. Fetch via REST
	// ------------------------------------------------------------

	resp, err = http.Get(
		"http://localhost:8080/messages/" + restPublished.ID,
	)
	if err != nil {
		panic(err)
	}

	restBytes, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	// ------------------------------------------------------------
	// 3. Publish via GraphQL
	// ------------------------------------------------------------

	publishQuery := map[string]string{
		"query": `
		mutation {
			publish(payload: "GraphQL Research") {
				id
			}
		}`,
	}

	body, _ := json.Marshal(publishQuery)

	resp, err = http.Post(
		"http://localhost:8081/graphql",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		panic(err)
	}

	var gqlPublish struct {
		Data struct {
			Publish struct {
				ID string `json:"id"`
			} `json:"publish"`
		} `json:"data"`
	}

	json.NewDecoder(resp.Body).Decode(&gqlPublish)
	resp.Body.Close()

	// ------------------------------------------------------------
	// 4. GraphQL -> status only
	// ------------------------------------------------------------

	statusQuery := map[string]string{
		"query": fmt.Sprintf(`
		query {
			message(id: "%s") {
				status
			}
		}`, gqlPublish.Data.Publish.ID),
	}

	body, _ = json.Marshal(statusQuery)

	resp, err = http.Post(
		"http://localhost:8081/graphql",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		panic(err)
	}

	statusBytes, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	// ------------------------------------------------------------
	// 5. GraphQL -> all fields
	// ------------------------------------------------------------

	allQuery := map[string]string{
		"query": fmt.Sprintf(`
		query {
			message(id: "%s") {
				id
				payload
				status
				createdAt
			}
		}`, gqlPublish.Data.Publish.ID),
	}

	body, _ = json.Marshal(allQuery)

	resp, err = http.Post(
		"http://localhost:8081/graphql",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		panic(err)
	}

	allBytes, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	// ------------------------------------------------------------
	// Results
	// ------------------------------------------------------------

	fmt.Println("REST Response")
	fmt.Println(string(restBytes))
	fmt.Printf("Bytes: %d\n\n", len(restBytes))

	fmt.Println("GraphQL (status only)")
	fmt.Println(string(statusBytes))
	fmt.Printf("Bytes: %d\n\n", len(statusBytes))

	fmt.Println("GraphQL (all fields)")
	fmt.Println(string(allBytes))
	fmt.Printf("Bytes: %d\n", len(allBytes))
}