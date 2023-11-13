package concurrency

import (
	"reflect"
	"testing"
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
		t.Errorf("got %v, want %v", got, want)
	}
}
