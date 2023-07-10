# Golang API

## Project setup


``` bash
go get -u github.com/gorilla/mux
```

``` bash
go build
./api-rest-project
```

## Endpoints

### List items
``` bash
GET items
```

### Create Book
``` bash
POST items
```

### Delete Book
``` bash
DELETE items/{id}
```

### Update Item
``` bash
PUT items/{id}
```

```
# Request example :

# {
#   "id":"1",
#   "name":"item name",
# }
```
