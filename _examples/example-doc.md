
# Blog

This is a simple blogging service provide **REST API** for consumer to build their own blogging platform on demand.

## Indices

* [Articles](#articles)

  * [Create article](#1-create-article)
  * [List articles](#2-list-articles)

* [Users](#users)

  * [Create user](#1-create-user)
  * [Delete user](#2-delete-user)
  * [Fetch user](#3-fetch-user)
  * [Fetch users](#4-fetch-users)
  * [Update user](#5-update-user)

* [Users/V2](#usersv2)

  * [Create user](#1-create-user-1)
  * [Update user](#2-update-user)

* [Ungrouped](#ungrouped)

  * [Health check](#1-health-check)


--------


## Articles
`Articles` directory contains all the article related APIs. Use `JWT` toekn for authentications for articles API.



### 1. Create article


Create article endpoint is accept form-data to create an article with binary data like file


***Endpoint:***

```bash
Method: POST
Type: FORMDATA
URL: {{blog-server}}/v1/articles
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| ExampleHeader1 | ExampleHeader1Value |  |
| ExampleHeader2 | ExampleHeader2Value |  |



***Body:***

| Key | Value | Description |
| --- | ------|-------------|
| author_id | 1 | Accept `author_id` as the primary *id* of author |
| title | This is title one | The `title` field must be between *1-255* chars |
| body | This is body one | The `body` field must be between *1-2000* chars |



***More example Requests/Responses:***


##### I. Example Request: Validation error


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| ExampleHeader1 | ExampleHeader1Value |  |
| ExampleHeader2 | ExampleHeader2Value |  |
| Content-Type | application/x-www-form-urlencoded |  |



***Body:***

| Key | Value | Description |
| --- | ------|-------------|
| author_id | 1 | Accept `author_id` as the primary *id* of author |



##### I. Example Response: Validation error
```js
{
    "errors": {
        "title": [
            "Title can not be empty"
        ],
        "body": [
            "Body can not be empty"
        ]
    },
    "message": "Validation error"
}
```


***Status Code:*** 422

<br>



##### II. Example Request: Invalid username/password



##### II. Example Response: Invalid username/password
```js
{
    "message": "Unauthorized attempt"
}
```


***Status Code:*** 401

<br>



##### III. Example Request: Success


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| ExampleHeader1 | ExampleHeader1Value |  |
| ExampleHeader2 | ExampleHeader2Value |  |



***Body:***

| Key | Value | Description |
| --- | ------|-------------|
| author_id | 1 | Accept `author_id` as the primary *id* of author |
| title | This is title one | The `title` field must be between *1-255* chars |
| body | This is body one | The `body` field must be between *1-2000* chars |



##### III. Example Response: Success
```js
{
    "data": {
        "id": 1,
        "author_id": 1,
        "title": "This is title one",
        "body": "This is body one"
    },
    "message": "Article created successfully"
}
```


***Status Code:*** 201

<br>



### 2. List articles


List articles endpoint provide article listing with *filtering*, *patination*


***Endpoint:***

```bash
Method: GET
Type: 
URL: {{blog-server}}/v1/articles
```



***Query params:***

| Key | Value | Description |
| --- | ------|-------------|
| page | {{id}} | Page is a `unsigned integer` which represents the page numer |
| limit | {{limit}} | Limit represents maximum numer of results in the response |



***More example Requests/Responses:***


##### I. Example Request: Invalid token


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Authorization | Bearer inalid_token | Providing invalid token cause `401` |



***Query:***

| Key | Value | Description |
| --- | ------|-------------|
| page | {{id}} | Page is a `unsigned integer` which represents the page numer |
| limit | {{limit}} | Limit represents maximum numer of results in the response |
| user_id | {{user_id}} | Filter the articles using *author_id* |



##### I. Example Response: Invalid token
```js
{
    "message": "Unauthorized token"
}
```


***Status Code:*** 401

<br>



##### II. Example Request: Success


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Authorization | Bearer {{jwt_token}} | Providing valid token won't cause *401* |



***Query:***

| Key | Value | Description |
| --- | ------|-------------|
| page | {{id}} | Page is a `unsigned integer` which represents the page numer |
| limit | {{limit}} | Limit represents maximum numer of results in the response |
| user_id | {{user_id}} | Filter the articles using *author_id* |



##### II. Example Response: Success
```js
{
    "data": [
        {
            "id": 3,
            "user_id": 10,
            "title": "Article 13",
            "body": "This is article three"
        },
        {
            "id": 2,
            "user_id": 10,
            "title": "Article 2",
            "body": "This is article two"
        },
        {
            "id": 1,
            "user_id": 10,
            "title": "Article 1",
            "body": "This is article one"
        }
    ]
}
```


***Status Code:*** 200

<br>



##### III. Example Request: Without filtering


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Authorization | Bearer {{jwt_token}} | Providing valid token won't cause *401* |



***Query:***

| Key | Value | Description |
| --- | ------|-------------|
| page | {{id}} | Page is a `unsigned integer` which represents the page numer |
| limit | {{limit}} | Limit represents maximum numer of results in the response |



##### III. Example Response: Without filtering
```js
{
    "data": [
        {
            "id": 3,
            "user_id": 10,
            "title": "Article 13",
            "body": "This is article three"
        },
        {
            "id": 2,
            "user_id": 10,
            "title": "Article 2",
            "body": "This is article two"
        },
        {
            "id": 1,
            "user_id": 10,
            "title": "Article 1",
            "body": "This is article one"
        }
    ]
}
```


***Status Code:*** 200

<br>



## Users
`Users` directory contains all the user related APIs. For authentication these apis requrie `BasicAuth`



### 1. Create user


Create user use `JSON` payload to create a user


***Endpoint:***

```bash
Method: POST
Type: RAW
URL: {{blog-server}}/v1/users
```



***Body:***

```js        
{
	"name": "Captain Jack Sparrow",
	"bio": "Captain Jack Sparrow is a fictional character and the main protagonist of the Pirates of the Caribbean film series. The character was created by screenwriters Ted Elliott and Terry Rossio and is portrayed by Johnny Depp"
}
```



***More example Requests/Responses:***


##### I. Example Request: Invalid username/password



##### I. Example Response: Invalid username/password
```js
{
    "message": "Unauthorized attempt"
}
```


***Status Code:*** 401

<br>



##### II. Example Request: Empty body



##### II. Example Response: Empty body
```js
{
    "message": "Invalid request body"
}
```


***Status Code:*** 400

<br>



##### III. Example Request: Validation error



***Body:***

```js        
{
	"name": "",
	"bio": ""
}
```



##### III. Example Response: Validation error
```js
{
    "errors": {
        "bio": [
            "Bio can not be empty"
        ],
        "name": [
            "Name can not be empty"
        ]
    },
    "message": "Validation error"
}
```


***Status Code:*** 422

<br>



##### IV. Example Request: Success



***Body:***

```js        
{
	"name": "Tom Hanks",
	"bio": "Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon."
}
```



##### IV. Example Response: Success
```js
{
    "data": {
        "id": 1,
        "name": "Tom Hanks",
        "bio": "Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon."
    },
    "message": "User created successfully"
}
```


***Status Code:*** 201

<br>



### 2. Delete user


Delete a single user using `id`


***Endpoint:***

```bash
Method: DELETE
Type: 
URL: {{blog-server}}/v1/users/{{id}}
```



***More example Requests/Responses:***


##### I. Example Request: Success



##### I. Example Response: Success
```js
{
    "message": "User deleted successfully"
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Invalid user id



##### II. Example Response: Invalid user id
```js
{
    "message": "Invalid user id"
}
```


***Status Code:*** 400

<br>



##### III. Example Request: Invalid username/password



##### III. Example Response: Invalid username/password
```js
{
    "message": "Unauthorized attempt"
}
```


***Status Code:*** 401

<br>



### 3. Fetch user


Fetch a single user using `id`


***Endpoint:***

```bash
Method: GET
Type: 
URL: {{blog-server}}/v1/users/{{id}}
```



***More example Requests/Responses:***


##### I. Example Request: Success



##### I. Example Response: Success
```js
{
    "data": {
        "id": 1,
        "name": "Tom Hanks",
        "bio": "Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon."
    }
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Invalid user id



##### II. Example Response: Invalid user id
```js
{
    "message": "Invalid user id"
}
```


***Status Code:*** 400

<br>



##### III. Example Request: Invalid username/password



##### III. Example Response: Invalid username/password
```js
{
    "message": "Unauthorized attempt"
}
```


***Status Code:*** 401

<br>



### 4. Fetch users


Fetch list of users using `pagination`


***Endpoint:***

```bash
Method: GET
Type: 
URL: {{blog-server}}/v1/users
```



***Query params:***

| Key | Value | Description |
| --- | ------|-------------|
| page | 1 |  |
| limit | 20 |  |



***More example Requests/Responses:***


##### I. Example Request: Success



***Query:***

| Key | Value | Description |
| --- | ------|-------------|
| page | 1 | Page numer is a `integer` which represents the page your are requesting |
| limit | 20 | Limit it the maximum number of users listing in the result. |



##### I. Example Response: Success
```js
{
    "data": [
        {
            "id": 1,
            "name": "Tom Hanks",
            "bio": "Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon."
        },
        {
            "id": 2,
            "name": "Captain Jack Sparrow",
            "bio": "Captain Jack Sparrow is a fictional character and the main protagonist of the Pirates of the Caribbean film series. The character was created by screenwriters Ted Elliott and Terry Rossio and is portrayed by Johnny Depp"
        }
    ]
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Invalid user id



##### II. Example Response: Invalid user id
```js
{
    "message": "Invalid user id"
}
```


***Status Code:*** 400

<br>



##### III. Example Request: Invalid username/password



##### III. Example Response: Invalid username/password
```js
{
    "message": "Unauthorized attempt"
}
```


***Status Code:*** 401

<br>



### 5. Update user


Update user use `JSON` payload to create a user


***Endpoint:***

```bash
Method: PUT
Type: RAW
URL: {{blog-server}}/v1/users/{{id}}
```



***Body:***

```js        
{
	"name": "Captain Jack Sparrow",
	"bio": "Captain Jack Sparrow is a fictional character and the main protagonist of the Pirates of the Caribbean film series. The character was created by screenwriters Ted Elliott and Terry Rossio and is portrayed by Johnny Depp"
}
```



***More example Requests/Responses:***


##### I. Example Request: Success



***Body:***

```js        
{
	"name": "Captain Jack Sparrow",
	"bio": "Captain Jack Sparrow is a fictional character and the main protagonist of the Pirates of the Caribbean film series. The character was created by screenwriters Ted Elliott and Terry Rossio and is portrayed by Johnny Depp"
}
```



##### I. Example Response: Success
```js
{
    "data": {
        "id": 1,
        "name": "Captain Jack Sparrow",
        "bio": "Captain Jack Sparrow is a fictional character and the main protagonist of the Pirates of the Caribbean film series. The character was created by screenwriters Ted Elliott and Terry Rossio and is portrayed by Johnny Depp"
    },
    "message": "User updated successfully"
}
```


***Status Code:*** 200

<br>



##### II. Example Request: Validation error



***Body:***

```js        
{
	"name": "",
	"bio": ""
}
```



##### II. Example Response: Validation error
```js
{
    "errors": {
        "bio": [
            "Bio can not be empty"
        ],
        "name": [
            "Name can not be empty"
        ]
    },
    "message": "Validation error"
}
```


***Status Code:*** 422

<br>



##### III. Example Request: Empty request body



##### III. Example Response: Empty request body
```js
{
    "message": "Invalid request body"
}
```


***Status Code:*** 400

<br>



##### IV. Example Request: Invalid username/password



##### IV. Example Response: Invalid username/password
```js
{
    "message": "Unauthorized attempt"
}
```


***Status Code:*** 401

<br>



##### V. Example Request: Invalid user id



***Body:***

```js        
{
	"name": "Captain Jack Sparrow",
	"bio": "Captain Jack Sparrow is a fictional character and the main protagonist of the Pirates of the Caribbean film series. The character was created by screenwriters Ted Elliott and Terry Rossio and is portrayed by Johnny Depp"
}
```



##### V. Example Response: Invalid user id
```js
{
    "message": "Invalid user id"
}
```


***Status Code:*** 400

<br>



## Users/V2



### 1. Create user


Create user use `JSON` payload to create a user


***Endpoint:***

```bash
Method: POST
Type: URLENCODED
URL: {{blog-server}}/v2/users
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| name | Tom Hanks |  |
| bio | Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon. |  |



***More example Requests/Responses:***


##### I. Example Request: Validation Error



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| name |  |  |
| bio |  |  |



##### I. Example Response: Validation Error
```js
{
    "errors": {
        "bio": [
            "Bio can not be empty"
        ],
        "name": [
            "Name can not be empty"
        ]
    },
    "message": "Validation error"
}
```


***Status Code:*** 422

<br>



##### II. Example Request: Success



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| name | Tom Hanks |  |
| bio | Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon. |  |



##### II. Example Response: Success
```js
{
    "data": {
        "id": 1,
        "name": "Tom Hanks",
        "bio": "Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon."
    },
    "message": "User created successfully"
}
```


***Status Code:*** 201

<br>



##### III. Example Request: Invalid username/password



##### III. Example Response: Invalid username/password
```js
{
    "message": "Unauthorized attempt"
}
```


***Status Code:*** 401

<br>



### 2. Update user


Create user use `JSON` payload to create a user


***Endpoint:***

```bash
Method: PATCH
Type: URLENCODED
URL: {{blog-server}}/v2/users/1
```



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| name | Mr. Tom Hanks |  |



***More example Requests/Responses:***


##### I. Example Request: Invalid username/password



##### I. Example Response: Invalid username/password
```js
{
    "message": "Unauthorized attempt"
}
```


***Status Code:*** 401

<br>



##### II. Example Request: Success



***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| name | Mr. Tom Hanks |  |



##### II. Example Response: Success
```js
{
    "data": {
        "id": 1,
        "name": "Mr. Tom Hanks",
        "bio": "Thomas Jeffrey Hanks is an American actor and filmmaker. Known for both his comedic and dramatic roles, Hanks is one of the most popular and recognizable film stars worldwide, and is widely regarded as an American cultural icon."
    },
    "message": "User partial update successful"
}
```


***Status Code:*** 206

<br>



## Ungrouped



### 1. Health check


System health check endpoint provide system health status for `probe`


***Endpoint:***

```bash
Method: GET
Type: 
URL: {{blog-server}}/
```



***More example Requests/Responses:***


##### I. Example Request: Error



##### I. Example Response: Error
```js
{
    "system": "Failed"
}
```


***Status Code:*** 500

<br>



##### II. Example Request: Success



##### II. Example Response: Success
```js
{
    "system": "OK"
}
```


***Status Code:*** 200

<br>



---
[Back to top](#blog)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2020-05-14 14:48:48 by [docgen](https://github.com/thedevsaddam/docgen)
