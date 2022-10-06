package lib

func MapValueToArray[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func MapKeyToArray[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

//func MapToArrayValue(maps map[any]any) []any {
//	r := make([]interface{}, len(maps))
//	for _, val := range maps {
//		r = append(r, val)
//	}
//	return r
//}
//
//func MapToArrayKey(maps map[any]any) []any {
//	r := make([]interface{}, len(maps))
//	for key := range maps {
//		r = append(r, key)
//	}
//	return r
//}
