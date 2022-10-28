# actuated

Actuated brings blazingly fast, secure builds to self-hosted CI runners.

## Building containers on self-hosted runners is slow and insecure

Most solutions that use containers for running Docker or Kubernetes in CI have very poor security boundaries. They require either privileged containers (root on the host), a shared Docker socket (root on the host), third-party tools which don't integrate well and still require root to mount folders, or user namespaces which come with their own limitations. The chances are, if you use Docker or K8s in your CI, and run with: [actions-runner-controller](https://github.com/actions-runner-controller/actions-runner-controller), Jenkins, or GitLab, then you may be compromising on security or user experience.

## Management is a nightmare

Self-hosted CI runners are continually out of date, and require fine-tuning to get all the right packages in place and Kernel modules to build containers and cloud-native software. You'll also have to spend extra time making sure builds don't conflict, and that they can't cause side effects to system-level packages. What if you need two different version of some software?

If you haven't felt this pain yet, then perhaps you're blissfully unaware or are not updating your packages?

> Are you [running privileged containers](https://learn.snyk.io/lessons/container-runs-in-privileged-mode/kubernetes/) for CI in your organisation? Are you sharing a Docker Socket (just as bad!)? Are you running [Docker in Docker (DIND)](https://jpetazzo.github.io/2015/09/03/do-not-use-docker-in-docker-for-ci/)? 🙈

## Self-managed runners are inefficient and overprovisioned

Self-hosted runners are typically over-provisioned meaning you're spending too much money.

Why are they over-provisioned? Because you never know how many jobs you'll have to run, so you have to make them bigger, or have too many hosts available.

Why are they inefficient?

By default, the self-hosted runner will only schedule one job per host at a time, because GitHub has no knowledge of the capacity of your machines. So each and every build you run could consume all the resources on the host. The second reason is that builds often conflict with one another causing side effects that only happen in CI and are really difficult to track down and reproduce.

Actuated uses VMs to slice up the whole machine, and can run many builds in parallel. The net effect is that your build queue will get cleared down much more quickly.

## Hands-free, VM-level isolation

Actuated provides a fast-booting microVM which can run Docker, Kubernetes and anything else you need, with full root on the VM, and no access to the host. Each environment is created just in time to take a build, and is removed immediately after.

Boot time is usually ~1-2 seconds for the VM, that extra second is because we start Docker as part of the boot-up process.

> What does "actuated" mean?
> 
> Something that activates or impels itself; specifically (a machine, device, etc.) that causes itself to begin operating automatically, self-activating.

We maintain a VM image that is updated regularly through an automated build, so you don't have to install SDKs, runtimes or language packs on your build machines.

Just enable automated updates on your server then install the actuated agent. We'll do the rest including managing efficient allocation across your fleet of servers, and updating the CI image.

And actuated will run your jobs efficiently across a fleet of hosts, or a single machine. They each need to be either bare-metal hosts (think: AWS Metal / Graviton, Equinix Metal, etc), or support nested virtualization (a feature available on GCP and DigitalOcean)

**Conceptual overview**

![Conceptual flow of starting up a new ephemeral runner](images/conceptual-high-level.png)

> Actuated will schedule builds across your fleet of agents, packing them in densely, without overloading the host. Each microVM will run just one build before being destroyed to ensure a clean, isolated build. 

Learn more in the [FAQ](faq.md)

## Watch a live demo

Alex shows you how actuated uses an isolated, immutable microVM to run K3s inside of a GitHub Action, followed by a matrix build that causes 5 VMs to be launched. You'll see how quick and easy it is to enable actuated, and how it can buffer and queue up jobs, when there's no remaining capacity in your fleet of agents.

<iframe width="560" height="315" src="https://www.youtube.com/embed/2o28iUC-J1w" title="YouTube video player" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

You can also watch a webinar that Alex recorded with Richard Case from Weaveworks on how microVMs compare to containers and legacy VMs, [you'll see Alex's demo at: 1:13:19](https://www.youtube.com/watch?v=CYCsa5e2vqg?t=4399).

## Get started

* [Register for the pilot](register.md)
* [Read the FAQ](faq.md)
* [Enable actuated for an existing repository](test-build.md)

## Got questions, comments or suggestions?

actuated is trademark of OpenFaaS Ltd.

You can contact the team working on actuated via email at: [contact@openfaas.com](mailto:contact@openfaas.com)

Follow [@selfactuated on Twitter](https://twitter.com/selfactuated) for updates and announcements

