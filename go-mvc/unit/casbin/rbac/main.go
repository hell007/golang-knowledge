package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin"
)

func main() {

	e := casbin.NewEnforcer("rbac.conf", "rbac.csv")

	fmt.Printf("RBAC test start\n") // output for debug

	// user
	// Enforce(uid, ctx.Path(), ctx.Method(), ".*")
	if e.Enforce("2", "/user/list/*", "GET", ".*") {
		log.Println("2 can read project")
	} else {
		log.Fatal("ERROR: 2 can not read project")
	}

	if e.Enforce("1", "/user/save", "POST", ".*") {
		log.Println("1 can read project")
	} else {
		log.Fatal("ERROR: 1 can not read project")
	}

}
