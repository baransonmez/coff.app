package main

import (
	"fmt"
	"github.com/baransonmez/coff.app/business/core/coffee"
	coffeeData "github.com/baransonmez/coff.app/business/core/coffee/data"
	"github.com/baransonmez/coff.app/business/core/user"
	userData "github.com/baransonmez/coff.app/business/core/user/data"
	"time"
)

func main() {
	fmt.Println("mod file generated")
	coffStore := coffeeData.NewInMem()
	service := coffee.NewService(coffStore)
	beanId, _ := service.CreateCoffeeBean(nil, coffee.NewCoffeeBean{
		Name:      "Yirgaciffe",
		Roaster:   "Montag",
		Origin:    "Etiopia",
		Price:     23,
		RoastDate: time.Now().AddDate(2, 3, 4),
	})
	bean, _ := service.GetCoffeeBean(nil, beanId)
	fmt.Println(bean)

	userStore := userData.NewInMem()
	userService := user.NewService(userStore)
	userId, _ := userService.CreateNewUser(nil, user.NewUser{
		Name: "Baran",
	})
	newUser, _ := userService.GetUser(nil, userId)
	fmt.Println(newUser)

}
