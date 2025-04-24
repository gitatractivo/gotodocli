package utils

// GetStatusEmoji returns an emoji based on task status
func GetStatusEmoji(status bool) string {
	if status {
		return "✅"
	}
	return "❌"
}