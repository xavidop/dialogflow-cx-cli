# Frequently Asked Questions (FAQ)

## How does it work?

Actuated has three main parts:

1. an agent which knows how to run VMs, you install this on your hosts
2. a VM image and Kernel that we build which has everything required for Docker, KinD and K3s
3. a multi-tenant control plane that we host, which tells your agents to start VMs and register a runner on your GitHub organisation

The multi-tenant control plane is run and operated by OpenFaaS Ltd.

![Conceptual flow of starting up a new ephemeral runner](images/conceptual.png)

> The conceptual overview showing how a MicroVM is requested by the control plane.

MicroVMs are only started when needed, and are registered with GitHub by the official GitHub Actions runner, using a short-lived registration token. The token is been encrypted with the public key of the agent. This ensures no other agent could use the token to bootstrap a token to the wrong organisation.

Learn more: [Self-hosted GitHub Actions API](https://docs.github.com/en/rest/actions/self-hosted-runners#create-a-registration-token-for-an-organization)

## How does actuated compare to a self-hosted runner?

A self-hosted runner is a machine on which you've installed and registered the a GitHub runner.

Quite often these machines suffer from some, if not all of the following issues:

* They require several hours to get all the required packages correctly installed to mirror a hosted runner
* You never update them out of fear of wasting time or breaking something which is working, meaning your supply chain is at risk
* Builds clash, if you're building a container image, or running a KinD cluster, names will clash, dirty state will be left over

We've heard in user interviews that the final point of dirty state can cause engineers to waste several days of effort chasing down problems.

Actuated uses a one-shot VM that is destroyed immediately after a build is completed.

## Who is actuated for?

actuated is primarily for software engineering teams who are currently using GitHub Actions. A GitHub organisation is required for installation, and runners are attached to individual repositories as required, to execute builds.

## What kind of machines do I need for the agent?

You'll need either: a bare-metal host (your own, AWS i3.metal or Equinix Metal), or a VM that supports nested virtualisation such as those provided by GCP and DigitalOcean.

## When will Jenkins, GitLab CI, BitBucket Pipeline Runners, Drone or Azure DevOps be supported?

For the pilot phase, we're targeting GitHub Actions because it has fine-grained access controls and the ability to schedule exactly one build to a runner. Most other CI systems expect self-hosted runners to perform many builds, and we believe that to be an anti-pattern. We'll offer advice to teams accepted into the pilot who wish to evaluate GitHub Actions or migrate away from another solution.

That said, if you're using these tools within your organisation, and face similar issues or concerns, we'd like to hear from you. And we have a proof of concept that works with GitLab CI, so feel free to reach out to us if you feel actuated would be a good fit for your team.

Feel free to contact us at: [contact@openfaas.com](mailto:contact@openfaas.com)

## What kind of access is required to my GitHub Organisation?

GitHub Apps provide fine-grained privileges, access control, and event data.

Actuated integrates with GitHub using a [GitHub App](https://docs.github.com/en/developers/apps/getting-started-with-apps/about-apps).

The actuated GitHub App will request:

* Administrative access to add/remove GitHub Actions Runners to individual repositories
* Events via webhook for Workflow Runs and Workflow Jobs

## Can GitHub's self-hosted runner be used on public repos?

The GitHub team recommends only running their self-hosted runners on private repositories.

Why?

On first glance, it seems like this might be due to how most people re-use a runner, and register it to process many jobs. It may even be because a bad actor could scan the local network of the runner and attempt to gain access to other systems. Actuated and iptables can largely fix both of these issues.

So, can you use a self-hosted runner on a public repo?

Through VM-level isolation, the primary concerns is resolved, because every run is started in an immutable VM.

> A bad actor could compromise the system or install malware leaving side-effects for future builds.

The second issue is that a bad actor could use the runner to run network scans or attacks against remote hosts.

This is a very hard problem to solve because a GitHub Action is a remote code execution (RCE) environment.

With Actuated, we can restrict builds on public repositories to organisation members only. If that's of interest, let us know.

## How many builds does a single actuated VM run?

When a VM starts up, it runs the GitHub Actions Runner ephemeral (aka one-shot) mode, so in can run at most one build. After that, the VM will be destroyed.

See also: [GitHub: ephemeral runners](https://docs.github.com/en/actions/hosting-your-own-runners/autoscaling-with-self-hosted-runners#using-ephemeral-runners-for-autoscaling)

## How are VMs scheduled?

VMs are placed efficiently across your Actuated Agents using a simple algorithm based upon the amount of RAM reserved for the VM.

Autoscaling of VMs is automatic. Let's say that you had 10 jobs pending, but given the RAM configuration, only enough capacity to run 8 of them? The second two would be queued until capacity one or more of those 8 jobs completed.

If you find yourself regularly getting into a queued state, there are three potential changes to consider:

1. Using Actuated Agents with more RAM
2. Allocated less RAM to each job
3. Adding more Actuated Agents

The plan you select will determine how many Actuated Agents you can run, so consider 1. and 2. before 3.

## Do I need to auto-scale the Actuated Agents?

If you haven't, read the previous section.

Most teams that we've interviewed said that a small static pool of Actuated Agents would satisfy their build requirements. For the pilot period, we are not offering auto-scaling of Actuated Agents.

If you feel that is a requirement for your team, set up some time to tell us why and we'll see if we can help.

## What's in the VM image and how is it built?

The VM image contains similar software to the hosted runner image: `ubuntu-latest` offered by GitHub. Unfortunately, GitHub does not publish this image, so we've done our best through user-testing to reconstruct it, including all the Kernel modules required to run Kubernetes and Docker.

The image is built automatically using GitHub Actions and is available on a container registry.

## How easy is it to debug a runner?

OpenSSH is pre-installed, but it will be inaccessible from your workstation by default.

So to connect to it, you can use an [inlets tunnel](https://inlets.dev/), Wireguard VPN or Tailscale ephemeral token (beware, Tailscale is not free for your commercial use) to log into any agent.

We recommend you add your SSH key and disable login with a password.

We're also considering an automated SSH gateway and a convenient CLI for actuated customers. Let us know if you'd like to try this out.

## What do I need to change in my workflows?

Very little, just add / set `runs-on: actuated`

## Is ARM64 supported?

Yes, actuated is built to run on both Intel/AMD and ARM64 hosts, check your subscription plan to see if ARM64 is included. This includes a Raspberry Pi 4B, AWS Graviton, Oracle Cloud ARM instances and potentially any other ARM64 instances which support virtualisation.

## How does actuated compare to a actions-runtime-controller (ARC)?

[actions-runtime-controller (ARC))](https://github.com/actions-runner-controller/actions-runner-controller) is maintained by [Yusuke Kuoka](https://github.com/mumoshu).

Its primary use-case is scale GitHub's self-hosted actions runner using Pods in a Kubernetes cluster. ARC is self-hosted software which means its setup and operation are complex, requiring you to create an properly configure a GitHub App along with its keys. For actuated, you only need to run a single binary on each of your runner hosts and send us an encrypted bootstrap token.

If you're running `npm install` or `maven`, then this may be a suitable isolation boundary for you.

The default mode for ARC is a reuseable runner, which can run many jobs, and each job could leave side-effects or poison the runner for future job runs.

If you need to build a container, in a container, on a Kubernetes node offers little isolation or security boundary.

What if ARC is configured to use "rootless" containers? With a rootless container, you lose access to "root" and `sudo`, both of which are essential in any kind of CI job. Actuated users get full access to root, and can run `docker build` without any tricks or losing access to `sudo`. That's the same experience you get from a hosted runner by GitHub, but it's faster because it's on your own hardware.

You can even run minikube, KinD, K3s and OpenShift with actuated without any changes.

ARC runs a container, so that should work on any machine with a modern Kernel, however actuated runs a VM, in order to provide proper isolation.

That means ARC runners can run pretty much anywhere, but actuated runners need to be on a bare-metal machine, or a VM that supports nested virtualisation.

See also: [Where can I run my agents?](/add-agent.txt)

## Are Windows or MacOS supported?

Linux is the only supported platform for actuated at this time on a AMD64 or ARM64 architecture. We may consider other operating systems in the future, feel free to reach out to us.

### Doesn't Kaniko fix all this for ARC?

[Kaniko](https://github.com/GoogleContainerTools/kaniko), by Google is an open source project for building containers. It's usually run as a container itself, and usually will require root privileges in order to mount the various filesystems layers required.

If you're an ARC user and for various reasons, cannot migrate away to a more secure solution, Kaniko may be a step in the right direction. Google Cloud users could also create a dedicated node pool with gVisor enabled, for some additional isolation.

However, it can only build containers, and still requires root, and itself is often run in Docker, so we're getting back to the same problems that actuated set out to solve.

In addition, Kaniko cannot and will not help you to run that container that you've just built to validate it to run end to end tests, neither can it run a KinD cluster, or a Minikube cluster.

## Is Actuated free and open-source?

Actuated is commercial software developed by OpenFaaS Ltd. A subscription will be required to use the software.

[Read the End User License Agreement (EULA)](https://github.com/self-actuated/actuated/blob/master/EULA.md)

The website and documentation are available on GitHub and we plan to release some open source tools in the future for actuated customers. 

## Is there a risk that we could get "locked-in" to actuated?

No, you can move back to either hosted runners (pay per minute from GitHub) or self-managed self-hosted runners at any time. Bear in mind that actuated solves for a certain set of issues with both of those approaches.

## Why is the brand called "actuated" and "selfactuated"?

The name of the software is actuated, in some places "actuated" is not available, and we liked "selfactuated" more than "actuatedhq" or "actuatedio" because it refers to the hybrid experience of self-hosted runners.

