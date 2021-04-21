package main

type HelpCommand struct{
}

func (cmd *HelpCommand) Run() int {
    app.ShowHelp(true)

    return ErrOk
}
