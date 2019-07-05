package file

import (
	"time"
)

type Data struct {
	filepath     string
	modifiedTime time.Time
}

func NewData(filepath string, modifiedTime time.Time) *Data {
	createdData := Data{}

	createdData.SetFilepath(filepath)
	createdData.SetModifiedTime(modifiedTime)

	return &createdData
}

func (d *Data) GetFilepath() string {
	return d.filepath
}

func (d *Data) GetModifiedTime() time.Time {
	return d.modifiedTime
}

func (d *Data) SetFilepath(filepath string) {
	d.filepath = filepath
}

func (d *Data) SetModifiedTime(modifiedTime time.Time) {
	d.modifiedTime = modifiedTime
}
