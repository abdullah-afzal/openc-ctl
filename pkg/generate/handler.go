package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/olegsu/openc-ctl/configs/templates"
	"github.com/spf13/viper"
)

type (
	// Handler - exposed struct that implementd Handler interface
	Handler struct{}
)

// Handle - the function that will be called from the CLI with viper config
// to provide access to all flags
func (g *Handler) Handle(cnf *viper.Viper) error {
	var file *os.File
	fileName := "main.go.tmpl"
	main := templates.TemplatesMap()[fileName]
	file, err := os.Create("main.go")
	if err != nil {
		return err
	}
	_, err = file.Write([]byte(main))
	if err != nil {
		return err
	}
	fmt.Println(cnf.GetString("user"))
	fmt.Println(cnf.GetString("name"))
	_, err = ioutil.ReadFile("go.mod")
	if os.IsNotExist(err) {
		fmt.Println("go.mod not exist, creating module")
		init := exec.Command("go", "mod", "init", fmt.Sprintf("github.com/%s/%s", cnf.GetString("user"), cnf.GetStringSlice("name")[0]))
		init.Stdout = os.Stdout
		init.Stderr = os.Stderr
		init.Env = os.Environ()
		err = init.Run()
		if err != nil {
			return err
		}

		tidy := exec.Command("go", "mod", "tidy")
		tidy.Stdout = os.Stdout
		tidy.Stderr = os.Stderr
		tidy.Env = os.Environ()
		err = tidy.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
