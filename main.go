package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	r "github.com/parnurzeal/gorequest"
	"github.com/x6r/haste/cmd"
)

type Data struct {
	Key string `json:"key,omitempty"`
}

const (
	ver   = "1.1.6"
	endpt = "documents"
)

var (
	instance = flag.String("i", "https://p.x4.pm", "hastebin instance where the content will be uploaded")
	filename = flag.String("f", "", "upload a file")
	raw      = flag.Bool("r", false, "returns raw url")
	prntver  = flag.Bool("v", false, "prints program version")
)

func init() {
	log.SetFlags(0)
}

func main() {
	flag.Parse()

	if *prntver {
		fmt.Println(ver)
		os.Exit(0)
	}

	var content, ext string
	if *filename == "" {
		content = flag.Arg(0)
	} else {
		file, err := os.ReadFile(*filename)
		if err != nil {
			log.Fatalln(err)
		}
		content = string(file)
		ext = filepath.Ext(*filename)
	}

	if content == "" {
		if pipe() {
			buf := bytes.NewBuffer(nil)
			_, err := buf.ReadFrom(os.Stdin)
			if err != nil {
				log.Fatalln(err)
			}
			content = buf.String()
		} else {
			content, *instance, *raw = cmd.Execute()
		}
	}

	res := upload(*instance, content)
	res += ext

	fmt.Println(res)
}

func upload(url string, s string) (res string) {
	req := r.New()
	resp, body, errs := req.Post(fmt.Sprintf("%s/%s", url, endpt)).Type("text").Send(s).End()

	if errs != nil {
		log.Fatalln(errs)
	}

	if resp.StatusCode != 200 {
		log.Fatalln(resp.Status)
	}

	var data Data
	err := json.Unmarshal([]byte(body), &data)
	if err != nil {
		log.Fatalln(err)
	}

	if *raw {
		res = fmt.Sprintf("%s/raw/%s", url, data.Key)
	} else {
		res = fmt.Sprintf("%s/%s", url, data.Key)
	}

	return
}

func pipe() bool {
	fi, _ := os.Stdin.Stat()
	return fi.Mode()&os.ModeCharDevice == 0
}
