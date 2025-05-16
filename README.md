# ticktick CLI client

This is [ticktick](https://ticktick.com/) cli client, written by golang.

## Description

The [ticktick](https://ticktick.com/) is a ToDo service.
This program provides command line interface for [ticktick](https://ticktick.com/).

## Demo (with [fzf](https://github.com/junegunn/fzf))

TODO: prepare demo screen image

### List tasks

### Add task

### Close task

## Usage

```bash
$ ttick --help

NAME:
  ttick - ticktick CLI client

USAGE:
  titick [global options] command [command options] [args ...]

VERSION:
   0.0.1

COMMANDS:
     list, l                  Show all tasks
     show                     Show task detail
     completed-list, c-l, cl  Show all completed tasks (only premium users)
     add, a                   Add task
     modify, m                Modify task
     close, c                 Close task
     delete, d                Delete task
     labels                   Show all labels
     projects                 Show all projects
     sync, s                  Sync cache
     quick, q                 Quick add a task
     help, h                  Show a list of commands or help for one command

GLOBAL OPTIONS:
   --color              colorize output
   --csv                output in CSV format
   --debug              output logs
   --namespace          display parent task like namespace
   --indent             display children task with indent
   --project-namespace  display parent project like namespace
   --help, -h           show help
   --version, -v        print the version
```
