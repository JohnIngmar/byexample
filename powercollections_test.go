package byexample

import (
	"testing"
)

func TestPowerMap(t *testing.T) {
	var map1 = NewPowerMap()

	set1 := map[string]int{"a": 1, "b": 2}

	set2 := map[string]int{"b": 2, "a": 1}

	set3 := map[string]int{"b": 2}

	set3["a"] = 1

	value := "val1"
	map1.Set(set1, value)

	val1, _ := map1.Get(set1)
	if val1 != value {
		t.Fatalf("bad, Expected values to be the same: %s=%s", val1, value)
	}

	val2, _ := map1.Get(set2)
	if val1 != val2 {
		t.Fatalf("bad, Expected values to be the same: %s=%s", val1, val2)
	}

	map1.Delete(set3)

	val3, _ := map1.Get(set1)

	if val3 != nil {
		t.Fatalf("bad, Expected value to be removed: %s", val3)
	}
}
