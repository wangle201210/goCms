package main

import (
	"fmt"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	afs "github.com/alibabacloud-go/afs-20180112/client"
	"github.com/galaxy-book/captcha-golang-demo/sdk"
	"log"
	"net/http"
)

const(
	captchaId = "12482cad39f140d6b43cebf91af4334c"
	secretId = "6e3672c4a93260dfeaf0365868e3bbd5"
	secretKey = "31c752b79963a119dbeb8f6345d1aef1"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("web")))
	http.HandleFunc("/verify", verify)

	fmt.Println("server is success listen on 8080.")
	fmt.Println("web entry: http://127.0.0.1:8080/index.html")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func verify(w http.ResponseWriter, r *http.Request) {
	validate := r.FormValue("validate")
	verifier, err := sdk.New(captchaId, secretId, secretKey)
	user := "nico"
		w.Write([]byte(err.Error()))
		return
	}
	verifyResult, err := verifier.Verify(validate, user)
	if err != nil{
		w.Write([]byte(err.Error()))
		return
	}
	if verifyResult.Result{
		w.Write([]byte("验证成功！"))
	}else{
		w.Write([]byte("验证失败！"))
	}
}

func aliyun()  {

	/**
	  示例
	  import (

	  )

	{csessionid: "01meJRDQhjKvuhvkU5cu47-BmOueMnl5OTwaLUzuAy4fBdMSWV…RjOD1nU1uKpcoNxKsUMe2CR-EMOVCcgR9rW9P1ZVJqv_U-Gzw",
	value: "pass",
	sig: "05XqrtZ0EaFgmmqIQes-s-CAdD4GEr3YmqBzF8eqnj9HMCfjom…eChX3uINoSf6ii2wWoFU_usbLLACzjGJJLew_tRwYbW7hqg_Y",
	token: "FFFF0N00000000009DFF:1618314285310:0.28581050440491595"}
	*/
	config := new(rpc.Config)
	config.SetAccessKeyId("*** Provide your AccessKeyId ***").
		SetAccessKeySecret("*** Provide your AccessKeySecret ***").
		SetRegionId("cn-hangzhou").
		SetEndpoint("afs.aliyuncs.com")
	client, _ := afs.NewClient(config)
}
