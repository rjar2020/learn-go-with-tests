# Learn Golang through TDD
Forked from [Learn Go with Tests by Chris James]

## Prerequisites (besides Git, of course)
- Go Tools

## Testing
From your CI or the terminal, for running the developer tests in a particular package execute the following:

```bash
go test github.com/rjar2020/learn-go-with-tests/ints
```
If you want to have a more verbose output, including coverage, test and examples information, execute the following:

```bash
go test -v -cover github.com/rjar2020/learn-go-with-tests/iteration
```
If you want to run all the tests in every package at once, execute the following from the repo root:
```bash
go test ./...
```
To get some help finding out race conditions in your implementation, you could use the [Go Race detector] as follows:
```bash
go test ./... -race
```

## Benchmarks
From your CI or the terminal, for executing benchmarks in a particular package use the following command:

```bash
go test -bench=. github.com/rjar2020/learn-go-with-tests/iteration
```

## Docs
From your terminal, to see the documentation execute the following steps from the project's root:

```bash
godoc -http=:6060
#You should see the followin message
# using module mode; GOMOD=/Users/<your-user>/repo/go/src/github.com/rjar2020/learn-go-with-tests/go.mod
```
Then go to http://localhost:6060/pkg/ and search for Third Party section, your packages should be there.

## More to explore
- Include a linter for ensuring exported methods are documented (like is reported by VS Code but in the CI, if it exists)
- Parallel testing, if not in the tutorial
- Implement a DB connection pool
- Implement a web crawler 
- Implement a scoring board algorithm
- Generate sample data in slices/maps for benchmarks (see #Dictionary.Add benchmark)
- Refactor (or create v2) mocking package to use Testify mocks

[Learn Go with Tests by Chris James]: https://quii.gitbook.io/learn-go-with-tests/
[Go Race detector]: https://blog.golang.org/race-detector