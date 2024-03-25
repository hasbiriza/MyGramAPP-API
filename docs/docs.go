// Code generated by swaggo/swag. DO NOT EDIT
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
        "/comments": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get comments",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Get comments",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetCommentResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Add new comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Add new comment",
                "parameters": [
                    {
                        "description": "body request for add new comment",
                        "name": "dto.NewCommentRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewCommentRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.GetCommentResponse"
                        }
                    }
                }
            }
        },
        "/comments/{commentId}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Update comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "commentId",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body request for update comment",
                        "name": "dto.UpdateCommentRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateCommentRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetCommentResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete comment",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comments"
                ],
                "summary": "Delete comment",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "commentId",
                        "name": "commentId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetCommentResponse"
                        }
                    }
                }
            }
        },
        "/photos": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get photos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photos"
                ],
                "summary": "Get photos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetPhotoResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Add new photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photos"
                ],
                "summary": "Add new photo",
                "parameters": [
                    {
                        "description": "body request for add new photo",
                        "name": "dto.NewPhotoRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewPhotoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.GetPhotoResponse"
                        }
                    }
                }
            }
        },
        "/photos/{photoId}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photos"
                ],
                "summary": "Update photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "photoId",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body request for update photo",
                        "name": "dto.PhotoUpdateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.PhotoUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetPhotoResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete photo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Photos"
                ],
                "summary": "Delete photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "photoId",
                        "name": "photoId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetPhotoResponse"
                        }
                    }
                }
            }
        },
        "/socialmedias": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get social medias",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Get social medias",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetSocialMediaHttpResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Add new social media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Add new social media",
                "parameters": [
                    {
                        "description": "body request for add new social media",
                        "name": "dto.NewSocialMediaRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewSocialMediaRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.GetSocialMediaResponse"
                        }
                    }
                }
            }
        },
        "/socialmedias/{socialMediaId}": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update social media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Update social media",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "socialMediaId",
                        "name": "socialMediaId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body request for update social media",
                        "name": "dto.UpdateSocialMediaRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateSocialMediaRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetSocialMediaResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete social media",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Social Media"
                ],
                "summary": "Delete social media",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "socialMediaId",
                        "name": "socialMediaId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetSocialMediaResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "User update",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User update",
                "parameters": [
                    {
                        "description": "body request for user update",
                        "name": "dto.UserUpdateRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetUserResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create new Users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create new User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetUserResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "User login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "body request for user login",
                        "name": "dto.UserLoginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetUserResponse"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "User register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User register",
                "parameters": [
                    {
                        "description": "body request for user register",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewUserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.GetUserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.GetCommentResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "dto.GetPhotoResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "dto.GetSocialMedia": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2023-10-09T05:14:35.19324086+07:00"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Monday Weeekly Official"
                },
                "social_media_url": {
                    "type": "string",
                    "example": "https://www.instagram.com/_weeekly/"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2023-10-09T05:14:35.19324086+07:00"
                },
                "user": {
                    "$ref": "#/definitions/dto.SocialMediaUser"
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "dto.GetSocialMediaHttpResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "social media successfully fetched"
                },
                "social_media": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.GetSocialMedia"
                    }
                },
                "status_code": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "dto.GetSocialMediaResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "dto.GetUserResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "dto.NewCommentRequest": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "so beautiful"
                },
                "photo_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "dto.NewPhotoRequest": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string",
                    "example": "Hello I'm Monday from Weeekly, hopefully You can do this!"
                },
                "photo_url": {
                    "type": "string",
                    "example": "https://www.pinterest.com/pin/807973989398829161/"
                },
                "title": {
                    "type": "string",
                    "example": "monday awesome"
                }
            }
        },
        "dto.NewSocialMediaRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Monday Weeekly Official"
                },
                "social_media_url": {
                    "type": "string",
                    "example": "https://www.instagram.com/_weeekly/"
                }
            }
        },
        "dto.NewUserRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 21
                },
                "email": {
                    "type": "string",
                    "example": "monday.day@email.com"
                },
                "password": {
                    "type": "string",
                    "example": "secret"
                },
                "username": {
                    "type": "string",
                    "example": "monday"
                }
            }
        },
        "dto.PhotoUpdateRequest": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string",
                    "example": "Hello I'm Monday from Weeekly, stay strong!"
                },
                "photo_url": {
                    "type": "string",
                    "example": "https://www.pinterest.com/pin/807973989398829161/"
                },
                "title": {
                    "type": "string",
                    "example": "monday awesome"
                }
            }
        },
        "dto.SocialMediaUser": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "profile_image_url": {
                    "type": "string",
                    "example": "https://www.pinterest.com/pin/807973989398829161/"
                },
                "username": {
                    "type": "string",
                    "example": "monday"
                }
            }
        },
        "dto.UpdateCommentRequest": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "omg so beautiful"
                }
            }
        },
        "dto.UpdateSocialMediaRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Weeekly Monday Official"
                },
                "social_media_url": {
                    "type": "string",
                    "example": "https://www.instagram.com/_weeekly/"
                }
            }
        },
        "dto.UserLoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "monday.day@email.com"
                },
                "password": {
                    "type": "string",
                    "example": "secret"
                }
            }
        },
        "dto.UserUpdateRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "monday.day@weeekly.com"
                },
                "username": {
                    "type": "string",
                    "example": "monday"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer\" followed by a space and token you got from the User Login api.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
