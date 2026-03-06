package tiktokads

import (
	"log"
	"strings"
	"testing"
)

func TestPreviewAd(t *testing.T) {
	initTestSession()
	catalogId := "7595212727484827393"
	videoId := "v10033g50000d6j92l7og65qb0aa5jhg"
	imageId := "ad-site-i18n-sg/202603045d0dcca471a9e75240b7945c"
	identityId := "7483102948432887825"

	ads := []*Ad{
		// test from existing ad
		{Id: "1858747646691602"},
		// test from new ad
		NewCatalogAdWithCustomVideo("Custom video & link", identityId, catalogId, "https://www.google.com", videoId, imageId),
	}

	for _, ad := range ads {
		ad.SetAdText("Ad Preview")
		preview, e := PreviewAd(getTestAccount(), ad)

		if nil != e {
			log.Println(LastRequestPayload)
			log.Println(LastResponse)
			t.Error(e)
		} else if !strings.HasPrefix(preview.Iframe, "<iframe ") {
			t.Error("Iframe should start with '<iframe'")
		}
	}
}
