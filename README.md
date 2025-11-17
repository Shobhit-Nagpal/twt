# twt

> Tweet from the terminal

A simple CLI tool to post tweets to X (Twitter) without leaving your terminal.

## Installation
```bash
go install github.com/Shobhit-Nagpal/twt@latest
```

Make sure `~/go/bin` is in your `PATH`.

## Setup

### 1. Create an X Developer App

1. Go to [developer.x.com](https://developer.x.com)
2. Create a new app with **read and write** permissions
3. Generate your API keys and access tokens

### 2. Create Config File

Create `~/.config/twt/config.json` with your credentials:
```json
{
  "USERNAME": "your_twitter_handle",
  "API_KEY": "your_api_key",
  "API_KEY_SECRET": "your_api_key_secret",
  "ACCESS_TOKEN": "your_access_token",
  "ACCESS_TOKEN_SECRET": "your_access_token_secret"
}
```

## Usage
```bash
# Post a tweet
twt post "Hello from the terminal!"

# View help
twt --help
```

## License

MIT
