package benchmarking_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hiteshrepo/grpc-demo/dist"
	pb "github.com/hiteshrepo/grpc-demo/dist/proto"
)

func BenchmarkJSONMarshal(b *testing.B) {
	s := dist.JsonSmall
	m := dist.JsonMedium
	l := dist.JsonLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(&s)
			_ = d
		}
		b.ReportAllocs()
	})
	b.Run("MediumData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(&m)
			_ = d
		}
		b.ReportAllocs()
	})
	b.Run("LargeData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(&l)
			_ = d
		}
		b.ReportAllocs()
	})
	fmt.Printf("\n")
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	s := dist.JsonSmall
	m := dist.JsonMedium
	l := dist.JsonLarge

	sd, _ := json.Marshal(&s)
	md, _ := json.Marshal(&m)
	ld, _ := json.Marshal(&l)

	var sf []byte
	var mf []byte
	var lf []byte

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(sd, &sf)
		}
		b.ReportAllocs()
	})
	b.Run("MediumData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(md, &mf)
		}
		b.ReportAllocs()
	})
	b.Run("LargeData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(ld, &lf)
		}
		b.ReportAllocs()
	})
	fmt.Printf("\n")
}

func BenchmarkProtocMarshal(b *testing.B) {
	s := dist.ProtocSmall
	m := dist.ProtocMedium
	l := dist.ProtocLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&s)
			_ = d
		}
		b.ReportAllocs()
	})
	b.Run("MediumData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&m)
			_ = d
		}
		b.ReportAllocs()
	})
	b.Run("LargeData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&l)
			_ = d
		}
		b.ReportAllocs()
	})
	fmt.Printf("\n")
}

func BenchmarkProtocUnMarshal(b *testing.B) {
	s := dist.ProtocSmall
	m := dist.ProtocMedium
	l := dist.ProtocLarge

	sd, _ := proto.Marshal(&s)
	md, _ := proto.Marshal(&m)
	ld, _ := proto.Marshal(&l)

	var sf pb.BenchSmall
	var mf pb.BenchMedium
	var lf pb.BenchLarge

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(sd, &sf)
		}
		b.ReportAllocs()
	})
	b.Run("MediumData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(md, &mf)
		}
		b.ReportAllocs()
	})
	b.Run("LargeData", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(ld, &lf)
		}
		b.ReportAllocs()
	})
	fmt.Printf("\n")
}
