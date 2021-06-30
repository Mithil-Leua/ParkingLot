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

func (p *ParkingLot) Unpark(space int) {
	if p == nil {
		fmt.Println("No parking lot")
		return
	}
	if space < 1 || space > p.Capacity {
		fmt.Println("Invalid slot number.")
	}
	fmt.Printf("Slot number %d is free\n", space)
	p.Space[space-1].FreeSpace()
	p.FreeSlots++
}

func (p ParkingLot) PrintParkingLot() {
	fmt.Printf("Slot No.    Registration No    Colour\n")
	for i := 0; i < p.Capacity; i++ {
		if p.Space[i].Occupied {
			car := p.Space[i].GetCarGivenSpace()
			fmt.Printf("%d           %s      %s\n", i+1, car.GetRegNo(), car.GetColor())
		}
	}
}

func (p ParkingLot) GetCarDetailsOfAGivenColor(color string) {
	for i := 0; i < p.Capacity; i++ {
		if p.Space[i].Occupied {
			car := p.Space[i].GetCarGivenSpace()
			if &car != nil && car.GetColor() == color {
				fmt.Printf("%s\n", car.GetRegNo())
			}
		}
	}
}

func (p ParkingLot) GetSlotNoOfColor(color string) {
	for i := 0; i < p.Capacity; i++ {
		if p.Space[i].Occupied {
			car := p.Space[i].GetCarGivenSpace()
			if &car != nil && car.GetColor() == color {
				fmt.Println("%d", i+1)
			}
		}
	}
}

func (p ParkingLot) GetSpaceOfACar(regNo string) {
	for i := 0; i < p.Capacity; i++ {
		if p.Space[i].Occupied {
			car:= p.Space[i].GetCarGivenSpace()
			if &car != nil && car.GetRegNo() == regNo {
				fmt.Printf("%d\n", i+1)
				return
			}
		}
	}
	fmt.Printf("Not found\n")
}
