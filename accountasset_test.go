package tiktokads

import (
	"fmt"
	"net/url"
	"testing"
	"time"
)

func TestGetAccountVideos(t *testing.T) {
	initTestSession()

	videos, err := GetAccountVideos(getTestAccount())
	assertNotEmptyResult(t, videos, err)
	for _, video := range videos {
		if len(video.Id) == 0 {
			t.Error("video id should not be empty")
		}
		if video.Size*video.Width*video.Height == 0 {
			t.Error("video resolution or size should not be zero")
		}
	}
}

func TestUploadAccountVideo(t *testing.T) {
	initTestSession()

	u := "https://storage.googleapis.com/feedcast-storage/tests/wood-vid-portrait.mp4"
	fileName := fmt.Sprintf("video.%s.mp4", time.Now().Format(time.DateTime))
	vid, err := UploadAccountVideo(getTestAccount(), fileName, u)

	if nil != err {
		t.Fatal(err)
	}
	if len(vid.Id) == 0 {
		t.Error("vid id should not be empty")
	}
}

func TestGetAccountImages(t *testing.T) {
	initTestSession()
	images, err := GetAccountImages(getTestAccount())
	assertNotEmptyResult(t, images, err)
	for _, image := range images {
		if len(image.Id) == 0 {
			t.Error("image id should not be empty")
		}
		if image.Size*image.Width*image.Height == 0 {
			t.Error("image resolution or size should not be zero")
		}
		if _, err = url.Parse(image.Url); err != nil {
			t.Error(err)
		}
	}
}

func TestUploadAccountImage(t *testing.T) {
	initTestSession()

	u := "https://www.orixa-media.com/wp-content/themes/orixa/dist/8dd25ecf07e1745079f1.jpeg"
	fileName := fmt.Sprintf("video.%s.mp4", time.Now().Format(time.DateTime))

	image, err := UploadAccountImage(getTestAccount(), fileName, u)
	if nil != err {
		t.Fatal(err)
	}
	if len(image.Id) == 0 {
		t.Error("vid id should not be empty")
	}
	if image.Size*image.Width*image.Height == 0 {
		t.Error("video resolution or size should not be zero")
	}
	if image.FileName != fileName {
		t.Error("file name should be equal")
	}
	t.Logf("URL = %s", image.Url)
}
