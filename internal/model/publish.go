package model

type Video struct {
	ID            uint64  `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Author        Account `json:"author" gorm:"-"`
	FromWho       uint64  `json:"-"` //表示视频所属的人
	PlayUrl       string  `json:"play_url"`
	CoverUrl      string  `json:"cover_url"`
	FavoriteCount uint64  `json:"favorite_count"`
	CommentCount  uint64  `json:"comment_count"`
	IsFavorite    bool    `json:"is_favorite"`
	Title         string  `json:"title"`
}
