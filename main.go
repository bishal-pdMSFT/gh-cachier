package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/cli/go-gh"
	flag "github.com/spf13/pflag"
)

type RepoCacheUsage struct {
	FullName                string `json:"full_name"`
	ActiveCachesSizeInBytes int64  `json:"active_caches_size_in_bytes"`
	ActiveCachesCount       int32  `json:"active_caches_count"`
}

type OrgCacheUsage struct {
	TotalCachesSize   int64 `json:"total_active_caches_size_in_bytes"`
	ActiveCachesCount int32 `json:"total_active_caches_count"`
}

type OrgCacheUsageByRepo struct {
	TotalCount      int32            `json:"total_count"`
	RepoCacheUsages []RepoCacheUsage `json:"repository_cache_usages"`
}

type Scope interface {
	getCacheUsage()
	getCacheUsageEndpoint() string
}

type RepoScope struct {
	Name       string
	CacheUsage RepoCacheUsage
}

type OrgScope struct {
	Name             string
	IsUsageByRepo    bool
	CacheUsage       OrgCacheUsage
	CacheUsageByRepo OrgCacheUsageByRepo
}

func (r *RepoScope) getCacheUsage() {
	fmt.Println("Getting Actions cache usage for repo: " + r.Name)
	fetchCacheUsageFromGitHub(r.getCacheUsageEndpoint(), &r.CacheUsage)
	prettyPrint(r.CacheUsage)
}

func (r *RepoScope) getCacheUsageEndpoint() string {
	return fmt.Sprintf("repos/%s/actions/cache/usage", r.Name)
}

func (o *OrgScope) getCacheUsage() {
	fmt.Println("Getting Actions cache usage for organization: " + o.Name)

	if o.IsUsageByRepo {
		fetchCacheUsageFromGitHub(o.getCacheUsageEndpoint(), &o.CacheUsageByRepo)
		prettyPrint(o.CacheUsageByRepo)
		return
	}

	fetchCacheUsageFromGitHub(o.getCacheUsageEndpoint(), &o.CacheUsage)
	prettyPrint(o.CacheUsage)
}

func (o *OrgScope) getCacheUsageEndpoint() string {
	if o.IsUsageByRepo {
		return fmt.Sprintf("orgs/%s/actions/cache/usage-by-repository", o.Name)
	}

	return fmt.Sprintf("orgs/%s/actions/cache/usage", o.Name)
}

func main() {
	var repo *string = flag.String("repo", "", "Name of the repository to fetch cache usage for")
	var org *string = flag.String("org", "", "Name of the organization to fetch cache usage for")
	//var ent *string = flag.String("enterprise", "", "Name of the enterprise to fetch cache usage for")
	var byRepo *bool = flag.BoolP("by-repo", "R", false, "Flag to fetch cache usage per repo for an org. To be used with `--org` option")

	flag.Parse()
	args := flag.Args()

	if args[0] == "usage" {
		if *repo != "" {
			var repoScope RepoScope = RepoScope{Name: *repo}
			repoScope.getCacheUsage()
		} else if *org != "" {
			var orgScope OrgScope = OrgScope{Name: *org, IsUsageByRepo: *byRepo}
			orgScope.getCacheUsage()
		}
	}

}

func fetchCacheUsageFromGitHub(endpoint string, response interface{}) {
	client, err := gh.RESTClient(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Get(endpoint, &response)
	if err != nil {
		log.Fatal(err)
	}
}

func prettyPrint(i interface{}) {
	fmt.Println()
	s, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(s))
	fmt.Println()
}
