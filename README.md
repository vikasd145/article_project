 To format proto file 
 
 clang-format -style=file -i pkg/proto/article_db.proto 
 
 ### `scripts`
 
 contains the docker and docker compose file
 
 ### `more can do`
 
can apply rate limit each api so that server does not go down in high traffic just drop request
 
can also use redis to cache the request for id
 
 