package model

type Video struct {
	ID            uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	FromWho       uint   `json:"-"` //表示视频所属的人
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount uint   `json:"favorite_count"`
	CommentCount  uint   `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}
