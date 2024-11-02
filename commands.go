package main

import "flag"

type CmdFlags struct {
	Add    string
	Delete int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add A New Todo Specify Title.")
	flag.StringVar(&cf.Edit, "edit", "", "Edit A Todo By Index &  Specify Title. id:new_title")

	flag.IntVar(&cf.Delete, "delete", -1, "Delete Todo Specify Index.")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle Todo Specify Index.")
	flag.BoolVar(&cf.List, "list", false, "List All Todo")
	flag.Parse()

	return &cf
}
