<!--

When adding a new project to the list, add an entry to the table under the CNCF Projects section. The table is sorted
by Maturity, than by project name. Use the below template to add project details in the corresponding location in the
doc itself.

If a project changes maturity level, update the location in the project table and relocate the project details to
the new location within the specific maturity level section.

NOTE: Frequently used links have been given their own markdown reference links. These cover things such as licenses and
CLAs used by the various projects. You will find a list below the project template available to use for your own
convenience.


##### PROJECT TEMPLATE #####

### ${PROJECT-FULL-NAME}

*"Quote or brief project description"* - [ ][${PROJECT}-overview]

- **Project Repository:**
- **Contributor Guide:** [ ][${PROJECT}-contributor-guide]
- **Chat:** [ ][${PROJECT}-chat]
- **Developer List/Forum:** [ ][${PROJECT}-dev-list]
- **License:**  [ ][${PROJECT}-license]
- **Legal Requirements:** [ ][${PROJECT}-legal]

[${PROJECT}-overview]:
[${PROJECT}-contributor-guide]:
[${PROJECT}-chat]:
[${PROJECT}-dev-list]:
[${PROJECT}-license]:
[${PROJECT}-legal]:

##### PROJECT TEMPLATE #####

-->

<!-- Frequently Used Links -->
[CNCF-CLA]: https://github.com/cncf/cla
[DCO]: https://developercertificate.org/
[apache-license]: https://choosealicense.com/licenses/apache-2.0/
[mit-license]: https://choosealicense.com/licenses/mit/
[freenode]: https://freenode.net


# CNCF Projects

All projects of the Cloud Native Computing Foundation are classified with one of three stages of maturity:

- [Graduated](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#graduation-stage)
- [Incubated](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#incubating-stage)
- [Sandboxed](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#sandbox-stage)

[CNCF Graduation Criteria](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc) are documented in
the CNCF TOC repo. The document describes the maturity stages of the projects.

| Project Name                            | Maturity                           | Focus                   |
|:---------------------------------------:|:----------------------------------:|:-----------------------:|
|        [containerd](#containerd)        |  [Graduated](#graduated-projects)  |    Container Runtime    |
|           [CoreDNS](#coredns)           |  [Graduated](#graduated-projects)  |    Service Discovery    |
|             [Envoy](#envoy)             |  [Graduated](#graduated-projects)  |      Service Mesh       |
|        [Kubernetes](#kubernetes)        |  [Graduated](#graduated-projects)  |      Orchestration      |
|        [Prometheus](#prometheus)        |  [Graduated](#graduated-projects)  |       Monitoring        |
|               [CNI](#cni)               | [Incubating](#incubating-projects) |     Networking API      |
|           [Fluentd](#fluentd)           | [Incubating](#incubating-projects) |         Logging         |
|              [etcd](#etcd)              | [Incubating](#incubating-projects) |     Key/Value Store     |
|              [gRPC](#grpc)              | [Incubating](#incubating-projects) |  Remote Procedure Call  |
|            [Harbor](#harbor)            | [Incubating](#incubating-projects) |        Registry         |
|              [Helm](#helm)              | [Incubating](#incubating-projects) |   Package Management    |
|            [Jaeger](#jaeger)            | [Incubating](#incubating-projects) |   Distributed Tracing   |
|           [Linkerd](#linkerd)           | [Incubating](#incubating-projects) |      Service Mesh       |
|              [NATS](#nats)              | [Incubating](#incubating-projects) |        Messaging        |
|            [Notary](#notary)            | [Incubating](#incubating-projects) |        Security         |
|       [OpenTracing](#opentracing)       | [Incubating](#incubating-projects) | Distributed Tracing API |
|               [rkt](#rkt)               | [Incubating](#incubating-projects) |    Container Runtime    |
|              [Rook](#rook)              | [Incubating](#incubating-projects) |         Storage         |
|               [TUF](#tuf)               | [Incubating](#incubating-projects) |  Software Update Spec   |
|            [Vitess](#vitess)            | [Incubating](#incubating-projects) |         Storage         |
|        [Buildpacks](#buildpacks)        |    [Sandbox](#sandbox-projects)    |        Packaging        |
|       [CloudEvents](#cloudevents)       |    [Sandbox](#sandbox-projects)    |       Serverless        |
|            [Cortex](#cortex)            |    [Sandbox](#sandbox-projects)    |       Monitoring        |
|         [Dragonfly](#dragonfly)         |    [Sandbox](#sandbox-projects)    |   Image Distribution    |
|             [Falco](#falco)             |    [Sandbox](#sandbox-projects)    |   Container Security    |
| [Open Policy Agent](#open-policy-agent) |    [Sandbox](#sandbox-projects)    |         Policy          |
|       [OpenMetrics](#openmetrics)       |    [Sandbox](#sandbox-projects)    |         Tooling         |
|            [SPIFFE](#spiffe)            |    [Sandbox](#sandbox-projects)    |      Identity Spec      |
|             [SPIRE](#spire)             |    [Sandbox](#sandbox-projects)    |        Identity         |
|      [Telepresence](#telepresence)      |    [Sandbox](#sandbox-projects)    |         Tooling         |
|              [TiKV](#tikv)              |    [Sandbox](#sandbox-projects)    |     Key/Value Store     |
|   [Virtual Kubelet](#virtual-kubelet)   |    [Sandbox](#sandbox-projects)    |        Nodeless         |

---

Graduated Projects
------------------

### containerd

*"containerd is an industry-standard core container runtime with an emphasis on simplicity, robustness and portability.
It is available as a daemon for Linux and Windows, which can manage the complete container lifecycle of its host
system: image transfer and storage, container execution and supervision, low-level storage and network attachments,
etc.."* - [About containerd - containerd.io][about-containerd]

- **Project Repository:** https://github.com/containerd/containerd
- **Contributor Guide:** [containerd/containerd/contributing][containerd-contributor-guide]
- **Chat:** Slack: `#containerd` in [dockercommunity.slack.com][containerd-chat]
- **Developer Mailing List/Forum:** None
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[container-overview]: https://containerd.io/#about-containerd
[containerd-contributor-guide]: https://github.com/containerd/containerd/blob/master/CONTRIBUTING.md
[containerd-chat]: https://join.slack.com/t/dockercommunity/shared_invite/enQtNDY4MDc1Mzc0MzIwLTgxZDBlMmM4ZGEyNDc1N2FkMzlhODJkYmE1YTVkYjM1MDE3ZjAwZjBkOGFlOTJkZjRmZGYzNjYyY2M3ZTUxYzQ

### CoreDNS

*"CoreDNS is a DNS server. It is written in Go. It can be used in a multitude of environments because of its
flexibility. CoreDNS is licensed under the Apache License Version 2, and completely open source."* -
[What is it? - coredns.io][coredns-overview]

- **Project Repository:** https://github.com/coredns/coredns
- **Contributor Guide:** [coredns/coredns/contributing][coredns-contributor-guide]
- **Chat:** Slack: `#coredns` in [slack.cncf.io][coredns-chat]
- **Developer Mailing List/Forum:** [Coredns-Discuss Mailing List][coredns-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** None

[coredns-overview]: https://coredns.io/
[coredns-contributor-guide]: https://github.com/coredns/coredns/blob/master/CONTRIBUTING.md
[coredns-chat]: https://cloud-native.slack.com/messages/C4DF7FP71/
[coredns-dev-list]: https://groups.google.com/forum/#!forum/coredns-discuss

### Envoy

*"Originally built at Lyft, Envoy is a high performance C++ distributed proxy designed for single services and
applications, as well as a communication bus and “universal data plane” designed for large microservice “service mesh”
architectures. Built on the learnings of solutions such as NGINX, HAProxy, hardware load balancers, and cloud load
balancers, Envoy runs alongside every application and abstracts the network by providing common features in a
platform-agnostic manner. When all service traffic in an infrastructure flows via an Envoy mesh, it becomes easy to
visualize problem areas via consistent observability, tune overall performance, and add substrate features in a single
place."* - [Why Envoy? - envoyproxy.io][envoy-overview]

- **Project Repository:** https://github.com/envoyproxy/envoy
- **Contributor Guide:** [envoyproxy/envoy/contributing][envoy-contributor-guide]
- **Chat:** Slack: [envoyslack.cncf.io][envoy-chat]
- **Developer Mailing List/Forum:** [Envoy-Dev Mailing List][envoy-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[envoy-overview]: https://www.envoyproxy.io/#why-envoy
[envoy-contributor-guide]: https://github.com/envoyproxy/envoy/blob/master/CONTRIBUTING.md
[envoy-chat]: https://envoyslack.cncf.io/
[envoy-dev-list]: https://groups.google.com/forum/#!forum/envoy-dev

### Kubernetes

*"Kubernetes is a portable, extensible open-source platform for managing containerized workloads and services, that
facilitates both declarative configuration and automation. It has a large, rapidly growing ecosystem. Kubernetes
services, support, and tools are widely available.*

*Google open-sourced the Kubernetes project in 2014. Kubernetes builds upon a [decade and a half of experience that
Google has with running production workloads at scale][kubernetes-borg-paper], combined with best-of-breed ideas and
practices from the community."* - [What is Kubernetes? - kubernetes.io][kubernetes-overview]

- **Project Repository:** https://github.com/kubernetes/kubernetes
- **Contributor Guide:** [kubernetes/community/contributors/guide][kubernetes-contributor-guide]
- **Chat:** Slack: [slack.k8s.io][kubernetes-chat]
- **Developer List/Forum:** [Kubernetes-dev Mailing List][kubernetes-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [CNCF CLA][CNCF-CLA]

[kubernetes-borg-paper]: https://research.google.com/pubs/pub43438.html
[kubernetes-overview]: https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/
[kubernetes-contributor-guide]: https://github.com/kubernetes/community/tree/master/contributors/guide
[kubernetes-chat]: https://slack.k8s.io/
[kubernetes-dev-list]: https://groups.google.com/forum/#!forum/kubernetes-dev

### Prometheus

*"Prometheus is an open-source systems monitoring and alerting toolkit originally built at SoundCloud. Since its
inception in 2012, many companies and organizations have adopted Prometheus, and the project has a very active
developer and user community. It is now a standalone open source project and maintained independently of any company.
To emphasize this, and to clarify the project's governance structure, Prometheus joined the Cloud Native Computing
Foundation in 2016 as the second hosted project, after Kubernetes."* -
[Introduction to Prometheus - prometheus.io][prometheus-overview]

- **Project Repository:** https://github.com/prometheus/prometheus
- **Contributor Guide:** [prometheus/contributing][prometheus-contributor-guide]
- **Chat:** IRC: `#prometheus` on [freenode][freenode] (join via [Riot][prometheus-chat]\)
- **Developer Mailing List/Forum:** [Prometheus-Developers Mailing List][prometheus-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[prometheus-overview]: https://prometheus.io/docs/introduction/overview/
[prometheus-contributor-guide]: https://github.com/prometheus/prometheus/blob/master/CONTRIBUTING.md
[prometheus-chat]: https://riot.im/app/#/room/#freenode_#prometheus:matrix.org
[prometheus-dev-list]: https://groups.google.com/forum/#!forum/prometheus-developers

---

Incubating Projects
-------------------

### CNI

*"CNI (Container Network Interface), a Cloud Native Computing Foundation project, consists of a specification and
libraries for writing plugins to configure network interfaces in Linux containers, along with a number of supported
plugins. CNI concerns itself only with network connectivity of containers and removing allocated resources when the
container is deleted. Because of this focus, CNI has a wide range of support and the specification is simple to
implement."* - [What is CNI?- CNI Readme][cni-overview]

- **Project Repository:** https://github.com/containernetworking/cni
- **Contributor Guide:** [containernetworking/cni/contributing][cni-contributor-guide]
- **Chat:** Slack: [containernetworking.slack.com][cni-chat]
- **Developer Mailing List/Forum:** [CNI-dev Mailing List][cni-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[cni-overview]: https://github.com/containernetworking/cni/blob/master/README.md#what-is-cni
[cni-contributor-guide]: https://github.com/containernetworking/cni/blob/master/CONTRIBUTING.md
[cni-chat]: https://containernetworking.slack.com/
[cni-dev-list]: https://groups.google.com/forum/#!forum/cni-dev

### Fluentd

*"Fluentd is an open source data collector for building the unified logging layer. Once installed on a server, it runs
in the background to collect, parse, transform, analyze and store various types of data."* -
[What is Fluentd? - fluentd.org faq][fluentd-overview]

- **Project Repository:** https://github.com/fluent/fluentd
- **Contributor Guide:** [fluent/fluentd/contributing][fluentd-contributor-guide]
- **Chat:** Slack: [slack.fluentd.org][fluentd-chat]
- **Developer Mailing List/Forum:** [Fluentd Mailing List][fluentd-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[fluentd-overview]: https://www.fluentd.org/faqs
[fluentd-contributor-guide]: https://github.com/fluent/fluentd/blob/master/CONTRIBUTING.md
[fluentd-chat]: https://slack.fluentd.org/
[fluentd-dev-list]: https://groups.google.com/group/fluentd

### etcd

*"etcd is a distributed reliable key-value store for the most critical data of a distributed system, with a focus on being:*
* *Simple: well-defined, user-facing API (gRPC)*
* *Secure: automatic TLS with optional client cert authentication*
* *Fast: benchmarked 10,000 writes/sec*
* *Reliable: properly distributed using Raft"*
\- [etcd readme][etcd-overview]

- **Project Repository:** https://github.com/etcd-io/etcd
- **Contributor Guide:** [etcd-io/etcd/contributing][etcd-contributor-guide]
- **Chat:** `#etcd` on [freenode][freenode] (join via [Riot][etcd-chat]\)
- **Developer List/Forum:** [etcd-dev Mailing List][etcd-dev-list]
- **License:**  [Apache 2.0][apache-license]

[etcd-overview]: https://github.com/etcd-io/etcd#etcd
[etcd-contributor-guide]: https://github.com/etcd-io/etcd/blob/master/CONTRIBUTING.md
[etcd-dev-list]: https://groups.google.com/forum/?hl=en#!forum/etcd-dev
[etcd-chat]:  https://riot.im/app/#/room/#freenode_#etcd:matrix.org

### gRPC

*"gRPC is a modern open source high performance RPC framework that can run in any environment. It can efficiently
connect services in and across data centers with pluggable support for load balancing, tracing, health checking and
authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and
browsers to backend services."* - [About - grpc.io][grpc-overview]

- **Project Repository:** https://github.com/grpc/grpc
- **Contributor Guide:** [grpc/grpc/contributing][grpc-contributor-guide]
- **Chat:** Gitter: [gitter.im/grpc/grpc][grpc-chat]
- **Developer List/Forum:** [gRPC-io Mailing List][grpc-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [CNCF CLA][CNCF-CLA]

[grpc-overview]: https://grpc.io/about/
[grpc-contributor-guide]: https://github.com/grpc/grpc/blob/master/CONTRIBUTING.md
[grpc-chat]: https://gitter.im/grpc/grpc
[grpc-dev-list]: https://groups.google.com/forum/#!forum/grpc-io

### Harbor

*"Project Harbor is an an open source trusted cloud native registry project that stores, signs, and scans content.
Harbor extends the open source Docker Distribution by adding the functionalities usually required by users such as
security, identity and management. Having a registry closer to the build and run environment can improve the image
transfer efficiency. Harbor supports replication of images between registries, and also offers advanced security
features such as user management, access control and activity auditing."* - [Harbor Readme][harbor-overview]

- **Project Repository:** https://github.com/goharbor/harbor
- **Contributor Guide:** [vmware/harbor/contributing][harbor-contributor-guide]
- **Chat:** Slack: `#harbor-dev` in [slack.cncf.io][harbor-chat]
- **Developer Mailing List/Forum:** [Harbor-Dev Mailing List][harbor-dev-list]
- **License:** [Apache 2.0][apache-license]

[harbor-overview]: https://github.com/goharbor/harbor
[harbor-contributor-guide]: https://github.com/goharbor/harbor/blob/master/CONTRIBUTING.md
[harbor-chat]: https://slack.cncf.io
[harbor-dev-list]: https://groups.google.com/forum/#!forum/harbor-dev

### Helm

*"Helm helps you manage Kubernetes applications — Helm Charts helps you define, install, and upgrade even the most
complex Kubernetes application.* *Charts are easy to create, version, share, and publish — so start using Helm and
stop the copy-and-paste madness.* *The latest version of Helm is maintained by the CNCF - in collaboration with
Microsoft, Google, Bitnami and the Helm contributor community."* - [What is Helm? - helm.sh][helm-overview]

- **Project Repository:** https://github.com/helm/helm
- **Contributor Guide:** [helm/helm/contributing][helm-contributor-guide]
- **Chat:** Slack `#helm-dev` in [slack.k8s.io][helm-chat]
- **Developer List/Forum:** [CNCF-Helm Mailing List][helm-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [CNCF CLA][CNCF-CLA]

[helm-overview]: https://www.helm.sh/
[helm-contributor-guide]: https://github.com/helm/helm/blob/master/CONTRIBUTING.md
[helm-chat]: https://slack.k8s.io/
[helm-dev-list]: https://lists.cncf.io/g/cncf-helm

### Jaeger

*"Jaeger, inspired by Dapper and OpenZipkin, is a distributed tracing system released as open source by Uber
Technologies. It is used for monitoring and troubleshooting microservices-based distributed systems."* -
[About - jaegertracing.io][jaeger-overview]

- **Project Repository:** https://github.com/jaegertracing/jaeger
- **Contributor Guide:** [jaegertracing/jaeger/contributing][jaeger-contributor-guide]
- **Chat:** Gitter: [gitter.im/jaegertracing/Lobby][jaeger-chat]
- **Developer Mailing List/Forum:** [Jaeger-Tracing Mailing List][jaeger-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[jaeger-overview]: https://www.jaegertracing.io/docs/#about
[jaeger-contributor-guide]: https://github.com/jaegertracing/jaeger/blob/master/CONTRIBUTING.md
[jaeger-chat]: https://gitter.im/jaegertracing/Lobby
[jaeger-dev-list]: https://groups.google.com/forum/#!forum/jaeger-tracing

### Linkerd

*"Linkerd is a transparent service mesh, designed to make modern applications safe and sane by transparently adding
service discovery, load balancing, failure handling, instrumentation, and routing to all inter-service
communication."* - [Linkerd Readme][linkerd-overview]

- **Project Repository:** https://github.com/linkerd/linkerd
- **Contributor Guide:** [linkerd/linkerd/contributing][linkerd-contributor-guide]
- **Chat:** Slack: [slack.linkerd.io][linkerd-chat]
- **Developer Mailing List/Forum:** [Linkerd Forum][linkerd-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[linkerd-overview]: https://github.com/linkerd/linkerd
[linkerd-contributor-guide]: https://github.com/linkerd/linkerd/blob/master/CONTRIBUTING.md
[linkerd-chat]: https://slack.linkerd.io/
[linkerd-dev-list]: https://discourse.linkerd.io/

### NATS

*"NATS is an open source, lightweight, high-performance cloud native infrastructure messaging system. It implements a
highly scalable and elegant publish-subscribe (pub/sub) distribution model. The performant nature of NATS make it an
ideal base for building modern, reliable, scalable cloud native distributed systems."* -
[What is NATS? - nats.io][nats-overview]

- **Project Repository:** https://github.com/nats-io
- **Contributor Guide:** [nats.io/documentation/contributing/][nats-contributor-guide]
- **Chat:** Slack: [natsio.slack.com][nats-chat]
- **Developer Mailing List/Forum:** [natsio Mailing List][nats-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** None

[nats-overview]: https://nats.io/documentation/faq/
[nats-contributor-guide]: https://nats.io/documentation/contributing/
[nats-chat]: https://natsio.slack.com/join/shared_invite/enQtMzE2NDkxNDI2NTE1LTc5ZDEzYTkwYWZkYWQ5YjY1MzBjMWZmYzA5OGQxMzlkMGQzMjYxNGM3MWYxMjNiYmNjNzIwMTVjMWE2ZDgxZGM
[nats-dev-list]: https://groups.google.com/forum/#!forum/natsio

### Notary

*"Notary aims to make the internet more secure by making it easy for people to publish and verify content. We often
rely on TLS to secure our communications with a web server, which is inherently flawed, as any compromise of the server
enables malicious content to be substituted for the legitimate content."* [Overview - Notary Readme][notary-overview]

- **Project Repository:** https://github.com/theupdateframework/notary
- **Contributor Guide:** [theupdateframework/notary/contributing][notary-contributor-guide]
- **Chat:** None
- **Developer Mailing List/Forum:** [theupateframework Mailing List][notary-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[notary-overview]: https://github.com/theupdateframework/notary/blob/master/README.md#overview
[notary-contributor-guide]: https://github.com/theupdateframework/notary/blob/master/CONTRIBUTING.md
[notary-dev-list]: https://groups.google.com/forum/?fromgroups#!forum/theupdateframework

### OpenTracing

Vendor-neutral APIs and instrumentation for distributed tracing.

- **Project Repository:** https://github.com/opentracing
- **Contributor Guide:** [opentracing-contrib/meta][opentracing-contributor-guide]
- **Chat:** Gitter: [gitter.im/opentracing/public][opentracing-chat]
- **Developer Mailing List/Forum:** [OpenTracing Mailing List][opentracing-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** None

[opentracing-contributor-guide]: https://github.com/opentracing-contrib/meta#opentracing-contributions
[opentracing-chat]: https://gitter.im/opentracing/public
[opentracing-dev-list]: https://groups.google.com/forum/#!forum/opentracing

### rkt

rkt is a pod-native container engine for Linux. It is composable, secure, and built on standards.

- **Project Repository:** https://github.com/rkt/rkt
- **Contributor Guide:** [rkt/rkt/contributing][rkt-contributor-guide]
- **Chat:** `#rkt-dev` on [freenode][freenode] (join via [Riot][rkt-chat]\)
- **Developer Mailing List/Forum:** [rkt-dev Mailing List][rkt-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[rkt-contributor-guide]: https://github.com/rkt/rkt/blob/master/CONTRIBUTING.md
[rkt-chat]: https://riot.im/app/#/room/#freenode_#rkt-dev:matrix.org
[rkt-dev-list]: https://groups.google.com/forum/#!forum/rkt-dev

### Rook

*"Rook is an open source cloud-native storage orchestrator for Kubernetes, providing the platform, framework, and
support for a diverse set of storage solutions to natively integrate with cloud-native environments."* -
[What is Rook? - Rook Readme][rook-overview]

- **Project Repository:** https://github.com/rook/rook
- **Contributor Guide:** [rook/rook/contributing][rook-contributor-guide]
- **Chat:** Slack: [rook-io.slack.com][rook-chat]
- **Developer Mailing List/Forum:** [Rook-Dev Mailing List][rook-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[rook-overview]: https://github.com/rook/rook/blob/master/README.md
[rook-contributor-guide]: https://github.com/rook/rook/blob/master/CONTRIBUTING.md
[rook-chat]: https://rook-io.slack.com/
[rook-dev-list]: https://groups.google.com/forum/#!forum/rook-dev

### TUF

*"The Update Framework (TUF) helps developers maintain the security of a software update system, even against attackers
that compromise the repository or signing keys. TUF provides a flexible framework and specification that developers can
adopt into any software update system."* - [TUF Readme][tuf-overview]

- **Project Repository:** https://github.com/theupdateframework/specification
- **Contributor Guide:** [theupdateframework/tuf/contributors][tuf-contributor-guide]
- **Chat:** None
- **Developer Mailing List/Forum:** [TUF Mailing List][tuf-dev-list]
- **License:** Dual Licensed [Apache 2.0][apache-license] / [MIT][mit-license]
- **Legal Requirements:** [DCO][DCO]

[tuf-overview]: https://github.com/theupdateframework/tuf/blob/develop/README.md
[tuf-contributor-guide]: https://github.com/theupdateframework/tuf/blob/develop/docs/CONTRIBUTORS.rst
[tuf-dev-list]: https://groups.google.com/forum/?fromgroups#!forum/theupdateframework

### Vitess

*"Vitess is a database solution for deploying, scaling and managing large clusters of MySQL instances. It's architected
to run as effectively in a public or private cloud architecture as it does on dedicated hardware. It combines and
extends many important MySQL features with the scalability of a NoSQL database."* -
[Overview - vitess.io][vitess-overview]

- **Project Repository:** https://github.com/vitessio/vitess
- **Contributor Guide:** [vitessio/vitess/contributing][vitess-contributor-guide]
- **Chat:** Slack: [vitess.slack.com][vitess-chat]
- **Developer Mailing List/Forum:** [Vitess Mailing List][vitess-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [Google Corporate CLA][vitess-corporate-cla] / [Google Individual CLA][vitess-individual-cla]

[vitess-overview]: https://vitess.io/overview/
[vitess-contributor-guide]: https://github.com/vitessio/vitess/blob/master/CONTRIBUTING.md
[vitess-chat]: https://bit.ly/vitess-slack
[vitess-dev-list]: https://groups.google.com/forum/#!forum/vitess
[vitess-corporate-cla]: https://cla.developers.google.com/about/google-corporate
[vitess-individual-cla]: https://cla.developers.google.com/about/google-individual

---

Sandbox Projects
----------------

### Buildpacks

*"Buildpacks provide a higher-level abstraction for building apps compared to
Dockerfiles."* - [What Are Buildpacks? - buildpacks.io][buildpacks-overview]

- **Project Repository:** https://github.com/buildpack
- **Chat:** Slack: [slack.buildpacks.io][buildpacks-chat]
- **License:**  [Apache 2.0][apache-license]

[buildpacks-overview]: https://buildpacks.io/
[buildpacks-chat]: https://slack.buildpacks.io/

### CloudEvents

CloudEvents Specification

- **Project Repository:** https://github.com/cloudevents/spec
- **Contributor Guide:** [cloudevents/spec/contributing][cloudevents-contributor-guide]
- **Chat:** Slack: `#cloudevents` in [slack.cncf.io][cloudevents-chat]
- **Developer Mailing List/Forum:** [CNCF-wg-Serverless Mailing List][cloudevents-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[cloudevents-contributor-guide]: https://github.com/cloudevents/spec/blob/master/CONTRIBUTING.md
[cloudevents-chat]: https://cloud-native.slack.com/messages/C9DB5ABAA/
[cloudevents-dev-list]: https://lists.cncf.io/g/cncf-wg-serverless

### Cortex

*"Cortex provides horizontally scalable, multi-tenant, long term storage for Prometheus metrics when used as a remote
write destination, and a horizontally scalable, Prometheus-compatible query API."* - [Cortex Readme][cortex-overview]

- **Project Repository:** https://github.com/cortexproject/cortex
- **Contributor Guide:** [cortextproject/cortex/readme][cortex-contributor-guide]
- **Chat:** Slack: `#cortext` in [slack.cncf.io][cortex-chat]
- **License:**  [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[cortex-overview]: https://github.com/cortexproject/cortex#open-source-horizontally-scalable-multi-tenant-prometheus-as-a-service
[cortex-contributor-guide]: https://github.com/cortexproject/cortex#for-developers
[cortex-chat]: https://cloud-native.slack.com/messages/cortex/

### Dragonfly

*"Dragonfly is an intelligent P2P-based image and file distribution tool. It aims to improve the efficiency and success
  rate of file transferring, and maximize the usage of network bandwidth, especially for the distribution of larget
  amounts of data, such as application distribution, cache distribution, log distribution, and image
  distribution."* - [Overview - d7y.io][dragonfly-overview]

- **Project Repository:** https://github.com/dragonflyoss/dragonfly
- **Contributor Guide:** [dragonflyoss/dragonfly/CONTRIBUTING][dragonfly-contributor-guide]
- **Chat:** [gitter.im/alibaba/Dragonfly][dragonfly-chat]
- **License:**  [Apache 2.0][apache-license]

[dragonfly-overview]: https://d7y.io/
[dragonfly-contributor-guide]: https://github.com/dragonflyoss/Dragonfly/blob/master/CONTRIBUTING.md
[dragonfly-chat]: https://gitter.im/alibaba/Dragonfly

### Falco

*"Falco is a behavioral activity monitor designed to detect anomalous activity in your applications. Powered by
sysdig’s system call capture infrastructure, Falco lets you continuously monitor and detect container, application,
host, and network activity... all in one place, from one source of data, with one set of rules."* -
[Overview - Falco Readme][falco-overview]

- **Project Repository:** https://github.com/falcosecurity/falco
- **Chat:** Slack: `#falco` in [slack.sysdig.com][falco-chat]
- **License:**  [Apache 2.0][apache-license]
- **Legal Requirements:** [Falco CLA][falco-legal]

[falco-overview]: https://github.com/falcosecurity/falco#overview
[falco-chat]: https://slack.sysdig.com/
[falco-legal]: https://github.com/falcosecurity/falco#contributor-license-agreements

### Open Policy Agent

*"OPA is a lightweight general-purpose policy engine that can be co-located with your service. You can integrate OPA
as a sidecar, host-level daemon, or library."* - [What is OPA? - openpolicyagent.org][opa-overview]

- **Project Repository:** https://github.com/open-policy-agent/opa
- **Contributor Guide:** [open-policy-agent/opa/contributing][opa-contributor-guide]
- **Chat:** Slack: [slack.openpolicyagent.org][opa-chat]
- **Developer Mailing List/Forum:** None
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** None

[opa-overview]: https://www.openpolicyagent.org/docs/#what-is-opa
[opa-contributor-guide]: https://github.com/open-policy-agent/opa/blob/master/CONTRIBUTING.md
[opa-chat]: https://slack.openpolicyagent.org/

### OpenMetrics

*"An effort to create an open standard for transmitting metrics at scale, with support for both text representation and
Protocol Buffers."* - [openmetrics.io][openmetrics-overview]

- **Project Repository:** https://github.com/OpenObservability/OpenMetrics
- **Contributor Guide:** [TBD]
- **Chat:** Slack: `#openmetrics` in [slack.cncf.io][openmetrics-chat]
- **Developer Mailing List/Forum:** [OpemMetrics Mailing List][openmetrics-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** None

[openmetrics-overview]: https://openmetrics.io/
[openmetrics-chat]: https://cloud-native.slack.com/messages/CC6CPDEJV/
[openmetrics-dev-list]: https://groups.google.com/forum/m/#!forum/openmetrics

### SPIFFE

*"SPIFFE (Secure Production Identity Framework For Everyone) provides a secure identity, in the form of a specially
crafted X.509 certificate, to every workload in a modern production environment. SPIFFE removes the need for
application-level authentication and complex network-level ACL configuration."* -
[What is SPIFFE? - spiffe.io][spiffe-overview]

- **Project Repository:** https://github.com/spiffe/spiffe
- **Contributor Guide:** [spiffe/spiffe/contributing][spiffe-contributor-guide]
- **Chat:** Slack: [slack.spiffe.io][spiffe-chat]
- **Developer Mailing List/Forum:** [SPIFFE Dev Discussion Mailing List][spiffe-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[spiffe-overview]: https://spiffe.io/
[spiffe-contributor-guide]: https://github.com/spiffe/spiffe/blob/master/CONTRIBUTING.md
[spiffe-chat]: https://slack.spiffe.io/
[spiffe-dev-list]: https://groups.google.com/a/spiffe.io/forum/#!forum/dev-discussion

### SPIRE

*"SPIRE (the SPIFFE Runtime Environment) is a tool-chain for establishing trust between software systems across a wide
variety of hosting platforms. Concretely, SPIRE exposes the SPIFFE Workload API, which can attest running software
systems and issue SPIFFE IDs and SVIDs to them."* - [Spire Readme][spire-overview]

- **Project Repository:** https://github.com/spiffe/spire
- **Contributor Guide:** [spiffe/spire/contributing][spire-contributor-guide]
- **Chat:** Slack: [slack.spiffe.io][spire-chat]
- **Developer Mailing List/Forum:** [SPIFFE Dev Discussion Mailing List][spire-dev-list]
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[spire-overview]: https://github.com/spiffe/spire/blob/master/README.md
[spire-contributor-guide]: https://github.com/spiffe/spire/blob/master/CONTRIBUTING.md
[spire-chat]: https://slack.spiffe.io/
[spire-dev-list]: https://groups.google.com/a/spiffe.io/forum/#!forum/dev-discussion

### Telepresence

*"Telepresence is an open source tool that lets you run a single service locally, while connecting that service to a
remote Kubernetes cluster."* - [Overview - telepresene.io][telepresence-overview]

- **Project Repository:** https://github.com/telepresenceio/telepresence
- **Contributor Guide:** [telepresenceio/telepresence/docs/reference/developing][telepresence-contributor-guide]
- **Chat:** Gitter: [gitter.im/datawire/telepresence][telepresence-chat]
- **Developer Mailing List/Forum:** None
- **License:** [Apache 2.0][apache-license]
- **Legal Requirements:** None

[telepresence-overview]: https://www.telepresence.io/discussion/overview
[telepresence-contributor-guide]: https://github.com/telepresenceio/telepresence/blob/master/docs/reference/developing.md
[telepresence-chat]: https://gitter.im/datawire/telepresence

### TiKV

*"TiKV ("Ti" stands for Titanium) is a distributed transactional key-value database, originally created to complement
TiDB, a distributed HTAP database compatible with the MySQL protocol. TiKV is built in Rust and powered by Raft, and
was inspired by the design of Google Spanner and HBase, but without dependency on any specific distributed file
system."* - [TiKV Readme][tikv-overview]

- **Project Repository:** https://github.com/tikv/tikv
- **Contributor Guide:** [tkiv/tkiv/contributing][tikv-contributor-guide]
- **License:**  [Apache 2.0][apache-license]
- **Legal Requirements:** [DCO][DCO]

[tikv-overview]: https://github.com/tikv/tikv#tikv
[tikv-contributor-guide]: https://github.com/tikv/tikv/blob/master/CONTRIBUTING.md

### Virtual Kubelet

*"Virtual Kubelet is an open source Kubernetes kubelet implementation that masquerades as a kubelet for the purposes of
connecting Kubernetes to other APIs. This allows the nodes to be backed by other services like ACI, AWS Fargate,
Hyper.sh, IoT Edge etc. The primary scenario for VK is enabling the extension of the Kubernetes API into serverless
container platforms like ACI, Fargate, and Hyper.sh, though we are open to others. However, it should be noted that
VK is explicitly not intended to be an alternative to Kubernetes federation."* -
[Virtual Kubelet Readme][virtual-kubelet-overview]

- **Project Repository:** https://github.com/virtual-kubelet/virtual-kubelet
- **Contributor Guide:** [virtual-kubelet/virtual-kubelet/contributing][virtual-kubelet-contributor-guide]
- **License:**  [Apache 2.0](apache-license)
- **Legal Requirements:** [CNCF CLA][CNCF-CLA]

[virtual-kubelet-overview]: https://github.com/virtual-kubelet/virtual-kubelet
[virtual-kubelet-contributor-guide]: https://github.com/virtual-kubelet/virtual-kubelet/blob/master/CONTRIBUTING.md
