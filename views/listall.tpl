<center>
<div>
  <ul>
  {{ if .err }}
    <li>{{ .err }}</li>
  {{ else }}
    {{ with .json }}
    {{ range . }}
    <li>Owner: {{ .Owner }}</li>
    <ul>
      <li>Account: {{ .Account }}</li>
    </ul>
    {{ end }}
    {{ end }}
  {{ end }}
  </ul>
</div>
</center>
