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
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple API",
    "title": "Simple Inventory API",
    "contact": {
      "email": "you@your-company.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "virtserver.swaggerhub.com",
  "basePath": "/nxmmob/helloworld/1.0.0",
  "paths": {
    "/inventory": {
      "get": {
        "description": "By passing in the appropriate options, you can search for\navailable inventory in the system\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "developers"
        ],
        "summary": "searches inventory",
        "operationId": "searchInventory",
        "parameters": [
          {
            "type": "string",
            "description": "pass an optional search string for looking up inventory",
            "name": "searchString",
            "in": "query"
          },
          {
            "type": "integer",
            "format": "int32",
            "description": "number of records to skip for pagination",
            "name": "skip",
            "in": "query"
          },
          {
            "maximum": 50,
            "type": "integer",
            "format": "int32",
            "description": "maximum number of records to return",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "search results matching criteria",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/InventoryItem"
              }
            }
          },
          "400": {
            "description": "bad input parameter"
          }
        }
      },
      "post": {
        "description": "Adds an item to the system",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "admins"
        ],
        "summary": "adds an inventory item",
        "operationId": "addInventory",
        "parameters": [
          {
            "description": "Inventory item to add",
            "name": "inventoryItem",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/InventoryItem"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "item created"
          },
          "400": {
            "description": "invalid input, object invalid"
          },
          "409": {
            "description": "an existing item already exists"
          }
        }
      }
    }
  },
  "definitions": {
    "InventoryItem": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "string",
          "example": "d290f1ee-6c54-4b01-90e6-d701748f0851"
        },
        "manufacturer": {
          "$ref": "#/definitions/Manufacturer"
        },
        "name": {
          "type": "string",
          "example": "Widget Adapter"
        },
        "releaseDate": {
          "type": "string",
          "format": "date-time",
          "example": "2016-08-29T09:12:33.001Z"
        }
      }
    },
    "Manufacturer": {
      "required": [
        "name"
      ],
      "properties": {
        "homePage": {
          "type": "string",
          "format": "url",
          "example": "https://www.acme-corp.com"
        },
        "name": {
          "type": "string",
          "example": "ACME Corporation"
        },
        "phone": {
          "type": "string",
          "example": "408-867-5309"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Secured Admin-only calls",
      "name": "admins"
    },
    {
      "description": "Operations available to regular developers",
      "name": "developers"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple API",
    "title": "Simple Inventory API",
    "contact": {
      "email": "you@your-company.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "virtserver.swaggerhub.com",
  "basePath": "/nxmmob/helloworld/1.0.0",
  "paths": {
    "/inventory": {
      "get": {
        "description": "By passing in the appropriate options, you can search for\navailable inventory in the system\n",
        "produces": [
          "application/json"
        ],
        "tags": [
          "developers"
        ],
        "summary": "searches inventory",
        "operationId": "searchInventory",
        "parameters": [
          {
            "type": "string",
            "description": "pass an optional search string for looking up inventory",
            "name": "searchString",
            "in": "query"
          },
          {
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "description": "number of records to skip for pagination",
            "name": "skip",
            "in": "query"
          },
          {
            "maximum": 50,
            "minimum": 0,
            "type": "integer",
            "format": "int32",
            "description": "maximum number of records to return",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "search results matching criteria",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/InventoryItem"
              }
            }
          },
          "400": {
            "description": "bad input parameter"
          }
        }
      },
      "post": {
        "description": "Adds an item to the system",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "admins"
        ],
        "summary": "adds an inventory item",
        "operationId": "addInventory",
        "parameters": [
          {
            "description": "Inventory item to add",
            "name": "inventoryItem",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/InventoryItem"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "item created"
          },
          "400": {
            "description": "invalid input, object invalid"
          },
          "409": {
            "description": "an existing item already exists"
          }
        }
      }
    }
  },
  "definitions": {
    "InventoryItem": {
      "type": "object",
      "required": [
        "id"
      ],
      "properties": {
        "id": {
          "type": "string",
          "example": "d290f1ee-6c54-4b01-90e6-d701748f0851"
        },
        "manufacturer": {
          "$ref": "#/definitions/Manufacturer"
        },
        "name": {
          "type": "string",
          "example": "Widget Adapter"
        },
        "releaseDate": {
          "type": "string",
          "format": "date-time",
          "example": "2016-08-29T09:12:33.001Z"
        }
      }
    },
    "Manufacturer": {
      "required": [
        "name"
      ],
      "properties": {
        "homePage": {
          "type": "string",
          "format": "url",
          "example": "https://www.acme-corp.com"
        },
        "name": {
          "type": "string",
          "example": "ACME Corporation"
        },
        "phone": {
          "type": "string",
          "example": "408-867-5309"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Secured Admin-only calls",
      "name": "admins"
    },
    {
      "description": "Operations available to regular developers",
      "name": "developers"
    }
  ]
}`))
}
