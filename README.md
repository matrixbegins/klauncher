# klauncher

Clone repo

## Build Docker Image
```
docker build -t counter:v1 .
```

## Setup
Configure env file
```
cp .env.example .env
```

Make your chagnes in `.env` file

## Build Project
```
go build
```

## Run Project

```
./klauncher

Laucnh a Kubernetes Job or CronJob based on certain configurations and env values.

Usage:
  klauncher [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  cron        Schedules a Kubernetes CronJob.
  help        Help about any command
  job         Schedules a Kubernetes Batch Job.

Flags:
  -h, --help     help for klauncher
  -t, --toggle   Help message for toggle

Use "klauncher [command] --help" for more information about a command.
```


**To Schedule an on-demand Batch Job run**
```
./klauncher job
```

**To Schedule a CronJob run**
```
./klauncher cron
```

