
type FeedInfo struct {
	feed_publisher_name string // Required
	feed_publisher_url string // Required
	feed_lang string // Required
	default_lang string // Optional
	feed_start_date string // Date // Recommended
	feed_end_date string // Date // Recommended
	feed_version string // Recommended
	feed_contact_email string // Optional
	feed_contact_url string // Optional
}
