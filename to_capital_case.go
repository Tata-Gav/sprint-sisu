package sprint_test
func ToCapitalCase(s string) string {
	var result []byte
	var capitalizeNext bool

	for i := 0; i < len(s); i++ {
		// Check if the character is alphanumeric
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') {
			if capitalizeNext {
				// Capitalize the current character
				if s[i] >= 'a' && s[i] <= 'z' {
					result = append(result, s[i]-32) // Convert lowercase to uppercase
				} else {
					result = append(result, s[i]) // Already uppercase
				}
				capitalizeNext = false
			} else {
				// Lowercase the current character
				if s[i] >= 'A' && s[i] <= 'Z' {
					result = append(result, s[i]+32) // Convert uppercase to lowercase
				} else {
					result = append(result, s[i]) // Already lowercase
				}
			}
		} else {
			// Non-alphanumeric character, just append it
			result = append(result, s[i])
			// Next character should be capitalized if it is alphanumeric
			if s[i] == ' ' || s[i] == '!' || s[i] == '-' || s[i] == '?' {
				capitalizeNext = true
			}
		}
	}

	return string(result)
}