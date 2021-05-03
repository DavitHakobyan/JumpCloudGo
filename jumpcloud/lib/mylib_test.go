package lib

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLibrary_AddAction(t *testing.T) {
	lib := NewLibrary()
	actionTime1 := ActionTime{"jump", 100}
	res, err := json.Marshal(actionTime1)
	if err != nil {
		fmt.Println(err)
	}

	if err := lib.addAction(string(res)); err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, len(lib.counter))
}

func TestLibrary_AddAction_empty_param(t *testing.T) {
	lib := NewLibrary()
	lib.addAction("")
	assert.Equal(t, 0, len(lib.counter))
}

func TestLibrary_GetStats_Empty(t *testing.T) {
	lib := NewLibrary()
	assert.Equal(t, "", lib.getStats(), "Result should be an empty string.")
}

func TestLibrary_GetStats_is_not_empty(t *testing.T) {
	lib := NewLibrary()
	actionTime1 := ActionTime{"jump", 100}
	res, err := json.Marshal(actionTime1)
	if err != nil {
		fmt.Println(err)
	}
	lib.addAction(string(res))

	assert.NotEqual(t, "", lib.getStats(), "Result is not empty string.")
}

func TestLibrary_AddActions_GetStats(t *testing.T) {
	// convert input parameters to json string
	json1 := "{\"action\":\"jump\",\"time\":100}"
	json2 := "{\"action\":\"run\",\"time\":75}"
	json3 := "{\"action\":\"jump\",\"time\":200}"

	// initialize the library
	lib := NewLibrary()

	// add input one
	if err := lib.addAction(json1); err != nil {
		t.Error(err)
	}

	// add input two
	if err := lib.addAction(json2); err != nil {
		t.Error(err)
	}

	// add input three
	if err := lib.addAction(json3); err != nil {
		t.Error(err)
	}

	// get stats and print
	result := lib.getStats()
	fmt.Println(result)
}

