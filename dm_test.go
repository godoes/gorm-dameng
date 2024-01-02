package dameng

import (
	"bytes"
	"testing"

	"gorm.io/gorm"
)

func TestDialector_QuoteTo(t *testing.T) {
	testData := []struct {
		raw    string
		expect string
	}{
		{"database.tableUser", "\"database\".\"tableUser\""},
		{"database.table`User", "\"database\".\"table`User\""},
		{"`a`.`b`", "\"`a`\".\"`b`\""},
		{"`a`.b`", "\"`a`\".\"b`\""},
		{"a.`b`", "\"a\".\"`b`\""},
		{"`a`.b`c", "\"`a`\".\"b`c\""},
		{"`a`.`b`c`", "\"`a`\".\"`b`c`\""},
		{"`a`.b", "\"`a`\".\"b\""},
		{"`ab`", "\"`ab`\""},
		{"`a``b`", "\"`a``b`\""},
		{"`a```b`", "\"`a```b`\""},
		{"a`b", "\"a`b\""},
		{"ab", "\"ab\""},
		{"`a.b`", "\"`a\".\"b`\""},
		{"a.b", "\"a\".\"b\""},
	}

	dialector := Open("")
	for _, item := range testData {
		buf := &bytes.Buffer{}
		dialector.QuoteTo(buf, item.raw)
		if buf.String() != item.expect {
			t.Fatalf("quote %q fail, got %q, expect %q", item.raw, buf.String(), item.expect)
		}
	}
}

// BenchmarkDialector_QuoteTo
// Result:
// goos: darwin
// goarch: amd64
// pkg: gorm.io/driver/mysql
// cpu: Intel(R) Core(TM) i7-8750H CPU @ 2.20GHz
// BenchmarkDialector_QuoteTo               9184232               113.2 ns/op
// BenchmarkDialector_QuoteTo-2             9782818               112.3 ns/op
// BenchmarkDialector_QuoteTo-4            10726722               109.0 ns/op
// BenchmarkDialector_QuoteTo-8             9656778               113.1 ns/op
// BenchmarkDialector_QuoteTo-12           10729615               112.7 ns/op
func BenchmarkDialector_QuoteTo(b *testing.B) {
	dialector := Open("")
	buf := &bytes.Buffer{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dialector.QuoteTo(buf, "database.table`User")
		buf.Reset()
	}
}

func TestNew(t *testing.T) {
	testWaitInit()
	type args struct {
		config Config
	}
	options := map[string]string{
		"schema":         dmSchema,
		"appName":        "dm_TestNew",
		"connectTimeout": "30000",
	}

	tests := []struct {
		name string
		args args
		//want gorm.Dialector
	}{
		{"TestNew", args{Config{
			DriverName:              DriverName,
			DSN:                     BuildUrl(dmUsername, dmPassword, dmHost, dmPort, options),
			VarcharSizeIsCharLength: true,
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDB, gotErr := gorm.Open(New(tt.args.config))
			if gotErr != nil {
				t.Errorf("New() got error: %v", gotErr)
			}
			var version string
			gotErr = gotDB.Raw("SELECT BANNER FROM V$VERSION WHERE BANNER LIKE 'DB Version:%'").Row().Scan(&version)
			if gotErr == nil {
				t.Log("DM Version:", version)
			} else {
				t.Errorf("Scan() got error: %v", gotErr)
			}
		})
	}
}
