// Code generated by 'yaegi extract github.com/carlmjohnson/requests'. DO NOT EDIT.

// MIT License
//
// Copyright (c) 2020 - 2023 Jan-Lukas Else
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package yaegiwrappers

import (
	"github.com/carlmjohnson/requests"
	"net/http"
	"reflect"
)

func init() {
	Symbols["github.com/carlmjohnson/requests/requests"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"BodyBytes":          reflect.ValueOf(requests.BodyBytes),
		"BodyFile":           reflect.ValueOf(requests.BodyFile),
		"BodyForm":           reflect.ValueOf(requests.BodyForm),
		"BodyJSON":           reflect.ValueOf(requests.BodyJSON),
		"BodyReader":         reflect.ValueOf(requests.BodyReader),
		"BodyWriter":         reflect.ValueOf(requests.BodyWriter),
		"Caching":            reflect.ValueOf(requests.Caching),
		"ChainHandlers":      reflect.ValueOf(requests.ChainHandlers),
		"CheckContentType":   reflect.ValueOf(requests.CheckContentType),
		"CheckPeek":          reflect.ValueOf(requests.CheckPeek),
		"CheckStatus":        reflect.ValueOf(requests.CheckStatus),
		"CopyHeaders":        reflect.ValueOf(requests.CopyHeaders),
		"DefaultValidator":   reflect.ValueOf(&requests.DefaultValidator).Elem(),
		"ErrHandler":         reflect.ValueOf(requests.ErrHandler),
		"ErrInvalidHandled":  reflect.ValueOf(&requests.ErrInvalidHandled).Elem(),
		"ErrRequest":         reflect.ValueOf(requests.ErrRequest),
		"ErrTransport":       reflect.ValueOf(requests.ErrTransport),
		"ErrURL":             reflect.ValueOf(requests.ErrURL),
		"ErrValidator":       reflect.ValueOf(requests.ErrValidator),
		"ErrorJSON":          reflect.ValueOf(requests.ErrorJSON),
		"GzipConfig":         reflect.ValueOf(requests.GzipConfig),
		"HasStatusErr":       reflect.ValueOf(requests.HasStatusErr),
		"LogTransport":       reflect.ValueOf(requests.LogTransport),
		"MaxFollow":          reflect.ValueOf(requests.MaxFollow),
		"NewCookieJar":       reflect.ValueOf(requests.NewCookieJar),
		"NoFollow":           reflect.ValueOf(&requests.NoFollow).Elem(),
		"PermitURLTransport": reflect.ValueOf(requests.PermitURLTransport),
		"Record":             reflect.ValueOf(requests.Record),
		"Replay":             reflect.ValueOf(requests.Replay),
		"ReplayFS":           reflect.ValueOf(requests.ReplayFS),
		"ReplayString":       reflect.ValueOf(requests.ReplayString),
		"ToBufioReader":      reflect.ValueOf(requests.ToBufioReader),
		"ToBufioScanner":     reflect.ValueOf(requests.ToBufioScanner),
		"ToBytesBuffer":      reflect.ValueOf(requests.ToBytesBuffer),
		"ToFile":             reflect.ValueOf(requests.ToFile),
		"ToHTML":             reflect.ValueOf(requests.ToHTML),
		"ToHeaders":          reflect.ValueOf(&requests.ToHeaders).Elem(),
		"ToJSON":             reflect.ValueOf(requests.ToJSON),
		"ToString":           reflect.ValueOf(requests.ToString),
		"ToWriter":           reflect.ValueOf(requests.ToWriter),
		"URL":                reflect.ValueOf(requests.URL),
		"UserAgentTransport": reflect.ValueOf(requests.UserAgentTransport),
		"ValidatorHandler":   reflect.ValueOf(requests.ValidatorHandler),

		// type definitions
		"BodyGetter":          reflect.ValueOf((*requests.BodyGetter)(nil)),
		"Builder":             reflect.ValueOf((*requests.Builder)(nil)),
		"CheckRedirectPolicy": reflect.ValueOf((*requests.CheckRedirectPolicy)(nil)),
		"Config":              reflect.ValueOf((*requests.Config)(nil)),
		"ErrorKind":           reflect.ValueOf((*requests.ErrorKind)(nil)),
		"ResponseError":       reflect.ValueOf((*requests.ResponseError)(nil)),
		"ResponseHandler":     reflect.ValueOf((*requests.ResponseHandler)(nil)),
		"RoundTripFunc":       reflect.ValueOf((*requests.RoundTripFunc)(nil)),
		"Transport":           reflect.ValueOf((*requests.Transport)(nil)),

		// interface wrapper definitions
		"_Transport": reflect.ValueOf((*_github_com_carlmjohnson_requests_Transport)(nil)),
	}
}

// _github_com_carlmjohnson_requests_Transport is an interface wrapper for Transport type
type _github_com_carlmjohnson_requests_Transport struct {
	IValue     interface{}
	WRoundTrip func(a0 *http.Request) (*http.Response, error)
}

func (W _github_com_carlmjohnson_requests_Transport) RoundTrip(a0 *http.Request) (*http.Response, error) {
	return W.WRoundTrip(a0)
}
