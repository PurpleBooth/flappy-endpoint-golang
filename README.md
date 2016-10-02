# A flappy endpoint: Let's go!

A go version of the [flappy endpoint](https://github.com/PurpleBooth/flappy-endpoint) from [Ben Wain](https://github.com/LoveSoftware)'s [Logstash talk](https://github.com/LoveSoftware/application-logging-with-logstash) which is useful for training. 

It returns a 200, 404, 500 or a slow response at random intervals.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development. I recommend 
running this on a docker based environment, so your prod is the same as your dev.

### Prerequisities

You'll need [docker](https://www.docker.com/products/overview).

### Trying it out

Then run the application with
```
docker-compose up
```

You can then see the application running by

```
curl http://$(boot2docker ip):8080/flappy
Slow response
```

## Contributing
Feel free to submit pull requests and issues. If it's a particularly large PR, you may wish to discuss it in an issue first.

Please note that this project is released with a [Contributor Code of Conduct](CONDUCT.md). By participating in this project you agree to abide by its terms.

## Versioning

We use [SemVer](http://semver.io/) for the version tags available See the tags on this repository. 

## Authors

See the list of [contributors](https://github.com/PurpleBooth/flappy-endpoint-golang/contributors) who participated in this project.

