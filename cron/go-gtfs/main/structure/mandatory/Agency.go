
type Agency struct {
	agency_id string // Conditionally Required

	agency_name string // Required
	agency_url string // Required
	agency_timezone string // Required

	agency_lang string // Optional
	agency_phone string // Optional
	agency_fare_url string // Optional
	agency_email string // Optional
}
