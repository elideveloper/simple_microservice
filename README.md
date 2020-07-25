# simple_microservice

From _current_ path
docker build .
docker run -a stdin -a stdout -it -p 8080:8080 _dockerimage_

POST /create
```json
{
	"first_name":"Mihail",
	"last_name" : "Petrovich",
	"age":33
}
```

GET /get?user_id=1