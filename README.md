# gqlgen-extra

This repository offers packages that can be used on a [gqlgen](https://gqlgen.com/) project to enable extra features.

## Packages

### types

`types` offers the following types that can be used in your GraphQL schemas.

 - Decimal (based on `github.com/cockroachdb/adp`)

You may configure you `gqlgen` project with the following example to use those types.

```
models:
  Decimal:
    model:
      - github.com/mfcochauxlabeerge/gqlgen-extra/types.Decimal
```

## Demo

You can find a demo gqlgen project in `demo/` that uses all of the packages presented in this repository.
