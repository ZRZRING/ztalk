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
        "/posts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "可按社区按时间或分数排序查询帖子列表接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "升级版帖子列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "社区 ID",
                        "name": "community_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "time",
                        "description": "排序依据",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 10,
                        "description": "每页数据量",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.PostListType"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Community": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.PostDetail": {
            "type": "object",
            "required": [
                "community_id",
                "content",
                "title"
            ],
            "properties": {
                "author_id": {
                    "description": "作者id",
                    "type": "integer"
                },
                "author_name": {
                    "type": "string"
                },
                "community": {
                    "$ref": "#/definitions/models.Community"
                },
                "community_id": {
                    "description": "社区id",
                    "type": "integer"
                },
                "content": {
                    "description": "帖子内容",
                    "type": "string"
                },
                "create_time": {
                    "description": "帖子创建时间",
                    "type": "string"
                },
                "id": {
                    "description": "帖子id",
                    "type": "string",
                    "example": "0"
                },
                "status": {
                    "description": "帖子状态",
                    "type": "integer"
                },
                "title": {
                    "description": "帖子标题",
                    "type": "string"
                },
                "vote_num": {
                    "type": "integer"
                }
            }
        },
        "response.PostListType": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "allOf": [
                        {
                            "$ref": "#/definitions/response.ResCode"
                        }
                    ]
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.PostDetail"
                    }
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "response.ResCode": {
            "type": "integer",
            "enum": [
                1000,
                1001,
                1002,
                1003,
                1004,
                1005,
                1006,
                1007
            ],
            "x-enum-varnames": [
                "CodeSuccess",
                "CodeInvalidParam",
                "CodeUserExist",
                "CodeUserNotExist",
                "CodeInvalidPassword",
                "CodeServerBusy",
                "CodeNeedLogin",
                "CodeInvalidToken"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "http://127.0.0.1",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "ztalk",
	Description:      "在线论坛平台",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
