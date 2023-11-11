
curl http://localhost:8080/hello
curl http://localhost:8080/hello -XPOST
curl http://localhost:8080/hello -XPUT

//body
curl http://localhost:8080/hello -XPOST --data "p" // for test data 
curl http://localhost:8080/hello\?name\=kha -XPOST --data "data"