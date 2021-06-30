package main

import (
	"bufio"
	"fmt"
	"os"
	parkingLot "parkingLot/src"
	"strconv"
	"strings"
)

var plot *parkingLot.ParkingLot

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		if text == "exit" {
			break
		}
		processCommand(text)
	}
}

func processCommand(s string) {
	values := strings.Split(s, " ")

	if len(values) < 2 && values[0] != "status"{
		fmt.Println("ACTION: Invalid command")
		return
	}

	switch values[0] {
	case "create_parking_lot":
		n, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println("CREATE: Invalid value for number of slots.")
		} else {
			plot = parkingLot.NewParkingLot(n, make([]parkingLot.Space, n), n)
		}
	case "park":
		plot.Park(values[1], values[2])
	case "leave":
		i, err := strconv.Atoi(values[1])
		if err != nil {
			fmt.Println("Invalid value for slot number.")
		} else {
			plot.Unpark(i)
		}
	case "status":
		plot.PrintParkingLot()
	case "registration_numbers_for_cars_with_colour":
		plot.GetCarDetailsOfAGivenColor(values[1])
	case "slot_numbers_for_cars_with_colour":
		plot.GetSlotNoOfColor(values[1])
	case "slot_number_for_registration_number":
		plot.GetSpaceOfACar(values[1])
	default:
		fmt.Println("ACTION: Invalid command")
	}
}
