package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
    // pod defaults
    DefaultNamespace = "default"
    DefaultRestartCommand = "use default"

    // env defaults
    DefaultEditor = "autodetect"
    DefaultTerminalEmulator = "autodetect"

    // log defaults
    DefaultLogPrefix = "index"
)

// InitCommand will ask the user a bunch of questions to setup their podctl env
type InitCommand struct {
}

// Run is the entry point for the InitCommand
func (cmd *InitCommand) Run() int {
    if _, err := loadConfig(); err == nil {
        fmt.Println("A .podctl.toml file has been found in this directory\nNothing to do!")
        return ErrOk
    }

    var buff bytes.Buffer

    cmd.printHeader()
    cmd.askPodQuestions(&buff)
    cmd.askEnvQuestions(&buff)
    cmd.askLogQuestions(&buff)

    if err := ioutil.WriteFile(ConfigFileName, buff.Bytes(), 0644); err != nil {
        panic(err)
    }

    fmt.Print("\nThanks, you can now start using podctl\n\n")
    return ErrOk
}

// printHeader will display he opening speel for the init command
func (cmd *InitCommand) printHeader() {
    // this looks odd but its fine
    fmt.Println(
`    ____            __     __  __   ____      _ __ 
   / __ \____  ____/ /____/ /_/ /  /  _/___  (_) /_
  / /_/ / __ \/ __  / ___/ __/ /   / // __ \/ / __/
 / ____/ /_/ / /_/ / /__/ /_/ /  _/ // / / / / /_  
/_/    \____/\__,_/\___/\__/_/  /___/_/ /_/_/\__/  
                                                   `,
    )

    fmt.Println(`We will now help you setup your config.
Defaut values will be shown in parenteses and can be selected by hitting the Enter key.
 `);
}

// askPodQuestions asks the user questions about their pod
func (cmd *InitCommand) askPodQuestions(buff *bytes.Buffer) {
    buff.WriteString("[pod]\n")

    name := cmd.ask("What is your pods name?", "")
    buff.WriteString(fmt.Sprintf("name = \"%s\"\n", name))

    namespace := cmd.ask("What is your pods namespace?", DefaultNamespace)
    buff.WriteString(fmt.Sprintf("namespace = \"%s\"\n", namespace))

    restart := cmd.ask("Specify your pods restart command: ", DefaultRestartCommand)
    if restart == DefaultRestartCommand {
        buff.WriteString("# restart_cmd = \"\"\n")
    } else {
        buff.WriteString(fmt.Sprintf("restart_cmd = \"%s\"\n", strings.ReplaceAll(restart, `"`, `\"`)))
    }

}

// askEnvQuestions asks the user questions about their environment
func (cmd *InitCommand) askEnvQuestions(buff *bytes.Buffer) {
    buff.WriteString("\n[env]\n")

    configDir := cmd.ask("Where would you like your kubernetes yaml files stored?", DefaultConfigPath)
    if configDir == DefaultConfigPath {
        buff.WriteString("# config_dir = \"\"\n")
    } else {
        buff.WriteString(fmt.Sprintf("config_dir = \"%s\"\n", strings.ReplaceAll(configDir, `"`, `\"`)))
    }

    editor := cmd.ask("What is your prefered editor?", DefaultEditor)
    if editor == DefaultEditor {
        buff.WriteString("# editor = \"\"\n")
    } else {
        buff.WriteString(fmt.Sprintf("editor = \"%s\"\n", strings.ReplaceAll(editor, `"`, `\"`)))
    }

    terminal := cmd.ask("What is your prefered terminal emulator?", DefaultTerminalEmulator)
    if terminal == DefaultTerminalEmulator {
        buff.WriteString("# terminal_emulator = \"\"\n")
    } else {
        buff.WriteString(fmt.Sprintf("terminal_emulator = \"%s\"\n", strings.ReplaceAll(terminal, `"`, `\"`)))
    }
}

// askLogQuestions will ask the user questions about the logs
func (cmd *InitCommand) askLogQuestions(buff *bytes.Buffer) {
    buff.WriteString("\n[logs]\n")

    prefixOptions := []string {
        "index",
        "podId",
        "pod",
        "server",
        "server-pod",
        "none",
    }

    prefix := cmd.ask("What is your prefered log prefix?", DefaultLogPrefix, prefixOptions...)
    buff.WriteString(fmt.Sprintf("prefix = \"%s\"\n", prefix))
}

// ask will ask the user a question and return the answer
func (cmd *InitCommand) ask(question, defaultValue string, options ...string) string {
    if defaultValue != "" {
        fmt.Printf("\n%s (%s) ", question, defaultValue)
    } else {
        fmt.Print("\n", question, " ")
    }

    if len(options) > 0 {
        fmt.Printf("\n%v ", options)
    }

    var response string
    if _, err := fmt.Scanln(&response); err != nil {
        if err.Error() != "unexpected newline" || defaultValue == "" {
            fmt.Println("\nError: ", err)
            return cmd.ask(question, defaultValue, options...)
        }
    }

    value := response
    if response == "" {
        value = defaultValue
    }

    if value == "" {
        fmt.Println("\nError: You must provide an answer to this question")
        return cmd.ask(question, defaultValue)
    }

    if !cmd.valid(value, options) {
        fmt.Println("\nError: Invalid option")
        return cmd.ask(question, defaultValue)
    }

    return value
}

// valid will check if the given value is a valid option
func (cmd *InitCommand) valid(value string, options []string) bool {
    if len(options) == 0 {
        return true
    }

    for _, option := range options {
        if value == option {
            return true
        }
    }

    return false
}
