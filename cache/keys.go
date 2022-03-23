package cache

import (
	"fmt"
	"strconv"
)

const (
	ClicksVideoList = "video:clicks:list"
)

func VideoClicksKey(id int) string {
	return fmt.Sprintf("video:clicks:%s", strconv.Itoa(id))
}

func VideoFavoriteKey(id int) string {
	return fmt.Sprintf("video:favorite:%s", strconv.Itoa(id))
}

func VideoLikeKey(id int) string {
	return fmt.Sprintf("video:like:%s", strconv.Itoa(id))
}

func CodeKey(email string) string {
	return fmt.Sprintf("code:%s", email)
}
