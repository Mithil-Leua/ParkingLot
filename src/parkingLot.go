package parkingLot

type ParkingLot struct {
	Capacity  int
	Space     []Space
	FreeSlots int
}

func NewParkingLot(capacity int, space []Space, freeSlots int) *ParkingLot {
	if capacity <= 0 {
		return nil
	}
	p := ParkingLot{}
	p.Capacity = capacity
	p.Space = space
	p.FreeSlots = freeSlots
	return &p
}

func (p *ParkingLot) parkCarToSpace(c Car, i int) {
	p.Space[i].AssignSpaceToCar(c)
	p.FreeSlots--
}

func (p ParkingLot) availableSpace() int {
	if p.FreeSlots > 0 {
		n := p.Capacity
		for i := 0; i < n; i++ {
			if !p.Space[i].Occupied {
				return i
			}
		}
	}
	return -1
}
