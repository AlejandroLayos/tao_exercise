package model

type DataCSV struct {
	Question string `csv:"Question text"`
	Created  string `csv:"Created At"`
	Choice1  string `csv:"Choice 1"`
	Choice   string `csv:"Choice"`
	Choice2  string `csv:"Choice 3"`
}
