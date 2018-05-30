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
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "denn nur lebendiges food tut gut",
    "title": "Savood",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0"
  },
  "basePath": "/v2/",
  "paths": {
    "/feed": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "feed"
        ],
        "summary": "Display a feed",
        "operationId": "getFeed",
        "parameters": [
          {
            "type": "string",
            "name": "location",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Feed"
            }
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
          }
        }
      }
    },
    "/healthcheck": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "health"
        ],
        "summary": "Check Server status",
        "operationId": "healthcheckGet",
        "responses": {
          "200": {
            "description": "Server is healthy"
          },
          "501": {
            "description": "Server is unhealthy",
            "schema": {
              "$ref": "#/definitions/ErrorModel"
            }
          }
        }
      }
    },
    "/message": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "message"
        ],
        "summary": "Add a new message",
        "operationId": "createNewMessage",
        "parameters": [
          {
            "description": "Message that needs to be added",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Message"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Message"
            }
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
          }
        }
      }
    },
    "/message/{id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "message"
        ],
        "summary": "Display a message",
        "operationId": "getMessageById",
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Message"
            }
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "message"
        ],
        "summary": "Delete a message",
        "operationId": "deleteMessageById",
        "responses": {
          "204": {
            "description": "Removed; No response."
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
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
          "message"
        ],
        "summary": "Update a message",
        "operationId": "updateMessageById",
        "parameters": [
          {
            "description": "New parameters of the message",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Message"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Accepted; No response."
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/offering": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "offering"
        ],
        "summary": "Add a new offering",
        "operationId": "createNewOffering",
        "parameters": [
          {
            "description": "Offering that needs to be added",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Offering"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Offering"
            }
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
          }
        }
      }
    },
    "/offering/{id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "offering"
        ],
        "summary": "Display an offering",
        "operationId": "getOfferingById",
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Offering"
            }
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "offering"
        ],
        "summary": "Delete an offering",
        "operationId": "deleteOfferingById",
        "responses": {
          "204": {
            "description": "Removed; No response."
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
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
          "offering"
        ],
        "summary": "Update an offering",
        "operationId": "updateOfferingById",
        "parameters": [
          {
            "description": "New parameters of the offering",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Offering"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Accepted; No response."
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/profile": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "summary": "Add a new profile",
        "operationId": "createNewProfile",
        "parameters": [
          {
            "description": "Profile that needs to be added",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
          }
        }
      }
    },
    "/profile/{id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "summary": "Display a profile",
        "operationId": "getProfileById",
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "summary": "Delete a profile",
        "operationId": "deleteProfileById",
        "responses": {
          "204": {
            "description": "Removed; No response."
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
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
          "profile"
        ],
        "summary": "Update a profile",
        "operationId": "updateProfileById",
        "parameters": [
          {
            "description": "New parameters of the profile",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Accepted; No response."
          },
          "404": {
            "$ref": "#/responses/InvalidParameterInput"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Comment": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string"
        },
        "creator-id": {
          "type": "integer",
          "format": "int64"
        },
        "id": {
          "type": "string"
        },
        "time": {
          "type": "string",
          "format": "date-time"
        }
      },
      "xml": {
        "name": "Comment"
      }
    },
    "ErrorModel": {
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Feed": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Offering"
      },
      "xml": {
        "name": "Feed"
      }
    },
    "InvalidParameterInput": {
      "type": "object",
      "required": [
        "attribute",
        "message"
      ],
      "properties": {
        "attribute": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Message": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string"
        },
        "from-id": {
          "type": "integer",
          "format": "int64"
        },
        "id": {
          "type": "string"
        },
        "important": {
          "type": "boolean"
        },
        "offering-id": {
          "type": "integer",
          "format": "int64"
        },
        "time": {
          "type": "string",
          "format": "date-time"
        },
        "to-id": {
          "type": "integer",
          "format": "int64"
        }
      },
      "xml": {
        "name": "Message"
      }
    },
    "Offering": {
      "type": "object",
      "properties": {
        "avatar-url": {
          "type": "string"
        },
        "best-by-date": {
          "type": "string",
          "format": "date"
        },
        "comments": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Comment"
          }
        },
        "creator-id": {
          "type": "integer",
          "format": "int64"
        },
        "header": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "location": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "requested-by": {
          "type": "integer",
          "format": "int64"
        },
        "time": {
          "type": "string",
          "format": "date-time"
        },
        "total_comments": {
          "type": "string"
        },
        "total_likes": {
          "type": "string"
        },
        "total_savoods": {
          "type": "string"
        }
      },
      "xml": {
        "name": "Offering"
      }
    },
    "Profile": {
      "type": "object",
      "properties": {
        "address": {
          "type": "object",
          "properties": {
            "city": {
              "type": "string"
            },
            "number": {
              "type": "string"
            },
            "street": {
              "type": "string"
            },
            "zip": {
              "type": "integer",
              "format": "int64"
            }
          }
        },
        "avatar-id": {
          "type": "string"
        },
        "background-id": {
          "type": "string"
        },
        "badges": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "description": {
          "type": "string"
        },
        "email": {
          "type": "string",
          "format": "email"
        },
        "firstname": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "lastname": {
          "type": "string"
        },
        "phone": {
          "type": "string",
          "format": "phone"
        }
      },
      "xml": {
        "name": "Profile"
      },
      "example": {
        "address": {
          "city": "Musterstadt",
          "number": 1337,
          "street": "Musterstraße",
          "zip": 42069
        },
        "avatar-id": "",
        "background-id": "",
        "badges": [
          true,
          false,
          true,
          true,
          true,
          false,
          false,
          true,
          false
        ],
        "description": "I save the wrap and the world",
        "email": "apiteam@swagger.io",
        "firstname": "Marty",
        "id": 5,
        "lastname": "McFlfy",
        "phone": "202-555-0191"
      }
    }
  },
  "responses": {
    "InvalidParameterInput": {
      "description": "Invalid parameter input was passed",
      "schema": {
        "$ref": "#/definitions/InvalidParameterInput"
      }
    },
    "NotFound": {
      "description": "The specified resource was not found",
      "schema": {
        "$ref": "#/definitions/ErrorModel"
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/ErrorModel"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    },
    "petstore_auth": {
      "type": "oauth2",
      "flow": "implicit",
      "authorizationUrl": "http://petstore.swagger.io/oauth/dialog",
      "scopes": {
        "read:pets": "read your pets",
        "write:pets": "modify pets in your account"
      }
    }
  },
  "tags": [
    {
      "description": "Management of profiles and user representation",
      "name": "profile"
    },
    {
      "description": "Management of offered food items",
      "name": "offering"
    },
    {
      "description": "Collection of nearby offerings",
      "name": "feed"
    },
    {
      "description": "Operations about user",
      "name": "message"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "denn nur lebendiges food tut gut",
    "title": "Savood",
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0"
  },
  "basePath": "/v2/",
  "paths": {
    "/feed": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "feed"
        ],
        "summary": "Display a feed",
        "operationId": "getFeed",
        "parameters": [
          {
            "type": "string",
            "name": "location",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Feed"
            }
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
            }
          }
        }
      }
    },
    "/healthcheck": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "health"
        ],
        "summary": "Check Server status",
        "operationId": "healthcheckGet",
        "responses": {
          "200": {
            "description": "Server is healthy"
          },
          "501": {
            "description": "Server is unhealthy",
            "schema": {
              "$ref": "#/definitions/ErrorModel"
            }
          }
        }
      }
    },
    "/message": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "message"
        ],
        "summary": "Add a new message",
        "operationId": "createNewMessage",
        "parameters": [
          {
            "description": "Message that needs to be added",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Message"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Message"
            }
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
            }
          }
        }
      }
    },
    "/message/{id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "message"
        ],
        "summary": "Display a message",
        "operationId": "getMessageById",
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Message"
            }
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
            }
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "message"
        ],
        "summary": "Delete a message",
        "operationId": "deleteMessageById",
        "responses": {
          "204": {
            "description": "Removed; No response."
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
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
          "message"
        ],
        "summary": "Update a message",
        "operationId": "updateMessageById",
        "parameters": [
          {
            "description": "New parameters of the message",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Message"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Accepted; No response."
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/offering": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "offering"
        ],
        "summary": "Add a new offering",
        "operationId": "createNewOffering",
        "parameters": [
          {
            "description": "Offering that needs to be added",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Offering"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Offering"
            }
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
            }
          }
        }
      }
    },
    "/offering/{id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "offering"
        ],
        "summary": "Display an offering",
        "operationId": "getOfferingById",
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Offering"
            }
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
            }
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "offering"
        ],
        "summary": "Delete an offering",
        "operationId": "deleteOfferingById",
        "responses": {
          "204": {
            "description": "Removed; No response."
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
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
          "offering"
        ],
        "summary": "Update an offering",
        "operationId": "updateOfferingById",
        "parameters": [
          {
            "description": "New parameters of the offering",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Offering"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Accepted; No response."
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/profile": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "summary": "Add a new profile",
        "operationId": "createNewProfile",
        "parameters": [
          {
            "description": "Profile that needs to be added",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
            }
          }
        }
      }
    },
    "/profile/{id}": {
      "get": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "summary": "Display a profile",
        "operationId": "getProfileById",
        "responses": {
          "200": {
            "description": "Object found and returned",
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
            }
          }
        }
      },
      "delete": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "summary": "Delete a profile",
        "operationId": "deleteProfileById",
        "responses": {
          "204": {
            "description": "Removed; No response."
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
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
          "profile"
        ],
        "summary": "Update a profile",
        "operationId": "updateProfileById",
        "parameters": [
          {
            "description": "New parameters of the profile",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Profile"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Accepted; No response."
          },
          "404": {
            "description": "Invalid parameter input was passed",
            "schema": {
              "$ref": "#/definitions/InvalidParameterInput"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Comment": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string"
        },
        "creator-id": {
          "type": "integer",
          "format": "int64"
        },
        "id": {
          "type": "string"
        },
        "time": {
          "type": "string",
          "format": "date-time"
        }
      },
      "xml": {
        "name": "Comment"
      }
    },
    "ErrorModel": {
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Feed": {
      "type": "array",
      "items": {
        "$ref": "#/definitions/Offering"
      },
      "xml": {
        "name": "Feed"
      }
    },
    "InvalidParameterInput": {
      "type": "object",
      "required": [
        "attribute",
        "message"
      ],
      "properties": {
        "attribute": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Message": {
      "type": "object",
      "properties": {
        "content": {
          "type": "string"
        },
        "from-id": {
          "type": "integer",
          "format": "int64"
        },
        "id": {
          "type": "string"
        },
        "important": {
          "type": "boolean"
        },
        "offering-id": {
          "type": "integer",
          "format": "int64"
        },
        "time": {
          "type": "string",
          "format": "date-time"
        },
        "to-id": {
          "type": "integer",
          "format": "int64"
        }
      },
      "xml": {
        "name": "Message"
      }
    },
    "Offering": {
      "type": "object",
      "properties": {
        "avatar-url": {
          "type": "string"
        },
        "best-by-date": {
          "type": "string",
          "format": "date"
        },
        "comments": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Comment"
          }
        },
        "creator-id": {
          "type": "integer",
          "format": "int64"
        },
        "header": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "location": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "requested-by": {
          "type": "integer",
          "format": "int64"
        },
        "time": {
          "type": "string",
          "format": "date-time"
        },
        "total_comments": {
          "type": "string"
        },
        "total_likes": {
          "type": "string"
        },
        "total_savoods": {
          "type": "string"
        }
      },
      "xml": {
        "name": "Offering"
      }
    },
    "Profile": {
      "type": "object",
      "properties": {
        "address": {
          "$ref": "#/definitions/profileAddress"
        },
        "avatar-id": {
          "type": "string"
        },
        "background-id": {
          "type": "string"
        },
        "badges": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "description": {
          "type": "string"
        },
        "email": {
          "type": "string",
          "format": "email"
        },
        "firstname": {
          "type": "string"
        },
        "id": {
          "type": "string"
        },
        "lastname": {
          "type": "string"
        },
        "phone": {
          "type": "string",
          "format": "phone"
        }
      },
      "xml": {
        "name": "Profile"
      },
      "example": {
        "address": {
          "city": "Musterstadt",
          "number": 1337,
          "street": "Musterstraße",
          "zip": 42069
        },
        "avatar-id": "",
        "background-id": "",
        "badges": [
          true,
          false,
          true,
          true,
          true,
          false,
          false,
          true,
          false
        ],
        "description": "I save the wrap and the world",
        "email": "apiteam@swagger.io",
        "firstname": "Marty",
        "id": 5,
        "lastname": "McFlfy",
        "phone": "202-555-0191"
      }
    },
    "profileAddress": {
      "type": "object",
      "properties": {
        "city": {
          "type": "string"
        },
        "number": {
          "type": "string"
        },
        "street": {
          "type": "string"
        },
        "zip": {
          "type": "integer",
          "format": "int64"
        }
      },
      "x-go-gen-location": "models"
    }
  },
  "responses": {
    "InvalidParameterInput": {
      "description": "Invalid parameter input was passed",
      "schema": {
        "$ref": "#/definitions/InvalidParameterInput"
      }
    },
    "NotFound": {
      "description": "The specified resource was not found",
      "schema": {
        "$ref": "#/definitions/ErrorModel"
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/ErrorModel"
      }
    }
  },
  "securityDefinitions": {
    "api_key": {
      "type": "apiKey",
      "name": "api_key",
      "in": "header"
    },
    "petstore_auth": {
      "type": "oauth2",
      "flow": "implicit",
      "authorizationUrl": "http://petstore.swagger.io/oauth/dialog",
      "scopes": {
        "read:pets": "read your pets",
        "write:pets": "modify pets in your account"
      }
    }
  },
  "tags": [
    {
      "description": "Management of profiles and user representation",
      "name": "profile"
    },
    {
      "description": "Management of offered food items",
      "name": "offering"
    },
    {
      "description": "Collection of nearby offerings",
      "name": "feed"
    },
    {
      "description": "Operations about user",
      "name": "message"
    }
  ]
}`))
}
