# Example: Debug a job with SSH

If your tier and subscription includes debugging with SSH, then you can follow these instructions to get a shell into your self-hosted runner.

Certified for:

- [x] `x86_64`

!!! info "Use a private repository"
    GitHub recommends using a private repository with self-hosted runners. Learn why in the [FAQ](/faq).

## Try out the action on your agent

Create a secret for the repo or organisation for `SSH_GATEWAY_IP` using the IP address, or DNS address that you were provided with by your support team.

Create a `.github/workflows/workflow.yaml` file

```yaml
name: connect

on:
  pull_request:
    branches:
      - '*'
  push:
    branches:
      - master

permissions:
  id-token: write
  contents: read
  actions: read

jobs:
  connect:
    name: connect
    runs-on: actuated
    steps:
      - name: Setup SSH server for Actor
        uses: alexellis/setup-sshd-actor@master
      - name: Connect to the actuated SSH gateway
        uses: alexellis/actuated-ssh-gateway-action@master
        with:
          gatewayaddr: ${{ secrets.SSH_GATEWAY_IP }}
          secure: true
      - name: Setup a blocking tmux session
        uses: alexellis/block-with-tmux-action@master
```

Next, trigger a build.

Open `https://$SSH_GATEWAY_IP/list` in your browser and look for your session, you can log in using the SSH command outputted for you.

Watch a demo:

<iframe width="560" height="315" src="https://www.youtube.com/embed/l9VuQZ4a5pc" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
