# Go simple proxy

HTTP-server that simply proxies GET-requests. Useful for getting data from locally locked resources.

You must specify the endpoint with GET-parameter `t`, for example:
```bash
curl http://localhost:8080/?t=https://dummyjson.com/todos/random
```

## TODO

* Access restriction by referrer
* CI build stage (remove app building from Dockerfile)
* Optimize error handling
