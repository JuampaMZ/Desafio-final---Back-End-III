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
        "/dentistas": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dentista"
                ],
                "summary": "Listar todos los dentistas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Dentist"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dentista"
                ],
                "summary": "Agregar un nuevo dentista",
                "parameters": [
                    {
                        "description": "Dentista",
                        "name": "dentista",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Dentist"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Dentist"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/dentistas/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dentista"
                ],
                "summary": "Obtener un dentista por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del dentista",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Dentist"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dentista"
                ],
                "summary": "Actualizar un dentista",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del dentista",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Dentista",
                        "name": "dentista",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Dentist"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Dentist"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dentista"
                ],
                "summary": "Eliminar un dentista",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del dentista",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Dentista"
                ],
                "summary": "Actualizar algunos campos de un dentista",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del dentista",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Dentist"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/pacientes": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Listar todos los pacientes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Patient"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Agregar un nuevo paciente",
                "parameters": [
                    {
                        "description": "Paciente",
                        "name": "paciente",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    }
                }
            }
        },
        "/pacientes/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Obtener un paciente por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del paciente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Actualizar un paciente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del paciente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Paciente",
                        "name": "paciente",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Eliminar un paciente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del paciente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Paciente"
                ],
                "summary": "Actualizar algunos campos de un paciente",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del paciente",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Patient"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/turnos": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Listar todos los turnos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Appointment"
                            }
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Agregar un nuevo turno",
                "parameters": [
                    {
                        "description": "Turno",
                        "name": "turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Appointment"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Appointment"
                        }
                    }
                }
            }
        },
        "/turnos/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Obtener un turno por ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del turno",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Appointment"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Actualizar un turno",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del turno",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Turno",
                        "name": "turno",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Appointment"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Appointment"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Eliminar un turno",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del turno",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Turno"
                ],
                "summary": "Actualizar algunos campos de un turno",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del turno",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Appointment"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Appointment": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "dentist_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "patient_id": {
                    "type": "integer"
                },
                "time": {
                    "type": "string"
                }
            }
        },
        "models.Dentist": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "license": {
                    "type": "string"
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "Código del error (por ejemplo, 400, 404)",
                    "type": "integer"
                },
                "message": {
                    "description": "Mensaje descriptivo del error",
                    "type": "string"
                }
            }
        },
        "models.Patient": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "dni": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "registration_date": {
                    "type": "string"
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
