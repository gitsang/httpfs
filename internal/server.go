package internal

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gitsang/httpfs/pkg/netx"
)

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

	slog.Info("serving...", logs...)
	http.Handle("/", http.FileServer(http.Dir(dir)))
	if err := http.ListenAndServe(listen, nil); err != nil {
		slog.Error("serve failed", append(logs, slog.Any("err", err))...)
	}
}
