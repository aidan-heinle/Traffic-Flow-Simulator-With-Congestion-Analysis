package sim

// represents a single piece of roadway
type RoadSegment struct {
	SegmentID    int
	Length       float64 // meters
	LaneCount    int
	CarCount     int
	AverageSpeed float64 // meters per second
}

// telementry data for CSV
type TickData struct {
	TickNumber   int
	LaneCount    int
	CarCount     int
	AverageSpeed float64
	Density      float64
	Capacity     float64
}

// simulation settings + container
type World struct {
	SegmentList []RoadSegment
	TickCount   int
	// settings
}
