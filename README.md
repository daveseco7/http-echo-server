# http-echo-server

Simple http server that outputs headers and bodies of requests to stdout. It's also possible to configure duration of a
request execution by increasing or decreasing the timeout config for the dummy workload.

### Endpoints available:

#### GET

- /success returns 200 - ok
- /error400 returns 400 - bad request
- /error500 returns 500 - internal server error

#### POST

- /success returns 200 - ok
- /error400 returns 400 - bad request
- /error500 returns 500 - internal server error

#### PUT

- /success returns 200 - ok
- /error400 returns 400 - bad request
- /error500 returns 500 - internal server error
