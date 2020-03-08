package client

const (
	registerTenant          string = "register-tenant"
	registerTenantHandler   string = "register-tenant-handler"
	reloadTenant            string = "reload-tenant"
	unregisterTenant        string = "unregister-tenant"
	unregisterTenantHandler string = "unregister-tenant-handler"
)

// TODO: responses, test the commands

// RegisterTenantHandlerRequest requires a tenant id, htype and hargs
type RegisterTenantHandlerRequest struct {
	TenantID int    `json:"id"`
	Htype    string `json:"htype"`
	Hargs    int    `json:"hargs"`
}

// UnRegisterTenantHandlerRequest requires a tenant id, htype and hargs
type UnRegisterTenantHandlerRequest struct {
	TenantID int    `json:"id"`
	Htype    string `json:"htype"`
	Hargs    int    `json:"hargs"`
}

// RegisterTenantRequest requires a tenant id and the tenant yaml file
type RegisterTenantRequest struct {
	TenantID       int    `json:"id"`
	TenantYAMLFile string `json:"yaml"`
}

// ReloadTenantRequest requires a tenant id and the tenant yaml file to reload tenant
type ReloadTenantRequest struct {
	TenantID       int    `json:"id"`
	TenantYAMLFile string `json:"yaml"`
}

// UnRegisterTenantRequest requires a tenant id to unregister tenant
type UnRegisterTenantRequest struct {
	TenantID int `json:"id"`
}

// RegisterTenantHandlerResponse is response from "register-tentant-handler <id> <hargs> <hint>"
type RegisterTenantHandlerResponse struct{}

// UnRegisterTenantHandlerResponse is response from "unregister-tentant-handler <id> <hargs> <hint>"
type UnRegisterTenantHandlerResponse struct{}

// RegisterTenantResposne is response from "register-tenant <id> <yaml>"
type RegisterTenantResposne struct{}

// UnRegisterTenantResponse is response from "unregister-tenant <id>"
type UnRegisterTenantResponse struct{}

// ReloadTenantResponse is response from "reload-tenant <id> <yaml>"
type ReloadTenantResponse struct{}

// RegisterTenantHandlerCommand does register-tenant-handler
func (s *Socket) RegisterTenantHandlerCommand(registerTenantHandlerRequest RegisterTenantHandlerRequest) (*RegisterTenantHandlerResponse, error) {
	registerTenantHandlerResp := &RegisterTenantHandlerResponse{}
	err := s.DoCommand(registerTenantHandler, registerTenantHandlerRequest, registerTenantHandlerResp)
	return registerTenantHandlerResp, err
}

// UnRegisterTenantHandlerCommand does unregister-tenant-handler
func (s *Socket) UnRegisterTenantHandlerCommand(unRegisterTenantHandlerRequest RegisterTenantHandlerRequest) (*UnRegisterTenantHandlerResponse, error) {
	unRegisterTenantHandlerResp := &UnRegisterTenantHandlerResponse{}
	err := s.DoCommand(unregisterTenantHandler, unRegisterTenantHandlerRequest, unRegisterTenantHandlerResp)
	return unRegisterTenantHandlerResp, err
}

// RegisterTenantCommand does register-tenant
func (s *Socket) RegisterTenantCommand(registerTenantRequest RegisterTenantRequest) (*RegisterTenantResposne, error) {
	registerTenantResp := &RegisterTenantResposne{}
	err := s.DoCommand(registerTenant, registerTenantRequest, registerTenantResp)
	return registerTenantResp, err
}

// ReloadTenantCommand does reload-tenant
func (s *Socket) ReloadTenantCommand(reloadTenantRequest ReloadTenantRequest) (*ReloadTenantResponse, error) {
	reloadTenantResp := &ReloadTenantResponse{}
	err := s.DoCommand(reloadTenant, reloadTenantRequest, reloadTenantResp)
	return reloadTenantResp, err
}

// UnRegisterTenantCommand does unregister-tenant
func (s *Socket) UnRegisterTenantCommand(unRegisterTenantRequest UnRegisterTenantRequest) (*UnRegisterTenantResponse, error) {
	unRegisterTenantResp := &UnRegisterTenantResponse{}
	err := s.DoCommand(unregisterTenant, unRegisterTenantRequest, unRegisterTenantResp)
	return unRegisterTenantResp, err
}
