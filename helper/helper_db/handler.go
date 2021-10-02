package helper_db

type Handler struct {
	service Servicer
}

func NewHandler(service Servicer) *Handler {
	return &Handler{service: service}
}
