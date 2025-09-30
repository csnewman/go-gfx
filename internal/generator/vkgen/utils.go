package main

func shouldIgnore(attrs map[string]string) bool {
	v, ok := take(attrs, "api")
	if !ok {
		return false
	}

	return v != "vulkan"
}
