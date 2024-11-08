# 🐊 GATOR
A blog aggreGATOR

## Config
Add `~/.gatorconfig.json` to your home directory and populate with db config settings:
```json
{
  "db_url": "connection_string_goes_here",
  "current_user_name": "username_goes_here"
}
```
## Install
From project root:
```bash
go install
```

## Run
From project root:
```bash
gator <command> [args...]
```