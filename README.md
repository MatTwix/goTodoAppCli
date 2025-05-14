# Todo App that works in terminal

## How to run the app

Firstly, you need to be in project's directory. Open the terminal and run next commands:

```cmd
go build -o todoapp
./todoapp [command]
```

## Command types

1. `./todoapp list` - gets all current tasks
1. `./todoapp add "Task name"` - creates new task
1. `./todoapp done 1` - makes task 1 done
1. `./todoapp remove 1` - deletes task 1

* Use -h, --help flags after any of the commands to get short instructions!
