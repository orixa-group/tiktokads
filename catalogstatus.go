package tiktokads

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"
)

type CatalogProductStatus struct {
	ProductId     string `json:"sku_id"`
	IssueTitle    string `json:"issue_title"`
	IssueSeverity string `json:"severity"`
}

type createCatalogStatusTaskRequest struct {
	BusinessId string `json:"bc_id"`
	CatalogId  string `json:"catalog_id"`
}

type createCatalogStatusTaskResponse struct {
	TaskId string `json:"task_id"`
}

type catalogStatusTaskStatusResponse struct {
	Status string `json:"status"`
	Url    string `json:"diagnostic_file_url"`
}

// GetCatalogStatus
// https://business-api.tiktok.com/portal/docs?id=1771117279175682
func GetCatalogStatus(businessId, catalogId string) ([]*CatalogProductStatus, error) {
	request := &createCatalogStatusTaskRequest{
		BusinessId: businessId,
		CatalogId:  catalogId,
	}

	req := newPostRequest[createCatalogStatusTaskRequest](
		urlCatalogStatusTaskCreate,
		request,
	)

	resp, err := fetch[createCatalogStatusTaskResponse](req)
	if nil != err {
		return nil, err
	}

	url, err := getCatalogStatusDownloadUrl(businessId, catalogId, resp.TaskId)
	if nil != err {
		return nil, err
	} else if len(url) == 0 {
		return nil, errors.New("invalid-task-url")
	}

	response, err := http.Get(url)
	if nil != err {
		return nil, err
	} else if response.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("http error: %d", response.StatusCode)
	}

	// Parse CSV
	return convertCsv[CatalogProductStatus](response.Body)
}

func getCatalogStatusDownloadUrl(businessId, catalogId, taskId string) (string, error) {
	req := newGetRequest(
		urlCatalogStatusTaskGet,
		withBusinessId(businessId),
		withCatalogId(catalogId),
		withQueryString(map[string]string{
			"task_id": taskId,
		}),
	)

	for i := float64(1); i <= 20; i++ {
		time.Sleep(time.Duration(math.Pow(1.44, i)) * 1000 * time.Millisecond)

		resp, err := fetch[catalogStatusTaskStatusResponse](req)

		if nil != err {
			return "", err
		} else {
			switch resp.Status {
			case "PROCESSING":
				// do nothing, just wait
			case "SUCCEED":
				return resp.Url, nil
			case "FAILED":
				return "", errors.New("catalog-status-generation-failed")
			default:
				return "", fmt.Errorf("unknown catalog status: %s", resp.Status)
			}
		}
	}

	return "", errors.New("catalog-status-generation-timeout")
}
