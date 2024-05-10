# main go-server with pg

make sure you have go and pg installed.

initiate .env file in ./server folder.
```
API_KEY=RETRIEVE FROM GOOGLE API
DB_HOST=localhost
DB_PORT=5432
DB_USER=yourUser
DB_NAME=dbName
DB_PASSWORD=pass
```

run the go server.
```
go run main.go
```
## endpoints
/signUp:
```
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "email":"example@example.com",
    "profile_picture":1,
    "user_name":"example_user"
    }' \
  http://localhost:8081/signUp
```

/login:
```
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"email":"example@example.com"}' \
  http://localhost:8081/login
```

/org:
```
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"user_id":"your_user_id","org_name":"example_org_name","location":"example_location","description":"example_description"}' \
  http://localhost:8081/org

curl -X GET \
  "http://localhost:8081/org?org_name=example_org_name"
```

/rewardPoints:
```
curl -X PUT \
  -H "Content-Type: application/json" \
  -d '{"userID":"your_user_id","points":100}' \
  http://localhost:8081/rewardPoints

curl -X GET \
  "http://localhost:8081/rewardPoints?userID=your_user_id"

```

/recyclingData/videos:
```
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"objectDetected":"your_object"}' \
  http://localhost:8081/recyclingData/videos
```

/recyclingData/maps:
```
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"latitude":your_latitude,"longitude":your_longitude,"object":"object"}' \
  http://localhost:8081/recyclingData/maps

```