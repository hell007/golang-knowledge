package jwt

import "github.com/dgrijalva/jwt-go"

const (
	DefaultContextKey = "iris-jwt"
)

// Config is a struct for specifying configuration options for the jwts middleware.
type Config struct {
	ValidationKeyGetter jwt.Keyfunc
	ContextKey          string
	ErrorHandler        errorHandler
	CredentialsOptional bool
	Extractor           TokenExtractor
	Debug               bool
	EnableAuthOnOptions bool
	SigningMethod       jwt.SigningMethod
	Expiration          bool
}
