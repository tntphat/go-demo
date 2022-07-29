package serializers

// UserLoginResp : struct
type UserLoginResp struct {
	Token   string `json:"token"`
	Expired int64  `json:"expired"`
}
