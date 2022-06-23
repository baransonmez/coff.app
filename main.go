package main

import (
	"fmt"
	"github.com/baransonmez/coff.app/business/core/coffee"
	"github.com/baransonmez/coff.app/business/core/coffee/db"
	"github.com/google/uuid"
)

func main() {
	fmt.Println("mod file generated")
	coffStore := db.NewStore()
	service := coffee.NewService(coffStore)
	bean, _ := service.GetCoffeeBean(nil, uuid.New())
	fmt.Println(bean)
}
