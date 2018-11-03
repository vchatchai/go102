package user

import (
	"fmt"
	"os/user"
)

func User() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println("The user home directory: ", user.HomeDir)

}
