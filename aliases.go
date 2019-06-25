package main

import (
	"fmt"
	"strings"
)

func printAliases(cfg *Cfg, strictMode, fallbackMode bool) {
	commands := cfg.ListCommands()
	outputs := make([]string, len(commands))

	flags := make([]string, 0, 2)
	if strictMode {
		flags = append(flags, "--strict")
	}

	if fallbackMode {
		flags = append(flags, "--fallback")
	}

	for index, command := range commands {
		output := append(flags, command)
		outputs[index] = strings.Join(output, " ")
	}

	fmt.Println()
	for index, command := range commands {
		fmt.Printf("alias %s='donner run %s'\n", command, outputs[index])
	}

	aliasCommand := strings.Join(append([]string{"donner", "aliases"}, flags...), " ")

	fmt.Println("")
	fmt.Printf("# copy and paste the output into your terminal or run\n")
	fmt.Printf("#  eval $(%s)", aliasCommand)
	fmt.Println("")
}
