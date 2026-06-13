package main

import (
	"fmt"
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo string
}

type ElectricTruck struct {
	id      string
	cargo   string
	battery float64
}

func (t *ElectricTruck) LoadCargo() error {
	fmt.Printf("Loading cargo for Electric Truck: %s\n", t.id)
	t.cargo = "loaded cargo"
	return nil
}

func (t *ElectricTruck) UnloadCargo() error {
	fmt.Printf("Unloading cargo for Electric Truck: %s\n", t.id)
	t.cargo = "no cargo."
	t.battery -= 10.0 // Simulate battery usage during unloading
	return nil
}

func (t *NormalTruck) LoadCargo() error {
	fmt.Printf("Loading cargo for Truck: %s\n", t.id)
	t.cargo = "loaded cargo"
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	fmt.Printf("Unloading cargo for Truck: %s\n", t.id)
	t.cargo = "no cargo."
	return nil
}

// processTruck handles the loading and unloading of a truck.
func processTruck(truck Truck) error {

	fmt.Printf("%+v is arrived.\n", truck)

	if err := truck.LoadCargo(); err != nil {
		return fmt.Errorf("failed to load cargo: %w", err)
	}

	fmt.Printf("%+v is going to the destination\n", truck)
	fmt.Printf("Arrived at the destination\n")

	if err := truck.UnloadCargo(); err != nil {
		return fmt.Errorf("failed to unload cargo: %w", err)
	}
	fmt.Printf("%+v is leaving.\n", truck)
	return nil
}

// func main() {
// 	fmt.Printf("Loading Normal Truck.\n")
// 	err := processTruck(&NormalTruck{id: "001"})
// 	if err != nil {
// 		log.Fatalf("Error in processing truck: %s", err)
// 	}

// 	fmt.Printf("Loading Electric Truck.\n")
// 	err = processTruck(&ElectricTruck{id: "002", cargo: "loaded cargo", battery: 57.63})
// 	if err != nil {
// 		log.Fatalf("Error in processing truck: %s", err)
// 	}
// }
