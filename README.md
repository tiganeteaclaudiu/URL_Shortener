# URL Shortener
### Complete package used to host URL Shortening/Redirection service, along with all requirements.

Services ran using `docker-compose`:

- **GRPC** Service which handles CRUD operations for URL Shortening
- **HTTP** Service using multiplexing to handle HTTP requests on different import, but same `go` instance
- **Redis** instance, configurable using `redis.conf` file
- **Nginx** reverse proxy, used to handle all incoming requests on a single port and interface
- Optional: **GRPC UI** instance which acts as a playground for the GRPC Service
- Optional: **Documentation** generated using `protoc-gen-doc`, which is always updated to latest `.proto` specs
- Optional: **Redis Commander** instance, used to debug and test the Redis instance

### Setup and running

##### Requirements:

- Docker: https://docs.docker.com/get-docker/
- Docker-compose (version 1.28.0 or later): https://docs.docker.com/compose/install/

##### Installation (Including all optional services determined by "testing" profile):
```
docker-compose --profile testing up
```

##### Installation (Excluding optional services such as Redis Commander):
```
docker-compose up
```

##### Reverse proxy should listen to port `40`.

### Usage

A single host is exposed (`localhost:40`), services should be accessible via:

- `/api` for GRPC calls
- `/<shortenedURLkey>` for redirecting to previously shortened URL (Ex. key: `nlkajsXsakPP`)
- `/ui` for accessing GRPC UI instance (if enabled via the `--profile testing` parameter on `docker-compose up`)
- `/docs/` for accessing GRPC service documentation

### Configuration

`.env` file can be used to configure most parameters of services ran:

- **REDIS_PORT**: defines port at which Redis instance will run. Default: `6379`
- **REDIS_URL**: defines host of Redis service within docker-compose network. Default: `redis:6379`
- **REDIS_COMMANDER_PORT**: defines port at which Redis Commander will be exposed. Default: `8081`
- **REDIS_HOSTS**: defines port at which Redis Commander can communicate with Redis. Default: `local:redis:6379`
- **SERVICE_PORT** (**WORK IN PROGRESS**): defines port at which service will run. Not exposed, but only used by Nginx. Default: `4040`
- **ENABLE_GRPC_UI**: for value `yes`, enables GRPC UI instance which acts as playground for GRPC service
- **EXPOSE_DOCS**: for value `yes`, exposes GRPC Documentation at `/docs/` path

#### `nginx/nginx.conf` file can also be used further to configure reverse proxy.

### Redis Cache Configuration

`redis.conf` file can be used to configure Redis instance.
By default, persistence is enabled via Docker volumes, using RDB (Redis Database).

This means results are not saved on each request, but rather periodically, based on the recent number of requests.
The policy enforced for this can be found in the same configuration file, under the `save` directives.
https://redis.io/topics/persistence

For backup, `.rdb` file generated on Redis backup can be found withing docker volume created for Redis service.

## GRPC Service Documentation

##### See `DOCS.md` file for documentation generated based on `.proto` file.






