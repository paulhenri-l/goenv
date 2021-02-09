package goenv

import (
	"github.com/spf13/cast"
	"os"
	"strings"
)

const Production = "production"
const Staging = "staging"
const Local = "local"

func Get(key string, def interface{}) interface{} {
	val := os.Getenv(key)

	if val == "" {
		return def
	}

	return val
}

func AppEnv() string {
	return StringOr("APP_ENV", "production")
}

func StringOr(key string, def string) string {
	return cast.ToString(
		Get(key, def),
	)
}

func BoolOr(key string, def bool) bool {
	return cast.ToBool(
		Get(key, def),
	)
}

func IntOr(key string, def int) int {
	return cast.ToInt(
		Get(key, def),
	)
}

func Int32Or(key string, def int32) int32 {
	return cast.ToInt32(
		Get(key, def),
	)
}

func Int64Or(key string, def int64) int64 {
	return cast.ToInt64(
		Get(key, def),
	)
}

func UintOr(key string, def uint) uint {
	return cast.ToUint(
		Get(key, def),
	)
}

func Uint32Or(key string, def uint32) uint32 {
	return cast.ToUint32(
		Get(key, def),
	)
}

func Uint64Or(key string, def uint64) uint64 {
	return cast.ToUint64(
		Get(key, def),
	)
}

func Float64Or(key string, def float64) float64 {
	return cast.ToFloat64(
		Get(key, def),
	)
}

func IntSliceOr(key string, def []int) []int {
	value := cast.ToString(Get(key, ""))
	parts := strings.Split(value, ",")
	final := make([]int, len(parts))

	if parts[0] == value {
		return def
	}

	for i, part := range parts {
		final[i] = cast.ToInt(part)
	}

	return final
}

func StringSliceOr(key string, def []string) []string {
	value := cast.ToString(Get(key, ""))
	parts := strings.Split(value, ",")
	final := make([]string, len(parts))

	if parts[0] == value {
		return def
	}

	for i, part := range parts {
		final[i] = cast.ToString(part)
	}

	return final
}
