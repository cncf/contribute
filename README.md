Contribution Guides to the CNCF Ecosystem
=========================================

Welcome! Are you interested in contributing to one of CNCF hosted projects? This repository should help you contains information and guidelines about contributions to CNCF projects.

CNCF offers multiple ways to start contributing to the CNCF ecosystem, including either foundation-wide and project-wide opportunities.

TOC
===

The CNCF TOC is the technical governing body of the CNCF Foundation. The detailed information on CNCF TOC, incuding its duties and responsibilities, togeteter with the information on collaboration is listed on [CNCF TOC repo](https://github.com/cncf/toc/).

Working Groups
--------------

Working groups (WG's) are the community-driven groups with the goal of continuous collaboration in the speficic areas. CNCF WG's are created and curated by the CNCF TOC and driven by the community members. CNCF TOC repo provides more [details](https://github.com/cncf/toc/tree/master/workinggroups#cncf-working-groups) on the purpose and goals of WG's, together with the [list of them](https://github.com/cncf/toc/blob/master/README.md#working-groups).

TOC Contributors
----------------

The recommended way to start contributing to CNCF TOC - acting as a TOC Contributor. Here are more [details](https://github.com/cncf/toc/blob/master/CONTRIBUTORS.md) on goals and purpose of the initiative.

Projects
========

The Projects of the Cloud Native Computing Foundation are listed below, together with the brief information on engaging with them:

### CNCF Project Grades

All projects of the Cloud Native Computing Foundation are classified with one of three stages of maturity:

-	Graduated
-	Incubated
-	Sandboxed

[CNCF Graduation Criteria](https://github.com/cncf/toc/blob/master/process/graduation_criteria.adoc) are documented in the CNCF TOC repo. The document describes the maturity stages of the projects.

| Project Name                            | Maturity                           | Focus                   |
|:---------------------------------------:|:----------------------------------:|:-----------------------:|
|        [Kubernetes](#kubernetes)        |  [Graduated](#graduated-projects)  |      Orchestration      |
|        [Prometheus](#prometheus)        |  [Graduated](#graduated-projects)  |       Monitoring        |
|               [CNI](#cni)               | [Incubating](#incubating-projects) |     Networking API      |
|        [containerd](#containerd)        | [Incubating](#incubating-projects) |    Container Runtime    |
|           [CoreDNS](#coredns)           | [Incubating](#incubating-projects) |    Service Discovery    |
|             [Envoy](#envoy)             | [Incubating](#incubating-projects) |      Service Mesh       |
|           [Fluentd](#fluentd)           | [Incubating](#incubating-projects) |         Logging         |
|              [gRPC](#grpc)              | [Incubating](#incubating-projects) |  Remote Procedure Call  |
|              [Helm](#helm)              | [Incubating](#incubating-projects) |   Package Management    |
|            [Jaeger](#jaeger)            | [Incubating](#incubating-projects) |   Distributed Tracing   |
|           [Linkerd](#linkerd)           | [Incubating](#incubating-projects) |      Service Mesh       |
|              [NATS](#nats)              | [Incubating](#incubating-projects) |        Messaging        |
|            [Notary](#notary)            | [Incubating](#incubating-projects) |        Security         |
|       [OpenTracing](#opentracing)       | [Incubating](#incubating-projects) | Distributed Tracing API |
|               [rkt](#rkt)               | [Incubating](#incubating-projects) |    Container Runtime    |
|               [TUF](#tuf)               | [Incubating](#incubating-projects) |  Software Update Spec   |
|            [Vitess](#vitess)            | [Incubating](#incubating-projects) |         Storage         |
|       [CloudEvents](#cloudevents)       |    [Sandbox](#sandbox-projects)    |       Serverless        |
|            [Harbor](#harbor)            |    [Sandbox](#sandbox-projects)    |        Registry         |
| [Open Policy Agent](#open-policy-agent) |    [Sandbox](#sandbox-projects)    |         Policy          |
|              [Rook](#rook)              |    [Sandbox](#sandbox-projects)    |         Storage         |
|            [SPIFFE](#spiffe)            |    [Sandbox](#sandbox-projects)    |      Identity Spec      |
|             [SPIRE](#spire)             |    [Sandbox](#sandbox-projects)    |        Identity         |
|      [Telepresence](#telepresence)      |    [Sandbox](#sandbox-projects)    |         Tooling         |
|       [OpenMetrics](#openmetrics)       |    [Sandbox](#sandbox-projects)    |         Tooling         |

---

Graduated Projects
------------------

### Kubernetes

*"Kubernetes is a portable, extensible open-source platform for managing containerized workloads and services, thatfacilitates both declarative configuration and automation. It has a large, rapidly growing ecosystem. Kubernetes services, support, and tools are widely available.*

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
-	**Chat:** IRC: `#prometheus` on [freenode](https://freenode.net/) (join via [Riot](https://riot.im/app/#/room/#freenode_#prometheus:matrix.org)\)
-	**Developer Mailing List/Forum:** [Prometheus-Developers Mailing List](https://groups.google.com/forum/#!forum/prometheus-developers)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

---

Incubating Projects
-------------------

### CNI

*"CNI (Container Network Interface), a Cloud Native Computing Foundation project, consists of a specification and libraries for writing plugins to configure network interfaces in Linux containers, along with a number of supported plugins. CNI concerns itself only with network connectivity of containers and removing allocated resources when the container is deleted. Because of this focus, CNI has a wide range of support and the specification is simple to implement."* - [What is CNI?- CNI Readme](https://github.com/containernetworking/cni/blob/master/README.md#what-is-cni)

-	**Project Repository:** https://github.com/containernetworking/cni
-	**Contributor Guide:** [containernetworking/cni/contributing](https://github.com/containernetworking/cni/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [containernetworking.slack.com](https://containernetworking.slack.com/)
-	**Developer Mailing List/Forum:** [CNI-dev Mailing List](https://groups.google.com/forum/#!forum/cni-dev)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### containerd

*"containerd is an industry-standard core container runtime with an emphasis on simplicity, robustness and portability. It is available as a daemon for Linux and Windows, which can manage the complete container lifecycle of its host system: image transfer and storage, container execution and supervision, low-level storage and network attachments, etc.."* - [About containerd - containerd.io](https://containerd.io/#about-containerd)

-	**Project Repository:** https://github.com/containerd/containerd
-	**Contributor Guide:** [containerd/containerd/contributing](https://github.com/containerd/containerd/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [slack.containerd.io](https://slack.containerd.io/)
-	**Developer Mailing List/Forum:** None
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### CoreDNS

*"CoreDNS is a DNS server. It is written in Go. It can be used in a multitude of environments because of its flexibility. CoreDNS is licensed under the Apache License Version 2, and completely open source."* - [What is it? - coredns.io](https://coredns.io/)

-	**Project Repository:** https://github.com/coredns/coredns
-	**Contributor Guide:** [coredns/coredns/contributing](https://github.com/coredns/coredns/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: `#coredns` in [slack.cncf.io](https://slack.cncf.io/)
-	**Developer Mailing List/Forum:** [Coredns-Discuss Mailing List](https://groups.google.com/forum/#!forum/coredns-discuss)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:**

### Envoy

*"Originally built at Lyft, Envoy is a high performance C++ distributed proxy designed for single services and applications, as well as a communication bus and “universal data plane” designed for large microservice “service mesh” architectures. Built on the learnings of solutions such as NGINX, HAProxy, hardware load balancers, and cloud load balancers, Envoy runs alongside every application and abstracts the network by providing common features in a platform-agnostic manner. When all service traffic in an infrastructure flows via an Envoy mesh, it becomes easy to visualize problem areas via consistent observability, tune overall performance, and add substrate features in a single place."* - [Why Envoy? - envoyproxy.io](https://www.envoyproxy.io/#why-envoy)

-	**Project Repository:** https://github.com/envoyproxy/envoy
-	**Contributor Guide:** [envoyproxy/envoy/contributing](https://github.com/envoyproxy/envoy/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [envoyslack.cncf.io](https://envoyslack.cncf.io/)
-	**Developer Mailing List/Forum:** [Envoy-Dev Mailing List](https://groups.google.com/forum/#!forum/envoy-dev)
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

### gRPC

*"gRPC is a modern open source high performance RPC framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services."* - [About - grpc.io](https://grpc.io/about/)

-	**Project Repository:** https://github.com/grpc/grpc
-	**Contributor Guide:** [grpc/grpc/contributing](https://github.com/grpc/grpc/blob/master/CONTRIBUTING.md)
-	**Chat:** Gitter: [gitter.im/grpc/grpc](https://gitter.im/grpc/grpc)
-	**Developer List/Forum:** [gRPC-io Mailing List](https://groups.google.com/forum/#!forum/grpc-io)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [CNCF CLA](https://github.com/cncf/cla)

### Helm

*"Helm helps you manage Kubernetes applications — Helm Charts helps you define, install, and upgrade even the most complex Kubernetes application.* *Charts are easy to create, version, share, and publish — so start using Helm and stop the copy-and-paste madness.* *The latest version of Helm is maintained by the CNCF - in collaboration with Microsoft, Google, Bitnami and the Helm contributor community."* - [What is Helm? - helm.sh](https://www.helm.sh/)

-	**Project Repository:** https://github.com/helm/helm
-	**Contributor Guide:** [helm/helm/contributing](https://github.com/helm/helm/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack `#helm-dev` in [slack.k8s.io](https://slack.k8s.io/)
-	**Developer List/Forum:** [CNCF-Helm Mailing List](https://lists.cncf.io/g/cncf-helm)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [CNCF CLA](https://github.com/cncf/cla)

### Jaeger

*"Jaeger, inspired by Dapper and OpenZipkin, is a distributed tracing system released as open source by Uber Technologies. It is used for monitoring and troubleshooting microservices-based distributed systems."* - [About - jaegertracing.io](https://www.jaegertracing.io/docs/#about)

-	**Project Repository:** https://github.com/jaegertracing/jaeger
-	**Contributor Guide:** [jaegertracing/jaeger/contributing](https://github.com/jaegertracing/jaeger/blob/master/CONTRIBUTING.md)
-	**Chat:** Gitter: [gitter.im/jaegertracing/Lobby](https://gitter.im/jaegertracing/Lobby)
-	**Developer Mailing List/Forum:** [Jaeger-Tracing Mailing List](https://groups.google.com/forum/#!forum/jaeger-tracing)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Linkerd

*"Linkerd is a transparent service mesh, designed to make modern applications safe and sane by transparently adding service discovery, load balancing, failure handling, instrumentation, and routing to all inter-service communication."* - [Linkerd Readme](https://github.com/linkerd/linkerd)

-	**Project Repository:** https://github.com/linkerd/linkerd
-	**Contributor Guide:** [linkerd/linkerd/contributing](https://github.com/linkerd/linkerd/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [slack.linkerd.io](https://slack.linkerd.io/)
-	**Developer Mailing List/Forum:** [Linkerd Forum](https://discourse.linkerd.io/)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### NATS

*"NATS is an open source, lightweight, high-performance cloud native infrastructure messaging system. It implements a highly scalable and elegant publish-subscribe (pub/sub) distribution model. The performant nature of NATS make it an ideal base for building modern, reliable, scalable cloud native distributed systems."* - [What is NATS? - nats.io](https://nats.io/documentation/faq/)

-	**Project Repository:** https://github.com/nats-io
-	**Contributor Guide:** [nats.io/documentation/contributing/](https://nats.io/documentation/contributing/)
-	**Chat:** Slack: [natsio.slack.com](https://natsio.slack.com/join/shared_invite/enQtMzE2NDkxNDI2NTE1LTc5ZDEzYTkwYWZkYWQ5YjY1MzBjMWZmYzA5OGQxMzlkMGQzMjYxNGM3MWYxMjNiYmNjNzIwMTVjMWE2ZDgxZGM)
-	**Developer Mailing List/Forum:** [natsio Mailing List](https://groups.google.com/forum/#!forum/natsio)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** None

### Notary

*"Notary aims to make the internet more secure by making it easy for people to publish and verify content. We often rely on TLS to secure our communications with a web server, which is inherently flawed, as any compromise of the server enables malicious content to be substituted for the legitimate content."* [Overview - Notary Readme](https://github.com/theupdateframework/notary/blob/master/README.md#overview)

-	**Project Repository:** https://github.com/theupdateframework/notary
-	**Contributor Guide:** [theupdateframework/notary/contributing](https://github.com/theupdateframework/notary/blob/master/CONTRIBUTING.md)
-	**Chat:** None
-	**Developer Mailing List/Forum:** [theupateframework Mailing List](https://groups.google.com/forum/?fromgroups#!forum/theupdateframework)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### OpenTracing

Vendor-neutral APIs and instrumentation for distributed tracing.

-	**Project Repository:** https://github.com/opentracing
-	**Contributor Guide:** [opentracing-contrib/meta](https://github.com/opentracing-contrib/meta#opentracing-contributions)
-	**Chat:** Gitter: [gitter.im/opentracing/public](https://gitter.im/opentracing/public)
-	**Developer Mailing List/Forum:** [OpenTracing Mailing List](https://groups.google.com/forum/#!forum/opentracing)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** None

### rkt

rkt is a pod-native container engine for Linux. It is composable, secure, and built on standards.

-	**Project Repository:** https://github.com/rkt/rkt
-	**Contributor Guide:** [rkt/rkt/contributing](https://github.com/rkt/rkt/blob/master/CONTRIBUTING.md)
-	**Chat:** `#rkt-dev` on [freenode](https://freenode.net/) (join via [Riot](https://riot.im/app/#/room/#freenode_#rkt-dev:matrix.org)\)
-	**Developer Mailing List/Forum:** [rkt-dev Mailing List](https://groups.google.com/forum/#!forum/rkt-dev)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### TUF

*"The Update Framework (TUF) helps developers maintain the security of a software update system, even against attackers that compromise the repository or signing keys. TUF provides a flexible framework and specification that developers can adopt into any software update system."* - [TUF Readme](https://github.com/theupdateframework/tuf/blob/develop/README.md)

-	**Project Repository:** https://github.com/theupdateframework/specification
-	**Contributor Guide:** [theupdateframework/tuf/contributors](https://github.com/theupdateframework/tuf/blob/develop/docs/CONTRIBUTORS.rst)
-	**Chat:** None
-	**Developer Mailing List/Forum:** [TUF Mailing List](https://groups.google.com/forum/?fromgroups#!forum/theupdateframework)
-	**License:** Dual Licensed [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/) / [MIT](https://choosealicense.com/licenses/mit/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Vitess

*"Vitess is a database solution for deploying, scaling and managing large clusters of MySQL instances. It's architected to run as effectively in a public or private cloud architecture as it does on dedicated hardware. It combines and extends many important MySQL features with the scalability of a NoSQL database."* - [Overview - vitess.io](https://vitess.io/overview/)

-	**Project Repository:** https://github.com/vitessio/vitess
-	**Contributor Guide:** [vitessio/vitess/contributing](https://github.com/vitessio/vitess/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [vitess.slack.com](https://bit.ly/vitess-slack)
-	**Developer Mailing List/Forum:** [Vitess Mailing List](https://groups.google.com/forum/#!forum/vitess)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [Google Corporate CLA](https://cla.developers.google.com/about/google-corporate) / [Google Individual CLA](https://cla.developers.google.com/about/google-individual)

---

Sandbox Projects
----------------

### CloudEvents

CloudEvents Specification

-	**Project Repository:** https://github.com/cloudevents/spec
-	**Contributor Guide:** [cloudevents/spec/contributing](https://github.com/cloudevents/spec/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: `#cloudevents` in [slack.cncf.io](https://slack.cncf.io/)
-	**Developer Mailing List/Forum:** [CNCF-wg-Serverless Mailing List](https://lists.cncf.io/g/cncf-wg-serverless)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Harbor

*"Project Harbor is an an open source trusted cloud native registry project that stores, signs, and scans content. Harbor extends the open source Docker Distribution by adding the functionalities usually required by users such as security, identity and management. Having a registry closer to the build and run environment can improve the image transfer efficiency. Harbor supports replication of images between registries, and also offers advanced security features such as user management, access control and activity auditing."* - [Harbor Readme](https://github.com/goharbor/harbor)

-	**Project Repository:** https://github.com/goharbor/harbor
-	**Contributor Guide:** [vmware/harbor/contributing](https://github.com/goharbor/harbor/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [code.vmware.com](https://code.vmware.com/join/)
-	**Developer Mailing List/Forum:** [Harbor-Dev Mailing List](https://groups.google.com/forum/#!forum/harbor-dev)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [VMware CLA](https://cla.vmware.com/faq)

### Open Policy Agent

*"OPA is a lightweight general-purpose policy engine that can be co-located with your service. You can integrate OPA as a sidecar, host-level daemon, or library."* - [What is OPA? - openpolicyagent.org](https://www.openpolicyagent.org/docs/#what-is-opa)

-	**Project Repository:** https://github.com/open-policy-agent/opa
-	**Contributor Guide:** [open-policy-agent/opa/contributing](https://github.com/open-policy-agent/opa/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [slack.openpolicyagent.org](https://slack.openpolicyagent.org/)
-	**Developer Mailing List/Forum:** None
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** None

### Rook

*"Rook is an open source cloud-native storage orchestrator for Kubernetes, providing the platform, framework, and support for a diverse set of storage solutions to natively integrate with cloud-native environments."* [What is Rook? - Rook Readme](https://github.com/rook/rook/blob/master/README.md)

-	**Project Repository:** https://github.com/rook/rook
-	**Contributor Guide:** [rook/rook/contributing](https://github.com/rook/rook/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [rook-io.slack.com](https://rook-io.slack.com/)
-	**Developer Mailing List/Forum:** [Rook-Dev Mailing List](https://groups.google.com/forum/#!forum/rook-dev)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### SPIFFE

*"SPIFFE (Secure Production Identity Framework For Everyone) provides a secure identity, in the form of a specially crafted X.509 certificate, to every workload in a modern production environment. SPIFFE removes the need for application-level authentication and complex network-level ACL configuration."* - [What is SPIFFE? - spiffe.io](https://spiffe.io/)

-	**Project Repository:** https://github.com/spiffe/spiffe
-	**Contributor Guide:** [spiffe/spiffe/contributing](https://github.com/spiffe/spiffe/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [slack.spiffe.io](https://slack.spiffe.io/]
-	**Developer Mailing List/Forum:** [SPIFFE Dev Discussion Mailing List](https://groups.google.com/a/spiffe.io/forum/#!forum/dev-discussion)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### SPIRE

*"SPIRE (the SPIFFE Runtime Environment) is a tool-chain for establishing trust between software systems across a wide variety of hosting platforms. Concretely, SPIRE exposes the SPIFFE Workload API, which can attest running software systems and issue SPIFFE IDs and SVIDs to them."* - [Spire Readme](https://github.com/spiffe/spire/blob/master/README.md)

-	**Project Repository:** https://github.com/spiffe/spire
-	**Contributor Guide:** [spiffe/spire/contributing](https://github.com/spiffe/spire/blob/master/CONTRIBUTING.md)
-	**Chat:** Slack: [slack.spiffe.io](https://slack.spiffe.io/]
-	**Developer Mailing List/Forum:** [SPIFFE Dev Discussion Mailing List](https://groups.google.com/a/spiffe.io/forum/#!forum/dev-discussion)
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** [DCO](https://developercertificate.org/)

### Telepresence

*"Telepresence is an open source tool that lets you run a single service locally, while connecting that service to a remote Kubernetes cluster."* - [Overview - telepresene.io](https://www.telepresence.io/discussion/overview)

-	**Project Repository:** https://github.com/telepresenceio/telepresesence
-	**Contributor Guide:** [telepresenceio/telepresence/docs/reference/developing](https://github.com/telepresenceio/telepresence/blob/master/docs/reference/developing.md)
-	**Chat:** Gitter: [gitter.im/datawire/telepresence](https://gitter.im/datawire/telepresence)
-	**Developer Mailing List/Forum:** None
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** None

### OpenMetrics

*"An effort to create an open standard for transmitting metrics at scale, with support for both text representation and Protocol Buffers."* - [openmetrics.io](https://openmetrics.io/)

-	**Project Repository:** https://github.com/OpenObservability/OpenMetrics
-	**Contributor Guide:** [TBD]
-	**Chat:** Slack: [https://cloud-native.slack.com/messages/openmetrics/](https://cloud-native.slack.com/messages/CC6CPDEJV/)
-	**Developer Mailing List/Forum:** https://groups.google.com/forum/m/#!forum/openmetrics
-	**License:** [Apache 2.0](https://choosealicense.com/licenses/apache-2.0/)
-	**Legal Requirements:** None

Community engagement
====================

Ambassadors
-----------

[Cloud Native Ambassadors](https://www.cncf.io/people/ambassadors/) (CNAs) are individuals who are passionate about [Cloud Native Computing Foundation](https://www.cncf.io/) technology and projects, recognized for their expertise, and willing to help others learn about the framework and community.

Successful ambassadors are people such as bloggers, influencers, evangelists who are already engaged with a CNCF project in some way including contributing to forums, online groups, community events, etc.

Details on the Ambassadors program, and information on how to join CNCF as an Ambassador is available [here](https://github.com/cncf/ambassadors).

Meetups
-------

The Cloud Native Computing Foundation supports the worldwide community of the Cloud Native meetups (they are listed on [meetup.com](https://www.meetup.com/pro/cncf/)) website.

CNCF is currently working on expanding the Cloud Native community around the globe, and we are happy to accept the new meetup communities to join our network, and become one of the official CNCF meetups.

Details on the Meetups program, together with the best practices on running CNCF Meetups is available [here](https://github.com/cncf/meetups).

CNCF + Summer of Code
---------------------

The Cloud Native Computing Foundation participates in [Google Summer of Code](https://summerofcode.withgoogle.com/) (GSoC). CNCF is a great place to spend a summer learning, coding, participating and contributing. We are an exciting open source foundation with a vibrant community of projects, and we look forward to your application and your project ideas!

CNCF and SoC information is available [here](https://github.com/cncf/soc/blob/master/README.md).
