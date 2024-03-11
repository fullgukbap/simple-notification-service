package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"notification-service/db"
	"notification-service/dto"
	"notification-service/ent/friendship"
	"notification-service/ent/friendshipstatus"
	"notification-service/ent/notificationobject"
	"notification-service/ent/user"
	"notification-service/notification"
	"strconv"
)

func UserFriendRelative(w http.ResponseWriter, r *http.Request) {
	body := new(dto.UserFriendRelativeRequest)
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b := db.I.Friendship.Query().Where(
		friendship.And(
			friendship.HasReceiverWith(user.ID(body.ReceiverUserID)),
			friendship.HasSenderWith(user.ID(body.SenderUserID)),
			friendship.HasFriendshipStatusWith(friendshipstatus.ID(notification.AddingFriendType)),
		),
	).ExistX(context.Background())
	if b {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "이미 신청하였습니다.")
		return
	}

	friendship := db.I.Friendship.Create().
		SetFriendshipStatusID(notification.AddingFriendType).
		SetReceiverID(body.ReceiverUserID).
		SetSenderID(body.SenderUserID).
		SaveX(context.Background())

	notification.Data{
		EntityTypeID: notification.AddingFriendType,
		EntityID:     friendship.ID,
		ActorID:      body.SenderUserID,
		NotifierID:   body.ReceiverUserID,
	}.AutoMigrateX()

	fmt.Fprint(w, http.StatusCreated)
}

func UserNotifications(w http.ResponseWriter, r *http.Request) {
	uid, _ := strconv.Atoi(r.PathValue("uid"))
	json.NewEncoder(w).Encode(notification.AutoRefineX(uid))
}

func UserList(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(db.I.User.Query().AllX(context.Background()))
}

func UserFriendAccept(w http.ResponseWriter, r *http.Request) {
	receiverUserID, _ := strconv.Atoi(r.PathValue("uid"))

	body := new(dto.UserFriendAcceptRequest)
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		panic(err)
	}

	f, err := db.I.Friendship.Query().Where(
		friendship.And(
			friendship.HasReceiverWith(user.ID(receiverUserID)),
			friendship.HasSenderWith(user.ID(body.SenderUserID)),
			friendship.HasFriendshipStatusWith(friendshipstatus.ID(notification.AddingFriendType)),
		),
	).Only(context.Background())

	if err == nil {
		db.I.Friendship.
			UpdateOneID(f.ID).
			SetFriendshipStatusID(notification.FriendType).
			SaveX(context.Background())
		fmt.Fprint(w, "성곡적으로 친추를 수락하였습니다.")

		u := db.I.Friendship.QueryReceiver(f).OnlyX(context.Background())
		noti := db.I.User.QueryNotifications(u).WithNotificationObject().OnlyX(context.Background())
		notiOb := db.I.NotificationObject.Query().Where(notificationobject.ID(noti.Edges.NotificationObject.ID)).OnlyX(context.Background())
		notiOb.Update().SetEntityTypeID(notification.FriendType).SaveX(context.Background())
	} else {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "잘못된 요청입니다. %s", err.Error())
	}

}
