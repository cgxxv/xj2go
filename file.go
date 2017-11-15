package xj2go

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func checkFile(filename, pkg string) (string, error) {
	if ok, err := pathExists(pkg); !ok {
		os.Mkdir(pkg, 0755)
		if err != nil {
			return "", err
		}
	}

	filename = path.Base(filename)
	if filename[:1] == "." {
		return "", errors.New("File could not start with '.'")
	}

	filename = pkg + "/" + filename + ".go"
	if ok, _ := pathExists(filename); ok {
		if err := os.Remove(filename); err != nil {
			log.Fatal(err)
			return "", err
		}
	}

	return filename, nil
}

func writeStruct(filename, pkg string, strcts *[]strctMap) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	pkgLines := make(map[string]string)
	strctLines := []string{}
	for _, strct := range *strcts {
		for root, sns := range strct {
			strctLines = append(strctLines, "type "+strings.Title(root)+" struct {\n")
			for i := 0; i < len(sns); i++ {
				if sns[i].Type == "time.Time" {
					pkgLines["time.Time"] = "import \"time\"\n"
				}
				strctLines = append(strctLines, "\t"+strings.Title(sns[i].Name)+"\t"+sns[i].Type+"\t"+sns[i].Tag+"\n")
			}
			strctLines = append(strctLines, "}\n")
		}
	}
	strctLines = append(strctLines, "\n")

	file.WriteString("package " + pkg + "\n\n")
	for _, pl := range pkgLines {
		file.WriteString(pl)
	}
	for _, sl := range strctLines {
		file.WriteString(sl)
	}

	ft := exec.Command("go", "fmt", filename)
	if err := ft.Run(); err != nil {
		log.Fatal(err)
		return err
	}

	vt := exec.Command("go", "vet", filename)
	if err := vt.Run(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
