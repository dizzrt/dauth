package dto

type ListServiceProviderRequest struct {
	Page     uint32 `json:"page"`
	PageSize uint32 `json:"page_size"`
}
