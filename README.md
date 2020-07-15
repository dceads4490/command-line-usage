### This is a golang demo program to setup custom help usage in a generic way.  I store the CustomUsage function in a library that is used by most all of my golang programs. I like to have long and short options for all of my command line parameters.  The problem with the golang default usage is that it is sorted lexically which makes it much harder to read when you have a lot of parameters.  This CustomUsage func allows me to control the order of the flag display and group my long and short options together.

#### This is a simple standalone example. The CustomUsage function controls the display.  The getParameters function contains the program flag setup and the UsageOrder array controls the display order.

### Build
#### go build command-line-usage.go

### Run
#### command-line-usage -h
Usage of command-line-usage:

  -debug  int	debug level (0)
	  -D  	debug shorthand

  -scan_sleep  int	Sleep between scans (1)
	  -w  	sleep shorthand

  -start_safe  	Specifies whether to requeue at startup (false)
	  -S  	start_safe shorthand

  -max_sort  int	Specifies whether to requeue at startup (100000)
	  -m  	max_sort shorthand

  -base_dir  string	[Required] Base direcory for scan. This should contain depot and queued at minimum.
	  -d  	base_dir shorthand.

  -msg_queue  string	[Required] Message queue to communicate with procs
	  -q  	msg_queue shorthand

  -prog_name  string	New program name: -prog_name=<new_name> (command-line-usage)
	  -P  	prog_name shorthand

  -file_suffix  string	Filename suffix to include.
	  -s  	file_suffix shorthand.

  -file_prefix  string	Filename prefix to include.
	  -p  	file_prefix shorthand.

