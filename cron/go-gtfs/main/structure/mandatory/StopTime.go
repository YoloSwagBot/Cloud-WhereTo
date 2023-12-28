
type StopTime struct {
	trip_id string // Required
	arrival_time string // Enum // Conditionally Required
	departure_time string // Enum // Conditionally Required
	stop_id string // Required
	stop_sequence string // Required
	stop_headsign string // Optional
	pickup_type int // Enum // Optional
	drop_off_type int // Enum // Optional
	continuous_pickup int // Enum // Optional
	continuous_drop_off int // Enum // Optional
	shape_dist_traveled float32 // Optional
	timepoint int // Enum // Recommended
}
