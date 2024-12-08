package db

import (
	"reflect"
	"testing"
)

func TestRootSuggestion(t *testing.T) {
	d := RootSuggestion("عل", 8)
	for _, e := range d {
		t.Logf("%q", e)
	}
}

func TestSearchByRoot(t *testing.T) {
	d, err := SearchByRoot("علم", 10)
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range d {
		t.Logf("%#v", e)
	}
}

func TestSearchByTxt(t *testing.T) {
	limit := 50
	d1, err := SearchByTxt("علم", limit, "")
	if err != nil {
		t.Fatal(err)
	}

	d2, err := SearchByTxt("علم", limit, "")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(d1, d2) {
		t.Fatal("res 1 is not as res 2")
	}
	d2, err = SearchByTxt("علم", limit, "")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(d1, d2) {
		t.Fatal("res 1 is not as res 2")
	}
}
