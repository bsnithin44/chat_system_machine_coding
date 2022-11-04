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

	nithin := users.NewUser("Nithin")
	alpit := users.NewUser("alpit")
	santosh := users.NewUser("santosh")

	uService.CreateUser(nithin)
	uService.CreateUser(alpit)
	uService.CreateUser(santosh)

	uService.SendFriendRequest(nithin.GetID(), alpit.GetID())
	uService.ApproveFriendReq(alpit.GetID(), nithin.GetID())
	uService.SendFriendRequest(nithin.GetID(), santosh.GetID())
	uService.RejectFriendReq(santosh.GetID(), nithin.GetID())

	msg1 := msgs.NewMessage("hello alpit")
	msg2 := msgs.NewMessage("you write bad code")

	uService.Send1to1Message(nithin.GetID(), alpit.GetID(), msg1)
	uService.Send1to1Message(nithin.GetID(), alpit.GetID(), msg2)

	msg3 := msgs.NewMessage("so do you")

	uService.Send1to1Message(alpit.GetID(), nithin.GetID(), msg3)

	grp1 := grps.NewGroup("bellandur", nithin.GetID())
	gService.CreateGroup(grp1)

	gService.AddUserToGroup(alpit.GetID(), grp1.GetID())

	msg4 := msgs.NewMessage("please dont get added")
	gService.SendGroupMessage(grp1.GetID(), santosh.GetID(), msg4)

	msg5 := msgs.NewMessage("please  get added nithin")
	gService.SendGroupMessage(grp1.GetID(), nithin.GetID(), msg5)

	msg6 := msgs.NewMessage("please  get added alput")
	gService.SendGroupMessage(grp1.GetID(), alpit.GetID(), msg6)

	msgs11 := gService.GetGroupMessages(grp1.GetID())
	for _, m := range msgs11 {
		fmt.Println(m.Text, m.Sender)

	}
}
