package ports_out


type AuthenticationProviderCommunication interface {
    SaveUser(email string, password string) (string, error)
}
