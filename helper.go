package main

import(
	"bytes"
	"html/template"
	"net/http"
	"runtime"
	"path/filepath"
	"log"
)

func concat(params ...string) string{
	var buffer bytes.Buffer
	for _,str := range params{
		buffer.WriteString(str);
	}

	return buffer.String();
}

func view(templateName string, data interface{},res http.ResponseWriter){
_,cwd,_,_ := runtime.Caller(0)

	full_path:= filepath.Join(filepath.Dir(cwd),concat("templates/",templateName,`.html`))

t,err:= template.ParseFiles(full_path)
if err != nil { // if there is an error
  	  log.Print("template parsing error: ", err) // log it
}
	err = t.Execute(res, data) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
  	  log.Print("template executing error: ", err) //log it
  	}
}
