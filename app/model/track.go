package model

type Track struct {
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Duration string `jsoin:"duration"`
	CoverURL string `json:"cover_url"`
	FileName string `json:"file_name"`
	UUID     string `json:"id"`
}
