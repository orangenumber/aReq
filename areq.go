// Written by Gon Yi
package areq

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

// New will create AReq object and takes PlugFns.
// However, before each request, Buf, Req, Log, FnsPre/Post will be reset.
func New(optionalFn ...func(d *AReq)) *AReq {
	r := &AReq{
		Ctx: context.Background(),
		Cli: http.Client{},
		Log: bytes.Buffer{},
		Buf: bytes.Buffer{},
		Mux: sync.Mutex{},
	}

	// Loading optional functions
	for _, f := range optionalFn {
		f(r)
	}
	return r
}

type AReq struct {
	Ctx       context.Context
	Cli       http.Client
	Req       *http.Request
	Res       *http.Response
	Buf       bytes.Buffer
	Log       bytes.Buffer
	Mux       sync.Mutex
	NoOutput  bool
	PlugFns   []PluginFn
	PlugFnErr []error
}

func (r *AReq) Request(method, url string, plugFn ...PluginFn) (statusCode int, err error) {
	// For each new Request, reset buffers
	r.Mux.Lock()
	defer r.Mux.Unlock()

	r.Buf.Reset()
	r.Log.Reset()
	r.ResetPluginFn()
	r.ResetPluginFnErrs()
	r.NoOutput = false

	// Load optional functions
	r.PlugFns = plugFn

	r.Req, err = http.NewRequestWithContext(r.Ctx, method, url, nil)
	if err != nil {
		return -1, err
	}
	// Apply Fn Pre
	for _, f := range r.PlugFns {
		if f.FnPre != nil {
			f.FnPre(r)
			r.Logf("[PlugFns.FnPre] %s applied", f.Name)
		}
	}
	r.Req.Header.Set("User-Agent", "AReq/0.4")
	r.Res, err = r.Cli.Do(r.Req)
	defer func() {
		if r.Res == nil {
			r.Log.WriteString("[AReq] no response received\n")
		} else if r.Res.Body != nil {
			if err = r.Res.Body.Close(); err != nil {
				r.Log.WriteString(err.Error())
			}
			r.Log.WriteString("[AReq] body closed\n")
		}
	}()

	if err != nil {
		return -1, err
	}
	if r.Res == nil {
		return -1, fmt.Errorf("no response received")
	}

	// Apply FnPost
	for _, f := range r.PlugFns {
		if f.FnPost != nil {
			f.FnPost(r)
			r.Logf("[PlugFns.FnPost] %s applied", f.Name)
		}
	}

	// If no FnPost, then save output to buffer
	if r.NoOutput == false {
		size, err := io.Copy(&r.Buf, r.Res.Body)
		r.Logf("[Save.ResBody] %d kb", size/1024)
		if err != nil {
			return -1, err
		}
	} else {
		if _, err = io.Copy(ioutil.Discard, r.Res.Body); err != nil {
			r.Log.WriteString(err.Error())
		}
		r.Log.WriteString("[Save.ResBody] No standard output (NoOutput=True)\n")
	}
	if err == nil && len(r.PlugFnErr) > 0 {
		err = fmt.Errorf("there are %d errors from plugin", len(r.PlugFnErr))
	}
	return r.Res.StatusCode, err
}

func (r *AReq) Logf(format string, a ...interface{}) {
	r.Log.WriteString(fmt.Sprintf(format, a...))
	r.Log.WriteRune('\n')
}

func (r *AReq) AddPluginFnErr(err error) {
	if err != nil {
		r.PlugFnErr = append(r.PlugFnErr, err)
	}
}

func (r *AReq) ResetPluginFnErrs() {
	r.PlugFnErr = r.PlugFnErr[:0]
}

func (r *AReq) ResetPluginFn() {
	r.PlugFns = r.PlugFns[:0]
}
