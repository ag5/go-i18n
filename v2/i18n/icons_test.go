package i18n

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"testing"
)

type iconResult struct {
	tag                 language.Tag
	appleStoreLink      string
	googlePlayStoreLink string
}

var expectedResults = []iconResult{
	{
		tag:                 language.Arabic, // ar
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/ar.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/ar.png",
	},
	{
		tag:                 language.ModernStandardArabic, // ar-001
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/ar.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/ar.png",
	},
	{
		tag:                 language.Czech, // cs
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/cs.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/cs.png",
	},
	{
		tag:                 language.Danish, // da
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/da.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/da.png",
	},
	{
		tag:                 language.German, // de
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/de.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/de.png",
	},
	{
		tag:                 language.English, // en
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/en.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/en.png",
	},
	{
		tag:                 language.AmericanEnglish, // en-US
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/en.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/en.png",
	},
	{
		tag:                 language.BritishEnglish, // en-GB
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/en.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/en.png",
	},
	{
		tag:                 language.Spanish, // es
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/es.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/es.png",
	},
	{
		tag:                 language.EuropeanSpanish, // es-ES
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/es.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/es.png",
	},
	{
		tag:                 language.LatinAmericanSpanish, // es-419
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/es.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/es.png",
	},
	{
		tag:                 language.Finnish, // fi
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/fi.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/fi.png",
	},
	{
		tag:                 language.French, // fr
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/fr.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/fr.png",
	},
	{
		tag:                 language.Hebrew, // he
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/he.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/he.png",
	},
	{
		tag:                 language.Indonesian, // id
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/id.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/id.png",
	},
	{
		tag:                 language.Italian, // it
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/it.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/it.png",
	},
	{
		tag:                 language.Japanese, // ja
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/ja.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/ja.png",
	},
	{
		tag:                 language.Korean, // ko
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/ko.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/ko.png",
	},
	{
		tag:                 language.Dutch, // nl
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/nl.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/nl.png",
	},
	{
		tag:                 language.Norwegian, // no
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/no.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/no.png",
	},
	{
		tag:                 language.Polish, // pl
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/pl.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/pl.png",
	},
	{
		tag:                 language.Portuguese, // pt
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/pt.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/pt.png",
	},
	{
		tag:                 language.BrazilianPortuguese, // pt-BR
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/pt-br.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/pt-br.png",
	},
	{
		tag:                 language.EuropeanPortuguese, // pt-PT
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/pt.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/pt.png",
	},
	{
		tag:                 language.Russian, // ru
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/ru.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/ru.png",
	},
	{
		tag:                 language.Swedish, // sv
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/se.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/sv.png",
	},
	{
		tag:                 language.Thai, // th
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/th.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/th.png",
	},
	{
		tag:                 language.Turkish, // tr
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/tr.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/tr.png",
	},
	{
		tag:                 language.Ukrainian, // uk
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/uk.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/uk.png",
	},
	{
		tag:                 language.Vietnamese, // vi
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/vi.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/vi.png",
	},
	{
		tag:                 language.Chinese, // zh
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/zh.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/zh.png",
	},
	{
		tag:                 language.SimplifiedChinese, // zh-Hans
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/zh.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/zh.png",
	},
	{
		tag:                 language.TraditionalChinese, // zh-Hant
		appleStoreLink:      "https://cdn.ag5.com/mail-assets/mobile/appstore/zh.png",
		googlePlayStoreLink: "https://cdn.ag5.com/mail-assets/mobile/googleplay/zh.png",
	},
}

func TestStoreIconLinks(t *testing.T) {
	for _, iconRes := range expectedResults {
		bundle := NewBundle(iconRes.tag)
		loc := NewLocalizer(bundle, iconRes.tag.String())
		assert.Equal(t, loc.AppleStoreIconLink(), iconRes.appleStoreLink, iconRes.tag)
		assert.Equal(t, loc.GooglePlayStoreIconLink(), iconRes.googlePlayStoreLink, iconRes.tag)
	}
}
