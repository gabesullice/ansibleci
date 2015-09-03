package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	http.HandleFunc("/", http.NotFound)
	http.HandleFunc("/deploy", DeployFunc)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func DeployFunc(w http.ResponseWriter, r *http.Request) {
	if err := os.Chdir("/app/deploykit-master"); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		r.Body.Close()
		log.Fatal(err)
	}

	decoder := json.NewDecoder(r.Body)
	var b map[string]interface{}
	if err := decoder.Decode(&b); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		r.Body.Close()
		return
	}

	log.Println(r.Body)

	go func() {
		os.Setenv("ANSIBLE_SSH_ARGS", "-o \"StrictHostKeyChecking no\" -o \"ForwardAgent yes\"")
		cmd := exec.Command("make", "deploy.dev")
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Start(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			r.Body.Close()
			log.Println(err)
		}

		if err := cmd.Wait(); err != nil {
			log.Println(err)
		}
	}()

	w.WriteHeader(http.StatusOK)
	r.Body.Close()
}
