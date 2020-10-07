// Written by Gon Yi
package areq_test

import (
	"github.com/orangenumber/areq"
	"testing"
)

func TestNew(t *testing.T) {
	r := areq.New()
	r.Request("GET", "https://httpbin.org/anything/1232")
	println(r.Buf.String())
}

func BenchmarkNew(b *testing.B) {
	b.ReportAllocs()
	r := areq.New()
	for i := 0; i < b.N; i++ {
		r.Request("GET", "https://httpbin.org/anything/1232")
		// println(r.Buf.String())
	}
}

func TestNew_Naked(t *testing.T) {
	areq.Request("GET", "https://httpbin.org/anything/12q32")
	println(areq.Buf.String())
}

func TestNewPluginFn(t *testing.T) {
	fPre := func(d *areq.AReq) {
		println("start func")
	}
	fPost := func(d *areq.AReq) {
		println("end func")
	}
	f := areq.NewPluginFn("test", fPre, fPost)

	areq.Request("GET", "https://httpbin.org/anything/12q32", f)
}
