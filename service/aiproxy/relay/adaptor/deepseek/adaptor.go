package deepseek

import (
	"github.com/labring/sealos/service/aiproxy/model"
	"github.com/labring/sealos/service/aiproxy/relay/adaptor"
	"github.com/labring/sealos/service/aiproxy/relay/adaptor/openai"
	"github.com/labring/sealos/service/aiproxy/relay/meta"
)

var _ adaptor.Adaptor = (*Adaptor)(nil)

type Adaptor struct {
	openai.Adaptor
}

const baseURL = "https://api.deepseek.com"

func (a *Adaptor) GetRequestURL(meta *meta.Meta) (string, error) {
	if meta.Channel.BaseURL == "" {
		meta.Channel.BaseURL = baseURL
	}
	return a.Adaptor.GetRequestURL(meta)
}

func (a *Adaptor) GetModelList() []*model.ModelConfig {
	return ModelList
}

func (a *Adaptor) GetChannelName() string {
	return "deepseek"
}