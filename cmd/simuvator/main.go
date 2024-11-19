package main

import (
	"fmt"
	"time"

	"github.com/dshaneg/elevator/internal/elevator/bank"
	"github.com/dshaneg/elevator/internal/elevator/car"
	"github.com/dshaneg/elevator/internal/passenger"
)

func main() {
	floorCount := 5
	cars := []bank.Member{
		car.NewCar(floorCount),
		car.NewCar(floorCount),
		car.NewCar(floorCount),
	}

	b, err := bank.New(floorCount, cars)
	if err != nil {
		panic(err)
	}

	passengers := []*passenger.Passenger{
		passenger.New(b, passenger.WithPrimaryFloor(3)),
	}
	runSim(passengers)
}

func runSim(passengers []*passenger.Passenger) {
	simTime := time.Date(0, 1, 1, 0, 0, 0, 0, time.Local)

	for range time.Tick(1 * time.Second) {
		simTime = simTime.Add(1 * time.Minute)
		fmt.Printf("%v Tick\n", simTime)

		for _, p := range passengers {
			p.Tick(simTime)
			fmt.Printf("Passenger on floor %d\n", p.Floor())
		}
	}
}
