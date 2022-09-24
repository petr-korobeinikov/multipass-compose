# 2. Use json format instead of yaml when gathering information from multipass

Date: 2022-09-24

## Status

Accepted

## Context

`multipass` produces semantically different objects in both `yaml` and `json`
format.

For example, the following command `multipass list --format yaml|json` produces
these snippets of code:

```yaml
backend:
  - state: Running
    ipv4:
      - 192.168.64.23
    release: 20.04 LTS
  ...
```

```json
{
  "list": [
    {
      "ipv4": [
        "192.168.64.23"
      ],
      "name": "backend",
      "release": "20.04 LTS",
      "state": "Running"
    },
    ...
  ]
}
```

## Decision

Use `json` format for gathering information about current state of virtual
machines.

## Consequences

It helps me to:

- keep structures in code clean and simple;
- omit useless slices;
- write more readable, understandable, and maintainable code.
