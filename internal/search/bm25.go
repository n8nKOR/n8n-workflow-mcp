package search

import (
	"math"
)

// BM25Parameters BM25 알고리즘의 매개변수
type BM25Parameters struct {
	K1 float64 // 용어 빈도 정규화 매개변수 (일반적으로 1.2)
	B  float64 // 문서 길이 정규화 매개변수 (일반적으로 0.75)
}

// DefaultBM25Parameters 기본 BM25 매개변수
func DefaultBM25Parameters() BM25Parameters {
	return BM25Parameters{
		K1: 1.2,
		B:  0.75,
	}
}

// BM25Scorer BM25 점수 계산기
type BM25Scorer struct {
	params       BM25Parameters
	documents    [][]string
	docFreq      map[string]int
	totalDocs    int
	avgDocLength float64
	docLengths   []int
}

// NewBM25Scorer 새 BM25 점수 계산기 생성
func NewBM25Scorer(documents [][]string, params BM25Parameters) *BM25Scorer {
	scorer := &BM25Scorer{
		params:     params,
		documents:  documents,
		docFreq:    make(map[string]int),
		totalDocs:  len(documents),
		docLengths: make([]int, len(documents)),
	}

	scorer.buildIndex()
	return scorer
}

// buildIndex 인덱스 구축
func (bm25 *BM25Scorer) buildIndex() {
	totalLength := 0

	// 각 문서의 길이 계산 및 용어 빈도 수집
	for i, doc := range bm25.documents {
		bm25.docLengths[i] = len(doc)
		totalLength += len(doc)

		// 문서의 고유 용어 집합
		uniqueTerms := make(map[string]bool)
		for _, term := range doc {
			uniqueTerms[term] = true
		}

		// 각 고유 용어에 대해 문서 빈도 증가
		for term := range uniqueTerms {
			bm25.docFreq[term]++
		}
	}

	// 평균 문서 길이 계산
	if bm25.totalDocs > 0 {
		bm25.avgDocLength = float64(totalLength) / float64(bm25.totalDocs)
	}
}

// Score 주어진 쿼리에 대한 문서의 BM25 점수 계산
func (bm25 *BM25Scorer) Score(queryTerms []string, docIndex int) float64 {
	if docIndex < 0 || docIndex >= len(bm25.documents) {
		return 0.0
	}

	doc := bm25.documents[docIndex]
	docLength := float64(bm25.docLengths[docIndex])

	// 문서 내 용어 빈도 계산
	termFreq := make(map[string]int)
	for _, term := range doc {
		termFreq[term]++
	}

	score := 0.0

	for _, queryTerm := range queryTerms {
		tf := float64(termFreq[queryTerm])

		if tf > 0 {
			// IDF 계산
			df := float64(bm25.docFreq[queryTerm])
			if df == 0 {
				continue
			}

			idf := math.Log(float64(bm25.totalDocs) / df)

			// BM25 공식
			numerator := tf * (bm25.params.K1 + 1)
			denominator := tf + bm25.params.K1*(1-bm25.params.B+bm25.params.B*(docLength/bm25.avgDocLength))

			score += idf * (numerator / denominator)
		}
	}

	return score
}

// ScoreAll 모든 문서에 대한 BM25 점수 계산
func (bm25 *BM25Scorer) ScoreAll(queryTerms []string) []float64 {
	scores := make([]float64, len(bm25.documents))
	for i := range bm25.documents {
		scores[i] = bm25.Score(queryTerms, i)
	}
	return scores
}

// GetDocumentFrequency 특정 용어의 문서 빈도 반환
func (bm25 *BM25Scorer) GetDocumentFrequency(term string) int {
	return bm25.docFreq[term]
}

// GetAverageDocumentLength 평균 문서 길이 반환
func (bm25 *BM25Scorer) GetAverageDocumentLength() float64 {
	return bm25.avgDocLength
}
