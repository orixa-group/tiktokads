package tiktokads

// OAuth
const urlOAuthInstall = "https://business-api.tiktok.com/portal/auth"
const urlOAuthFetchTokenFromAuthCode = "https://business-api.tiktok.com/open_api/v1.3/oauth2/access_token/"

// Business API
const urlAccountsFetch = "https://business-api.tiktok.com/open_api/v1.3/oauth2/advertiser/get/"
const urlBusinessCentersFetch = "https://business-api.tiktok.com/open_api/v1.3/bc/get/"

// catalogs
const urlCatalogCreate = "https://business-api.tiktok.com/open_api/v1.3/catalog/create/"
const urlCatalogGet = "https://business-api.tiktok.com/open_api/v1.3/catalog/get/"
const urlCatalogDelete = "https://business-api.tiktok.com/open_api/v1.3/catalog/delete/"
const urlCatalogFeedGet = "https://business-api.tiktok.com/open_api/v1.3/catalog/feed/get/"
const urlCatalogFeedCreate = "https://business-api.tiktok.com/open_api/v1.3/catalog/feed/create/"
const urlCatalogFeedUpdate = "https://business-api.tiktok.com/open_api/v1.3/catalog/feed/update/"
const urlCatalogSetGet = "https://business-api.tiktok.com/open_api/v1.3/catalog/set/get/"
const urlCatalogSetUpdate = "https://business-api.tiktok.com/open_api/v1.3/catalog/set/update/"
const urlCatalogSetCreate = "https://business-api.tiktok.com/open_api/v1.3/catalog/set/create/"
const urlCatalogAddEvent = "https://business-api.tiktok.com/open_api/v1.3/catalog/eventsource/bind/"

const urlBusinessPixelGet = "https://business-api.tiktok.com/open_api/v1.3/bc/pixel/get/"
const urlPixelGet = "https://business-api.tiktok.com/open_api/v1.3/pixel/list/"

// campaigns
const urlCampaignSmartCreate = "https://business-api.tiktok.com/open_api/v1.3/smart_plus/campaign/create/"
const urlCampaignSmartUpdate = "https://business-api.tiktok.com/open_api/v1.3/smart_plus/campaign/update/"
const urlCampaignsFetch = "https://business-api.tiktok.com/open_api/v1.3/campaign/get/"
const urlCampaignCreate = "https://business-api.tiktok.com/open_api/v1.3/campaign/create/"
const urlCampaignUpdate = "https://business-api.tiktok.com/open_api/v1.3/campaign/update/"
const urlCampaignStatusUpdate = "https://business-api.tiktok.com/open_api/v1.3/campaign/status/update/"

// adgroups
const urlAdGroupGet = "https://business-api.tiktok.com/open_api/v1.3/adgroup/get/"
const urlAdGroupCreate = "https://business-api.tiktok.com/open_api/v1.3/adgroup/create/"
const urlAdGroupUpdate = "https://business-api.tiktok.com/open_api/v1.3/adgroup/update/"

const urlSmartAdGroupCreate = "https://business-api.tiktok.com/open_api/v1.3/smart_plus/adgroup/create/"
const urlSmartAdGroupUpdate = "https://business-api.tiktok.com/open_api/v1.3/smart_plus/adgroup/update/"

// ads
const urlAdCreate = "https://business-api.tiktok.com/open_api/v1.3/ad/create/"
const urlAdUpdate = "https://business-api.tiktok.com/open_api/v1.3/ad/update/"
const urlAdGet = "https://business-api.tiktok.com/open_api/v1.3/ad/get/"
const urlAdPreview = "https://business-api.tiktok.com/open_api/v1.3/creative/ads_preview/create/"
const urlAdStatusUpdate = "https://business-api.tiktok.com/open_api/v1.3/ad/status/update/"

const urlSmartAdStatusCreate = "https://business-api.tiktok.com/open_api/v1.3/smart_plus/ad/status/create/"
const urlSmartAdStatusUpdate = "https://business-api.tiktok.com/open_api/v1.3/smart_plus/ad/status/update/"

// acccount assets (image&vids)
const urlAccountVideoGet = "https://business-api.tiktok.com/open_api/v1.3/file/video/ad/search/"
const urlAccountVideoUpload = "https://business-api.tiktok.com/open_api/v1.3/file/video/ad/upload/"
const urlAccountImageGet = "https://business-api.tiktok.com/open_api/v1.3/file/image/ad/search/"
const urlAccountImageUpload = "https://business-api.tiktok.com/open_api/v1.3/file/image/ad/upload/"

// audiences
const urlAudienceGet = "https://business-api.tiktok.com/open_api/v1.3/dmp/custom_audience/list/"

// const urlInterestsGet = "https://business-api.tiktok.com/open_api/v1.3/targeting/search/"
const urlInterestsGet = "https://business-api.tiktok.com/open_api/v1.3/tool/interest_keyword/recommend/"
const urlIdentityGet = "https://business-api.tiktok.com/open_api/v1.3/identity/get/"

// reporting
const urlReporting = "https://business-api.tiktok.com/open_api/v1.3/report/integrated/get/"

// musics
const urlGetMusics = "https://business-api.tiktok.com/open_api/v1.3/file/music/get"

// ctas
const urlGetCtas = "https://business-api.tiktok.com/open_api/v1.3/creative/cta/recommend/"

// product report
const urlCatalogStatusTaskCreate = "https://business-api.tiktok.com/open_api/v1.3/diagnostic/catalog/product/task/create/"
const urlCatalogStatusTaskGet = "https://business-api.tiktok.com/open_api/v1.3/diagnostic/catalog/product/task/get/"
