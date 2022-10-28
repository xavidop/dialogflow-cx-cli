# Example: Publish an OpenFaaS function

This example will publish an OpenFaaS function to GitHub's Container Registry (GHCR).

* The example uses Docker's buildx and QEMU for a multi-arch build
* Dynamic variables to inject the SHA and OWNER name from the repo
* Uses the token that GitHub assigns to the action to publish the containers.

You can also run this example on GitHub's own hosted runners.

[Docker CE](https://docker.io) is preinstalled in the actuated VM image, and will start upon boot-up.

Certified for:

- [x] `x86_64`

!!! info "Use a private repository"
    GitHub recommends using a private repository with self-hosted runners. Learn why in the [FAQ](/faq).

## Try out the action on your agent

For alexellis' repository called [alexellis/autoscaling-functions](https://github.com/alexellis/autoscaling-functions), then check out the `.github/workspaces/publish.yml` file:

* The "Setup QEMU" and "Set up Docker Buildx" steps configure the builder to produce a multi-arch image.
* The "OWNER" variable means this action can be run on any organisation without having to hard-code a username for GHCR.
* Only the bcrypt function is being built with the `--filter` command added, remove it to build all functions in the stack.yml.
* `--platforms linux/amd64,linux/arm64,linux/arm/v7` will build for regular Intel/AMD machines, 64-bit ARM and 32-bit ARM i.e. Raspberry Pi, most users can reduce this list to just "linux/amd64" for a speed improvement

Make sure you edit `runs-on:` and set it to `runs-on: actuated`

```yaml
name: publish

on:
  push:
    branches:
      - '*'
  pull_request:
    branches:
      - '*'

permissions:
  actions: read
  checks: write
  contents: read
  packages: write

jobs:
  publish:
    runs-on: actuated
    steps:
      - uses: actions/checkout@master
        with:
          fetch-depth: 1
      - name: Get faas-cli
        run: curl -sLSf https://cli.openfaas.com | sudo sh
      - name: Pull custom templates from stack.yml
        run: faas-cli template pull stack
      - name: Set up QEMU
        uses: docker/setup-qemu-action@v1
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v1
      - name: Get TAG
        id: get_tag
        run: echo ::set-output name=TAG::latest-dev
      - name: Get Repo Owner
        id: get_repo_owner
        run: >
          echo ::set-output name=repo_owner::$(echo ${{ github.repository_owner }} |
          tr '[:upper:]' '[:lower:]')
      - name: Docker Login
        run: > 
          echo ${{secrets.GITHUB_TOKEN}} | 
          docker login ghcr.io --username 
          ${{ steps.get_repo_owner.outputs.repo_owner }} 
          --password-stdin
      - name: Publish functions
        run: >
          OWNER="${{ steps.get_repo_owner.outputs.repo_owner }}" 
          TAG="latest"
          faas-cli publish
          --extra-tag ${{ github.sha }}
          --build-arg GO111MODULE=on
          --platforms linux/amd64,linux/arm64,linux/arm/v7
          --filter bcrypt
```

If you'd like to deploy the function, check out a more comprehensive example of how to log in and deploy in [Serverless For Everyone Else](https://store.openfaas.com/l/serverless-for-everyone-else)
