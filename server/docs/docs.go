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
        "/login": {
            "post": {
                "description": "Выполняет вход в систему для пользователя с переданными данными в запросе",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Вход пользователя",
                "parameters": [
                    {
                        "description": "Email и пароль для пользователя",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserLoginData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешный вход в систему",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Невалидный JSON в теле запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Неверный пароль",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Пользователь не найден",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при создании пользователя",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Регистрирует нового пользователя с переданными данными в запросе",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "register"
                ],
                "summary": "Регистрация нового пользователя",
                "parameters": [
                    {
                        "description": "Email и пароль для пользователя",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserLoginData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Пользователь успешно зарегистрирован",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Невалидный JSON в теле запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "409": {
                        "description": "Пользователь уже существует",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при создании пользователя",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Выводит список тасок. В запросе нужно передать Bearer Token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Вывод списка тасок",
                "responses": {
                    "200": {
                        "description": "Список тасок",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Task"
                            }
                        }
                    },
                    "401": {
                        "description": "Необходимо выполнить вход для пользователя",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при создании пользователя",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Обновляет существующую таску. В запросе нужно передать Bearer Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Обновление таски",
                "parameters": [
                    {
                        "description": "Данные существующей таски",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TaskToUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное создание",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Невалидный JSON в теле запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Необходимо выполнить вход для пользователя",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при создании пользователя",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Создает новую таску. В запросе нужно передать Bearer Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Создание таски",
                "parameters": [
                    {
                        "description": "Данные новой таски",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TaskToCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Успешное создание",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Невалидный JSON в теле запроса",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Необходимо выполнить вход для пользователя",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Ошибка при создании пользователя",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Task": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/model.TaskStatus"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.TaskStatus": {
            "type": "integer",
            "enum": [
                1,
                2,
                3,
                4
            ],
            "x-enum-varnames": [
                "TaskStatusToDo",
                "TaskStatusInProgress",
                "TaskStatusDone",
                "TaskStatusDeleted"
            ]
        },
        "model.TaskToCreate": {
            "type": "object",
            "properties": {
                "status": {
                    "$ref": "#/definitions/model.TaskStatus"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.TaskToUpdate": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/model.TaskStatus"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.UserLoginData": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Токен для запросов в формате: Bearer {YOUR_TOKEN}",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "To-Do List API",
	Description:      "This is a sample server with To-Do List API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
