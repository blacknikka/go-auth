<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>{{ .Title }}</title>
  </head>
  <body>
    <form action="/upload" enctype="multipart/form-data" method="post">
      <input type="file" name="upload" id="upload" multiple="multiple">
      <input type="submit" value="Upload file" />
    </form>
  </body>
</html>