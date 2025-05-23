package models

type QueryRequestFilterObj struct {
	Name     string      `json:"name"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

type QueryRequestSorting struct {
	Asc   bool   `json:"asc"`
	Field string `json:"field"`
}

type QueryRequestDialect struct {
	AssociatedData map[string]string `json:"associatedData"`
	QueryMode      string            `json:"queryMode"`
}

type QueryRequestParam struct {
	Filters       []*QueryRequestFilterObj `json:"filters"`
	Dialect       *QueryRequestDialect     `json:"dialect"`
	Paging        bool                     `json:"paging"`
	Pageable      *PageInfo                `json:"pageable"`
	Sorting       *QueryRequestSorting     `json:"sorting"`
	ResultColumns []string                 `json:"resultColumns"`
}

type GetRequestTemplateRolesParam struct {
	RequestTemplateIds []string `json:"requestTemplateIds"`
}

type TransFiltersParam struct {
	IsStruct   bool
	StructObj  interface{}
	Prefix     string
	KeyMap     map[string]string
	PrimaryKey string
}

type PlatformRequestParam struct {
	Tab            string `json:"tab"`             // 标签,取值有:pending 待处理,myPending    hasProcessed 已处理  submit 我提交的 draft 我的暂存  collect 收藏
	Action         int    `json:"action"`          // 行为, 1表示发布行为,2请求,3问题,4事件, 5变更
	Type           int    `json:"type"`            // 0 所有,1表示请求定版,2 任务处理,3 审批 4确认请求
	Rollback       int    `json:"rollback"`        // 0代表所有,1表示被退回,2表示其他,3表示自己撤销(submit 我提交的tab下才有这个筛选生效)
	QueryTimeStart string `json:"queryTimeStart" ` // 查询时间-请求创建时间开始
	QueryTimeEnd   string `json:"queryTimeEnd" `   // 查询时间-请求创建时间结束
	CommonRequestParam
}

type FilterRequestParam struct {
	StartTime string `json:"startTime"` //开始时间
}

// QueryCollectTemplateParam 模板查询条件
type QueryCollectTemplateParam struct {
	Action           int                  `json:"action"`           // 行为, 1表示发布行为,2请求,3问题,4事件, 5变更
	Id               string               `json:"id"`               // ID
	Name             string               `json:"name"`             // Name
	TemplateGroupId  []string             `json:"templateGroupId"`  // 模板组id
	OperatorObjType  []string             `json:"operatorObjType"`  // 操作对象类型
	ProcDefName      []string             `json:"procDefName"`      // 使用编排
	ManageRole       []string             `json:"manageRole"`       // 属主角色
	Owner            []string             `json:"owner"`            // 属主
	UseRole          []string             `json:"useRole"`          // 使用角色
	Tags             []string             `json:"tags" `            // 标签
	CreatedStartTime string               `json:"createdTimeStart"` // 创建开始时间
	CreatedEndTime   string               `json:"createdTimeEnd"`   // 创建结束时间
	UpdatedStartTime string               `json:"updatedTimeStart"` // 更新开始时间
	UpdatedEndTime   string               `json:"updatedTimeEnd"`   // 更新结束时间
	StartIndex       int                  `json:"startIndex"`
	PageSize         int                  `json:"pageSize"`
	Sorting          *QueryRequestSorting `json:"sorting"` // 排序字段
}

type CountPlatformParam struct {
	Tab            string `json:"tab"`             // 标签,取值有:pending 待处理,myPending 本人待处理,hasProcessed 已处理  submit 我提交的 draft 我的暂存, all表示查询所有
	QueryTimeStart string `json:"queryTimeStart" ` // 查询时间-开始
	QueryTimeEnd   string `json:"queryTimeEnd" `   // 查询时间结束
}

// AddCollectTemplateParam 添加收藏
type AddCollectTemplateParam struct {
	TemplateId string `json:"templateId"` //模板id
	Role       string `json:"role"`       //模板使用角色
}

type RequestHistoryParam struct {
	Tab        string `json:"tab"`        // 标签,取值有: draft 暂存(不包含被退回和撤回)  commit 已经提交  rollback 被退回  revoke 撤回
	Permission string `json:"permission"` // 权限,取值有: group 本组,  all 表示所有
	Action     int    `json:"action"`     // 行为: 0表示所有,1表示发布行为,2请求,3问题,4事件, 5变更
	CommonRequestParam
}

type RequestConfirmParam struct {
	Id             string   `json:"id"`             // 请求ID
	TaskId         string   `json:"taskId"`         // 请求确认任务ID
	MarkTaskId     []string `json:"markTaskId"`     // 任务id
	CompleteStatus string   `json:"completeStatus"` // 任务完成状态 取值：complete or uncompleted
	Notes          string   `json:"notes"`          // 备注
}

type CommonRequestParam struct {
	Id                         string               `json:"id"`                          // ID
	Name                       string               `json:"name"`                        // Name
	Query                      string               `json:"query"`                       // Id and Name 通用查询
	TemplateId                 []string             `json:"templateId"`                  // 模版id
	Status                     []string             `json:"status"`                      // 请求状态 Pending InProgress(Faulted)
	OperatorObj                string               `json:"operatorObj"`                 // 操作对象
	CreatedBy                  []string             `json:"createdBy"`                   // 创建人
	OperatorObjType            []string             `json:"operatorObjType"`             // 操作对象类型
	ProcDefName                []string             `json:"procDefName"`                 // 使用编排
	Handler                    []string             `json:"handler"`                     // 当前处理人
	CreatedStartTime           string               `json:"createdTimeStart"`            // 创建开始时间
	CreatedEndTime             string               `json:"createdTimeEnd"`              // 创建结束时间
	ReportStartTime            string               `json:"reportTimeStart"`             // 请求提交开始时间
	ReportEndTime              string               `json:"reportTimeEnd"`               // 请求提交结束时间
	ApprovalStartTime          string               `json:"approvalTimeStart"`           // 请求审批开始时间
	ApprovalEndTime            string               `json:"approvalTimeEnd" `            // 请求审批结束时间
	UpdatedStartTime           string               `json:"updatedTimeStart" `           // 更新开始时间
	UpdatedEndTime             string               `json:"updatedTimeEnd" `             // 更新结束时间
	ExpectStartTime            string               `json:"expectTimeStart" `            // 期望开始时间
	ExpectEndTime              string               `json:"expectTimeEnd" `              // 期望结束时间
	TaskName                   string               `json:"taskName"`                    // 任务名称
	TaskCreatedStartTime       string               `json:"taskCreatedTimeStart" `       // 任务提交开始时间
	TaskCreatedEndTime         string               `json:"taskCreatedTimeEnd" `         // 任务提交结束时间
	TaskApprovalStartTime      string               `json:"taskApprovalTimeStart" `      // 任务审批开始时间
	TaskApprovalEndTime        string               `json:"taskApprovalTimeEnd" `        // 任务审批结束时间
	TaskExpectStartTime        string               `json:"taskExpectTimeStart" `        // 任务期望开始时间
	TaskExpectEndTime          string               `json:"taskExpectTimeEnd" `          // 任务期望结束时间
	TaskHandleUpdatedStartTime string               `json:"taskHandleUpdatedTimeStart" ` // 处理节点更新开始时间
	TaskHandleUpdatedEndTime   string               `json:"taskHandleUpdatedTimeEnd" `   // 处理结束时间
	RequestRefId               string               `json:"requestRefId" `               // 请求RefId
	StartIndex                 int                  `json:"startIndex"`
	PageSize                   int                  `json:"pageSize"`
	Sorting                    *QueryRequestSorting `json:"sorting"` // 排序字段
}

type PlatDataParam struct {
	Param      CommonRequestParam
	QueryParam []interface{}
	User       string
	UserToken  string
	Tab        string // 标签,取值有:pending 待处理    hasProcessed 已处理  submit 我提交的 draft 我的暂存  collect 收藏
}

type RequestQueryParam struct {
	TemplateId string `json:"templateId"` // 模板id
	RequestId  string `json:"requestId"`  // 请求id
}

type TerminateInstanceParam struct {
	ProcInstId  string `json:"procInstId"`
	ProcInstKey string `json:"procInstKey"`
}

// CommonParam 通用参数
type CommonParam struct {
	User     string   // 登录用户
	Roles    []string // 用户角色
	Token    string   // 鉴权token
	Language string   // 环境语言
}

type FormTemplateLibraryParam struct {
	Name     string                        `json:"name"`
	FormType string                        `json:"formType"`
	Items    []*FormItemTemplateLibraryDto `json:"items"`
}

type QueryFormTemplateLibraryParam struct {
	Name       string `json:"name"`      // 组件库名称
	FormType   string `json:"formType"`  // 表单类型
	Type       string `json:"type"`      // message 信息表单, data 数据表单全量查询
	CreatedBy  string `json:"createdBy"` // 创建人
	StartIndex int    `json:"startIndex"`
	PageSize   int    `json:"pageSize"`
}

type RequestAssociationParam struct {
	Query           string `json:"query"`  // 请求名称,模糊
	Action          int    `json:"action"` // 类型
	StartIndex      int    `json:"startIndex"`
	PageSize        int    `json:"pageSize"`
	ReportStartTime string `json:"reportTimeStart"` // 请求提交开始时间
	ReportEndTime   string `json:"reportTimeEnd"`   // 请求提交结束时间
}

type BatchExportRequestTemplateParam struct {
	RequestTemplateIds []string `json:"requestTemplateIds"`
}

type RequestFormSensitiveDataParam struct {
	RequestId    string `json:"requestId"`
	CiType       string `json:"ciType"`
	AttrName     string `json:"attrName"`
	AttrVal      string `json:"attrVal"`
	Guid         string `json:"guid"`
	TmpId        string `json:"tmpId"`
	TaskHandleId string `json:"taskHandleId"`
}
