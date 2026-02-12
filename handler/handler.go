package handler

import "workshop-pwa-api/api/vallaris"

type handler struct {
	vaAPI *vallaris.VallarisAPI
}

func NewHandler(vaAPI *vallaris.VallarisAPI) *handler {
	return &handler{
		vaAPI: vaAPI,
	}
}
