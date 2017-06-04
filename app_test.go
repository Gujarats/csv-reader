package main

import (
	"testing"
)

func TestInsertChan(t *testing.T) {
	myChan := make(chan map[string]bool, 3)

	data3 := make(map[string]bool)
	data3["data2"] = true
	data3["data4"] = true

	data2 := make(map[string]bool)
	data2["data2"] = true
	data2["data1"] = true

	data1 := make(map[string]bool)
	data1["data2"] = true
	data1["data3"] = true
	data1["data5"] = true
	data3["data1"] = true

	myChan <- data3
	myChan <- data2
	myChan <- data1
	close(myChan)

	actual := getSameValues(myChan)
	expected := []string{"data2"}
	if !eqSlice(actual, expected) {
		t.Errorf("actual = %v, expected = %v\n", actual, expected)
	}
}

func eqSlice(a, b []string) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
