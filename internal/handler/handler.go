package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"calc/internal/model"
	"calc/internal/service"
)

type Handler struct {
	srv service.Service
	mux *http.ServeMux
}

func (h *Handler) CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req model.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Expression == "" {
		resp := model.Response{Error: "Expression is not valid"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(resp)
		return
	}

	result, err := h.srv.Calc(req.Expression)
	var resp model.Response

	if err != nil {
		resp.Error = err.Error()
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		resp.Result = strconv.FormatFloat(result, 'f', -1, 64)
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func New(srv service.Service) http.Handler {
	handler := &Handler{
		srv: srv,
		mux: http.NewServeMux(),
	}

	handler.mux.HandleFunc("POST /api/v1/calculate", handler.CalcHandler)

	return handler.mux
}
