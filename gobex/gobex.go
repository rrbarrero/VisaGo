package gobex

import (
	"log"

	auth "github.com/korylprince/go-ad-auth/v3"
)

func LdapInit() Gobex {
	config := auth.Config{
		Server:   "server",
		Port:     389,
		BaseDN:   "ou=ORGANIZACION,dc=server,dc=es",
		Security: auth.SecurityInsecureStartTLS,
	}
	ad := Gobex{config: config}
	return ad
}

type Gobex struct {
	config auth.Config
}

// AuthUser validate username and pass against AD server
func (gobex *Gobex) AuthUser(username, password string) bool {

	status, err := auth.Authenticate(&gobex.config, username, password)
	if err != nil {
		log.Fatal(err)
	}
	if !status {
		log.Printf("Failed to login with %s and %s", username, password)
		return false
	}
	log.Printf("Login success for %s.", username)
	return true
}

func (gobex *Gobex) ChangeUserPassword(username, password, newPassword string) bool {
	if err := auth.UpdatePassword(&gobex.config, username, password, newPassword); err != nil {
		log.Printf("Failed to change password with %s, %s and %s", username, password, newPassword)
		return false
	}
	return true
}
