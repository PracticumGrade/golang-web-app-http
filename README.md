# web-app-http
Код для курса "Создание основы веб-приложения с помощью стандартной библиотеки Go"

```Makefile
~$ make dep_tidy # to update project deps
~$ make build    # to build project
~$ make tests    # to run project tests
```

## Endpoints
 * `GET /printer/hello` - text;
 * `GET /printer/api` - text;
 * `GET /printer/version` - text, current app version from `Makefile`;
 * `GET /printer/randomnumber` - random integer as text;
 * `GET,POST,PUT,DELETE /restapi/users` - JSON, `UserData`;
 * `GET /restapi/users/filtered` - JSON, `UserData`, with params:
   * `page` - integer, shows the number of page with limited amount of entries;
   * `limit` - integer, limits amount of entries;
   * `name` - string, filters entries by `UserData.UserEmail`;
   * `email` - string, filters entries by `UserData.UserName`.

### UserData
```go
type UserData struct {
	UserID    int64  `json:"user_id,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	UserEmail string `json:"user_email,omitempty"`
}
```