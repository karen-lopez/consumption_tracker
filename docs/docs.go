// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/consumption": {
            "get": {
                "description": "Get energy consumption data for specific meters within a date range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "consumption"
                ],
                "summary": "Get energy consumption data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Comma-separated list of meter IDs",
                        "name": "meters_ids",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Start date in YYYY-MM-DD format",
                        "name": "start_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "End date in YYYY-MM-DD format",
                        "name": "end_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Period type (daily, weekly, monthly)",
                        "name": "kind_period",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ConsumptionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.ConsumptionResponse": {
            "type": "object",
            "properties": {
                "data_graph": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.MeterData"
                    }
                },
                "period": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "dtos.MeterData": {
            "type": "object",
            "properties": {
                "active": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "address": {
                    "type": "string"
                },
                "exported": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "meter_id": {
                    "type": "integer"
                },
                "reactive_capacitive": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "reactive_inductive": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "handlers.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "error message"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
