package sim

// represents a single piece of roadway
type RoadSegment struct {
	SegmentID    int
	Length       float64 // meters
	LaneCount    int
	CarCount     int
	AverageSpeed float64 // meters per second
	//
	NextSegment int // will hold the ID for next segment

}

// telementry data for CSV
type TickData struct {
	TickNumber   int
	LaneCount    int
	CarCount     int
	SegmentID    int
	AverageSpeed float64
	Density      float64
	Capacity     float64
}

// simulation settings + container
type World struct {
	SegmentList     []RoadSegment
	CommandChannels []chan SegmentCommand
	ResultChannels  []chan TickData
	TickCount       int
	SegmentCount    int
	// settings
}

// Sent to segment workers via World
type SegmentCommand struct {
	Tick int
}
