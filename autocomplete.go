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
                    "apply": predict.Nothing,
                    "a": predict.Nothing,
                },
            },
            "exec": &complete.Command{
                Args: predict.Nothing,
            },
            "init": &complete.Command{},
            "list": &complete.Command{},
            "logs": &complete.Command{},
            "restart": &complete.Command{},
        },
        Flags: map[string]complete.Predictor{
            "help": predict.Nothing,
            "h": predict.Nothing,
        },
    }

    cmd.Complete("podctl")

    return cmd
}
