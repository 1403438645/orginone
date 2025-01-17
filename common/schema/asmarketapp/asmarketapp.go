// Code generated by entc, DO NOT EDIT.

package asmarketapp

import (
	"orginone/common/tools/date"
)

const (
	// Label holds the string label denoting the asmarketapp type in the database.
	Label = "as_market_app"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAppName holds the string denoting the app_name field in the database.
	FieldAppName = "app_name"
	// FieldIcon holds the string denoting the icon field in the database.
	FieldIcon = "icon"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldContact holds the string denoting the contact field in the database.
	FieldContact = "contact"
	// FieldContactName holds the string denoting the contact_name field in the database.
	FieldContactName = "contact_name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldFile holds the string denoting the file field in the database.
	FieldFile = "file"
	// FieldSaleStatus holds the string denoting the sale_status field in the database.
	FieldSaleStatus = "sale_status"
	// FieldTenantID holds the string denoting the tenant_id field in the database.
	FieldTenantID = "tenant_id"
	// FieldPlatform holds the string denoting the platform field in the database.
	FieldPlatform = "platform"
	// FieldTargetUser holds the string denoting the target_user field in the database.
	FieldTargetUser = "target_user"
	// FieldDeployStatus holds the string denoting the deploy_status field in the database.
	FieldDeployStatus = "deploy_status"
	// FieldDeployAddress holds the string denoting the deploy_address field in the database.
	FieldDeployAddress = "deploy_address"
	// FieldDeployType holds the string denoting the deploy_type field in the database.
	FieldDeployType = "deploy_type"
	// FieldPublishTime holds the string denoting the publish_time field in the database.
	FieldPublishTime = "publish_time"
	// FieldAppType holds the string denoting the app_type field in the database.
	FieldAppType = "app_type"
	// FieldApplyTime holds the string denoting the apply_time field in the database.
	FieldApplyTime = "apply_time"
	// FieldAppAddress holds the string denoting the app_address field in the database.
	FieldAppAddress = "app_address"
	// FieldAppMail holds the string denoting the app_mail field in the database.
	FieldAppMail = "app_mail"
	// FieldAppPhoto holds the string denoting the app_photo field in the database.
	FieldAppPhoto = "app_photo"
	// FieldAppField holds the string denoting the app_field field in the database.
	FieldAppField = "app_field"
	// FieldAppCategory holds the string denoting the app_category field in the database.
	FieldAppCategory = "app_category"
	// FieldAppProjectSource holds the string denoting the app_project_source field in the database.
	FieldAppProjectSource = "app_project_source"
	// FieldAppStar holds the string denoting the app_star field in the database.
	FieldAppStar = "app_star"
	// FieldAppFoundsSource holds the string denoting the app_founds_source field in the database.
	FieldAppFoundsSource = "app_founds_source"
	// FieldInnerURL holds the string denoting the inner_url field in the database.
	FieldInnerURL = "inner_url"
	// FieldOutURL holds the string denoting the out_url field in the database.
	FieldOutURL = "out_url"
	// FieldReformStatus holds the string denoting the reform_status field in the database.
	FieldReformStatus = "reform_status"
	// FieldOutIPURL holds the string denoting the out_ip_url field in the database.
	FieldOutIPURL = "out_ip_url"
	// FieldIsDeleted holds the string denoting the is_deleted field in the database.
	FieldIsDeleted = "is_deleted"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreateUser holds the string denoting the create_user field in the database.
	FieldCreateUser = "create_user"
	// FieldUpdateUser holds the string denoting the update_user field in the database.
	FieldUpdateUser = "update_user"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// EdgeAppMenus holds the string denoting the appmenus edge name in mutations.
	EdgeAppMenus = "appMenus"
	// EdgeAppRoles holds the string denoting the approles edge name in mutations.
	EdgeAppRoles = "appRoles"
	// EdgeUseds holds the string denoting the useds edge name in mutations.
	EdgeUseds = "useds"
	// EdgeAppAlerts holds the string denoting the appalerts edge name in mutations.
	EdgeAppAlerts = "appAlerts"
	// EdgeAppRedeploys holds the string denoting the appredeploys edge name in mutations.
	EdgeAppRedeploys = "appRedeploys"
	// EdgeAppKeys holds the string denoting the appkeys edge name in mutations.
	EdgeAppKeys = "appKeys"
	// EdgeAppPurchases holds the string denoting the apppurchases edge name in mutations.
	EdgeAppPurchases = "appPurchases"
	// EdgeAppComponents holds the string denoting the appcomponents edge name in mutations.
	EdgeAppComponents = "appComponents"
	// EdgeAppGroupDistribs holds the string denoting the appgroupdistribs edge name in mutations.
	EdgeAppGroupDistribs = "appGroupDistribs"
	// EdgeAppGroupDistribConfigs holds the string denoting the appgroupdistribconfigs edge name in mutations.
	EdgeAppGroupDistribConfigs = "appGroupDistribConfigs"
	// EdgeAppGroupDistribsRelation holds the string denoting the appgroupdistribsrelation edge name in mutations.
	EdgeAppGroupDistribsRelation = "appGroupDistribsRelation"
	// Table holds the table name of the asmarketapp in the database.
	Table = "as_market_app"
	// AppMenusTable is the table that holds the appMenus relation/edge.
	AppMenusTable = "as_market_menu"
	// AppMenusInverseTable is the table name for the AsMarketMenu entity.
	// It exists in this package in order to avoid circular dependency with the "asmarketmenu" package.
	AppMenusInverseTable = "as_market_menu"
	// AppMenusColumn is the table column denoting the appMenus relation/edge.
	AppMenusColumn = "app_id"
	// AppRolesTable is the table that holds the appRoles relation/edge.
	AppRolesTable = "as_market_app_role"
	// AppRolesInverseTable is the table name for the AsMarketAppRole entity.
	// It exists in this package in order to avoid circular dependency with the "asmarketapprole" package.
	AppRolesInverseTable = "as_market_app_role"
	// AppRolesColumn is the table column denoting the appRoles relation/edge.
	AppRolesColumn = "app_id"
	// UsedsTable is the table that holds the useds relation/edge.
	UsedsTable = "as_market_used_app"
	// UsedsInverseTable is the table name for the AsMarketUsedApp entity.
	// It exists in this package in order to avoid circular dependency with the "asmarketusedapp" package.
	UsedsInverseTable = "as_market_used_app"
	// UsedsColumn is the table column denoting the useds relation/edge.
	UsedsColumn = "app_id"
	// AppAlertsTable is the table that holds the appAlerts relation/edge.
	AppAlertsTable = "as_market_app_alert"
	// AppAlertsInverseTable is the table name for the AsMarketAppAlert entity.
	// It exists in this package in order to avoid circular dependency with the "asmarketappalert" package.
	AppAlertsInverseTable = "as_market_app_alert"
	// AppAlertsColumn is the table column denoting the appAlerts relation/edge.
	AppAlertsColumn = "alert_release_app_id"
	// AppRedeploysTable is the table that holds the appRedeploys relation/edge.
	AppRedeploysTable = "as_redeploy_data"
	// AppRedeploysInverseTable is the table name for the AsRedeployData entity.
	// It exists in this package in order to avoid circular dependency with the "asredeploydata" package.
	AppRedeploysInverseTable = "as_redeploy_data"
	// AppRedeploysColumn is the table column denoting the appRedeploys relation/edge.
	AppRedeploysColumn = "app_id"
	// AppKeysTable is the table that holds the appKeys relation/edge.
	AppKeysTable = "as_market_app_key_secret"
	// AppKeysInverseTable is the table name for the AsMarketAppKeySecret entity.
	// It exists in this package in order to avoid circular dependency with the "asmarketappkeysecret" package.
	AppKeysInverseTable = "as_market_app_key_secret"
	// AppKeysColumn is the table column denoting the appKeys relation/edge.
	AppKeysColumn = "app_id"
	// AppPurchasesTable is the table that holds the appPurchases relation/edge.
	AppPurchasesTable = "as_market_app_purchase"
	// AppPurchasesInverseTable is the table name for the AsMarketAppPurchase entity.
	// It exists in this package in order to avoid circular dependency with the "asmarketapppurchase" package.
	AppPurchasesInverseTable = "as_market_app_purchase"
	// AppPurchasesColumn is the table column denoting the appPurchases relation/edge.
	AppPurchasesColumn = "app_id"
	// AppComponentsTable is the table that holds the appComponents relation/edge.
	AppComponentsTable = "as_market_app_component"
	// AppComponentsInverseTable is the table name for the AsMarketAppComponent entity.
	// It exists in this package in order to avoid circular dependency with the "asmarketappcomponent" package.
	AppComponentsInverseTable = "as_market_app_component"
	// AppComponentsColumn is the table column denoting the appComponents relation/edge.
	AppComponentsColumn = "app_id"
	// AppGroupDistribsTable is the table that holds the appGroupDistribs relation/edge.
	AppGroupDistribsTable = "as_market_app_group_distribution"
	// AppGroupDistribsInverseTable is the table name for the AsMarketAppGroupDistribution entity.
	// It exists in this package in order to avoid circular dependency with the "asmarketappgroupdistribution" package.
	AppGroupDistribsInverseTable = "as_market_app_group_distribution"
	// AppGroupDistribsColumn is the table column denoting the appGroupDistribs relation/edge.
	AppGroupDistribsColumn = "app_id"
	// AppGroupDistribConfigsTable is the table that holds the appGroupDistribConfigs relation/edge.
	AppGroupDistribConfigsTable = "as_app_group_distribution_data"
	// AppGroupDistribConfigsInverseTable is the table name for the AsAppGroupDistributionData entity.
	// It exists in this package in order to avoid circular dependency with the "asappgroupdistributiondata" package.
	AppGroupDistribConfigsInverseTable = "as_app_group_distribution_data"
	// AppGroupDistribConfigsColumn is the table column denoting the appGroupDistribConfigs relation/edge.
	AppGroupDistribConfigsColumn = "app_id"
	// AppGroupDistribsRelationTable is the table that holds the appGroupDistribsRelation relation/edge.
	AppGroupDistribsRelationTable = "as_market_app_group_distribution_relation"
	// AppGroupDistribsRelationInverseTable is the table name for the AsMarketAppGroupDistributionRelation entity.
	// It exists in this package in order to avoid circular dependency with the "asmarketappgroupdistributionrelation" package.
	AppGroupDistribsRelationInverseTable = "as_market_app_group_distribution_relation"
	// AppGroupDistribsRelationColumn is the table column denoting the appGroupDistribsRelation relation/edge.
	AppGroupDistribsRelationColumn = "app_id"
)

// Columns holds all SQL columns for asmarketapp fields.
var Columns = []string{
	FieldID,
	FieldAppName,
	FieldIcon,
	FieldVersion,
	FieldContact,
	FieldContactName,
	FieldDescription,
	FieldFile,
	FieldSaleStatus,
	FieldTenantID,
	FieldPlatform,
	FieldTargetUser,
	FieldDeployStatus,
	FieldDeployAddress,
	FieldDeployType,
	FieldPublishTime,
	FieldAppType,
	FieldApplyTime,
	FieldAppAddress,
	FieldAppMail,
	FieldAppPhoto,
	FieldAppField,
	FieldAppCategory,
	FieldAppProjectSource,
	FieldAppStar,
	FieldAppFoundsSource,
	FieldInnerURL,
	FieldOutURL,
	FieldReformStatus,
	FieldOutIPURL,
	FieldIsDeleted,
	FieldStatus,
	FieldCreateUser,
	FieldUpdateUser,
	FieldCreateTime,
	FieldUpdateTime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSaleStatus holds the default value on creation for the "sale_status" field.
	DefaultSaleStatus int64
	// DefaultPlatform holds the default value on creation for the "platform" field.
	DefaultPlatform int64
	// DefaultDeployType holds the default value on creation for the "deploy_type" field.
	DefaultDeployType int64
	// DefaultIsDeleted holds the default value on creation for the "is_deleted" field.
	DefaultIsDeleted int64
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int64
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() date.DateTime
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() date.DateTime
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() date.DateTime
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() int64
)
