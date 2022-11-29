package main

import (
	"assumerole/internal/assumeRole"
	"fmt"
)

type Role struct {
	Roles *assumeRole.AssumeRole
}

func Run() error {
	fmt.Println("Connecting to Services using STS Token")
	assumeRoles := assumeRole.NewAssumeRole()
	Role := Role{
		Roles: assumeRoles,
	}
	err := Role.Roles.Connect()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Unable to Connect to the Application!")
	}
}
