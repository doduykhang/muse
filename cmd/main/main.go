package main

import (
	"fmt"
	"github.com/doduykhang/muse/pkg/config"
)

func main() {
	env := config.GetEnv()
	fmt.Printf("%+v \n", env)
}
