<center>
<div>
  <ul>
  {{ if .err }}
    <li>{{ .err }}</li>
  {{ else }}
    <p>Logging in successfully...</p>
    <li>Owner: {{ $.json.Owner }}</li>
    <ul>
      <li>Account: {{ $.json.Account }}</li>
    </ul>
  {{ end }}
  </ul>
</div>
</center>
