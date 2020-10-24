package main

/*********************************************************************
**      IMPORTS
 */
import (
	"flag"
	"fmt"
	"os"
)

// Define some variables to receive command-line values
var scan_sleep int = 1
var start_safe bool = false
var max_sort int = 100000
var prog_name string
var g_debug int
var base_dir string
var msg_queue string
var file_suffix string
var file_prefix string

// Array to store the usage order
var UsageOrder []string

/**PROC+**********************************************************************/
/*                                                                           */
/* Name:        command-line-usage                                           */
/* Purpose:     demo program to display custom help output in a              */
/*                  generic way.  In practice CustomUsage and UsageOrder are */
/*                  stored in a lib package that is used by all programs     */
/* Returns:     None                                                         */
/* Params:      None                                                         */
/*                                                                           */
/**PROC-**********************************************************************/
func main() {
	// Comment the following line to see before and after effect with -h
	flag.Usage = CustomUsage
	// Setup flags and parse command line parameters
	setParameters()
	flag.Parse()

	// Print the variable values
	fmt.Println("Variables:")
	fmt.Println("scan_sleep:", scan_sleep)
	fmt.Println("start_safe:", start_safe)
	fmt.Println("max_sort:", max_sort)
	fmt.Println("prog_name:", prog_name)
	fmt.Println("g_debug:", g_debug)
	fmt.Println("base_dir:", base_dir)
	fmt.Println("msg_queue:", msg_queue)
	fmt.Println("file_suffix:", file_suffix)
	fmt.Println("file_prefix:", file_prefix)
}

/**PROC+**********************************************************************/
/*                                                                           */
/* Name:        CustomUsage                                                  */
/* Purpose:     Sets up custom usage display                                 */
/* Returns:     None                                                         */
/* Params:      None                                                         */
/*                                                                           */
/**PROC-**********************************************************************/
func CustomUsage() {
	/********************************************
	 * Need to put a safety here in case UsageOrder
	 * is not set.
	 * Because this is a lib function, we do not want
	 * to fail.  So, we will just print in lexical order
	 * like default if UsageOrder is not provided.
	 */
	if len(UsageOrder) == 0 {
		fmt.Fprintf(os.Stderr, "Error: Attempting to use CustomUsage, but UsageOrder is not set\n")
		fmt.Fprintf(os.Stderr, "\t -This output might be really ugly\n")
		flag.VisitAll(func(f *flag.Flag) {
			// append f.Name to UsageOrder
			UsageOrder = append(UsageOrder, f.Name)
		})
	}
	/********************************************
	 * Create usage map to store usage for each parameter
	 * we will use this in combination with UsageOrder
	 * to order the output
	 */
	usageMap := make(map[string]string, 0)

	// Loop through all defined flags, set or unset
	flag.VisitAll(func(f *flag.Flag) {
		// Get the flag type and store in tp
		tp, _ := flag.UnquoteUsage(f)
		if len(f.Name) > 1 && f.DefValue != "" {
			// Longname usage with default value
			usageMap[f.Name] = fmt.Sprintf("\n  -%s  %s\t%s (%s)", f.Name, tp, f.Usage, f.DefValue)
		} else if len(f.Name) > 1 {
			// Longname usage without default value
			usageMap[f.Name] = fmt.Sprintf("\n  -%s  %s\t%s", f.Name, tp, f.Usage)
		} else {
			// Shorthand usage
			usageMap[f.Name] = fmt.Sprintf("\t  -%s  \t%s", f.Name, f.Usage)
		}
	})
	fmt.Printf("Usage of %s:\n", os.Args[0])
	for s := range UsageOrder {
		fmt.Println(usageMap[UsageOrder[s]])
	}
}

/**PROC+**********************************************************************/
/*                                                                           */
/* Name:        setParameters                                                */
/* Purpose:     Set up flags for command line parse                          */
/* Returns:     None                                                         */
/* Params:      None                                                         */
/*                                                                           */
/**PROC-**********************************************************************/
func setParameters() {
	// Set usage order for display
	UsageOrder = []string{"debug", "D", "scan_sleep", "w", "start_safe", "S", "max_sort", "m", "base_dir", "d", "msg_queue", "q", "prog_name", "P", "file_suffix", "s", "file_prefix", "p"}

	/* ----  debug           ----*/
	flag.IntVar(&g_debug, "debug", g_debug, "debug level")
	flag.IntVar(&g_debug, "D", g_debug, "debug shorthand")
	/* ----  sleep time      ----*/
	flag.IntVar(&scan_sleep, "scan_sleep", scan_sleep, "Sleep between scans")
	flag.IntVar(&scan_sleep, "w", scan_sleep, "sleep shorthand")

	/* ----  Safe mode           ----*/
	flag.BoolVar(&start_safe, "start_safe", start_safe, "Specifies whether to requeue at startup")
	flag.BoolVar(&start_safe, "S", start_safe, "start_safe shorthand")

	/* ----  sort size      ----*/
	flag.IntVar(&max_sort, "max_sort", max_sort, "Specifies whether to requeue at startup")
	flag.IntVar(&max_sort, "m", max_sort, "max_sort shorthand")

	/* ----  directory      ----*/
	flag.StringVar(&base_dir, "base_dir", base_dir, "[Required] Base direcory for scan. This should contain depot and queued at minimum.")
	flag.StringVar(&base_dir, "d", base_dir, "base_dir shorthand.")

	/* ----  MSG Queue      ----*/
	flag.StringVar(&msg_queue, "msg_queue", msg_queue, "[Required] Message queue to communicate with procs")
	flag.StringVar(&msg_queue, "q", msg_queue, "msg_queue shorthand")

	/* ----  prog name      ----*/
	flag.StringVar(&prog_name, "prog_name", os.Args[0], "New program name: -prog_name=<new_name>")
	flag.StringVar(&prog_name, "P", os.Args[0], "prog_name shorthand")

	/* ----  file suffix   ----*/
	flag.StringVar(&file_suffix, "file_suffix", file_suffix, "Filename suffix to include.")
	flag.StringVar(&file_suffix, "s", file_suffix, "file_suffix shorthand.")

	/* ----  file prefix   ----*/
	flag.StringVar(&file_prefix, "file_prefix", file_prefix, "Filename prefix to include.")
	flag.StringVar(&file_prefix, "p", file_prefix, "file_prefix shorthand.")

}
