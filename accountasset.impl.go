package tiktokads

func (i *AssetImage) GetId() string {
	return i.MaterialId
}

func (i *AssetImage) GetName() string {
	return i.FileName
}

func (i *AssetImage) GetFormat() string {
	return i.Format
}

func (i *AssetImage) GetWidth() int {
	return i.Width
}

func (i *AssetImage) GetHeight() int {
	return i.Height
}

func (i *AssetImage) GetSize() int {
	return i.Size
}

func (i *AssetImage) GetUrl() string {
	return i.Url
}

func (v *AdVideo) GetId() string {
	return v.MaterialId
}

func (v *AdVideo) GetName() string {
	return v.FileName
}

func (v *AdVideo) GetFormat() string {
	return v.Format
}

func (v *AdVideo) GetWidth() int {
	return v.Width
}

func (v *AdVideo) GetHeight() int {
	return v.Height
}

func (v *AdVideo) GetSize() int {
	return v.Size
}

func (v *AdVideo) GetUrl() string {
	return v.PreviewUrl
}
