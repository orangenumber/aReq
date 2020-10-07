// Written by Gon Yi
package areq

import (
	"compress/gzip"
	"io"
	"io/ioutil"
)

var Plugin plugin

type plugin struct{}

func (p *plugin) SetBody(body io.Reader, fileExt string) PluginFn {
	Name := "SET-BODY"

	rc, ok := body.(io.ReadCloser)
	if !ok && body != nil {
		rc = ioutil.NopCloser(body)
	}
	return PluginFn{
		Name: Name,
		FnPre: func(d *AReq) {
			d.Req.Body = rc
			// d.Logf("[%s] applied", Name)
			if fileExt != "" {
				d.Req.Header.Set("Content-Type", MimeByFileExt(fileExt, "json"))
			}
		},
	}
}

func (p *plugin) AcceptEncoding(encType string) PluginFn {
	Name := "ACCEPT-ENCODING"
	return PluginFn{
		Name: Name,
		FnPre: func(d *AReq) {
			d.Logf("[%s] requesting response in %s", Name, encType)
			d.Req.Header.Add("Accept-Encoding", encType)
		},
		FnPost: func(d *AReq) {
			if d.Res.Header.Get("Content-Encoding") == encType {
				println("ct", d.Res.ContentLength)
				d.Logf("[%s] received %s response (%d bytes)", Name, encType, d.Res.ContentLength)
				var err error
				if d.Res.Body, err = gzip.NewReader(d.Res.Body); err != nil {
					d.Logf("[%s] ERROR: %s", Name, err.Error())
				}
			}
		},
	}
}

func (p *plugin) SetContentType(fileExt string) PluginFn {
	Name := "SET-CONTENT-TYPE"
	return PluginFn{
		Name: Name,
		FnPre: func(d *AReq) {
			d.Req.Header.Set("Content-Type", MimeByFileExt(fileExt, "json"))
		},
	}
}
