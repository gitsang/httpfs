package internal

import (
	_ "embed"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gitsang/httpfs/pkg/netx"
)

//go:embed upload.html
var uploadHtml string

func Serve(listen, dir string) {
	logs := make([]any, 0)
	logs = append(logs, slog.String("dir", dir), slog.String("listen", listen))

	ips, err := netx.GetIPv4s()
	if err != nil {
		slog.Error("get ipv4 failed", append(logs, slog.Any("err", err))...)
	}

	logG := make([]any, 0, len(ips))
	for _, ip := range ips {
		logG = append(logG, slog.String("addr", fmt.Sprintf("http://%s%s", ip, listen)))
	}
	logs = append(logs, slog.Group("visits", logG...))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			file, header, err := r.FormFile("file")
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer file.Close()

			dst, err := os.Create(filepath.Join(dir, r.URL.Path, header.Filename))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer dst.Close()

			_, err = io.Copy(dst, file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if strings.HasSuffix(r.URL.Path, "/") {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, uploadHtml)
		}

		http.FileServer(http.Dir(dir)).ServeHTTP(w, r)
	})

	slog.Info("serving...", logs...)
	if err := http.ListenAndServe(listen, nil); err != nil {
		slog.Error("serve failed", append(logs, slog.Any("err", err))...)
	}
}
