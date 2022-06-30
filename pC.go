package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const Creator = "https://t.me/www_1c_bitrix_dev"

var Port int
var URLReq string

func main() {

	cli()
	hello()
	go IamRun()

	r := chi.NewRouter()

	r.Get("/*", MainHandlFuncGet)
	r.Post("/*", MainHandlFuncPost)
	http.ListenAndServe(":"+strconv.Itoa(Port), r)

}

func MainHandlFuncGet(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get(URLReq + r.URL.RequestURI())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer resp.Body.Close()
	bodyRu, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(bodyRu)
	return

}

func MainHandlFuncPost(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	rb := bytes.NewReader(b)
	resp, err := http.Post(URLReq+r.URL.Path, r.Header["Content-Type"][0], rb)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	bodyRu, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(bodyRu)
	return

}

func IamRun() {
	for {
		time.Sleep(time.Second * 60)
		fmt.Println(`Я работаю, не выключай меня`)
	}
}

func hello() {
	fmt.Printf("Привет!\nЯ программа проксятор.\nЯ уменю принимать запросы на %v порт Post и Get и проксировать их нужный сервис.\n", Port)
	fmt.Printf("Если тебе нужны будут другие программы - пиши создателю.\n+++ %v +++ \n\n\n", Creator)

	fmt.Printf("Использую порт - %v\n", Port)
	fmt.Printf("Проксируем сервер - %v\n", URLReq)
}

func cli() {

	var urlReqCLI = flag.String("urlProxy", "", "The 'urlProxy' option server request")
	var portReqCLI = flag.Int("port", 8081, "The 'port' option port server request")
	flag.Parse()

	URLReq = *urlReqCLI
	Port = *portReqCLI

	if Port < 1 && Port > 65536 {
		panic("Не правильный порт \n")
	}

	if len(URLReq) < 10 {
		panic("Слишком короткий сервер\n")
	}

	if URLReq[:7] != "http://" && URLReq[:8] != "https://" {
		panic("Нет указания типа сервера \n")
	}
}
