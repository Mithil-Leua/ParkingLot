package parkingLot

type Space struct {
	Occupied bool
	Car      Car
}

func NewSpace(occupied bool, car Car) *Space {
	s := Space{}
	s.Occupied = occupied
	s.Car = car
	return &s
}

func (s *Space) AssignSpaceToCar(c Car) {
	s.Occupied = true
	s.Car = c
}

func (s *Space) FreeSpace() {
	s.Occupied = false
}

func (s Space) GetCarGivenSpace() Car {
	return s.Car
}
