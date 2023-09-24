# xdb-apis
API protocol between xdb SDKs and [xdb service](https://github.com/xdblab/xdb)

It's based on OpenAPI 3.1, see [specification](https://spec.openapis.org/oas/v3.0.3)

## Development

## Prerequisites

Install openapi-generator if you haven't. See more [documentation](https://openapi-generator.tech/docs/installation)

An easy way to install openapi-generator CLI is to use Homebrew:
```
brew install openapi-generator
```

And to upgrade it:
```
brew update && brew upgrade openapi-generator
```

You may also need to upgrade it to the latest if it's older than what we are currently using.
It's okay to upgrade but not downgrade the version in a PR.


## Build
Just run `make` to re-generate the API code based on the OpenAPI schema. 

