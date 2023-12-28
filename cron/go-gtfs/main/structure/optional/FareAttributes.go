
type FareAttribute struct {
	fare_id string // Required
	price float32 // Required
	currency_type string // Required
	payment_type int // Enum // Required
	transfers int // Enum // Required
	agency_id string // Required
	transfer_duration int // Optional
}
