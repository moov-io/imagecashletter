package files

import "github.com/moov-io/imagecashletter"

type testICLFileRepository struct {
	err error

	file *imagecashletter.File
}

func (r *testICLFileRepository) GetFiles() ([]*imagecashletter.File, error) {
	if r.err != nil {
		return nil, r.err
	}
	return []*imagecashletter.File{r.file}, nil
}

func (r *testICLFileRepository) GetFile(fileId string) (*imagecashletter.File, error) {
	if r.err != nil {
		return nil, r.err
	}

	if r.file != nil && r.file.ID == fileId {
		return r.file, nil
	}

	return nil, nil
}

func (r *testICLFileRepository) SaveFile(file *imagecashletter.File) error {
	if r.err == nil { // only persist if we're not error'ing
		r.file = file
	}
	return r.err
}

func (r *testICLFileRepository) DeleteFile(fileId string) error {
	return r.err
}
