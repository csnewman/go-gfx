package main

import (
	"strings"
)

func shouldIgnore(attrs map[string]string) bool {
	v, ok := take(attrs, "api")
	if !ok {
		return false
	}

	for _, api := range strings.Split(v, ",") {
		if api == "vulkan" {
			return false
		}
	}

	return true
}

func shouldIgnoreExport(attrs map[string]string) bool {
	v, ok := take(attrs, "export")
	if !ok {
		return false
	}

	for _, api := range strings.Split(v, ",") {
		if api == "vulkan" {
			return false
		}
	}

	return true
}

func shouldIgnoreSupported(attrs map[string]string) bool {
	v, ok := take(attrs, "supported")
	if !ok {
		return false
	}

	for _, api := range strings.Split(v, ",") {
		if api == "vulkan" {
			return false
		}
	}

	return true
}

func take[K comparable, V any](m map[K]V, k K) (V, bool) {
	v, ok := m[k]

	delete(m, k)

	return v, ok
}
