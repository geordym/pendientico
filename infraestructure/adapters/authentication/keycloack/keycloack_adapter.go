package adapters

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v13"
	ports_out "github.com/geordym/pendientico/domain/ports/out"
	configuration "github.com/geordym/pendientico/infraestructure/configuration/environment"
	"github.com/geordym/pendientico/infraestructure/middlewares"
)

type KeycloakAdapter struct {
	client *gocloak.GoCloak
	token  string
	realm  string
}

func NewKeycloakAdapter(client *gocloak.GoCloak, token, realm string) *KeycloakAdapter {
	return &KeycloakAdapter{
		client: client,
		token:  token,
		realm:  realm,
	}
}

func (k *KeycloakAdapter) SaveUser(email string, password string) (string, error) {
	kcUser := gocloak.User{
		Username: gocloak.StringP(email),
		Email:    gocloak.StringP(email),
		Enabled:  gocloak.BoolP(true),
	}

	userID, err := k.client.CreateUser(context.Background(), k.token, k.realm, kcUser)
	if err != nil {
		return "", err
	}

	err = k.client.SetPassword(context.Background(), k.token, userID, k.realm, password, false)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (k *KeycloakAdapter) GetUserAuthProviderIdLogged(ctx context.Context) (string, error) {
	claimsInterface := ctx.Value(middlewares.ClaimsKey{})
	if claimsInterface == nil {
		return "", fmt.Errorf("claims not found in context")
	}

	claims, ok := claimsInterface.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("claims have wrong type")
	}

	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return "", fmt.Errorf("sub claim not found")
	}

	return sub, nil
}

func NewKeycloakAdapterFromEnv(env configuration.Environment) (ports_out.AuthenticationProviderCommunication, error) {
	client := gocloak.NewClient(env.KeycloakURL)

	token, err := client.LoginClient(context.Background(), env.KeycloakClientID, env.KeycloakClientSecret, env.KeycloakRealm)
	if err != nil {
		return nil, err
	}

	return &KeycloakAdapter{
		client: client,
		token:  token.AccessToken,
		realm:  env.KeycloakRealm,
	}, nil
}

var _ ports_out.AuthenticationProviderCommunication = (*KeycloakAdapter)(nil)
