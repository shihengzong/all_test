package return_file_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func ReturnFile() {
	fmt.Println("程序开始...")
	// mux := http.NewServeMux()
	// mux.HandleFunc("/hello", hello)
	// listener, _ := net.Listen("tcp", ":8080")
	// // 创建一个http服务
	// httpServer := &http.Server{
	// 	ReadTimeout:  5000,
	// 	WriteTimeout: 5000,
	// 	Handler:      mux,
	// }

	// // 启动服务端
	// httpServer.Serve(listener)
	http.HandleFunc("/read", Read)
	http.ListenAndServe(":8080", nil)
}

func Read(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/vnd.ms-excel")
	w.Header().Add("Content-Disposition", "attachment")
	w.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	fileName := `../legal_admin/excels/3b02e1727034646atemplate.bd3ec28.xlsx`
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0600)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	cByte, err := ioutil.ReadAll(f)
	w.Write(cByte)
}
