package response

import "TikTok/internal/model"

type VideoListResponse struct {
	CommResponse
	VideoList []model.Video `json:"video_list"`
}
