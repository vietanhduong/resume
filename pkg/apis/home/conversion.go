package home

import "fmt"

// ConvertMetadataToMap metadata input section to map
func ConvertMetadataToMap(metadata Metadata) map[string]interface{} {
	return map[string]interface{}{
		"address":  metadata.Address,
		"phone":    metadata.Phone,
		"email":    metadata.Email,
		"websites": metadata.Websites,
	}
}

// ConvertSectionToMap convert input section to map
func ConvertSectionToMap(section Section) map[string]interface{} {
	var achievements []map[string]interface{}
	for _, achievement := range section.Achievements {
		achievements = append(achievements, ConvertAchievementToMap(achievement))
	}
	return map[string]interface{}{
		"name":         section.Name,
		"achievements": achievements,
	}
}

// ConvertAchievementToMap convert input achievement to map
func ConvertAchievementToMap(achievement Achievement) map[string]interface{} {
	result := map[string]interface{}{
		"name":     achievement.Name,
		"duration": ConvertDurationToString(achievement.Duration),
		"details":  achievement.Details,
	}
	return result
}

// ConvertDurationToString convert input duration to string
// the result should be: start (- end)
func ConvertDurationToString(duration Duration) string {
	if duration.End == "" {
		return duration.Start
	}
	return fmt.Sprintf("%s - %s", duration.Start, duration.End)
}
