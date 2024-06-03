package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestEnglish(t *testing.T) {
	name := "Ben"
	language := "english"
	expected := "Hello Ben"
	actual, err := Hello(name, language)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}



// package main

// import "testing"

// func TestEnglish(t *testing.T) {
// 	name := "Ben"
// 	language := "english"
// 	expected := "Hello Ben"
// 	actual, err := Hello(name, language)

// 	if err != nil {
// 		t.Errorf("Should not produce an error")
// 	}

// 	if expected != actual {
// 		t.Errorf("Result was incorrect, got %s, want: %s", expected, actual)
// 	}
// }
