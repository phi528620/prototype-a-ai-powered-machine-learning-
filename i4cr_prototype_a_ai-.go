package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type I4CRPrototypeA struct {
	ID        string                  `json:"id"`
	ModelName  string                  `json:"modelName"`
	ModelType  string                  `json:"modelType"`
	Inputs     []InputFeature          `json:"inputs"`
	Outputs    []OutputFeature         `json:"outputs"`
	Training   []TrainingDataPoint     `json:"training"`
	HyperParams HyperParameterConfig   `json:"hyperParams"`
	ModelConfig ModelConfiguration     `json:"modelConfig"`
	TrainingLog []TrainingLogEntry     `json:"trainingLog"`
	Inferences  []InferenceResult     `json:"inferences"`
}

type InputFeature struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	DataType    string `json:"dataType"`
}

type OutputFeature struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	DataType    string `json:"dataType"`
}

type TrainingDataPoint struct {
	ID          string `json:"id"`
	InputValues []float64 `json:"inputValues"`
	OutputValues []float64 `json:"outputValues"`
}

type HyperParameterConfig struct {
	LearningRate float64 `json:"learningRate"`
	BatchSize   int    `json:"batchSize"`
	EPOCHs      int    `json:"ePOCHs"`
}

type ModelConfiguration struct {
	ModelArchitecture string `json:"modelArchitecture"`
	Optimizer         string `json:"optimizer"`
	LossFunction      string `json:"lossFunction"`
}

type TrainingLogEntry struct {
	EPOCH int    `json:"ePOCH"`
	Loss  float64 `json:"loss"`
}

type InferenceResult struct {
	InputValues []float64 `json:"inputValues"`
	OutputValues []float64 `json:"outputValues"`
	PredictedOutput float64 `json:"predictedOutput"`
}

func NewI4CRPrototypeA() *I4CRPrototypeA {
	return &I4CRPrototypeA{
		ID: uuid.New().String(),
	}
}

func main() {
	i4cr := NewI4CRPrototypeA()
	i4cr.ModelName = "I4CR Prototype A"
	i4cr.ModelType = "Neural Network"

	i4cr.Inputs = append(i4cr.Inputs, InputFeature{
		Name:        "Input 1",
		Description: "Input feature 1",
		DataType:    "float64",
	})

	i4cr.Outputs = append(i4cr.Outputs, OutputFeature{
		Name:        "Output 1",
		Description: "Output feature 1",
		DataType:    "float64",
	})

	i4cr.Training = append(i4cr.Training, TrainingDataPoint{
		ID: "TrainingDataPoint_1",
		InputValues: []float64{1.0, 2.0},
		OutputValues: []float64{3.0},
	})

	i4cr.HyperParams = HyperParameterConfig{
		LearningRate: 0.01,
		BatchSize: 32,
		EPOCHs: 10,
	}

	i4cr.ModelConfig = ModelConfiguration{
		ModelArchitecture: "FeedForward",
		Optimizer: "SGD",
		LossFunction: "MeanSquaredError",
	}

	json, err := json.MarshalIndent(i4cr, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json))
}