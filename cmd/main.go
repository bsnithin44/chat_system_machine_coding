package main

import (
	"chatss/pkg/grps"
	"chatss/pkg/msgs"
	"chatss/pkg/users"
	"fmt"
)

func main() {

	mService := msgs.NewMessageService()
	gService := grps.NewGroupService(mService)

	uService := users.NewUserService(mService)

	nithin := uService.CreateUser("Nithin")
	alpit := uService.CreateUser("alpit")
	santosh := uService.CreateUser("santosh")

	uService.SendFriendRequest(nithin, alpit)
	uService.ApproveFriendReq(alpit, nithin)
	uService.SendFriendRequest(nithin, santosh)
	uService.RejectFriendReq(santosh, nithin)

	msg1 := msgs.NewMessage("hello alpit")
	msg2 := msgs.NewMessage("you write bad code")

	uService.Send1to1Message(nithin, alpit, msg1)
	uService.Send1to1Message(nithin, alpit, msg2)

	msg3 := msgs.NewMessage("so do you")

	uService.Send1to1Message(alpit, nithin, msg3)

	grp1 := gService.CreateGroup("bellandur", nithin)

	gService.AddUserToGroup(alpit, grp1)

	msg4 := msgs.NewMessage("please dont get added")
	gService.SendGroupMessage(grp1, santosh, msg4)

	msg5 := msgs.NewMessage("please  get added nithin")
	gService.SendGroupMessage(grp1, nithin, msg5)

	msg6 := msgs.NewMessage("please  get added alput")
	gService.SendGroupMessage(grp1, alpit, msg6)

	msgs11 := gService.GetGroupMessages(grp1)
	for _, m := range msgs11 {
		fmt.Println(m.Text, m.Sender)

	}
}
