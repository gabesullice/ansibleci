package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	//"os/exec"
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
		log.Fatal(err)
	}

	decoder := json.NewDecoder(r.Body)
	var b map[string]interface{}
	if err := decoder.Decode(&b); err != nil {
		log.Println(err)
		return
	}

	log.Println(b)

	//os.Setenv("ANSIBLE_SSH_ARGS", "-o \"StrictHostKeyChecking no\" -o \"ForwardAgent yes\"")
	//cmd := exec.Command("make", "deploy.dev")
	//cmd.Stderr = os.Stderr
	//cmd.Stdout = os.Stdout
	//if err := cmd.Run(); err != nil {
	//	log.Println(err)
	//}
}
