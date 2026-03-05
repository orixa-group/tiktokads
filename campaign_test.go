package tiktokads

import (
	"cmp"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func getTestAccount() string {
	return cmp.Or(os.Getenv("TIKTOKADS_TEST_ACCOUNT"), "7101293445955731458")
}

func TestGetCampaigns(t *testing.T) {
	initTestSession()

	campaigns, err := GetCampaigns(getTestAccount())
	assertNotEmptyResult(t, campaigns, err)

	for _, campaign := range campaigns {
		if id, _ := strconv.ParseInt(campaign.Id, 10, 64); 0 == id {
			t.Error("campaign id should be numeric")
		}
		if campaign.Budget <= 0 {
			t.Error("campaign budget should be positive")
		}
	}
}

func TestUpdateShoppingCampaign(t *testing.T) {
	initTestSession()

	businessId := "7489850869144485895"
	pixelId := "7595236536677974034"
	catalogId := "7595212727484827393"

	identityId := ""
	identities, err := GetAccountIdentities(getTestAccount())
	assertNotEmptyResult(t, identities, err)
	for _, identity := range identities {
		if identity.IsAvailable() && identity.IsCustomizedUser() {
			identityId = identity.Id
		}
	}

	if len(identityId) == 0 {
		t.Fatal("identity id should not be empty")
	}

	campaign := NewShoppingCatalogCampaign(fmt.Sprintf("[TestUnit] Test Shopping Campaign %s", time.Now().Format(time.DateTime)))
	campaign.SetBudget(12345)
	campaign.SetDynamicBudget(true)

	campaignCreated, err := UpdateCampaign(getTestAccount(), campaign)

	if nil != err {
		t.Fatal(err)
	} else if len(campaignCreated.Id) == 0 {
		t.Fatal("campaign id not found after creation")
	}

	campaignCreated.SetEnabled(!campaignCreated.GetEnabled())
	campaignCreated.SetBudget(campaignCreated.GetBudget() * 2)
	campaignCreated.SetName(fmt.Sprintf("[TestUnit] Test Shopping Campaign Updated %s", time.Now().Format(time.DateTime)))

	updated, err := UpdateCampaign(getTestAccount(), campaignCreated)
	if nil != err {
		t.Fatal(err)
	} else if !reflect.DeepEqual(updated, campaignCreated) {
		t.Error("updated campaign not equal to campaignCreated campaign")
	}

	// Create adgroup
	adGroup := NewShoppingAdGroup("Default adgroup", businessId, pixelId, catalogId)
	adGroup.LocationIds = []string{"3012874", "3023519"}
	adGroup.AgeGroups = []AdGroupAgeGroup{AdGroupAgeGroup_25_34}

	adGroupCreated, err := UpdateAdGroup(getTestAccount(), campaignCreated.Id, adGroup)
	if nil != err {
		t.Error(cmp.Or(err, DeleteCampaign(getTestAccount(), campaignCreated.Id)))
	} else if len(adGroupCreated.Id) == 0 {
		t.Fatal("no id found after adgroup creation")
	}

	ad := NewCatalogAd("Default Ad", identityId, catalogId)
	ad.SetAdText("Best products")
	adCreated, err := UpdateAd(getTestAccount(), adGroupCreated.Id, ad)
	if nil != err {
		t.Error(err)
	} else if len(adCreated.Id) == 0 {
		t.Error("no id found after ad creation")
	}

	adCreated.SetEnabled(false)
	if adUpdated, e := UpdateAd(getTestAccount(), adGroupCreated.Id, adCreated); e != nil {
		t.Error(e)
	} else {
		adCreated.SecondaryStatus = ""
		adUpdated.SecondaryStatus = ""

		if !reflect.DeepEqual(adCreated, adUpdated) {
			buf, _ := json.Marshal(adCreated)
			log.Println(string(buf))
			buf, _ = json.Marshal(adUpdated)
			log.Println(string(buf))
			t.Error("ad not equal after update")
		}
	}

	if err = DeleteAd(getTestAccount(), adCreated.Id); err != nil {
		t.Error(err)
	}

	ad = NewCatalogAdWithCustomVideo("custom vid", identityId, catalogId, "https://jolivetboisnegoce.fr", "v10033g50000d6j92l7og65qb0aa5jhg", "ad-site-i18n-sg/202603045d0dcca471a9e75240b7945c")
	ad.SetAdText("Best products")
	if _, e := UpdateAd(getTestAccount(), adGroupCreated.Id, ad); e != nil {
		t.Error(e)
	}
}

func TestDeleteCampaign(t *testing.T) {
	initTestSession()

	err := DeleteCampaign(getTestAccount(), "1856300822349010")
	if nil != err {
		t.Fatal(err)
	}
}
