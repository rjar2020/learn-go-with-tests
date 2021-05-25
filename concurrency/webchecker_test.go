package concurrency

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func mockWebsiteCheker(url string) bool {
	if url == "watt://furhurterwe.geds" {
		return false
	}
	return true
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestChecWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"watt://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"watt://furhurterwe.geds":    false,
	}

	got := CheckWebsites(mockWebsiteCheker, websites)

	require.Equal(t, want, got)
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 10)

	for i := 0; i < len(urls); i++ {
		urls[i] = "An url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}

func ExampleCheckWebsites() {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"watt://furhurterwe.geds",
	}
	got := CheckWebsites(mockWebsiteCheker, websites)

	for _, url := range websites {
		fmt.Printf("%s -> %t\n", url, got[url])
	}
	/* Output:
http://google.com -> true
http://blog.gypsydave5.com -> true
watt://furhurterwe.geds -> false*/
}
