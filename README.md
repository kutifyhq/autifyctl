<a href="https://github.com/kutifyhq">
    <img src="https://avatars.githubusercontent.com/u/84083798?s=200&v=4" alt="Kutifyhq logo" title="Kutify" align="right" height="80" />
</a>

# Autifyctl

## Usage

```console
$ autifyctl schedule run 1234 --access-token $ACCESS_TOKEN
```

## Docker Image

```console
$ docker run ghcr.io/kutifyhq/autifyctl:latest schedule run 1234 --access-token $ACCESS_TOKEN
```

## Environment variables

Access token can be configure either by `--access-token` flag (`-t` for short) or `AUTIFYCTL_ACCESS_TOKEN` enivornment variable.

## License

See [LICENSE](./LICENSE) for details.
