## GoIDS
GoIDS Is basic  fasthttp server + intrusion detection On Post Method

##### How Does It Work
Any incoming POST request  passes through 78 regex that are included in `filters.json` file. in case of incompatibility Post Body with Rule Filters , Client receives `OK` message .

#### Run Project

```shell script
go run main.go
```

 
```shell script
curl -d 'name=<script>window.onload'  localhost:8080
```

```shell script
curl -d 'name=mehdi-shokohi'  localhost:8080
```