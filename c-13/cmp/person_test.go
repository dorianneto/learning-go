package person

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_NewPerson(t *testing.T) {
	expected := &Person{
		Name: "Dennis",
		Age:  37,
	}
	result := NewPerson("Dennis", 37)

	comparer := cmp.Comparer(func(x, y Person) bool {
		return x.Name == y.Name && x.Age == y.Age
	})

	if diff := cmp.Diff(expected, result, comparer); diff != "" {
		t.Error(diff)
	}
}
