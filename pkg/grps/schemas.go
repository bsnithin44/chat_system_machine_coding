package grps

import "github.com/google/uuid"

type Group struct {
	id     string
	name   string
	admin  string
	Users  []string
	chatID string
}

func NewGroup(name string, admin string) Group {
	return Group{
		id:    "random id",
		name:  name,
		admin: admin,
		Users: []string{admin},
	}
}
func (g *Group) GetID() string {
	return g.id
}

func (g *Group) addUser(userID string) {
	g.Users = append(g.Users, userID)
}

func (g *Group) setChatID() {
	g.chatID = uuid.New().String()
}

func (g *Group) getChatID() string {
	return g.chatID
}
