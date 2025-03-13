package main

import "fmt"

func main() {
	adminUser := User{"Alice", "admin"}
	customerUser := User{"Charlie", "customer"}

	fmt.Println(adminUser.Username, "can manage users: ", HasPermission(adminUser, "manage_users"))
	fmt.Println(customerUser.Username, "Can't manage users", HasPermission(customerUser, "manage_users"))
}
