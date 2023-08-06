# Golang API

## Project setup

``` bash
go build
./api-rest-project
```

## Endpoints

### List items
``` bash
GET items
```

### Create item
``` bash
POST items
```

### Delete item
``` bash
DELETE items/{id}
```

### Update item
``` bash
PUT items/{id}
```

### Curl examples
```
curl http://localhost:8000/items
curl -X POST -H "content-type: application/json" http://localhost:8000/items -d '{"id": 4, "name": "mouse"}'
curl -X PUT -H "content-type: application/json" http://localhost:8000/item{id} -d ''
curl -X DELETE -H "content-type: application/json" http://localhost:8000/item{id} -d ''
```
