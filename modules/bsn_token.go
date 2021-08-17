package modules

import (
	"context"
	"github.com/bianjieai/irita-sdk-go/types"
)

// Token token
type Token struct {
	projectId        string
	projectKey       string
	chainAccountAddr string
}

const (
	projectIdHeader           = "projectId"
	projectKeyHeader          = "projectKey"
	chainAccountAddressHeader = "chainAccountAddress"
)

// GetRequestMetadata 获取当前请求认证所需的元数据
func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{projectIdHeader: t.projectId, projectKeyHeader: t.projectKey, chainAccountAddressHeader: t.chainAccountAddr}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输
func (t *Token) RequireTransportSecurity() bool {
	return false
}

func NewBsnToken(info types.BSNProjectInfo) *Token {
	return &Token{
		projectId:        info.ProjectId,
		projectKey:       info.ProjectKey,
		chainAccountAddr: info.ChainAccountAddress,
	}
}
