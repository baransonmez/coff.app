package main

import (
	"fmt"
	"github.com/baransonmez/coff.app/business/core/coffee"
	"github.com/baransonmez/coff.app/business/core/coffee/data"
	"time"
)

func main() {
	fmt.Println("mod file generated")
	coffStore := data.NewInMem()
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
}
