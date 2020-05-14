
# Docgen Examples

A collection of APIs to showcase docgen.

## Indices

* [/feature-requests](#1-feature-requests)
* [/home](#2-home)


--------

## Routes

### /feature-request


Request a new feature.


```bash
Method: POST
Type: 
URL: http://localhost:3000/api/feature-request
```

***Request Body:***

```js        
{
	"title": "Feature title",
	"description": "Feature description."
}
```

***Examples:***


##### I. Request a new feature

***URL:***
```bash
http://localhost:3000/api/feature-request
```

***Request Headers:***

| Key | Value | Description |
| --- | ------|-------------|
| Content-Type | application/json |  |

***Request Body:***

```js        
{
	"title": "More example details",
	"description": "Include the URL, HTTP method, query params, and headers for examples"
}
```

***Response Code:*** 201


***Response Body:***

```js
{
	"message": "Success"
}
```

<br>



### 2. /home

Home page where the latest news about the project can be found.

```bash
Method: GET
Type: 
URL: http://localhost:3000/api/home
```

***Examples:***


##### I. Get the Latest News


***Request Query:***

| Key | Value | Description |
| --- | ------|-------------|
| news | latest |  |


***Response Code:*** 200


***Response Body:***
```js
{
	"data": {
		"news": [
			{
				"title": "Welcome to Docgen!"
			}
		]
	}
}
```

<br>



---
[Back to top](#docgen-examples)
> Made with &#9829; by [thedevsaddam](https://github.com/thedevsaddam) | Generated at: 2020-05-14 07:07:11 by [docgen](https://github.com/thedevsaddam/docgen)
