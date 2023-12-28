
type Transfer struct {
	from_stop_id string // Conditionally Required
	to_stop_id string // Conditionally Required
	from_route_id string // Optional
	to_route_id string // Optional
	from_trip_id string // Conditionally Required
	to_trip_id string // Conditionally Required
	transfer_type int // Enum // Required
	min_transfer_time int // Optional
}
