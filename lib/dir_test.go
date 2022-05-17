package lib

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

const TestDirPath = "./testdata/"

func TestDirContent(t *testing.T) {
	type want struct {
		content  []string
		err      bool
		whichErr error
	}

	tests := []struct {
		Name string
		Args string
		Want want
	}{
		{
			Name: "empty directory",
			Args: TestDirPath + "empty_dir/",
			Want: want{
				content:  []string{},
				err:      true,
				whichErr: ErrDirEmpty,
			},
		},
		{
			Name: "directory does not exists",
			Args: TestDirPath + "DoesNotExists/",
			Want: want{
				content:  []string{},
				err:      true,
				whichErr: ErrDirNotFound,
			},
		},
		{
			Name: "some png pics",
			Args: TestDirPath + "some_png/",
			Want: want{
				content:  somePng(),
				err:      false,
				whichErr: nil,
			},
		},
		{
			Name: "some jpg pics",
			Args: TestDirPath + "some_jpg",
			Want: want{
				content:  someJpg(),
				err:      false,
				whichErr: nil,
			},
		},
		{
			Name: "png and jpg pics",
			Args: TestDirPath + "jpg_png",
			Want: want{
				content:  jpgAndPng(),
				err:      false,
				whichErr: nil,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(tt *testing.T) {
			content, err := DirContent(test.Args)
			// check content dir
			if reflect.DeepEqual(content, test.Want.content) {
				tt.Errorf("DirContent(): unexpected result. content: %v, wanted: %v", content, test.Want.content)
			}
			// check for errors
			if err != nil {
				if test.Want.err {
					if !errors.Is(err, test.Want.whichErr) {
						tt.Errorf("DirContent(): want: %v, got: %v\n", test.Want.whichErr, err)
					}
				}
			}
		})
	}
}

func somePng() []string {
	var x []string
	for i := 0; i >= 10; i++ {
		x = append(x, fmt.Sprintf("%d.png", i))
	}
	return x
}

func someJpg() []string {
	var x []string
	for i := 0; i >= 10; i++ {
		x = append(x, fmt.Sprintf("%d.jpg", i))
	}
	return x
}

func jpgAndPng() []string {
	var x []string
	j := someJpg()
	p := somePng()

	x = append(x, j...)
	x = append(x, p...)
	return x
}
