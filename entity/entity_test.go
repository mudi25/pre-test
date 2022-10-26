package entity_test

import (
	"testing"

	"github.com/mudi25/pretest/entity"
)

func TestMain(m *testing.M) {
	m.Run()
}
func TestIsPalindrom(t *testing.T) {
	data := []struct {
		s           string
		IsPalindrom bool
	}{
		{s: "apa", IsPalindrom: true},
		{s: "1001", IsPalindrom: true},
		{s: "foo", IsPalindrom: false},
		{s: "bar", IsPalindrom: false},
	}
	for _, d := range data {
		if result := entity.IsPalindrom(d.s); result != d.IsPalindrom {
			t.Errorf("expect result %t actual %t", d.IsPalindrom, result)
		}
	}

}
