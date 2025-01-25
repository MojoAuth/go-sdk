// The mojobody package holds the structs to be encoded as the body in POST and PUT calls
// These structs are meant to serve as convenient measures assisting API calls provided by the MojoAuth
// Go SDK
// All functions in this SDK takes interface{} as the body, but initiating your
// data in the appropriate struct and passing in place of the body when calling the SDK functions
// will ensure the parameters submitted are correctly formatted and named for the MojoAuth APIs
// The usage of the structs in this package is optional and provided for convenience only
// Majority of methods take map[string]string as body parameter as well.
// These structs provide reference only, and do not include optional parameters
package mojobody

import "time"

type LoginResponse struct {
	Guid   string    `json:"guid"`
	Status bool      `json:"status"`
	Expiry time.Time `json:"expiry"`
}
type TokenResponse struct {
	IsValid bool   `json:"isValid"`
	Token   string `json:"access_token"`
}

type Key struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

type JwkResponse struct {
	Keys []Key  `json:"keys"`
}

