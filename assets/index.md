@{{- if .Data.Info.Name -}}@
# @{{ .Data.Info.Name }}@
@{{ end }}@
@{{ if .Data.Info.Description }}@
@{{- .Data.Info.Description -}}@
@{{ end }}@

## Indices
@{{- range $index, $c := .Data.Collections }}@
##### [@{{ $c.Name }}@](@{{ $c.Name | snake }}@)
    @{{ range $i, $item := $c.Items }}@
* [@{{ $item.Name }}@](@{{ $c.Name | snake }}@-@{{ $item.Name | snake }}@)

    @{{ end }}@
@{{ end }}@

--------

@{{ range $di, $d := .Data.Collections }}@
### [@{{ $d.Name }}@](@{{ $d.Name | snake }}@)
@{{ $d.Description }}@

@{{ end }}@
