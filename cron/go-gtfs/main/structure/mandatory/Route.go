
type Route struct {
	route_id string // Required
	agency_id string // Conditionally Required
	route_short_name string // Conditionally Required
	route_long_name string // Conditionally Required
	route_desc string // Optional
	route_type Int // Enum // Required
	route_url string // Optional
	route_color string // Color // Optional
	route_text_color string // Color // Optional
	route_sort_order int // Optional
	continuous_pickup int // Enum // Optional
	continuous_drop_off int // Enum // Optional

	network_id string // Conditionally Forbidden
}
