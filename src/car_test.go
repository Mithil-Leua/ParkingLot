package parkingLot

import (
	"testing"
)

func TestNewCar(t *testing.T) {
	regNo := "registration"
	color := "white"

	car := NewCar(regNo, color)
	if car == nil {
		t.Fatalf("expected %s actual %s", car, "nil")
	}
}

func TestMakeCar(t *testing.T) {
	regNo := "registration"
	color := "white"

	car := MakeCar(regNo, color)
	if &car == nil {
		t.Fatalf("expected %s actual %s", car, "nil")
	}
}

func TestCar_GetColor(t *testing.T) {
	regNo := "registration"
	color := "white"

	car := NewCar(regNo, color)
	answer := car.GetColor()

	if answer != color {
		t.Fatalf("expected %s actual %s", color, answer)
	}
}

func TestCar_GetRegNo(t *testing.T) {
	regNo := "registration"
	color := "white"

	car := NewCar(regNo, color)
	answer := car.GetRegNo()

	if answer != regNo {
		t.Fatalf("expected %s actual %s", regNo, answer)
	}
}
