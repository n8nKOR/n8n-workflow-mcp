package config

import (
	"flag"
	"os"
	"strconv"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/search"
)

// Config holds application configuration
type Config struct {
	Name    string
	Version string

	// 검색 관련 설정
	SearchWeights search.SearchWeights
	DataPath      string
}

// New creates a new configuration instance
func New() *Config {
	cfg := &Config{
		Name:    "n8n-workflow-mcp-server",
		Version: "0.1.0",
	}

	// 검색 관련 설정 초기화
	cfg.initializeSearchConfig()

	return cfg
}

// initializeSearchConfig 검색 관련 설정을 초기화합니다
func (c *Config) initializeSearchConfig() {
	// 기본 가중치 설정
	c.SearchWeights = search.DefaultSearchWeights()

	// 기본 데이터 경로 설정
	c.DataPath = "internal/data/nodes"

	// 환경 변수에서 가중치 설정 읽기
	if nodeNameWeight := os.Getenv("SEARCH_WEIGHT_NODE_NAME"); nodeNameWeight != "" {
		if weight, err := strconv.ParseFloat(nodeNameWeight, 64); err == nil {
			c.SearchWeights.NodeName = weight
		}
	}

	if overviewWeight := os.Getenv("SEARCH_WEIGHT_OVERVIEW"); overviewWeight != "" {
		if weight, err := strconv.ParseFloat(overviewWeight, 64); err == nil {
			c.SearchWeights.Overview = weight
		}
	}

	if useCasesWeight := os.Getenv("SEARCH_WEIGHT_USE_CASES"); useCasesWeight != "" {
		if weight, err := strconv.ParseFloat(useCasesWeight, 64); err == nil {
			c.SearchWeights.UseCases = weight
		}
	}

	if operationWeight := os.Getenv("SEARCH_WEIGHT_OPERATION"); operationWeight != "" {
		if weight, err := strconv.ParseFloat(operationWeight, 64); err == nil {
			c.SearchWeights.Operation = weight
		}
	}

	// 환경 변수에서 데이터 경로 설정 읽기
	if dataPath := os.Getenv("N8N_NODES_DATA_PATH"); dataPath != "" {
		c.DataPath = dataPath
	}
}

// ApplyFlags applies command line flags to the configuration
func (c *Config) ApplyFlags() {
	// 명령행 인자로 가중치 설정 (환경 변수보다 우선)
	nodeNameWeight := flag.Float64("search-weight-node-name", c.SearchWeights.NodeName, "노드명에 대한 검색 가중치")
	overviewWeight := flag.Float64("search-weight-overview", c.SearchWeights.Overview, "개요에 대한 검색 가중치")
	useCasesWeight := flag.Float64("search-weight-use-cases", c.SearchWeights.UseCases, "사용 사례에 대한 검색 가중치")
	operationWeight := flag.Float64("search-weight-operation", c.SearchWeights.Operation, "오퍼레이션에 대한 검색 가중치")
	dataPath := flag.String("data-path", c.DataPath, "n8n 노드 데이터 디렉토리 경로")

	// flag.Parse()는 main에서 호출되므로 여기서는 호출하지 않음

	// 파싱된 값들을 설정에 적용 (flag.Parse() 이후에 호출되어야 함)
	c.SearchWeights.NodeName = *nodeNameWeight
	c.SearchWeights.Overview = *overviewWeight
	c.SearchWeights.UseCases = *useCasesWeight
	c.SearchWeights.Operation = *operationWeight
	c.DataPath = *dataPath
}
