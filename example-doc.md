<!--- Collection name and description -->

# SMS

Full `Rest API` collection and documentation of *School management system*

<!--- Request items indices -->

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
<!--- Iterate main collection -->


### Github API
Contains github ***API*** collection

<!--- Iterate collection items -->


###### 1. Fetch profile


Get ***github*** profile information


```bash
Method: GET
Type: raw
URL: https://api.github.com/users/thedevsaddam
```

<!--- headers items -->

***Headers:***

<!--- Iterate headers items -->
| Key | Value | Description |
| --- | ------|-------------|
| Authorization | {{access_token}} | Valid `access_token` |

<!--- End Iterate headers items -->

<!--- End  headers items -->


<!--- Query param items -->

<!--- End query param items -->

<!--- Body mode -->

<!--- Raw body data -->



<!---End Raw body data -->

<!---FormData -->

<!---End FormData -->


<!---x-urlencoded data -->

<!---End x-urlencoded data -->

<!--- End Body mode -->


<!--- Items response -->

***Response***

        
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



<!--- End Items response -->


<!--- End Iterate collection items -->

###### 2. create profile


To create a new profile for user you must provide `Authorization` header with *valid* `access_token`.


```bash
Method: POST
Type: raw
URL: https://api.github.com/users/thedevsaddam
```

<!--- headers items -->

***Headers:***

<!--- Iterate headers items -->
| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/json | `content-type` must be `application/json` or `application/text` |
| Authorization | Bearer {{access_token}} | Provide `access_token` |

<!--- End Iterate headers items -->

<!--- End  headers items -->


<!--- Query param items -->

<!--- End query param items -->

<!--- Body mode -->

<!--- Raw body data -->


***Body:***

```js        
{
	"name": "John Doe",
	"age": 30,
	"is_bangladeshi": true
}
```


<!---End Raw body data -->

<!---FormData -->

<!---End FormData -->


<!---x-urlencoded data -->

<!---End x-urlencoded data -->

<!--- End Body mode -->


<!--- Items response -->

***Response***

        
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



<!--- End Items response -->


<!--- End Iterate collection items -->


<!--- End Iterate main collection -->

### Student
Contains Students `API` collection

<!--- Iterate collection items -->


###### 1. Fetch students


Fetch list of all students


```bash
Method: GET
Type: raw
URL: {{base_url}}/students
```

<!--- headers items -->

***Headers:***

<!--- Iterate headers items -->
| Key | Value | Description |
| --- | ------|-------------|
| Authorization | Bearer {{access_token}} | Must provide valid `access_token` |

<!--- End Iterate headers items -->

<!--- End  headers items -->


<!--- Query param items -->

***Query params:***

<!--- Query param items -->
| Key | Value | Description |
| --- | ------|-------------|
| page | 1 | page number must be number (`int`) |
| limit | 20 | list limit must be a number (`int`) |


<!--- End query param items -->

<!--- Body mode -->

<!--- Raw body data -->



<!---End Raw body data -->

<!---FormData -->

<!---End FormData -->


<!---x-urlencoded data -->

<!---End x-urlencoded data -->

<!--- End Body mode -->


<!--- Items response -->


<!--- End Iterate collection items -->


<!--- End Iterate main collection -->

### Teacher
Teacher contains teacher's api collection

<!--- Iterate collection items -->


###### 1. Fetch teachers


Get list of teachers


```bash
Method: GET
Type: raw
URL: {{base_url}}/teachers
```

<!--- headers items -->


<!--- Query param items -->

***Query params:***

<!--- Query param items -->
| Key | Value | Description |
| --- | ------|-------------|
| q | john | q can be name, email, phone etc |
| page | 1 | page as interger number |


<!--- End query param items -->

<!--- Body mode -->

<!--- Raw body data -->



<!---End Raw body data -->

<!---FormData -->

<!---End FormData -->


<!---x-urlencoded data -->

<!---End x-urlencoded data -->

<!--- End Body mode -->


<!--- Items response -->


<!--- End Iterate collection items -->

###### 2. Create teacher



```bash
Method: POST
Type: formdata
URL: {{base_url}}/teachers
```

<!--- headers items -->

***Headers:***

<!--- Iterate headers items -->
| Key | Value | Description |
| --- | ------|-------------|
| Authorization | Bearer {{access_token}} | `access_token` must be a valid *oAuth2* access token. You can also use ~~custom token~~ but it will be removed from next version |

<!--- End Iterate headers items -->

<!--- End  headers items -->


<!--- Query param items -->

<!--- End query param items -->

<!--- Body mode -->

<!--- Raw body data -->

<!---End Raw body data -->

<!---FormData -->

<!--- Formdata items -->

***Body:***

| Key | Value | Description |
| --- | ------|-------------|
| name | John Doe | `name` field must be between *3 to 20* chars |
| age | 33 | `age` field must be a valid numeric value |



<!---End FormData -->


<!---x-urlencoded data -->

<!---End x-urlencoded data -->

<!--- End Body mode -->


<!--- Items response -->

***Response***

        
Status: Validation Error | Code: 422



```js
{
    "errors": [
        "Name field is required",
        "age field is required"
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



<!--- End Items response -->


<!--- End Iterate collection items -->

###### 3. Update teacher


update a teacher using api


```bash
Method: PUT
Type: urlencoded
URL: {{base_url}}/teachers
```

<!--- headers items -->

***Headers:***

<!--- Iterate headers items -->
| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/x-www-form-urlencoded |  |

<!--- End Iterate headers items -->

<!--- End  headers items -->


<!--- Query param items -->

<!--- End query param items -->

<!--- Body mode -->

<!--- Raw body data -->

<!---End Raw body data -->

<!---FormData -->

<!---End FormData -->


<!---x-urlencoded data -->

***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| name | John Doe |  |
| age | 33 |  |



<!---End x-urlencoded data -->

<!--- End Body mode -->


<!--- Items response -->


<!--- End Iterate collection items -->

###### 4. Update teacher partially


This api need header for update teacher


```bash
Method: PATCH
Type: urlencoded
URL: {{base_url}}/teachers
```

<!--- headers items -->

***Headers:***

<!--- Iterate headers items -->
| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/x-www-form-urlencoded |  |
| Authorization | Bearer {{access_token}} |  |

<!--- End Iterate headers items -->

<!--- End  headers items -->


<!--- Query param items -->

<!--- End query param items -->

<!--- Body mode -->

<!--- Raw body data -->

<!---End Raw body data -->

<!---FormData -->

<!---End FormData -->


<!---x-urlencoded data -->

***Body:***


| Key | Value | Description |
| --- | ------|-------------|
| name | John Doe |  |
| age | 33 |  |



<!---End x-urlencoded data -->

<!--- End Body mode -->


<!--- Items response -->


<!--- End Iterate collection items -->

###### 5. Remove teacher


update a teacher using api


```bash
Method: DELETE
Type: 
URL: {{base_url}}/teachers/{{id}}
```

<!--- headers items -->


<!--- Query param items -->

<!--- End query param items -->

<!--- Body mode -->


<!--- Items response -->


<!--- End Iterate collection items -->


<!--- End Iterate main collection -->

### Teacher/v2


<!--- Iterate collection items -->


###### 1. Fetch teachers


Get list of teachers


```bash
Method: GET
Type: raw
URL: {{base_url}}/v2/teachers
```

<!--- headers items -->


<!--- Query param items -->

***Query params:***

<!--- Query param items -->
| Key | Value | Description |
| --- | ------|-------------|
| q | john | q can be name, email, phone etc |
| page | 1 | page as interger number |


<!--- End query param items -->

<!--- Body mode -->

<!--- Raw body data -->



<!---End Raw body data -->

<!---FormData -->

<!---End FormData -->


<!---x-urlencoded data -->

<!---End x-urlencoded data -->

<!--- End Body mode -->


<!--- Items response -->


<!--- End Iterate collection items -->


<!--- End Iterate main collection -->

### Default


<!--- Iterate collection items -->


###### 1. Login


Inorder to access the private ***API*** you must get an access token by providing `username/password`


```bash
Method: POST
Type: formdata
URL: {{base_url}}/login
```

<!--- headers items -->


<!--- Query param items -->

<!--- End query param items -->

<!--- Body mode -->

<!--- Raw body data -->

<!---End Raw body data -->

<!---FormData -->

<!--- Formdata items -->

***Body:***

| Key | Value | Description |
| --- | ------|-------------|
| username | user | username email/phone |
| password | pass | pass must be greater than `5` chars |



<!---End FormData -->


<!---x-urlencoded data -->

<!---End x-urlencoded data -->

<!--- End Body mode -->


<!--- Items response -->


<!--- End Iterate collection items -->


<!--- End Iterate main collection -->


---
[Back to top](#SMS)
> ___Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam)___
