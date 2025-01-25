package go_mojoauth

import (
	"errors"

	"github.com/mojoauth/go-sdk/mojoerror"
)

// domain is the default domain for API calls to MojoAuth
const domain = "https://api.mojoauth.com"

// MojoAuth struct holds context for intializing the MojoAuth client and the domain for API calls
// Domain can be changed after intialization
type Mojoauth struct {
	Context *Context
	Domain  string
}

// Config struct contains MojoAuth credentials and is used when initalizing the MojoAuth API client struct
type Config struct {
	ApiKey string
}

// Context struct is a field in the MojoAuth struct
type Context struct {
	ApiKey string
	Jwks string
	Token  string
}

// NewMojoAuth initializes a new MojoAuth struct with a Config struct
// Config struct must contain the ApiKey and ApiSecret of your MojoAuth site
// Example:
// 			cfg := lr.Config{
// 				ApiKey:    os.Getenv("APIKEY"),
// 				ApiSecret: os.Getenv("APISECRET"),
// 			}
// 			lrclient, _ := lr.NewMojoAuth(&cfg)
// It also takes optional variadic arguments
// Some APIs require for an access token to be passed with Authorization Bearer header
// Initialize MojoAuth struct with a token passed in the variadic argument like so:
// 			lrclient, _ := lr.NewMojoAuth(&cfg, map[string]string{"token": "9c3208ae-2848-4ac5-baef-41dd4103e263"})
func NewMojoAuth(cfg *Config, optionalArgs ...map[string]string) (*Mojoauth, error) {

	if cfg.ApiKey == "" {
		errMsg := "Must initialize MojoAuth client with ApiKey"
		err := mojoerror.New("IntializationError", errMsg, errors.New(errMsg))
		return nil, err
	}

	ctx := Context{
		ApiKey: cfg.ApiKey,
	}

	// If an access token is passed on initiation, set it in Context
	for _, arg := range optionalArgs {
		if arg["token"] != "" {
			ctx.Token = arg["token"]
		} else {
			ctx.Token = ""
		}
	}

	return &Mojoauth{
		Context: &ctx,
		Domain:  domain,
	}, nil
}
