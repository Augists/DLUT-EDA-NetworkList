<center>
  <div>
    <style>
      .btn {
        display: inline-block;
        padding: 6px 12px;
        margin-bottom: 0;
        font-size: 14px;
        font-weight: 400;
        line-height: 1.42857143;
        text-align: center;
        white-space: nowrap;
        vertical-align: middle;
        -ms-touch-action: manipulation;
        touch-action: manipulation;
        cursor: pointer;
        -webkit-user-select: none;
        -moz-user-select: none;
        -ms-user-select: none;
        user-select: none;
        background-image: none;
        border: 1px solid transparent;
        border-radius: 4px;
      }
    </style>
  {{ if .get }}
    <form action="/delete" method="post">
      <input type="text" name="account" value="201992222">
      <button class="btn" type="submit">Delete</button>
    </form>
  {{ else }}
    {{ if .err }}
    <p style="color: red;">{{ .err }}</p>
    {{ else }}
    <p style="color: green;">{{ .delete }}</p>
    {{ end }}
    <button class="btn" onclick="window.location.href='/'">Back</button>
  {{ end }}
  </div>
</center>
