# gh-cachier
CLI extension for managing Actions cache. Lets you see the usage of caching at various scopes:
- Repository
- Organization, and per repository in an Organization
- [future] Enterprise 

This CLI is built on top of [Actions cache REST APIs](https://docs.github.com/en/rest/actions/cache) and exposes similar functionality.

## Install
Run following `gh cli` command:
```
gh extension install bishal-pdmsft/gh-cachier
```

## Usage
```
  -R, --by-repo --org   Flag to fetch cache usage per repo for an org. To be used with --org option
      --org string      Name of the organization to fetch cache usage for
      --repo string     Name of the repository to fetch cache usage for
```

#### Get cache usage in a repo
```
gh cachier usage --repo <owner/repo> 
```

#### Get cache usage in an organization
```
gh cachier usage --org <org> 
```

#### Get per repo cache usage in an organization
```
gh cachier usage --org <org> -R
```
