package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/companieshouse/insolvency-poc/models"
	"io"
	"net"
	"net/http"
	"time"
)

func HandlePostAttachment(w http.ResponseWriter, req *http.Request) {

	attachmentResponse := models.AttachmentResponse{
		AttachmentId: "3fa85f64-5717-4562-b3fc-2c963f66afa6",
		Links:        models.AttachmentResponseLinks{Self: "/transactions/987654321/insolvency/attachment/3fa85f64-5717-4562-b3fc-2c963f66afa6"},
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(attachmentResponse)
	if err != nil {
		fmt.Println(fmt.Errorf("error encoding attachment to response: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleGetAttachment(w http.ResponseWriter, req *http.Request) {

	attachmentResourceResponse := models.AttachmentResourceResponse{
		Etag:           "etag123",
		Kind:           "insolvency#attachment",
		Status:         "submitted",
		AttachmentType: "special-resolution",
		File: models.FileDetails{
			Name:        "Statement.pdf",
			Size:        1024,
			ContentType: "application/pdf",
		},
		Links: models.AttachmentResourceResponseLinks{
			Self:     "/transactions/123456789/insolvency/attachment/3fa85f64-5717-4562-b3fc-2c963f66afa6",
			Download: "/transactions/123456789/insolvency/attachment/3fa85f64-5717-4562-b3fc-2c963f66afa6/download",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(attachmentResourceResponse)
	if err != nil {
		fmt.Println(fmt.Errorf("error encoding attachment details to response: [%v]", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func HandleDeleteAttachment(w http.ResponseWriter, req *http.Request) {

	// Delete attachment

	w.WriteHeader(http.StatusOK)
}

func HandleDownloadAttachment(w http.ResponseWriter, req *http.Request) {
	url := "https://upload.wikimedia.org/wikipedia/commons/f/fb/Companies_House%2C_Cardiff_-_geograph.org.uk_-_1521422.jpg"

	timeout := time.Duration(5) * time.Second
	transport := &http.Transport{
		ResponseHeaderTimeout: timeout,
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, timeout)
		},
		DisableKeepAlives: true,
	}
	client := &http.Client{
		Transport: transport,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	w.Header().Set("Content-Disposition", "attachment; filename=CompaniesHouse.png")
	w.Header().Set("Content-Type", req.Header.Get("Content-Type"))
	w.Header().Set("Content-Length", req.Header.Get("Content-Length"))

	//stream the body to the client without fully loading it into memory
	io.Copy(w, resp.Body)
}
