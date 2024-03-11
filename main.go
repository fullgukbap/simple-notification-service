package main

import (
	"net/http"
	_ "notification-service/ent/runtime"
	"notification-service/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// if err := db.I.Schema.Create(context.Background()); err != nil {
	// 	panic(err)
	// }

	// db.I.FriendshipStatus.CreateBulk(
	// 	db.I.FriendshipStatus.Create().
	// 		SetDescription("friend"),

	// 	db.I.FriendshipStatus.Create().
	// 		SetDescription("request to add a friend"),
	// ).SaveX(context.Background())

	// db.I.EntityType.CreateBulk(
	// 	db.I.EntityType.Create().
	// 		SetEntityName("friendship").
	// 		SetNotificationDescription("fried"),

	// 	db.I.EntityType.Create().
	// 		SetEntityName("friendship").
	// 		SetNotificationDescription("request to add a friend"),
	// ).SaveX(context.Background())

	// db.I.User.Create().
	// 	SetUsername("junbumhan").
	// 	SaveX(context.Background())
	// db.I.User.Create().
	// 	SetUsername("jahyun").
	// 	SaveX(context.Background())

	http.HandleFunc("POST /api/user/friend/relative", handler.UserFriendRelative)
	http.HandleFunc("GET /api/{uid}/notifications", handler.UserNotifications)
	http.HandleFunc("POST /api/{uid}/friend/accept", handler.UserFriendAccept)
	http.HandleFunc("GET /api/users", handler.UserList)

	http.ListenAndServe(":9190", nil)
}

// 내가 다른 유저들에게 친추
// [POST] /api/user/friend/relative

// 현재 나에게 날라온 알림(친추)
// [GET] /api/:uid/notifications

// 친추 수락
// [POST] /api/:uid/friend/accept

// 현재 존재하는 유저들
// [GET] /api/users

/*
	friendship status
	1. friend
	2. request to add a friend
*/
