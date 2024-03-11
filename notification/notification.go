package notification

import (
	"context"
	"fmt"
	"notification-service/db"
	"notification-service/ent/user"
)

const (
	FriendType       = 1
	AddingFriendType = 2
)

type Data struct {
	EntityTypeID int
	EntityID     int
	ActorID      int
	NotifierID   int
}

func (n Data) AutoMigrateX() {

	//TODO: 같은 요청인지 확인하기

	notiObject := db.I.NotificationObject.Create().
		SetEntityTypeID(n.EntityTypeID).
		SetEntityID(n.EntityID).
		SaveX(context.Background())

	db.I.NotificationChange.Create().
		SetActorID(n.ActorID).
		SetNotificationObjectID(notiObject.ID).
		SaveX(context.Background())

	db.I.Notification.Create().
		SetNotifierID(n.NotifierID).
		SetNotificationObjectID(notiObject.ID).
		SaveX(context.Background())
}

func AutoRefineX(userID int) (notifications []string) {
	user := db.I.User.Query().
		Where(user.ID(userID)).
		WithNotifications().
		OnlyX(context.Background())

	for _, noti := range user.Edges.Notifications {
		notiOb := db.I.Notification.QueryNotificationObject(noti).OnlyX(context.Background())
		notiCh := db.I.NotificationObject.QueryNotificationChanges(notiOb).OnlyX(context.Background())
		actor := db.I.NotificationChange.QueryActor(notiCh).OnlyX(context.Background())
		entityType := db.I.NotificationObject.QueryEntityType(notiOb).OnlyX(context.Background())
		notifications = append(notifications, fmt.Sprintf("%s가 %s에게 %s합니다.", actor.Username, "너", entityType.NotificationDescription))
	}

	return notifications
}
