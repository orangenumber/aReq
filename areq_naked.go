// Written by Gon Yi
package areq

var naked = New()
var Buf = &naked.Buf
var Log = &naked.Log

func Request(method, url string, plugFn ...PluginFn) (statusCode int, err error) {
	return naked.Request(method, url, plugFn...)
}
