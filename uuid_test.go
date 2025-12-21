package uuidbench

import (
	"testing"

	coolaj86 "github.com/coolaj86/uuidv7"
	toolbox "github.com/eucatur/go-toolbox/uuid"
	gofrs "github.com/gofrs/uuid"
	google "github.com/google/uuid"
	lsluuid "github.com/lsl/uuid"
	proto "github.com/uuid6/uuid6go-proto"
)

// https://github.com/gofrs/uuid/blob/master/uuid.go
// https://github.com/gofrs/uuid/blob/master/generator.go#L412
func BenchmarkGofrsV7(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gofrs.NewV7()
	}
}

// https://github.com/google/uuid/blob/2d3c2a9cc518326daf99a383f07c4d3c44317e4d/version7.go#L48
func BenchmarkGoogleV7(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = google.NewV7()
	}
}

// https://github.com/eucatur/go-toolbox/blob/6796f092db6ee99c261cd985a509a6719aa65ab6/uuid/uuid_v7.go#L17
func BenchmarkToolboxV7(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = toolbox.GenerateUUIDv7()
	}
}

// https://github.com/uuid6/uuid6go-proto/blob/d8acedd7f415e747b00a65c47e4e242f4f479c52/uuid_v7.go
func BenchmarkUUID6ProtoV7(b *testing.B) {
	var test proto.UUIDv7Generator
	test.SubsecondPrecisionLength = 16
	test.NodePrecisionLength = 4
	test.Node = 15
	test.CounterPrecisionLength = 8

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = test.Next()
	}
}

// https://github.com/coolaj86/uuidv7/blob/5eddd0f27f07fc2d4f2f93ff953f44e7cce8bd5b/uuidv7.go#L1
func BenchmarkCoolAJV7(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = coolaj86.New()
	}
}

// https://github.com/lsl/uuid/blob/main/uuidv7.go
func BenchmarkLSLV7(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = lsluuid.NewV7()
	}
}
