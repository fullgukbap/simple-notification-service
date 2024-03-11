package dto

type UserFriendRelativeRequest struct {
	SenderUserID   int `json:"sender_user_id"`
	ReceiverUserID int `json:"receiver_user_id"`
}

type UserFriendAcceptRequest struct {
	SenderUserID int `json:"sender_user_id"`
}
