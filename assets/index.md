<!--- Collection name and description -->
@{{if .Data.Info.Name -}}@
# @{{ .Data.Info.Name | trim }}@
@{{ end }}@
@{{ if .Data.Info.Description }}@
@{{- .Data.Info.Description -}}@
@{{ end }}@

<!--- Table of contents -->
@{{- $numCollections := len .Data.Collections}}@
<!--- If we have only one group/collection, then no need for the "ungrouped" heading -->
@{{- if eq $numCollections 1 }}@ 
@{{- range $index, $c := .Data.Collections -}}@
@{{- range $i, $item := $c.Items }}@
1. [@{{ $item.Name | trim }}@](#@{{ merge $i $item.Name | trim | glink | glinkInc }}@)
@{{- range $ri, $response := $item.Responses }}@
@{{- $riRoman := ($ri | addOne | roman) }}@
   1. [@{{ $response.Name | trim }}@](#@{{ (print $riRoman ". Example Request: " $response.Name) | trim | glink | glinkInc }}@)
@{{- /* End iterate responses for the request */}}@
@{{- end }}@
@{{- /* End iterate requests in the collection */}}@
@{{- end }}@
@{{- /* End iterate collections */}}@
@{{- end }}@
@{{- /* End if we have more than one collection */}}@
@{{- end }}@

<!--- Variables --->
@{{ if .Data.Variables }}@
## Variables

<!--- Iterate variables -->
| Key | Value | Type |
| --- | ------|-------------|
@{{ range $ih, $v := .Data.Variables -}}@
| @{{ $v.Key }}@ | @{{ $v.Value }}@ | @{{ $v.Type }}@ |
@{{ end }}@
<!--- End Iterate headers items -->

<!--- End  headers items -->
@{{ end }}@
<!--- End Variables --->

## Endpoints

<!--- If we have more than one group/collection, then display each group name heading -->
@{{- if gt $numCollections 1 }}@ 
@{{- range $index, $c := .Data.Collections }}@
* [@{{ $c.Name | trim }}@](#@{{ $c.Name | trim | glink }}@)
@{{- range $i, $item := $c.Items }}@
    1. [@{{ $item.Name | trim }}@](#@{{ merge $i $item.Name | trim | glink | glinkInc }}@)
@{{- range $ri, $response := $item.Responses }}@
@{{- $riRoman := ($ri | addOne | roman) }}@
        * [@{{ $response.Name | trim }}@](#@{{ (print $riRoman ". Example Request: " $response.Name) | trim | glink | glinkInc }}@)
@{{- /* End iterate responses for the request */}}@
@{{- end }}@
@{{- /* End iterate requests in the collection */}}@
@{{- end }}@
@{{- /* End iterate collections */}}@
@{{- end }}@
@{{- /* End if we have more than one collection */}}@
@{{- end }}@

--------
<!--- Iterate main collection -->

@{{ range $di, $d := .Data.Collections }}@
<!--- Only show the collection name if there is more than the "ungrouped" collection -->
@{{ if gt $numCollections 1 }}@
## @{{ $d.Name | trim  }}@
@{{ end }}@
@{{ $d.Description }}@

<!--- Iterate collection items -->

@{{ range $ii, $item := $d.Items }}@
### @{{ $ii | addOne }}@. @{{ if $item.Name }}@@{{ $item.Name | trim }}@@{{ end }}@

@{{ if $item.Request.Description }}@
@{{ $item.Request.Description }}@
@{{ end }}@

***Endpoint:***

```bash
Method: @{{ $item.Request.Method | upper }}@
Type: @{{ $item.Request.Body.Mode | upper }}@
URL: @{{ $item.Request.URL.Raw | trimQueryParams | e }}@
```

<!--- headers items -->
@{{ if $item.Request.Headers }}@
***Headers:***

<!--- Iterate headers items -->
| Key | Value | Description |
| --- | ------|-------------|
@{{ range $ih, $h := $item.Request.Headers -}}@
| @{{ $h.Key | e }}@ | @{{ $h.Value | e }}@ | @{{ $h.Description | e }}@ |
@{{ end }}@
<!--- End Iterate headers items -->

<!--- End  headers items -->
@{{ end }}@

<!--- Query param items -->
@{{ if $item.Request.URL.Query }}@
***Query params:***

<!--- Query param items -->
| Key | Value | Description |
| --- | ------|-------------|
@{{ range $iq, $q := $item.Request.URL.Query -}}@
| @{{ $q.Key | e }}@ | @{{ $q.Value | e }}@ | @{{ $q.Description | e }}@ |
@{{ end }}@
@{{ end }}@
<!--- End query param items -->

<!--- URL variables items -->
@{{ if $item.Request.URL.Variables }}@
***URL variables:***

<!--- URL variables items -->
| Key | Value | Description |
| --- | ------|-------------|
@{{ range $iq, $q := $item.Request.URL.Variables -}}@
| @{{ $q.Key | e }}@ | @{{ $q.Value | e }}@ | @{{ $q.Description | e }}@ |
@{{ end }}@
@{{ end }}@
<!--- End URL variables items -->

<!--- Body mode -->
@{{ if $item.Request.Body.Mode}}@
<!--- Raw body data -->
@{{ if eq $item.Request.Body.Mode "raw"}}@
@{{ if $item.Request.Body.Raw }}@
***Body:***

```js        
@{{ $item.Request.Body.Raw }}@
```
@{{ end }}@
@{{ end }}@
<!---End Raw body data -->

<!---FormData -->
@{{ if eq $item.Request.Body.Mode "formdata"}}@
<!--- Formdata items -->
@{{ if $item.Request.Body.FormData }}@
***Body:***

| Key | Value | Description |
| --- | ------|-------------|
@{{ range $if, $f := $item.Request.Body.FormData -}}@
| @{{ $f.Key | e }}@ | @{{ $f.Value | e }}@ | @{{ $f.Description | e }}@ |
@{{ end }}@
@{{ end }}@
@{{ end }}@
<!---End FormData -->


<!---x-urlencoded data -->
@{{ if eq $item.Request.Body.Mode "urlencoded"}}@
***Body:***

@{{ if $item.Request.Body.URLEncoded }}@
| Key | Value | Description |
| --- | ------|-------------|
@{{ range $iu, $u := $item.Request.Body.URLEncoded -}}@
| @{{ $u.Key | e }}@ | @{{ $u.Value | e }}@ | @{{ $u.Description | e }}@ |
@{{ end }}@
@{{ end }}@
@{{ end }}@
<!---End x-urlencoded data -->

<!--- End Body mode -->
@{{ end }}@

<!--- Items response -->
@{{ if $item.Responses }}@
***More example Requests/Responses:***
@{{ range $ir, $resp := $item.Responses }}@
@{{ if $resp.Name }}@
#### @{{ $ir | addOne | roman }}@. Example Request: @{{ $resp.Name }}@

<!--- headers items -->
@{{ if $resp.OriginalRequest.Headers }}@
***Headers:***

<!--- Iterate headers items -->
| Key | Value | Description |
| --- | ------|-------------|
@{{ range $ih, $h := $resp.OriginalRequest.Headers -}}@
| @{{ $h.Key | e }}@ | @{{ $h.Value | e }}@ | @{{ $h.Description | e }}@ |
@{{ end }}@
<!--- End Iterate headers items -->

<!--- End  headers items -->
@{{ end }}@


<!--- query items -->
@{{ if $resp.OriginalRequest.URL.Query }}@
***Query:***

<!--- Iterate query items -->
| Key | Value | Description |
| --- | ------|-------------|
@{{ range $ih, $h := $resp.OriginalRequest.URL.Query -}}@
| @{{ $h.Key | e }}@ | @{{ $h.Value | e }}@ | @{{ $h.Description | e }}@ |
@{{ end }}@
<!--- End Iterate query items -->

<!--- End  query items -->
@{{ end }}@

<!--- url variable items -->
@{{ if $resp.OriginalRequest.URL.Variables }}@
***Query:***

<!--- Iterate url variable items -->
| Key | Value | Description |
| --- | ------|-------------|
@{{ range $ih, $h := $resp.OriginalRequest.URL.Variables -}}@
| @{{ $h.Key | e }}@ | @{{ $h.Value | e }}@ | @{{ $h.Description | e }}@ |
@{{ end }}@
<!--- End Iterate url variable items -->

<!--- End url variable items -->
@{{ end }}@

<!--- Body mode -->
@{{ if ne $resp.OriginalRequest.Body.Mode "" }}@
<!--- Raw body data -->
@{{ if eq $resp.OriginalRequest.Body.Mode "raw"}}@
@{{ if $resp.OriginalRequest.Body.Raw }}@
***Body:***

```js        
@{{ $resp.OriginalRequest.Body.Raw }}@
```
@{{ end }}@
@{{ end }}@
<!---End Raw body data -->

<!---FormData -->
@{{ if eq $resp.OriginalRequest.Body.Mode "formdata"}}@
<!--- Formdata items -->
@{{ if $resp.OriginalRequest.Body.FormData }}@
***Body:***

| Key | Value | Description |
| --- | ------|-------------|
@{{ range $if, $f := $resp.OriginalRequest.Body.FormData -}}@
| @{{ $f.Key | e }}@ | @{{ $f.Value | e }}@ | @{{ $f.Description | e }}@ |
@{{ end }}@
@{{ end }}@
@{{ end }}@
<!---End FormData -->


<!---x-urlencoded data -->
@{{ if eq $item.Request.Body.Mode "urlencoded"}}@
***Body:***

@{{ if $resp.OriginalRequest.Body.URLEncoded }}@
| Key | Value | Description |
| --- | ------|-------------|
@{{ range $iu, $u := $resp.OriginalRequest.Body.URLEncoded -}}@
| @{{ $u.Key | e }}@ | @{{ $u.Value | e }}@ | @{{ $u.Description | e }}@ |
@{{ end }}@
@{{ end }}@
@{{ end }}@
<!---End x-urlencoded data -->

@{{ else }}@
***Body: None***
<!--- End Body mode -->
@{{ end }}@

@{{ if $resp.Body }}@
#### @{{ $ir | addOne | roman }}@. Example Response: @{{ $resp.Name }}@
```js
@{{ $resp.Body }}@
```
@{{ end }}@

***Status Code:*** @{{ $resp.Code }}@

<br>

<!--- Iterate Items response end -->
@{{ end }}@

<!--- if response exist response end -->
@{{ end }}@

<!--- End Items response -->
@{{ end }}@

<!--- End Iterate collection items -->
@{{ end }}@

<!--- End Iterate main collection -->
@{{ end }}@


---
[Back to top](#@{{ .Data.Info.Name | trim | glink }}@)

>Generated at @{{date_time}}@ by [docgen](https://github.com/thedevsaddam/docgen)
