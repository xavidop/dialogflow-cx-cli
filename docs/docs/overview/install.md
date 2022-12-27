# Install

You can install the pre-compiled binary (in several ways), using Docker or compiling it from source.

Below you can find the steps for each of them.

## Install the pre-compiled binary

### homebrew tap
1. Add the Hombrew tab:
```sh
brew tap xavidop/tap git@github.com:xavidop/homebrew-tap.git
brew update
```
2. Install the Dialogflow CX CLI:
```sh
brew install cxcli
```

### snapcraft

```sh
sudo snap install cxcli
```

### scoop

```powershell
scoop bucket add cxcli https://github.com/xavidop/scoop-bucket.git
scoop install cxcli
```

### chocolatey

```powershell
choco install cxcli
```

### apt

```sh
echo 'deb [trusted=yes] https://apt.fury.io/xavidop/ /' | sudo tee /etc/apt/sources.list.d/cxcli.list
sudo apt update
sudo apt install cxcli
```


### yum

```sh
echo '[cxcli]
name=Dialogflow CX CLI Repo
baseurl=https://yum.fury.io/xavidop/
enabled=1
gpgcheck=0' | sudo tee /etc/yum.repos.d/cxcli.repo
sudo yum install cxcli
```

### aur

```sh
yay -S cxcli-bin
```

### deb, rpm and apk packages

Download the `.deb`, `.rpm` or `.apk` packages from the [OSS releases page][releases] and install them with the appropriate tools.

### go install

```sh
go install github.com/xavidop/dialogflow-cx-cli@latest
```

### bash script

```sh
curl -sfL https://cxcli.xavidop.me/static/run | bash
```

#### Additional Options
You can also set the `VERSION` variable to specify
a version instead of using latest.

You can also pass flags and args to cxcli:

```bash
curl -sfL https://cxcli.xavidop.me/static/run |
    VERSION=__VERSION__ bash -s -- version
```

!!! tip
    This script does not install anything, it just downloads, verifies and
    runs cxcli.
    Its purpose is to be used within scripts and CIs.

### manually

Download the pre-compiled binaries from the [releases page][releases] and copy them to the desired location.


## Verifying the artifacts

### binaries

All artifacts are checksummed, and the checksum file is signed with [cosign][].

1. Download the files you want, and the `checksums.txt`, `checksum.txt.pem` and `checksums.txt.sig` files from the [releases][releases] page:
    ```sh
    wget https://github.com/xavidop/dialogflow-cx-cli/releases/download/__VERSION__/checksums.txt
    wget https://github.com/xavidop/dialogflow-cx-cli/releases/download/__VERSION__/checksums.txt.sig
    wget https://github.com/xavidop/dialogflow-cx-cli/releases/download/__VERSION__/checksums.txt.pem
    ```
1. Verify the signature:
    ```sh
    COSIGN_EXPERIMENTAL=1 cosign verify-blob \
    --cert checksums.txt.pem \
    --signature checksums.txt.sig \
    checksums.txt
    ```
1. If the signature is valid, you can then verify the SHA256 sums match with the downloaded binary:
    ```sh
    sha256sum --ignore-missing -c checksums.txt
    ```

### docker images

Our Docker images are signed with [cosign][].

Verify the signatures:

```sh
COSIGN_EXPERIMENTAL=1 cosign verify xavidop/cxcli
```

!!! info
    The `.pem` and `.sig` files are the image `name:tag`, replacing `/` and `:` with `-`.

## Running with Docker

You can also use it within a Docker container.
To do that, you'll need to execute something more-or-less like the examples below.

Registries:

- [`xavidop/cxcli`](https://hub.docker.com/r/xavidop/cxcli)
- [`ghcr.io/xavidop/cxcli`](https://github.com/xavidop/dialogflow-cx-cli/pkgs/container/cxcli)

Example usage:

```sh
docker run --rm \
    xavidop/cxcli cxcli version
```

Note that the image will almost always have the last stable Go version.

If you need more things, you are encouraged to keep your own image. You can
always use cxcli's [own Dockerfile][dockerfile] as an example though
and iterate from that.

[dockerfile]: https://github.com/xavidop/dialogflow-cx-cli/blob/master/Dockerfile
[releases]: https://github.com/xavidop/dialogflow-cx-cli/releases
[cosign]: https://github.com/sigstore/cosign

## Compiling from source

Here you have two options:

If you want to contribute to the project, please follow the
steps on our [contributing guide](/community/contributing/).

If you just want to build from source for whatever reason, follow these steps:

**clone:**

```sh
git clone https://github.com/xavidop/dialogflow-cx-cli
cd dialogflow-cx-cli
```

**get the dependencies:**

```sh
go mod tidy
```

**build:**

```sh
go build -o cxcli .
```

**verify it works:**

```sh
./cxcli version
```