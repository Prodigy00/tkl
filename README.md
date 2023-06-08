# tkl - short for task list

tkl is a command line productivity tool that enables devs to have all their `//TODO:` comments in one place instead of having it littered across projects. At any given time as a dev, you could be working on multiple (large) codebases and having `//TODO:` comments in your codebase that you could likely forget except you visit that particular file again is suboptimal, why not have one place that centralizes all your tasks on the fly? That is what **tkl** aims to solve. What's better? You never have to leave the terminal! 

## Features

- Add tasks defined by their level of urgency(red emojis for mission critical tasks, yellow for important-but-not-urgent tasks and green for normal tasks)
- List tasks
- Update tasks, marking them done/not done(incoming, wip)
- Delete tasks(incoming, wip)


## Dependencies

The project requires the following dependencies:

- github.com/enescakir/emoji

## Documentation of subcommands and basic usage
This will be updated as the feature set gets completed
```
subcommand: add
usage:
   subcommand name: add
   subcommand options:
      -t : task name(string),
      -l : level of importance (int)(0,1,2)
      -h : description for add subcommand
      -help: description for add subcommand
   default actions:
      adds the date of creation automatically

subcommand: delete
usage:
   subcommand name: delete
   subcommand options:
      None yet
   default actions:
      delete task(s)

subcommand: list
usage:
   subcommand name: list
   subcommand options:
      -h : description for list subcommand
      -help : description for list subcommand
   default actions:
      lists tasks

subcommand: update
usage:
   subcommand name: update
   subcommand options:
      None yet
   default actions:
      update a task

```

## Installation



```sh
git clone https://github.com/<your_username>/tkl.git
cd tkl
go build
```

## Sample
 <img width="1440" alt="tkl sample ss 2" src="https://github.com/Prodigy00/tkl/assets/19712590/5bbf52af-a4ef-43f8-8c26-258415063c4d">

## License

This project is licensed under the [MIT License](LICENSE).
