
type Frequency struct {
	trip_id string // Required
	start_time string // Time // Required
	end_time string // Time // Required
	headway_secs int // Required
	exact_times int // Enum // Optional
}
