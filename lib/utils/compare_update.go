package utils

import (
	"github.com/artnoi43/todong/data/model"
	"github.com/artnoi43/todong/internal"
	"github.com/artnoi43/todong/lib/enums"
)

func UpdatedTodo(
	uuid string, // targetTodo's UUID - it is from the path param, not req body
	imgStrReq string, // New imageStr from req,
	updatesReq *internal.TodoReqBody, // This is updated fields from the
	targetTodo *model.Todo, // This is existing to-do
	u *model.Todo, // This one is going to be written to data store
) {
	compareVal(targetTodo.UUID, uuid, &u.UUID)
	compareVal(targetTodo.UserUUID, "", &u.UserUUID)
	compareVal(targetTodo.Title, updatesReq.Title, &u.Title)
	compareVal(targetTodo.Description, updatesReq.Description, &u.Description)
	compareVal(targetTodo.TodoDate, updatesReq.TodoDate, &u.TodoDate)
	compareVal(targetTodo.Image, imgStrReq, &u.Image)
	compareStatus(targetTodo.Status, enums.Status(updatesReq.Status), &u.Status)
}

func compareVal(old, new string, target *string) {
	var nullString string
	if new == nullString {
		*target = old
		return
	}
	*target = new
}

func compareStatus(old, new enums.Status, target *enums.Status) {
	if new.IsValid() {
		*target = new
	}
	*target = old
}
