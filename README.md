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
