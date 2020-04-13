# Discord bot

## env variables

```shell
BOT_DISCORD_TOKEN - discord token - required
BOT_COMMAND_PREFIX - all bot commands prefix (example bot) - required
BOT_INSTAGRAM_ENABLED - true or false to switch instagram bot features
BOT_INSTAGRAM_USERNAME
BOT_INSTAGRAM_PASSWORD
BOT_INSTAGRAM_PASSWORD
BOT_N2YO_ENABLED - true or false to switch ny2o bot features
BOT_N2YO_API_KEY
BOT_CACHE_TYPE - memory or redis (default memory)
BOT_REDIS_URL - required if BOT_CACHE_TPE=redis
BOT_SCHEDULED_MESSAGES_ENABLED  - true or false to switch sending scheduled messages
BOT_SCHEDULED_CONFIG_FILE_URL=config.yaml
BOT_DATABASE_URL
BOT_BURZE_DZIS_API_KEY
```

## Deploy

```shell
heroku container:push --recursive
heroku container:release web worker
```
