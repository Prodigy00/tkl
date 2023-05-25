package tasklist

/**
 * subcommand name: add
 * subcommand options:
 *   -t : task name(string)
 *   -l : level of importance (int)(0,1,2)
 * default actions:
 * adds the date of creation automatically
 */

//type taskMap map[string]string
//
//func NewAddTaskCommand() *AddCommand {
//	atcmd := &AddCommand{
//		fs: flag.NewFlagSet("add", flag.ContinueOnError),
//	}
//	atcmd.fs.StringVar(&atcmd.taskname, "task", "new task", "name of new task to be added")
//	return atcmd
//}
//
//type AddCommand struct {
//	fs       *flag.FlagSet
//	taskname string
//}
//
//func (ac *AddCommand) Name() string {
//	return ac.fs.Name()
//}
//
//func (ac *AddCommand) Init(args []string) error {
//	return ac.fs.Parse(args)
//}
//
//func (ac *AddCommand) Run() error {
//	taskMap[ac.taskname]
//	fmt.Println("new task added: ", ac.taskname)
//	return nil
//}

// should exist as a subcommand e.g tkl add
// should be able to add task with -t argument as title
// i.e tkl add -m "taskdescription" -t "tag_which could be project name"
// could also be tkl add -tm "taskdescription":"tag"
