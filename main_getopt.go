package main

import (
	"fmt";
	"os";
	"strings";
	"github.com/tcler/cmdline-go/getopt";
)

func main() {
	var nargv []string = os.Args[1:]
	var cmdline getopt.Cmdline

	var options = []getopt.Option {
		{Help: "Options group1:"},
		{Names: []string{"h", "H", "help"}, Argtype: getopt.N, Help: "out put the usage info"},
		{Names: []string{"f", "F", "file"}, Argtype: getopt.M, Help: "file to be parse"},
		{Names: []string{"wenj", "wenjian"}, Link: "f", Hide: true, Help: "deprecated use f instead"},
		{Names: []string{"o"}, Argtype: getopt.O, Help: "mount option"},
		{Names: []string{"v"}, Argtype: getopt.N, Help: "verbose output, -vvv means verbose level 3"},
		{Names: []string{"x"}, Argtype: getopt.Y, Help: "dump binary file to text"},
		{Names: []string{"s"}, Argtype: getopt.Y, Help: "enable smart mode", Hide: false},
		{Names: []string{"S"}, Link: "s", Hide: true},

		{Help: "\nOptions group2:"},
		{Names: []string{"e"}, Argtype: getopt.M, Help: "sed -e option, will forward to child sed process", Forward: true},
		{Names: []string{"r"}, Argtype: getopt.N, Help: "sed -r option, will forward to child sed process", Forward: true},
		{Names: []string{"n"}, Argtype: getopt.N, Help: "sed -n option, will forward to child sed process", Forward: true},
	}

	// debug info
	fmt.Println("Debug info:")
	fmt.Println(strings.Repeat("-", 80))
	fmt.Println(nargv)
	getopt.GetUsage(options)

	cmdline = getopt.GetOptions(options, nargv)
	for k, v := range cmdline.OptionMap {
		fmt.Printf("opt(%s) = %#v\n", k, v)
	}
	fmt.Println()

	for _, v := range cmdline.InvalidOptions {
		fmt.Println(v)
	}
	fmt.Println()

	fmt.Printf("params: %#v\n", cmdline.Args)
	fmt.Printf("forward: %#v\n", cmdline.ForwardOptions)
	fmt.Println(strings.Repeat("-", 80))

	// start your code
	if _, ok := cmdline.OptionMap["help"]; ok {
		fmt.Println("Usage: ...")
		//getopt.GetUsage(options)
	}
	if val, ok := cmdline.OptionMap["file"]; ok {
		filelist := val
		fmt.Printf("file list: %#v\n", filelist)
	}
	if val, ok := cmdline.OptionMap["v"]; ok {
		verboselevel := len(val)
		fmt.Println("verbose level:", verboselevel)
	}
	if val, ok := cmdline.OptionMap["s"]; ok {
		smartmode := val[0]
		fmt.Printf("smart mode: \"%s\"\n", smartmode)
	}
}