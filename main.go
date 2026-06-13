package main

import (
	"fmt"
	"sync"
)

type Truck interface {
	LoadTruck() error
	UnloadTruck() error
}

type NormalTruck struct {
	id    int
	cargo bool
}

type ElectricTruck struct {
	id      int
	battery float64
	cargo   bool
}

func (t *NormalTruck) LoadTruck() error {
	t.cargo = true
	fmt.Printf("Normal Truck %d loaded.\n", t.id)
	return nil
}

func (t *NormalTruck) UnloadTruck() error {
	t.cargo = false
	fmt.Printf("Normal Truck %d unloaded.\n", t.id)
	return nil
}

func (t *ElectricTruck) LoadTruck() error {
	t.cargo = true
	fmt.Printf("Electric Truck %d loaded.\n", t.id)
	return nil
}

func (t *ElectricTruck) UnloadTruck() error {
	t.cargo = false
	fmt.Printf("Electric Truck %d unloaded.\n", t.id)
	return nil
}

func processTrucks(t []Truck) error {
	for _, truck := range t {
		if err := truck.LoadTruck(); err != nil {
			fmt.Printf("Error loading truck: %v\n", err)
			return err
		}
		if err := truck.UnloadTruck(); err != nil {
			fmt.Printf("Error unloading truck: %v\n", err)
			return err
		}
	}
	return nil
}

func ProcessFleet(t []Truck) error {

	var wg sync.WaitGroup
	for _, truck := range t {
		wg.Add(1)

		go func(truck Truck) {
			if err := processTrucks([]Truck{truck}); err != nil {
				fmt.Printf("Error processing truck: %v\n", err)
			}
			wg.Done()
		}(truck)
	}
	wg.Wait()

	return nil
}

func main() {
	trucks := []Truck{
		&NormalTruck{id: 1},
		&ElectricTruck{id: 2, battery: 100.0},
		&NormalTruck{id: 3},
		&ElectricTruck{id: 4, battery: 80.0},
	}

	if err := ProcessFleet(trucks); err != nil {
		fmt.Printf("Error processing fleet: %v\n", err)
	}
}
