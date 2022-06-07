package models

func stringFromAny(val any) string {
	return val.(string)
}

func floatFromAny(val any) float64 {
	return val.(float64)
}

func mapFromAny(val any) map[string]any {
	return val.(map[string]any)
}
