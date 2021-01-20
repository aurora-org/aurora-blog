### Aurora System API Document

#### 请求地址

http://127.0.0.1:8080/api/v1

#### 响应码对应状态

```text
OK      = 200
Created = 201
Updated = 202
Deleted = 203

BadRequest = 400
Forbidden  = 403
NotFound   = 404

InternalServerError = 500
```


#### Account 账户相关

##### 账号登录

POST /mng/account/login

- Request

```json
{
    "Body": {
        "userName": "aurora",
        "password": "aurora"
    }
}
```

- Response

> 该响应返回 token，用于需要授权的操作，token 有效期默认为一小时。

```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "token": "${token}"
    }
}
```

##### 修改密码

PUT /mng/account/password

- Request

```json
{
    "Headers": {
        "Authorization": "${token}"
    },
    "Body": {
        "old": "aurora",
        "new": "123456"
    }
}
```

- Response

```json
{
    "code": 202,
    "status": "Updated",
    "data": null
}
```

#### Author 作者相关

##### 获取作者信息

GET /mng/author 

- Response

```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "age": 23,
        "avatar": "https://image.qmdx00.cn/black_avatar.jpeg",
        "createdAt": "2021-01-19 10:07:08",
        "extra": "",
        "gender": 1,
        "hobby": "sing dance and rap",
        "name": "wimi",
        "nickName": "big ass",
        "updatedAt": "2021-01-20 10:00:06"
    }
}
```

##### 修改作者信息

PUT /mng/author

- Request

> 将需要修改的属性写入 request body 中即可，无需传所有属性。

```json
{
    "Headers": {
        "Authorization": "${token}"
    },
    "Body": {
        "gender": 0
    }
}
```

- Response

```json
{
    "code": 202,
    "status": "Updated",
    "data": {
        "age": 23,
        "avatar": "https://image.qmdx00.cn/black_avatar.jpeg",
        "createdAt": "2021-01-15 10:07:08",
        "extra": "",
        "gender": 0,
        "hobby": "sing dance and rap",
        "name": "wimi",
        "nickName": "big ass",
        "updatedAt": "2021-01-20 10:00:06"
    }
}
```

#### Site 站点相关

##### 获取站点信息

GET /mng/site

- Response

```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "beiAn": "",
        "createdAt": "2021-01-18 17:20:40",
        "extra": "",
        "name": "Skyscraper",
        "powered": "Wimi",
        "startAt": "2021-01-18T17:20:40+08:00",
        "updatedAt": "2021-01-18 17:20:40"
    }
}
```

##### 修改站点信息

PUT /mng/site

- Request

```json
{
    "Headers": {
        "Authorization": "${token}"
    },
    "Body": {
        "name": "Aurora"
    }
}
```

- Response

```json
{
    "code": 202,
    "status": "Updated",
    "data": {
        "beiAn": "",
        "createdAt": "2021-01-18 17:20:40",
        "extra": "",
        "name": "Aurora",
        "powered": "Wimi",
        "startAt": "2021-01-18T17:20:40+08:00",
        "updatedAt": "2021-01-20 10:54:53"
    }
}
```

#### Article 文章相关

##### 获取文章列表

GET /mng/articles?{page, size, tag, category, order, by}

- Request

> 该接口默认过滤设置为不可见的文章
>
> page 必填，为文章列表页码 ，1~N
>
> size 必填，为每页显示文章数 ，1~N （N <= 50)
>
> tag 可选，通过标签筛选文章
>
> category 可选，通过分类筛选文章
>
> order 可选，排序规则，默认为创建时间，可选 [ CREATE | UPDATE | TIMES ]
>
> by 可选，排序方式，默认降序排列，可选 [ ASC | DESC ]

```json
{
    "Params": {
        "page": 1,
        "size": 10,
        "tag": "default",
        "category": "default",
        "order": "TIMES",
        "by": "ASC"
    }
}
```

- Response

```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "pagination": {
            "total": 2,
            "page": 1,
            "size": 10
        },
        "objects": [
            {
                "author": "big ass",
                "banner": "https://www.baidu.com",
                "category": "default",
                "content": "test test test",
                "createdAt": "2021-01-20T16:26:04+08:00",
                "extra": "",
                "id": 4,
                "tag": "default",
                "times": 0,
                "title": "第四篇博客",
                "updatedAt": "2021-01-20T16:26:04+08:00",
                "visible": true
            },
            {
                "author": "big ass",
                "banner": "https://www.baidu.com",
                "category": "default",
                "content": "hello world",
                "createdAt": "2021-01-20T16:25:44+08:00",
                "extra": "",
                "id": 3,
                "tag": "default",
                "times": 1,
                "title": "第三篇博客",
                "updatedAt": "2021-01-20T16:25:44+08:00",
                "visible": true
            }
        ]
    }
}
```

##### 获取指定文章

GET /mng/articles/{id}

- Response

```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "author": "big ass",
        "banner": "https://www.baidu.com",
        "category": "default",
        "content": "hello world",
        "createdAt": "2021-01-20T16:25:44+08:00",
        "extra": "",
        "id": 3,
        "tag": "default",
        "times": 1,
        "title": "第三篇博客",
        "updatedAt": "2021-01-20T16:25:44+08:00",
        "visible": true
    }
}
```

##### 新增文章

POST /mng/articles

- Request

```json
{
    "Headers": {
        "Authorization": "${token}"
    },
    "Body": {
        "title": "测试文章",
        "content": "test test test",
        "banner": "https://www.baidu.com",
        "tag": "default",
        "category": "default",
        "visible": true
    }
}
```

- Response

```json
{
    "code": 201,
    "status": "Created",
    "data": {
        "author": "big ass",
        "banner": "https://www.baidu.com",
        "category": "default",
        "content": "test test test",
        "createdAt": "2021-01-20 17:02:55",
        "extra": "",
        "id": 5,
        "tag": "default",
        "times": 0,
        "title": "测试文章",
        "updatedAt": "2021-01-20 17:02:55",
        "visible": true
    }
}
```

##### 阅读指定文章

POST /mng/articles/{id}

> 调用该接口使文章阅读次数加一，无请求参数和响应体

##### 修改文章

PUT /mng/articles/{id}

- Request

> 可修改文章标题，主图，内容，标签，分类以及可见性，传需要修改的字段即可

```json
{
    "Headers": {
        "Authorization": "${token}"
    },
    "Body": {
 		"visible": false       
    }
}
```

- Response

```json
{
    "code": 202,
    "status": "Updated",
    "data": {
        "author": "big ass",
        "banner": "https://www.baidu.com",
        "category": "default",
        "content": "test test test",
        "createdAt": "2021-01-20 17:02:55",
        "extra": "",
        "id": 5,
        "tag": "default",
        "times": 0,
        "title": "测试文章",
        "updatedAt": "2021-01-20 17:13:31",
        "visible": false
    }
}
```

##### 删除文章

DELETE /mng/articles/{id}

- Request

```json
{
    "Headers": {
        "Authorization": "${token}"
    }
}
```

- Response

```json
{
    "code": 203,
    "status": "Deleted",
    "data": "${id}"
}
```

#### Category 分类相关

##### 获取分类列表

GET /mng/categories

- Response

```json
{
    "code": 200,
    "status": "OK",
    "data": [
        {
            "createdAt": "2021-01-20 11:05:02",
            "description": "编程相关",
            "extra": "",
            "id": 1,
            "name": "学习",
            "updatedAt": "2021-01-20 11:05:02"
        }
    ]
}
```

##### 创建分类

POST /mng/categories

- Request

```json
{
    "Headers": {
        "Authorization": "${token}"
    },
    "Body": {
        "name": "Go",
        "description": "编程语言"
    }
}
```

- Response

```json
{
    "code": 201,
    "status": "Created",
    "data": {
        "createdAt": "2021-01-20 11:06:50",
        "description": "编程语言",
        "extra": "",
        "id": 2,
        "name": "Go",
        "updatedAt": "2021-01-20 11:06:50"
    }
}
```

##### 删除分类

DELETE /mng/categories/{id}

- Request

```json
{
    "Headers": {
        "Authorization": "${token}"
    }
}
```

- Response

> data 为删除的分类的 ID

```json
{
    "code": 203,
    "status": "Deleted",
    "data": "${id}"
}
```

#### Tag 标签相关

##### 获取标签列表

GET /mng/tags

- Response

```json
{
    "code": 200,
    "status": "OK",
    "data": [
        {
            "createdAt": "2021-01-20 11:14:33",
            "extra": "",
            "id": 1,
            "name": "Go",
            "updatedAt": "2021-01-20 11:14:33"
        }
    ]
}
```

##### 创建标签

POST /mng/tags

- Request

```json
{
	"Headers": {
		"Authorization": "${token}"
    },
    "Body": {
        "name": "Go"
    }
}
```

- Response

```json
{
    "code": 201,
    "status": "Created",
    "data": {
        "createdAt": "2021-01-20 11:16:14",
        "extra": "",
        "id": 2,
        "name": "Algorithms",
        "updatedAt": "2021-01-20 11:16:14"
    }
}
```

##### 删除标签

DELETE /mng/tags/{id}

- Request

```json
{
    "Headers": {
        "Authorization": "${token}"
    }
}
```

- Response

```json
{
    "code": 203,
    "status": "Deleted",
    "data": "${id}"
}
```

