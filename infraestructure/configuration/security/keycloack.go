package security

import (
    "context"
    "log"
    "os"

    "github.com/coreos/go-oidc"
    "github.com/joho/godotenv"
)

var (
    ctx      = context.Background()
    provider *oidc.Provider
    Verifier *oidc.IDTokenVerifier
)

func InitKeycloak() {
    if err := godotenv.Load(); err != nil {
        log.Println("No se pudo cargar .env, usando variables de entorno del sistema")
    }

    keycloakURL := os.Getenv("KEYCLOAK_URL")
    realm := os.Getenv("KEYCLOAK_REALM")
    clientID := os.Getenv("KEYCLOAK_CLIENT_ID")

    if keycloakURL == "" || realm == "" || clientID == "" {
        log.Fatal("Faltan variables de entorno de Keycloak")
    }

    var err error
    provider, err = oidc.NewProvider(ctx, keycloakURL+"/realms/"+realm)
    if err != nil {
        log.Fatalf("Error al crear proveedor OIDC: %v", err)
    }

    Verifier = provider.Verifier(&oidc.Config{
        ClientID: clientID,
    })
}
