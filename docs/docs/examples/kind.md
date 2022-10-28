# Example: Kubernetes with KinD

You may need to access Kubernetes within your build. [KinD](https://kind.sigs.k8s.io) is a popular option, and easy to run in an action.

Certified for:

- [x] `x86_64`
- [x] `arm64` including Raspberry Pi 4

!!! info "Use a private repository"
    GitHub recommends using a private repository with self-hosted runners. Learn why in the [FAQ](/faq).

## Try out the action on your agent

Create a new file at: `.github/workspaces/build.yml` and commit it to the repository.

Note that it's important to make sure Kubernetes is responsive before performing any commands like running a Pod or installing a helm chart.

```yaml
name: build

on: push
jobs:
  start-kind:
    runs-on: actuated
    steps:
      - uses: actions/checkout@master
        with:
          fetch-depth: 1
      - name: get arkade
        uses: alexellis/setup-arkade@v1
      - name: get kubectl and kubectl
        uses: alexellis/arkade-get@master
        with:
          kubectl: latest
          kind: latest
      - name: Install Kubernetes kind
        run: |
          mkdir -p $HOME/.kube/
          kind create cluster --wait 300s
      - name: Wait until CoreDNS is ready
        run: |
          kubectl rollout status deploy/coredns -n kube-system --timeout=300s
      - name: Explore nodes
        run: kubectl get nodes -o wide
      - name: Explore pods
        run: kubectl get pod -A -o wide
      - name: Show kubelet logs
        run: docker exec kind-control-plane journalctl -u kubelet
```

To run this on ARM64, just change the actuated label to `actuated-aarch64`.
