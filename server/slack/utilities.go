package slack

import "fmt"

func getChannelInfo(channel string) (info Channel) {
	c, err := api.GetChannelInfo(channel)
	if err != nil {
		fmt.Println("Error retrieving channel information")
	}
	info = Channel{
		Name: c.Name,
		ID:   c.ID,
	}
	return
}

func getUserInfo(user string) (info User) {
	u, err := api.GetUserInfo(user)
	if err != nil {
		fmt.Println("Error retrieving channel information")
	}
	info = User{
		Name: u.Name,
		ID:   u.ID,
	}
	return
}

func ListChannels() (channels []map[string]string) {
	c, err := api.GetChannels(true)
	if err != nil {
		fmt.Println("Error retrieving channel list")
	}
	for _, channel := range c {
		channels = append(channels, map[string]string{
			"ChannelName": channel.Name,
			"ID":          channel.ID,
		})
	}
	return channels
}

func ListUsers() (users []map[string]string) {
	u, err := api.GetUsers()
	if err != nil {
		fmt.Println("Error retrieving user list")
	}
	for _, user := range u {
		users = append(users, map[string]string{
			"UserName": user.Name,
			"ID":       user.ID,
		})
	}
	return users
}

func ListGroups() (groups []map[string]string) {
	g, err := api.GetGroups(true)
	if err != nil {
		fmt.Println("Error retrieving groups list")
	}
	for _, group := range g {
		groups = append(groups, map[string]string{
			"GroupName": group.Name,
			"ID":        group.ID,
		})
	}
	return groups
}
