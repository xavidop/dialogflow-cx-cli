# Frequently Asked Questions (FAQ)

## How does it work?

`cxcli` has three main purposes:

1. Make the interaction with your Dialogflow CX agents from your laptop or your continuous integration pipelines easier than ever
2. Create testing tools that will help users build their Dialogflow CX agent
3. Interact with other Google Cloud APIs such as TTS and STT in a very easy way

## Who is `cxcli` for?

`cxcli` is primarily for software engineering teams who are currently using Dialogflow CX. It is recommended for machine learning engineers that usually work with STT, TTS, NLU and NLP technologies.

## What kind of machines/containers do I need for the `cxcli`?

You'll need either: a bare-metal host (your own, AWS i3.metal or Equinix Metal) or a VM that supports nested virtualisation such as those provided by Google Cloud, Azure, AWS DigitalOcean, etc. or a Linux or Windows container.

## When will Jenkins, GitLab CI, BitBucket Pipeline Runners, Drone or Azure DevOps be supported?

For the current phase, we're targeting GitHub Actions because it has fine-grained access controls and the ability to schedule exactly one build to a runner. The other CI systems will be available soon.

That said, if you're using these tools within your organisation, we'd like to hear from you.
So feel free to reach out to us if you feel `cxcli` would be a good fit for your team.

Feel free to contact us at: [dialogflowcxcli@gmail.com](mailto:dialogflowcxcli@gmail.com)

## What kind of access is required in my Google Cloud project?

Refer to the Authentication page [here](/overview/authentication)

## Can cxcli be used on public repos?

Yes, `cxcli` can be used on public and private repos.

## What's in the Container image and how is it built?

The Container image contains uses `alpine:latest` and the `cxcli` installed on it.

The image is built automatically using GitHub Actions and is available on a container registry.

## Is ARM64 supported?

Yes, `cxcli` is built to run on both Intel/AMD and ARM64 hosts. This includes a Raspberry Pi 4B, AWS Graviton, Oracle Cloud ARM instances and potentially any other ARM64 instances that support virtualisation.

## Are Windows or macOS supported?

Yes, in addition to Linux, Windows and macOS are also supported platforms for `cxcli` at this time on a AMD64 or ARM64 architecture.

## Is `cxcli` free and open-source?

`cxcli` is an open source tool, however, it interacts with Google Cloud APIs, so a Google Cloud account is required.

The website and documentation are available on GitHub and we plan to release some open source tools in the future for cxcli customers.

## Is there a risk that we could get "locked-in" to `cxcli`?

No, you can switch back to using either the `gcloud` CLI tool or the Google Cloud APIs at any time. Keep in mind that `cxcli` not only solves for a certain set of issues with both of those approaches but also simplifies the interaction with Google Cloud.

## Why is the brand called "cxcli" and "Dialogflow CX CLI" ?

The name of the software is `cxcli`, in some places "cxcli" is not available, and we liked "Dialogflow CX CLI" because it refers to the what the tool does.
