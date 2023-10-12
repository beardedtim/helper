package view

type PageData struct {
	Title string
	// This will allow us to use /assets in local development
	// or use some sort of CDN/Blobl Storage in production
	AssetURL string
}
