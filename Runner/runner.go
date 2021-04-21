package Runner

import (
	"333exec/Runner/Code"
	"333exec/Runner/Data"
	"333exec/Runner/Packages"
	"333exec/Runner/Types"
	"333exec/Runner/Variable"
	"333exec/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

type Runner struct {
}

func (r *Runner) Handler(c *fiber.Ctx) error {
	urlPath := c.Path()
	switch *utils.TrimSlashes(&urlPath) {
	case "get/exe":
		imports := []*string{
			&Packages.LOG,
		}

		printHello, err := Code.LogPrintln(Variable.Var{
			Value: "hello " + c.IP(),
			Type:  Types.STRING,
		})
		if err != nil {
			return err
		}

		mainCode := []*string{
			&printHello,
		}

		t, err := template.ParseFiles("Templates/serverexe/exe.ccc")
		if err != nil {
			log.Print(err)
			return err
		}

		userPath := strings.Replace(c.IP(), ".", "_", -1)

		err = os.MkdirAll("files/"+userPath, os.ModeDir)
		if err != nil {
			return err
		}

		f, err := os.Create("files/" + userPath + "/main.go")
		if err != nil {
			log.Println("create file: ", err)
			return err
		}

		err = t.Execute(f, Data.Data{
			Imports:  imports,
			Vars:     nil,
			MainCode: mainCode,
			Funcs:    nil,
		})
		if err != nil {
			log.Print("execute: ", err)
			return err
		}

		myPath, err := os.Getwd()
		if err != nil {
			log.Println(err)
			return err
		}

		cmd := exec.Command("go", "build")
		cmd.Dir = myPath + "/files/" + userPath

		buildOutput, err := cmd.CombinedOutput()
		if err != nil {
			log.Println(err)
			fmt.Println(string(buildOutput))
			return err
		}

		return c.SendFile(myPath+"/files/"+userPath+"/"+userPath+".exe", true)
	}

	return nil
}
