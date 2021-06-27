# test path ["/"]
```
curl http://localhost:9999/  
```
result
```
url.path = "/"
```

# test path ["/hello"]
```
curl http://localhost:9999/hello
```
result
```
header["User-Agent"]=["curl/7.71.1"]
header["Accept"]=["*/*"]
```