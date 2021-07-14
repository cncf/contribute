# Contributing Guide

Want to contribute to this website? Our [Contributing Guide] has detailed
information on how to edit the site and submit a pull request.

The site is composed of two repositories:

* cncf/contribute: This repository contains the pages under the /contributors
  section of the site.
* [cncf/tag-contributor-strategy]: Contains the website infrastructure, and
  pages under the /maintainers section of the site.

[Contributing Guide]: https://cncf-contribute.netlify.app/about/contributing/
[cncf/tag-contributor-strategy]: https://github.com/cncf/sig-contributor-strategy/

## Netlify Deployments

The main website is in another repository, [cncf/tag-contributor-strategy]. When
a pull request is created for this repository, Netlify clones the main repo and
builds the website using the content from this repository. A preview of the
website is available in the GitHub pull request.

When commits are pushed to the production branch of this repository, a build of
the main website is triggered in Netlify.
