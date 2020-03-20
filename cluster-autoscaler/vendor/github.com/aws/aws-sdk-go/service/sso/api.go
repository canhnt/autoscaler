// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sso

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

const opGetRoleCredentials = "GetRoleCredentials"

// GetRoleCredentialsRequest generates a "aws/request.Request" representing the
// client's request for the GetRoleCredentials operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetRoleCredentials for more information on using the GetRoleCredentials
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetRoleCredentialsRequest method.
//    req, resp := client.GetRoleCredentialsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sso-2019-06-10/GetRoleCredentials
func (c *SSO) GetRoleCredentialsRequest(input *GetRoleCredentialsInput) (req *request.Request, output *GetRoleCredentialsOutput) {
	op := &request.Operation{
		Name:       opGetRoleCredentials,
		HTTPMethod: "GET",
		HTTPPath:   "/federation/credentials",
	}

	if input == nil {
		input = &GetRoleCredentialsInput{}
	}

	output = &GetRoleCredentialsOutput{}
	req = c.newRequest(op, input, output)
	req.Config.Credentials = credentials.AnonymousCredentials
	return
}

// GetRoleCredentials API operation for AWS Single Sign-On.
//
// Returns the STS short-term credentials for a given role name that is assigned
// to the user.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Single Sign-On's
// API operation GetRoleCredentials for usage and error information.
//
// Returned Error Types:
//   * InvalidRequestException
//   Indicates that a problem occurred with the input to the request. For example,
//   a required parameter might be missing or out of range.
//
//   * UnauthorizedException
//   Indicates that the request is not authorized. This can happen due to an invalid
//   access token in the request.
//
//   * TooManyRequestsException
//   Indicates that the request is being made too frequently and is more than
//   what the server can handle.
//
//   * ResourceNotFoundException
//   The specified resource doesn't exist.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sso-2019-06-10/GetRoleCredentials
func (c *SSO) GetRoleCredentials(input *GetRoleCredentialsInput) (*GetRoleCredentialsOutput, error) {
	req, out := c.GetRoleCredentialsRequest(input)
	return out, req.Send()
}

// GetRoleCredentialsWithContext is the same as GetRoleCredentials with the addition of
// the ability to pass a context and additional request options.
//
// See GetRoleCredentials for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SSO) GetRoleCredentialsWithContext(ctx aws.Context, input *GetRoleCredentialsInput, opts ...request.Option) (*GetRoleCredentialsOutput, error) {
	req, out := c.GetRoleCredentialsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAccountRoles = "ListAccountRoles"

// ListAccountRolesRequest generates a "aws/request.Request" representing the
// client's request for the ListAccountRoles operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListAccountRoles for more information on using the ListAccountRoles
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListAccountRolesRequest method.
//    req, resp := client.ListAccountRolesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sso-2019-06-10/ListAccountRoles
func (c *SSO) ListAccountRolesRequest(input *ListAccountRolesInput) (req *request.Request, output *ListAccountRolesOutput) {
	op := &request.Operation{
		Name:       opListAccountRoles,
		HTTPMethod: "GET",
		HTTPPath:   "/assignment/roles",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAccountRolesInput{}
	}

	output = &ListAccountRolesOutput{}
	req = c.newRequest(op, input, output)
	req.Config.Credentials = credentials.AnonymousCredentials
	return
}

// ListAccountRoles API operation for AWS Single Sign-On.
//
// Lists all roles that are assigned to the user for a given AWS account.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Single Sign-On's
// API operation ListAccountRoles for usage and error information.
//
// Returned Error Types:
//   * InvalidRequestException
//   Indicates that a problem occurred with the input to the request. For example,
//   a required parameter might be missing or out of range.
//
//   * UnauthorizedException
//   Indicates that the request is not authorized. This can happen due to an invalid
//   access token in the request.
//
//   * TooManyRequestsException
//   Indicates that the request is being made too frequently and is more than
//   what the server can handle.
//
//   * ResourceNotFoundException
//   The specified resource doesn't exist.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sso-2019-06-10/ListAccountRoles
func (c *SSO) ListAccountRoles(input *ListAccountRolesInput) (*ListAccountRolesOutput, error) {
	req, out := c.ListAccountRolesRequest(input)
	return out, req.Send()
}

// ListAccountRolesWithContext is the same as ListAccountRoles with the addition of
// the ability to pass a context and additional request options.
//
// See ListAccountRoles for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SSO) ListAccountRolesWithContext(ctx aws.Context, input *ListAccountRolesInput, opts ...request.Option) (*ListAccountRolesOutput, error) {
	req, out := c.ListAccountRolesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListAccountRolesPages iterates over the pages of a ListAccountRoles operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListAccountRoles method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListAccountRoles operation.
//    pageNum := 0
//    err := client.ListAccountRolesPages(params,
//        func(page *sso.ListAccountRolesOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *SSO) ListAccountRolesPages(input *ListAccountRolesInput, fn func(*ListAccountRolesOutput, bool) bool) error {
	return c.ListAccountRolesPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListAccountRolesPagesWithContext same as ListAccountRolesPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SSO) ListAccountRolesPagesWithContext(ctx aws.Context, input *ListAccountRolesInput, fn func(*ListAccountRolesOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListAccountRolesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListAccountRolesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*ListAccountRolesOutput), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

const opListAccounts = "ListAccounts"

// ListAccountsRequest generates a "aws/request.Request" representing the
// client's request for the ListAccounts operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListAccounts for more information on using the ListAccounts
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListAccountsRequest method.
//    req, resp := client.ListAccountsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sso-2019-06-10/ListAccounts
func (c *SSO) ListAccountsRequest(input *ListAccountsInput) (req *request.Request, output *ListAccountsOutput) {
	op := &request.Operation{
		Name:       opListAccounts,
		HTTPMethod: "GET",
		HTTPPath:   "/assignment/accounts",
		Paginator: &request.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAccountsInput{}
	}

	output = &ListAccountsOutput{}
	req = c.newRequest(op, input, output)
	req.Config.Credentials = credentials.AnonymousCredentials
	return
}

// ListAccounts API operation for AWS Single Sign-On.
//
// Lists all AWS accounts assigned to the user. These AWS accounts are assigned
// by the administrator of the account. For more information, see Assign User
// Access (https://docs.aws.amazon.com/singlesignon/latest/userguide/useraccess.html#assignusers)
// in the AWS SSO User Guide. This operation returns a paginated response.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Single Sign-On's
// API operation ListAccounts for usage and error information.
//
// Returned Error Types:
//   * InvalidRequestException
//   Indicates that a problem occurred with the input to the request. For example,
//   a required parameter might be missing or out of range.
//
//   * UnauthorizedException
//   Indicates that the request is not authorized. This can happen due to an invalid
//   access token in the request.
//
//   * TooManyRequestsException
//   Indicates that the request is being made too frequently and is more than
//   what the server can handle.
//
//   * ResourceNotFoundException
//   The specified resource doesn't exist.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sso-2019-06-10/ListAccounts
func (c *SSO) ListAccounts(input *ListAccountsInput) (*ListAccountsOutput, error) {
	req, out := c.ListAccountsRequest(input)
	return out, req.Send()
}

// ListAccountsWithContext is the same as ListAccounts with the addition of
// the ability to pass a context and additional request options.
//
// See ListAccounts for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SSO) ListAccountsWithContext(ctx aws.Context, input *ListAccountsInput, opts ...request.Option) (*ListAccountsOutput, error) {
	req, out := c.ListAccountsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// ListAccountsPages iterates over the pages of a ListAccounts operation,
// calling the "fn" function with the response data for each page. To stop
// iterating, return false from the fn function.
//
// See ListAccounts method for more information on how to use this operation.
//
// Note: This operation can generate multiple requests to a service.
//
//    // Example iterating over at most 3 pages of a ListAccounts operation.
//    pageNum := 0
//    err := client.ListAccountsPages(params,
//        func(page *sso.ListAccountsOutput, lastPage bool) bool {
//            pageNum++
//            fmt.Println(page)
//            return pageNum <= 3
//        })
//
func (c *SSO) ListAccountsPages(input *ListAccountsInput, fn func(*ListAccountsOutput, bool) bool) error {
	return c.ListAccountsPagesWithContext(aws.BackgroundContext(), input, fn)
}

// ListAccountsPagesWithContext same as ListAccountsPages except
// it takes a Context and allows setting request options on the pages.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SSO) ListAccountsPagesWithContext(ctx aws.Context, input *ListAccountsInput, fn func(*ListAccountsOutput, bool) bool, opts ...request.Option) error {
	p := request.Pagination{
		NewRequest: func() (*request.Request, error) {
			var inCpy *ListAccountsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.ListAccountsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}

	for p.Next() {
		if !fn(p.Page().(*ListAccountsOutput), !p.HasNextPage()) {
			break
		}
	}

	return p.Err()
}

const opLogout = "Logout"

// LogoutRequest generates a "aws/request.Request" representing the
// client's request for the Logout operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See Logout for more information on using the Logout
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the LogoutRequest method.
//    req, resp := client.LogoutRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sso-2019-06-10/Logout
func (c *SSO) LogoutRequest(input *LogoutInput) (req *request.Request, output *LogoutOutput) {
	op := &request.Operation{
		Name:       opLogout,
		HTTPMethod: "POST",
		HTTPPath:   "/logout",
	}

	if input == nil {
		input = &LogoutInput{}
	}

	output = &LogoutOutput{}
	req = c.newRequest(op, input, output)
	req.Config.Credentials = credentials.AnonymousCredentials
	req.Handlers.Unmarshal.Swap(restjson.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// Logout API operation for AWS Single Sign-On.
//
// Removes the client- and server-side session that is associated with the user.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS Single Sign-On's
// API operation Logout for usage and error information.
//
// Returned Error Types:
//   * InvalidRequestException
//   Indicates that a problem occurred with the input to the request. For example,
//   a required parameter might be missing or out of range.
//
//   * UnauthorizedException
//   Indicates that the request is not authorized. This can happen due to an invalid
//   access token in the request.
//
//   * TooManyRequestsException
//   Indicates that the request is being made too frequently and is more than
//   what the server can handle.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sso-2019-06-10/Logout
func (c *SSO) Logout(input *LogoutInput) (*LogoutOutput, error) {
	req, out := c.LogoutRequest(input)
	return out, req.Send()
}

// LogoutWithContext is the same as Logout with the addition of
// the ability to pass a context and additional request options.
//
// See Logout for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SSO) LogoutWithContext(ctx aws.Context, input *LogoutInput, opts ...request.Option) (*LogoutOutput, error) {
	req, out := c.LogoutRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// Provides information about your AWS account.
type AccountInfo struct {
	_ struct{} `type:"structure"`

	// The identifier of the AWS account that is assigned to the user.
	AccountId *string `locationName:"accountId" type:"string"`

	// The display name of the AWS account that is assigned to the user.
	AccountName *string `locationName:"accountName" type:"string"`

	// The email address of the AWS account that is assigned to the user.
	EmailAddress *string `locationName:"emailAddress" min:"1" type:"string"`
}

// String returns the string representation
func (s AccountInfo) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AccountInfo) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *AccountInfo) SetAccountId(v string) *AccountInfo {
	s.AccountId = &v
	return s
}

// SetAccountName sets the AccountName field's value.
func (s *AccountInfo) SetAccountName(v string) *AccountInfo {
	s.AccountName = &v
	return s
}

// SetEmailAddress sets the EmailAddress field's value.
func (s *AccountInfo) SetEmailAddress(v string) *AccountInfo {
	s.EmailAddress = &v
	return s
}

type GetRoleCredentialsInput struct {
	_ struct{} `type:"structure"`

	// The token issued by the CreateToken API call. For more information, see CreateToken
	// (https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/API_CreateToken.html)
	// in the AWS SSO OIDC API Reference Guide.
	//
	// AccessToken is a required field
	AccessToken *string `location:"header" locationName:"x-amz-sso_bearer_token" type:"string" required:"true" sensitive:"true"`

	// The identifier for the AWS account that is assigned to the user.
	//
	// AccountId is a required field
	AccountId *string `location:"querystring" locationName:"account_id" type:"string" required:"true"`

	// The friendly name of the role that is assigned to the user.
	//
	// RoleName is a required field
	RoleName *string `location:"querystring" locationName:"role_name" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRoleCredentialsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRoleCredentialsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRoleCredentialsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetRoleCredentialsInput"}
	if s.AccessToken == nil {
		invalidParams.Add(request.NewErrParamRequired("AccessToken"))
	}
	if s.AccountId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountId"))
	}
	if s.RoleName == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccessToken sets the AccessToken field's value.
func (s *GetRoleCredentialsInput) SetAccessToken(v string) *GetRoleCredentialsInput {
	s.AccessToken = &v
	return s
}

// SetAccountId sets the AccountId field's value.
func (s *GetRoleCredentialsInput) SetAccountId(v string) *GetRoleCredentialsInput {
	s.AccountId = &v
	return s
}

// SetRoleName sets the RoleName field's value.
func (s *GetRoleCredentialsInput) SetRoleName(v string) *GetRoleCredentialsInput {
	s.RoleName = &v
	return s
}

type GetRoleCredentialsOutput struct {
	_ struct{} `type:"structure"`

	// The credentials for the role that is assigned to the user.
	RoleCredentials *RoleCredentials `locationName:"roleCredentials" type:"structure"`
}

// String returns the string representation
func (s GetRoleCredentialsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetRoleCredentialsOutput) GoString() string {
	return s.String()
}

// SetRoleCredentials sets the RoleCredentials field's value.
func (s *GetRoleCredentialsOutput) SetRoleCredentials(v *RoleCredentials) *GetRoleCredentialsOutput {
	s.RoleCredentials = v
	return s
}

// Indicates that a problem occurred with the input to the request. For example,
// a required parameter might be missing or out of range.
type InvalidRequestException struct {
	_            struct{} `type:"structure"`
	respMetadata protocol.ResponseMetadata

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s InvalidRequestException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s InvalidRequestException) GoString() string {
	return s.String()
}

func newErrorInvalidRequestException(v protocol.ResponseMetadata) error {
	return &InvalidRequestException{
		respMetadata: v,
	}
}

// Code returns the exception type name.
func (s InvalidRequestException) Code() string {
	return "InvalidRequestException"
}

// Message returns the exception's message.
func (s InvalidRequestException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s InvalidRequestException) OrigErr() error {
	return nil
}

func (s InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s InvalidRequestException) StatusCode() int {
	return s.respMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s InvalidRequestException) RequestID() string {
	return s.respMetadata.RequestID
}

type ListAccountRolesInput struct {
	_ struct{} `type:"structure"`

	// The token issued by the CreateToken API call. For more information, see CreateToken
	// (https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/API_CreateToken.html)
	// in the AWS SSO OIDC API Reference Guide.
	//
	// AccessToken is a required field
	AccessToken *string `location:"header" locationName:"x-amz-sso_bearer_token" type:"string" required:"true" sensitive:"true"`

	// The identifier for the AWS account that is assigned to the user.
	//
	// AccountId is a required field
	AccountId *string `location:"querystring" locationName:"account_id" type:"string" required:"true"`

	// The number of items that clients can request per page.
	MaxResults *int64 `location:"querystring" locationName:"max_result" min:"1" type:"integer"`

	// The page token from the previous response output when you request subsequent
	// pages.
	NextToken *string `location:"querystring" locationName:"next_token" type:"string"`
}

// String returns the string representation
func (s ListAccountRolesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAccountRolesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAccountRolesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAccountRolesInput"}
	if s.AccessToken == nil {
		invalidParams.Add(request.NewErrParamRequired("AccessToken"))
	}
	if s.AccountId == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountId"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccessToken sets the AccessToken field's value.
func (s *ListAccountRolesInput) SetAccessToken(v string) *ListAccountRolesInput {
	s.AccessToken = &v
	return s
}

// SetAccountId sets the AccountId field's value.
func (s *ListAccountRolesInput) SetAccountId(v string) *ListAccountRolesInput {
	s.AccountId = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListAccountRolesInput) SetMaxResults(v int64) *ListAccountRolesInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListAccountRolesInput) SetNextToken(v string) *ListAccountRolesInput {
	s.NextToken = &v
	return s
}

type ListAccountRolesOutput struct {
	_ struct{} `type:"structure"`

	// The page token client that is used to retrieve the list of accounts.
	NextToken *string `locationName:"nextToken" type:"string"`

	// A paginated response with the list of roles and the next token if more results
	// are available.
	RoleList []*RoleInfo `locationName:"roleList" type:"list"`
}

// String returns the string representation
func (s ListAccountRolesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAccountRolesOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *ListAccountRolesOutput) SetNextToken(v string) *ListAccountRolesOutput {
	s.NextToken = &v
	return s
}

// SetRoleList sets the RoleList field's value.
func (s *ListAccountRolesOutput) SetRoleList(v []*RoleInfo) *ListAccountRolesOutput {
	s.RoleList = v
	return s
}

type ListAccountsInput struct {
	_ struct{} `type:"structure"`

	// The token issued by the CreateToken API call. For more information, see CreateToken
	// (https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/API_CreateToken.html)
	// in the AWS SSO OIDC API Reference Guide.
	//
	// AccessToken is a required field
	AccessToken *string `location:"header" locationName:"x-amz-sso_bearer_token" type:"string" required:"true" sensitive:"true"`

	// This is the number of items clients can request per page.
	MaxResults *int64 `location:"querystring" locationName:"max_result" min:"1" type:"integer"`

	// (Optional) When requesting subsequent pages, this is the page token from
	// the previous response output.
	NextToken *string `location:"querystring" locationName:"next_token" type:"string"`
}

// String returns the string representation
func (s ListAccountsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAccountsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAccountsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAccountsInput"}
	if s.AccessToken == nil {
		invalidParams.Add(request.NewErrParamRequired("AccessToken"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccessToken sets the AccessToken field's value.
func (s *ListAccountsInput) SetAccessToken(v string) *ListAccountsInput {
	s.AccessToken = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListAccountsInput) SetMaxResults(v int64) *ListAccountsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListAccountsInput) SetNextToken(v string) *ListAccountsInput {
	s.NextToken = &v
	return s
}

type ListAccountsOutput struct {
	_ struct{} `type:"structure"`

	// A paginated response with the list of account information and the next token
	// if more results are available.
	AccountList []*AccountInfo `locationName:"accountList" type:"list"`

	// The page token client that is used to retrieve the list of accounts.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListAccountsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAccountsOutput) GoString() string {
	return s.String()
}

// SetAccountList sets the AccountList field's value.
func (s *ListAccountsOutput) SetAccountList(v []*AccountInfo) *ListAccountsOutput {
	s.AccountList = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListAccountsOutput) SetNextToken(v string) *ListAccountsOutput {
	s.NextToken = &v
	return s
}

type LogoutInput struct {
	_ struct{} `type:"structure"`

	// The token issued by the CreateToken API call. For more information, see CreateToken
	// (https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/API_CreateToken.html)
	// in the AWS SSO OIDC API Reference Guide.
	//
	// AccessToken is a required field
	AccessToken *string `location:"header" locationName:"x-amz-sso_bearer_token" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s LogoutInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s LogoutInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *LogoutInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "LogoutInput"}
	if s.AccessToken == nil {
		invalidParams.Add(request.NewErrParamRequired("AccessToken"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccessToken sets the AccessToken field's value.
func (s *LogoutInput) SetAccessToken(v string) *LogoutInput {
	s.AccessToken = &v
	return s
}

type LogoutOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s LogoutOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s LogoutOutput) GoString() string {
	return s.String()
}

// The specified resource doesn't exist.
type ResourceNotFoundException struct {
	_            struct{} `type:"structure"`
	respMetadata protocol.ResponseMetadata

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s ResourceNotFoundException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceNotFoundException) GoString() string {
	return s.String()
}

func newErrorResourceNotFoundException(v protocol.ResponseMetadata) error {
	return &ResourceNotFoundException{
		respMetadata: v,
	}
}

// Code returns the exception type name.
func (s ResourceNotFoundException) Code() string {
	return "ResourceNotFoundException"
}

// Message returns the exception's message.
func (s ResourceNotFoundException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s ResourceNotFoundException) OrigErr() error {
	return nil
}

func (s ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s ResourceNotFoundException) StatusCode() int {
	return s.respMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s ResourceNotFoundException) RequestID() string {
	return s.respMetadata.RequestID
}

// Provides information about the role credentials that are assigned to the
// user.
type RoleCredentials struct {
	_ struct{} `type:"structure"`

	// The identifier used for the temporary security credentials. For more information,
	// see Using Temporary Security Credentials to Request Access to AWS Resources
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html)
	// in the AWS IAM User Guide.
	AccessKeyId *string `locationName:"accessKeyId" type:"string"`

	// The date on which temporary security credentials expire.
	Expiration *int64 `locationName:"expiration" type:"long"`

	// The key that is used to sign the request. For more information, see Using
	// Temporary Security Credentials to Request Access to AWS Resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html)
	// in the AWS IAM User Guide.
	SecretAccessKey *string `locationName:"secretAccessKey" type:"string" sensitive:"true"`

	// The token used for temporary credentials. For more information, see Using
	// Temporary Security Credentials to Request Access to AWS Resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_use-resources.html)
	// in the AWS IAM User Guide.
	SessionToken *string `locationName:"sessionToken" type:"string" sensitive:"true"`
}

// String returns the string representation
func (s RoleCredentials) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RoleCredentials) GoString() string {
	return s.String()
}

// SetAccessKeyId sets the AccessKeyId field's value.
func (s *RoleCredentials) SetAccessKeyId(v string) *RoleCredentials {
	s.AccessKeyId = &v
	return s
}

// SetExpiration sets the Expiration field's value.
func (s *RoleCredentials) SetExpiration(v int64) *RoleCredentials {
	s.Expiration = &v
	return s
}

// SetSecretAccessKey sets the SecretAccessKey field's value.
func (s *RoleCredentials) SetSecretAccessKey(v string) *RoleCredentials {
	s.SecretAccessKey = &v
	return s
}

// SetSessionToken sets the SessionToken field's value.
func (s *RoleCredentials) SetSessionToken(v string) *RoleCredentials {
	s.SessionToken = &v
	return s
}

// Provides information about the role that is assigned to the user.
type RoleInfo struct {
	_ struct{} `type:"structure"`

	// The identifier of the AWS account assigned to the user.
	AccountId *string `locationName:"accountId" type:"string"`

	// The friendly name of the role that is assigned to the user.
	RoleName *string `locationName:"roleName" type:"string"`
}

// String returns the string representation
func (s RoleInfo) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RoleInfo) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *RoleInfo) SetAccountId(v string) *RoleInfo {
	s.AccountId = &v
	return s
}

// SetRoleName sets the RoleName field's value.
func (s *RoleInfo) SetRoleName(v string) *RoleInfo {
	s.RoleName = &v
	return s
}

// Indicates that the request is being made too frequently and is more than
// what the server can handle.
type TooManyRequestsException struct {
	_            struct{} `type:"structure"`
	respMetadata protocol.ResponseMetadata

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s TooManyRequestsException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TooManyRequestsException) GoString() string {
	return s.String()
}

func newErrorTooManyRequestsException(v protocol.ResponseMetadata) error {
	return &TooManyRequestsException{
		respMetadata: v,
	}
}

// Code returns the exception type name.
func (s TooManyRequestsException) Code() string {
	return "TooManyRequestsException"
}

// Message returns the exception's message.
func (s TooManyRequestsException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s TooManyRequestsException) OrigErr() error {
	return nil
}

func (s TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s TooManyRequestsException) StatusCode() int {
	return s.respMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s TooManyRequestsException) RequestID() string {
	return s.respMetadata.RequestID
}

// Indicates that the request is not authorized. This can happen due to an invalid
// access token in the request.
type UnauthorizedException struct {
	_            struct{} `type:"structure"`
	respMetadata protocol.ResponseMetadata

	Message_ *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s UnauthorizedException) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UnauthorizedException) GoString() string {
	return s.String()
}

func newErrorUnauthorizedException(v protocol.ResponseMetadata) error {
	return &UnauthorizedException{
		respMetadata: v,
	}
}

// Code returns the exception type name.
func (s UnauthorizedException) Code() string {
	return "UnauthorizedException"
}

// Message returns the exception's message.
func (s UnauthorizedException) Message() string {
	if s.Message_ != nil {
		return *s.Message_
	}
	return ""
}

// OrigErr always returns nil, satisfies awserr.Error interface.
func (s UnauthorizedException) OrigErr() error {
	return nil
}

func (s UnauthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", s.Code(), s.Message())
}

// Status code returns the HTTP status code for the request's response error.
func (s UnauthorizedException) StatusCode() int {
	return s.respMetadata.StatusCode
}

// RequestID returns the service's response RequestID for request.
func (s UnauthorizedException) RequestID() string {
	return s.respMetadata.RequestID
}
