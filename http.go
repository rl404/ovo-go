package ovo

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Requester is http request interface.
type Requester interface {
	Call(ctx context.Context, method, url, appID, key string, header http.Header, request interface{}, response interface{}) (statusCode int, err error)
}

type requester struct {
	client *http.Client
	logger Logger
}

func defaultRequester(client *http.Client, logger Logger) *requester {
	return &requester{
		client: client,
		logger: logger,
	}
}

// Call to prepare request and execute.
func (r *requester) Call(ctx context.Context, method, url, appID, key string, header http.Header, request interface{}, response interface{}) (int, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)

	reqBody, err := json.Marshal(request)
	if err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}

	if header != nil {
		req.Header = header
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(fmt.Sprintf("%s%d", appID, now.Unix())))
	sha := hex.EncodeToString(h.Sum(nil))

	req.Header.Add("app-id", appID)
	req.Header.Add("random", strconv.FormatInt(now.Unix(), 10))
	req.Header.Add("hmac", sha)

	r.logger.Debug("%s %s", method, url)
	r.logRequestHeader(req.Header)
	r.logRequestBody(reqBody)
	defer func() { r.logger.Info("%s %s [%s]", method, url, time.Since(now)) }()

	return r.doRequest(req, response)
}

func (r *requester) doRequest(req *http.Request, response interface{}) (int, error) {
	resp, err := r.client.Do(req)
	if err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}

	r.logResponseBody(resp.StatusCode, respBody)

	if resp.StatusCode != http.StatusOK {
		var tx Transaction
		if err := json.Unmarshal(respBody, &tx); err != nil {
			r.logger.Error(err.Error())
			return resp.StatusCode, errors.New(string(respBody))
		}
		return resp.StatusCode, ovoErr[tx.ResponseCode]
	}

	if err := json.Unmarshal(respBody, &response); err != nil {
		r.logger.Error(err.Error())
		return http.StatusInternalServerError, ErrInternal
	}

	return resp.StatusCode, nil
}

func (r *requester) logRequestHeader(header http.Header) {
	if header == nil || len(header) == 0 {
		return
	}

	for k, h := range header {
		for _, v := range h {
			r.logger.Debug("header: %s: %s", k, v)
		}
	}
}

func (r *requester) logRequestBody(request []byte) {
	if request == nil {
		return
	}

	var out bytes.Buffer
	if err := json.Indent(&out, request, "", "  "); err != nil {
		r.logger.Error(err.Error())
		r.logger.Debug("request: %s", string(request))
		return
	}

	r.logger.Debug("request: %s", out.String())
}

func (r *requester) logResponseBody(code int, response []byte) {
	if response == nil {
		return
	}

	var out bytes.Buffer
	if err := json.Indent(&out, response, "", "  "); err != nil {
		r.logger.Error(err.Error())
		r.logger.Debug("response: %d %s", code, string(response))
		return
	}

	r.logger.Debug("response: %d %s", code, out.String())
}
