package goenv

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	mockEnv(t, "SOME_ENV", "world")

	assert.Equal(t, "world", Get("SOME_ENV", ""))
	assert.Equal(t, "default", Get("NOT_SET", "default"))
}

func TestAppEnv(t *testing.T) {
	mockEnv(t, "APP_ENV", "local")

	assert.Equal(t, "local", AppEnv())
}

func TestAppEnv_DefaultsToProduction(t *testing.T) {
	assert.Equal(t, "production", AppEnv())
}

func TestEnvString(t *testing.T) {
	mockEnv(t, "SOME_ENV", "world")

	assert.Equal(t, "world", StringOr("SOME_ENV", ""))
	assert.Equal(t, "default", StringOr("NOT_SET", "default"))
}

func TestEnvBool(t *testing.T) {
	mockEnv(t, "SOME_ENV", "1")

	assert.Equal(t, true, BoolOr("SOME_ENV", false))
	assert.Equal(t, true, BoolOr("NOT_SET", true))
}

func TestEnvInt(t *testing.T) {
	mockEnv(t, "SOME_ENV", "1")

	assert.Equal(t, 1, IntOr("SOME_ENV", 0))
	assert.Equal(t, 2, IntOr("NOT_SET", 2))
}

func TestEnvInt32(t *testing.T) {
	mockEnv(t, "SOME_ENV", "1")

	assert.Equal(t, int32(1), Int32Or("SOME_ENV", int32(2)))
	assert.Equal(t, int32(2), Int32Or("NOT_SET", int32(2)))
}

func TestEnvInt64(t *testing.T) {
	mockEnv(t, "SOME_ENV", "1")

	assert.Equal(t, int64(1), Int64Or("SOME_ENV", int64(2)))
	assert.Equal(t, int64(2), Int64Or("NOT_SET", int64(2)))
}

func TestEnvUint(t *testing.T) {
	mockEnv(t, "SOME_ENV", "1")

	assert.Equal(t, uint(1), UintOr("SOME_ENV", uint(2)))
	assert.Equal(t, uint(2), UintOr("NOT_SET", uint(2)))
}

func TestEnvUint32(t *testing.T) {
	mockEnv(t, "SOME_ENV", "1")

	assert.Equal(t, uint32(1), Uint32Or("SOME_ENV", uint32(2)))
	assert.Equal(t, uint32(2), Uint32Or("NOT_SET", uint32(2)))
}

func TestEnvUint64(t *testing.T) {
	mockEnv(t, "SOME_ENV", "1")

	assert.Equal(t, uint64(1), Uint64Or("SOME_ENV", uint64(2)))
	assert.Equal(t, uint64(2), Uint64Or("NOT_SET", uint64(2)))
}

func TestEnvFloat64(t *testing.T) {
	mockEnv(t, "SOME_ENV", "1.2")

	assert.Equal(t, 1.2, Float64Or("SOME_ENV", float64(2)))
	assert.Equal(t, 2.2, Float64Or("NOT_SET", 2.2))
}

func TestEnvIntSlice(t *testing.T) {
	mockEnv(t, "SOME_ENV", "1,2,3,4")

	assert.Equal(t, []int{1, 2, 3, 4}, IntSliceOr("SOME_ENV", []int{}))
	assert.Equal(t, []int{1, 2}, IntSliceOr("NOT_SET", []int{1, 2}))
}

func TestEnvStringSlice(t *testing.T) {
	mockEnv(t, "SOME_ENV", "a,b,c,d")

	assert.Equal(t, []string{"a", "b", "c", "d"}, StringSliceOr("SOME_ENV", []string{}))
	assert.Equal(t, []string{"a", "b"}, StringSliceOr("NOT_SET", []string{"a", "b"}))
}

func mockEnv(t *testing.T, key, value string) {
	orig := os.Getenv(key)

	_ = os.Setenv(key, value)

	t.Cleanup(func() {
		_ = os.Setenv(key, orig)
	})
}
