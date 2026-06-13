package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {

		fmt.Printf("Loading Normal Truck.\n")

		t.Run("Should load and unload cargo", func(t *testing.T) {
			err := processTruck(&NormalTruck{id: "001"})
			if err != nil {
				t.Fatalf("Error in processing truck: %s", err)
			}

			fmt.Printf("Loading Electric Truck.\n")

			err = processTruck(&ElectricTruck{id: "002", cargo: "loaded cargo", battery: 57.63})
			if err != nil {
				t.Fatalf("Error in processing truck: %s", err)
			}
		})
	})
}
