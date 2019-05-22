# Buildpack for Consul agent as Sidecar

This project is a starting point for running a Consul agent as a sidcar next to Cloud Foundry applications.

Status:

* [x] Works in Linux containers
* [x] Installs `consul` in $PATH
* [ ] Works in Windows containers
* [x] Application's `manifest.yml` can run any arbitrary `consul agent ...` command or wrapper script
* [ ] Demonstrate joining remote Consul cluster
* [ ] Demonstrate joining remote Consul cluster on CFDev
* [ ] Native support for joining existing Consul clusters
* [ ] Configure CF app as a Consul service
* [ ] Setup [Consul Connect](https://learn.hashicorp.com/consul/getting-started/connect) for proxied traffic to other Consul client apps

## Demonstration

Locally, run a `consul agent -server` node:

```plain
consul agent -server -bootstrap -bootstrap-expect 1  \
    -advertise 10.144.0.34 -bind 0.0.0.0 \
    -config-dir tmp/demo.consul.d -data-dir tmp/demo.consul-data.d
```

```plain
cf v3-create-app app-using-consul
cf v3-apply-manifest -f fixtures/rubyapp/manifest.yml
cf v3-push app-using-consul -p fixtures/rubyapp
```

```plain
cf ssh app-using-consul
/tmp/lifecycle/shell
```

```plain
consul agent -bind 0.0.0.0 \
     -join host.cfdev.sh \
    -node cfdev-$INSTANCE_GUID \
    -config-dir .config.d/ -data-dir .consul-data.d/
```

Unfortunately, this fails.

In the consul server logs we see:

```plain
    2019/05/23 08:27:23 [INFO] consul: member 'cfdev-5fbaaf3c-4d37-4be2-5019-6068' joined, marking health alive
    2019/05/23 08:27:25 [INFO] memberlist: Suspect cfdev-5fbaaf3c-4d37-4be2-5019-6068 has failed, no acks received
    2019/05/23 08:27:28 [INFO] memberlist: Suspect cfdev-5fbaaf3c-4d37-4be2-5019-6068 has failed, no acks received
    2019/05/23 08:27:29 [INFO] memberlist: Marking cfdev-5fbaaf3c-4d37-4be2-5019-6068 as failed, suspect timeout reached (0 peer confirmations)
```

The consul client logs show:

```plain
    2019/05/22 22:27:23 [INFO] serf: EventMemberJoin: cfdev-5fbaaf3c-4d37-4be2-5019-6068 10.255.103.9
    2019/05/22 22:27:23 [INFO] agent: Started DNS server 127.0.0.1:8600 (tcp)
    2019/05/22 22:27:23 [INFO] agent: Started DNS server 127.0.0.1:8600 (udp)
    2019/05/22 22:27:23 [INFO] agent: Started HTTP server on 127.0.0.1:8500 (tcp)
    2019/05/22 22:27:23 [INFO] agent: (LAN) joining: [host.cfdev.sh]
    2019/05/22 22:27:23 [INFO] serf: EventMemberJoin: starkair.gateway 10.144.0.34
    2019/05/22 22:27:23 [INFO] agent: (LAN) joined: 1 Err: <nil>
    2019/05/22 22:27:23 [INFO] agent: started state syncer
    2019/05/22 22:27:23 [INFO] consul: adding server starkair.gateway (Addr: tcp/10.144.0.34:8300) (DC: dc1)
    2019/05/22 22:27:23 [WARN] manager: No servers available
    2019/05/22 22:27:23 [ERR] agent: failed to sync remote state: No known Consul servers
    2019/05/22 22:27:24 [ERR] consul: "Catalog.NodeServices" RPC failed to server 10.144.0.34:8300: rpc error getting client: failed to get conn: dial tcp <nil>->10.144.0.34:8300: connect: connection refused
    2019/05/22 22:27:24 [ERR] agent: failed to sync remote state: rpc error getting client: failed to get conn: dial tcp <nil>->10.144.0.34:8300: connect: connection refused
    2019/05/22 22:27:25 [INFO] memberlist: Suspect starkair.gateway has failed, no acks received
```

## Buildpack User Documentation

### Building the Buildpack

To build this buildpack, run the following command from the buildpack's directory:

1. Source the .envrc file in the buildpack directory.

    ```bash
    source .envrc
    ```

    To simplify the process in the future, install [direnv](https://direnv.net/) which will automatically source .envrc when you change directories.

1. Install buildpack-packager

    ```bash
    ./scripts/install_tools.sh
    ```

1. Build the buildpack

    ```bash
    buildpack-packager build
    ```

1. Use in Cloud Foundry

    Upload the buildpack to your Cloud Foundry and optionally specify it by name

    ```bash
    cf create-buildpack [BUILDPACK_NAME] [BUILDPACK_ZIP_FILE_PATH] 1
    cf push my_app [-b BUILDPACK_NAME]
    ```

### Testing

Buildpacks use the [Cutlass](https://github.com/cloudfoundry/libbuildpack/cutlass) framework for running integration tests.

To test this buildpack, run the following command from the buildpack's directory:

1. Source the .envrc file in the buildpack directory.

    ```bash
    source .envrc
    ```

    To simplify the process in the future, install [direnv](https://direnv.net/) which will automatically source .envrc when you change directories.

1. Run unit tests

    ```bash
    ./scripts/unit.sh
    ```

1. Run integration tests

    ```bash
    ./scripts/integration.sh
    ```

    More information can be found on Github [cutlass](https://github.com/cloudfoundry/libbuildpack/cutlass).

### Reporting Issues

Open an issue on this project

## Disclaimer

This buildpack is experimental and not yet intended for production use.
