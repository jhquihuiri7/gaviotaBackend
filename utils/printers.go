package utils

import (
	"fmt"
	"net/http"
	"os/exec"
)

func GetPrinters(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Cmd{
		Path: `C:\WINDOWS\system32\cmd.exe`,
	}
	cmd.Args = []string{"/C", "wmic", "printer", "get", "name,default"}
	salida, err := cmd.Output()
	fmt.Println(string(salida))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(w, string(salida))

}

// `salida, err := exec.Command("cmd", "/C", "wmic", "printer", "get", "name,default").Output()
//	if err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(string(salida))
//	salida, err = exec.Command("cmd", "/C", "wmic", "printer", "where", "name='EPSON L380 Series'", "call","setdefaultprinter").Output()
//	if err != nil {
//		log.Fatalln(err)
//	}
//	fmt.Println(string(salida))
//	p, err := filepath.Abs("./Libro1.pdf")
//	fmt.Println(p)
//	file := fmt.Sprintf(`"%s"`,p)
//	salida, err = exec.Command("cmd", "/C", "rundll32.exe", "mshtml.dll,PrintHTML", file).Output()
//	if err != nil {
//		log.Fatalln(err)
//	}`
