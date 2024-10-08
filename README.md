# Webhook Bridge

This application is designed to take in github events and feed them to a Redis stream, 
which allows for event driven architectures.

## Features

1. A stream per event type
2. Additional metadata on streams for interesting parts of an event (e.g. `action` of: merged, labelled on a pull request)
3. Prometheus metric endpoint
4. Swagger ui
5. Configuration via disk and environment variables, allowing you to not put secrets in the env for github.

This application exposes a web server on port `3000` by default and expects to get events on `api/v1/webhook/github`
Metrics are on `/metrics` and health is on `/healthz`

If it cannot connect to redis the health will fail and events will be checked when it next comes up. It will use the 
GitHub app to look up events between it's last posted message id and the current event(s) that are marked as failed. 

## Configuration

The application can be configured using environment variables or a configuration file. Environment variables take precedence over configuration file settings when both are present.

### Environment Variables

| Name                  | Description                                | Default                                      |
|-----------------------|--------------------------------------------|----------------------------------------------|
| `LOG_LEVEL`           | Log level for the application              | `"info"`                                     |
| `API_VERSION`         | The version of the api                     | (Set in Docker at build time to version tag) |
| `API_HOSTNAME`        | Hostname for this api                      | `"localhost:3000"`                           |
| `REDIS_HOSTNAME`      | Redis server hostname                      |                                              |
| `REDIS_PASSWORD`      | Redis server password                      |                                              | 
| `REDIS_DB`            | Redis database number                      | `0`                                          | 
| `GITHUB_HMAC_ENABLED` | Enable HMAC validation for GitHub webhooks | `true`                                       |
| `GITHUB_HMAC_SECRET`  | HMAC secret for GitHub webhooks            |                                              |

### Configuration File

The application can also be configured using a YAML configuration file. The default configuration file name is `config.yaml`, and it can be placed in `/etc/app/` or the current directory.
Using the cli args you can also set a custom configuration file path.

```yaml
log:
  level: info

api:
  version: dev
  hostname: localhost:3000

redis:
  hostname: your-redis-hostname
  password: your-redis-password
  db: 0

github:
  hmac:
    enabled: true
    secret: your-hmac-secret
```

## Code decisions

For this app it is decided that the models for the app should sit in with their appropriate package, so each package will
have it's models locally and provide interfaces for others to consume. 
