# Todo App that works in terminal

## Initiallisation

1. Clone the repo:

    ```bash
    git clone https://github.com/MatTwix GoTodoAppCli.git
    cd GoTodoAppCli
    ```

1. Bluild the application:

    ```bash
    go build -o todoapp
    ./todoapp [command]
    ```

1. Run the application:

    ```bash
    ./todoapp --help
    ```

## Command types*

1. `./todoapp list` - gets all current tasks
1. `./todoapp add "Task name"` - creates new task
1. `./todoapp done 1` - makes task 1 done
1. `./todoapp remove 1` - deletes task 1

\*Use -h, --help flags after any of the commands to get short instructions!
