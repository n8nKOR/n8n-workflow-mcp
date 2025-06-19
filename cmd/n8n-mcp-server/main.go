package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/n8nKOR/n8n-workflow-mcp/internal/server"
)

var (
	version = "0.2.3"
	name    = "n8n-workflow-validator"
)

func main() {
	var (
		showVersion    = flag.Bool("version", false, "Show version and exit")
		standalone     = flag.Bool("standalone", true, "임베딩된 데이터를 사용하는 standalone 모드로 실행 (외부 파일 불필요)")
		nameWeight     = flag.Float64("name-weight", 3.0, "노드명에 대한 검색 가중치")
		overviewWeight = flag.Float64("overview-weight", 2.0, "개요에 대한 검색 가중치")
		caseWeight     = flag.Float64("case-weight", 1.5, "사용 사례에 대한 검색 가중치")
		opWeight       = flag.Float64("op-weight", 1.0, "오퍼레이션에 대한 검색 가중치")
		dataPath       = flag.String("data-path", "internal/data/nodes", "n8n 노드 데이터 디렉토리 경로 (standalone 모드에서는 무시됨)")
		lang           = flag.String("lang", "ko", "언어 설정 (ko, en, jp)")
	)

	// flag 파싱
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s version %s\n", name, version)
		os.Exit(0)
	}

	var app *server.App

	// standalone 모드 여부에 따라 다른 초기화 방식 사용
	if *standalone {
		app = server.NewStandaloneWithFlags(*nameWeight, *overviewWeight, *caseWeight, *opWeight, *lang)
	} else {
		app = server.NewWithFlags(*nameWeight, *overviewWeight, *caseWeight, *opWeight, *dataPath, *lang)
	}

	// 애플리케이션 시작
	if err := app.Start(); err != nil {
		os.Exit(1)
	}
}
