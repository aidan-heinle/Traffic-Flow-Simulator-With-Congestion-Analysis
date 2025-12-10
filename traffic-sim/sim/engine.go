package sim

// create world + road segments
func InitalizeWorld() World {
	S1 := RoadSegment{
		SegmentID:    1,
		Length:       25.0,
		LaneCount:    4,
		CarCount:     10,
		AverageSpeed: 31.3,
	}

	S2 := RoadSegment{
		SegmentID:    2,
		Length:       50.0,
		LaneCount:    3,
		CarCount:     7,
		AverageSpeed: 33.4,
	}

	S3 := RoadSegment{
		SegmentID:    3,
		Length:       45.0,
		LaneCount:    2,
		CarCount:     5,
		AverageSpeed: 30.0,
	}

	SimulationWorld := World{
		SegmentList: []RoadSegment{S1, S2, S3},
		TickCount:   100,
	}

	return SimulationWorld

}

// updates a single RoadSegment for a single tick
// returns Density
func UpdateSegment(seg *RoadSegment, newCarCount int, newAvgSpeed float64) float64 {
	seg.CarCount = newCarCount
	seg.AverageSpeed = newAvgSpeed

	density := float64(seg.CarCount) / (float64(seg.LaneCount) * seg.Length)

	return density
}

// Runs one tick of the World
// Updates every segment once
func RunStep(world *World, tickNumber int) []TickData {

	tickOutput := make([]TickData, 0, len(world.SegmentList))

	for i := range world.SegmentList {
		seg := &world.SegmentList[i]

		newCarCount := seg.CarCount     // Placeholder value
		newAvgSpeed := seg.AverageSpeed // Placeholder value

		density := UpdateSegment(seg, newCarCount, newAvgSpeed)

		tickOutput = append(tickOutput, TickData{
			TickNumber:   tickNumber,
			LaneCount:    seg.LaneCount,
			CarCount:     seg.CarCount,
			AverageSpeed: seg.AverageSpeed,
			Density:      density,
			Capacity:     0.0, // Placeholder
		})
	}
	return tickOutput

}

