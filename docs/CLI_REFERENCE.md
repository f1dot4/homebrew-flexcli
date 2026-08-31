# FlexCLI Reference

This document is auto-generated. Do not edit manually.

### `flexcli`

```
FlexCLI - FlexCoach Command Line Interface

Configuration precedence (highest to lowest):
  1. Command-line flags (--server, --key)
  2. Environment variables (FLEXCLI_SERVER_URL, FLEXCLI_API_KEY)
  3. Config file context (--config file, or $HOME/.flexcli.json if --config is not set)

Note: --config only selects which file is read; it does not itself take part in the
server/key precedence above.

Usage:
  flexcli [command]

Available Commands:
  admin       System administration commands
  completion  Generate the autocompletion script for the specified shell
  config      Configure CLI settings for an environment
  connect     Manage device connections and system status
  context     Manage environments (contexts)
  help        Help about any command
  plan        Manage training plans
  profile     User Profile Hub

Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
  -h, --help             help for flexcli
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
  -v, --version          version for flexcli

Use "flexcli [command] --help" for more information about a command.
```

### `flexcli admin`

```
System administration commands

Usage:
  flexcli admin [command]

Available Commands:
  backup        Manage system backups
  health-metric Correct or delete a single raw health metric reading
  settings      Manage global system settings
  status        Get system-wide status and health
  sync-all      Trigger background sync for all users
  users         List all user profiles

Flags:
  -h, --help   help for admin

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli admin [command] --help" for more information about a command.
```

### `flexcli admin backup`

```
Manage system backups

Usage:
  flexcli admin backup [command]

Available Commands:
  config      List backup configurations
  create      Trigger immediate backup
  list        List backup history
  set-config  Update a backup configuration

Flags:
  -h, --help   help for backup

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli admin backup [command] --help" for more information about a command.
```

### `flexcli admin backup config`

```
List backup configurations

Usage:
  flexcli admin backup config [flags]

Flags:
  -h, --help   help for config

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli admin backup create`

```
Trigger immediate backup

Usage:
  flexcli admin backup create [flags]

Flags:
  -h, --help   help for create

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli admin backup list`

```
List backup history

Usage:
  flexcli admin backup list [flags]

Flags:
  -h, --help   help for list

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli admin backup set-config`

```
Update a backup configuration

Usage:
  flexcli admin backup set-config [key] [value] [flags]

Flags:
  -h, --help   help for set-config

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli admin health-metric`

```
Correct or delete a single raw health metric reading

Usage:
  flexcli admin health-metric [command]

Available Commands:
  correct     Overwrite one raw health metric reading and re-aggregate that day
  delete      Delete one raw health metric reading and re-aggregate that day

Flags:
  -h, --help   help for health-metric

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli admin health-metric [command] --help" for more information about a command.
```

### `flexcli admin health-metric correct`

```
Overwrite one raw health metric reading and re-aggregate that day

Usage:
  flexcli admin health-metric correct <user_id> <date> <metric_name> <value> [flags]

Flags:
  -h, --help            help for correct
      --source string   Reading source (withings, garmin, manual, telegram) (default "withings")
      --unit string     Unit for the corrected value

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli admin health-metric delete`

```
Delete one raw health metric reading and re-aggregate that day

Usage:
  flexcli admin health-metric delete <user_id> <date> <metric_name> [flags]

Flags:
  -h, --help            help for delete
      --source string   Reading source (withings, garmin, manual, telegram) (default "withings")

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli admin settings`

```
Manage global system settings

Usage:
  flexcli admin settings [command]

Available Commands:
  list             List all global settings
  merge-strategies Get a one-time URL to the health metrics merge strategy config UI
  set              Update a global setting

Flags:
  -h, --help   help for settings

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli admin settings [command] --help" for more information about a command.
```

### `flexcli admin settings list`

```
List all global settings

Usage:
  flexcli admin settings list [flags]

Flags:
  -h, --help   help for list

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli admin settings merge-strategies`

```
Get a one-time URL to the health metrics merge strategy config UI

Usage:
  flexcli admin settings merge-strategies [flags]

Flags:
  -h, --help   help for merge-strategies

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli admin settings set`

```
Update a global setting

Usage:
  flexcli admin settings set [key] [value] [flags]

Flags:
  -h, --help   help for set

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli admin status`

```
Get system-wide status and health

Usage:
  flexcli admin status [flags]

Flags:
  -h, --help   help for status
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli admin sync-all`

```
Trigger background sync for all users

Usage:
  flexcli admin sync-all [flags]

Flags:
  -h, --help            help for sync-all
      --source string   Specific sync source (garmin, withings)

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli admin users`

```
List all user profiles

Usage:
  flexcli admin users [flags]

Flags:
  -h, --help   help for users

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli config`

```
Configure CLI settings for an environment

Usage:
  flexcli config [flags]

Flags:
  -h, --help            help for config
      --key string      API Key
      --name string     Context name (default "default")
      --server string   FlexCoach server URL

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
```

### `flexcli connect`

```
Manage device connections and system status

Usage:
  flexcli connect [command]

Available Commands:
  garmin      Manage Garmin connection and settings
  status      Get system status
  withings    Manage Withings connection and settings

Flags:
  -h, --help   help for connect

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli connect [command] --help" for more information about a command.
```

### `flexcli connect garmin`

```
Manage Garmin connection and settings

Usage:
  flexcli connect garmin [command]

Available Commands:
  config      Manage garmin expert settings

Flags:
  -h, --help   help for garmin

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli connect garmin [command] --help" for more information about a command.
```

### `flexcli connect garmin config`

```
Manage garmin expert settings

Usage:
  flexcli connect garmin config [command]

Available Commands:
  get         Get expert settings
  set         Update expert settings

Flags:
  -h, --help   help for config

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli connect garmin config [command] --help" for more information about a command.
```

### `flexcli connect garmin config get`

```
Get expert settings

Usage:
  flexcli connect garmin config get [flags]

Flags:
  -h, --help   help for get

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli connect garmin config set`

```
Update expert settings

Usage:
  flexcli connect garmin config set [flags]

Flags:
  -h, --help                    help for set
      --interval int            Sync interval in hours
      --lookback-manual int     Days to look back for manual sync
      --lookback-schedule int   Days to look back for scheduled sync

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli connect status`

```
Get system status

Usage:
  flexcli connect status [flags]

Flags:
  -h, --help   help for status
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli connect withings`

```
Manage Withings connection and settings

Usage:
  flexcli connect withings [command]

Available Commands:
  config      Manage withings expert settings

Flags:
  -h, --help   help for withings

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli connect withings [command] --help" for more information about a command.
```

### `flexcli connect withings config`

```
Manage withings expert settings

Usage:
  flexcli connect withings config [command]

Available Commands:
  get         Get expert settings
  set         Update expert settings

Flags:
  -h, --help   help for config

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli connect withings config [command] --help" for more information about a command.
```

### `flexcli connect withings config get`

```
Get expert settings

Usage:
  flexcli connect withings config get [flags]

Flags:
  -h, --help   help for get

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli connect withings config set`

```
Update expert settings

Usage:
  flexcli connect withings config set [flags]

Flags:
  -h, --help           help for set
      --interval int   Sync interval in hours

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli context`

```
Manage environments (contexts)

Usage:
  flexcli context [command]

Available Commands:
  delete      Remove a context
  list        List all contexts
  use         Switch the active context

Flags:
  -h, --help   help for context

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli context [command] --help" for more information about a command.
```

### `flexcli context delete`

```
Remove a context

Usage:
  flexcli context delete [name] [flags]

Flags:
  -h, --help   help for delete

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli context list`

```
List all contexts

Usage:
  flexcli context list [flags]

Flags:
  -h, --help   help for list

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli context use`

```
Switch the active context

Usage:
  flexcli context use [name] [flags]

Flags:
  -h, --help   help for use

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli plan`

```
Manage training plans

Usage:
  flexcli plan [command]

Available Commands:
  activate    Manually activate a training plan
  generate    Generate today's plan (or meso/macro plan)
  get         Get today's plan
  list        List all training plans
  modify      Modify today's plan
  skip        Skip today's plan (or a specific plan by ID)

Flags:
  -h, --help   help for plan

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli plan [command] --help" for more information about a command.
```

### `flexcli plan activate`

```
Manually activate a training plan

Usage:
  flexcli plan activate [plan-id] [flags]

Flags:
  -h, --help   help for activate
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli plan generate`

```
Generate today's plan (or meso/macro plan)

Usage:
  flexcli plan generate [flags]

Flags:
  -h, --help                  help for generate
  -i, --instructions string   Optional instructions for generation
      --json                  Output in JSON format
      --macro                 Generate a macro (4-week) plan
      --meso                  Generate a meso (weekly) plan

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli plan get`

```
Get today's plan

Usage:
  flexcli plan get [flags]

Flags:
  -h, --help   help for get
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli plan list`

```
List all training plans

Usage:
  flexcli plan list [flags]

Flags:
      --end-date string     Only include plans on or before this date (YYYY-MM-DD)
  -h, --help                help for list
      --json                Output in JSON format
      --start-date string   Only include plans on or after this date (YYYY-MM-DD)
  -s, --status string       Filter by status (active, scheduled, inactive)

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli plan modify`

```
Modify today's plan

Usage:
  flexcli plan modify [flags]

Flags:
  -h, --help                  help for modify
  -i, --instructions string   Required instructions for modification
      --json                  Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli plan skip`

```
Skip today's plan (or a specific plan by ID)

Usage:
  flexcli plan skip [plan-id] [flags]

Flags:
  -h, --help            help for skip
      --json            Output in JSON format
      --reason string   Reason for skipping (default "other")

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile`

```
User Profile Hub

Usage:
  flexcli profile [command]

Available Commands:
  body        Body metrics and thresholds
  constraint  Manage physical constraints
  data        Sync & data: manual sync, activities, health metrics
  delete      Permanently delete user profile and all data
  get         View full profile
  goal        Manage training goals
  insights    View AI insights (coach, sleep or doc)
  preferences Manage preferences (expert settings, custom list)
  stats       View training statistics and reports

Flags:
  -h, --help   help for profile

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile [command] --help" for more information about a command.
```

### `flexcli profile body`

```
Body metrics and thresholds

Usage:
  flexcli profile body [command]

Available Commands:
  threshold   Manage training thresholds
  vitals      Manage body vitals (weight, height, sex, birthdate)

Flags:
  -h, --help   help for body

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile body [command] --help" for more information about a command.
```

### `flexcli profile body threshold`

```
Manage training thresholds

Usage:
  flexcli profile body threshold [command]

Available Commands:
  get         Get current thresholds
  set         Set thresholds

Flags:
  -h, --help   help for threshold

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile body threshold [command] --help" for more information about a command.
```

### `flexcli profile body threshold get`

```
Get current thresholds

Usage:
  flexcli profile body threshold get [flags]

Flags:
  -h, --help   help for get
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile body threshold set`

```
Set thresholds

Usage:
  flexcli profile body threshold set [flags]

Flags:
      --cycling-ftp int       Cycling FTP (W)
      --cycling-lthr int      Cycling LTHR (bpm)
      --cycling-pace string   Cycling Pace (e.g. 1:20/km)
  -h, --help                  help for set
      --json                  Output in JSON format
      --running-ftp int       Running FTP (W)
      --running-lthr int      Running LTHR (bpm)
      --running-pace string   Running Pace (e.g. 4:30/km)

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile body vitals`

```
Manage body vitals (weight, height, sex, birthdate)

Usage:
  flexcli profile body vitals [command]

Available Commands:
  get         View current body vitals
  set         Update body vitals

Flags:
  -h, --help   help for vitals

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile body vitals [command] --help" for more information about a command.
```

### `flexcli profile body vitals get`

```
View current body vitals

Usage:
  flexcli profile body vitals get [flags]

Flags:
  -h, --help   help for get
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile body vitals set`

```
Update body vitals

Usage:
  flexcli profile body vitals set [flags]

Flags:
      --birthdate string   Birthdate (YYYY-MM-DD)
      --height float       Height in cm
  -h, --help               help for set
      --json               Output in JSON format
      --sex string         Sex (male/female/other)
      --weight float       Weight in kg

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile constraint`

```
Manage physical constraints

Usage:
  flexcli profile constraint [command]

Available Commands:
  add         Add a new physical constraint
  delete      Delete a constraint by index
  list        List all constraints

Flags:
  -h, --help   help for constraint

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile constraint [command] --help" for more information about a command.
```

### `flexcli profile constraint add`

```
Add a new physical constraint

Usage:
  flexcli profile constraint add [text] [flags]

Flags:
  -h, --help   help for add
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile constraint delete`

```
Delete a constraint by index

Usage:
  flexcli profile constraint delete [index] [flags]

Flags:
  -h, --help   help for delete
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile constraint list`

```
List all constraints

Usage:
  flexcli profile constraint list [flags]

Flags:
  -h, --help   help for list
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data`

```
Sync & data: manual sync, activities, health metrics

Usage:
  flexcli profile data [command]

Available Commands:
  activity     Manage Garmin activities (alias: act): list, show, download, upload, delete
  fitness      View imported fitness data: personal records
  garmin       Garmin gamification, equipment, goals, workouts, and devices
  healthmetric View imported health metrics (alias: hm): list, show, delete
  nutrition    Manage nutrition logs (alias: nut): log, list, delete
  sync         Manually trigger Garmin or Withings synchronization

Flags:
  -h, --help   help for data

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile data [command] --help" for more information about a command.
```

### `flexcli profile data activity`

```
Manage Garmin activities (alias: act): list, show, download, upload, delete

Usage:
  flexcli profile data activity [command]

Aliases:
  activity, act

Available Commands:
  delete        Delete an activity from Garmin Connect (defaults to 'latest', currently disabled)
  download      Download an activity file from Garmin Connect (defaults to 'latest')
  download-bulk Download multiple activities in a ZIP bundle
  list          List synced activities with their Garmin activity IDs
  rename        Rename an activity in Garmin Connect (defaults to 'latest')
  show          Show details for a specific activity (defaults to 'latest')
  sync          Sync activity weather, splits, HR zones, and exercise sets now (normally daily)
  upload        Upload a FIT/GPX/TCX file to Garmin Connect

Flags:
  -h, --help   help for activity

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile data activity [command] --help" for more information about a command.
```

### `flexcli profile data activity delete`

```
Delete an activity from Garmin Connect (defaults to 'latest', currently disabled)

Usage:
  flexcli profile data activity delete [activity_id] [flags]

Flags:
  -h, --help   help for delete

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data activity download`

```
Download an activity file from Garmin Connect (defaults to 'latest')

Usage:
  flexcli profile data activity download [activity_id] [flags]

Flags:
  -f, --format string   File format (fit, gpx, tcx, csv, kml) (default "fit")
  -h, --help            help for download
  -o, --output string   Output file path (default: <activity_id>.<ext>)

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data activity download-bulk`

```
Download multiple activities in a ZIP bundle

Usage:
  flexcli profile data activity download-bulk [flags]

Flags:
  -d, --day int         Day to download (1-31, optional)
  -f, --format string   File format (fit, gpx, tcx, csv, kml) (default "fit")
  -h, --help            help for download-bulk
  -m, --month int       Month to download (1-12, optional)
  -o, --output string   Output ZIP file path
  -y, --year int        Year to download activities for (required)

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data activity list`

```
List synced activities with their Garmin activity IDs

Usage:
  flexcli profile data activity list [flags]

Flags:
      --end-date string     Only include activities on or before this date (YYYY-MM-DD)
  -h, --help                help for list
      --json                Output as JSON
      --page int            Page number (default 1)
      --page-size int       Number of activities per page (default 20)
      --start-date string   Only include activities on or after this date (YYYY-MM-DD)

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data activity rename`

```
Rename an activity in Garmin Connect (defaults to 'latest')

Usage:
  flexcli profile data activity rename <title> [activity_id] [flags]

Flags:
  -h, --help   help for rename

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data activity show`

```
Show details for a specific activity (defaults to 'latest')

Usage:
  flexcli profile data activity show [activity_id] [flags]

Flags:
  -h, --help   help for show
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data activity sync`

```
Sync activity weather, splits, HR zones, and exercise sets now (normally daily)

Usage:
  flexcli profile data activity sync [flags]

Flags:
  -h, --help   help for sync

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data activity upload`

```
Upload a FIT/GPX/TCX file to Garmin Connect

Usage:
  flexcli profile data activity upload <file> [flags]

Flags:
  -h, --help   help for upload

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data fitness`

```
View imported fitness data: personal records

Usage:
  flexcli profile data fitness [command]

Available Commands:
  records     List all personal records (PRs) from Garmin

Flags:
  -h, --help   help for fitness

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile data fitness [command] --help" for more information about a command.
```

### `flexcli profile data fitness records`

```
List all personal records (PRs) from Garmin

Usage:
  flexcli profile data fitness records [flags]

Flags:
  -h, --help   help for records
      --json   Output as JSON

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data garmin`

```
Garmin gamification, equipment, goals, workouts, and devices

Usage:
  flexcli profile data garmin [command]

Available Commands:
  badges      List all earned and available Garmin badges
  challenges  List all Garmin challenges
  devices     List all registered Garmin devices
  gear        List all Garmin gear (shoes, bikes, etc.)
  goals       List all Garmin training goals
  sync        Sync Garmin badges, challenges, gear, goals, workouts, and devices now (normally weekly)
  workouts    List all Garmin scheduled workouts

Flags:
  -h, --help   help for garmin

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile data garmin [command] --help" for more information about a command.
```

### `flexcli profile data garmin badges`

```
List all earned and available Garmin badges

Usage:
  flexcli profile data garmin badges [flags]

Flags:
  -h, --help   help for badges
      --json   Output as JSON

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data garmin challenges`

```
List all Garmin challenges

Usage:
  flexcli profile data garmin challenges [flags]

Flags:
  -h, --help   help for challenges
      --json   Output as JSON

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data garmin devices`

```
List all registered Garmin devices

Usage:
  flexcli profile data garmin devices [flags]

Flags:
  -h, --help   help for devices
      --json   Output as JSON

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data garmin gear`

```
List all Garmin gear (shoes, bikes, etc.)

Usage:
  flexcli profile data garmin gear [flags]

Flags:
  -h, --help   help for gear
      --json   Output as JSON

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data garmin goals`

```
List all Garmin training goals

Usage:
  flexcli profile data garmin goals [flags]

Flags:
  -h, --help   help for goals
      --json   Output as JSON

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data garmin sync`

```
Sync Garmin badges, challenges, gear, goals, workouts, and devices now (normally weekly)

Usage:
  flexcli profile data garmin sync [flags]

Flags:
  -h, --help   help for sync

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data garmin workouts`

```
List all Garmin scheduled workouts

Usage:
  flexcli profile data garmin workouts [flags]

Flags:
  -h, --help   help for workouts
      --json   Output as JSON

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data healthmetric`

```
View imported health metrics (alias: hm): list, show, delete

Usage:
  flexcli profile data healthmetric [command]

Aliases:
  healthmetric, hm

Available Commands:
  delete      Delete a health metric (currently disabled)
  list        List imported health metrics (paginated)
  show        Show aggregated health metric for a specific date (YYYY-MM-DD, defaults to today)

Flags:
  -h, --help   help for healthmetric

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile data healthmetric [command] --help" for more information about a command.
```

### `flexcli profile data healthmetric delete`

```
Delete a health metric (currently disabled)

Usage:
  flexcli profile data healthmetric delete <id> [flags]

Flags:
  -h, --help   help for delete

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data healthmetric list`

```
List imported health metrics (paginated)

Usage:
  flexcli profile data healthmetric list [flags]

Flags:
      --end-date string     Only include metrics on or before this date (YYYY-MM-DD)
  -h, --help                help for list
      --json                Output as JSON
      --page int            Page number (default 1)
      --page-size int       Number of metrics per page (default 20)
      --start-date string   Only include metrics on or after this date (YYYY-MM-DD)

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data healthmetric show`

```
Show aggregated health metric for a specific date (YYYY-MM-DD, defaults to today)

Usage:
  flexcli profile data healthmetric show [date] [flags]

Flags:
  -h, --help   help for show
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data nutrition`

```
Manage nutrition logs (alias: nut): log, list, delete

Usage:
  flexcli profile data nutrition [command]

Aliases:
  nutrition, nut

Available Commands:
  delete      Delete a nutrition log entry
  list        List nutrition entries and daily totals
  log         Log a food/nutrition entry

Flags:
  -h, --help   help for nutrition

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile data nutrition [command] --help" for more information about a command.
```

### `flexcli profile data nutrition delete`

```
Delete a nutrition log entry

Usage:
  flexcli profile data nutrition delete <id> [flags]

Flags:
  -h, --help   help for delete

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data nutrition list`

```
List nutrition entries and daily totals

Usage:
  flexcli profile data nutrition list [flags]

Flags:
  -d, --days int            Number of days to list up to today
      --end-date string     Only include entries on or before this date (YYYY-MM-DD)
  -h, --help                help for list
      --json                Output as JSON
      --start-date string   Only include entries on or after this date (YYYY-MM-DD)

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data nutrition log`

```
Log a food/nutrition entry

Usage:
  flexcli profile data nutrition log [flags]

Flags:
      --calories int    Calories (kcal)
      --carbs float     Carbohydrates (g)
      --fat float       Fat (g)
      --fiber float     Fiber (g)
  -h, --help            help for log
      --meal string     Meal type (breakfast, lunch, dinner, snack)
      --name string     Name of the food item or meal
      --notes string    Optional notes
      --protein float   Protein (g)
      --sodium int      Sodium (mg)
      --sugar float     Sugar (g)
      --time string     Time of entry (YYYY-MM-DD or YYYY-MM-DD HH:MM, default: now)

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data sync`

```
Manually trigger Garmin or Withings synchronization

Usage:
  flexcli profile data sync [command]

Available Commands:
  garmin      Sync data from Garmin Connect
  withings    Sync data from Withings

Flags:
  -h, --help   help for sync

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile data sync [command] --help" for more information about a command.
```

### `flexcli profile data sync garmin`

```
Sync data from Garmin Connect

Usage:
  flexcli profile data sync garmin [flags]

Flags:
  -h, --help   help for garmin

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile data sync withings`

```
Sync data from Withings

Usage:
  flexcli profile data sync withings [flags]

Flags:
  -h, --help   help for withings

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile delete`

```
Permanently delete user profile and all data

Usage:
  flexcli profile delete [flags]

Flags:
      --force   Skip confirmation prompt
  -h, --help    help for delete

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile get`

```
View full profile

Usage:
  flexcli profile get [flags]

Flags:
  -h, --help   help for get
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile goal`

```
Manage training goals

Usage:
  flexcli profile goal [command]

Available Commands:
  add         Add a new performance goal
  delete      Delete a goal
  list        List active and pending goals
  suggest     Suggest measurable targets for a qualitative goal using AI

Flags:
  -h, --help   help for goal

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile goal [command] --help" for more information about a command.
```

### `flexcli profile goal add`

```
Add a new performance goal

Usage:
  flexcli profile goal add [name] [flags]

Flags:
      --description string   Goal description
  -h, --help                 help for add
      --json                 Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile goal delete`

```
Delete a goal

Usage:
  flexcli profile goal delete [id] [flags]

Flags:
  -h, --help   help for delete
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile goal list`

```
List active and pending goals

Usage:
  flexcli profile goal list [flags]

Flags:
  -h, --help   help for list
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile goal suggest`

```
Suggest measurable targets for a qualitative goal using AI

Usage:
  flexcli profile goal suggest [goal description] [flags]

Flags:
  -h, --help   help for suggest
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile insights`

```
View AI insights (coach, sleep or doc)

Usage:
  flexcli profile insights [command]

Available Commands:
  coach       View latest AI coaching insights
  doc         View latest health analysis from 'The Doc'
  sleep       View latest sleep investigation insights

Flags:
  -h, --help   help for insights

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile insights [command] --help" for more information about a command.
```

### `flexcli profile insights coach`

```
View latest AI coaching insights

Usage:
  flexcli profile insights coach [flags]

Flags:
      --force   Force regeneration of insights
  -h, --help    help for coach
      --json    Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile insights doc`

```
View latest health analysis from 'The Doc'

Usage:
  flexcli profile insights doc [flags]

Aliases:
  doc, health

Flags:
      --force   Force regeneration of health analysis
  -h, --help    help for doc
      --json    Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile insights sleep`

```
View latest sleep investigation insights

Usage:
  flexcli profile insights sleep [flags]

Flags:
      --force   Regenerate report even if one exists for today
  -h, --help    help for sleep
      --json    Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile preferences`

```
Manage preferences (expert settings, custom list)

Usage:
  flexcli profile preferences [command]

Aliases:
  preferences, pref

Available Commands:
  custom      Manage custom training preferences (free-text list)
  expert      Manage expert settings (SLEEP_LOG_ENABLED, sync intervals, LANGUAGE, etc.)

Flags:
  -h, --help   help for preferences

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile preferences [command] --help" for more information about a command.
```

### `flexcli profile preferences custom`

```
Manage custom training preferences (free-text list)

Usage:
  flexcli profile preferences custom [command]

Available Commands:
  add         Add a new custom training preference
  list        List custom training preferences
  remove      Remove a custom training preference by index

Flags:
  -h, --help   help for custom

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile preferences custom [command] --help" for more information about a command.
```

### `flexcli profile preferences custom add`

```
Add a new custom training preference

Usage:
  flexcli profile preferences custom add [text] [flags]

Flags:
  -h, --help   help for add
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile preferences custom list`

```
List custom training preferences

Usage:
  flexcli profile preferences custom list [flags]

Flags:
  -h, --help   help for list
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile preferences custom remove`

```
Remove a custom training preference by index

Usage:
  flexcli profile preferences custom remove [index] [flags]

Flags:
  -h, --help   help for remove
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile preferences expert`

```
Manage expert settings (SLEEP_LOG_ENABLED, sync intervals, LANGUAGE, etc.)

Usage:
  flexcli profile preferences expert [command]

Available Commands:
  get         View current preferences
  set         Update user preferences using KEY=VALUE pairs

Flags:
  -h, --help   help for expert

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile preferences expert [command] --help" for more information about a command.
```

### `flexcli profile preferences expert get`

```
View current preferences

Usage:
  flexcli profile preferences expert get [flags]

Flags:
  -h, --help   help for get
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile preferences expert set`

```
Update user preferences. 
Expert settings and basic settings can also be set via KEY=VALUE positional arguments. Use KEY= to reset a setting to its system default.
Example: flexcli profile preferences expert set WITHINGS_SYNC_INTERVAL_HOURS=2 LANGUAGE=Deutsch timezone=Europe/Vienna
Example (reset): flexcli profile preferences expert set WITHINGS_SYNC_INTERVAL_HOURS=

Usage:
  flexcli profile preferences expert set [KEY=VALUE...] [flags]

Flags:
  -h, --help   help for set
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile stats`

```
View training statistics and reports

Usage:
  flexcli profile stats [command]

Available Commands:
  dashboard    View training dashboard
  healthtrends View health trends (7d vs 30d)
  report       View training reports
  sleep        Manage sleep logs

Flags:
  -h, --help   help for stats

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile stats [command] --help" for more information about a command.
```

### `flexcli profile stats dashboard`

```
View training dashboard

Usage:
  flexcli profile stats dashboard [flags]

Flags:
  -h, --help   help for dashboard
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile stats healthtrends`

```
View health trends (7d vs 30d)

Usage:
  flexcli profile stats healthtrends [flags]

Aliases:
  healthtrends, health

Flags:
  -d, --days int   Lookback days for trend analysis (default 30)
      --force      Force regeneration of health analysis
  -h, --help       help for healthtrends
      --json       Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile stats report`

```
View training reports

Usage:
  flexcli profile stats report [command]

Available Commands:
  list        List recent training reports
  show        Show detailed training report

Flags:
  -h, --help   help for report

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile stats report [command] --help" for more information about a command.
```

### `flexcli profile stats report list`

```
List recent training reports

Usage:
  flexcli profile stats report list [flags]

Flags:
  -h, --help   help for list
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile stats report show`

```
Show detailed training report

Usage:
  flexcli profile stats report show [report-id] [flags]

Flags:
  -h, --help   help for show
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile stats sleep`

```
Manage sleep logs

Usage:
  flexcli profile stats sleep [command]

Available Commands:
  get         Get a sleep log for a specific date
  list        List recent sleep logs
  log         Submit a daily sleep log

Flags:
  -h, --help   help for sleep

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)

Use "flexcli profile stats sleep [command] --help" for more information about a command.
```

### `flexcli profile stats sleep get`

```
Get a sleep log for a specific date

Usage:
  flexcli profile stats sleep get [date] [flags]

Flags:
  -h, --help   help for get
      --json   Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile stats sleep list`

```
List recent sleep logs

Usage:
  flexcli profile stats sleep list [flags]

Flags:
  -d, --days int   Number of days to list (default 7)
  -h, --help       help for list
      --json       Output in JSON format

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

### `flexcli profile stats sleep log`

```
Submit a daily sleep log

Usage:
  flexcli profile stats sleep log [flags]

Flags:
      --alcohol int       Alcohol units consumed
      --caffeine string   Last caffeine bucket (before_noon, before_2pm, before_5pm, after_5pm) (default "before_noon")
      --date string       Date (YYYY-MM-DD), defaults to today
  -h, --help              help for log
      --meal              Had a heavy meal after 7 PM
      --notes string      Optional notes
      --restedness int    Subjective restedness (1-5) (default 3)

Global Flags:
      --config string    config file (default is $HOME/.flexcli.json)
      --context string   Use specific context from config
      --key string       FlexCoach API key override (env: FLEXCLI_API_KEY)
      --server string    FlexCoach server URL override (env: FLEXCLI_SERVER_URL)
```

