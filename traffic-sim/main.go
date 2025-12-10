package main

import (
	"fmt"
	"log"
	"traffic-sim/output"
	"traffic-sim/sim"
)

func main() {

	world := sim.InitalizeWorld()

	Data := make([][]sim.TickData, 0, world.TickCount)

	for tick := 0; tick < world.TickCount; tick++ {
		tickData := sim.RunStep(&world, tick)
		Data = append(Data, tickData)
	}

	err := output.WriteCSV("simulation_output.csv", Data)
	if err != nil {
		log.Fatalf("Failed to write CSVL %v", err)
	}

	fmt.Println("Simulation Completed. CSV file generated: simulation_output.csv")

}
