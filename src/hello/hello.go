package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/o1egl/govatar"
)

func encodeImage(img image.Image, ext string) (*bytes.Buffer, string, error) {
	buffer := new(bytes.Buffer)

	if ext == "png" {
		e := png.Encode(buffer, img)
		return buffer, "image/png", e
	}

	e := jpeg.Encode(buffer, img, nil)
	return buffer, "image/jpeg", e
}

func writeImage(w http.ResponseWriter, img image.Image, ext string) {
	buffer, contentType, _ := encodeImage(img, ext)

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	w.Write(buffer.Bytes())
}

func femaleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	result, err := govatar.Generate(govatar.FEMALE)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	writeImage(w, result, vars["ext"])
}

func maleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	result, err := govatar.Generate(govatar.MALE)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	writeImage(w, result, vars["ext"])
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("sugar-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "sugar-cookie",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}

	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)

	fmt.Fprintf(w, ` <h1>GOvatar Generator</h1>

		<form action="/girlpower">
			<input type="submit" value="Female Avatar" />
		</form>
		<form action="/male">
			<input type="submit" value="Male Avatar" />
		</form>

		`)

	fmt.Fprintf(w, "<p>Times visited according to cookie counter (who doesn't like double cookies??): %s </p>", cookie.Value)

	result := division(count)
	strResult := strconv.Itoa(result)
	fmt.Fprintf(w, "<p>Times visited accoring to how many times you've hit refresh (very scientific calculation): %s </p>", strResult)
}

func division(x int) (result int) {
	result = x/2 + 1
	return result
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/girlpower", femaleHandler)
	http.HandleFunc("/male", maleHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
