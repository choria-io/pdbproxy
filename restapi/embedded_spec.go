package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API to do Choria Discovery via REST services",
    "title": "Choria Discovery API",
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/backup": {
      "get": {
        "produces": [
          "application/octet-stream"
        ],
        "summary": "Produce a DB backup",
        "responses": {
          "200": {
            "$ref": "#/responses/success"
          },
          "500": {
            "$ref": "#/responses/error"
          }
        }
      }
    },
    "/discover": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Perform a discovery request",
        "parameters": [
          {
            "description": "Filter description",
            "name": "request",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/discoveryRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/nodes"
          },
          "400": {
            "$ref": "#/responses/error"
          }
        }
      }
    },
    "/set": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Create a new set",
        "parameters": [
          {
            "description": "Node set to store",
            "name": "set",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/set"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/success"
          },
          "400": {
            "$ref": "#/responses/error"
          },
          "500": {
            "$ref": "#/responses/error"
          }
        }
      }
    },
    "/set/{set}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "summary": "Retrieves the query or nodes for a set",
        "parameters": [
          {
            "type": "string",
            "description": "Node set to retrieve",
            "name": "set",
            "in": "path",
            "required": true
          },
          {
            "type": "boolean",
            "default": false,
            "description": "Also include the matching nodes if the set is a PQL query",
            "name": "discover",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/set"
          },
          "404": {
            "description": "Not found"
          },
          "500": {
            "$ref": "#/responses/error"
          }
        }
      },
      "put": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "appilcation/json"
        ],
        "summary": "Update a set",
        "parameters": [
          {
            "type": "string",
            "description": "Node set to retrieve",
            "name": "set",
            "in": "path",
            "required": true
          },
          {
            "description": "Set Properties",
            "name": "newSet",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/set"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/success"
          },
          "404": {
            "description": "Not found"
          },
          "500": {
            "$ref": "#/responses/error"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "summary": "Delete a set",
        "parameters": [
          {
            "type": "string",
            "description": "Node set to delete",
            "name": "set",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/success"
          },
          "404": {
            "description": "Not found"
          },
          "500": {
            "$ref": "#/responses/error"
          }
        }
      }
    },
    "/sets": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Retrieve list of known sets",
        "responses": {
          "200": {
            "$ref": "#/responses/sets"
          }
        }
      }
    }
  },
  "definitions": {
    "agentsFilter": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/wordOrRegex"
      }
    },
    "classesFilter": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/wordOrRegex"
      }
    },
    "collectiveFilter": {
      "$ref": "#/definitions/word"
    },
    "discoveryRequest": {
      "type": "object",
      "properties": {
        "agents": {
          "$ref": "#/definitions/agentsFilter"
        },
        "classes": {
          "$ref": "#/definitions/classesFilter"
        },
        "collective": {
          "$ref": "#/definitions/collectiveFilter"
        },
        "facts": {
          "$ref": "#/definitions/factsFilter"
        },
        "identities": {
          "$ref": "#/definitions/identitiesFilter"
        },
        "query": {
          "description": "PQL Query",
          "type": "string",
          "minLength": 1
        }
      }
    },
    "discoverySuccessModel": {
      "type": "object",
      "properties": {
        "nodes": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "status": {
          "description": "HTTP Status Code",
          "type": "integer"
        }
      },
      "additionalProperties": false
    },
    "errorModel": {
      "type": "object",
      "properties": {
        "detail": {
          "description": "Extra details about the error",
          "type": "string"
        },
        "message": {
          "description": "Short description of the problem",
          "type": "string"
        },
        "status": {
          "description": "HTTP Status Code",
          "type": "integer"
        }
      },
      "additionalProperties": false
    },
    "factFilter": {
      "type": "object",
      "required": [
        "fact",
        "operator",
        "value"
      ],
      "properties": {
        "fact": {
          "$ref": "#/definitions/word"
        },
        "operator": {
          "type": "string",
          "enum": [
            "==",
            "=~",
            "!=",
            "\u003e=",
            "\u003e",
            "\u003c=",
            "\u003c"
          ]
        },
        "value": {
          "$ref": "#/definitions/wordOrRegex"
        }
      }
    },
    "factsFilter": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/factFilter"
      }
    },
    "identitiesFilter": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/wordOrRegex"
      }
    },
    "set": {
      "type": "object",
      "required": [
        "set",
        "query"
      ],
      "properties": {
        "nodes": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "query": {
          "type": "string",
          "minLength": 1
        },
        "set": {
          "$ref": "#/definitions/word"
        }
      }
    },
    "sets": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/word"
      }
    },
    "successModel": {
      "type": "object",
      "properties": {
        "detail": {
          "type": "string"
        },
        "status": {
          "description": "HTTP Status Code",
          "type": "integer"
        }
      }
    },
    "word": {
      "type": "string",
      "minLength": 1,
      "pattern": "^[a-zA-Z0-9_\\-\\.]+$"
    },
    "wordOrRegex": {
      "type": "string",
      "minLength": 1,
      "pattern": "^((\\\\/[a-zA-Z0-9\\\\.\\\\-_]+\\/)|([a-zA-Z][a-zA-Z0-9\\\\.\\\\-_]+))$"
    }
  },
  "responses": {
    "error": {
      "description": "Standard Error Format",
      "schema": {
        "$ref": "#/definitions/errorModel"
      }
    },
    "nodes": {
      "description": "Basic successful discovery request",
      "schema": {
        "$ref": "#/definitions/discoverySuccessModel"
      }
    },
    "set": {
      "description": "Node Set",
      "schema": {
        "$ref": "#/definitions/set"
      }
    },
    "sets": {
      "description": "Known Sets",
      "schema": {
        "$ref": "#/definitions/sets"
      }
    },
    "success": {
      "description": "Basic successful request",
      "schema": {
        "$ref": "#/definitions/successModel"
      }
    }
  }
}`))
}
