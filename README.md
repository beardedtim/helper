## development

_**Starting**_

```sh
docker-commpose -f docker-compose.dev.yaml up
```

_**Seeing API Available**_

```sh
curl http://localhost:8080/openapi.json
```

### Building Static Assets

__**Build CSS*__

```sh
cd static
pnpm install
npx postcss ./css --dir ../public/css
```

### Building Template Files

```sh
templ generate
```

If that complains about `templ` not being installed run

```sh
go install github.com/a-h/templ/cmd/templ@latest
```