# multipass-compose

> Compose multipass™ virtual machines like a ... charm!

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

---

This project is licensed under the terms of the MIT license.
