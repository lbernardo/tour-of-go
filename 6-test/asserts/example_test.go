package asserts

import "testing"

func TestSum(t *testing.T) {
	result := Sum(10, 10)
	if result != 20 {
		t.Errorf("Was expected 20 and return %v", result)
	}
}

func TestMult(t *testing.T) {
	result := Mult(5, 5)
	if result != 25 {
		t.Errorf("Was expected 25 and return %v", result)
	}
}

func TestConvert(t *testing.T) {
	t.Run("With error", func(t *testing.T) {
		_, err := Convert("Number 42")
		if err == nil {
			t.Error("Was expected error but return nil")
		}
	})
	t.Run("Without error", func(t *testing.T) {
		i, err := Convert("42")
		if err != nil {
			t.Errorf("Was expected nil but return %v ", err)
		}
		if i != 42 {
			t.Errorf("Was expected 42 but return %v", i)
		}
	})
}
