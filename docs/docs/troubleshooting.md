# Troubleshooting

## Getting support

All customers have access to a public Slack channel for support and collaboration.

Enterprise customers may also have an upgraded SLA for support tickets via email and access to a private Slack channel.

## You need to rotate the authentication token used for your agent

There should not be many reasons to rotate this token, however, if something's happened and it's been leaked or an employee has left the company, contact us via email for the update procedure.

## You need to rotate your private/public keypair

Your private/public keypair is comparable to an SSH key, although it cannot be used to gain access to your agent via SSH.

If you need to rotate it for some reason, please contact us by email as soon as you can.

## Disk space is running out

This can be observed by running `df -h` or `df -h /`.

Over time, the various "fat" VM images that we ship to you may fill up the disk space on your machine.

You can delete all images, including the current image with the following command.

```bash
sudo ctr -n mvm image ls -q | xargs sudo ctr -n mvm image rm
```

We do not recommend running this on a cron schedule since the maintenance command will cause your agent to download the latest fat VM image ~ 1GB+/- again.

An alternative for a cron schedule would need to exclude the current image being used:

```bash
CURRENT="ghcr.io/openfaasltd/actuated-ubuntu:20.0.4-2022-09-30-1357"
sudo ctr -n mvm image ls -q |grep -v $CURRENT | xargs sudo ctr -n mvm image rm
```

## Your agent has been offline or unavailable for a significant period of time

If your agent has been offline for a significant period of time, then our control plane will have disconnected it from its pool of available agents.

Contact us via Slack to have it reinstated.

## The devmapper snapshotter is not available or running

The actuated agent uses the devmapper snapshotter for containerd, which emulates a thin-provisioned block device. Performance can be improved by attaching a dedicated disk or partition, but in our testing the devmapper works well enough for most workloads.

The dmsetup.sh script must be run upon every fresh boot-up of the host. It enables Firecracker to use snapshots to save disk space when launching new VMs.

If you see an error about the "devmapper" snapshot driver, then run the `dmsetup.sh` shell script then restart containerd:

```bash
./dmsetup.sh
sudo systemctl daemon-reload
sudo systemctl restart containerd
```

## Your builds are slower than expected

* Check free disk space (`df -h`)
* Check for unattended updates/upgrades (`ps -ef | grep unattended-upgrades`) and (`ps -ef | grep apt`)

If you're using spinning disks, then consider switching to SSDs. If you're already using SSDs, consider using PCIe/NVMe SSDs.

Finally, we do have another way to speed up microVMs by attaching another drive or partition to your host. Contact us for more information.