package main

import (
	"fmt"
	"log"

	"github.com/cli/go-gh"

	flag "github.com/spf13/pflag"
)

func main() {
	// fmt.Println("hi world, this is the gh-cachier extension!")
	// client, err := gh.RESTClient(nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// response := struct {Login string}{}
	// err = client.Get("user", &response)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("running as %s\n", response.Login)

	var repo *string = flag.String("repo", "", "Name of the repository to fetch cache usage for")
	flag.Parse()

	args := flag.Args()
	// for i, arg := range args {
	// 	fmt.Printf("arg %d - %s", i, arg)
	// }
	//
	// fmt.Println("repo: " + *repo)

	if args[0] == "usage" {
		getRepoCacheUsage(repo)
	}

}

// usage <owner/repo>
func getRepoCacheUsage(repo *string) {
	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	response := struct {
		full_name                   string
		active_caches_size_in_bytes int64
		active_caches_count         int32
	}{}

	fmt.Println("repo: " + *repo)

	endpoint := fmt.Sprintf("repos/%s/actions/cache/usage", *repo)
	fmt.Println("endpoint: " + endpoint)
	err = client.Get(endpoint, &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
}

// For more examples of using go-gh, see:
// https://github.com/cli/go-gh/blob/trunk/example_gh_test.go
