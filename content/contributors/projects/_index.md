---
title: CNCF Projects
---
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





##### PROJECT TEMPLATE #####

-->

<!-- Frequently Used Links -->

All projects of the Cloud Native Computing Foundation are classified with one of three stages of maturity:

-	[Graduated](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#graduation-stage)
-	[Incubated](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#incubating-stage)
-	[Sandboxed](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc#sandbox-stage)

[CNCF Graduation Criteria](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc) are documented in the CNCF TOC repo. The document describes the maturity stages of the projects.

---

Graduated Projects
------------------

### Kubernetes

*"Kubernetes is a portable, extensible open-source platform for managing containerized workloads and services, that facilitates both declarative configuration and automation. It has a large, rapidly growing ecosystem. Kubernetes services, support, and tools are widely available.*

*Google open-sourced the Kubernetes project in 2014. Kubernetes builds upon a [decade and a half of experience that Google has with running production workloads at scale](https://research.google.com/pubs/pub43438.html), combined with best-of-breed ideas and practices from the community."* - [What is Kubernetes? - kubernetes.io](https://kubernetes.io/docs/concepts/overview/what-is-kubernetes/)

-	**Project Repository:** https://github.com/kubernetes/kubernetes
-	**Contributor Guide:** [kubernetes/community/contributors/guide](https://github.com/kubernetes/community/tree/master/contributors/guide)
-	**Chat:** Slack: [slack.k8s.io](https://slack.k8s.io/)
-	**Developer List/Forum:** [Kubernetes-dev Mailing List](https://groups.google.com/forum/#!forum/kubernetes-dev)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [CNCF CLA](https://github.com/cncf/cla)

### Prometheus

*"Prometheus is an open-source systems monitoring and alerting toolkit originally built at SoundCloud. Since its inception in 2012, many companies and organizations have adopted Prometheus, and the project has a very active developer and user community. It is now a standalone open source project and maintained independently of any company. To emphasize this, and to clarify the project's governance structure, Prometheus joined the Cloud Native Computing Foundation in 2016 as the second hosted project, after Kubernetes."* - [Introduction to Prometheus - prometheus.io](https://prometheus.io/docs/introduction/overview/)

-	**Project Repository:** https://github.com/prometheus/prometheus
-	**Contributor Guide:** [prometheus/contributing](https://github.com/prometheus/prometheus/blob/master/CONTRIBUTING.md)
-	**Chat:** IRC: `#prometheus` on [freenode](https://freenode.net) (join via [Riot](https://riot.im/app/#/room/#freenode_#prometheus:matrix.org)\)
-	**Developer Mailing List/Forum:** [Prometheus-Developers Mailing List](https://groups.google.com/forum/#!forum/prometheus-developers)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Envoy

*"Originally built at Lyft, Envoy is a high performance C++ distributed proxy designed for single services and applications, as well as a communication bus and “universal data plane” designed for large microservice “service mesh” architectures. Built on the learnings of solutions such as NGINX, HAProxy, hardware load balancers, and cloud load balancers, Envoy runs alongside every application and abstracts the network by providing common features in a platform-agnostic manner. When all service traffic in an infrastructure flows via an Envoy mesh, it becomes easy to visualize problem areas via consistent observability, tune overall performance, and add substrate features in a single place."* - [Why Envoy? - envoyproxy.io](https://www.envoyproxy.io/#why-envoy)

-	**Project Repository:** https://github.com/envoyproxy/envoy
-	**Contributor Guide:** [envoyproxy/envoy/contributing](https://github.com/envoyproxy/envoy/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [envoyslack.cncf.io](https://envoyslack.cncf.io/)
-	**Developer Mailing List/Forum:** [Envoy-Dev Mailing List](https://groups.google.com/forum/#!forum/envoy-dev)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### CoreDNS

*"CoreDNS is a DNS server. It is written in Go. It can be used in a multitude of environments because of its flexibility. CoreDNS is licensed under the Apache License Version 2, and completely open source."* - [What is it? - coredns.io](https://coredns.io/)

-	**Project Repository:** https://github.com/coredns/coredns
-	**Contributor Guide:** [coredns/coredns/contributing](https://github.com/coredns/coredns/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: `#coredns` in [slack.cncf.io](https://cloud-native.slack.com/messages/C4DF7FP71/)
-	**Developer Mailing List/Forum:** [Coredns-Discuss Mailing List](https://groups.google.com/forum/#!forum/coredns-discuss)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** None

### containerd

*"containerd is an industry-standard core container runtime with an emphasis on simplicity, robustness and portability. It is available as a daemon for Linux and Windows, which can manage the complete container lifecycle of its host system: image transfer and storage, container execution and supervision, low-level storage and network attachments, etc.."* - [About containerd - containerd.io][about-containerd]

-	**Project Repository:** https://github.com/containerd/containerd
-	**Contributor Guide:** [containerd/containerd/contributing](https://github.com/containerd/containerd/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: `#containerd` in [dockercommunity.slack.com](https://join.slack.com/t/dockercommunity/shared_invite/enQtNDY4MDc1Mzc0MzIwLTgxZDBlMmM4ZGEyNDc1N2FkMzlhODJkYmE1YTVkYjM1MDE3ZjAwZjBkOGFlOTJkZjRmZGYzNjYyY2M3ZTUxYzQ)
-	**Developer Mailing List/Forum:** None
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Fluentd

*"Fluentd is an open source data collector for building the unified logging layer. Once installed on a server, it runs in the background to collect, parse, transform, analyze and store various types of data."* - [What is Fluentd? - fluentd.org faq](https://www.fluentd.org/faqs)

-	**Project Repository:** https://github.com/fluent/fluentd
-	**Contributor Guide:** [fluent/fluentd/contributing](https://github.com/fluent/fluentd/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [slack.fluentd.org](https://slack.fluentd.org/)
-	**Developer Mailing List/Forum:** [Fluentd Mailing List](https://groups.google.com/group/fluentd)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Jaeger

*"Jaeger, inspired by Dapper and OpenZipkin, is a distributed tracing system released as open source by Uber Technologies. It is used for monitoring and troubleshooting microservices-based distributed systems."* - [About - jaegertracing.io](https://www.jaegertracing.io/docs/#about)

-	**Project Repository:** https://github.com/jaegertracing/jaeger
-	**Contributor Guide:** [jaegertracing/jaeger/contributing](https://github.com/jaegertracing/jaeger/blob/master/CONTRIBUTING.md)
-	**Chat:** Gitter: [gitter.im/jaegertracing/Lobby](https://gitter.im/jaegertracing/Lobby)
-	**Developer Mailing List/Forum:** [Jaeger-Tracing Mailing List](https://groups.google.com/forum/#!forum/jaeger-tracing)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Vitess

*"Vitess is a database solution for deploying, scaling and managing large clusters of MySQL instances. It's architected to run as effectively in a public or private cloud architecture as it does on dedicated hardware. It combines and extends many important MySQL features with the scalability of a NoSQL database."* - [Overview - vitess.io](https://vitess.io/overview/)

-	**Project Repository:** https://github.com/vitessio/vitess
-	**Contributor Guide:** [vitessio/vitess/contributing](https://github.com/vitessio/vitess/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [vitess.slack.com](https://bit.ly/vitess-slack)
-	**Developer Mailing List/Forum:** [Vitess Mailing List](https://groups.google.com/forum/#!forum/vitess)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [Google Corporate CLA](https://cla.developers.google.com/about/google-corporate) / [Google Individual CLA](https://cla.developers.google.com/about/google-individual)

### TUF

*"The Update Framework (TUF) helps developers maintain the security of a software update system, even against attackers that compromise the repository or signing keys. TUF provides a flexible framework and specification that developers can adopt into any software update system."* - [TUF Readme](https://github.com/theupdateframework/tuf/blob/develop/README.md)

-	**Project Repository:** https://github.com/theupdateframework/specification
-	**Contributor Guide:** [theupdateframework/tuf/contributors](https://github.com/theupdateframework/tuf/blob/develop/docs/CONTRIBUTORS.rst)
-	**Chat:** None
-	**Developer Mailing List/Forum:** [TUF Mailing List](https://groups.google.com/forum/?fromgroups#!forum/theupdateframework)
-	**License:** Dual Licensed [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/) / [MIT](https://choosealicense.com/licenses/mit/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Helm

*"Helm helps you manage Kubernetes applications — Helm Charts helps you define, install, and upgrade even the most complex Kubernetes application.* *Charts are easy to create, version, share, and publish — so start using Helm and stop the copy-and-paste madness.* *The latest version of Helm is maintained by the CNCF - in collaboration with Microsoft, Google, Bitnami and the Helm contributor community."* - [What is Helm? - helm.sh](https://www.helm.sh/)

-	**Project Repository:** https://github.com/helm/helm
-	**Contributor Guide:** [helm/helm/contributing](https://github.com/helm/helm/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack `#helm-dev` in [slack.k8s.io](https://slack.k8s.io/)
-	**Developer List/Forum:** [CNCF-Helm Mailing List](https://lists.cncf.io/g/cncf-helm)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Harbor

*"Project Harbor is an an open source trusted cloud native registry project that stores, signs, and scans content. Harbor extends the open source Docker Distribution by adding the functionalities usually required by users such as security, identity and management. Having a registry closer to the build and run environment can improve the image transfer efficiency. Harbor supports replication of images between registries, and also offers advanced security features such as user management, access control and activity auditing."* - [Harbor Readme](https://github.com/goharbor/harbor)

-	**Project Repository:** https://github.com/goharbor/harbor
-	**Contributor Guide:** [vmware/harbor/contributing](https://github.com/goharbor/harbor/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: `#harbor-dev` in [slack.cncf.io](https://slack.cncf.io)
-	**Developer Mailing List/Forum:** [Harbor-Dev Mailing List](https://groups.google.com/forum/#!forum/harbor-dev)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)

### Rook

*"Rook is an open source cloud-native storage orchestrator for Kubernetes, providing the platform, framework, and support for a diverse set of storage solutions to natively integrate with cloud-native environments."* - [What is Rook? - Rook Readme](https://github.com/rook/rook/blob/master/README.md)

-	**Project Repository:** https://github.com/rook/rook
-	**Contributor Guide:** [rook/rook/contributing](https://github.com/rook/rook/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [rook-io.slack.com](https://rook-io.slack.com/)
-	**Developer Mailing List/Forum:** [Rook-Dev Mailing List](https://groups.google.com/forum/#!forum/rook-dev)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### TiKV

*"TiKV ("Ti" stands for Titanium) is a distributed transactional key-value database, originally created to complement TiDB, a distributed HTAP database compatible with the MySQL protocol. TiKV is built in Rust and powered by Raft, and was inspired by the design of Google Spanner and HBase, but without dependency on any specific distributed file system."* - [TiKV Readme](https://github.com/tikv/tikv#tikv)

-	**Project Repository:** https://github.com/tikv/tikv
-	**Contributor Guide:** [tkiv/tkiv/contributing](https://github.com/tikv/tikv/blob/master/CONTRIBUTING.md)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)
---

Incubating Projects
-------------------

### OpenTracing

Vendor-neutral APIs and instrumentation for distributed tracing.

-	**Project Repository:** https://github.com/opentracing
-	**Contributor Guide:** [opentracing-contrib/meta](https://github.com/opentracing-contrib/meta#opentracing-contributions)
-	**Chat:** Gitter: [gitter.im/opentracing/public](https://gitter.im/opentracing/public)
-	**Developer Mailing List/Forum:** [OpenTracing Mailing List](https://groups.google.com/forum/#!forum/opentracing)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** None

### gRPC

*"gRPC is a modern open source high performance RPC framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services."* - [About - grpc.io](https://grpc.io/about/)

-	**Project Repository:** https://github.com/grpc/grpc
-	**Contributor Guide:** [grpc/grpc/contributing](https://github.com/grpc/grpc/blob/master/CONTRIBUTING.md)
-	**Chat:** Gitter: [gitter.im/grpc/grpc](https://gitter.im/grpc/grpc)
-	**Developer List/Forum:** [gRPC-io Mailing List](https://groups.google.com/forum/#!forum/grpc-io)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [CNCF CLA](https://github.com/cncf/cla)

### CNI

*"CNI (Container Network Interface), a Cloud Native Computing Foundation project, consists of a specification and libraries for writing plugins to configure network interfaces in Linux containers, along with a number of supported plugins. CNI concerns itself only with network connectivity of containers and removing allocated resources when the container is deleted. Because of this focus, CNI has a wide range of support and the specification is simple to implement."* - [What is CNI?- CNI Readme](https://github.com/containernetworking/cni/blob/master/README.md#what-is-cni)

-	**Project Repository:** https://github.com/containernetworking/cni
-	**Contributor Guide:** [containernetworking/cni/contributing](https://github.com/containernetworking/cni/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [containernetworking.slack.com](https://containernetworking.slack.com/)
-	**Developer Mailing List/Forum:** [CNI-dev Mailing List](https://groups.google.com/forum/#!forum/cni-dev)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Notary

*"Notary aims to make the internet more secure by making it easy for people to publish and verify content. We often rely on TLS to secure our communications with a web server, which is inherently flawed, as any compromise of the server enables malicious content to be substituted for the legitimate content."* [Overview - Notary Readme](https://github.com/theupdateframework/notary/blob/master/README.md#overview)

-	**Project Repository:** https://github.com/theupdateframework/notary
-	**Contributor Guide:** [theupdateframework/notary/contributing](https://github.com/theupdateframework/notary/blob/master/CONTRIBUTING.md)
-	**Chat:** None
-	**Developer Mailing List/Forum:** [theupateframework Mailing List](https://groups.google.com/forum/?fromgroups#!forum/theupdateframework)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### NATS

*"NATS is an open source, lightweight, high-performance cloud native infrastructure messaging system. It implements a highly scalable and elegant publish-subscribe (pub/sub) distribution model. The performant nature of NATS make it an ideal base for building modern, reliable, scalable cloud native distributed systems."* - [What is NATS? - nats.io](https://docs.nats.io/nats-concepts/intro)

-	**Project Repository:** https://github.com/nats-io
-	**Contributor Guide:** [nats.io/contributing/](https://nats.io/contributing/)
-	**Chat:** Slack: [natsio.slack.com](https://natsio.slack.com/join/shared_invite/enQtMzE2NDkxNDI2NTE1LTc5ZDEzYTkwYWZkYWQ5YjY1MzBjMWZmYzA5OGQxMzlkMGQzMjYxNGM3MWYxMjNiYmNjNzIwMTVjMWE2ZDgxZGM)
-	**Developer Mailing List/Forum:** [natsio Mailing List](https://groups.google.com/forum/#!forum/natsio)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** None

### Linkerd

*"Linkerd is a transparent service mesh, designed to make modern applications safe and sane by transparently adding service discovery, load balancing, failure handling, instrumentation, and routing to all inter-service communication."* - [Linkerd Readme](https://github.com/linkerd/linkerd)

-	**Project Repository:** https://github.com/linkerd/linkerd
-	**Contributor Guide:** [linkerd/linkerd/contributing](https://github.com/linkerd/linkerd/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [slack.linkerd.io](https://slack.linkerd.io/)
-	**Developer Mailing List/Forum:** [Linkerd Forum](https://discourse.linkerd.io/)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)


### etcd

*"etcd is a distributed reliable key-value store for the most critical data of a distributed system, with a focus on being:*

-	*Simple: well-defined, user-facing API (gRPC)*
-	*Secure: automatic TLS with optional client cert authentication*
-	*Fast: benchmarked 10,000 writes/sec*
-	*Reliable: properly distributed using Raft"*\- [etcd readme](https://github.com/etcd-io/etcd#etcd)
-	**Project Repository:** https://github.com/etcd-io/etcd
-	**Contributor Guide:** [etcd-io/etcd/contributing](https://github.com/etcd-io/etcd/blob/master/CONTRIBUTING.md)
-	**Chat:** `#etcd` on [freenode](https://freenode.net) (join via [Riot](https://riot.im/app/#/room/#freenode_#etcd:matrix.org)\)
-	**Developer List/Forum:** [etcd-dev Mailing List](https://groups.google.com/forum/?hl=en#!forum/etcd-dev)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)

### Open Policy Agent

*"OPA is a lightweight general-purpose policy engine that can be co-located with your service. You can integrate OPA as a sidecar, host-level daemon, or library."* - [What is OPA? - openpolicyagent.org](https://www.openpolicyagent.org/docs/#what-is-opa)

-	**Project Repository:** https://github.com/open-policy-agent/opa
-	**Contributor Guide:** [open-policy-agent/opa/contributing](https://github.com/open-policy-agent/opa/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [slack.openpolicyagent.org](https://slack.openpolicyagent.org/)
-	**Developer Mailing List/Forum:** None
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** None

### CRI-O

*"CRI-O is an implementation of the Kubernetes CRI (Container Runtime Interface) to enable using OCI (Open Container Initiative) compatible runtimes. It is a lightweight alternative to using Docker as the runtime for kubernetes. It allows Kubernetes to use any OCI-compliant runtime as the container runtime for running pods. Today it supports runc and Kata Containers as the container runtimes but any OCI-conformant runtime can be plugged in principle."* - [What is CRI-O? - CRI-O.org](https://cri-o.io/)

-	**Project Repository:** https://github.com/cri-o
-	**Contributor Guide:** [cri-o/cri-o/CONTRIBUTING.md](https://github.com/cri-o/cri-o/blob/master/CONTRIBUTING.md)
-	**Chat:** `#cri-o` on [freenode](https://freenode.net) (join via [Riot](https://riot.im/app/#/room/#freenode_#cri-o:matrix.org)\)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)


### CloudEvents

CloudEvents Specification

-	**Project Repository:** https://github.com/cloudevents/spec
-	**Contributor Guide:** [cloudevents/spec/contributing](https://github.com/cloudevents/spec/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: `#cloudevents` in [slack.cncf.io](https://cloud-native.slack.com/messages/C9DB5ABAA/)
-	**Developer Mailing List/Forum:** [CNCF-wg-Serverless Mailing List](https://lists.cncf.io/g/cncf-wg-serverless)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Falco

*"Falco is a behavioral activity monitor designed to detect anomalous activity in your applications. Powered by sysdig’s system call capture infrastructure, Falco lets you continuously monitor and detect container, application, host, and network activity... all in one place, from one source of data, with one set of rules."* - [Overview - Falco Readme](https://github.com/falcosecurity/falco#overview)

-	**Project Repository:** https://github.com/falcosecurity/falco
-	**Chat:** Slack: `#falco` in [slack.sysdig.com](https://slack.sysdig.com/)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [Falco CLA](https://github.com/falcosecurity/falco#contributor-license-agreements)

### Argo

*"Open source Kubernetes native workflows, events, CI and CD"* - [https://argoproj.github.io/ ][$Argo overview]

-	**Project Repository:** https://github.com/argoproj/
-	**Contributor Guide:** [argo/community/][(https://github.com/argoproj/argo/tree/master/community#contributing-to-argo\)
-	**Chat:** [Argo Slack](https://argoproj.slack.com/)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [Argo CLA](https://github.com/argoproj/argo/tree/master/community#contributing-to-argo)

### Dragonfly

*"Dragonfly is an intelligent P2P-based image and file distribution tool. It aims to improve the efficiency and success rate of file transferring, and maximize the usage of network bandwidth, especially for the distribution of larget amounts of data, such as application distribution, cache distribution, log distribution, and image distribution."* - [Overview - d7y.io](https://d7y.io/)

-	**Project Repository:** https://github.com/dragonflyoss/dragonfly
-	**Contributor Guide:** [dragonflyoss/dragonfly/CONTRIBUTING](https://github.com/dragonflyoss/Dragonfly/blob/master/CONTRIBUTING.md)
-	**Chat:** [gitter.im/alibaba/Dragonfly](https://gitter.im/alibaba/Dragonfly)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)

### SPIFFE

*"SPIFFE (Secure Production Identity Framework For Everyone) provides a secure identity, in the form of a specially crafted X.509 certificate, to every workload in a modern production environment. SPIFFE removes the need for application-level authentication and complex network-level ACL configuration."* - [What is SPIFFE? - spiffe.io](https://spiffe.io/)

-	**Project Repository:** https://github.com/spiffe/spiffe
-	**Contributor Guide:** [spiffe/spiffe/contributing](https://github.com/spiffe/spiffe/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [slack.spiffe.io](https://slack.spiffe.io/)
-	**Developer Mailing List/Forum:** [SPIFFE Dev Discussion Mailing List](https://groups.google.com/a/spiffe.io/forum/#!forum/dev-discussion)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### SPIRE

*"SPIRE (the SPIFFE Runtime Environment) is a tool-chain for establishing trust between software systems across a wide variety of hosting platforms. Concretely, SPIRE exposes the SPIFFE Workload API, which can attest running software systems and issue SPIFFE IDs and SVIDs to them."* - [Spire Readme](https://github.com/spiffe/spire/blob/master/README.md)

-	**Project Repository:** https://github.com/spiffe/spire
-	**Contributor Guide:** [spiffe/spire/contributing](https://github.com/spiffe/spire/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [slack.spiffe.io](https://slack.spiffe.io/)
-	**Developer Mailing List/Forum:** [SPIFFE Dev Discussion Mailing List](https://groups.google.com/a/spiffe.io/forum/#!forum/dev-discussion)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)


### Contour

*"Contour is an open source Kubernetes ingress controller providing the control plane for the Envoy edge and service proxy. Contour supports dynamic configuration updates and multi-team ingress delegation out of the box while maintaining a lightweight profile."* - [projectcontour.io](https://projectcontour.io)

-	**Project Repository:** https://github.com/projectcontour/contour
-	**Contributor Guide:** [projectcontour/community](https://github.com/projectcontour/community)
-	**Chat:** Slack: `#contour` in [kubernetes.slack.com](https://kubernetes.slack.com/messages/contour)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

---

Sandbox Projects
----------------

### Telepresence

*"Telepresence is an open source tool that lets you run a single service locally, while connecting that service to a remote Kubernetes cluster."* - [Overview - telepresence.io](https://www.telepresence.io/discussion/overview)

-	**Project Repository:** https://github.com/telepresenceio/telepresence
-	**Contributor Guide:** [telepresenceio/telepresence/docs/reference/developing](https://github.com/telepresenceio/telepresence/blob/master/docs/reference/developing.md)
-	**Chat:** Slack: `#telepresence` in [datawire-oss.slack.com](https://d6e.co/slack)
-	**Developer Mailing List/Forum:** None
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### OpenMetrics

*"An effort to create an open standard for transmitting metrics at scale, with support for both text representation and Protocol Buffers."* - [openmetrics.io](https://openmetrics.io/)

-	**Project Repository:** https://github.com/OpenObservability/OpenMetrics
-	**Contributor Guide:** [TBD]
-	**Chat:** Slack: `#openmetrics` in [slack.cncf.io](https://cloud-native.slack.com/messages/CC6CPDEJV/)
-	**Developer Mailing List/Forum:** [OpemMetrics Mailing List](https://groups.google.com/forum/m/#!forum/openmetrics)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)

### Cortex

*"Cortex provides horizontally scalable, multi-tenant, long term storage for Prometheus metrics when used as a remote write destination, and a horizontally scalable, Prometheus-compatible query API."* - [Cortex Readme](https://github.com/cortexproject/cortex#open-source-horizontally-scalable-multi-tenant-prometheus-as-a-service)

-	**Project Repository:** https://github.com/cortexproject/cortex
-	**Contributor Guide:** [cortextproject/cortex/readme](https://github.com/cortexproject/cortex#for-developers)
-	**Chat:** Slack: `#cortext` in [slack.cncf.io](https://cloud-native.slack.com/messages/cortex/)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Buildpacks

*"Buildpacks provide a higher-level abstraction for building apps compared to Dockerfiles."* - [What Are Buildpacks? - buildpacks.io](https://buildpacks.io/)

-	**Project Repository:** https://github.com/buildpack
-	**Chat:** Slack: [slack.buildpacks.io](https://slack.buildpacks.io/)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)

### Virtual Kubelet

*"Virtual Kubelet is an open source Kubernetes kubelet implementation that masquerades as a kubelet for the purposes of connecting Kubernetes to other APIs. This allows the nodes to be backed by other services like ACI, AWS Fargate, Hyper.sh, IoT Edge etc. The primary scenario for VK is enabling the extension of the Kubernetes API into serverless container platforms like ACI, Fargate, and Hyper.sh, though we are open to others. However, it should be noted that VK is explicitly not intended to be an alternative to Kubernetes federation."* - [Virtual Kubelet Readme](https://github.com/virtual-kubelet/virtual-kubelet)

-	**Project Repository:** https://github.com/virtual-kubelet/virtual-kubelet
-	**Contributor Guide:** [virtual-kubelet/virtual-kubelet/contributing](https://github.com/virtual-kubelet/virtual-kubelet/blob/master/CONTRIBUTING.md)
-	**License:** [Apache 2.0](apache-license)
-	**Legal Requirements:** [CNCF CLA](https://github.com/cncf/cla)

### KubeEdge

*"KubeEdge is an open source system for extending native containerized application orchestration capabilities to hosts at Edge."* - [KubeEdge website](https://kubeedge.io/)

-	**Project Repository:** https://github.com/kubeedge/kubeedge
-	**Contributor Guide:** [KubeEdge-contributor-guide](https://github.com/kubeedge/kubeedge/blob/master/CONTRIBUTING.md)
-	**Chat:** [KubeEdge](https://kubeedge.slack.com/)
-	**Developer List/Forum:** [KubeEdge](https://groups.google.com/forum/?hl=en#!forum/kubeedge)
-	**License:** [Apache 2.0](apache-license)

### Keptn
*"Keptn is a control-plane for continuous delivery and automated operations."* - [Keptn.sh](https://keptn.sh)

-	**Project Repository:** https://github.com/keptn/keptn
-	**Contributor Guide:** [Keptn-contributor-guide](https://github.com/keptn/keptn/blob/master/CONTRIBUTING.md)
-	**Chat:** [Keptn Slack](https://slack.keptn.sh/)
-	**Developer List/Forum:** [Keptn mailing list](https://groups.google.com/forum/#!forum/keptn)
-	**License:** [Apache 2.0](apache-license)

### Brigade

*"Brigade is a tool for running scriptable, automated tasks in the cloud — as part of your Kubernetes cluster."* - [Brigade-overview](https://brigade.sh/)

-	**Project Repository:** https://github.com/brigadecore/brigade/
-	**Contributor Guide:** [Brigade-contributor-guide](https://docs.brigade.sh/topics/developers/)
-	**Chat:** [Brigade-chat](https://kubernetes.slack.com/messages/C87MF1RFD/)

### Network Service Mesh

*"Network Service Mesh (NSM) is a novel approach to solving complicated L2/L3 use cases in Kubernetes that are tricky to address withing the existing Kubernetes Network Model. Inspired by Istio, Network Service Mesh maps the concept of a service mesh to L2/L3 payloads."* - [What is Network Service Mesh? - networkservicemesh.io ](https://networkservicemesh.io/docs/concepts/what-is-nsm)

-	**Project Repository:** https://github.com/networkservicemesh
-	**Chat:** Slack `#nsm-dev` in [slack.cncf.io](https://cloud-native.slack.com/messages/CHSKJ4849/)
-	**Developer List/Forum:** [Network Service Mesh Mailing List](https://groups.google.com/forum/#!forum/networkservicemesh)

### OpenTelemetry

*"OpenTelemetry is made up of an integrated set of APIs and libraries as well as a collection mechanism via an agent and collector. These components are used to generate, collect, and describe telemetry about distributed systems. This data includes basic context propagation, distributed traces, metrics, and other signals in the future. OpenTelemetry is designed to make it easy to get critical telemetry data out of your services and into your backend(s) of choice. For each supported language it offers a single set of APIs, libraries, and data specifications, and developers can take advantage of whichever components they see fit.* - [What is OpenTelemetry? - opentelemetry.io](https://opentelemetry.io/)

-	**Project Repository:** https://github.com/open-telemetry
-	**Chat:** Gitter: [open-telemetry/community](https://gitter.im/open-telemetry/community)
-	**Developer List/Forum:** [OpenTelemetry Dev Mailing List](https://lists.cncf.io/g/cncf-opentelemetry-contributors)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)

### OpenEBS

*"OpenEBS is the leading open-source project for container-attached and container-native storage on Kubernetes. OpenEBS adopts Container Attached Storage (CAS) approach, where each workload is provided with a dedicated storage controller. OpenEBS implements granular storage policies and isolation that enable users to optimize storage for each specific workload. OpenEBS runs in user space and does not have any Linux kernel module dependencies."* - [Introduction - OpenEBS.io](https://docs.openebs.io/#font-size-6-introduction-font)

-	**Project Repository:** https://github.com/openebs
-	**Contributor Guide:** [openebs/openebs/CONTRIBUTING](https://github.com/openebs/openebs/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [openebs-community.slack.com](https://openebs-community.slack.com/messages/openebs-users/)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Thanos

*"Thanos is a set of components that can be composed into a highly available metric system with unlimited storage capacity, which can be added seamlessly on top of existing Prometheus deployments."* - [Overview - Thanos readme](https://github.com/thanos-io/thanos#overview)

-	**Project Repository:** https://github.com/thanos-io/thanos
-	**Contributor Guide:** [thanos.io/contributing](https://thanos.io/tip/contributing/)
-	**Chat:** Slack: `#thanos` in [slack.cncf.io](https://app.slack.com/client/T08PSQ7BQ/CK5RSSC10)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Flux

*"Flux is a tool that automatically ensures that the state of your Kubernetes cluster matches the configuration you’ve supplied in Git. It uses an operator in the cluster to trigger deployments inside Kubernetes, which means that you don’t need a separate continuous delivery tool."* - [Flux - fluxcd.io](https://fluxcd.io/)

-	**Project Repository:** https://github.com/fluxcd/flux
-	**Contributor Guide:** [fluxcd/flux/CONTRIBUTING](https://github.com/fluxcd/flux/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: `#flux` in [slack.cncf.io](https://app.slack.com/client/T08PSQ7BQ/CLAJ40HV3)
-	**Developer List/Forum:** [Flux Dev Mailing List](https://lists.cncf.io/g/cncf-flux-dev)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### in-toto

*"in-toto provides a framework to protect the integrity of the software supply chain. It does so by verifying that each task in the chain is carried out as planned, by authorized personnel only, and that the product is not tampered with in transit."* - [in-toto - in-toto Readme](https://github.com/in-toto/in-toto#in-toto----)

-	**Project Repository:** https://github.com/in-toto/in-toto
-	**Contributor Guide:** [Instructions for Contributors](https://github.com/in-toto/in-toto#instructions-for-contributors)
-	**Developer List/Forum:** [in-toto mailing list](mailto:in-toto-dev@googlegroups.com)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Litmus
*"Litmus is a toolset to perform cloud-native chaos engineering. Litmus provides tools to orchestrate chaos on Kubernetes and helps SREs find weaknesses in their deployments. SREs use Litmus to run chaos experiments initially in the staging environment and eventually in production to find bugs, vulnerabilities. Fixing the weaknesses leads to increased resilience of the system."* - [Litmus](https://litmuschaos.io/)

-	**Project Repository:** https://github.com/litmuschaos/litmus
-	**Contributor Guide:** [Litmus-contributor-guide](https://github.com/litmuschaos/litmus/blob/master/CONTRIBUTING.md)
-	**Chat:** [Litmus Slack](https://slack.litmuschaos.io)
-	**License:** [Apache 2.0](apache-license)

### Tinkerbell

*"Tinkerbell is a bare metal provisioning engine. Tinkerbell standardizes infrastructure and application management using the same API-centric, declarative configuration and automation approach pioneered by the Kubernetes community."* - [Tinkerbell](https://tinkerbell.org/)

-	**Project Repository:** https://github.com/tinkerbell
-	**Contributor Guide:** [tinkerbell/tink/CONTRIBUTING](https://github.com/tinkerbell/tink/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: `#tinkerbell` in [slack.cncf.io](https://app.slack.com/client/T08PSQ7BQ/C01SRB41GMT)
-	**Developer List/Forum:** [Tinkerbell Contributors Mailing List](https://groups.google.com/g/tinkerbell-contributors?pli=1)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Service Mesh Performance

*"Service Mesh Performance is a standard for capturing and characterizing the details of infrastructure capacity, service mesh configuration, and workload metadata."* - [Service Mesh Performance](https://smp-spec.io/)

-	**Project Repository:** https://github.com/service-mesh-performance
-	**Contributor Guide:** [smp-spec.io/contributing](https://github.com/service-mesh-performance/service-mesh-performance/blob/master/CONTRIBUTING.md). Service Mesh Performance prominently and consistently engages with the Meshery project. For a more complete set of contributing guides see [docs.meshery.io/project/contributing](https://docs.meshery.io/project/contributing)
-	**Chat:** Slack: `#tinkerbell` in [slack.layer5.io](https://slack.layer5.io)
-	**Developer List/Forum:** [Service Mesh Performance Mailing Lists](https://smp-spec.io/subscribe)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)
---

Archived Projects
-----------------

### rkt

rkt is a pod-native container engine for Linux. It is composable, secure, and built on standards.

-	**Project Repository:** https://github.com/rkt/rkt
-	**Contributor Guide:** [rkt/rkt/contributing](https://github.com/rkt/rkt/blob/master/CONTRIBUTING.md)
-	**Chat:** `#rkt-dev` on [freenode](https://freenode.net) (join via [Riot](https://riot.im/app/#/room/#freenode_#rkt-dev:matrix.org)\)
-	**Developer Mailing List/Forum:** [rkt-dev Mailing List](https://groups.google.com/forum/#!forum/rkt-dev)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)
