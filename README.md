# Bet API

## Overview

Bet API is a Go-based project designed to handle betting operations efficiently.using native http library with some minimal functions that let's users to make a bet for a event,settle the bet for the event and check users balances

## Assumptions

- users ara pre registerd (there is no user auth or create)
- events ara already added
- setlling the bets are done by sending a post call to the /settle_bet endpoint using the event id and the result win/lose
-

## Features

- User authentication and management
- Bet placement and tracking
- Mutual Exclution
- Scalable architecture

## Installation

1. Clone the repository:

```bash
git clone https://github.com/Kushan2k/bet-api-golang.git
```

2. Navigate to the project directory:

```bash
cd bet_api
```

3. Install dependencies:

```bash
go mod tidy
```

## Usage

1. Start the server:

```bash
go run cmd/main.go
```

2. Access the API at `http://localhost:8080`.

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


Thank you