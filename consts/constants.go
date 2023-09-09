package consts

type CompanyType string

const (
	CompanyTypeCorporations       CompanyType = "Corporations"
	CompanyTypeNonProfit          CompanyType = "NonProfit"
	CompanyTypeCooperative        CompanyType = "Cooperative"
	CompanyTypeSoleProprietorship CompanyType = "Sole Proprietorship"
)

const (
	CorrelationID string = "correlation-id"
)

const (
	KafkaEventNotifierTopic = "kafka-event-notifier"
	KafkaGroupId            = "kafka-group-id"
)
