package main

import (
	"chaochaogege.com/onlinecode/global"
	. "chaochaogege.com/onlinecode/languages"
)

func langFactory(lang string) global.Runnable {
	switch lang {
	case "golang":
		return &GolangRunner{}
	case "python":
		return &PythonRunner{}
	default:
		{
			return nil
		}
	}

}
