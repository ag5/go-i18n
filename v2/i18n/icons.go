package i18n

import "golang.org/x/text/language"

// Apple store icon links are hosted via frontend repo: https://github.com/ag5/js-monorepo
func generateAppleStoreIconLink(lang string) string {
	return "https://cdn.ag5.com/mail-assets/mobile/appstore/" + lang + ".png"
}

var appleStoreLink = map[language.Tag]string{
	language.Arabic:               generateAppleStoreIconLink("ar"),    // ar
	language.ModernStandardArabic: generateAppleStoreIconLink("ar"),    // ar-001
	language.Czech:                generateAppleStoreIconLink("cs"),    // cs
	language.Danish:               generateAppleStoreIconLink("da"),    // da
	language.German:               generateAppleStoreIconLink("de"),    // de
	language.English:              generateAppleStoreIconLink("en"),    // en
	language.AmericanEnglish:      generateAppleStoreIconLink("en"),    // en-US
	language.BritishEnglish:       generateAppleStoreIconLink("en"),    // en-GB
	language.Spanish:              generateAppleStoreIconLink("es"),    // es
	language.EuropeanSpanish:      generateAppleStoreIconLink("es"),    // es-ES
	language.LatinAmericanSpanish: generateAppleStoreIconLink("es"),    // es-419
	language.Finnish:              generateAppleStoreIconLink("fi"),    // fi
	language.French:               generateAppleStoreIconLink("fr"),    // fr
	language.Hebrew:               generateAppleStoreIconLink("he"),    // he
	language.Indonesian:           generateAppleStoreIconLink("id"),    // id
	language.Italian:              generateAppleStoreIconLink("it"),    // it
	language.Japanese:             generateAppleStoreIconLink("ja"),    // ja
	language.Korean:               generateAppleStoreIconLink("ko"),    // ko
	language.Dutch:                generateAppleStoreIconLink("nl"),    // nl
	language.Norwegian:            generateAppleStoreIconLink("no"),    // no
	language.Polish:               generateAppleStoreIconLink("pl"),    // pl
	language.Portuguese:           generateAppleStoreIconLink("pt"),    // pt
	language.BrazilianPortuguese:  generateAppleStoreIconLink("pt-br"), // pt-BR
	language.EuropeanPortuguese:   generateAppleStoreIconLink("pt"),    // pt-PT
	language.Russian:              generateAppleStoreIconLink("ru"),    // ru
	language.Swedish:              generateAppleStoreIconLink("se"),    // sv
	language.Thai:                 generateAppleStoreIconLink("th"),    // th
	language.Turkish:              generateAppleStoreIconLink("tr"),    // tr
	language.Ukrainian:            generateAppleStoreIconLink("uk"),    // uk
	language.Vietnamese:           generateAppleStoreIconLink("vi"),    // vi
	language.Chinese:              generateAppleStoreIconLink("zh"),    // zh
	language.SimplifiedChinese:    generateAppleStoreIconLink("zh"),    // zh-Hans
	language.TraditionalChinese:   generateAppleStoreIconLink("zh"),    // zh-Hant
}

func (l *Localizer) AppleStoreIconLink() string {
	if len(l.tags) <= 0 {
		return generateAppleStoreIconLink("en")
	}

	for _, tag := range l.tags {
		value, exists := appleStoreLink[tag]
		if exists {
			return value
		}
	}

	return generateAppleStoreIconLink("en")
}

// Google Play Store icon links are hosted via frontend repo: https://github.com/ag5/js-monorepo
func generateGooglePlayStoreIconLink(lang string) string {
	return "https://cdn.ag5.com/mail-assets/mobile/googleplay/" + lang + ".png"
}

var googlePlayStoreLink = map[language.Tag]string{
	language.Arabic:               generateGooglePlayStoreIconLink("ar"),    // ar
	language.ModernStandardArabic: generateGooglePlayStoreIconLink("ar"),    // ar-001
	language.Czech:                generateGooglePlayStoreIconLink("cs"),    // cs
	language.Danish:               generateGooglePlayStoreIconLink("da"),    // da
	language.German:               generateGooglePlayStoreIconLink("de"),    // de
	language.English:              generateGooglePlayStoreIconLink("en"),    // en
	language.AmericanEnglish:      generateGooglePlayStoreIconLink("en"),    // en-US
	language.BritishEnglish:       generateGooglePlayStoreIconLink("en"),    // en-GB
	language.Spanish:              generateGooglePlayStoreIconLink("es"),    // es
	language.EuropeanSpanish:      generateGooglePlayStoreIconLink("es"),    // es-ES
	language.LatinAmericanSpanish: generateGooglePlayStoreIconLink("es"),    // es-419
	language.Finnish:              generateGooglePlayStoreIconLink("fi"),    // fi
	language.French:               generateGooglePlayStoreIconLink("fr"),    // fr
	language.Hebrew:               generateGooglePlayStoreIconLink("he"),    // he
	language.Indonesian:           generateGooglePlayStoreIconLink("id"),    // id
	language.Italian:              generateGooglePlayStoreIconLink("it"),    // it
	language.Japanese:             generateGooglePlayStoreIconLink("ja"),    // ja
	language.Korean:               generateGooglePlayStoreIconLink("ko"),    // ko
	language.Dutch:                generateGooglePlayStoreIconLink("nl"),    // nl
	language.Norwegian:            generateGooglePlayStoreIconLink("no"),    // no
	language.Polish:               generateGooglePlayStoreIconLink("pl"),    // pl
	language.Portuguese:           generateGooglePlayStoreIconLink("pt"),    // pt
	language.BrazilianPortuguese:  generateGooglePlayStoreIconLink("pt-br"), // pt-BR
	language.EuropeanPortuguese:   generateGooglePlayStoreIconLink("pt"),    // pt-PT
	language.Russian:              generateGooglePlayStoreIconLink("ru"),    // ru
	language.Swedish:              generateGooglePlayStoreIconLink("sv"),    // sv
	language.Thai:                 generateGooglePlayStoreIconLink("th"),    // th
	language.Turkish:              generateGooglePlayStoreIconLink("tr"),    // tr
	language.Ukrainian:            generateGooglePlayStoreIconLink("uk"),    // uk
	language.Vietnamese:           generateGooglePlayStoreIconLink("vi"),    // vi
	language.Chinese:              generateGooglePlayStoreIconLink("zh"),    // zh
	language.SimplifiedChinese:    generateGooglePlayStoreIconLink("zh"),    // zh-Hans
	language.TraditionalChinese:   generateGooglePlayStoreIconLink("zh"),    // zh-Hant
}

func (l *Localizer) GooglePlayStoreIconLink() string {
	if len(l.tags) <= 0 {
		return generateGooglePlayStoreIconLink("en")
	}

	for _, tag := range l.tags {
		value, exists := googlePlayStoreLink[tag]
		if exists {
			return value
		}
	}

	return generateGooglePlayStoreIconLink("en")
}
