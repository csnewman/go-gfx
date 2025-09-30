package main

func shouldIgnore(attrs map[string]string) bool {
	v, ok := take(attrs, "api")
	if !ok {
		return false
	}

	return v != "vulkan"
}

func take[K comparable, V any](m map[K]V, k K) (V, bool) {
	v, ok := m[k]

	delete(m, k)

	return v, ok
}
