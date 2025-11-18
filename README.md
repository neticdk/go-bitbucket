

# go-bitbucket

go-bitbucket is a Go client library for accessing the [Bitbucket REST API](https://developer.atlassian.com/server/bitbucket/rest).

## Features

*   Comprehensive coverage of the Bitbucket Server API.
*   Easy-to-use interface for managing projects, repositories, pull requests, users, and more.
*   Supports authentication with basic authentication.
*   Provides a mock server for testing.
*   Includes webhook parsing and validation.
   
## Installation

```bash
go get github.com/neticdk/go-bitbucket
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/neticdk/go-bitbucket/bitbucket"
)

func main() {
	// If using basic authentication, set the credentials - could also use standard OAuth 2 library.
	hc := (&bitbucket.BasicAuthTransport{
		Username: "your_username",
		Password: "your_password",
	}).Client()

    // Create a new Bitbucket client.  Replace with your Bitbucket URL and credentials.
	client, err := bitbucket.NewClient("https://your-bitbucket-server.com", hc)
	if err != nil {
		log.Fatal(err)
	}

	// Create a context.
	ctx := context.Background()

    // Example: List repositories.
	opts := &bitbucket.RepositorySearchOptions{Permission: bitbucket.PermissionRepoWrite, ListOptions: bitbucket.ListOptions{Limit: 10}}
	all := make([]*model.Repo, 0)
	for {
		repos, resp, err := client.Projects.SearchRepositories(ctx, opts)
		if err != nil {
            log.Fatal(err)
		}
		for _, r := range repos {
			all = append(all, convertRepo(r, perms, ""))
		}
		if resp.LastPage {
			break
		}
		opts.Start = resp.NextPageStart
	}

	fmt.Println("Repositories:")
	for _, repo := range all {
		fmt.Printf("- %s (%s)\n", repo.Name, repo.Key)
	}
}
```

Replace placeholders:

* `https://your-bitbucket-server.com`: Your Bitbucket Server URL.
* `your_username`: Your Bitbucket username.
* `your_password`: Your Bitbucket password.

## Mock Server

The library includes a mock server for testing your applications without needing a real Bitbucket Server instance.

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/neticdk/go-bitbucket/bitbucket"
	"github.com/neticdk/go-bitbucket/mock"
)

func main() {
	// Create a mock server.
	mockServer := mock.NewMockServer(
		mock.WithRequestMatch(mock.ListProjects, bitbucket.ProjectList{
			Projects: []*bitbucket.Project{
				{
					ID:   1,
					Key:  "PRJ1",
					Name: "Project 1",
				},
			},
		}),
	)
	defer mockServer.Close()

	// Create a client using the mock server's URL.
	client, err := bitbucket.NewClient(mockServer.URL, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Use the client as you would with a real Bitbucket Server.
	ctx := context.Background()
	projects, _, err := client.Projects.ListProjects(ctx, &bitbucket.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Projects (from mock server):")
	for _, project := range projects {
		fmt.Printf("- %s (%s)\n", project.Name, project.Key)
	}
}
```

## Webhooks

The library provides functions for parsing and validating Bitbucket Server webhook payloads.

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/neticdk/go-bitbucket/bitbucket"
)

func main() {
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	// Replace with your webhook secret.
	secret := []byte("your_webhook_secret")

	event, payload, err := bitbucket.ParsePayload(r, secret)
	if err != nil {
		log.Printf("Error parsing webhook: %v", err)
		http.Error(w, "Invalid webhook", http.StatusBadRequest)
		return
	}

	switch ev := event.(type) {
	case *bitbucket.RepositoryPushEvent:
		fmt.Printf("Repository push event for %s\n", ev.Repository.Slug)
		// Process the push event...
	case *bitbucket.PullRequestEvent:
		fmt.Printf("Pull request event: %s\n", ev.EventKey)
		// Process the pull request event...
	default:
		fmt.Printf("Unhandled event type: %T\n", ev)
		fmt.Printf("Payload: %s\n", string(payload))
	}

	w.WriteHeader(http.StatusOK)
}
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request on GitHub.
