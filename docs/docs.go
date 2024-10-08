// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Jason Field",
            "url": "https://github.com/xorima"
        },
        "license": {
            "name": "MIT",
            "url": "https://github.com/xorima/webhook-bridge/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/webhook/github": {
            "post": {
                "description": "Once received, this will add them to the event stream for consumers.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Webhooks"
                ],
                "summary": "This API is there to receive the GitHub events.",
                "responses": {
                    "202": {
                        "description": "Successful Response",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/healthz": {
            "get": {
                "description": "Returns the health of this instance.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "This API is there to receive the health of this instance.",
                "responses": {
                    "200": {
                        "description": "Healthy",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "app.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    },
    "externalDocs": {
        "description": "GitHub",
        "url": "https://github.com/xorima/webhook-bridge"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Webhook Bridge API",
	Description:      "This is a bridge to receive various webhook events and publish them to a channel.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
