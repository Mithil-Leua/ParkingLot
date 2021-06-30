package parkingLot

import "testing"

func TestNewParkingLot(t *testing.T) {
	p := NewParkingLot(3, make([]Space, 3), 3)

	if p == nil {
		t.Fatalf("expected %v actual %v", p, nil)
	}
}

func TestNewParkingLotInvalidCapacity(t *testing.T) {
	p := NewParkingLot(-1, make([]Space, 3), -1)

	if p != nil {
		t.Fatalf("expected %v actual %v", nil, p)
	}
}

func TestParkSpaceToCar(t *testing.T) {
	p := NewParkingLot(3, make([]Space, 3), 3)

	regNo := "registration"
	color := "white"
	car := MakeCar(regNo, color)

	p.parkCarToSpace(car, 0)

	if p.FreeSlots != 2 && p.Space[0].Occupied {
		t.Fatalf("expected %d actual %d", 2, p.FreeSlots)
	}
}

func TestAvailableSpace(t *testing.T) {
	p := NewParkingLot(3, make([]Space, 3), 3)

	answer := p.availableSpace()

	if answer != 0 {
		t.Fatalf("expected %d actual %d", 0, answer)
	}
}

func TestAvailableSpaceInvalidSpace(t *testing.T) {
	p := NewParkingLot(3, make([]Space, 3), 0)

	answer := p.availableSpace()

	if answer != -1 {
		t.Fatalf("expected %d actual %d", 0, answer)
	}
}