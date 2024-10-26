package model

import (
	"fmt"
)

// PredictRequest represents the prediction request data
type PredictRequest struct {
	Features []float64 `json:"features"`
}

// PredictResponse represents the prediction response data
type PredictResponse struct {
	Prediction float64 `json:"prediction"`
}

// Model interface defines the methods that any ML model must implement
type Model interface {
	Predict(features []float64) (float64, error)
	Load(modelPath string) error
}

// SimpleModel is a basic implementation of the Model interface
// This is a placeholder that we'll replace with real ML model loading later
type SimpleModel struct{}

// Predict makes a prediction using the model
func (m *SimpleModel) Predict(features []float64) (float64, error) {
	if len(features) == 0 {
		return 0, fmt.Errorf("empty feature set")
	}

	// This is a dummy implementation - just averages the features
	// We'll replace this with real model inference later
	sum := 0.0
	for _, f := range features {
		sum += f
	}
	return sum / float64(len(features)), nil
}

// Load loads the model from a file
func (m *SimpleModel) Load(modelPath string) error {
	// This is a placeholder - we'll implement real model loading later
	return nil
}

// NewModel creates a new instance of the model
func NewModel() Model {
	return &SimpleModel{}
}
