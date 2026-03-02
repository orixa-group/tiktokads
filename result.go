package tiktokads

// apiResult
type apiResult[T any] struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"request_id"`
	Data      *T     `json:"data"` // polymorphic. Can either be structure or array of structure
}

type emptyResult struct {
}

// listResult structure for array of structure in apiResult
// Results can be in 'list' json attribute or other attributes ('identity_list') depending on resource type
type listResult[T any] struct {
	List         []*T `json:"list"`
	IdentityList []*T `json:"identity_list"`
}

func (l *listResult[T]) GetResults() []*T {
	if len(l.IdentityList) > 0 {
		return l.IdentityList
	}

	return l.List
}

type listFeedResult[T any] struct {
	FeedList []*T `json:"feed_list"`
}

// getResultsFromListResults Helper to get inner results from listResult
func getResultsFromListResults[L any](res *apiResult[listResult[L]], err error) ([]*L, error) {
	if err != nil {
		return nil, err
	}

	return res.Data.List, nil
}
