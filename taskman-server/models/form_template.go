package models

type FormTemplateTable struct {
	Id          string `json:"id" xorm:"id"`
	Name        string `json:"name" xorm:"name"`
	Description string `json:"description" xorm:"description"`
	Version     string `json:"version" xorm:"version"`
	ConfirmTime string `json:"confirmTime" xorm:"confirm_time"`
	CreatedBy   string `json:"createdBy" xorm:"created_by"`
	CreatedTime string `json:"createdTime" xorm:"created_time"`
	UpdatedBy   string `json:"updatedBy" xorm:"updated_by"`
	UpdatedTime string `json:"updatedTime" xorm:"updated_time"`
	DelFlag     int    `json:"delFlag" xorm:"del_flag"`
}

type FormTable struct {
	Id           string `json:"id" xorm:"id"`
	Name         string `json:"name" xorm:"name"`
	Description  string `json:"description" xorm:"description"`
	FormTemplate string `json:"formTemplate" xorm:"form_template"`
	FormVersion  string `json:"formVersion" xorm:"form_version"`
	CreatedBy    string `json:"createdBy" xorm:"created_by"`
	CreatedTime  string `json:"createdTime" xorm:"created_time"`
	UpdatedBy    string `json:"updatedBy" xorm:"updated_by"`
	UpdatedTime  string `json:"updatedTime" xorm:"updated_time"`
	DelFlag      int    `json:"delFlag" xorm:"del_flag"`
}

type FormItemTemplateTable struct {
	Id              string `json:"id" xorm:"id"`
	RecordId        string `json:"record_id" xorm:"record_id"`
	Version         string `json:"version" xorm:"version"`
	Name            string `json:"name" xorm:"name"`
	Description     string `json:"description" xorm:"description"`
	FormTemplate    string `json:"formTemplate" xorm:"form_template"`
	DefaultValue    string `json:"defaultValue" xorm:"default_value"`
	Sort            int    `json:"sort" xorm:"sort"`
	PackageName     string `json:"packageName" xorm:"package_name"`
	Entity          string `json:"entity" xorm:"entity"`
	AttrDefId       string `json:"attrDefId" xorm:"attr_def_id"`
	AttrDefName     string `json:"attrDefName" xorm:"attr_def_name"`
	AttrDefDataType string `json:"attrDefDataType" xorm:"attr_def_data_type"`
	ElementType     string `json:"elementType" xorm:"element_type"`
	Title           string `json:"title" xorm:"title"`
	Width           int    `json:"width" xorm:"width"`
	RefPackageName  string `json:"refPackageName" xorm:"ref_package_name"`
	RefEntity       string `json:"refEntity" xorm:"ref_entity"`
	DataOptions     string `json:"dataOptions" xorm:"data_options"`
	Required        string `json:"required" xorm:"required"`
	Regular         string `json:"regular" xorm:"regular"`
	IsEdit          string `json:"isEdit" xorm:"is_edit"`
	IsView          string `json:"isView" xorm:"is_view"`
	IsOutput        string `json:"isOutput" xorm:"is_output"`
}

type FormItemTable struct {
	Id               string `json:"id" xorm:"id"`
	Form             string `json:"form" xorm:"form"`
	FormItemTemplate string `json:"formItemTemplate" xorm:"form_item_template"`
	Name             string `json:"name" xorm:"name"`
	Value            string `json:"value" xorm:"value"`
}

type RequestFormTemplateDto struct {
	Id          string                   `json:"id"`
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	UpdatedTime string                   `json:"updatedTime"`
	Items       []*FormItemTemplateTable `json:"items"`
}