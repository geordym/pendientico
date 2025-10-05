package ports_out

import "context"

type AuthenticationProviderCommunication interface {
	SaveUser(email string, password string) (string, error)
	GetUserAuthProviderIdLogged(ctx context.Context) (string, error)
}
