// Code generated by ogen, DO NOT EDIT.

package v1

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
)

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	sec       SecuritySource
	baseClient
}
type errorHandler interface {
	NewError(ctx context.Context, err error) *ErrStatusCode
}

var _ Handler = struct {
	errorHandler
	*Client
}{}

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, sec SecuritySource, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		sec:        sec,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// AttemptAttemptIDGet invokes GET /attempt/{attempt_id} operation.
//
// GET /attempt/{attempt_id}
func (c *Client) AttemptAttemptIDGet(ctx context.Context, params AttemptAttemptIDGetParams) (*Attempt, error) {
	res, err := c.sendAttemptAttemptIDGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendAttemptAttemptIDGet(ctx context.Context, params AttemptAttemptIDGetParams) (res *Attempt, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AttemptAttemptIDGet",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/attempt/"
	{
		// Encode "attempt_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "attempt_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.AttemptID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "AttemptAttemptIDGet", r); err {
			case nil:
				satisfied[0] |= 1 << 0
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, errors.New("no security requirement satisfied")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAttemptAttemptIDGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AuthLoginPost invokes POST /auth/login operation.
//
// POST /auth/login
func (c *Client) AuthLoginPost(ctx context.Context, request *AuthLoginPostReq) (*Jwt, error) {
	res, err := c.sendAuthLoginPost(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendAuthLoginPost(ctx context.Context, request *AuthLoginPostReq) (res *Jwt, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AuthLoginPost",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/auth/login"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAuthLoginPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAuthLoginPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AuthVerifyGet invokes GET /auth/verify operation.
//
// GET /auth/verify
func (c *Client) AuthVerifyGet(ctx context.Context) (*OkResponse, error) {
	res, err := c.sendAuthVerifyGet(ctx)
	_ = res
	return res, err
}

func (c *Client) sendAuthVerifyGet(ctx context.Context) (res *OkResponse, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "AuthVerifyGet",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/auth/verify"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "AuthVerifyGet", r); err {
			case nil:
				satisfied[0] |= 1 << 0
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, errors.New("no security requirement satisfied")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeAuthVerifyGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TaskPost invokes POST /task operation.
//
// POST /task
func (c *Client) TaskPost(ctx context.Context, request *TaskPostReq) (*Task, error) {
	res, err := c.sendTaskPost(ctx, request)
	_ = res
	return res, err
}

func (c *Client) sendTaskPost(ctx context.Context, request *TaskPostReq) (res *Task, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "TaskPost",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/task"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeTaskPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAdminAuth"
			switch err := c.securityBearerAdminAuth(ctx, "TaskPost", r); err {
			case nil:
				satisfied[0] |= 1 << 0
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAdminAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, errors.New("no security requirement satisfied")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTaskPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TaskTaskIDAttemptPost invokes POST /task/{task_id}/attempt operation.
//
// POST /task/{task_id}/attempt
func (c *Client) TaskTaskIDAttemptPost(ctx context.Context, request *TaskTaskIDAttemptPostReq, params TaskTaskIDAttemptPostParams) (*Attempt, error) {
	res, err := c.sendTaskTaskIDAttemptPost(ctx, request, params)
	_ = res
	return res, err
}

func (c *Client) sendTaskTaskIDAttemptPost(ctx context.Context, request *TaskTaskIDAttemptPostReq, params TaskTaskIDAttemptPostParams) (res *Attempt, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "TaskTaskIDAttemptPost",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/task/"
	{
		// Encode "task_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "task_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.TaskID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/attempt"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeTaskTaskIDAttemptPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "TaskTaskIDAttemptPost", r); err {
			case nil:
				satisfied[0] |= 1 << 0
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, errors.New("no security requirement satisfied")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTaskTaskIDAttemptPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TaskTaskIDAttemptsGet invokes GET /task/{task_id}/attempts operation.
//
// GET /task/{task_id}/attempts
func (c *Client) TaskTaskIDAttemptsGet(ctx context.Context, params TaskTaskIDAttemptsGetParams) ([]Attempt, error) {
	res, err := c.sendTaskTaskIDAttemptsGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendTaskTaskIDAttemptsGet(ctx context.Context, params TaskTaskIDAttemptsGetParams) (res []Attempt, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "TaskTaskIDAttemptsGet",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/task/"
	{
		// Encode "task_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "task_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.TaskID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/attempts"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "TaskTaskIDAttemptsGet", r); err {
			case nil:
				satisfied[0] |= 1 << 0
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, errors.New("no security requirement satisfied")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTaskTaskIDAttemptsGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TaskTaskIDGet invokes GET /task/{task_id} operation.
//
// GET /task/{task_id}
func (c *Client) TaskTaskIDGet(ctx context.Context, params TaskTaskIDGetParams) (*Task, error) {
	res, err := c.sendTaskTaskIDGet(ctx, params)
	_ = res
	return res, err
}

func (c *Client) sendTaskTaskIDGet(ctx context.Context, params TaskTaskIDGetParams) (res *Task, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "TaskTaskIDGet",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/task/"
	{
		// Encode "task_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "task_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.TaskID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "TaskTaskIDGet", r); err {
			case nil:
				satisfied[0] |= 1 << 0
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, errors.New("no security requirement satisfied")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTaskTaskIDGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TasksGet invokes GET /tasks operation.
//
// GET /tasks
func (c *Client) TasksGet(ctx context.Context) ([]Task, error) {
	res, err := c.sendTasksGet(ctx)
	_ = res
	return res, err
}

func (c *Client) sendTasksGet(ctx context.Context) (res []Task, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "TasksGet",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/tasks"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "TasksGet", r); err {
			case nil:
				satisfied[0] |= 1 << 0
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, errors.New("no security requirement satisfied")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeTasksGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UserGet invokes GET /user operation.
//
// GET /user
func (c *Client) UserGet(ctx context.Context) (*User, error) {
	res, err := c.sendUserGet(ctx)
	_ = res
	return res, err
}

func (c *Client) sendUserGet(ctx context.Context) (res *User, err error) {
	var otelAttrs []attribute.KeyValue

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, elapsedDuration.Microseconds(), otelAttrs...)
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, otelAttrs...)

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, "UserGet",
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, otelAttrs...)
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/user"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u, nil)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	{
		type bitset = [1]uint8
		var satisfied bitset
		{
			stage = "Security:BearerAuth"
			switch err := c.securityBearerAuth(ctx, "UserGet", r); err {
			case nil:
				satisfied[0] |= 1 << 0
			case ogenerrors.ErrSkipClientSecurity:
				// Skip this security.
			default:
				return res, errors.Wrap(err, "security \"BearerAuth\"")
			}
		}

		if ok := func() bool {
		nextRequirement:
			for _, requirement := range []bitset{
				{0b00000001},
			} {
				for i, mask := range requirement {
					if satisfied[i]&mask != mask {
						continue nextRequirement
					}
				}
				return true
			}
			return false
		}(); !ok {
			return res, errors.New("no security requirement satisfied")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()

	stage = "DecodeResponse"
	result, err := decodeUserGetResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
