<p align="center">
  <a href="https://www.mojoauth.com">
    <img alt="MojoAuth" src="https://mojoauth.com/blog/assets/images/logo.svg" width="200" />
  </a>
</p>

<h1 align="center">
  MojoAuth Go SDK
</h1>


## Documentation

* [Configuration](https://mojoauth.com/docs/) - Everything you need to begin using the MojoAuth SDK.

## Installation

To install, run:
`go get github.com/mojoauth/go-sdk`

Import the package:

`import "github.com/mojoauth/go-sdk"`

Install all package dependencies by running `go get ./...` in the root folder of this SDK.  

## Usage

Take a peek:

Before making any API calls, the MojoAuth API client must be initialized with your MojoAuth API key.

Sample code:

```
cfg := go-sdk.Config{
  ApiKey: "<Enter ApiKey>",
}

mojoclient, err := go_mojoauth.NewMojoAuth(&cfg)

if err != nil {
  errors = errors + err.(mojoerror.Error).OrigErr().Error()
  //		respCode = 500
}

```
## How to contribute

We appreciate all kinds of contributions from anyone, be it finding an issue or writing a blog.

Please check the [contributing guide](CONTRIBUTING.md) to become a contributor.

## License

For more information on licensing, please refer to [License](https://github.com/MojoAuth/mojoauth-go/blob/main/LICENSE)
