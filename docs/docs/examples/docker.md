# Example: Kubernetes with KinD

[Docker CE](https://docker.io) is preinstalled in the actuated VM image, and will start upon boot-up.

Certified for:

- [x] `x86_64`
- [x] `arm64` including Raspberry Pi 4

!!! info "Use a private repository"
    GitHub recommends using a private repository with self-hosted runners. Learn why in the [FAQ](/faq).

## Try out the action on your agent

Create a new file at: `.github/workspaces/build.yml` and commit it to the repository.

Try running a container to ping Google for 3 times:

```yaml
name: build

on: push
jobs:
  ping-google:
    runs-on: actuated
    steps:
      - uses: actions/checkout@master
        with:
          fetch-depth: 1
      - name: Run a ping to Google with Docker
        run: |
          docker run --rm -i alpine:latest ping -c 3 google.com
```

Build a container with Docker:

```yaml
name: build

on: push
jobs:
  build-in-docker:
    runs-on: actuated
    steps:
      - uses: actions/checkout@master
        with:
          fetch-depth: 1
      - name: Build inlets-connect using Docker
        run: |
          git clone --depth=1 https://github.com/alexellis/inlets-connect
          cd inlets-connect
          docker build -t inlets-connect .
          docker images
```

To run this on ARM64, just change the actuated label to `actuated-aarch64`.