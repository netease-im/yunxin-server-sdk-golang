package endpoint

// Endpoints 端点集合
type Endpoints struct {
	DefaultEndpoint string   `json:"default_endpoint"` // 默认端点
	BackupEndpoints []string `json:"backup_endpoints"` // 备用端点列表
}

// GetDefaultEndpoint 获取默认端点
func (e *Endpoints) GetDefaultEndpoint() string {
	return e.DefaultEndpoint
}

// SetDefaultEndpoint 设置默认端点
func (e *Endpoints) SetDefaultEndpoint(defaultEndpoint string) {
	e.DefaultEndpoint = defaultEndpoint
}

// GetBackupEndpoints 获取备用端点列表
func (e *Endpoints) GetBackupEndpoints() []string {
	return e.BackupEndpoints
}

// SetBackupEndpoints 设置备用端点列表
func (e *Endpoints) SetBackupEndpoints(backupEndpoints []string) {
	e.BackupEndpoints = backupEndpoints
}
