package core

import (
	"fmt"
	"os"
)

var DIRGEN_VERSION = "3"

func GenerateIn() {
	if len(os.Args) != 2 {
		fmt.Println("Too much or too low args , need one.")
		fmt.Println("\tHint. Use ./binary (pwd) \n\tTo init go dir where you right now ")
		panic(os.Args)
	}
	path := os.Args[1]
	var err error
	fmt.Println("Dirgen V" + DIRGEN_VERSION + "\nFor other info :")
	fmt.Println("https://github.com/golang-standards/project-layout")
	dirs := []string{
		"/cmd",
		"/internal",
		"/pkg",
		"/deployments",
		"/test",
		"/assets",
		"/scripts",
		"/api",
	}
	for _, dirname := range dirs {
		err = os.Mkdir(path+dirname, 0755)
		if err != nil {
			fmt.Fprint(os.Stderr, "\nosmagic/dirgen cannot create:\t"+dirname+"\nat"+path+"\n\nerror is:\t"+err.Error())
		}
	}
	dgn, err := os.Create(path + "/DirgenNote.md")
	if err != nil {
		fmt.Fprint(os.Stderr, "\nosmagic/dirgen cannot create:\tDirgenNote.md\nat:"+path+"\n\nerror is:\t"+err.Error())

	}
	notes := []string{"./cmd is for main",
		"./internal is for project-specific sources",
		"\tinside internal should be directories that represents layers of service ",
		"./pkg is for sources that may be re-used in other projects",
		"./deployments is for docker k8s etc",
		"./test is for test, isnt it obvious",
		"./assets is for assets ",
		"./scripts is for makefiles , bashscripts etc ",
		"./api is for .proto ",
	}
	for _, note := range notes {
		_, err := dgn.WriteString(note)
		if err != nil {
			fmt.Fprint(os.Stderr, "\nosmagic/dirgen cannot write to:\tDirgenNote.md\nat:"+path+"\n\nerror is:\t"+err.Error())
		}
		_, err = dgn.Write([]byte{'\n'})
		if err != nil {
			fmt.Fprint(os.Stderr, "\nosmagic/dirgen cannot write to:\tDirgenNote.md\nat:"+path+"\n\nerror is:\t"+err.Error())
		}
		fmt.Println(note)
	}
	_, err = dgn.WriteString("\nFor other info :\nhttps://github.com/golang-standards/project-layout")
	if err != nil {
		fmt.Fprint(os.Stderr, "\nosmagic/dirgen cannot write to:\tDirgenNote.md\nat:"+path+"\n\nerror is:\t"+err.Error())
	}
	err = dgn.Close()
	if err != nil {
		fmt.Fprint(os.Stderr, "osmagic/dirgen cannot close connect to file")
	}

}
