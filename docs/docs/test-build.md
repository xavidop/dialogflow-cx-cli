# Start a build on your agent

Once you've [registered your GitHub organisation](register.md), and [set up your first runner](add-agent.md), you can either add actuated to an existing GitHub workflow, or create a test repository to see it in action.

We suggest creating a test build so that you can see how everything works before moving over to an existing repository. 

!!! warning "Don't use a public repository"
    Due to limitations in the design of GitHub's runner, we recommend using a private repository. Learn more in the [FAQ](/faq.md).

## Create a test build

This build will show you the specs, OS and Kernel name reported by the MicroVM.

1. Create a test repository and a GitHub Action

    Create `./.github/workflows/ci.yaml`:

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
    specs:
        name: specs
        runs-on: actuated
        steps:
        - uses: actions/checkout@v1
        - name: sleep
            run: |
            sleep 2
        - name: Check specs
            run: |
            ./specs.sh

    ```

    Note that the `runs-on:` field says `actuated` and not `ubuntu-latest`. This is how the actuated control plane knows to send this job to your agent.

    Then add `specs.sh` to the root of the repository:

    ```bash
    #!/bin/bash

    echo Information on main disk
    df -h /

    echo Memory info
    free -h

    echo Total CPUs:
    echo CPUs: $(nproc)

    echo CPU Model
    cat /proc/cpuinfo |grep "model name"

    echo Kernel and OS info
    uname -a

    echo OS
    cat /etc/os-release

    echo Egress IP:
    curl -s -L -S https://checkip.amazonaws.com
    ```

2. Hit commit, and watch the VM boot up.

    Do you have any questions or comments?

    Feel free to reach out to us over Slack in the public channel for support.

## Enable an existing repository

To add actuated to an existing repository, simply edit the workflow YAML file and change `runs-on:` to `runs-on: actuated`.

If you want to go back to a hosted runner, edit the field back to `runs-on: ubuntu-latest` or whatever you used prior to that.
