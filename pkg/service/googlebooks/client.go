package googlebooks

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

// Client is a client for the Google Books API
type Client struct {
	*resty.Client
}

func NewClient() *Client {
	client := resty.New().
		SetDebug(false).
		SetDisableWarn(true).
		SetLogger(log.StandardLogger()).
		SetRetryCount(2).
		SetRetryWaitTime(500 * time.Millisecond).
		SetRetryMaxWaitTime(20 * time.Second).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r.IsError()
			},
		).
		OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
			log.WithFields(log.Fields{
				"url":    r.Request.URL,
				"status": r.Status(),
				"time":   r.Time(),
			}).Info("Response")
			return nil
		}).
		SetBaseURL("https://www.googleapis.com/books/v1")
	return &Client{
		Client: client,
	}
}

type RequestSearchBooks func(q string) (*Response, error)

func NewRequestSearchBooks(c *Client) RequestSearchBooks {
	return func(q string) (*Response, error) {
		books := new(Response)
		r, err := c.R().
			SetResult(books).
			SetQueryParams(map[string]string{
				"q": q,
			}).
			Get("/volumes")

		if err != nil {
			return nil, err
		}

		if r.IsError() {
			return handleError(r)
		}

		return books, nil
	}
}

func handleError(r *resty.Response) (*Response, error) {
	if strings.Contains(r.Header().Get("Content-Type"), "application/json") {
		var httpGoogleError ErrorResponse
		if err := json.Unmarshal(r.Body(), &httpGoogleError); err != nil {
			return nil, err
		}
		return nil, errors.Errorf("[URL: %s, StatusCode: %d, Message: %s]",
			r.Request.URL, r.StatusCode(), httpGoogleError.Error.Message)
	}
	return nil, errors.New("unknown error (no JSON)")
}
