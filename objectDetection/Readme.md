# object recognition server [ go v1.11.2 ]

Build the image.
```
$ docker build -t recognize .
```

Run service in a container.
```
$ docker run -p 8080:8080 --rm recognize
```

Call the service.
```
curl -X POST -F "image=@"./path to a local image"" http://localhost:8080/recognize
```
