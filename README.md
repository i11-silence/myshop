# Myshop



## 接口



### 用户相关

#### 注册

请求路径：POST /user/register



 请求参数：application/json

| 名称     | 位置 | 类型   | 必选 | 说明   |
| -------- | ---- | ------ | ---- | ------ |
| username | body | string | 是   | 用户名 |
| password | body | string | 是   | 密码   |

 

 返回示例：

成功返回：

```JSON
{
  "status": 10000,
  "info": "success"
}
```

#### 登录（获取token）

 请求路径： GET /user/token

 

请求参数：application/json

| 名称     | 类型   | 必选 | 说明          |
| -------- | ------ | ---- | ------------- |
| username | string | 是   | 用户名 |
| password | string | 是   | 密码          |

返回参数：

| 字段名        | 必选 | 类型          | 说明          |
| ------------- | ---- | ------------- | ------------- |
| refresh_token | 是   | Bearer $token | refresh_token |
| token         | 是   | Bearer $token | token         |

返回示例：

成功返回

```JSON
 {
   "status": 10000,
   "info": "success",
   "data": {
     "refresh_token": "{refresh_token}",
     "token": "{token}"
   }
 }
```

#### 刷新token（维持登录状态）

请求路径： GET /user/token/refresh

请求头：

| 字段名        | 必选 | 数值          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

请求参数:

| 名称          | 位置  | 类型   | 必选 | 说明          |
| ------------- | ----- | ------ | ---- | ------------- |
| refresh_token | query | string | 是   | refresh_token |

返回参数：

| 字段名        | 必选 | 类型          | 说明            |
| ------------- | ---- | ------------- | --------------- |
| refresh_token | 是   | Bearer $token | 新refresh_token |
| token         | 是   | Bearer $token | 新token         |

返回示例：

成功返回

```JSON
{
  "status": 10000,
  "info": "success",
  "data": {
    "refresh_token": "{refresh_token}",
    "token": "{token}"
  }
}
```

#### 修改用户密码

请求路径：PUT /user/password

请求头：

| 字段名        | 必选 | 数值          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

 请求参数：application/json

| 名称         | 类型   | 必选 | 说明   |
| ------------ | ------ | ---- | ------ |
| old_password | string | 是   | 旧密码 |
| new_password | string | 是   | 新密码 |



返回示例：

成功返回

```JSON
{
  "info": "success",
  "status": 10000
}
```


#### 修改用户信息

请求路径：PUT /user/info

该接口用于更改用户信息，除token外，全部请求参数都是非必选的，要求该接口有某个参数才修改对应的信息。比如我请求时只选了nickname和introduction字段，那么我们只改用户昵称和简介的信息，其他信息保持原状。

请求头:

| 字段名        | 必选 | 数值          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

请求参数：application/json

| 名称    | 位置 | 类型   | 必选 | 说明                   |
| ----- | ---- | ------ | ---- | ---------------------- |
| nickname | body | string | 否   | 昵称                   |
|       
                 |
| telephone | body | string | 否   | 手机号                 |
                  |
| gender | body | string | 否   | 性别                   |
                   |
| birthday | body | string | 否   | 生日，格式为YYYY-MM-DD |



返回示例：

成功返回

```JSON
{
   "info": "success",
   "status": 10000
 }
```

#### 获取用户信息

请求路径：GET /user/info/{user_id}

请求头:

| 字段名        | 必选 | 数值          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

请求参数：

| 名称    | 位置 | 类型   | 必选 | 说明   |
| ------- | ---- | ------ | ---- | ------ |
| user_id | path | string | 是   | 用户id |

返回参数：

| 字段名 | 必选 | 类型         | 说明     |
| ------ | ---- | ------------ | -------- |
| user   | 是   | 复杂数据类型 | 用户信息 |

#### **返回示例：**

**成功返回**

```JSON
{
  "status": 10000,
  "info": "success",
  "data": {
      "user": {
      "nickname": "你好世界",
      "phone": 12345678900,
      "gender": "未知",
      "birthday": "2022-01-02"
    }
  }
}
```

### 商品相关

#### 获取商品列表

请求路径：GET /product/list



 返回参数：

| 字段名   | 类型             | 说明       |
|-------| ---------------- | ---------- |
| goods | 复杂数据类型数组 | 商品的信息 |

#### 返回示例：

**成功返回**

```JSON
{
  "status": 10000,
  "info": "success",
  "data": {
    "goods":[
      {
        "id": "1",
        "name": "傲慢与偏见",
        "description": "一本书",
        "type":"book",
        "comment_num": 35,
        "price": 9.8,
        "is_addedCart":true
        "cover": "http://127.0.0.1/picture_url1",
        "publish_time": "1980-11-07",
        "link": "http://127.0.0.1/test1"
      },
      {
        "id": "2",
        "name": "T-shirt",
        "description": "一件短袖",
        "type":"clothes",
        "comment_num": 100,
        "price": 88.88,
        "is_addedCart":false
        "cover": "http://127.0.0.1/picture_url2",
        "publish_time": "1980-11-07",
        "link": "http://127.0.0.1/test2"
      }
    ]
  }
}
```

#### 搜索商品

 请求路径 ：GET /book/search

 请求头：

| 字段名        | 必选 | 数值          | 说明                                        |
| ------------- | ---- | ------------- | ------------------------------------------- |
| Authorization | 否   | Bearer $token | 验证token，没token则is_addedCart字段为false |

#### 请求参数：query

| 名称         | 位置  | 类型   | 必选 | 说明                              |
| ------------ | ----- | ------ | ---- | --------------------------------- |
| product_name | query | string | 是   | 商品名称/代号（如果有相关的字段） |

#### **返回参数：**

| 字段名  | 必选 | 类型         | 说明     |
|------| ---- | ------------ | -------- |
| good | 是   | 复杂数据类型 | 商品信息 |

#### 返回示例：

成功返回

```JSON
{
  "status": 10000,
  "info": "success",
  "data": {
    "good": {
      "product_id": "1",
      "name": "傲慢与偏见",
      "description": "一本书",
      "type": "book",
      "comment_num": 35,
      "price": 9.8,
      "is_addedCart": true,
      "cover": "http://127.0.0.1/picture_url1",
      "publish_time": "1980-11-07",
      "link": "http://127.0.0.1/test1"
    }
  }
}
```

#### 加入购物车

#### 请求路径：PUT /product/addCart

#### 请求头：

| 字段名        | 必选 | 数值          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

#### 请求参数：x-www-form-urlencoded/form-data

| 名称       | 类型   | 必选 | 说明   |
| ---------- | ------ | ---- | ------ |
| product_id | string | 是   | 商品id |

#### 返回参数：

#### 返回示例：

**成功返回**

```JSON
{
  "info": "success",
  "status": 10000
}
```

#### 获取购物车商品列表

#### 请求路径 ： GET /product/cart

#### 请求头：

| 字段名        | 必选 | 数值          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |



#### 返回参数

| 字段名   | 类型         | 说明    |
|-------| ------------ | ------ |
| goods | 复杂数据类型数组 | 商品的信息 |

#### 返回示例：

```JSON
{
  "status": 10000,
  "info": "success",
  "data": {
    "goods":[
      {
        "id": "1",
        "name": "傲慢与偏见",
        "type":"book",
        "price": 9.8,
        "cover": "http://127.0.0.1/picture_url1",
        "link": "http://127.0.0.1/test1",
        "comment_num":114,
        "is_addedCart":true
      },
      {
        "id": "2",
        "name": "T-shirt",
        "type":"clothes",
        "price": 88.88,
        "cover": "http://127.0.0.1/picture_url2",
        "link": "http://127.0.0.1/test2",
        "comment_num":514,
        "is_addedCart":true
      }
    ]
  }
}
```

#### 获取商品详情

 请求路径：GET /product/info/{product_id}


 请求参数：

| 名称       | 位置 | 类型   | 必选 | 说明   |
| ---------- | ---- | ------ | ---- | ------ |
| product_id | path | string | 是   | 商品id |

返回参数：

| 字段名  | 类型     | 说明   |
|------|--------|------|
| good | 复杂数据类型 | 商品信息 |



#### 返回示例：

```JSON
{
  "status": 10000,
  "info": "success",
  "data": {
    "good": {
      "product_id": "1",
      "name": "傲慢与偏见",
      "description": "一本书",
      "type": "book",
      "comment_num": 35,
      "price": 9.8,
      "is_addedCart": true,
      "cover": "http://127.0.0.1/picture_url1",
      "publish_time": "1980-11-07",
      "link": "http://127.0.0.1/test1"
    }
  }
  }
```

#### 获取相应标签的商品列表

请求路径 GET /product/{type}



 请求参数：

| 名称   | 位置   | 类型   | 必选 | 说明     |
| ---- |------| ------ | ---- | -------- |
| type | path | string | 否   | 商品标签 |

 返回参数：

| 字段名   | 必选 | 类型     | 说明     |
|-------| ---- |--------| -------- |
| goods | 是   | 复杂数据数组 | 商品信息 |

 返回示例：

成功返回

```JSON
{
  "data": [
    {
      "id": 2,
      "name": "3",
      "description": "4",
      "type": "5",
      "price": 6,
      "cover": "7",
      "link": "8",
      "Publish_time": "2025-02-15T20:29:41Z",
      "publish_time": "Sat Feb 15 20:29:41 2025",
      "comment_num": 0,
      "is_added_cart": true
    },
    {
      "id": 3,
      "name": "a",
      "description": "b",
      "type": "5",
      "price": -1,
      "cover": "c",
      "link": "d",
      "Publish_time": "2025-02-15T21:33:30Z",
      "publish_time": "Sat Feb 15 21:33:30 2025",
      "comment_num": 0,
      "is_added_cart": false
    }
  ],
  "info": "success",
  "status": 10000
}
```

### 评论相关

 #### 获取商品的评论

 请求路径： GET /comment/{product_id}


 请求参数：

| 名称       | 位置 | 类型   | 必选 | 说明   |
| ---------- | ---- | ------ | ---- | ------ |
| product_id | path | string | 是   | 商品id |

返回参数：

| 字段名   | 必选 | 类型             | 说明                       |
| -------- | ---- | ---------------- | -------------------------- |
| comments | 是   | 复杂数据类型数组 | 信息不完全的书评信息的数组 |

 返回示例:

成功返回

```JSON
{
  "info": "success",
  "status": 10000,
  "comments": [
    {
      "id": "79",
      "publish_time": 503270377478,
      "content": "occaecat est voluptate qui officia",
      "username": "abc",
      "nickname": "康静",
      "praise_count": 65,
      "is_praised": 1,（0为未处理， 1为点赞， 2为点踩）
      "good_id": "99"
    },
    {
      "id": "80",
      "publish_time": 503270377478,
      "content": "occaecat est voluptate qui officia",
      "username": "dfg",
      "nickname": "康静",
      "praise_count": 65,
      "is_praised": 1,
      "good_id": "99"
    }
  ]
}
```

#### 给商品评论

 请求路径： POST /comment/{product_id}

 请求头：

| 字段名        | 必选 | 数值          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

 请求参数：application/json

| 名称         | 类型   | 必选 | 说明     |
|------------| ------ | ---- | -------- |
| product_id | string | 是   | 商品id   |
| content    | string | 是   | 评论内容 |

 返回参数：

| 字段名  | 必选 | 类型 | 说明       |
|------| ---- | ---- | ---------- |
| data | 是   | int  | 此评论的id |

返回示例：

```JSON
{
  "info": "success",
  "status": 10000,
  "data": "{comment_id}"
}
```

#### 删除书评



请求路径：DELETE /comment/{comment_id}

 请求头：

| 字段名        | 必选 | 数值          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

 请求参数：

| 名称       | 位置 | 类型   | 必选 | 说明   |
| ---------- | ---- | ------ | ---- | ------ |
| comment_id | path | string | 是   | 评论id |



返回示例：

成功返回

```JSON
{
  "info": "success",
  "status": 10000
}
```

#### 更新评论

请求路径： PUT /comment/{comment_id}

 请求头：

| 字段名        | 必选 | 数值          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

返回参数：

 请求参数：application/json

| 名称       | 类型   | 必选 | 说明     |
| ---------- | ------ | ---- | -------- |
| comment_id | string | 是   | 评论id   |
| content    | string | 是   | 评论内容 |

 返回示例：

成功返回

```JSON
{
  "info": "success",
  "status": 10000
}
```

### 操作相关

#### 点赞点踩

 请求路径： PUT /comment/praise

请求头： 

| 字段名        | 必选 | 数值          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

 请求参数：form-data/x-www-form-urlencoded

| 字段名     | 必选 | 类型   | 说明                                               |
| ---------- | ---- | ------ | -------------------------------------------------- |
| model      | 是   | int    | 模型标识。当model为1时，为点赞；model为2时，为点踩 |
| comment_id | 是   | string | 要点赞的id。                                       |



 返回示例：

成功返回

```JSON
{
    "info": "success",
    "status": 10000
}
```

 #### 下单

 请求路径： POST /operate/order

 请求头：

| 字段名        | 必选 | 数值          | 说明      |
| ------------- | ---- | ------------- | --------- |
| Authorization | 是   | Bearer $token | 验证token |

 请求参数：application/json

| 字段名  | 必选 | 类型     | 说明     |
| ------- | ---- |--------| -------- |
| user_id | 是   | string | 用户id   |
| orders  | 是   | 复杂类型数组 | 订单内容 |
| address | 是   | string | 地址信息 |
| total   | 是   | float  | 总价     |

返回参数：

| 名称     | 类型   | 必选 | 说明   |
| -------- | ------ | ---- | ------ |
| order_id | string | 是   | 订单id |

 返回示例：

成功返回

```JSON
{
  "info": "success",
  "status": 10000,
  "order_id":233
}
```

## 状态码
- 正常：10000
- 处理出错：30000
