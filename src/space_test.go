package parkingLot

import "testing"

func TestNewSpace(t *testing.T) {
	s := NewSpace(false, Car{"reg", "white"})

	if s == nil {
		t.Fatalf("expected %v actual %s", s, "nil")
	}
}

func TestSpace_AssignSpaceToCar(t *testing.T) {
	car := Car{"reg", "white"}
	s := NewSpace(false, car)
	s.AssignSpaceToCar(car)

	if !s.Occupied {
		t.Fatalf("expected %v actual %v", true, s.Occupied)
	}
}

func TestSpace_FreeSpace(t *testing.T) {
	car := Car{"reg", "white"}
	s := NewSpace(true, car)
	s.FreeSpace()

	if s.Occupied {
		t.Fatalf("expected %v actual %v", true, s.Occupied)
	}
}

func TestSpace_GetCarGivenSpace(t *testing.T) {
	car := Car{"reg", "white"}
	s := NewSpace(false, car)

	answer := s.GetCarGivenSpace()

	if answer.color != car.color && answer.regNo != car.color {
		t.Fatalf("expected %v actual %v", car, answer)
	}
}