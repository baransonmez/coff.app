package main

import (
	"fmt"
	"github.com/baransonmez/coff.app/business/core/coffee"
	"github.com/baransonmez/coff.app/business/core/coffee/db"
)

func main() {
	fmt.Println("mod file generated")
	coffStore := db.NewStore()
	_ = coffee.NewService(coffStore)

}
