package entity

// PageSetting 分页信息
type PageSetting struct {
    PageNum  int `json:"page_num,omitempty"`
    PageSize int `json:"page_size,omitempty"`
}
