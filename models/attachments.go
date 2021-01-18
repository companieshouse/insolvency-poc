package models

type AttachmentResponse struct {
	AttachmentId string                  `json:"attachment_id"`
	Links        AttachmentResponseLinks `json:"links"`
}

type AttachmentResponseLinks struct {
	Self string `json:"self"`
}

type AttachmentResourceResponse struct {
	Etag           string                          `json:"etag"`
	Kind           string                          `json:"kind"`
	Status         string                          `json:"status"`
	AttachmentType string                          `json:"attachment_type"`
	File           FileDetails                     `json:"file"`
	Links          AttachmentResourceResponseLinks `json:"links"`
}

type FileDetails struct {
	Name        string `json:"name"`
	Size        int    `json:"size"`
	ContentType string `json:"content_type"`
}

type AttachmentResourceResponseLinks struct {
	Self     string `json:"self"`
	Download string `json:"download"`
}
