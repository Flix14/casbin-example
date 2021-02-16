package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin"
)

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatalf("Error al crear Casbin Enforcer: %v", err)
	}

	felixRoles, err := e.GetRolesForUser("felix")
	if err != nil {
		log.Fatalf("Error al obtener roles de felix: %v", err)
	}
	felixPerms, err := e.GetImplicitPermissionsForUser("felix")
	if err != nil {
		log.Fatalf("No se pueden obtener permisos para felix: %v", err)
	}
	fmt.Printf(
		"Roles de felix: %v. Permisos de felix: %v.\n",
		felixRoles,
		felixPerms,
	)

	maribelRoles, err := e.GetRolesForUser("maribel")
	if err != nil {
		log.Fatalf("Error al obtener roles de maribel: %v", err)
	}
	maribelPerms, err := e.GetImplicitPermissionsForUser("maribel")
	if err != nil {
		log.Fatalf("No se pueden obtener permisos para maribel: %v", err)
	}
	fmt.Printf(
		"Roles de maribel: %v. Permisos de maribel: %v.\n",
		maribelRoles,
		maribelPerms,
	)
}
