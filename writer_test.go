package writer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const timeInSec = 3600

func TestNewLogWriter(t *testing.T) {
	cases := []struct {
		filename    string
		inputString []string
		expected    []string
		expectedErr error
	}{
		{
			filename: "/tmp/test.log",
			inputString: []string{
				"Hello World",
				"Lorem Ipsum",
			},
			expected: []string{
				"Hello WorldLorem Ipsum",
			},
			expectedErr: nil,
		},
		{
			filename: "../../../test.log",
			inputString: []string{
				"Hello World\n",
				"Lorem Ipsum",
			},
			expected: []string{
				"Hello World",
				"Lorem Ipsum",
			},
			expectedErr: nil,
		},
		{
			filename: "/asdfasdf/test.log",
			inputString: []string{
				"Hello World\n",
				"Lorem Ipsum",
			},
			expected: []string{
				"Hello World",
				"Lorem Ipsum",
			},
			expectedErr: errCreatingFile,
		},
	}

	Convey("Log Writer", t, func() {
		for _, c := range cases {
			Convey("File Creation: "+c.filename, func() {
				wr, err := NewWriter(c.filename, timeInSec)
				if err != nil && err != c.expectedErr {
					t.Errorf("Error Initializing log file: " + c.filename + "\n\t" + err.Error())
				}
				if err == nil {
					wr.Close()
				}

			})

		}

		for _, c := range cases {
			Convey("Writing To file:"+c.filename, func() {
				rotateWriter, err := NewWriter(c.filename, timeInSec)
				if err != nil && err != c.expectedErr {
					t.Errorf("Error creating NewWriter: \n\t" + err.Error())
				}
				if rotateWriter != nil {
					for _, input := range c.inputString {
						_, err := rotateWriter.Write([]byte(input))
						if err != nil && err != c.expectedErr {
							t.Errorf("Error writing log to file: " + c.filename + "\n\t" + err.Error())
						}
					}
					rotateWriter.Close()
				}
			})

		}
		for _, c := range cases {
			Convey("Read From Log File: "+c.filename, func() {
				rotateWriter, err := NewWriter(c.filename, timeInSec)
				if rotateWriter != nil {
					for _, input := range c.inputString {
						_, err := rotateWriter.Write([]byte(input))
						if err != nil && err != c.expectedErr {
							t.Errorf("Error writing log to file: " + c.filename + "\n\t" + err.Error())
						}
					}
					_, err = rotateWriter.read()
					if err != nil && err != c.expectedErr {
						t.Errorf("Error reading log file: " + c.filename + "\n\t" + err.Error())
					}
				}

			})
		}
		for _, c := range cases {
			Convey("Match Written & Read Content: "+c.filename, func() {
				rotateWriter, _ := NewWriter(c.filename, timeInSec)
				if rotateWriter != nil {
					for _, input := range c.inputString {
						_, err := rotateWriter.Write([]byte(input))
						if err != nil && err != c.expectedErr {
							t.Errorf("Error writing log to file: " + c.filename + "\n\t" + err.Error())
						}
					}

					logContent, err := rotateWriter.read()
					if err != nil && err != c.expectedErr {
						t.Errorf("Error reading log file: " + c.filename + "\n\t" + err.Error())
					}

					for idx, val := range c.expected {
						if val != logContent[idx] {
							So(val, ShouldEqual, logContent[idx])
						}
					}
				}
			})
		}
	})
}
