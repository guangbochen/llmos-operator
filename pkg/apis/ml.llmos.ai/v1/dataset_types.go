package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/llmos-ai/llmos-operator/pkg/apis/common"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Model",type="string",JSONPath=`.spec.model`
// +kubebuilder:printcolumn:name="Ready",type="integer",JSONPath=".status.readyReplicas"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.state"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// Dataset represents the main structure of an LLM dataset.
type Dataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatasetSpec   `json:"spec,omitempty"`
	Status DatasetStatus `json:"status,omitempty"`
}

type DatasetSpec struct {
	// Format specifies the file format of the dataset (e.g., "json", "jsonl", "csv", "parquet").
	Format string `json:"format"`
	// SourceURL is the URL where the dataset is sourced from.
	SourceURL string `json:"sourceUrl"`
	// DestinationURL is the URL where the dataset is stored or hosted.
	DestinationURL string `json:"destinationUrl"`
	// Tags are keywords or labels associated with the dataset.
	Tags []string `json:"tags,omitempty"`
	// Features describes the structure of the dataset (e.g., columns and their types),
	// ensuring compatibility with various data formats. This is particularly useful for validating and processing LLM datasets,
	// which often have structured data like instructions and outputs.
	Features []Feature `json:"features,omitempty"`
	// FineTuningTechniques lists the supported fine-tuning techniques (e.g., "LoRA", "QLoRA", "GRPO").
	FineTuningTechniques []string `json:"fine_tuning_techniques"`
	// Preprocessing is an optional configuration for data preprocessing.
	Preprocessing *PreprocessingConfig `json:"preprocessing,omitempty"`
}

// Feature represents a single feature (column) in the dataset.
type Feature struct {
	// Name is the name of the feature (column).
	Name string `json:"name"`
	// Type is the data type of the feature (e.g., "string", "int64", "float64", "array").
	Type string `json:"type"`
}

// PreprocessingConfig holds optional configurations for data preprocessing.
type PreprocessingConfig struct {
	// TokenizerModel specifies the tokenizer model used (e.g., "openapi/tiktoken", "bert-base-uncased").
	TokenizerModel string `json:"tokenizer_model"`
	// MaxLength is the maximum length for tokenization (e.g., 512).
	MaxLength int `json:"max_length"`
	// FilteringRules is a list of rules for filtering data (e.g., ["remove_duplicates", "remove_empty"]).
	FilteringRules []string `json:"filtering_rules"`
	// ChunkingMethod specifies the method for chunking data (e.g., "fixed_size", "by_sentence").
	ChunkingMethod string `json:"chunking_method"`
}

type DatasetStatus struct {
	// Version is the version number of the dataset to track updates.
	Version string `json:"version"`
	// Creator is the name or identifier of the dataset creator.
	Creator string `json:"creator,omitempty"`
	// Conditions is a list of conditions associated with the dataset.
	Conditions []common.Condition `json:"conditions"`
	// State is the current state of the dataset (e.g., "ready", "error").
	State string `json:"state"`
}
