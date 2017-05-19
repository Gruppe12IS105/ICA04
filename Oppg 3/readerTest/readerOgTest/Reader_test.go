package readerOgTest

import "testing"

func BenchmarkReadWholeToMem1000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReadWholeToMem()

	}
}
