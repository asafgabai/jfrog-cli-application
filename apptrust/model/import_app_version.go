package model

type ImportAppVersionOptions struct {
	Mode         string                    `json:"mode,omitempty"`
	PathMappings []DistributionPathMapping `json:"path_mappings,omitempty"`
}
