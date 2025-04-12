## Description

This tool automates the process of publishing a Go library by tagging a version, pushing the tag to the remote repository, and updating the Go module proxy.

## Prerequisites

* [Git](https://git-scm.com/) installed.
* [Go](https://go.dev/) installed.
* You must be in the root directory of your Go module's Git repository.
* You must have push access to the remote repository (`origin`).
## Install
```bash
go install github.com/Raezil/publisher@latest
```
## Usage

```bash
publisher <version> <module-path>
```

### Example
```bash
publisher v1.0.0 github.com/Raezil/publisher
```
