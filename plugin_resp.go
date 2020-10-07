// Written by Gon Yi
package areq

import (
	"bytes"
	"encoding/json"
	"io"
)

func (p *plugin) RespToJSON(dst interface{}) PluginFn {
	Name := "RES-TO-JSON"
	return PluginFn{
		Name: Name,
		FnPost: func(d *AReq) {
			d.NoOutput = true
			if err := json.NewDecoder(d.Res.Body).Decode(&dst); err != nil {
				d.Logf("[%s] cannot decode JSON", Name)
				d.AddPluginFnErr(err)
				d.NoOutput = false // if failed, use standard out
			}
		},
	}
}

func (p *plugin) RespToByteBuf(dst *bytes.Buffer) PluginFn {
	Name := "RES-TO-BYTE-BUF"
	return PluginFn{
		Name: Name,
		FnPost: func(d *AReq) {
			d.NoOutput = true
			count, err := io.Copy(dst, d.Res.Body)
			d.Logf("[%s] copied %d KB\n", Name, count/1024)
			if err != nil {
				d.Logf("[%s] %s\n", Name, err.Error())
			}
		},
	}
}
