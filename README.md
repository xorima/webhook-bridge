# Webhook Bridge

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



## Code decisions

For this app it is decided that the models for the app should sit in with their appropriate package, so each package will
have it's models locally and provide interfaces for others to consume. 
