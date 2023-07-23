# Solution to the test
## Description of your solution

> This is a simple web server that returns a json with the analysis of the required data as mentioned in the test. The code is written in `Go (v1.18)`. The solution takes one API endpoint as a GET request, no other type of request are allowed (GET, PUT, PATCH, etc).

> The endpoint is `/analysis` and it takes two query parameters: `duration` and `dimension`. The `duration` parameter is required and it takes a string with the format `Xs` where `X` is a number and `s` is the unit of time (seconds, minutes or hours). The `dimension` parameter is also required and takes a string with the name of the dimension to be analyzed. The possible values are: `comments`, `likes`, `retweets` and `favorites`.

Here is an example output a valid request for the following url: `http://localhost:8080/analysis?duration=5s&dimension=likes`

```json
{
    "average_likes": 37640,
    "total_post": 26,
    "minimum_timestamp": 1493133950,
    "maximum_timestamp": 1690109306
} 
```

## Reasons behind my choices:

The first choice I had to make after knowing what to build was what language I have to code in, I chose Go because it is a good language for this kind of projects because it is fast and efficient. It is also part of your tech stack and can be integrated easily into your microservice architecture.

I chose to use the standard library because it is the most efficient way to build a web server in Go. I also used the standard library to parse the query parameters and to parse the json response.

I decided to use this type of file structure for my code, it is simple and can be easily maintained if new data types or new data types have to be aggregated.

```
├───.github
│   └───workflows
├───helpers
├───packages
├───routes
└───test
```

 ## Future Additions

1. I would have used go channels to make the code more concurrent and efficient.
2. I would deploy the application on kubernetes and use a load balancer to handle the requests load.
3. Add more test and create a test coverage report.
4. I would have made the Server Reader and Aggregator 2 different applications so it can be more easily scalable. 

## how to run your project.

### Through the Docker
The image is on docker hub so you can run it with the following command:
1. `docker pull melol/upf-test-img`
2. `docker run --publish 8080:8080 melol/upf-test-img`
3. Use postman/curl/web browser to make a GET request to `http://localhost:8080/analysis?duration=5s&dimension=likes`.

> P.S: I did test the docker image on my other PC with a different OS that I developed on and it worked fine :-).