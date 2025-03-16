package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error){
	var(
		method= r.Method
		uri= r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "url", uri )
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

func (app *application) clientError(w http.ResponseWriter, status int){
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data TemplateData) {
        
	ts, ok:= app.templateCache[page]
	if !ok{
		err:= fmt.Errorf("The template %s does not exist!", page)
		app.serverError(w, r, err)
	}
    //buffer to handle runtime errors gracefully
	buf:= new(bytes.Buffer)

	err:= ts.ExecuteTemplate(buf, "base", data)
	if err!=nil{
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
	
}

func (app *application) newTemplateData(r *http.Request) TemplateData{
	return TemplateData{
		CurrentYear: time.Now().Year(),
	}
}