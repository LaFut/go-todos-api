# go-todos-api
```shell script
docker-composer up -d
```

swagger: http://localhost:8080/doc/swagger.json

### get todos
```shell script
curl -v -H "Content-Type: application/json"  localhost:8080/api/todos
```
### get todo
```shell script
curl -v -H "Content-Type: application/json"  localhost:8080/api/todos/:id
```
### add
```shell script 
curl -d '{  "name": "test",   "parentid": 0 }' -H 'Content-Type: application/json'  localhost:8080/api/todos
```
### change
```shell script
curl -d '{  "name": "test",   "parentid": 0 }' -H 'Content-Type: application/json' -X PUT  localhost:8080/api/todos/:id
```
### delete
```shell script
curl -X DELETE  localhost:8080/api/todos/:id
```

