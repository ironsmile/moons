package moons_test

import (
	"fmt"
	"testing"

	"github.com/ironsmile/moons"
)

func TestBodyFromString(t *testing.T) {
	body := moons.BodyFromString("NePtUnE")
	if body != moons.Neptune {
		t.Errorf("Expected %s but got %s", moons.Neptune, body)
	}
}

func ExampleCount() {
	body := moons.BodyFromString("earth")
	count, err := moons.Count(body)
	if err != nil {
		fmt.Printf("Could not get Earth moons: %s\n", err)
	}
	fmt.Printf("%s has %d moon, right?\n", body, count)
	// Output:
	// Earth has 1 moon, right?
}
