# Learn Golang through TDD
Forked from [Learn Go with Tests by Chris James]

## Prerequisites (besides Git, of course)
- Go Tools

## Testing
From your CI or the terminal, for running the developer tests in a particular package execute the following:

```bash
go test github.com/rjar2020/learn-go-with-tests/ints
```
If you want to have a more verbose output, including coverage, test and examples information, execute the following

```bash
go test -v -cover github.com/rjar2020/learn-go-with-tests/iteration
```

## Benchmarks
From your CI or the terminal, for executing benchmarks in a particular package use the following command:

```bash
go test -bench=. github.com/rjar2020/learn-go-with-tests/iteration
```

[Learn Go with Tests by Chris James]: https://quii.gitbook.io/learn-go-with-tests/

## Docs
From your terminal, to see the documentation execute the following steps from the project's root:

```bash
godoc -http=:6060
#You should see the followin message
# using module mode; GOMOD=/Users/<your-user>/repo/go/src/github.com/rjar2020/learn-go-with-tests/go.mod
```
Then go to http://localhost:6060/pkg/ and search for Third Party section, your packages should be there.

[Learn Go with Tests by Chris James]: https://quii.gitbook.io/learn-go-with-tests/
