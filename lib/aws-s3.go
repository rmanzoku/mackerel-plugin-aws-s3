package mpawss3

import (
	"flag"

	mp "github.com/mackerelio/go-mackerel-plugin-helper"
)

var graphdef = map[string]mp.Graphs{
	"objects": {
		Label: "S3 Total Object",
		Unit:  "integer",
		Metrics: []mp.Metrics{
			{Name: "object", Label: "Total objects", Diff: false},
		},
	},
}

// S3Plugin mackerel plugin for S3
type S3Plugin struct {
	Prefix string
}

// FetchMetrics interface for mackerelplugin
func (s S3Plugin) FetchMetrics() (map[string]interface{}, error) {
	ret := make(map[string]interface{})

	ret["object"] = float64(100)

	return ret, nil

}

// GraphDefinition interface for mackerelplugin
func (s S3Plugin) GraphDefinition() map[string]mp.Graphs {

	return graphdef
}

// MetricKeyPrefix interface for PluginWithPrefix
func (s S3Plugin) MetricKeyPrefix() string {
	if s.Prefix == "" {
		s.Prefix = "s3"
	}
	return s.Prefix
}

// Do the plugin
func Do() {
	var (
		optPrefix   = flag.String("metric-key-prefix", "s3", "Metric key prefix")
		optTempfile = flag.String("tempfile", "", "Temp file name")
	)
	flag.Parse()

	var s3 S3Plugin
	s3.Prefix = *optPrefix

	helper := mp.NewMackerelPlugin(s3)
	helper.Tempfile = *optTempfile
	helper.Run()
}
