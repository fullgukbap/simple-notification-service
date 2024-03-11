package db

import "notification-service/ent"

func init() {
	client, err := ent.Open("mysql", "root:qwe123@tcp(localhost:3306)/test?parseTime=True")
	if err != nil {
		panic(err)
	}
	I = client
}

var I *ent.Client
