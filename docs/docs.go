// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/segments": {
            "post": {
                "description": "Create a new segment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "segments"
                ],
                "summary": "Create segment",
                "parameters": [
                    {
                        "description": "Segment data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Segment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Segment"
                        }
                    }
                }
            }
        },
        "/api/segments/{slug}": {
            "delete": {
                "description": "Delete a segment by slug",
                "tags": [
                    "segments"
                ],
                "summary": "Delete segment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Segment slug",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.statusResponse"
                        }
                    }
                }
            }
        },
        "/api/users/{user_id}/segments": {
            "get": {
                "description": "Get segments of a user",
                "tags": [
                    "users"
                ],
                "summary": "Get user segments",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Update segments for a user",
                "tags": [
                    "users"
                ],
                "summary": "Update user segments",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Segments data",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserSegmentsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.statusResponse"
                        }
                    }
                }
            }
        },
        "/api/users/{user_id}/segments/add_with_ttl": {
            "post": {
                "description": "Add user to a segment with a specified TTL",
                "tags": [
                    "users"
                ],
                "summary": "Add user to segment with TTL",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User to Segment with TTL request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserToSegmentWithTTLRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.statusResponse"
                        }
                    }
                }
            }
        },
        "/api/users/{user_id}/segments/history": {
            "get": {
                "description": "Generate a CSV report of user segment history for a specific user within a given year and month",
                "tags": [
                    "users"
                ],
                "summary": "Generate user segment history report",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Year",
                        "name": "year",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Month",
                        "name": "month",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "CSV report",
                        "schema": {
                            "type": "file"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.statusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "models.Segment": {
            "type": "object",
            "required": [
                "slug"
            ],
            "properties": {
                "slug": {
                    "type": "string"
                }
            }
        },
        "models.UserSegmentsRequest": {
            "type": "object",
            "properties": {
                "add_segments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "remove_segments": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "models.UserToSegmentWithTTLRequest": {
            "type": "object",
            "properties": {
                "segment_slug": {
                    "type": "string"
                },
                "ttl_minutes": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Avito Segment API",
	Description:      "API Server for AvitoSegment Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
