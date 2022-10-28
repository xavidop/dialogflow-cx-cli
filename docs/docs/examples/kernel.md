# Example: Test that compute time by compiling a Kernel

Use this sample to test the raw compute speed of your hosts by building a Kernel.

Certified for:

- [x] `x86_64`

!!! info "Use a private repository"
    GitHub recommends using a private repository with self-hosted runners. Learn why in the [FAQ](/faq).

## Try out the action on your agent

Create a new file at: `.github/workspaces/build.yml` and commit it to the repository.

```yaml
name: microvm-kernel

on: push
jobs:
  microvm-kernel:
    runs-on: actuated
    steps:
      - name: free RAM
        run: free -h
      - name: List CPUs
        run: nproc
      - name: get build toolchain
        run: |
          sudo apt update -qy
          sudo apt-get install -qy \
            git \
            build-essential \
            kernel-package \
            fakeroot \
            libncurses5-dev \
            libssl-dev \
            ccache \
            bison \
            flex \
            libelf-dev \
            dwarves
      - name: clone linux
        run: |
          time git clone https://github.com/torvalds/linux.git linux.git --depth=1 --branch v5.10
          cd linux.git
          curl -o .config -s -f https://raw.githubusercontent.com/firecracker-microvm/firecracker/main/resources/guest_configs/microvm-kernel-x86_64-5.10.config
          echo "# CONFIG_KASAN is not set" >> .config
      - name: make config
        run: |
          cd linux.git 
          make oldconfig
      - name: Make vmlinux
        run: |
          cd linux.git
          time make vmlinux -j$(nproc)
          du -h ./vmlinux
```

When you have a build time, why not change `runs-on: actuated` to `runs-on: ubuntu-latest` to compare it to a hosted runner from GitHub?

Here's our test, where our own machine built the Kernel 4x faster than a hosted runner:

[![Faster Kernel builds](https://pbs.twimg.com/media/FfGEFrxXoAAg9Vn?format=jpg&name=large)](https://twitter.com/alexellisuk/status/1581190198276526080/)
