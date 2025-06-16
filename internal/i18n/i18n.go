package i18n

import (
	"fmt"
	"log"
)

// SupportedLanguage는 지원되는 언어 타입입니다
type SupportedLanguage string

const (
	Korean   SupportedLanguage = "ko"
	English  SupportedLanguage = "en"
	Japanese SupportedLanguage = "jp"
)

// I18n은 다국어 지원을 위한 구조체입니다
type I18n struct {
	currentLang SupportedLanguage
	messages    map[SupportedLanguage]map[string]string
}

// New는 새로운 I18n 인스턴스를 생성합니다
func New(lang string) *I18n {
	supportedLang := Korean // 기본값

	switch lang {
	case "en":
		supportedLang = English
	case "jp":
		supportedLang = Japanese
	case "ko":
		supportedLang = Korean
	default:
		log.Printf("Unsupported language: %s, default 'ko' will be used", lang)
	}

	i18n := &I18n{
		currentLang: supportedLang,
		messages:    make(map[SupportedLanguage]map[string]string),
	}

	i18n.loadMessages()
	return i18n
}

// GetCurrentLanguage는 현재 설정된 언어를 반환합니다
func (i *I18n) GetCurrentLanguage() string {
	return string(i.currentLang)
}

// T는 키에 해당하는 번역된 메시지를 반환합니다
func (i *I18n) T(key string, args ...interface{}) string {
	if langMessages, exists := i.messages[i.currentLang]; exists {
		if message, exists := langMessages[key]; exists {
			if len(args) > 0 {
				return fmt.Sprintf(message, args...)
			}
			return message
		}
	}

	// 현재 언어에서 찾지 못한 경우 한국어에서 찾기 (fallback)
	if i.currentLang != Korean {
		if langMessages, exists := i.messages[Korean]; exists {
			if message, exists := langMessages[key]; exists {
				if len(args) > 0 {
					return fmt.Sprintf(message, args...)
				}
				return message
			}
		}
	}

	// 키를 찾지 못한 경우 키 자체를 반환
	return key
}

// loadMessages는 모든 언어의 메시지를 로드합니다
func (i *I18n) loadMessages() {
	i.messages[Korean] = i.getKoreanMessages()
	i.messages[English] = i.getEnglishMessages()
	i.messages[Japanese] = i.getJapaneseMessages()
}
