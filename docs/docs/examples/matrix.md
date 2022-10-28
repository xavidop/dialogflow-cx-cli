# Example: matrix-build - run a VM per each job in a matrix

Use this sample to test launching multiple VMs in parallel.

Certified for:

- [x] `x86_64`
- [x] `arm64` including Raspberry Pi 4

!!! info "Use a private repository"
    GitHub recommends using a private repository with self-hosted runners. Learn why in the [FAQ](/faq).

## Try out the action on your agent

Create a new file at: `.github/workspaces/build.yml` and commit it to the repository.

```yaml
name: CI

on:
  pull_request:
    branches:
      - '*'
  push:
    branches:
      - master
      - main

jobs:
  arkade-e2e:
    name: arkade-e2e
    runs-on: actuated
    strategy:
      matrix:
        apps: [run-job,k3sup,arkade,kubectl,faas-cli]
    steps:
      - name: Get arkade
        run: |
          curl -sLS https://get.arkade.dev | sudo sh
      - name: Download app
        run: |
          echo ${{ matrix.apps }}
          arkade get ${{ matrix.apps }}
          file /home/runner/.arkade/bin/${{ matrix.apps }}
```

The matrix will cause a new VM to be launched for each item in the "apps" array.

