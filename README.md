# Morpheus v2
[![CircleCI branch](https://img.shields.io/circleci/project/github/Nordgedanken/Morpheusv2/master.svg)](https://circleci.com/gh/Nordgedanken/Morpheusv2)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FNordgedanken%2FMorpheusv2.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FNordgedanken%2FMorpheusv2?ref=badge_shield)

[![Github All Releases](https://img.shields.io/github/downloads/Nordgedanken/Morpheusv2/total.svg)]()

[![Go Report Card](https://goreportcard.com/badge/github.com/Nordgedanken/Morpheusv2)](https://goreportcard.com/report/github.com/Nordgedanken/Morpheusv2)
---

A Matrix client written in Go-QT and a Reboot of my previous Project [Morpheus](https://github.com/Nordgedanken/Morpheus)

## Contributing - Matrix Room
Join the Matrix Room at [#Morpheus:matrix.ffslfl.net](https://matrix.to/#/#Morpheus:matrix.ffslfl.net)
Read the [Contribution Guideline](CONTRIBUTING.md)

## How to build
### Prerequisites
1. https://github.com/therecipe/qt
   * Follow https://github.com/therecipe/qt/wiki/Installation
2. Clone this repo by doing `go get -u github.com/Nordgedanken/Morpheusv2`

### Build
1. Run `qtdeploy build desktop` inside  `$GOPATH/src/github.com/Nordgedanken/Morpheusv2`
2. Run the Application from within `deploy/**`

## How to Build the Windows Installer
1. Get a Windows PC
2. Follow setup on https://github.com/mh-cbon/go-msi
3. Download the Morpheus.exe from the latest CI Build for your tag into the Source Folder or the build result (it needs to be static linked)
4. Run `go-msi make --msi Morpheusv2.msi --version 0.1.0 --src scripts/templates`

## Versioning
We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/Nordgedanken/Morpheus/tags).


## License
This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FNordgedanken%2FMorpheusv2.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FNordgedanken%2FMorpheusv2?ref=badge_large)

## Acknowledgments
* Inspired by [nheko](http://github.com/mujx/nheko)