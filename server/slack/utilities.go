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
