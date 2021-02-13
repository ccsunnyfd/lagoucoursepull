package utils

import (
	"lagoucoursepull/internal/model"
	"net/http"
)

// LagouHTTPRequest LagouHTTPRequest
type LagouHTTPRequest struct {
	r *http.Request
}

// NewLagouHTTPRequest NewLagouHTTPRequest
func NewLagouHTTPRequest(cr *model.LagouRequest) (*LagouHTTPRequest, error) {
	req, err := http.NewRequest(cr.Method, cr.Url+"?"+cr.Params, nil)
	if err != nil {
		return nil, err
	}
	return &LagouHTTPRequest{
		r: req,
	}, nil
}

// SetHeader SetHeader
func (rq *LagouHTTPRequest) SetHeader() {
	rq.r.Header.Set("origin", "https://kaiwu.lagou.com")
	rq.r.Header.Set("referer", "https://kaiwu.lagou.com/")
	rq.r.Header.Set("sec-ch-ua", `"Chromium";v="88", "Google Chrome";v="88", ";Not A Brand";v="99"`)
	rq.r.Header.Set("sec-ch-ua-mobile", `?0`)
	rq.r.Header.Set("sec-fetch-dest", `empty`)
	rq.r.Header.Set("sec-fetch-mode", `cors`)
	rq.r.Header.Set("sec-fetch-site", `same-site`)
	rq.r.Header.Set("user-agent", `Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36`)
	rq.r.Header.Set("x-l-req-header", `{"deviceType":1}`)
}

// SetAuth SetAuth
func (rq *LagouHTTPRequest) SetAuth(cookie string, auth string) {
	rq.r.Header.Set("cookie", cookie)
	rq.r.Header.Set("authorization", auth)
}
