package handlers

// 헬퍼 함수들
func generateBasicMermaid() string {
	return `flowchart TD
    A[Start] --> B[HTTP Request]
    B --> C[Process Data]
    C --> D[Code]
    D --> E[End]
    
    classDef startEnd fill:#e1f5fe
    classDef process fill:#f3e5f5
    class A,E startEnd
    class B,C,D process`
}

func generateBasicN8nJSON() string {
	// 기본 n8n 워크플로우 템플릿
	return `{
  "name": "Generated Workflow",
  "nodes": [
    {
      "parameters": {},
      "id": "start-node",
      "name": "Start",
      "type": "n8n-nodes-base.start",
      "typeVersion": 1,
      "position": [250, 300]
    },
    {
      "parameters": {
        "url": "https://api.example.com/data",
        "options": {}
      },
      "id": "http-node",
      "name": "HTTP Request",
      "type": "n8n-nodes-base.httpRequest",
      "typeVersion": 1,
      "position": [450, 300]
    }
  ],
  "connections": {
    "Start": {
      "main": [
        [
          {
            "node": "HTTP Request",
            "type": "main",
            "index": 0
          }
        ]
      ]
    }
  }
}`
}
