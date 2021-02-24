package main

import (
	"electronic-school-diary/env"
)

func main() {

	env := new(env.Env)
	env.InitDB()

}
