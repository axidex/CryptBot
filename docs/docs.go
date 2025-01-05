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
        "/api/v1/a5": {
            "post": {
                "description": "A5",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "A5",
                "parameters": [
                    {
                        "type": "string",
                        "default": "010011",
                        "description": "a",
                        "name": "a",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "1001011",
                        "description": "b",
                        "name": "b",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "10101011110",
                        "description": "c",
                        "name": "c",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "f34ff541f04fed5ee04fff59fa50",
                        "description": "Text",
                        "name": "text",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/aes": {
            "post": {
                "description": "Aes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "Aes",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Перу",
                        "description": "text",
                        "name": "text",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "Ключ",
                        "description": "key",
                        "name": "key",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/blowfish": {
            "post": {
                "description": "Blowfish",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "Blowfish",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 4,
                        "description": "Left block",
                        "name": "l",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 9,
                        "description": "Right block",
                        "name": "r",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Key 1",
                        "name": "k1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 5,
                        "description": "Key 2",
                        "name": "k2",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 3,
                        "description": "Key 3",
                        "name": "k3",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 4,
                        "description": "Key 4",
                        "name": "k4",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 5,
                        "description": "Key 5",
                        "name": "k5",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/des": {
            "post": {
                "description": "Des",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "Des",
                "parameters": [
                    {
                        "type": "string",
                        "default": "гана",
                        "description": "text",
                        "name": "text",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "куб",
                        "description": "key",
                        "name": "key",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/des3": {
            "post": {
                "description": "Des3",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "Des3",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 5,
                        "description": "Left block",
                        "name": "l",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "Right block",
                        "name": "r",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 3,
                        "description": "Key 1",
                        "name": "k1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 7,
                        "description": "Key 2",
                        "name": "k2",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 5,
                        "description": "Key 3",
                        "name": "k3",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/diffie": {
            "post": {
                "description": "Diffie",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "Diffie",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 17,
                        "description": "n",
                        "name": "n",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 11,
                        "description": "q",
                        "name": "q",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 2,
                        "description": "x",
                        "name": "x",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 3,
                        "description": "y",
                        "name": "y",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/diffie3": {
            "post": {
                "description": "Diffie3",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "Diffie3",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "a",
                        "name": "a",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 2,
                        "description": "b",
                        "name": "b",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 3,
                        "description": "c",
                        "name": "c",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 197,
                        "description": "g",
                        "name": "g",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 53,
                        "description": "p",
                        "name": "p",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/elgamal": {
            "post": {
                "description": "Elgamal",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "Elgamal",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 19,
                        "description": "p",
                        "name": "p",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 3,
                        "description": "g",
                        "name": "g",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 7,
                        "description": "k",
                        "name": "k",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 8,
                        "description": "x",
                        "name": "x",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 5,
                        "description": "m",
                        "name": "m",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/enigma": {
            "post": {
                "description": "Enigma",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "Enigma",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 315,
                        "description": "rotor",
                        "name": "rot",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "C",
                        "description": "reflector",
                        "name": "ref",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "A-U,D-Y",
                        "description": "patch panel",
                        "name": "pp",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "MJL",
                        "description": "text",
                        "name": "str",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/feistel": {
            "post": {
                "description": "Feistel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "Feistel",
                "parameters": [
                    {
                        "type": "string",
                        "default": "047b16c2",
                        "description": "text",
                        "name": "data",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "67e99b3c",
                        "description": "key",
                        "name": "keys",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/hash": {
            "post": {
                "description": "Hash",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "Hash",
                "parameters": [
                    {
                        "type": "string",
                        "default": "сотовыйтелефон",
                        "description": "word",
                        "name": "word",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "7",
                        "description": "h0",
                        "name": "h0",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 11,
                        "description": "p",
                        "name": "p",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/invMix": {
            "post": {
                "description": "InvMix",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "InvMix",
                "parameters": [
                    {
                        "type": "string",
                        "default": "h",
                        "description": "char",
                        "name": "char",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/md5": {
            "post": {
                "description": "Md5",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "Md5",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 5,
                        "description": "a",
                        "name": "a",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "b",
                        "name": "b",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 9,
                        "description": "c",
                        "name": "c",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "d",
                        "name": "d",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 6,
                        "description": "m",
                        "name": "m",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "k",
                        "name": "k",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/rsa": {
            "post": {
                "description": "Rsa",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "Rsa",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 3,
                        "description": "p",
                        "name": "p",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 7,
                        "description": "q",
                        "name": "q",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 5,
                        "description": "e",
                        "name": "e",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 101,
                        "description": "x",
                        "name": "x",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/api/v1/sBlock": {
            "post": {
                "description": "SBlock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Crypt"
                ],
                "summary": "SBlock",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 11,
                        "description": "a",
                        "name": "a",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 17,
                        "description": "b",
                        "name": "b",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 15,
                        "description": "c",
                        "name": "c",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 12,
                        "description": "z0",
                        "name": "z0",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 7,
                        "description": "n",
                        "name": "n",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Processed successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
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
