package gravatar_test

import "testing"
import "github.com/nowk/go-gravatar"

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestHash(t *testing.T) {
	got := gravatar.Hash("yung.kwon@damncarousel.com")
	exp := "abc5e4c1bc544767ddace4f3e28cd6a4"
	if got != exp {
		t.Errorf("Expected %s, got %s", exp, got)
	}
}

func TestLookup(t *testing.T) {
	grav := gravatar.Lookup("yung.kwon@damncarousel.com")

	got := grav.URL.String()
	exp := "http://www.gravatar.com/avatar/abc5e4c1bc544767ddace4f3e28cd6a4"
	if got != exp {
		t.Errorf("Expected %s, got %s", exp, got)
	}
}
