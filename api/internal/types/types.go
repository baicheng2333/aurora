// Code generated by goctl. DO NOT EDIT.
package types

type AuthenticateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthenticatedResponse struct {
	Token string `json:"token"`
}
