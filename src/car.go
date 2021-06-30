package parkingLot

type Car struct {
	regNo string
	color string
}

func NewCar(regNo string, color string) *Car {
	c := Car{}
	c.regNo = regNo
	c.color = color
	return &c
}

func MakeCar(regNo string, color string) Car {
	c := Car{}
	c.regNo = regNo
	c.color = color
	return c
}

func (v *Car) SetRegNo(regNo string) {
	v.regNo = regNo
}

func (v Car) GetRegNo() string {
	return v.regNo
}

func (v *Car) SetColor(color string) {
	v.color = color
}

func (v Car) GetColor() string {
	return v.color
}
