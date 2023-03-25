// Code generated by ogen, DO NOT EDIT.

package v1

import (
	"fmt"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type AuthLoginPostReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// GetLogin returns the value of Login.
func (s *AuthLoginPostReq) GetLogin() string {
	return s.Login
}

// GetPassword returns the value of Password.
func (s *AuthLoginPostReq) GetPassword() string {
	return s.Password
}

// SetLogin sets the value of Login.
func (s *AuthLoginPostReq) SetLogin(val string) {
	s.Login = val
}

// SetPassword sets the value of Password.
func (s *AuthLoginPostReq) SetPassword(val string) {
	s.Password = val
}

type BearerAdminAuth struct {
	Token string
}

// GetToken returns the value of Token.
func (s *BearerAdminAuth) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *BearerAdminAuth) SetToken(val string) {
	s.Token = val
}

type BearerAuth struct {
	Token string
}

// GetToken returns the value of Token.
func (s *BearerAuth) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *BearerAuth) SetToken(val string) {
	s.Token = val
}

// Ref: #/components/schemas/Error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetCode returns the value of Code.
func (s *Error) GetCode() int {
	return s.Code
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetCode sets the value of Code.
func (s *Error) SetCode(val int) {
	s.Code = val
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// Ref: #/components/schemas/Jwt
type Jwt struct {
	Token string `json:"token"`
}

// GetToken returns the value of Token.
func (s *Jwt) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *Jwt) SetToken(val string) {
	s.Token = val
}

// NewOptQuery returns new OptQuery with value set to v.
func NewOptQuery(v Query) OptQuery {
	return OptQuery{
		Value: v,
		Set:   true,
	}
}

// OptQuery is optional Query.
type OptQuery struct {
	Value Query
	Set   bool
}

// IsSet returns true if OptQuery was set.
func (o OptQuery) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptQuery) Reset() {
	var v Query
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptQuery) SetTo(v Query) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptQuery) Get() (v Query, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptQuery) Or(d Query) Query {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptTaskPostReq returns new OptTaskPostReq with value set to v.
func NewOptTaskPostReq(v TaskPostReq) OptTaskPostReq {
	return OptTaskPostReq{
		Value: v,
		Set:   true,
	}
}

// OptTaskPostReq is optional TaskPostReq.
type OptTaskPostReq struct {
	Value TaskPostReq
	Set   bool
}

// IsSet returns true if OptTaskPostReq was set.
func (o OptTaskPostReq) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptTaskPostReq) Reset() {
	var v TaskPostReq
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptTaskPostReq) SetTo(v TaskPostReq) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptTaskPostReq) Get() (v TaskPostReq, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptTaskPostReq) Or(d TaskPostReq) TaskPostReq {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Query
type Query struct {
	ID           int64  `json:"id"`
	QueryRow     string `json:"query_row"`
	QueryHash    string `json:"query_hash"`
	ResponseRaw  string `json:"response_raw"`
	ResponseHash string `json:"response_hash"`
}

// GetID returns the value of ID.
func (s *Query) GetID() int64 {
	return s.ID
}

// GetQueryRow returns the value of QueryRow.
func (s *Query) GetQueryRow() string {
	return s.QueryRow
}

// GetQueryHash returns the value of QueryHash.
func (s *Query) GetQueryHash() string {
	return s.QueryHash
}

// GetResponseRaw returns the value of ResponseRaw.
func (s *Query) GetResponseRaw() string {
	return s.ResponseRaw
}

// GetResponseHash returns the value of ResponseHash.
func (s *Query) GetResponseHash() string {
	return s.ResponseHash
}

// SetID sets the value of ID.
func (s *Query) SetID(val int64) {
	s.ID = val
}

// SetQueryRow sets the value of QueryRow.
func (s *Query) SetQueryRow(val string) {
	s.QueryRow = val
}

// SetQueryHash sets the value of QueryHash.
func (s *Query) SetQueryHash(val string) {
	s.QueryHash = val
}

// SetResponseRaw sets the value of ResponseRaw.
func (s *Query) SetResponseRaw(val string) {
	s.ResponseRaw = val
}

// SetResponseHash sets the value of ResponseHash.
func (s *Query) SetResponseHash(val string) {
	s.ResponseHash = val
}

// Ref: #/components/schemas/Task
type Task struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Query       OptQuery `json:"query"`
}

// GetID returns the value of ID.
func (s *Task) GetID() int64 {
	return s.ID
}

// GetName returns the value of Name.
func (s *Task) GetName() string {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *Task) GetDescription() string {
	return s.Description
}

// GetQuery returns the value of Query.
func (s *Task) GetQuery() OptQuery {
	return s.Query
}

// SetID sets the value of ID.
func (s *Task) SetID(val int64) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Task) SetName(val string) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *Task) SetDescription(val string) {
	s.Description = val
}

// SetQuery sets the value of Query.
func (s *Task) SetQuery(val OptQuery) {
	s.Query = val
}

type TaskPostReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	QueryRaw    string `json:"query_raw"`
}

// GetName returns the value of Name.
func (s *TaskPostReq) GetName() string {
	return s.Name
}

// GetDescription returns the value of Description.
func (s *TaskPostReq) GetDescription() string {
	return s.Description
}

// GetQueryRaw returns the value of QueryRaw.
func (s *TaskPostReq) GetQueryRaw() string {
	return s.QueryRaw
}

// SetName sets the value of Name.
func (s *TaskPostReq) SetName(val string) {
	s.Name = val
}

// SetDescription sets the value of Description.
func (s *TaskPostReq) SetDescription(val string) {
	s.Description = val
}

// SetQueryRaw sets the value of QueryRaw.
func (s *TaskPostReq) SetQueryRaw(val string) {
	s.QueryRaw = val
}

// Ref: #/components/schemas/User
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Login     string `json:"login"`
}

// GetID returns the value of ID.
func (s *User) GetID() int64 {
	return s.ID
}

// GetFirstName returns the value of FirstName.
func (s *User) GetFirstName() string {
	return s.FirstName
}

// GetLastName returns the value of LastName.
func (s *User) GetLastName() string {
	return s.LastName
}

// GetLogin returns the value of Login.
func (s *User) GetLogin() string {
	return s.Login
}

// SetID sets the value of ID.
func (s *User) SetID(val int64) {
	s.ID = val
}

// SetFirstName sets the value of FirstName.
func (s *User) SetFirstName(val string) {
	s.FirstName = val
}

// SetLastName sets the value of LastName.
func (s *User) SetLastName(val string) {
	s.LastName = val
}

// SetLogin sets the value of Login.
func (s *User) SetLogin(val string) {
	s.Login = val
}
