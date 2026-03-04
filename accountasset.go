package tiktokads

import (
	"encoding/json"
	"errors"
	"math"
	"time"
)

type AssetFields struct {
	MaterialId string    `json:"material_id,omitempty"`
	Format     string    `json:"format,omitempty"`
	FileName   string    `json:"file_name,omitempty"`
	Height     int       `json:"height,omitempty"`
	Width      int       `json:"width,omitempty"`
	Size       int       `json:"size,omitempty"`
	Created    time.Time `json:"create_time,omitempty"`
}

type AssetImage struct {
	*AssetFields
	Id               string `json:"image_id,omitempty"`
	Displayable      bool   `json:"displayable,omitempty"`
	IsCarouselUsable bool   `json:"is_carousel_usable,omitempty"`
	Url              string `json:"image_url,omitempty"`
}

type AdVideo struct {
	*AssetFields
	Id         string  `json:"video_id,omitempty"`
	Duration   float32 `json:"duration,omitempty"`
	PreviewUrl string  `json:"preview_url,omitempty"`
}

type uploadVideoRequest struct {
	AdvertiserId    string `json:"advertiser_id,omitempty"`
	Url             string `json:"video_url,omitempty"`
	FileName        string `json:"file_name,omitempty"`
	FlawDetect      bool   `json:"flaw_detect"`
	AutoBindEnabled bool   `json:"auto_bind_enabled"`
	AutoFixEnabled  bool   `json:"auto_fix_enabled"`
	UploadType      string `json:"upload_type"`
}

type uploadImageRequest struct {
	AdvertiserId string `json:"advertiser_id,omitempty"`
	UploadType   string `json:"upload_type"`
	FileName     string `json:"file_name,omitempty"`
	Url          string `json:"image_url,omitempty"`
}

type uploadVideoResponse struct {
	*AdVideo
}

func (u *uploadVideoResponse) UnmarshalJSON(bytes []byte) error {
	var vid AdVideo
	err := json.Unmarshal(bytes, &vid)
	if nil == err {
		if len(vid.Id) > 0 {
			u.AdVideo = &vid
		}
	} else {
		var vids []AdVideo
		err = json.Unmarshal(bytes, &vids)
		if nil == err && len(vids) == 1 {
			u.AdVideo = &vids[0]
		}
	}

	return nil
}

func GetAccountVideos(accountId string) ([]*AdVideo, error) {
	req := newGetRequest(
		urlAccountVideoGet,
		withAccountId(accountId),
	)

	return fetchAllPages[AdVideo](req, 100)
}

func UploadAccountVideo(accountId string, fileName, url string) (*AdVideo, error) {
	req := newPostRequest(
		urlAccountVideoUpload,
		&uploadVideoRequest{
			AdvertiserId:    accountId,
			Url:             url,
			FileName:        fileName,
			UploadType:      "UPLOAD_BY_URL",
			FlawDetect:      true,
			AutoBindEnabled: true,
			AutoFixEnabled:  true,
		},
	)
	minCreated := time.Now()

	// ignore result: video_id returned in api on creation is not the same as the on returned in video list api
	_, err := fetch[emptyResult](req)

	if nil != err {
		return nil, err
	} else {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * time.Duration(math.Pow(2, float64(i))))
			vids, _ := GetAccountVideos(accountId)
			for _, vid := range vids {
				// try match new video by creation date & filename
				if minCreated.Before(vid.Created) && vid.FileName == fileName {
					return vid, nil
				}
			}
		}
	}

	return nil, errors.New("video-not-found")
}

func GetAccountImages(accountId string) ([]*AssetImage, error) {
	req := newGetRequest(
		urlAccountImageGet,
		withAccountId(accountId),
	)

	return fetchAllPages[AssetImage](req, 100)
}

func UploadAccountImage(accountId string, fileName, url string) (*AssetImage, error) {
	req := newPostRequest(
		urlAccountImageUpload,
		&uploadImageRequest{
			AdvertiserId: accountId,
			UploadType:   "UPLOAD_BY_URL",
			FileName:     fileName,
			Url:          url,
		},
	)
	_, err := fetch[emptyResult](req)

	if nil != err {
		return nil, err
	} else {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * time.Duration(math.Pow(2, float64(i))))
			images, _ := GetAccountImages(accountId)
			for _, image := range images {
				if image.FileName == fileName {
					return image, nil
				}
			}
		}
	}

	return nil, errors.New("imag-not-found")
}
