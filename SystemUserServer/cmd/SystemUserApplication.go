package main

import (
	"SystemUserServer/internal/mapper"
	"SystemUserServer/internal/router"
)

func main() {
	mapper.DataBaseMapperInit()
	router.RouterInit()
}
