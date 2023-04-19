# Contributing

By participating in this project, you agree to abide by our
[code of conduct](https://github.com/xavidop/dialogflow-cx-cli/blob/master/CODE_OF_CONDUCT.md).

## Set up your machine

`cxcli` is written in [Go](https://golang.org/).

Prerequisites:

- [Go 1.19+](https://golang.org/doc/install)

Other things you might need to run the tests:

- [cosign](https://github.com/sigstore/cosign)
- [Docker](https://www.docker.com/)
- [Podman](https://podman.io/)
- [Snapcraft](https://snapcraft.io/)
- [Syft](https://github.com/anchore/syft)

Clone `cxcli` anywhere:

```sh
git clone git@github.com:xavidop/dialogflow-cx-cli.git
```

`cd` into the directory and install the dependencies:

```sh
go mod tidy
```

A good way of making sure everything is all right is by running the build:

```sh
go build -o cxcli .
```

## Test your change

You can create a branch for your changes and try to build from the source as you go:

```sh
go build -o cxcli .
```

## Create a commit

Commit messages should be well formatted, and to make that "standardized", we
are using Conventional Commits.

You can follow the documentation on
[their website](https://www.conventionalcommits.org).

## Submit a pull request

Push your branch to your `cxcli` fork and open a pull request against the main branch.

## Financial contributions

You can contribute in our GitHub Sponsors or to any of the contributors directly.
See [this page](https://cxcli.xavidop.me/sponsors) for more details.
