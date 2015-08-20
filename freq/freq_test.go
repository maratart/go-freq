package freq

import "testing"

func TestSortMapByValue(t *testing.T) {
	m := map[string]int{"s1": 2, "s2": 1, "s3": 6, "s4": 5, "s5": 4}
	r := []Pair{{"s2", 1}, {"s1", 2}, {"s5", 4}, {"s4", 5}, {"s3", 6}}
	s := sortMapByValue(m)
	for i := range s {
		if len(s) != len(r) {
			t.Error("Lens of result not equal.", "\nExpected:", len(r), "\nGot:", len(s))
			break
		} else if s[i] != r[i] {
			t.Error(r[i], "!=", s[i], "\n", r, "\n", s)
		}
	}
}
