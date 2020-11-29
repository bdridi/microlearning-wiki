# Workcale Microlearning webapp

Microlearning wiki service is a the wiki microservice which is a part of the workcale microlearning project.

### Service overview

- Search random wiki pages by category


### Start the server

`go run main.go`

### API documentation

- request : `http://localhost:8081/api/v1/wiki?category=health`
- response :

    ```json
    [{
        "title": "lorem"
        "category":  "ipsum"
        "url" : "http://wikipedia/lorem.html"
    }]
    ```

### Build docker image 

`docker build -t wiki-service .`
