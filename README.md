[![No Maintenance Intended](http://unmaintained.tech/badge.svg)](http://unmaintained.tech/)

# multipass-compose

> Compose multipass™ virtual machines like a ... charm!

## Deprecation warning

Multipass is a buggy and inconvenient project to manage virtual machines with Ubuntu.
So this project is now deprecated.
Think about using `lima`/`colima` instead which is more powerful and truly supported by community.

## Install

```shell
go install github.com/petr-korobeinikov/multipass-compose/cmd/multipass-compose@latest
```

## TL;DR

1. Define your `multipass-compose.yaml`:

   ```yaml
   services:
     web-server:
       image: focal
     database:
       image: focal
     backend:
       image: focal
   ```

2. Start machines up:

   ```shell
   multipass-compose up
   ```

3. Check the status:

   ```shell
   multipass-compose status
   ```

4. Do the job :hammer_and_wrench: :carpentry_saw: :hammer: :wrench:
5. Turn machines off:

   ```shell
   multipass-compose down
   ```

## Command reference

```shell
multipass-compose up          # Starts machines up
multipass-compose down        # Turns machines down
multipass-compose status      # Shows machines status
multipass-compose ip <name>   # Shows IPv4 address of given machine
```

## `multipass-compose.yaml` reference

```yaml
services: # Defines a list of services
  lb: # Custom machine name
    image: focal # Ubuntu release, see `multipass find`
    cpus: 1 # Number of CPUs to allocate.
    mem: 1G # Amount of memory to allocate. Positive integers, in bytes, or with K, M, G suffix.
    disk: 5G # Disk space to allocate. Positive integers, in bytes, or with K, M, G suffix.
    cloud-init: path/to/cloud/config.yaml # Path or URL to a user-data cloud-init configuration.
  database:
    image: focal
    cpus: 4
    mem: 4G
    disk: 80G
  backend:
    image: focal
    cpus: 2
    mem: 2G
    disk: 20G
```

A complete list of examples you can find out
in [this repo](https://github.com/petr-korobeinikov/multipass-compose-showcase).

## Development

`bats` helpers installed as git submodules. To fetch them out use:

```shell
git submodule update --init --recursive
```
