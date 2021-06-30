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
	args := strings.Split(s, " ")
	switch args[0] {
	case "create_parking_lot":
		n, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("CREATE: Invalid value for number of slots.")
		} else {
			plot = parkingLot.NewParkingLot(n, make([]parkingLot.Space, n), n)
		}
	}
}
