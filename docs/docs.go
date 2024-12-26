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
        "/questions": {
            "get": {
                "description": "Get list of questions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "questions"
                ],
                "summary": "Get questions",
                "responses": {
                    "200": {
                        "description": "Questions list response",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/main.LeetCodeQuestions"
                                }
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new question",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "questions"
                ],
                "summary": "Create question",
                "parameters": [
                    {
                        "description": "Question object",
                        "name": "question",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.LeetCodeQuestions"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Question created",
                        "schema": {
                            "$ref": "#/definitions/main.LeetCodeQuestions"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get list of users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get users",
                "responses": {
                    "200": {
                        "description": "User list response",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/main.User"
                                }
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "User object",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created",
                        "schema": {
                            "$ref": "#/definitions/main.User"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/main.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Difficulty": {
            "type": "string",
            "enum": [
                "Easy",
                "Medium",
                "Hard"
            ],
            "x-enum-varnames": [
                "Easy",
                "Medium",
                "Hard"
            ]
        },
        "main.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "main.LeetCodeQuestions": {
            "type": "object",
            "properties": {
                "difficulty": {
                    "$ref": "#/definitions/main.Difficulty"
                },
                "id": {
                    "type": "string"
                },
                "lastCompletedTime": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nextDueTime": {
                    "type": "string"
                },
                "notes": {
                    "type": "string"
                },
                "pattern": {
                    "type": "string"
                }
            }
        },
        "main.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "LeetCode Questions API",
	Description:      "API for managing LeetCode questions and users",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
