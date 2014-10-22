package svr

import (
	"io"
	"libs/log"
	"net/http"
)

func init() {
	println("svr.admin.init")
}

func HttpAdminHandler(w http.ResponseWriter, r *http.Request) {
	//　TODO：　做web控制后台
	log.Info("admin Handler: %s %s %s", r.Method, r.URL.RequestURI(), r.RemoteAddr)

	if !authOK(r) {
		noAuthResponse(w)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	io.WriteString(w, "TODO: 留做web控制后台")
}