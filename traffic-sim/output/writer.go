package output

import (
	"encoding/csv"
	"os"
	"strconv"

	"traffic-sim/sim"
)

func WriteCSV(filename string, Data [][]sim.TickData) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{
		"Tick",
		"LaneCount",
		"CarCount",
		"AverageSpeed",
		"Density",
		"Capacity",
	})

	for _, tickData := range Data {
		for _, row := range tickData {
			writer.Write([]string{
				strconv.Itoa(row.TickNumber),
				strconv.Itoa(row.LaneCount),
				strconv.Itoa(row.CarCount),
				strconv.FormatFloat(row.AverageSpeed, 'f', 4, 64),
				strconv.FormatFloat(row.Density, 'f', 4, 64),
				strconv.FormatFloat(row.Capacity, 'f', 4, 64),
			})
		}
	}

	return nil
}

