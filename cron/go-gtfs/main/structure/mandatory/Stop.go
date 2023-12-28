
type Stops struct {
	stop_id string // Required
	zone_id string // Conditionally Required

	stop_name string // Conditionally Required
	stop_lat string // Conditionally Required
	stop_long string // Conditionally Required
	parent_station string // Conditionally Required

	stop_code string // Optional
	tts_stop_name string // Optional
	stop_desc string // Optional
	stop_url string // Optional
	location_type string // Optional
	stop_timezone string // Optional
	wheelchair_boarding int // Optional
	level_id string // Optional
	platform_code string // Optional
}
