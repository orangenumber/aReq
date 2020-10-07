// Written by Gon Yi
package areq

func (p *plugin) BasicAuth(id, pwd string) PluginFn {
	Name := "BASIC-AUTH"
	return PluginFn{
		Name: Name,
		FnPre: func(d *AReq) {
			d.Req.SetBasicAuth(id, pwd)
			// d.Logf("[%s] applied (id: %s)", Name, id)
		},
	}
}
