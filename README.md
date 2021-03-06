<div align="center" style="text-align: center;">
  <img src="https://raw.githubusercontent.com/mfcochauxlaberge/gqlgen-extra/master/assets/logo.png" height="120">
  <br>
  <a href="https://github.com/mfcochauxlaberge/gqlgen-extra/actions?query=workflow%3ATest+branch%3Amaster">
    <img src="https://github.com/mfcochauxlaberge/gqlgen-extra/workflows/Test/badge.svg?branch=master">
  </a>
  <a href="https://github.com/mfcochauxlaberge/gqlgen-extra/actions?query=workflow%3ALint+branch%3Amaster">
    <img src="https://github.com/mfcochauxlaberge/gqlgen-extra/workflows/Lint/badge.svg?branch=master">
  </a>
  <a href="https://goreportcard.com/report/github.com/mfcochauxlaberge/gqlgen-extra">
    <img src="https://goreportcard.com/badge/github.com/mfcochauxlaberge/gqlgen-extra">
  </a>
  <a href="https://codecov.io/gh/mfcochauxlaberge/gqlgen-extra">
    <img src="https://img.shields.io/codecov/c/github/mfcochauxlaberge/gqlgen-extra">
  </a>
  <br>
  <a href="https://github.com/mfcochauxlaberge/gqlgen-extra/blob/master/go.mod">
    <img src="https://img.shields.io/badge/go%20version-1.13%2B-%2300acd7">
  </a>
  <a href="https://github.com/mfcochauxlaberge/gqlgen-extra/blob/master/go.mod">
    <img src="https://img.shields.io/github/v/release/mfcochauxlaberge/gqlgen-extra?include_prereleases&sort=semver">
  </a>
  <a href="https://github.com/mfcochauxlaberge/gqlgen-extra/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/mfcochauxlaberge/gqlgen-extra?color=a33">
  </a>
  <a href="https://pkg.go.dev/github.com/mfcochauxlaberge/gqlgen-extra?tab=doc">
    <img src="https://img.shields.io/static/v1?label=doc&message=pkg.go.dev&color=007d9c">
  </a>
</div>

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
