package video

import (
	"github.com/go-chi/chi"
	"gitlab.com/kayslay/vid_trimmer/internal/input"
	"gitlab.com/kayslay/vid_trimmer/pkg/video/handler"
	"gitlab.com/kayslay/vid_trimmer/pkg/video/service"
	"net/http"
)

func Router() http.Handler {
	r := chi.NewRouter()

	svc := service.NewBasicService(input.NewLink("file"))
	h := handler.NewVideo(svc)
	r.Get("/", h.Download)
	return r
}
