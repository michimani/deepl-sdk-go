package deepl

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/michimani/deepl-sdk-go/types"
)

const contentType string = "application/x-www-form-urlencoded"

type Requester struct {
	Endpoint string
	Params   types.RequestParams
	Body     io.Reader
}

func NewRequester(e string, p types.RequestParams) *Requester {
	return &Requester{
		Endpoint: e,
		Params:   p,
	}
}

func (r *Requester) Exec(res interface{}) (*types.ErrorResponse, error) {
	if err := r.prepare(); err != nil {
		return nil, err
	}

	raw, err := r.post()
	if err != nil {
		return nil, err
	}

	return raw.Unmarshal(&res)
}

func (r *Requester) prepare() error {
	body, err := r.Params.Body()
	if err != nil {
		return err
	}

	r.Body = body
	return nil
}

// POST is creating POST request to the endpoint with body.
func (r *Requester) post() (*types.RawResponse, error) {
	rr := types.RawResponse{
		StatusCode: http.StatusInternalServerError,
		Body:       nil,
	}

	req, err := http.NewRequest("POST", r.Endpoint, r.Body)
	if err != nil {
		return &rr, err
	}

	req.Header.Set("Content-Type", contentType)
	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return &rr, err
	}

	defer resp.Body.Close()
	if err != nil {
		return &rr, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &rr, err
	}

	rr.StatusCode = resp.StatusCode
	rr.Body = bytes
	return &rr, nil
}
