package sim

// create world + road segments
func InitalizeWorld() World {

	segments := []RoadSegment{
		{SegmentID: 1, Length: 25.0, LaneCount: 4, CarCount: 10, AverageSpeed: 31.3},
		{SegmentID: 2, Length: 50.0, LaneCount: 3, CarCount: 7, AverageSpeed: 33.4},
		{SegmentID: 3, Length: 45.0, LaneCount: 2, CarCount: 5, AverageSpeed: 30.0},
	}

	world := World{
		SegmentList: segments,
		TickCount:   100,
	}

	world.CommandChannels = make([]chan SegmentCommand, len(segments))
	world.ResultChannels = make([]chan TickData, len(segments))

	for i := range segments {

		cmdChan := make(chan SegmentCommand)
		resChan := make(chan TickData)

		world.CommandChannels[i] = cmdChan
		world.ResultChannels[i] = resChan

		go SegmentWorker(&world.SegmentList[i], cmdChan, resChan)

	}

	return world
}

// Runs one tick of the World
// Updates every segment once
func RunStep(world *World, tickNumber int) []TickData {

	results := make([]TickData, len(world.SegmentList))

	for i := range world.SegmentList {
		world.CommandChannels[i] <- SegmentCommand{Tick: tickNumber}
	}

	for i := range world.SegmentList {
		res := <-world.ResultChannels[i]
		results[i] = res
	}

	return results

}

func SegmentWorker(
	seg *RoadSegment,
	commandChan <-chan SegmentCommand,
	resultChan chan<- TickData,
) {

	for cmd := range commandChan {

		newCarCount := seg.CarCount // placeholder
		// newAvgSpeed := seg.AverageSpeed // placeholder

		density := float64(newCarCount) / (float64(seg.LaneCount) * seg.Length)

		resultChan <- TickData{
			TickNumber:   cmd.Tick,
			LaneCount:    seg.LaneCount,
			CarCount:     seg.CarCount,
			AverageSpeed: seg.AverageSpeed,
			Density:      density,
			Capacity:     0.0, // placeholder
		}

	}

}
