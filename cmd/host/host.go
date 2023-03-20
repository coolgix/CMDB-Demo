package host

import "context"

//定义host的数据模型

//定义interface
//host app service 的 接口定义(增删改查)
type Service interface {
	//录入主机信息
	CreateHost(context.Context, *Host) (*Host, error)
	//查询主机列表 ,有些数据不方便再列表返回，主机信息分为几个部分
	//一级目录前端给予通用信息的展示
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
	//查询主机详情，详情页给予主机的详细信息展示
	DescribeHost(context.Context, *DescribeHostRequest) (*Host, error)
	//更新主机信息
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	//删除主机，比如前端需要打印当前删除主机的ip信息，或者其他信息
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}

//返回主机列表对象，方便后期进行接口拓展
type HostSet struct {
	Items []*Host //host列表
	Total int     //分页
}

//host模型对象定义
type Host struct {
	//资源的公共属性部分
	*Resource
	//资源独有属性
	*Describe
}

//定义Vendor数据结构
type Vendor int

const (
	//枚举的默认值 默认是1  PrivateIDC
	Private_IDC Vendor = iota
	//默认加一 编译器自动增加
	ALIYUN
	TXYUN
	HUAWEIYUN
)

type Resource struct {
	Id     string `json:"id"`     // 全局唯一Id
	Vendor Vendor `json:"vendor"` // 厂商
	Region string `json:"region"` // 地域
	//Zone        string            `json:"zone"`        // 区域
	CreateAt int64 `json:"create_at"` // 创建时间
	ExpireAt int64 `json:"expire_at"` // 过期时间
	//Category    string            `json:"category"`    // 种类
	Type string `json:"type"` // 规格
	//InstanceId  string            `json:"instance_id"` // 实例ID
	Name        string            `json:"name"`        // 名称
	Description string            `json:"description"` // 描述
	Status      string            `json:"status"`      // 服务商中的状态
	Tags        map[string]string `json:"tags"`        // 标签
	UpdateAt    int64             `json:"update_at"`   // 更新时间
	SyncAt      int64             `json:"sync_at"`     // 同步时间
	Account     string            `json:"accout"`      // 资源所属账号
	PublicIP    string            `json:"public_ip"`   // 公网IP
	PrivateIP   string            `json:"private_ip"`  // 内网IP
	PayType     string            `json:"pay_type"`    // 实例付费方式
}

type Describe struct {
	//ResourceId   string `json:"resource_id"`   // 关联Resource
	CPU          int    `json:"cpu"`           // 核数
	Memory       int    `json:"memory"`        // 内存
	GPUAmount    int    `json:"gpu_amount"`    // GPU数量
	GPUSpec      string `json:"gpu_spec"`      // GPU类型
	OSType       string `json:"os_type"`       // 操作系统类型，分为Windows和Linux
	OSName       string `json:"os_name"`       // 操作系统名称
	SerialNumber string `json:"serial_number"` // 序列号
	//ImageID      string `json:"image_id"`      // 镜像ID
	//InternetMaxBandwidthOut int    `json:"internet_max_bandwidth_out"` // 公网出带宽最大值，单位为 Mbps
	//InternetMaxBandwidthIn  int    `json:"internet_max_bandwidth_in"`  // 公网入带宽最大值，单位为 Mbps
	//KeyPairName             string `json:"key_pair_name"`              // 秘钥对名称
	//SecurityGroups          string `json:"security_groups"`            // 安全组  采用逗号分隔
}

//定义QueryHost 请求对象
type QueryHostRequest struct {
}

//查询主机详情
type DescribeHostRequest struct {
}

//定义UpdateHost 对象
type UpdateHostRequest struct {
	//只可以跟新Describe 信息
	*Describe
}

//定义DeleteHost 对象
type DeleteHostRequest struct {
	//根据全局唯一ID删除
	Id string
}
