package main

func ensureAdminExists() {
	if !containsLogin(DefaultAdminLogin) {
		users[DefaultAdminLogin] = &User{Password: DefaultAdminHash}
	}
}
