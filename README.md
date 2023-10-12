## Tools Used

_**Languages or Development Tooling**_

- [Go](https://go.dev) as language
- [Docker](https://www.docker.com) as container
  - Also using Docker Compose for backing services and spinning up containers
- [PNPM](https://pnpm.io) for installing `static` dependencies
- [Air](https://github.com/cosmtrek/air) for hot reloading of Go code one changes

_**Major Frameworks/tools**_

- [Gin](https://gin-gonic.com) for basic Express-like Routing
- [Tonic](https://pkg.go.dev/github.com/loopfz/gadgeto/tonic) for interacting with Fizz
- [Fizz](https://github.com/wI2L/fizz) to generate OpenAPI spec for API
- [Templ](https://templ.guide) for SSR of HTML
- [Postcss](https://postcss.org) for generating CSS
- [Babel](https://babeljs.io) for generating JS 

## Development

_**Starting**_

```sh
docker-commpose -f docker-compose.dev.yaml up
```

_**Seeing API Available**_

```sh
curl http://localhost:8080/openapi.json
```

### Building Static Assets

__**Install Static Deps**_

```sh
cd static
pnpm install
```

__**Build CSS*__

```sh
cd static
npx postcss ./css --dir ../public/css
```

__**Build JS**__

```sh
cd static
npx babel js --out-dir ../public/js
```

### Building Template Files

```sh
templ generate
```

If that complains about `templ` not being installed run

```sh
go install github.com/a-h/templ/cmd/templ@latest
```