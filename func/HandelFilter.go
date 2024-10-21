package groupie

import (
	"net/http"
	"strconv"
	"strings"
)

func HandelFilter(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		Error(res, 405, "Method Not Allowed")
	}
	filter := req.FormValue("filter")
	if filter == "Filter" {
		Data.Filters = nil
		FromCreationDate := req.FormValue("FromCreationDate")
		ToCreationDate := req.FormValue("ToCreationDate")
		FromFirsetAlbum := req.FormValue("FromFirsetAlbum")
		ToFirsetAlbum := req.FormValue("ToFirsetAlbum")
		local := req.FormValue("local")
		local = strings.ReplaceAll(local, ", ", "-")
		local = strings.ToLower(local)
		req.ParseForm()
		Members := req.Form["members"]
		for i, artist := range Data.Arts {
			Form, err := strconv.Atoi(FromCreationDate)
			if err != nil {
				Error(res, 400, "Oops!! Bade Request")
				return
			}
			To, err := strconv.Atoi(ToCreationDate)
			if err != nil {
				Error(res, 400, "Oops!! Bade Request")
				return
			}

			FormAlbum, err := strconv.Atoi(FromFirsetAlbum)
			if err != nil {
				Error(res, 400, "Oops!! Bade Request")
				return
			}
			ToAlbum, err := strconv.Atoi(ToFirsetAlbum)
			if err != nil {
				Error(res, 400, "Oops!! Bade Request")
				return
			}
			FirstAlbum, err := strconv.Atoi(artist.FirstAlbum[6:])
			if err != nil {
				Error(res, 400, "Oops!! Bade Request")
				return
			}
			if (artist.CreationDate >= Form && artist.CreationDate <= To) && (FirstAlbum >= FormAlbum && FirstAlbum <= ToAlbum) {
				for _, nMembers := range Members {
					num, err := strconv.Atoi(nMembers)
					if err != nil {
						Error(res, 400, "Oops!! Bade Request")
						return
					}
					if num == len(artist.Members) {
						for _, location := range Data.DataLocals["index"][i].Locations {
							if strings.Contains(location, local) {
								Data.Filters = append(Data.Filters, artist)
								break
							}
						}
					}
				}
				if len(Members) == 0 {
					for _, location := range Data.DataLocals["index"][i].Locations {
						if strings.Contains(location, local) {
							Data.Filters = append(Data.Filters, artist)
							break
						}
					}
				}
			}

		}
	} else {
		Data.Filters = Data.Arts
	}
}
