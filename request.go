package go_mojoauth

import (
	"errors"

	"github.com/mojoauth/go-sdk/httprutils"
	"github.com/mojoauth/go-sdk/mojoerror"
)

var URLEncodedHeader = map[string]string{"content-Type": "application/x-www-form-urlencoded"}

var JSONHeader = map[string]string{"content-Type": "application/json"}

// NewGetRequest takes a uri and query parameters, then constructs a GET request for MojoAuth API endpoints requiring access tokens
// being passed in Authorization Bearer header
func (mojo Mojoauth) NewGetReqWithToken(path string, queries ...map[string]string) (*httprutils.Request, error) {
	if mojo.Context.Token == "" {
		errMsg := "Must initialize MojoAuth with access token for this API call."
		err := mojoerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	request := &httprutils.Request{
		Method: httprutils.Get,
		URL:    mojo.Domain + path,
		Headers: map[string]string{
			"content-Type":  "application/x-www-form-urlencoded",
			"Authorization": "Bearer " + mojo.Context.Token,
		},
		QueryParams: map[string]string{
			"X-API-Key": mojo.Context.ApiKey,
		},
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}

	return request, nil
}

// NewGetRequest takes a uri and query parameters, then constructs a GET request for a MojoAuth API endpoint
func (mojo Mojoauth) NewGetReq(path string, queries ...map[string]string) *httprutils.Request {
	request := &httprutils.Request{
		Method:      httprutils.Get,
		URL:         mojo.Domain + path,
		Headers: map[string]string{
			"content-Type":  "application/x-www-form-urlencoded",
			"X-API-Key": mojo.Context.ApiKey,
		},
		QueryParams: map[string]string{},
	}
	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}
	return request
}

// NewPostReqWithToken takes a uri, body, and query parameters, then constructs the request for MojoAuth PUT API end points requiring access tokens being passed in Authorization Bearer header
func (mojo Mojoauth) NewPostReqWithToken(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {

	if mojo.Context.Token == "" {
		errMsg := "Must initialize MojoAuth with access token for this API call."
		err := mojoerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Post,
		URL:    mojo.Domain + path,
		Headers: map[string]string{
			"content-Type":  "application/json",
			"Authorization": "Bearer " + mojo.Context.Token,
			"X-API-Key":     mojo.Context.ApiKey,
		},
		Body: encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}

	return request, nil
}

// NewPostReq takes a uri, body, and optional queries to construct a POST request for a MojoAuth POST API endpoint
func (mojo Mojoauth) NewPostReq(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {
	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Post,
		URL:    mojo.Domain + path,
		Headers: map[string]string{
			"X-API-Key":    mojo.Context.ApiKey,
			"content-Type": "application/json",
		},
		QueryParams: map[string]string{},
		Body:        encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}
	return request, nil
}

// NewPutReq takes a uri, body, and optional queries to construct a PUT request for a MojoAuth API endpoint
func (mojo Mojoauth) NewPutReq(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {
	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Put,
		URL:    mojo.Domain + path,
		Headers: map[string]string{
			"content-Type": "application/json",
		},
		QueryParams: map[string]string{
			"X-API-Key": mojo.Context.ApiKey,
		},
		Body: encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}
	return request, nil
}

// NewPutReqWithToken takes a uri and query parameters, then constructs a PUT request for MojoAuth API endpoints requiring access tokens
// being passed in Authorization Bearer header
func (mojo Mojoauth) NewPutReqWithToken(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {
	if mojo.Context.Token == "" {
		errMsg := "Must initialize MojoAuth with access token for this API call."
		err := mojoerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Put,
		URL:    mojo.Domain + path,
		Headers: map[string]string{
			"content-Type":  "application/json",
			"Authorization": "Bearer " + mojo.Context.Token,
		},
		QueryParams: map[string]string{
			"X-API-Key": mojo.Context.ApiKey,
		},
		Body: encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}
	return request, nil
}

// NewDeleteReq takes a uri, body, and optional queries to construct a DELETE request for a MojoAuth POST API endpoint
func (mojo Mojoauth) NewDeleteReq(path string, body ...interface{}) *httprutils.Request {
	if len(body) != 0 {
		encoded, err := httprutils.EncodeBody(body[0])
		if err != nil {
			return nil
		}
		return &httprutils.Request{
			Method:  httprutils.Delete,
			URL:     mojo.Domain + path,
			Headers: URLEncodedHeader,
			Body:    encoded,
		}
	} else {
		return &httprutils.Request{
			Method:  httprutils.Delete,
			URL:     mojo.Domain + path,
			Headers: URLEncodedHeader,
		}
	}
}

// NewDeleteReqWithToken takes a uri and query parameters, then constructs a PUT request for MojoAuth API endpoints requiring access tokens
// being passed in Authorization Bearer header
func (mojo Mojoauth) NewDeleteReqWithToken(path string, body interface{}, queries ...map[string]string) (*httprutils.Request, error) {
	if mojo.Context.Token == "" {
		errMsg := "Must initialize MojoAuth with access token for this API call."
		err := mojoerror.New("MissingTokenErr", errMsg, errors.New(errMsg))
		return nil, err
	}

	encodedBody, error := httprutils.EncodeBody(body)
	if error != nil {
		return nil, error
	}

	request := &httprutils.Request{
		Method: httprutils.Delete,
		URL:    mojo.Domain + path,
		Headers: map[string]string{
			"content-Type":  "application/json",
			"Authorization": "Bearer " + mojo.Context.Token,
		},
		QueryParams: map[string]string{
			"X-API-Key": mojo.Context.ApiKey,
		},
		Body: encodedBody,
	}

	for _, q := range queries {
		for k, v := range q {
			request.QueryParams[k] = v
		}
	}

	return request, nil
}

// AddApiCredentialsToReqHeader removes the X-API-Key query parameter from a constructed request
// and add MojoAuth app credentials in the request headers
func (mojo Mojoauth) AddApiCredentialsToReqHeader(req *httprutils.Request) {
	delete(req.QueryParams, "X-API-Key")
	req.Headers["X-API-Key"] = mojo.Context.ApiKey
}

// NormalizeApiKey normalizes the apikey parameter in queries for requests to be sent to
// MojoAuth endpoints that only accept "apikey"
func (mojo Mojoauth) NormalizeApiKey(req *httprutils.Request) {
	delete(req.QueryParams, "X-API-Key")
	req.QueryParams["apikey"] = mojo.Context.ApiKey
}
