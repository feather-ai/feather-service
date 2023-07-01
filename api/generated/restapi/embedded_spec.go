// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "FeatherAI ServiceCore API",
    "version": "0.1.0"
  },
  "paths": {
    "/v1/api/system/completePublish": {
      "put": {
        "security": [
          {
            "ApiKeyAuth": []
          }
        ],
        "description": "Complete the Publish of a system to feather",
        "parameters": [
          {
            "name": "definition",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/completePublishRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/completePublishResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          },
          "500": {
            "description": "Generic Error response",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          }
        }
      }
    },
    "/v1/api/system/preparePublish": {
      "put": {
        "security": [
          {
            "ApiKeyAuth": []
          }
        ],
        "description": "Prepare to Publish a system to feather. The API will return upload URLs to use to upload the binary model files",
        "parameters": [
          {
            "name": "definition",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/preparePublishRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK. Includes the upload URLs for the binary models",
            "schema": {
              "$ref": "#/definitions/preparePublishResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          },
          "500": {
            "description": "Generic Error response",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          }
        }
      }
    },
    "/v1/client/apikey": {
      "get": {
        "security": [
          {
            "FeatherToken": []
          }
        ],
        "description": "Get the list of all API keys for the logged in user",
        "responses": {
          "200": {
            "description": "List of API keys",
            "schema": {
              "type": "array",
              "items": {
                "type": "object",
                "properties": {
                  "created": {
                    "type": "string",
                    "format": "date-time"
                  },
                  "key": {
                    "type": "string",
                    "format": "uuid"
                  },
                  "name": {
                    "type": "string"
                  }
                }
              }
            }
          },
          "400": {
            "description": "Error fetching API keys"
          }
        }
      },
      "put": {
        "security": [
          {
            "FeatherToken": []
          }
        ],
        "description": "Create a new API key for the logged in user",
        "parameters": [
          {
            "type": "string",
            "description": "Name of the API key",
            "name": "name",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "New API Key",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Error"
          }
        }
      },
      "delete": {
        "security": [
          {
            "FeatherToken": []
          }
        ],
        "description": "Delete/Revoke an API key. After this call the API key is immediately unavailable and",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "key",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok, key revoked"
          },
          "400": {
            "description": "Error deleting the key"
          }
        }
      }
    },
    "/v1/client/login": {
      "put": {
        "description": "Login a client",
        "parameters": [
          {
            "type": "string",
            "name": "X-AUTH0-TOKEN",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "403": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/v1/client/refresh": {
      "put": {
        "description": "Refresh a feather token that has recently expired",
        "parameters": [
          {
            "type": "string",
            "name": "X-FEATHER-TOKEN",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "name": "X-AUTH0-TOKEN",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK - includes new token",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "403": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/v1/client/uploads": {
      "get": {
        "security": [
          {
            "FeatherToken": []
          }
        ],
        "description": "Get the list of all pending upload requests for this user",
        "responses": {
          "200": {
            "description": "List of upload requests",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/uploadRequest"
              }
            }
          }
        }
      }
    },
    "/v1/debug/executeRequestSchema": {
      "get": {
        "security": [
          {
            "FeatherToken": []
          }
        ],
        "description": "Internal",
        "parameters": [
          {
            "type": "string",
            "description": "SystemID",
            "name": "systemId",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "New API Key",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/v1/health": {
      "get": {
        "description": "Health check endpoint",
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/v1/public/system": {
      "get": {
        "description": "Get the full description for a specific system, by username and systemname.",
        "parameters": [
          {
            "type": "string",
            "description": "Username of the system's creator",
            "name": "username",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Name of the system to fetch (must be used in conjunction with  username)",
            "name": "systemname",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Detailed information about a system",
            "schema": {
              "$ref": "#/definitions/systemDetails"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          },
          "404": {
            "description": "User or system not found",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          }
        }
      }
    },
    "/v1/public/system/{systemId}": {
      "get": {
        "description": "Get the full description for a specific system",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "The ID of the system",
            "name": "systemId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Detailed information about a system",
            "schema": {
              "$ref": "#/definitions/systemDetails"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          }
        }
      }
    },
    "/v1/public/system/{systemId}/step/{stepIndex}": {
      "put": {
        "description": "Execute a specific step of a system. This API will automatically run the latest published version of the system.",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "The ID of the system",
            "name": "systemId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "description": "The Index of the step to run, starting at 0.",
            "name": "stepIndex",
            "in": "path",
            "required": true
          },
          {
            "description": "The inputs expected by the step  (See system schema)",
            "name": "inputData",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/systemInputs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/runSystemResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/runSystemError"
            }
          },
          "413": {
            "description": "Payload too large"
          }
        }
      }
    },
    "/v1/public/systems": {
      "get": {
        "description": "Get a list of systems. This API can be called with a username query argument, in which case all the systems for that user will be returned. If no argument is given, all systems are returned.",
        "parameters": [
          {
            "type": "string",
            "description": "Name of the owner to fetch systems for",
            "name": "username",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "List of systems",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/systemInfo"
              }
            }
          },
          "404": {
            "description": "Nothing found",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          },
          "429": {
            "description": "Too many requests",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          }
        }
      }
    },
    "/v1/public/user/{userName}": {
      "get": {
        "description": "Get user info",
        "parameters": [
          {
            "type": "string",
            "description": "The user name of the user to lookup",
            "name": "userName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "User Information",
            "schema": {
              "$ref": "#/definitions/userInfo"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "completePublishRequest": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "completePublishResponse": {
      "type": "object",
      "properties": {
        "system": {
          "type": "string"
        },
        "user": {
          "type": "string"
        }
      }
    },
    "genericError": {
      "type": "string"
    },
    "loginResponse": {
      "type": "object",
      "properties": {
        "expireAt": {
          "type": "string",
          "format": "date-time"
        },
        "featherToken": {
          "type": "string"
        }
      }
    },
    "preparePublishRequest": {
      "type": "object",
      "required": [
        "name",
        "slug",
        "files",
        "schema"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "files": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "filename": {
                "type": "string"
              },
              "filetype": {
                "type": "string"
              }
            }
          }
        },
        "name": {
          "type": "string"
        },
        "schema": {
          "type": "string"
        },
        "slug": {
          "type": "string"
        }
      }
    },
    "preparePublishResponse": {
      "type": "object",
      "properties": {
        "expiryTime": {
          "type": "string",
          "format": "date-time"
        },
        "files": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "filename": {
                "type": "string"
              },
              "uploadUrl": {
                "type": "string"
              }
            }
          }
        },
        "id": {
          "type": "string"
        }
      }
    },
    "runSystemError": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "tty": {
          "type": "string"
        }
      }
    },
    "runSystemResponse": {
      "type": "object",
      "properties": {
        "outputLocation": {
          "type": "string"
        },
        "outputs": {
          "type": "array",
          "items": {
            "type": "object"
          }
        },
        "tty": {
          "type": "string"
        }
      }
    },
    "systemDetails": {
      "description": "Detailed information about a system, including the system schema",
      "type": "object",
      "properties": {
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "description": {
          "type": "string"
        },
        "files": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "created": {
                "type": "string",
                "format": "date-time"
              },
              "name": {
                "type": "string"
              },
              "type": {
                "type": "string"
              }
            }
          }
        },
        "id": {
          "type": "string"
        },
        "lastUpdated": {
          "type": "string",
          "format": "date-time"
        },
        "name": {
          "type": "string"
        },
        "num_steps": {
          "type": "integer"
        },
        "ownerId": {
          "type": "string"
        },
        "schema": {
          "type": "object"
        },
        "slug": {
          "type": "string"
        },
        "system_id": {
          "type": "string"
        }
      }
    },
    "systemInfo": {
      "description": "Summary information for a system",
      "type": "object",
      "properties": {
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "ownerId": {
          "type": "string"
        },
        "ownerName": {
          "type": "string"
        },
        "slug": {
          "type": "string"
        }
      }
    },
    "systemInputs": {
      "type": "string",
      "format": "binary"
    },
    "uploadRequest": {
      "type": "object",
      "required": [
        "uploadUrl",
        "id",
        "expiryTime"
      ],
      "properties": {
        "expiryTime": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "string",
          "maxLength": 36,
          "minLength": 36
        },
        "uploadUrl": {
          "type": "string",
          "minLength": 8
        }
      }
    },
    "userInfo": {
      "type": "object",
      "properties": {
        "userId": {
          "type": "string"
        },
        "userName": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "ApiKeyAuth": {
      "type": "apiKey",
      "name": "X-FEATHER-API-KEY",
      "in": "header"
    },
    "FeatherToken": {
      "type": "apiKey",
      "name": "X-FEATHER-TOKEN",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "FeatherAI ServiceCore API",
    "version": "0.1.0"
  },
  "paths": {
    "/v1/api/system/completePublish": {
      "put": {
        "security": [
          {
            "ApiKeyAuth": []
          }
        ],
        "description": "Complete the Publish of a system to feather",
        "parameters": [
          {
            "name": "definition",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/completePublishRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/completePublishResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          },
          "500": {
            "description": "Generic Error response",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          }
        }
      }
    },
    "/v1/api/system/preparePublish": {
      "put": {
        "security": [
          {
            "ApiKeyAuth": []
          }
        ],
        "description": "Prepare to Publish a system to feather. The API will return upload URLs to use to upload the binary model files",
        "parameters": [
          {
            "name": "definition",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/preparePublishRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK. Includes the upload URLs for the binary models",
            "schema": {
              "$ref": "#/definitions/preparePublishResponse"
            }
          },
          "400": {
            "description": "Bad Request",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          },
          "500": {
            "description": "Generic Error response",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          }
        }
      }
    },
    "/v1/client/apikey": {
      "get": {
        "security": [
          {
            "FeatherToken": []
          }
        ],
        "description": "Get the list of all API keys for the logged in user",
        "responses": {
          "200": {
            "description": "List of API keys",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/GetV1ClientApikeyOKBodyItems0"
              }
            }
          },
          "400": {
            "description": "Error fetching API keys"
          }
        }
      },
      "put": {
        "security": [
          {
            "FeatherToken": []
          }
        ],
        "description": "Create a new API key for the logged in user",
        "parameters": [
          {
            "type": "string",
            "description": "Name of the API key",
            "name": "name",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "New API Key",
            "schema": {
              "type": "string"
            }
          },
          "400": {
            "description": "Error"
          }
        }
      },
      "delete": {
        "security": [
          {
            "FeatherToken": []
          }
        ],
        "description": "Delete/Revoke an API key. After this call the API key is immediately unavailable and",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "name": "key",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok, key revoked"
          },
          "400": {
            "description": "Error deleting the key"
          }
        }
      }
    },
    "/v1/client/login": {
      "put": {
        "description": "Login a client",
        "parameters": [
          {
            "type": "string",
            "name": "X-AUTH0-TOKEN",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "403": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/v1/client/refresh": {
      "put": {
        "description": "Refresh a feather token that has recently expired",
        "parameters": [
          {
            "type": "string",
            "name": "X-FEATHER-TOKEN",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "name": "X-AUTH0-TOKEN",
            "in": "header",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK - includes new token",
            "schema": {
              "$ref": "#/definitions/loginResponse"
            }
          },
          "403": {
            "description": "Unauthorized"
          }
        }
      }
    },
    "/v1/client/uploads": {
      "get": {
        "security": [
          {
            "FeatherToken": []
          }
        ],
        "description": "Get the list of all pending upload requests for this user",
        "responses": {
          "200": {
            "description": "List of upload requests",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/uploadRequest"
              }
            }
          }
        }
      }
    },
    "/v1/debug/executeRequestSchema": {
      "get": {
        "security": [
          {
            "FeatherToken": []
          }
        ],
        "description": "Internal",
        "parameters": [
          {
            "type": "string",
            "description": "SystemID",
            "name": "systemId",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "New API Key",
            "schema": {
              "type": "object"
            }
          },
          "400": {
            "description": "Error"
          }
        }
      }
    },
    "/v1/health": {
      "get": {
        "description": "Health check endpoint",
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/v1/public/system": {
      "get": {
        "description": "Get the full description for a specific system, by username and systemname.",
        "parameters": [
          {
            "type": "string",
            "description": "Username of the system's creator",
            "name": "username",
            "in": "query"
          },
          {
            "type": "string",
            "description": "Name of the system to fetch (must be used in conjunction with  username)",
            "name": "systemname",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Detailed information about a system",
            "schema": {
              "$ref": "#/definitions/systemDetails"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          },
          "404": {
            "description": "User or system not found",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          }
        }
      }
    },
    "/v1/public/system/{systemId}": {
      "get": {
        "description": "Get the full description for a specific system",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "The ID of the system",
            "name": "systemId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Detailed information about a system",
            "schema": {
              "$ref": "#/definitions/systemDetails"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          }
        }
      }
    },
    "/v1/public/system/{systemId}/step/{stepIndex}": {
      "put": {
        "description": "Execute a specific step of a system. This API will automatically run the latest published version of the system.",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "The ID of the system",
            "name": "systemId",
            "in": "path",
            "required": true
          },
          {
            "type": "integer",
            "description": "The Index of the step to run, starting at 0.",
            "name": "stepIndex",
            "in": "path",
            "required": true
          },
          {
            "description": "The inputs expected by the step  (See system schema)",
            "name": "inputData",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/systemInputs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/runSystemResponse"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/runSystemError"
            }
          },
          "413": {
            "description": "Payload too large"
          }
        }
      }
    },
    "/v1/public/systems": {
      "get": {
        "description": "Get a list of systems. This API can be called with a username query argument, in which case all the systems for that user will be returned. If no argument is given, all systems are returned.",
        "parameters": [
          {
            "type": "string",
            "description": "Name of the owner to fetch systems for",
            "name": "username",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "List of systems",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/systemInfo"
              }
            }
          },
          "404": {
            "description": "Nothing found",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          },
          "429": {
            "description": "Too many requests",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          }
        }
      }
    },
    "/v1/public/user/{userName}": {
      "get": {
        "description": "Get user info",
        "parameters": [
          {
            "type": "string",
            "description": "The user name of the user to lookup",
            "name": "userName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "User Information",
            "schema": {
              "$ref": "#/definitions/userInfo"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/genericError"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "GetV1ClientApikeyOKBodyItems0": {
      "type": "object",
      "properties": {
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "key": {
          "type": "string",
          "format": "uuid"
        },
        "name": {
          "type": "string"
        }
      }
    },
    "PreparePublishRequestFilesItems0": {
      "type": "object",
      "properties": {
        "filename": {
          "type": "string"
        },
        "filetype": {
          "type": "string"
        }
      }
    },
    "PreparePublishResponseFilesItems0": {
      "type": "object",
      "properties": {
        "filename": {
          "type": "string"
        },
        "uploadUrl": {
          "type": "string"
        }
      }
    },
    "SystemDetailsFilesItems0": {
      "type": "object",
      "properties": {
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "name": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "completePublishRequest": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "completePublishResponse": {
      "type": "object",
      "properties": {
        "system": {
          "type": "string"
        },
        "user": {
          "type": "string"
        }
      }
    },
    "genericError": {
      "type": "string"
    },
    "loginResponse": {
      "type": "object",
      "properties": {
        "expireAt": {
          "type": "string",
          "format": "date-time"
        },
        "featherToken": {
          "type": "string"
        }
      }
    },
    "preparePublishRequest": {
      "type": "object",
      "required": [
        "name",
        "slug",
        "files",
        "schema"
      ],
      "properties": {
        "description": {
          "type": "string"
        },
        "files": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/PreparePublishRequestFilesItems0"
          }
        },
        "name": {
          "type": "string"
        },
        "schema": {
          "type": "string"
        },
        "slug": {
          "type": "string"
        }
      }
    },
    "preparePublishResponse": {
      "type": "object",
      "properties": {
        "expiryTime": {
          "type": "string",
          "format": "date-time"
        },
        "files": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/PreparePublishResponseFilesItems0"
          }
        },
        "id": {
          "type": "string"
        }
      }
    },
    "runSystemError": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "tty": {
          "type": "string"
        }
      }
    },
    "runSystemResponse": {
      "type": "object",
      "properties": {
        "outputLocation": {
          "type": "string"
        },
        "outputs": {
          "type": "array",
          "items": {
            "type": "object"
          }
        },
        "tty": {
          "type": "string"
        }
      }
    },
    "systemDetails": {
      "description": "Detailed information about a system, including the system schema",
      "type": "object",
      "properties": {
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "description": {
          "type": "string"
        },
        "files": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/SystemDetailsFilesItems0"
          }
        },
        "id": {
          "type": "string"
        },
        "lastUpdated": {
          "type": "string",
          "format": "date-time"
        },
        "name": {
          "type": "string"
        },
        "num_steps": {
          "type": "integer"
        },
        "ownerId": {
          "type": "string"
        },
        "schema": {
          "type": "object"
        },
        "slug": {
          "type": "string"
        },
        "system_id": {
          "type": "string"
        }
      }
    },
    "systemInfo": {
      "description": "Summary information for a system",
      "type": "object",
      "properties": {
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "ownerId": {
          "type": "string"
        },
        "ownerName": {
          "type": "string"
        },
        "slug": {
          "type": "string"
        }
      }
    },
    "systemInputs": {
      "type": "string",
      "format": "binary"
    },
    "uploadRequest": {
      "type": "object",
      "required": [
        "uploadUrl",
        "id",
        "expiryTime"
      ],
      "properties": {
        "expiryTime": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "type": "string",
          "maxLength": 36,
          "minLength": 36
        },
        "uploadUrl": {
          "type": "string",
          "minLength": 8
        }
      }
    },
    "userInfo": {
      "type": "object",
      "properties": {
        "userId": {
          "type": "string"
        },
        "userName": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "ApiKeyAuth": {
      "type": "apiKey",
      "name": "X-FEATHER-API-KEY",
      "in": "header"
    },
    "FeatherToken": {
      "type": "apiKey",
      "name": "X-FEATHER-TOKEN",
      "in": "header"
    }
  }
}`))
}
