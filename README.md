# Go Pinger

A simple Go application that checks the availability of a remote host using ICMP (ping). It offers easy configuration, a
Telegram bot for notifications, and scheduled host status checks using Goroutines.

## Features

- **Host Monitoring**: Pings a list of remote hosts to check their availability.
- **Scheduled Checks**: Allows periodic status checks using Goroutines.
- **Telegram Bot Notifications**: Sends alerts to a Telegram bot when a host goes down or comes back online.
- **Configuration**: Easy configuration using environment variables (`.env`).

## Installation

### Prerequisites

Before running the project, make sure you have the following installed:

- [Docker](https://www.docker.com/get-started) (for containerization)

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/Phenix123/go-pinger.git
   cd go-pinger
   ```

2. Create a `.env` file in the root of the project with variables from `.env.example`:

3. Build and run the application using Docker:

   ```bash
   docker-compose up --build
   ```

   This will build the Docker image and start the container with the necessary environment configuration from the `.env`
   file.

## Configuration

- **BOT_TOKEN**: Your bot's token, which can be generated through BotFather on Telegram.
- **CHAT_ID**: The chat ID where the bot will send notifications. You can get this by starting a conversation with your
  bot and using the [GetUpdates](https://core.telegram.org/bots/api#getupdates) API.
- **PING_INTERVAL**: The interval (in seconds) between each ping check.
- **HOST**: The hostname that you want to monitor.

## Usage

### Running the Project

To run the project using Docker, execute the following command:

```bash
docker-compose up
```

This will build the Docker container (if needed) and run the application inside it.

### Stopping the Project

To stop the running container:

```bash
docker-compose down
```

### Logs

To view logs for the running container:

```bash
docker-compose logs -f
```

## Development

### Local Development with Docker

1. Clone the repository and navigate to the project directory:

   ```bash
   git clone https://github.com/Phenix123/go-pinger.git
   cd go-pinger
   ```

2. Ensure your `.env` file is configured as desired.

3. Build and run the project in development mode using Docker:

   ```bash
   docker-compose up --build
   ```

   This will rebuild the Docker image if any code changes are made and start the container.

4. Make any code changes as necessary. The project will be reloaded when the container is rebuilt.

### Dockerfile and docker-compose.yml

If you want to customize the Docker setup or environment variables, you can modify the `Dockerfile`
or `docker-compose.yml` to suit your needs.

### Testing

To test the functionality of your project, ensure you have the correct `.env` file setup, and then run the application
locally or inside a Docker container.

### Contribution

If you would like to contribute to this project, please fork the repository, create a feature branch, and submit a pull
request. Contributions are welcome!

### License

This project is licensed under the MIT License.

## Authors

- **Phenix123** - Initial work and development

## Acknowledgements

- [Telegram Bot API](https://core.telegram.org/bots)
- [Go Programming Language](https://golang.org/)

---

## Feature improvement

- Add few hosts for pinger;

---

This README now reflects the Docker setup for both running and developing the project. It provides clear steps on how to
use Docker for easy setup and configuration, as well as for local development.
