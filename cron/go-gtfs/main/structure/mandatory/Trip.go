
type Trip struct {
	route_id string // Required
	service_id string // Required
	trip_id string // Required
	trip_headsign string // Optional
	trip_short_name string // Optional
	direction_id int // Enum // Optional
	block_id string // Optional
	shape_id string // Conditionally Required
	wheelchair_accessible int // Enum // Optional
	bikes_allowed int // Enum // Optional
}
