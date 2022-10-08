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
<center>
  {{ if .err }}
    <div class="alert alert-danger" role="alert">
      Error Message: {{ .err }}
    </div>
  {{ end }}
  {{ if .id }}
    <div class="alert alert-success" role="alert">
      <h4 class="alert-heading">Success!</h4>
      <p>Successfully add a new account with ID: {{ .id }}</p>
    </div>
  {{ end }}
  <button class="btn" onclick="window.location.href='/listall'">List All Account</button><br><br>
  <button class="btn" onclick="window.location.href='/list'">Log in DLUT-EDA by a Random Account</button><br><br>
  <button class="btn" onclick="window.location.href='/add'">Add Account</button><br><br>
  <button class="btn" onclick="window.location.href='/delete'">Delete Account</button><br><br>
</center>
</div>
