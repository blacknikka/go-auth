<!DOCTYPE html>
<html>
<body>
    <table>
    <thead>
        <tr>
            <th>ID</th>
            <th>名前</th>
            <th>Email</th>
        </tr>
    </thead>
    <tbody>
        {{range $user := .}}
        <tr>
            <td>{{$user.ID}}</td>
            <td>{{$user.Name}}</td>
            <td>{{$user.Email}}</td>
        </tr>
        {{end}}
    </tbody>
</body>
</html>
