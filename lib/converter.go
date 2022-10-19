package lib

import (
	"reflect"
	"strings"
	"unsafe"
)

// go 1.18 or higher
func MapValueToArray[M ~map[K]V, K comparable, V any](m M) []V {
	r := make([]V, 0, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

// go 1.18 or higher
func MapKeyToArray[M ~map[K]V, K comparable, V any](m M) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

func ReplaceAtIndex1(str string, replacement rune, index int) string {
	out := []rune(str)
	out[index] = replacement
	return string(out)
}

func ReplaceAtIndex2(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func ReplaceAtIndex3(str string, replacement rune, index int) string {
	return strings.Join([]string{str[:index], str[index+1:]}, string(replacement))
}

func strToBytes(str string) []byte {
	string_header := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bytes_header := &reflect.SliceHeader{
		Data: string_header.Data,
		Len:  string_header.Len,
		Cap:  string_header.Len,
	}
	return *(*[]byte)(unsafe.Pointer(bytes_header))
}

func strToBytesCopy(str string) []byte {
	bytes_unsafe := strToBytes(str)
	bytes := make([]byte, len(bytes_unsafe))
	copy(bytes, bytes_unsafe)
	return bytes
}

func bytesToStr(bytes []byte) string {
	bytes_header := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	string_header := &reflect.StringHeader{
		Data: bytes_header.Data,
		Len:  bytes_header.Len,
	}
	return *(*string)(unsafe.Pointer(string_header))
}

func ReplaceAtIndex4(str string, replacement rune, index int) string {
	bytes := strToBytesCopy(str)
	bytes[index] = byte(replacement)
	return bytesToStr(bytes)
}

func ReplaceAtIndex5(str string, replacement rune, index int) string {
	bytes := strToBytes(str)
	bytes[index] = byte(replacement)
	return bytesToStr(bytes)
}

func GenerateString(n int) string {
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < n; i++ {
		b.WriteRune('a')
	}
	return b.String()
}
