package request

type SearchFriend struct {
	UserId   string `json:"userId" form:"userId"`
	UserName string `json:"userName" form:"userName"`
}
type AddFriend struct {
	UserId   uint `json:"userId"`
	FriendId uint `json:"friendId"`
}
