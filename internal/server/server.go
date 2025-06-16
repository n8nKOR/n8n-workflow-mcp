package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mark3labs/mcp-go/server"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/config"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/handlers"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/i18n"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/search"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/services"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/validator"
	"github.com/n8nKOR/n8n-workflow-mcp/internal/workflow"
)

// App represents the main application
type App struct {
	Config        *config.Config
	Server        *server.MCPServer
	Handler       *handlers.Handler
	SearchService *services.SearchService
	I18n          *i18n.I18n
}

// NewStandalone creates a new standalone application instance with embedded data
func NewStandalone() *App {
	// 설정 초기화
	cfg := config.New()

	// i18n 초기화 (기본 언어: 한국어)
	i18nInstance := i18n.New("ko")

	// 우아한 종료 처리
	setupGracefulShutdown()

	// 서비스 초기화 (파일 기반)
	n8nValidator := validator.NewN8nValidator(true, 50) // 줄 번호 포함, 최대 50개 오류
	workflowService := workflow.NewService(n8nValidator)

	// 핸들러 초기화
	handler := handlers.NewHandler(workflowService)

	// 검색 서비스 초기화 (임베딩된 데이터 사용)
	searchService := services.NewSearchService()

	// 임베딩된 데이터로 검색 서비스 초기화
	if err := searchService.InitializeEmbedded(cfg.SearchWeights); err != nil {
		log.Printf("검색 기능은 사용할 수 없습니다.")
	}

	// MCP 서버 생성 (Claude Desktop 호환성 옵션 추가)
	mcpServer := server.NewMCPServer(
		cfg.Name,
		cfg.Version,
		server.WithToolCapabilities(false),
	)

	return &App{
		Config:        cfg,
		Server:        mcpServer,
		Handler:       handler,
		SearchService: searchService,
		I18n:          i18nInstance,
	}
}

// NewStandaloneWithFlags creates a new standalone application instance with custom search weights and language
func NewStandaloneWithFlags(nodeNameWeight, overviewWeight, useCasesWeight, operationWeight float64, lang string) *App {
	// 설정 초기화
	cfg := config.New()

	// 명령행 인자로 전달받은 값들로 설정 업데이트
	cfg.SearchWeights = search.SearchWeights{
		NodeName:  nodeNameWeight,
		Overview:  overviewWeight,
		UseCases:  useCasesWeight,
		Operation: operationWeight,
	}

	// i18n 초기화
	i18nInstance := i18n.New(lang)

	// 우아한 종료 처리
	setupGracefulShutdown()

	// 서비스 초기화 (파일 기반)
	n8nValidator := validator.NewN8nValidator(true, 50) // 줄 번호 포함, 최대 50개 오류
	workflowService := workflow.NewService(n8nValidator)

	// 핸들러 초기화
	handler := handlers.NewHandler(workflowService)

	// 검색 서비스 초기화 (임베딩된 데이터 사용)
	searchService := services.NewSearchService()

	// 임베딩된 데이터로 검색 서비스 초기화
	if err := searchService.InitializeEmbedded(cfg.SearchWeights); err != nil {
		log.Printf("임베딩된 검색 서비스 초기화 실패: %v", err)
		log.Printf("검색 기능은 사용할 수 없습니다.")
	}

	// MCP 서버 생성 (Claude Desktop 호환성 옵션 추가)
	mcpServer := server.NewMCPServer(
		cfg.Name,
		cfg.Version,
		server.WithToolCapabilities(false),
	)

	return &App{
		Config:        cfg,
		Server:        mcpServer,
		Handler:       handler,
		SearchService: searchService,
		I18n:          i18nInstance,
	}
}

// New creates a new application instance (기존 파일 시스템 방식)
func New() *App {
	// 설정 초기화
	cfg := config.New()

	// i18n 초기화 (기본 언어: 한국어)
	i18nInstance := i18n.New("ko")

	// 우아한 종료 처리
	setupGracefulShutdown()

	// 서비스 초기화 (파일 기반)
	n8nValidator := validator.NewN8nValidator(true, 50) // 줄 번호 포함, 최대 50개 오류
	workflowService := workflow.NewService(n8nValidator)

	// 핸들러 초기화
	handler := handlers.NewHandler(workflowService)

	// 검색 서비스 초기화 (config의 설정 사용)
	searchService := services.NewSearchService()

	// 검색 서비스에 문서 로드 (config에서 가져온 경로와 가중치 사용)
	if err := searchService.InitializeWithWeights(cfg.SearchWeights, cfg.DataPath); err != nil {
		log.Printf("검색 서비스 초기화 실패: %v", err)
		log.Printf("검색 기능은 사용할 수 없습니다. 경로를 확인하세요: %s", cfg.DataPath)
	}

	// MCP 서버 생성 (Claude Desktop 호환성 옵션 추가)
	mcpServer := server.NewMCPServer(
		cfg.Name,
		cfg.Version,
		server.WithToolCapabilities(false),
	)

	return &App{
		Config:        cfg,
		Server:        mcpServer,
		Handler:       handler,
		SearchService: searchService,
		I18n:          i18nInstance,
	}
}

// NewWithFlags creates a new application instance with custom search weights, data path, and language (기존 파일 시스템 방식)
func NewWithFlags(nodeNameWeight, overviewWeight, useCasesWeight, operationWeight float64, dataPath, lang string) *App {
	// 설정 초기화
	cfg := config.New()

	// 명령행 인자로 전달받은 값들로 설정 업데이트
	cfg.SearchWeights = search.SearchWeights{
		NodeName:  nodeNameWeight,
		Overview:  overviewWeight,
		UseCases:  useCasesWeight,
		Operation: operationWeight,
	}
	cfg.DataPath = dataPath

	// i18n 초기화
	i18nInstance := i18n.New(lang)

	// 우아한 종료 처리
	setupGracefulShutdown()

	// 서비스 초기화 (파일 기반)
	n8nValidator := validator.NewN8nValidator(true, 50) // 줄 번호 포함, 최대 50개 오류
	workflowService := workflow.NewService(n8nValidator)

	// 핸들러 초기화
	handler := handlers.NewHandler(workflowService)

	// 검색 서비스 초기화 (config의 설정 사용)
	searchService := services.NewSearchService()

	// 검색 서비스에 문서 로드 (config에서 가져온 경로와 가중치 사용)
	if err := searchService.InitializeWithWeights(cfg.SearchWeights, cfg.DataPath); err != nil {
		log.Printf("검색 서비스 초기화 실패: %v", err)
		log.Printf("검색 기능은 사용할 수 없습니다. 경로를 확인하세요: %s", cfg.DataPath)
	}

	// MCP 서버 생성 (Claude Desktop 호환성 옵션 추가)
	mcpServer := server.NewMCPServer(
		cfg.Name,
		cfg.Version,
		server.WithToolCapabilities(false),
	)

	return &App{
		Config:        cfg,
		Server:        mcpServer,
		Handler:       handler,
		SearchService: searchService,
		I18n:          i18nInstance,
	}
}

// Start starts the application
func (a *App) Start() error {
	// 도구들 설정
	if err := SetupTools(a.Server, a.Handler, a.SearchService, a.I18n); err != nil {
		log.Fatalf("도구 설정 실패: %v", err)
	}

	// stdio 서버 시작
	if err := server.ServeStdio(a.Server); err != nil {
		log.Fatalf("서버 오류: %v", err)
	}

	return nil
}

// setupGracefulShutdown sets up graceful shutdown handling
func setupGracefulShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		if os.Getenv("MCP_QUIET") != "true" {
			log.Println("[INFO] 종료 신호를 받았습니다. 서버를 중지합니다...")
		}
		os.Exit(0)
	}()
}
