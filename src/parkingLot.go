package parkingLot

import "fmt"

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

func (p *ParkingLot) Park(regno string, color string) {
	if p == nil {
		fmt.Println("No parking lot")
	}

	allottedSpace := p.availableSpace()

	if allottedSpace == -1 {
		fmt.Println("Sorry, parking lot is full")
	}

	p.parkCarToSpace(MakeCar(regno, color), allottedSpace)
	fmt.Printf("Allocated slot number: %d\n", allottedSpace+1)
}
