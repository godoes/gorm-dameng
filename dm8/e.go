/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */

package dm8

import (
	"bytes"
	"errors"
	"io"
	"math"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
)

type dm_build_930 struct{}

var Dm_build_931 = &dm_build_930{}

func (Dm_build_933 *dm_build_930) Dm_build_932(dm_build_934 []byte, dm_build_935 int, dm_build_936 byte) int {
	dm_build_934[dm_build_935] = dm_build_936
	return 1
}

func (Dm_build_938 *dm_build_930) Dm_build_937(dm_build_939 []byte, dm_build_940 int, dm_build_941 int8) int {
	dm_build_939[dm_build_940] = byte(dm_build_941)
	return 1
}

func (Dm_build_943 *dm_build_930) Dm_build_942(dm_build_944 []byte, dm_build_945 int, dm_build_946 int16) int {
	dm_build_944[dm_build_945] = byte(dm_build_946)
	dm_build_945++
	dm_build_944[dm_build_945] = byte(dm_build_946 >> 8)
	return 2
}

func (Dm_build_948 *dm_build_930) Dm_build_947(dm_build_949 []byte, dm_build_950 int, dm_build_951 int32) int {
	dm_build_949[dm_build_950] = byte(dm_build_951)
	dm_build_950++
	dm_build_949[dm_build_950] = byte(dm_build_951 >> 8)
	dm_build_950++
	dm_build_949[dm_build_950] = byte(dm_build_951 >> 16)
	dm_build_950++
	dm_build_949[dm_build_950] = byte(dm_build_951 >> 24)
	dm_build_950++
	return 4
}

func (Dm_build_953 *dm_build_930) Dm_build_952(dm_build_954 []byte, dm_build_955 int, dm_build_956 int64) int {
	dm_build_954[dm_build_955] = byte(dm_build_956)
	dm_build_955++
	dm_build_954[dm_build_955] = byte(dm_build_956 >> 8)
	dm_build_955++
	dm_build_954[dm_build_955] = byte(dm_build_956 >> 16)
	dm_build_955++
	dm_build_954[dm_build_955] = byte(dm_build_956 >> 24)
	dm_build_955++
	dm_build_954[dm_build_955] = byte(dm_build_956 >> 32)
	dm_build_955++
	dm_build_954[dm_build_955] = byte(dm_build_956 >> 40)
	dm_build_955++
	dm_build_954[dm_build_955] = byte(dm_build_956 >> 48)
	dm_build_955++
	dm_build_954[dm_build_955] = byte(dm_build_956 >> 56)
	return 8
}

func (Dm_build_958 *dm_build_930) Dm_build_957(dm_build_959 []byte, dm_build_960 int, dm_build_961 float32) int {
	return Dm_build_958.Dm_build_977(dm_build_959, dm_build_960, math.Float32bits(dm_build_961))
}

func (Dm_build_963 *dm_build_930) Dm_build_962(dm_build_964 []byte, dm_build_965 int, dm_build_966 float64) int {
	return Dm_build_963.Dm_build_982(dm_build_964, dm_build_965, math.Float64bits(dm_build_966))
}

func (Dm_build_968 *dm_build_930) Dm_build_967(dm_build_969 []byte, dm_build_970 int, dm_build_971 uint8) int {
	dm_build_969[dm_build_970] = dm_build_971
	return 1
}

func (Dm_build_973 *dm_build_930) Dm_build_972(dm_build_974 []byte, dm_build_975 int, dm_build_976 uint16) int {
	dm_build_974[dm_build_975] = byte(dm_build_976)
	dm_build_975++
	dm_build_974[dm_build_975] = byte(dm_build_976 >> 8)
	return 2
}

func (Dm_build_978 *dm_build_930) Dm_build_977(dm_build_979 []byte, dm_build_980 int, dm_build_981 uint32) int {
	dm_build_979[dm_build_980] = byte(dm_build_981)
	dm_build_980++
	dm_build_979[dm_build_980] = byte(dm_build_981 >> 8)
	dm_build_980++
	dm_build_979[dm_build_980] = byte(dm_build_981 >> 16)
	dm_build_980++
	dm_build_979[dm_build_980] = byte(dm_build_981 >> 24)
	return 3
}

func (Dm_build_983 *dm_build_930) Dm_build_982(dm_build_984 []byte, dm_build_985 int, dm_build_986 uint64) int {
	dm_build_984[dm_build_985] = byte(dm_build_986)
	dm_build_985++
	dm_build_984[dm_build_985] = byte(dm_build_986 >> 8)
	dm_build_985++
	dm_build_984[dm_build_985] = byte(dm_build_986 >> 16)
	dm_build_985++
	dm_build_984[dm_build_985] = byte(dm_build_986 >> 24)
	dm_build_985++
	dm_build_984[dm_build_985] = byte(dm_build_986 >> 32)
	dm_build_985++
	dm_build_984[dm_build_985] = byte(dm_build_986 >> 40)
	dm_build_985++
	dm_build_984[dm_build_985] = byte(dm_build_986 >> 48)
	dm_build_985++
	dm_build_984[dm_build_985] = byte(dm_build_986 >> 56)
	return 3
}

func (Dm_build_988 *dm_build_930) Dm_build_987(dm_build_989 []byte, dm_build_990 int, dm_build_991 []byte, dm_build_992 int, dm_build_993 int) int {
	copy(dm_build_989[dm_build_990:dm_build_990+dm_build_993], dm_build_991[dm_build_992:dm_build_992+dm_build_993])
	return dm_build_993
}

func (Dm_build_995 *dm_build_930) Dm_build_994(dm_build_996 []byte, dm_build_997 int, dm_build_998 []byte, dm_build_999 int, dm_build_1000 int) int {
	dm_build_997 += Dm_build_995.Dm_build_977(dm_build_996, dm_build_997, uint32(dm_build_1000))
	return 4 + Dm_build_995.Dm_build_987(dm_build_996, dm_build_997, dm_build_998, dm_build_999, dm_build_1000)
}

func (Dm_build_1002 *dm_build_930) Dm_build_1001(dm_build_1003 []byte, dm_build_1004 int, dm_build_1005 []byte, dm_build_1006 int, dm_build_1007 int) int {
	dm_build_1004 += Dm_build_1002.Dm_build_972(dm_build_1003, dm_build_1004, uint16(dm_build_1007))
	return 2 + Dm_build_1002.Dm_build_987(dm_build_1003, dm_build_1004, dm_build_1005, dm_build_1006, dm_build_1007)
}

func (Dm_build_1009 *dm_build_930) Dm_build_1008(dm_build_1010 []byte, dm_build_1011 int, dm_build_1012 string, dm_build_1013 string, dm_build_1014 *DmConnection) int {
	dm_build_1015 := Dm_build_1009.Dm_build_1147(dm_build_1012, dm_build_1013, dm_build_1014)
	dm_build_1011 += Dm_build_1009.Dm_build_977(dm_build_1010, dm_build_1011, uint32(len(dm_build_1015)))
	return 4 + Dm_build_1009.Dm_build_987(dm_build_1010, dm_build_1011, dm_build_1015, 0, len(dm_build_1015))
}

func (Dm_build_1017 *dm_build_930) Dm_build_1016(dm_build_1018 []byte, dm_build_1019 int, dm_build_1020 string, dm_build_1021 string, dm_build_1022 *DmConnection) int {
	dm_build_1023 := Dm_build_1017.Dm_build_1147(dm_build_1020, dm_build_1021, dm_build_1022)

	dm_build_1019 += Dm_build_1017.Dm_build_972(dm_build_1018, dm_build_1019, uint16(len(dm_build_1023)))
	return 2 + Dm_build_1017.Dm_build_987(dm_build_1018, dm_build_1019, dm_build_1023, 0, len(dm_build_1023))
}

func (Dm_build_1025 *dm_build_930) Dm_build_1024(dm_build_1026 []byte, dm_build_1027 int) byte {
	return dm_build_1026[dm_build_1027]
}

func (Dm_build_1029 *dm_build_930) Dm_build_1028(dm_build_1030 []byte, dm_build_1031 int) int16 {
	var dm_build_1032 int16
	dm_build_1032 = int16(dm_build_1030[dm_build_1031] & 0xff)
	dm_build_1031++
	dm_build_1032 |= int16(dm_build_1030[dm_build_1031]&0xff) << 8
	return dm_build_1032
}

func (Dm_build_1034 *dm_build_930) Dm_build_1033(dm_build_1035 []byte, dm_build_1036 int) int32 {
	var dm_build_1037 int32
	dm_build_1037 = int32(dm_build_1035[dm_build_1036] & 0xff)
	dm_build_1036++
	dm_build_1037 |= int32(dm_build_1035[dm_build_1036]&0xff) << 8
	dm_build_1036++
	dm_build_1037 |= int32(dm_build_1035[dm_build_1036]&0xff) << 16
	dm_build_1036++
	dm_build_1037 |= int32(dm_build_1035[dm_build_1036]&0xff) << 24
	return dm_build_1037
}

func (Dm_build_1039 *dm_build_930) Dm_build_1038(dm_build_1040 []byte, dm_build_1041 int) int64 {
	var dm_build_1042 int64
	dm_build_1042 = int64(dm_build_1040[dm_build_1041] & 0xff)
	dm_build_1041++
	dm_build_1042 |= int64(dm_build_1040[dm_build_1041]&0xff) << 8
	dm_build_1041++
	dm_build_1042 |= int64(dm_build_1040[dm_build_1041]&0xff) << 16
	dm_build_1041++
	dm_build_1042 |= int64(dm_build_1040[dm_build_1041]&0xff) << 24
	dm_build_1041++
	dm_build_1042 |= int64(dm_build_1040[dm_build_1041]&0xff) << 32
	dm_build_1041++
	dm_build_1042 |= int64(dm_build_1040[dm_build_1041]&0xff) << 40
	dm_build_1041++
	dm_build_1042 |= int64(dm_build_1040[dm_build_1041]&0xff) << 48
	dm_build_1041++
	dm_build_1042 |= int64(dm_build_1040[dm_build_1041]&0xff) << 56
	return dm_build_1042
}

func (Dm_build_1044 *dm_build_930) Dm_build_1043(dm_build_1045 []byte, dm_build_1046 int) float32 {
	return math.Float32frombits(Dm_build_1044.Dm_build_1060(dm_build_1045, dm_build_1046))
}

func (Dm_build_1048 *dm_build_930) Dm_build_1047(dm_build_1049 []byte, dm_build_1050 int) float64 {
	return math.Float64frombits(Dm_build_1048.Dm_build_1065(dm_build_1049, dm_build_1050))
}

func (Dm_build_1052 *dm_build_930) Dm_build_1051(dm_build_1053 []byte, dm_build_1054 int) uint8 {
	return dm_build_1053[dm_build_1054] & 0xff
}

func (Dm_build_1056 *dm_build_930) Dm_build_1055(dm_build_1057 []byte, dm_build_1058 int) uint16 {
	var dm_build_1059 uint16
	dm_build_1059 = uint16(dm_build_1057[dm_build_1058] & 0xff)
	dm_build_1058++
	dm_build_1059 |= uint16(dm_build_1057[dm_build_1058]&0xff) << 8
	return dm_build_1059
}

func (Dm_build_1061 *dm_build_930) Dm_build_1060(dm_build_1062 []byte, dm_build_1063 int) uint32 {
	var dm_build_1064 uint32
	dm_build_1064 = uint32(dm_build_1062[dm_build_1063] & 0xff)
	dm_build_1063++
	dm_build_1064 |= uint32(dm_build_1062[dm_build_1063]&0xff) << 8
	dm_build_1063++
	dm_build_1064 |= uint32(dm_build_1062[dm_build_1063]&0xff) << 16
	dm_build_1063++
	dm_build_1064 |= uint32(dm_build_1062[dm_build_1063]&0xff) << 24
	return dm_build_1064
}

func (Dm_build_1066 *dm_build_930) Dm_build_1065(dm_build_1067 []byte, dm_build_1068 int) uint64 {
	var dm_build_1069 uint64
	dm_build_1069 = uint64(dm_build_1067[dm_build_1068] & 0xff)
	dm_build_1068++
	dm_build_1069 |= uint64(dm_build_1067[dm_build_1068]&0xff) << 8
	dm_build_1068++
	dm_build_1069 |= uint64(dm_build_1067[dm_build_1068]&0xff) << 16
	dm_build_1068++
	dm_build_1069 |= uint64(dm_build_1067[dm_build_1068]&0xff) << 24
	dm_build_1068++
	dm_build_1069 |= uint64(dm_build_1067[dm_build_1068]&0xff) << 32
	dm_build_1068++
	dm_build_1069 |= uint64(dm_build_1067[dm_build_1068]&0xff) << 40
	dm_build_1068++
	dm_build_1069 |= uint64(dm_build_1067[dm_build_1068]&0xff) << 48
	dm_build_1068++
	dm_build_1069 |= uint64(dm_build_1067[dm_build_1068]&0xff) << 56
	return dm_build_1069
}

func (Dm_build_1071 *dm_build_930) Dm_build_1070(dm_build_1072 []byte, dm_build_1073 int) []byte {
	dm_build_1074 := Dm_build_1071.Dm_build_1060(dm_build_1072, dm_build_1073)

	dm_build_1075 := make([]byte, dm_build_1074)
	copy(dm_build_1075[:int(dm_build_1074)], dm_build_1072[dm_build_1073+4:dm_build_1073+4+int(dm_build_1074)])
	return dm_build_1075
}

func (Dm_build_1077 *dm_build_930) Dm_build_1076(dm_build_1078 []byte, dm_build_1079 int) []byte {
	dm_build_1080 := Dm_build_1077.Dm_build_1055(dm_build_1078, dm_build_1079)

	dm_build_1081 := make([]byte, dm_build_1080)
	copy(dm_build_1081[:int(dm_build_1080)], dm_build_1078[dm_build_1079+2:dm_build_1079+2+int(dm_build_1080)])
	return dm_build_1081
}

func (Dm_build_1083 *dm_build_930) Dm_build_1082(dm_build_1084 []byte, dm_build_1085 int, dm_build_1086 int) []byte {

	dm_build_1087 := make([]byte, dm_build_1086)
	copy(dm_build_1087[:dm_build_1086], dm_build_1084[dm_build_1085:dm_build_1085+dm_build_1086])
	return dm_build_1087
}

func (Dm_build_1089 *dm_build_930) Dm_build_1088(dm_build_1090 []byte, dm_build_1091 int, dm_build_1092 int, dm_build_1093 string, dm_build_1094 *DmConnection) string {
	return Dm_build_1089.Dm_build_1183(dm_build_1090[dm_build_1091:dm_build_1091+dm_build_1092], dm_build_1093, dm_build_1094)
}

func (Dm_build_1096 *dm_build_930) Dm_build_1095(dm_build_1097 []byte, dm_build_1098 int, dm_build_1099 string, dm_build_1100 *DmConnection) string {
	dm_build_1101 := Dm_build_1096.Dm_build_1060(dm_build_1097, dm_build_1098)
	dm_build_1098 += 4
	return Dm_build_1096.Dm_build_1088(dm_build_1097, dm_build_1098, int(dm_build_1101), dm_build_1099, dm_build_1100)
}

func (Dm_build_1103 *dm_build_930) Dm_build_1102(dm_build_1104 []byte, dm_build_1105 int, dm_build_1106 string, dm_build_1107 *DmConnection) string {
	dm_build_1108 := Dm_build_1103.Dm_build_1055(dm_build_1104, dm_build_1105)
	dm_build_1105 += 2
	return Dm_build_1103.Dm_build_1088(dm_build_1104, dm_build_1105, int(dm_build_1108), dm_build_1106, dm_build_1107)
}

func (Dm_build_1110 *dm_build_930) Dm_build_1109(dm_build_1111 byte) []byte {
	return []byte{dm_build_1111}
}

func (Dm_build_1113 *dm_build_930) Dm_build_1112(dm_build_1114 int8) []byte {
	return []byte{byte(dm_build_1114)}
}

func (Dm_build_1116 *dm_build_930) Dm_build_1115(dm_build_1117 int16) []byte {
	return []byte{byte(dm_build_1117), byte(dm_build_1117 >> 8)}
}

func (Dm_build_1119 *dm_build_930) Dm_build_1118(dm_build_1120 int32) []byte {
	return []byte{byte(dm_build_1120), byte(dm_build_1120 >> 8), byte(dm_build_1120 >> 16), byte(dm_build_1120 >> 24)}
}

func (Dm_build_1122 *dm_build_930) Dm_build_1121(dm_build_1123 int64) []byte {
	return []byte{byte(dm_build_1123), byte(dm_build_1123 >> 8), byte(dm_build_1123 >> 16), byte(dm_build_1123 >> 24), byte(dm_build_1123 >> 32),
		byte(dm_build_1123 >> 40), byte(dm_build_1123 >> 48), byte(dm_build_1123 >> 56)}
}

func (Dm_build_1125 *dm_build_930) Dm_build_1124(dm_build_1126 float32) []byte {
	return Dm_build_1125.Dm_build_1136(math.Float32bits(dm_build_1126))
}

func (Dm_build_1128 *dm_build_930) Dm_build_1127(dm_build_1129 float64) []byte {
	return Dm_build_1128.Dm_build_1139(math.Float64bits(dm_build_1129))
}

func (Dm_build_1131 *dm_build_930) Dm_build_1130(dm_build_1132 uint8) []byte {
	return []byte{dm_build_1132}
}

func (Dm_build_1134 *dm_build_930) Dm_build_1133(dm_build_1135 uint16) []byte {
	return []byte{byte(dm_build_1135), byte(dm_build_1135 >> 8)}
}

func (Dm_build_1137 *dm_build_930) Dm_build_1136(dm_build_1138 uint32) []byte {
	return []byte{byte(dm_build_1138), byte(dm_build_1138 >> 8), byte(dm_build_1138 >> 16), byte(dm_build_1138 >> 24)}
}

func (Dm_build_1140 *dm_build_930) Dm_build_1139(dm_build_1141 uint64) []byte {
	return []byte{byte(dm_build_1141), byte(dm_build_1141 >> 8), byte(dm_build_1141 >> 16), byte(dm_build_1141 >> 24), byte(dm_build_1141 >> 32), byte(dm_build_1141 >> 40), byte(dm_build_1141 >> 48), byte(dm_build_1141 >> 56)}
}

func (Dm_build_1143 *dm_build_930) Dm_build_1142(dm_build_1144 []byte, dm_build_1145 string, dm_build_1146 *DmConnection) []byte {
	if dm_build_1145 == "UTF-8" {
		return dm_build_1144
	}

	if dm_build_1146 == nil {
		if e := dm_build_1188(dm_build_1145); e != nil {
			tmp, err := io.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1144), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1146.encodeBuffer == nil {
		dm_build_1146.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1146.encode = dm_build_1188(dm_build_1146.getServerEncoding())
		dm_build_1146.transformReaderDst = make([]byte, 4096)
		dm_build_1146.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1146.encode; e != nil {

		dm_build_1146.encodeBuffer.Reset()

		n, err := dm_build_1146.encodeBuffer.ReadFrom(
			Dm_build_1202(bytes.NewReader(dm_build_1144), e.NewEncoder(), dm_build_1146.transformReaderDst, dm_build_1146.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1146.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1148 *dm_build_930) Dm_build_1147(dm_build_1149 string, dm_build_1150 string, dm_build_1151 *DmConnection) []byte {
	return Dm_build_1148.Dm_build_1142([]byte(dm_build_1149), dm_build_1150, dm_build_1151)
}

func (Dm_build_1153 *dm_build_930) Dm_build_1152(dm_build_1154 []byte) byte {
	return Dm_build_1153.Dm_build_1024(dm_build_1154, 0)
}

func (Dm_build_1156 *dm_build_930) Dm_build_1155(dm_build_1157 []byte) int16 {
	return Dm_build_1156.Dm_build_1028(dm_build_1157, 0)
}

func (Dm_build_1159 *dm_build_930) Dm_build_1158(dm_build_1160 []byte) int32 {
	return Dm_build_1159.Dm_build_1033(dm_build_1160, 0)
}

func (Dm_build_1162 *dm_build_930) Dm_build_1161(dm_build_1163 []byte) int64 {
	return Dm_build_1162.Dm_build_1038(dm_build_1163, 0)
}

func (Dm_build_1165 *dm_build_930) Dm_build_1164(dm_build_1166 []byte) float32 {
	return Dm_build_1165.Dm_build_1043(dm_build_1166, 0)
}

func (Dm_build_1168 *dm_build_930) Dm_build_1167(dm_build_1169 []byte) float64 {
	return Dm_build_1168.Dm_build_1047(dm_build_1169, 0)
}

func (Dm_build_1171 *dm_build_930) Dm_build_1170(dm_build_1172 []byte) uint8 {
	return Dm_build_1171.Dm_build_1051(dm_build_1172, 0)
}

func (Dm_build_1174 *dm_build_930) Dm_build_1173(dm_build_1175 []byte) uint16 {
	return Dm_build_1174.Dm_build_1055(dm_build_1175, 0)
}

func (Dm_build_1177 *dm_build_930) Dm_build_1176(dm_build_1178 []byte) uint32 {
	return Dm_build_1177.Dm_build_1060(dm_build_1178, 0)
}

func (Dm_build_1180 *dm_build_930) Dm_build_1179(dm_build_1181 []byte, dm_build_1182 string) []byte {
	if dm_build_1182 == "UTF-8" {
		return dm_build_1181
	}

	if e := dm_build_1188(dm_build_1182); e != nil {

		tmp, err := io.ReadAll(
			transform.NewReader(bytes.NewReader(dm_build_1181), e.NewDecoder()),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return tmp
	}

	panic("Unsupported Charset!")

}

func (Dm_build_1184 *dm_build_930) Dm_build_1183(dm_build_1185 []byte, dm_build_1186 string, _ *DmConnection) string {
	return string(Dm_build_1184.Dm_build_1179(dm_build_1185, dm_build_1186))
}

func dm_build_1188(dm_build_1189 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1189); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1190 struct {
	dm_build_1191 io.Reader
	dm_build_1192 transform.Transformer
	dm_build_1193 error

	dm_build_1194                []byte
	dm_build_1195, dm_build_1196 int

	dm_build_1197                []byte
	dm_build_1198, dm_build_1199 int

	dm_build_1200 bool
}

const dm_build_1201 = 4096

func Dm_build_1202(dm_build_1203 io.Reader, dm_build_1204 transform.Transformer, dm_build_1205 []byte, dm_build_1206 []byte) *Dm_build_1190 {
	dm_build_1204.Reset()
	return &Dm_build_1190{
		dm_build_1191: dm_build_1203,
		dm_build_1192: dm_build_1204,
		dm_build_1194: dm_build_1205,
		dm_build_1197: dm_build_1206,
	}
}

func (dm_build_1208 *Dm_build_1190) Read(dm_build_1209 []byte) (int, error) {
	dm_build_1210, dm_build_1211 := 0, error(nil)
	for {

		if dm_build_1208.dm_build_1195 != dm_build_1208.dm_build_1196 {
			dm_build_1210 = copy(dm_build_1209, dm_build_1208.dm_build_1194[dm_build_1208.dm_build_1195:dm_build_1208.dm_build_1196])
			dm_build_1208.dm_build_1195 += dm_build_1210
			if dm_build_1208.dm_build_1195 == dm_build_1208.dm_build_1196 && dm_build_1208.dm_build_1200 {
				return dm_build_1210, dm_build_1208.dm_build_1193
			}
			return dm_build_1210, nil
		} else if dm_build_1208.dm_build_1200 {
			return 0, dm_build_1208.dm_build_1193
		}

		if dm_build_1208.dm_build_1198 != dm_build_1208.dm_build_1199 || dm_build_1208.dm_build_1193 != nil {
			dm_build_1208.dm_build_1195 = 0
			dm_build_1208.dm_build_1196, dm_build_1210, dm_build_1211 = dm_build_1208.dm_build_1192.Transform(dm_build_1208.dm_build_1194, dm_build_1208.dm_build_1197[dm_build_1208.dm_build_1198:dm_build_1208.dm_build_1199], dm_build_1208.dm_build_1193 == io.EOF)
			dm_build_1208.dm_build_1198 += dm_build_1210

			switch {
			case dm_build_1211 == nil:
				if dm_build_1208.dm_build_1198 != dm_build_1208.dm_build_1199 {
					dm_build_1208.dm_build_1193 = nil
				}

				dm_build_1208.dm_build_1200 = dm_build_1208.dm_build_1193 != nil
				continue
			case errors.Is(dm_build_1211, transform.ErrShortDst) && (dm_build_1208.dm_build_1196 != 0 || dm_build_1210 != 0):

				continue
			case errors.Is(dm_build_1211, transform.ErrShortSrc) && dm_build_1208.dm_build_1199-dm_build_1208.dm_build_1198 != len(dm_build_1208.dm_build_1197) && dm_build_1208.dm_build_1193 == nil:

			default:
				dm_build_1208.dm_build_1200 = true

				if dm_build_1208.dm_build_1193 == nil || dm_build_1208.dm_build_1193 == io.EOF {
					dm_build_1208.dm_build_1193 = dm_build_1211
				}
				continue
			}
		}

		if dm_build_1208.dm_build_1198 != 0 {
			dm_build_1208.dm_build_1198, dm_build_1208.dm_build_1199 = 0, copy(dm_build_1208.dm_build_1197, dm_build_1208.dm_build_1197[dm_build_1208.dm_build_1198:dm_build_1208.dm_build_1199])
		}
		dm_build_1210, dm_build_1208.dm_build_1193 = dm_build_1208.dm_build_1191.Read(dm_build_1208.dm_build_1197[dm_build_1208.dm_build_1199:])
		dm_build_1208.dm_build_1199 += dm_build_1210
	}
}
