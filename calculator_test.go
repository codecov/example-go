package calculator

import "testing"

func TestAdd(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
		e string
	}{
		{1, 2, 3, ""},
		{1.0, 2.0, 3.0, ""},
		{0, 2.0, 2.0, ""},
		{2.0, 0, 2.0, ""},
		{-4, 2.0, -2.0, ""},
	}

	for _, table := range tables {
		if total, err := Add(table.x, table.y); err != nil && err.Error() != table.e {
			t.Errorf("Add of (%d+%d) produced wrong error, got %v, want %v.", table.x, table.y, err.Error(), table.e)
		} else if total != table.n {
			t.Errorf("Add of (%d+%d) produced wrong result, got %d, want %d.", table.x, table.y, total, table.n)
		}
	}
}

func TestSubtract(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
		e string
	}{
		{1, 2, -1.0, ""},
		{2, 1, 1.0, ""},
		{1.0, 2.0, -1.0, ""},
		{0, 2.0, -2.0, ""},
		{2.0, 0, 2.0, ""},
		{-4, 2.0, -6.0, ""},
	}

	for _, table := range tables {
		if total, err := Subtract(table.x, table.y); err != nil && err.Error() != table.e {
			t.Errorf("Subtract of (%d-%d) produced wrong error, got %v, want %v.", table.x, table.y, err.Error(), table.e)
		} else if total != table.n {
			t.Errorf("Subtract of (%d-%d) produced wrong result, got %d, want %d.", table.x, table.y, total, table.n)
		}
	}
}

func TestMultiply(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
		e string
	}{
		{1, 2, 2.0, ""},
		{1.0, 2.0, 2.0, ""},
		{0, 2.0, 0, ""},
		{2.0, 0, 0, ""},
		{-4, 2.0, -8.0, ""},
	}

	for _, table := range tables {
		if total, err := Multiply(table.x, table.y); err != nil && err.Error() != table.e {
			t.Errorf("Multiply of (%d*%d) produced wrong error, got %v, want %v.", table.x, table.y, err.Error(), table.e)
		} else if total != table.n {
			t.Errorf("Multiply of (%d*%d) produced wrong result, got %d, want %d.", table.x, table.y, total, table.n)
		}
	}
}

func TestDivide(t *testing.T) {
	tables := []struct {
		x int
		y int
		n float64
		e string
	}{
		{1, 2, 0.5, ""},
		{1.0, 2.0, 0.5, ""},
		{0, 2.0, 0, ""},
		{-4, 2.0, -2.0, ""},
		// {2.0, 0, 0, "Cannot divide by 0"},
	}

	for _, table := range tables {
		if total, err := Divide(table.x, table.y); err != nil && err.Error() != table.e {
			t.Errorf("Divide of (%d/%d) produced wrong error, got %v, want %v.", table.x, table.y, err.Error(), table.e)
		} else if total != table.n {
			t.Errorf("Divide of (%d/%d) produced wrong result, got %f, want %f.", table.x, table.y, total, table.n)
		}
	}
}
