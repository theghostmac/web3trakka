# web3trakka

## Running the app
Explanation:
- `build`: Compiles the application locally using the Go compiler.
- `run`: Builds and then runs the application locally.
- `docker-build`: Builds the Docker image using `docker-compose`.
- `docker-up`: Builds (if necessary) and starts the Docker containers as defined in `docker-compose.yml`.
- `docker-down`: Stops the Docker containers.
- `clean`: Removes the built binary to clean up the project directory.

To use this `Makefile`:

Run the application locally: `make run`
Run the application in Docker: `make docker-up`
Stop the Docker containers: `make docker-down`
Clean up the built files: `make clean`

## Available commands
- You can track a crypto pair using the `track` command.
- You can set price alerts to buy a pair at a certain price using the `alert` command.
- You can also view your portfolio using the `view` command.

To see how a command works, run this to show help:
```shell
./web3trakka <command> help
```

## Tracking Crypto Pairs
To track a crypto pair, say ETHUSDT, just run:
```shell
make build
# then
./web3trakka track ETHUSDT
```
