package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockingWebsiteChecker(url string) bool {
	return url != "site1"
}
func TestCheckWebsites(t *testing.T) {
	sites := []string{
		"site1",
		"site2",
		"site3",
	}

	got := CheckWebsites(mockingWebsiteChecker, sites)

	want := map[string]bool{
		"site1": false,
		"site2": true,
		"site3": true,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}
func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = "a url"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
