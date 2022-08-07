package envconfig

import (
	"regexp"
	"testing"
)

var envvarNameRe2 = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)
var envvarNameRe = regexp.MustCompile(`\A[a-zA-Z_][a-zA-Z0-9_]*\z`)

// BenchmarkEnvvarNameRe2-8   	11781014	        93.28 ns/op	       3 B/op	       1 allocs/op
func BenchmarkEnvvarNameRe2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = envvarNameRe2.Match([]byte("ABC"))
	}
}

// BenchmarkEnvvarNameRe-8   	12946851	        92.65 ns/op	       3 B/op	       1 allocs/op
func BenchmarkEnvvarNameRe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = envvarNameRe.Match([]byte("ABC"))
	}
}

// BenchmarkValidateEnv-8   	187759477	         6.232 ns/op	       0 B/op	       0 allocs/op
func BenchmarkValidateEnv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = validateEnv("ABC")
	}
}
