# Add your first agent to actuated

actuated is split into three parts:

1. An agent that you run on your own machines or VMs, which can launch a VM with a single-use GitHub Actions runner
2. A VM image launched by the agent, with all the preinstalled software found on a hosted GitHub Actions runner
3. Our own control plane that talks to GitHub on your behalf, and schedules builds across your fleet of agents.

We look after 2 and 3 which means you just have to set up one or more agent to get started.

!!! info "Have you registered your organisation yet?"
    Before you can add an agent, you or your GitHub organisation admin will need to install the: [Actuated GitHub App](register.md).

## Decide where to run your agent

There are three places you can run an agent:

1. Bare-metal on-premises (cheap, convenient, high performance)

    This could be a machine racked in your server room, under your desk, or in a co-lo somewhere.

    It can be a cheap and convenient way to re-use existing hardware.

    Make sure you segment or isolate the agent into its own subnet, VLAN, DMZ, or VPC so that it cannot access the rest of your network. If you are thinking of running an actuated runner at home, we can share some iptables rules that worked well for our own testing.

    If you need to build for ARM64, we've shown that [a Raspberry Pi 4 is faster than emulating ARM on a GitHub Hosted Runner](https://twitter.com/alexellisuk/status/1583092051398524928?s=20&t=2SelTpdc5idJLmayIu3Djw)

2. Bare-metal on the cloud (higher cost, convenient, high performance)

    You can provision bare-metal hosts in the cloud using any number of providers like AWS, Alibaba Cloud, Cherry Servers, Equinix Metal, FastHosts, OVHcloud, Scaleway and Vultr, [see a list here](https://github.com/alexellis/awesome-baremetal#bare-metal-cloud) 
    
    For Intel/AMD builds on AWS, you'll need to use [AWS i3.metal](https://aws.amazon.com/ec2/instance-types/i3/).

    For ARM64 builds on AWS, the [a1.metal](https://aws.amazon.com/ec2/instance-types/a1/) is ideal.

    If you're not using AWS, or want larger machines, [Equinix Metal](https://metal.equinix.com/) offer bare-metal as a service for both Intel/AMD and ARM64 machines.

    This option is both convenient and offers the highest performance available, however bare-metal machines tends to be priced higher than you may be used to with VMs.

    Bear in mind that you may be able to run a single, larger bare-metal machine where you used to need half a dozen cloud VMs, since actuated can schedule builds much more efficiently than the built-in self-hosted runner from GitHub.

3. Cloud Virtual Machines (VMs) that support nested virtualization (lowest cost, convenient, mid-level performance)

    Both [DigitalOcean](https://m.do.co/c/8d4e75e9886f) and [Google Compute Platform (GCP)](https://cloud.google.com/compute) (new customers get 300 USD free credits from GCP) support nested virtualisation on their Virtual Machines (VMs).

    This option may not have the raw speed and throughput of a dedicated, bare-metal host, but keeps costs low and is convenient for getting started.

The recommended Operating System for an Actuated Agent is: Ubuntu Server 22.04 or Ubuntu Server 20.04.

## Review the End User License Agreement (EULA)

Make sure you've read the [Actuated EULA](https://github.com/self-actuated/actuated/blob/master/EULA.md) before registering your organisation with the actuated GitHub App, or starting the agent binary on one of your hosts.

## Set up your first agent

1. Download the agent and installation script

    Once you've decided where to set up your first agent, you'll need to download the installation package from a container registry

    Install [crane](https://github.com/google/go-containerregistry/releases):

    ```bash
    curl -sLS https://get.arkade.dev | sudo sh
    arkade get crane
    sudo mv $HOME/.arkade/bin/crane /usr/local/bin/
    ```

    ```bash
    rm -rf agent
    mkdir -p agent
    crane export ghcr.io/openfaasltd/actuated-agent:latest | tar -xvf - -C ./agent
    ```
    
    Install the agent binary to `/usr/local/bin/`:

    ```bash
    sudo mv ./agent/agent* /usr/local/bin/
    ```

    Run the setup.sh script which will install all the required dependencies like containerd, CNI and Firecracker.

    ```bash
    cd agent
    sudo ./install.sh
    ```

    Create a file to store your license from Gumroad:

    ```bash
    mkdir -p ~/.actuated

    # Paste the contents, hit enter, then Control + D
    # Or edit the file with nano/vim
    cat > $HOME/.actuated/LICENSE
    ```

2. Generate an RSA keypair

    ```bash
    cd ~/.actuated/
    agent keygen
    ```

    The RSA keypair is only used to encrypt messages and cannot. RSA keys are sometimes used with SSH sessions, however actuated does not use any form of SSH at this time.
    
    This will write: `key_rsa` and `key_rsa.pub` to the current working folder.

    Keep the `key_rsa` private, we will not ask you to share this file with us.
    Share `key_rsa.pub` with us via email or Slack. This key is not confidential, so don't worry about sharing it.

    ```bash
    cat ~/.actuated/key_rsa.pub
    ```

3. Install the agent's authentication token.

    Create an API token for us to present when we send jobs to your Actuated Agent:
    
    ```bash
    openssl rand -base64 32 > ~/.actuated/TOKEN
    ```

    Encrypt the token with our public key and email `.actuated/TOKEN.ENC` to us, or share it with us on Slack:

    ```bash
    cat <<EOF > actuated.pem
    -----BEGIN RSA PUBLIC KEY-----
    MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAo9EC7IrP8zTE9jm8agPa
    m0D/sFfmAlchhskLZksO4ZYzDHK9fuQ9oEhPYVkgrU5TifbL5UchdsSn//ELSy2Q
    TPRQoXVMdzPgLCrn15U+Xr7KpV3iNBV1go+ZzNE/ymdyS2kCCjxYiBLVuymn20hA
    ZzqkHSyOeM6IrG+A462KfmN0vqIpubpMkoK/wSkSSDjN0SoMWc9gaAqEFEHkSt9+
    t65fIdzG0sKSEMb613WG+K/A/WBrcdqGHWoMG2h2CpK12tNobZEt3yCL0WVgkAKU
    VwaHniNYHn5niJHH/DgvXMWDECoKA1ZJyMdWC3MuIlyfWVzT5N7a/HPTzyzlrdCl
    bwIDAQAB
    -----END RSA PUBLIC KEY-----

    EOF

    agent encrypt --key ./actuated.pem \
        --in $HOME/.actuated/TOKEN \
        --out $HOME/.actuated/TOKEN.enc
    ```

    Post-pilot, we will provide a more automated way to exchange this token.

4. Add HTTPS for the agent's endpoint

    The actuated control plane will only communicate with a HTTPS endpoint to ensure properly encryption is in place. An API token is used in addition with the TLS connection for all requests.

    In addition, any bootstrap tokens sent to the agent are further encrypted with the agent's public key.

    For hosts with public IPs, you will need to use the built-in TLS provisioning with Let's Encrypt. For hosts behind a firewall, NAT or in a private datacenter, you can use inlets to create a secure tunnel to the agent.

    We're considering other models for after the pilot, for instance GitHub's own API has the runner make an outbound connection and uses long-polling.

    See also: [expose the agent with HTTPS](expose-agent.md)

4. Start the agent

    For an Intel/AMD Actuated Agent, create a `start.sh` file:

    ```bash
    #!/bin/bash

    echo Running Agent from: ./agent
    DOMAIN=agent1.example.com

    sudo -E agent up \
        --image-ref=ghcr.io/openfaasltd/actuated-ubuntu20.04:x86-64-latest \
        --kernel-ref=ghcr.io/openfaasltd/actuated-kernel-5.10.77:x86-64-latest \
        --letsencrypt-domain $DOMAIN \
        --letsencrypt-email webmaster@$DOMAIN
    ```

    For an Actuated Agent behind an [inlets tunnel](https://inlets.dev):

    ```bash
    #!/bin/bash

    echo Running Agent from: ./agent
    sudo -E agent up \
        --image-ref=ghcr.io/openfaasltd/actuated-ubuntu20.04:aarch64-latest \
        --kernel-ref=ghcr.io/openfaasltd/actuated-kernel-5.10.77:aarch64-latest \
        --listen-addr 127.0.0.1:
    ```

    For ARM64 Actuated Agents, change the prefix of the image tags from `x86-64-` to `aarch64-`

    You can also run the Actuated Agent software as a systemd unit file for automatic restarts and to start upon boot-up.

## Next steps

You can now start your first build and see it run on your actuated agent.

[Start a build on your agent](test-build.md)

See also: [Troubleshooting your agent](troubleshooting.md)
