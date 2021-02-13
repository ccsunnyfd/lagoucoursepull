package model

// LagouRequest LagouRequest
type LagouRequest struct {
	Method string
	Url    string
	Params string
}

// NewLagouRequestWithParams NewLagouRequestWithParams
func NewLagouRequestWithParams(method, url, params string) *LagouRequest {
	return &LagouRequest{
		Method: method,
		Url:    url,
		Params: params,
	}
}

// NewLagouRequestWithoutParams NewLagouRequestWithoutParams
func NewLagouRequestWithoutParams(method, url string) *LagouRequest {
	return &LagouRequest{
		Method: method,
		Url:    url,
	}
}
