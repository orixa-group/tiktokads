package tiktokads

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestUpdateCatalogFeed(t *testing.T) {
	initTestSession()

	catalog := NewCatalogFeed("XML Feed", "https://storage.googleapis.com/feedcast-storage/tests/_toto_xml.xml")
	businessId := "7489850869144485895"
	catalogId := "7594833492484146945"

	created, err := UpdateCatalogFeed(businessId, catalogId, catalog)
	if nil != err {
		t.Fatal(err)
	} else if len(created.FeedId) == 0 {
		t.Fatal("created feed should have id")
	}

	created.FeedName = fmt.Sprintf("XML Feed updated [%s]", time.Now().Format(time.DateTime))
	created.LastUpdateParam.Uri = fmt.Sprintf("%s?rand=%d", created.LastUpdateParam.Uri, time.Now().Unix())

	updated, err := UpdateCatalogFeed(businessId, catalogId, created)

	if nil != err {
		t.Fatal(err)
	} else if !reflect.DeepEqual(updated, created) {
		t.Fatal("updated & created feeds should be equal")
	}
}

func TestGetCatalogFeeds(t *testing.T) {
	initTestSession()
	businessId := "7489850869144485895"
	catalogId := "7594833492484146945"

	feeds, err := GetCatalogFeeds(businessId, catalogId)
	assertNotEmptyResult(t, feeds, err)

	for _, feed := range feeds {
		if id, _ := strconv.Atoi(feed.FeedId); 0 == id {
			t.Error("feed id should be numeric")
		}
		if len(feed.LastUpdateParam.Uri) == 0 {
			t.Error("lastUpdateParam.Uri should be set")
		}
	}
}
