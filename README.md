# GitHub Bridge

This application is designed to take in github events and feed them to a Redis stream, 
which allows for event driven architectures.

## Features

1. A stream per event type
2. Additional metadata on streams for interesting parts of an event (e.g. merged, labelled on a pull request)
3. Prometheus metric endpoint & 
4. Swagger ui
5. GitHub App for event processing missed events
6. Configuration via disk and environment variables, allowing you to not put secrets in the env for github.

This application exposes a web server on port `3000` by default and expects to get events on `api/v1/webhook/github`
Metrics are on `/metrics` and health is on `/healthz`

If it cannot connect to redis the health will fail and events will be checked when it next comes up. It will use the 
GitHub app to look up events between it's last posted message id and the current event(s) that are marked as failed. 

## Configuration

TBC.

## Code decisions

For this app it is decided that the models for the app should sit in with their appropriate package, so each package will
have it's models locally and provide interfaces for others to consume. 

### Repository Structure

What would normally be seen as business domains here will be found in the `controllers` directory

Sure, here is a more generic and nicely formatted repository structure section for your `README.md`:

## Repository Structure

The project is organized into the following directories:

```
github-bridge/
├── cmd/
│   └── github-bridge/
│       └── main.go
├── internal/
│   ├── app/
│   │   ├── app.go
│   │   └── webhook.go
│   ├── controllers/
│   │   ├── githubController/
│   │   │   └── controller.go
│   │   └── gitlabController/
│   │       └── controller.go
│   ├── data/
│   │   ├── redis/
│   │   │   └── redis.go
│   │   └── githubapi/
│   │       └── githubapi.go
│   └── infrastructure/
│       ├── config/
│       │   └── config.go
│       └── observability/
│           ├── metrics.go
│           └── logging.go

├── go.mod
├── go.sum
└── README.md
```

### Explanation

- **cmd/github-bridge/main.go**: Entry point of the application.
- **internal/app/**: Contains the base application logic for starting and running the app.
  - `app.go`: Main application setup and run logic.
  - `webhook.go`: Webhook handler logic.
- **internal/controllers/**: Contains business domain logic.
  - `githubController/`: Logic specific to GitHub webhooks.
  - `gitlabController/`: Logic specific to GitLab webhooks.
- **internal/data/**: Contains logic for interacting with databases and external APIs.
  - `redis/`: Logic for interacting with Redis.
  - `githubapi/`: Logic for interacting with GitHub API.
- **internal/infrastructure/**: Contains infrastructure-related code such as observability and configuration.
  - `config/`: Configuration management.
  - `observability/`: Metrics and logging setup.
- `docs/`: Contains the dynamically generated swagger specifications

This structure helps in organizing the codebase into clear, logical sections, making it easier to maintain and extend.
