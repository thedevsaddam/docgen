
# SMS

Full `Rest API` collection and documentation of *School management system*

## Indices

* [Github API](#github-api)

  * [Fetch profile](#1-fetch-profile)
  * [create profile](#2-create-profile)

* [Student](#student)

  * [Fetch students](#1-fetch-students)

* [Teacher](#teacher)

  * [Fetch teachers](#1-fetch-teachers)
  * [Create teacher](#2-create-teacher)
  * [Update teacher](#3-update-teacher)
  * [Update teacher partially](#4-update-teacher-partially)
  * [Remove teacher](#5-remove-teacher)

* [Teacher/v2](#teacherv2)

  * [Fetch teachers](#1-fetch-teachers-1)

* [Default](#default)

  * [Login](#1-login)


--------


## Github API
Contains github ***API*** collection


### 1. Fetch profile


Get ***github*** profile information


***Endpoint:***

```bash
Method: GET
Type: RAW
URL: https://api.github.com/users/thedevsaddam
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Authorization | {{access_token}} | Valid `access_token` |


***Responses:***


Status: Success Response | Code: 200


```js
{
    "login": "thedevsaddam",
    "id": 9676798,
    "node_id": "MDQ6VXNlcjk2NzY3OTg=",
    "avatar_url": "https://avatars0.githubusercontent.com/u/9676798?v=4",
    "gravatar_id": "",
    "url": "https://api.github.com/users/thedevsaddam",
    "html_url": "https://github.com/thedevsaddam",
    "followers_url": "https://api.github.com/users/thedevsaddam/followers",
    "following_url": "https://api.github.com/users/thedevsaddam/following{/other_user}",
    "gists_url": "https://api.github.com/users/thedevsaddam/gists{/gist_id}",
    "starred_url": "https://api.github.com/users/thedevsaddam/starred{/owner}{/repo}",
    "subscriptions_url": "https://api.github.com/users/thedevsaddam/subscriptions",
    "organizations_url": "https://api.github.com/users/thedevsaddam/orgs",
    "repos_url": "https://api.github.com/users/thedevsaddam/repos",
    "events_url": "https://api.github.com/users/thedevsaddam/events{/privacy}",
    "received_events_url": "https://api.github.com/users/thedevsaddam/received_events",
    "type": "User",
    "site_admin": false,
    "name": "Saddam H",
    "company": "Pathao Inc.",
    "blog": "https://thedevsaddam.github.io/",
    "location": "Dhaka, Bangladesh",
    "email": null,
    "hireable": true,
    "bio": "Software Engineer | Open Source Enthusiast | Love to write elegant code | Golang | Postgres | MongoDB | Elasticsearch / sleepy head, silent guy ;)",
    "public_repos": 78,
    "public_gists": 38,
    "followers": 197,
    "following": 184,
    "created_at": "2014-11-11T14:07:13Z",
    "updated_at": "2019-01-25T16:17:23Z"
}
```


### 2. create profile


To create a new profile for user you must provide `Authorization` header with *valid* `access_token`.


***Endpoint:***

```bash
Method: POST
Type: RAW
URL: https://api.github.com/users/thedevsaddam
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/json | `content-type` must be `application/json` or `application/text` |
| Authorization | Bearer {{access_token}} | Provide `access_token` |


***Body:***

```js        
{
	"name": "John Doe",
	"age": 30,
	"is_bangladeshi": true
}
```


***Responses:***


Status: succes | Code: 200


```js
{
    "login": "thedevsaddam",
    "avatar_url": "https://avatars0.githubusercontent.com/u/9676798?v=4",
    "url": "https://api.github.com/users/thedevsaddam",
    "html_url": "https://github.com/thedevsaddam",
    "followers_url": "https://api.github.com/users/thedevsaddam/followers",
    "following_url": "https://api.github.com/users/thedevsaddam/following{/other_user}",
    "created_at": "2014-11-11T14:07:13Z",
    "updated_at": "2017-12-18T05:25:18Z"
}
```


## Student
Contains Students `API` collection


### 1. Fetch students


Fetch list of all students


***Endpoint:***

```bash
Method: GET
Type: RAW
URL: {{base_url}}/students
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Authorization | Bearer {{access_token}} | Must provide valid `access_token` |


***Query params:***

| Key | Value | Description |
| --- | ------|-------------|
| page | 1 | page number must be number (`int`) |
| limit | 20 | list limit must be a number (`int`) |


## Teacher
Teacher contains teacher's api collection


### 1. Fetch teachers


Get list of teachers


***Endpoint:***

```bash
Method: GET
Type: RAW
URL: {{base_url}}/teachers
```


***Query params:***

| Key | Value | Description |
| --- | ------|-------------|
| q | john | q can be name, email, phone etc |
| page | 1 | page as integer number |


### 2. Create teacher


***Endpoint:***

```bash
Method: POST
Type: FORMDATA
URL: {{base_url}}/teachers
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Authorization | Bearer {{access_token}} | `access_token` must be a valid *oAuth2* access token. You can also use ~~custom token~~ but it will be removed from next version |


***Body:***

| Key | Value | Description |
| --- | ------|-------------|
| name | John Doe | `name` field must be between *3 to 20* chars |
| age | 33 | `age` field must be a valid numeric value |


***Responses:***


Status: Validation Error | Code: 422


```js
{
    "errors": [
        "Name field is required",
        "Age field is required"
    ],
    "code": 42201
}
```


Status: Success | Code: 200


```js
{
	"message": "Teacher created successfully",
	"teacher_id": 1002
}
```


### 3. Update teacher


update a teacher using api


***Endpoint:***

```bash
Method: PUT
Type: URLENCODED
URL: {{base_url}}/teachers
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/x-www-form-urlencoded |  |


***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| name | John Doe |  |
| age | 33 |  |


### 4. Update teacher partially


This api need header for update teacher


***Endpoint:***

```bash
Method: PATCH
Type: URLENCODED
URL: {{base_url}}/teachers
```


***Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/x-www-form-urlencoded |  |
| Authorization | Bearer {{access_token}} |  |


***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| name | John Doe |  |
| age | 33 |  |


### 5. Remove teacher


update a teacher using api


***Endpoint:***

```bash
Method: DELETE
Type: 
URL: {{base_url}}/teachers/{{id}}
```


## Teacher/v2


### 1. Fetch teachers


Get list of teachers


***Endpoint:***

```bash
Method: GET
Type: RAW
URL: {{base_url}}/v2/teachers
```


***Query params:***

| Key | Value | Description |
| --- | ------|-------------|
| q | john | q can be name, email, phone etc |
| page | 1 | page as integer number |


## Default


### 1. Login


Inorder to access the private ***API*** you must get an access token by providing `username/password`


***Endpoint:***

```bash
Method: POST
Type: FORMDATA
URL: {{base_url}}/login
```


***Body:***

| Key | Value | Description |
| --- | ------|-------------|
| username | user | username email/phone |
| password | pass | pass must be greater than `5` chars |


---
[Back to top](#sms)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2019-02-12 19:45:27 by [docgen](https://github.com/thedevsaddam/docgen)
