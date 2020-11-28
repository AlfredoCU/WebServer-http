package main

import (
	"./data"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// Port
	PORT = ":9000"
	// Routers
	HOME = "/"
	FORM = "/form"
	DATA = "/data"
	GENERAL = "/general"
	STUDENT = "/student"
	SUBJECT = "/subject"
	// Header
	KEY = "Content-Type"
	VALUE = "text/html"
)

var dataset data.AllData

func loadingHTML(a string) string {
	html, _ := ioutil.ReadFile(a)
	return string(html)
}

func root(res http.ResponseWriter, _ *http.Request) {
	res.Header().Set(KEY, VALUE)
	_, _ = fmt.Fprint(res, loadingHTML("index.html"))
}

func form(res http.ResponseWriter, _ *http.Request) {
	res.Header().Set(KEY, VALUE)
	_, _ = fmt.Fprintf(res, loadingHTML("form.html"))
}

func student(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)

	if err := req.ParseForm(); err != nil {
		_, _ = fmt.Fprintf(res, "ParseForm() error %v", err)
		return
	}

	fmt.Println(req.PostForm)
	stu := req.FormValue("name")

	fmt.Println(stu)

	res.Header().Set(KEY, VALUE)
	_, _ = fmt.Fprintf(res, loadingHTML("student.html"), dataset.StudentAVG(stu))
}

func subject(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)

	if err := req.ParseForm(); err != nil {
		_, _ = fmt.Fprintf(res, "ParseForm() error %v", err)
		return
	}

	fmt.Println(req.PostForm)
	sub := req.FormValue("subject")

	fmt.Println(sub)

	res.Header().Set(KEY, VALUE)
	_, _ = fmt.Fprintf(res, loadingHTML("subject.html"), dataset.SubjectAVG(sub))
}

func general(res http.ResponseWriter, _ *http.Request) {
	res.Header().Set(KEY, VALUE)
	_,_ = fmt.Fprintf(res, loadingHTML("general.html"), dataset.GeneralAVG())
}

func datasets(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)

	switch req.Method {
		case "POST":
			if err := req.ParseForm(); err != nil {
				_, _ = fmt.Fprintf(res, "ParseForm() error %v", err)
				return
			}

			fmt.Println(req.PostForm)

			dts := 	data.Data{
				Student: req.FormValue("student"),
				Subject: req.FormValue("subject"),
				Grade: req.FormValue("grade")}

			dataset.Add(dts)
			fmt.Println(dataset)

			res.Header().Set(KEY, VALUE)
			_, _ = fmt.Fprintf(res, loadingHTML("register.html"), dts.Student)

		case "GET":
			res.Header().Set(KEY, VALUE)
			_, _ = fmt.Fprintf(res, loadingHTML("data.html"), dataset.String())
		}
}

func main() {
	http.HandleFunc(HOME, root)
	http.HandleFunc(FORM, form)
	http.HandleFunc(DATA, datasets)
	http.HandleFunc(GENERAL, general)
	http.HandleFunc(STUDENT, student)
	http.HandleFunc(SUBJECT, subject)
	fmt.Println("App listening at http://localhost" + PORT)
	_ = http.ListenAndServe(PORT, nil)
}
