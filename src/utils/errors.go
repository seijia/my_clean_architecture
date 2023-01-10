package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UnmarshalErrorResponse(data []byte) (ErrorResponse, error) {
	var r ErrorResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ErrorResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ErrorResponse struct {
	Errors []ErrorResp `json:"errors"`
}

type ErrorResp struct {
	Message          string `json:"message"`
	ReturnDetailCode string `json:"returnDetailCode"`
	StatusCode       int    `json:"-"`
	err              error  `json:"-"`
}

func (er *ErrorResp) Error() string {
	return fmt.Sprintf("%+v", er.err)
}

func (er *ErrorResp) ToResponse() *ErrorResponse {
	return &ErrorResponse{
		Errors: []ErrorResp{*er},
	}
}

func NewErrorResp(message, returnDetailCode string, statusCode int, err error) *ErrorResp {

	return &ErrorResp{
		Message:          message,
		ReturnDetailCode: returnDetailCode,
		StatusCode:       statusCode,
		err:              err,
	}
}

func NewBadRequestError(err error) *ErrorResp {
	return NewErrorResp("コンテンツが存在しません", "bad_request", http.StatusBadRequest, err)
}

func NewInvalidAccessToken(err error) *ErrorResp {
	return NewErrorResp("アクセストークンの認証に失敗しました", "invalid_access_token", http.StatusUnauthorized, err)
}

func NewExpiredAccessToken(err error) *ErrorResp {
	return NewErrorResp("アクセストークンの有効期限が切れました", "expired_access_token", http.StatusUnauthorized, err)
}

func NewInternalServerError(err error) *ErrorResp {
	return NewErrorResp("予期せぬサーバーエラーが発生しました", "internal_server_error", http.StatusInternalServerError, err)
}

func NewServiceUnavailableErrorOutside(err error) *ErrorResp {
	return NewErrorResp("メンテナンス中です、しばらくしてから再度お試しください", "service_unavailable_error_outside", http.StatusServiceUnavailable, err)
}
