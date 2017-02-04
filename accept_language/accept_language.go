package accept_language

import (
	"sort"
	"strconv"
	"strings"
	"webgo"
	"webgo/utils"
)

type (
	TLanguage struct {
		Language string
		Quality  float32
	}

	AcceptLanguages []TLanguage

	TAcceptLanguage struct {
	}
)

func NewAcceptLanguage() *TAcceptLanguage {
	return &TAcceptLanguage{}
}

func (self *TAcceptLanguage) NearestLang(hd *webgo.THandler) string {
	header := hd.Request.Header.Get("Accept-Language")
	if header != "" {
		acceptLanguageHeaderValues := strings.Split(header, ",")
		acceptLanguages := make(AcceptLanguages, len(acceptLanguageHeaderValues))

		for i, languageRange := range acceptLanguageHeaderValues {
			// Check if a given range is qualified or not
			if qualifiedRange := strings.Split(languageRange, ";q="); len(qualifiedRange) == 2 {
				quality, error := strconv.ParseFloat(qualifiedRange[1], 32)
				if error != nil {
					// When the quality is unparseable, assume it's 1
					acceptLanguages[i] = TLanguage{utils.Trim(qualifiedRange[0]), 1}
				} else {
					acceptLanguages[i] = TLanguage{utils.Trim(qualifiedRange[0]), float32(quality)}
				}
			} else {
				acceptLanguages[i] = TLanguage{utils.Trim(languageRange), 1}
			}
		}
		sort.Slice(acceptLanguages, func(i, j int) bool {
			return acceptLanguages[i].Quality >= acceptLanguages[j].Quality
		})
		return acceptLanguages[0].Language
	} else {
		// If we have no Accept-Language header just map an empty slice
		return ""
	}
}

func (self *TAcceptLanguage) Request(act interface{}, hd *webgo.THandler) {
}

func (self *TAcceptLanguage) Response(act interface{}, hd *webgo.THandler) {
}

func (self *TAcceptLanguage) Panic(act interface{}, hd *webgo.THandler) {
}
