package utils

func SafeIsEqualStrPointer(a *string, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}

	return *a == *b
}
