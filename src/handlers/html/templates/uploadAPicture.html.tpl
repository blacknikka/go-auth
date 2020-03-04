<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <title>{{ .Title }}</title>
  </head>
  <body>
    <form action="/file/upload_request" enctype="multipart/form-data" method="post">
      <input type="text" name="file_name" id="name" multiple="multiple">
      <input type="file" name="up_data" id="upload" multiple="multiple">
      <input type="submit" value="Upload file" />
    </form>
  </body>
</html>
