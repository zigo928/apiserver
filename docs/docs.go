// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2018-08-03 10:53:45.530222699 +0800 CST m=+0.089619603

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "contact": {},
        "license": {}
    },
    "paths": {
        "/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Login generates the authentication token",
                "parameters": [
                    {
                        "description": "Username",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    },
                    {
                        "description": "Password",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"token\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ\"}}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "List users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "list the users in the database.",
                "parameters": [
                    {
                        "description": "List users",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.ListRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"totalCount\":1,\"userList\":[{\"id\":0,\"username\":\"admin\",\"random\":\"user 'admin' get random string 'EnqntiSig'\",\"password\":\"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG\",\"createdAt\":\"2018-05-28 00:25:33\",\"updatedAt\":\"2018-05-28 00:25:33\"}]}}",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.SwaggerListResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Add new user to the database.",
                "parameters": [
                    {
                        "description": "Create a new user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"username\":\"kong\"}}",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/user.CreateResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "put": {
                "description": "Update an user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update a exist user account info.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The user's database id num",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The user info",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.UserModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":null}",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "delete an user by the user identifier.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The user's database id index num",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":null}",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/user/{username}": {
            "get": {
                "description": "Get an user by username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get an user by the user identifier.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":0,\"message\":\"OK\",\"data\":{\"username\":\"kong\",\"password\":\"$2a$10$E0kwtmtLZbwW/bDQ8qI8e.eHPqhQOW9tvjwpyo/p05f/f4Qvr3OmS\"}}",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.UserModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "sayHello": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.UserModel": {
            "type": "object",
            "required": [
                "username",
                "password"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.CreateRequest": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.CreateResponse": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "user.ListRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.SwaggerListResponse": {
            "type": "object",
            "properties": {
                "totalCount": {
                    "type": "integer"
                },
                "userList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.UserInfo"
                    }
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
