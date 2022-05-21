package response

import user "TikTok/internal/model"

type Response struct {
	StatusCode uint   `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     uint64 `json:"user_id"`
	Token      string `json:"token"`
}

const FailureCode uint = 1
const SuccessCode = 0
const FailureMsgToken = "token验证失败"

type CommResponse struct {
	StatusCode uint   `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type UserInfo struct {
	StatusCode uint         `json:"status_code"`
	StatusMsg  string       `json:"status_msg"`
	User       user.Account `json:"user"`
}
