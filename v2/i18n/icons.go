package i18n

import "golang.org/x/text/language"

var storeLinkLanguageMapper = map[language.Tag]string{
	language.Arabic:               "ar",    // ar
	language.ModernStandardArabic: "ar",    // ar-001
	language.Czech:                "cs",    // cs
	language.Danish:               "da",    // da
	language.German:               "de",    // de
	language.English:              "en",    // en
	language.AmericanEnglish:      "en",    // en-US
	language.BritishEnglish:       "en",    // en-GB
	language.Spanish:              "es",    // es
	language.EuropeanSpanish:      "es",    // es-ES
	language.LatinAmericanSpanish: "es",    // es-419
	language.Finnish:              "fi",    // fi
	language.French:               "fr",    // fr
	language.Hebrew:               "he",    // he
	language.Indonesian:           "id",    // id
	language.Italian:              "it",    // it
	language.Japanese:             "ja",    // ja
	language.Korean:               "ko",    // ko
	language.Dutch:                "nl",    // nl
	language.Norwegian:            "no",    // no
	language.Polish:               "pl",    // pl
	language.Portuguese:           "pt",    // pt
	language.BrazilianPortuguese:  "pt-br", // pt-BR
	language.EuropeanPortuguese:   "pt",    // pt-PT
	language.Russian:              "ru",    // ru
	language.Swedish:              "sv",    // sv
	language.Thai:                 "th",    // th
	language.Turkish:              "tr",    // tr
	language.Ukrainian:            "uk",    // uk
	language.Vietnamese:           "vi",    // vi
	language.Chinese:              "zh",    // zh
	language.SimplifiedChinese:    "zh",    // zh-Hans
	language.TraditionalChinese:   "zh",    // zh-Hant
}

// Apple store icon links are hosted via frontend repo: https://github.com/ag5/js-monorepo
func generateAppleStoreIconLink(lang string) string {
	return "https://cdn.ag5.com/mail-assets/mobile/appstore/" + lang + ".png"
}

func (l *Localizer) AppleStoreIconLink() string {
	if len(l.tags) <= 0 {
		return generateAppleStoreIconLink("en")
	}

	for _, tag := range l.tags {
		value, exists := storeLinkLanguageMapper[tag]
		if exists {
			return generateAppleStoreIconLink(value)
		}
	}

	return generateAppleStoreIconLink("en")
}

// Google Play Store icon links are hosted via frontend repo: https://github.com/ag5/js-monorepo
func generateGooglePlayStoreIconLink(lang string) string {
	return "https://cdn.ag5.com/mail-assets/mobile/googleplay/" + lang + ".png"
}

func (l *Localizer) GooglePlayStoreIconLink() string {
	if len(l.tags) <= 0 {
		return generateGooglePlayStoreIconLink("en")
	}

	for _, tag := range l.tags {
		value, exists := storeLinkLanguageMapper[tag]
		if exists {
			return generateGooglePlayStoreIconLink(value)
		}
	}

	return generateGooglePlayStoreIconLink("en")
}
