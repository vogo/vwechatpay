package vwxapply4sub

// 结算规则ID常量定义
// 参考：https://kf.qq.com/faq/220228IJb2UV220228uEjU3Q.html

// 小微商户结算规则
const (
	// SettlementRuleIDMicro 小微商户结算规则ID
	// 主体类型：小微
	// 结算规则：费率0.6%，入账周期T+1
	// 适用行业：提供网上交易场所或信息服务业务、餐饮、零售、交通出行等实体业务
	SettlementRuleIDMicro = "703"
)

// 个体户结算规则
const (
	// SettlementRuleIDIndividual 个体户结算规则ID
	// 主体类型：个体户
	// 结算规则：费率0.6%，入账周期T+1
	// 适用行业：提供网上交易场所或信息服务的业务、通讯业务、财经类业务及其他平台服务、餐饮、零售、交通出行等实体业务
	// 包含行业：餐饮、电商平台、零售、食品生鲜、咨询/娱乐票务、房产中介、宠物医院、共享服务、休闲娱乐/旅游服务等
	SettlementRuleIDIndividual = "719"

	// SettlementRuleIDIndividualVirtual 个体户虚拟业务结算规则ID
	// 主体类型：个体户
	// 结算规则：费率0.6%，入账周期T+1，虚拟限额
	// 适用行业：通讯业务
	// 包含行业：婚介平台/就业信息平台/其他信息服务平台、虚拟充值
	SettlementRuleIDIndividualVirtual = "720"

	// SettlementRuleIDIndividualOil 个体户加油结算规则ID
	// 主体类型：个体户
	// 结算规则：费率0.3%，入账周期T+1
	// 适用行业：加油
	// 包含行业：快递、物流、加油/加气
	SettlementRuleIDIndividualOil = "721"

	// SettlementRuleIDIndividualUtility 个体户民生缴费结算规则ID
	// 主体类型：个体户
	// 结算规则：费率0.2%，入账周期T+1
	// 适用行业：民生缴费
	// 包含行业：水电煤
	SettlementRuleIDIndividualUtility = "790"
)

// 平台收付通二级商户结算规则
const (
	// SettlementRuleIDPlatformMicro 平台收付通二级商户-小微商户结算规则ID
	// 主体类型：小微
	// 结算规则：费率0.6%，入账周期T+1
	// 适用行业：提供网上交易场所或信息服务的业务、餐饮、零售、交通出行等实体业务
	SettlementRuleIDPlatformMicro = "747"

	// SettlementRuleIDPlatformIndividual 平台收付通二级商户-个体户结算规则ID
	// 主体类型：个体户
	// 结算规则：费率0.6%，入账周期T+1
	// 适用行业：提供网上交易场所或信息服务的业务、餐饮、零售、医疗、交通出行等实体业务
	SettlementRuleIDPlatformIndividual = "802"

	// SettlementRuleIDPlatformLogisticsMicro 平台收付通二级商户-小微物流结算规则ID
	// 主体类型：小微
	// 结算规则：费率0.3%，入账周期T+1
	// 适用行业：物流快递服务
	SettlementRuleIDPlatformLogisticsMicro = "777"

	// SettlementRuleIDPlatformLogisticsIndividual 平台收付通二级商户-个体户物流结算规则ID
	// 主体类型：个体户
	// 结算规则：费率0.3%，入账周期T+1
	// 适用行业：物流快递服务
	SettlementRuleIDPlatformLogisticsIndividual = "779"

	// SettlementRuleIDPlatformLogisticsEnterprise 平台收付通二级商户-企业物流结算规则ID
	// 主体类型：企业
	// 结算规则：费率0.3%，入账周期T+1
	// 适用行业：物流快递服务
	SettlementRuleIDPlatformLogisticsEnterprise = "801"

	// SettlementRuleIDPlatformLogisticsInstitution 平台收付通二级商户-事业单位物流结算规则ID
	// 主体类型：事业单位
	// 结算规则：费率0.3%，入账周期T+1
	// 适用行业：物流快递服务
	SettlementRuleIDPlatformLogisticsInstitution = "778"

	// SettlementRuleIDPlatformLogisticsSocial 平台收付通二级商户-社会组织物流结算规则ID
	// 主体类型：社会组织
	// 结算规则：费率0.3%，入账周期T+1
	// 适用行业：物流快递服务
	SettlementRuleIDPlatformLogisticsSocial = "805"
)
