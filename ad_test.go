package tiktokads

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateCatalogAd(t *testing.T) {
	initTestSession()
	businessId := "7489850869144485895"
	//pixelId := "7595236536677974034"
	catalogId := "7595212727484827393"
	adGroupId := "1858747631696034"
	identityId := "7483102948432887825"
	//videoId := "v10033g50000d6j92l7og65qb0aa5jhg"
	//imageId := "ad-site-i18n-sg/202603045d0dcca471a9e75240b7945c"
	//lpUrl := "https://jolivetboisnegoce.fr"

	sets, err := GetCatalogSets(businessId, catalogId)
	assertNotEmptyResult(t, sets, err)

	ad := newShoppingAds(fmt.Sprintf("Catalog Ads %s", time.Now().Format(time.DateTime)), identityId, catalogId)
	ad.SetAdText("Best offers")
	ad.WithGeneratedCatalogVideo()
	ad.SetProductSetId(sets[0].Id)
	adCreated, err := UpdateAd(getTestAccount(), adGroupId, ad)
	if err != nil {
		t.Fatal(err)
	} else if adCreated.ProductSetId != sets[0].Id {
		t.Error("ad created product set id not equal to adCreated.ProductSetId")
	}

	// remove product set
	adCreated.SetProductSetId("")
	adCreated.Name = fmt.Sprintf("Ad Updated %s", time.Now().Format(time.DateTime))
	adCreated.SetEnabled(false)

	adUpdated, err := UpdateAd(getTestAccount(), adGroupId, adCreated)
	if err != nil {
		t.Fatal(err)
	} else if adUpdated.Id != adCreated.Id {
		t.Error("ad updated id not equal to adCreated.Id")
	}
}

func TestCreateCustomAd(t *testing.T) {
	initTestSession()
	businessId := "7489850869144485895"
	//pixelId := "7595236536677974034"
	catalogId := "7595212727484827393"
	adGroupId := "1858747631696034"
	identityId := "7483102948432887825"
	videoId := "v10033g50000d6j92l7og65qb0aa5jhg"
	imageId := "ad-site-i18n-sg/202603045d0dcca471a9e75240b7945c"
	lpUrl := "https://jolivetboisnegoce.fr"

	sets, err := GetCatalogSets(businessId, catalogId)
	assertNotEmptyResult(t, sets, err)
	ad := NewCatalogAdWithCustomVideo("Custom video", identityId, catalogId, lpUrl, videoId, imageId)
	ad.SetAdText("Best offers")
	ad.SetAiGeneratedContent(true)
	adCreated, err := UpdateAd(getTestAccount(), adGroupId, ad)
	if err != nil {
		t.Fatal(err)
	}
	adCreated.Name = fmt.Sprintf("Ad Updated %s", time.Now().Format(time.DateTime))
	adCreated.SetEnabled(false)

	videoId = "v10033g50000cvckupvog65h50275780"
	imageId = "ad-site-i18n-sg/202603055d0d2cec278078f641d886d8"
	lpUrl = "https://jolivetboisnegoce.fr?utm_source=tiktok"
	adCreated.WithCustomUrlAndAssets(lpUrl, videoId, imageId)
	adCreated.Name = fmt.Sprintf("Custom video updated %s", time.Now().Format(time.DateTime))
	//adCreated.SetAiGeneratedContent(true)

	adUpdated, err := UpdateAd(getTestAccount(), adGroupId, adCreated)
	if err != nil {
		t.Fatal(err)
	} else if adUpdated.Id != adCreated.Id {
		t.Error("ad updated id not equal to adCreated.Id")
	}

	if adUpdated.VideoId != adCreated.VideoId {
		t.Error("ad updated video id not equal to adCreated.VideoId")
	}

}
