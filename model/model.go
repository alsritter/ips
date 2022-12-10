package model

const (
	// IPv4 IP库版本 v4
	IPv4 uint16 = 0x01

	// IPv6 IP库版本 v6
	IPv6 uint16 = 0x02
)

// Meta 元数据
type Meta struct {

	// IPVersion IP库版本
	IPVersion uint16

	// Fields 数据字段列表
	Fields []string
}

// IsIPv4Support 是否支持 IPv4
func (m *Meta) IsIPv4Support() bool {
	return m.IPVersion&IPv4 == IPv4
}

// IsIPv6Support 是否支持 IPv6
func (m *Meta) IsIPv6Support() bool {
	return m.IPVersion&IPv6 == IPv6
}

// IPDBMeta ipdb 元数据
type IPDBMeta struct {

	// Build 构建时间 10位时间戳
	Build int `json:"build"`

	// IPVersion IP库版本
	IPVersion uint16 `json:"ip_version"`

	// Languages 支持语言
	// value为语言对应的fields偏移量
	Languages map[string]int `json:"languages"`

	// NodeCount 节点数量
	NodeCount int `json:"node_count"`

	// TotalSize 节点区域和数据区域大小统计
	TotalSize int `json:"total_size"`

	// Fields 数据字段列表
	// 城市级别数据库包含13个字段
	// "country_name": "国家名称"
	// "region_name": "省份名称"
	// "city_name": "城市名称"
	// "owner_domain": "所有者"
	// "isp_domain": "运营商"
	// "latitude": "纬度"
	// "longitude": "经度"
	// "timezone": "时区"
	// "utc_offset": "UTC偏移量"
	// "china_admin_code": "中国邮编"
	// "idd_code": "电话区号"
	// "country_code": "国家代码"
	// "continent_code": "大陆代码"
	Fields []string `json:"fields"`
}