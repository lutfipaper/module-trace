package interfaces

type TracingConfig struct {
	PrintOperation bool   `yaml:"printoperation"`
	Signoz         Signoz `yaml:"signoz"`
}

type Signoz struct {
	Enable   bool   `yaml:"enable" desc:"tracing:signoz:enable"`
	Url      string `yaml:"url" desc:"tracing:signoz:url" default:"127.0.0.1:v"`
	Endpoint string `yaml:"endpoint" desc:"tracing:signoz:endpoint"`
	Insecure string `yaml:"insecure" desc:"tracing:signoz:insecure"`
}

var TracingConfigManual = map[string]string{
	"tracing:signoz:enable":   `enable signoz tracing`,
	"tracing:signoz:url":      `like 10.10.10.10:v`,
	"tracing:signoz:endpoint": `like http://10.10.10.10:14268/api/traces`,
}

func SetManual(man map[string]string) map[string]string {
	for k, v := range TracingConfigManual {
		man[k] = v
	}
	return man
}
