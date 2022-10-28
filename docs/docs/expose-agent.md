## Expose the agent's API over HTTPS

The actuated agent serves HTTP, and must be accessible by the actuated control plane.

We expect most of our pilot customers to be using hosts with public IP addresses, and the combination of an API token plus TLS is a battle tested combination.

For anyone running with private hosts, OpenFaaS Ltd's inlets product can be used to get incoming traffic over a secure tunnel

## For a host on a public cloud

If you're running the agent on a host with a public IP, you can use the built-in TLS mechanism in the actuated agent to receive a certificate from Let's Encrypt, valid for 90 days. The certificate will be renewed by the actuated agent, so there are no additional administration tasks required.

![Accessing the agent's endpoint built-in TLS and Let's Encrypt](images/builtin-tls.png)
> Pictured: Accessing the agent's endpoint built-in TLS and Let's Encrypt

Determine the public IP of your instance:

```bash
# curl -s http://checkip.amazonaws.com

141.73.80.100
```

Now imagine that your sub-domain is `agent.example.com`, you need to create a DNS A record of `agent.example.com=141.73.80.100`, changing both the sub-domain and IP to your own.

Once created, edit the start.sh file on the agent and add two flags:

```
--letsencrypt-domain agent.example.com \
--letsencrypt-email webmaster@agent.example.com
```

Your agent's endpoint URL is going to be: `https://agent.example.com` on port 443

## Private hosts - on-premises, behind NAT or at home

You'll need a way to expose the client to the Internet, which includes HTTPS encryption and a sufficient amount of connections/traffic per minute.

[Inlets](https://inlets.dev/) provides a quick and secure solution here. It is available on a [monthly subscription](https://openfaas.gumroad.com/l/inlets-subscription), bear in mind that the "Personal" plan is not for this kind of commercial use.

![Accessing the agent's private endpoint using an inlets-pro tunnel](images/tunnel-server.png)
> Pictured: Accessing the agent's private endpoint using an inlets-pro tunnel

Reach out to us if you'd like us to host a tunnel server for you, alternatively, you can follow the instructions below to set up your own.

The [inletsctl](https://github.com/inlets/inletsctl) tool will create a HTTPS tunnel server with you on your favourite cloud with a HTTPS certificate obtained from Let's Encrypt.

If you have just the one Actuated Agent:

```bash
export AGENT_DOMAIN=agent1.example.com
export LE_EMAIL=webmaster@agent1.example.com

arkade get inletsctl
sudo mv $HOME/.arkade/bin/inletsctl /usr/local/bin/

inletsctl create \
    --provider digitalocean \
    --region lon1 \
    --token-file $HOME/do-token \
    --letsencrypt-email $LE_EMAIL \
    --letsencrypt-domain $AGENT_DOMAIN
```

Then note down the tunnel's wss:// URL and token.

Then run a HTTPS client to expose your agent:

```bash
inlets-pro http client \
    --url $WSS_URL \
    --token $TOKEN \
    --upstream http://127.0.0.1:8081
```

For two or more Actuated Agents:

```bash
export AGENT_DOMAIN1=agent1.example.com
export AGENT_DOMAIN2=agent2.example.com
export LE_EMAIL=webmaster@agent1.example.com

arkade get inletsctl
sudo mv $HOME/.arkade/bin/inletsctl /usr/local/bin/

inletsctl create \
    --provider digitalocean \
    --region lon1 \
    --token-file $HOME/do-token \
    --letsencrypt-email $LE_EMAIL \
    --letsencrypt-domain $AGENT_DOMAIN1 \
    --letsencrypt-domain $AGENT_DOMAIN2
```

Then note down the tunnel's wss:// URL and token.

Then run a HTTPS client to expose your agent, using the unique agent domain, run the inlets-pro client on the Actuated Agents:

```bash
export AGENT_DOMAIN1=agent1.example.com
inlets-pro http client \
    --url $WSS_URL \
    --token $TOKEN \
    --upstream $AGENT1_DOMAIN=http://127.0.0.1:8081
```

```bash
export AGENT_DOMAIN2=agent2.example.com
inlets-pro http client \
    --url $WSS_URL \
    --token $TOKEN \
    --upstream $AGENT1_DOMAIN=http://127.0.0.1:8081
```

You can generate a systemd service (so that inlets restarts upon disconnection, and reboot) by adding `--generate=systemd > inlets.service` and running:

```bash
sudo cp inlets.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable inlets.service
sudo systemctl start inlets

# Check status with:
sudo systemctl status inlets
```

Your agent's endpoint URL is going to be: `https://$AGENT_DOMAIN`.

### Preventing the runner from accessing your local network

!!! warning "Network segmentation"
    Proper network segmentation of hosts running the actuated agent is required. This is to prevent runners from making outbound connections to other hosts on your local network. We will not accept any responsibility for your configuration.

If hardware isolation is not available, iptables rules may provide an alternative for isolating the runners from your network.

Imagine you were using a LAN range of `192.168.0.0/24`, with a router of `192.168.0.1`, then the following probes and tests show that the runner cannot access the host 192.168.0.101, and that nmap's scan will come up dry.

We add a rule to allow access to the router, but reject packets going via TCP or UDP to any other hosts on the network.

```bash
sudo iptables --insert CNI-ADMIN \
    --destination  192.168.0.1 --jump ACCEPT
sudo iptables --insert CNI-ADMIN \
    --destination  192.168.0.0/24 --jump REJECT -p tcp  --reject-with tcp-reset
sudo iptables --insert CNI-ADMIN \
    --destination  192.168.0.0/24 --jump REJECT -p udp --reject-with icmp-port-unreachable
```

You can test the efficacy of these rules by running nmap, mtr, ping and any other probing utilities within a GitHub workflow.

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
      - name: addr
        run: ip addr
      - name: route
        run: ip route
      - name: pkgs
        run: |
             sudo apt-get update && \
              sudo apt-get install traceroute mtr nmap netcat -qy
      - name: traceroute
        run: traceroute  192.168.0.101
      - name: Connect to ssh
        run: echo | nc  192.168.0.101 22
      - name: mtr
        run: mtr -rw  -c 1  192.168.0.101
      - name: nmap for SSH
        run: nmap -p 22  192.168.0.0/24
      - name: Ping router
        run: |
          ping -c 1  192.168.0.1
      - name: Ping 101
        run: |
          ping -c 1  192.168.0.101
```
