package areq

// Optional custom function
// 1. FnPre:  Before-Request-Send
// 2. FnPost: After-Request-Sent
type PluginFn struct {
	Name   string
	FnPre  func(d *AReq)
	FnPost func(d *AReq)
}

func NewPluginFn(name string, fnPre func(d *AReq), fnPost func(d *AReq)) PluginFn {
	return PluginFn{
		Name:   name,
		FnPre:  fnPre,
		FnPost: fnPost,
	}
}
