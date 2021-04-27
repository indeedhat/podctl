package main

import (
	"github.com/posener/complete/v2"
	"github.com/posener/complete/v2/predict"
)

// RegisterAutoComplete will generate a list of and apply autocomplete rules for podctl
func RegisterAutoComplete() *complete.Command {
	cmd := &complete.Command{
        Sub: map[string]*complete.Command{
            "apply": &complete.Command{},
            "attach": &complete.Command{},
            "configure": &complete.Command{
                Flags: map[string]complete.Predictor{
                    "print": predict.Nothing,
                    "p": predict.Nothing,
                },
            },
            "exec": &complete.Command{
                Args: predict.Nothing,
            },
            "help": &complete.Command{},
            "init": &complete.Command{},
            "list": &complete.Command{},
            "logs": &complete.Command{},
            "restart": &complete.Command{},
        },
        Flags: map[string]complete.Predictor{},
    }

    cmd.Complete("podctl")

    return cmd
}
