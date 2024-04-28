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

type dm_build_0 struct{}

var Dm_build_1 = &dm_build_0{}

func (Dm_build_3 *dm_build_0) Dm_build_2(dm_build_4 []byte, dm_build_5 int, dm_build_6 byte) int {
	dm_build_4[dm_build_5] = dm_build_6
	return 1
}

func (Dm_build_8 *dm_build_0) Dm_build_7(dm_build_9 []byte, dm_build_10 int, dm_build_11 int8) int {
	dm_build_9[dm_build_10] = byte(dm_build_11)
	return 1
}

func (Dm_build_13 *dm_build_0) Dm_build_12(dm_build_14 []byte, dm_build_15 int, dm_build_16 int16) int {
	dm_build_14[dm_build_15] = byte(dm_build_16)
	dm_build_15++
	dm_build_14[dm_build_15] = byte(dm_build_16 >> 8)
	return 2
}

func (Dm_build_18 *dm_build_0) Dm_build_17(dm_build_19 []byte, dm_build_20 int, dm_build_21 int32) int {
	dm_build_19[dm_build_20] = byte(dm_build_21)
	dm_build_20++
	dm_build_19[dm_build_20] = byte(dm_build_21 >> 8)
	dm_build_20++
	dm_build_19[dm_build_20] = byte(dm_build_21 >> 16)
	dm_build_20++
	dm_build_19[dm_build_20] = byte(dm_build_21 >> 24)
	dm_build_20++
	return 4
}

func (Dm_build_23 *dm_build_0) Dm_build_22(dm_build_24 []byte, dm_build_25 int, dm_build_26 int64) int {
	dm_build_24[dm_build_25] = byte(dm_build_26)
	dm_build_25++
	dm_build_24[dm_build_25] = byte(dm_build_26 >> 8)
	dm_build_25++
	dm_build_24[dm_build_25] = byte(dm_build_26 >> 16)
	dm_build_25++
	dm_build_24[dm_build_25] = byte(dm_build_26 >> 24)
	dm_build_25++
	dm_build_24[dm_build_25] = byte(dm_build_26 >> 32)
	dm_build_25++
	dm_build_24[dm_build_25] = byte(dm_build_26 >> 40)
	dm_build_25++
	dm_build_24[dm_build_25] = byte(dm_build_26 >> 48)
	dm_build_25++
	dm_build_24[dm_build_25] = byte(dm_build_26 >> 56)
	return 8
}

func (Dm_build_28 *dm_build_0) Dm_build_27(dm_build_29 []byte, dm_build_30 int, dm_build_31 float32) int {
	return Dm_build_28.Dm_build_47(dm_build_29, dm_build_30, math.Float32bits(dm_build_31))
}

func (Dm_build_33 *dm_build_0) Dm_build_32(dm_build_34 []byte, dm_build_35 int, dm_build_36 float64) int {
	return Dm_build_33.Dm_build_52(dm_build_34, dm_build_35, math.Float64bits(dm_build_36))
}

func (Dm_build_38 *dm_build_0) Dm_build_37(dm_build_39 []byte, dm_build_40 int, dm_build_41 uint8) int {
	dm_build_39[dm_build_40] = dm_build_41
	return 1
}

func (Dm_build_43 *dm_build_0) Dm_build_42(dm_build_44 []byte, dm_build_45 int, dm_build_46 uint16) int {
	dm_build_44[dm_build_45] = byte(dm_build_46)
	dm_build_45++
	dm_build_44[dm_build_45] = byte(dm_build_46 >> 8)
	return 2
}

func (Dm_build_48 *dm_build_0) Dm_build_47(dm_build_49 []byte, dm_build_50 int, dm_build_51 uint32) int {
	dm_build_49[dm_build_50] = byte(dm_build_51)
	dm_build_50++
	dm_build_49[dm_build_50] = byte(dm_build_51 >> 8)
	dm_build_50++
	dm_build_49[dm_build_50] = byte(dm_build_51 >> 16)
	dm_build_50++
	dm_build_49[dm_build_50] = byte(dm_build_51 >> 24)
	return 3
}

func (Dm_build_53 *dm_build_0) Dm_build_52(dm_build_54 []byte, dm_build_55 int, dm_build_56 uint64) int {
	dm_build_54[dm_build_55] = byte(dm_build_56)
	dm_build_55++
	dm_build_54[dm_build_55] = byte(dm_build_56 >> 8)
	dm_build_55++
	dm_build_54[dm_build_55] = byte(dm_build_56 >> 16)
	dm_build_55++
	dm_build_54[dm_build_55] = byte(dm_build_56 >> 24)
	dm_build_55++
	dm_build_54[dm_build_55] = byte(dm_build_56 >> 32)
	dm_build_55++
	dm_build_54[dm_build_55] = byte(dm_build_56 >> 40)
	dm_build_55++
	dm_build_54[dm_build_55] = byte(dm_build_56 >> 48)
	dm_build_55++
	dm_build_54[dm_build_55] = byte(dm_build_56 >> 56)
	return 3
}

func (Dm_build_58 *dm_build_0) Dm_build_57(dm_build_59 []byte, dm_build_60 int, dm_build_61 []byte, dm_build_62 int, dm_build_63 int) int {
	copy(dm_build_59[dm_build_60:dm_build_60+dm_build_63], dm_build_61[dm_build_62:dm_build_62+dm_build_63])
	return dm_build_63
}

func (Dm_build_65 *dm_build_0) Dm_build_64(dm_build_66 []byte, dm_build_67 int, dm_build_68 []byte, dm_build_69 int, dm_build_70 int) int {
	dm_build_67 += Dm_build_65.Dm_build_47(dm_build_66, dm_build_67, uint32(dm_build_70))
	return 4 + Dm_build_65.Dm_build_57(dm_build_66, dm_build_67, dm_build_68, dm_build_69, dm_build_70)
}

func (Dm_build_72 *dm_build_0) Dm_build_71(dm_build_73 []byte, dm_build_74 int, dm_build_75 []byte, dm_build_76 int, dm_build_77 int) int {
	dm_build_74 += Dm_build_72.Dm_build_42(dm_build_73, dm_build_74, uint16(dm_build_77))
	return 2 + Dm_build_72.Dm_build_57(dm_build_73, dm_build_74, dm_build_75, dm_build_76, dm_build_77)
}

func (Dm_build_79 *dm_build_0) Dm_build_78(dm_build_80 []byte, dm_build_81 int, dm_build_82 string, dm_build_83 string, dm_build_84 *DmConnection) int {
	dm_build_85 := Dm_build_79.Dm_build_217(dm_build_82, dm_build_83, dm_build_84)
	dm_build_81 += Dm_build_79.Dm_build_47(dm_build_80, dm_build_81, uint32(len(dm_build_85)))
	return 4 + Dm_build_79.Dm_build_57(dm_build_80, dm_build_81, dm_build_85, 0, len(dm_build_85))
}

func (Dm_build_87 *dm_build_0) Dm_build_86(dm_build_88 []byte, dm_build_89 int, dm_build_90 string, dm_build_91 string, dm_build_92 *DmConnection) int {
	dm_build_93 := Dm_build_87.Dm_build_217(dm_build_90, dm_build_91, dm_build_92)

	dm_build_89 += Dm_build_87.Dm_build_42(dm_build_88, dm_build_89, uint16(len(dm_build_93)))
	return 2 + Dm_build_87.Dm_build_57(dm_build_88, dm_build_89, dm_build_93, 0, len(dm_build_93))
}

func (Dm_build_95 *dm_build_0) Dm_build_94(dm_build_96 []byte, dm_build_97 int) byte {
	return dm_build_96[dm_build_97]
}

func (Dm_build_99 *dm_build_0) Dm_build_98(dm_build_100 []byte, dm_build_101 int) int16 {
	var dm_build_102 int16
	dm_build_102 = int16(dm_build_100[dm_build_101] & 0xff)
	dm_build_101++
	dm_build_102 |= int16(dm_build_100[dm_build_101]&0xff) << 8
	return dm_build_102
}

func (Dm_build_104 *dm_build_0) Dm_build_103(dm_build_105 []byte, dm_build_106 int) int32 {
	var dm_build_107 int32
	dm_build_107 = int32(dm_build_105[dm_build_106] & 0xff)
	dm_build_106++
	dm_build_107 |= int32(dm_build_105[dm_build_106]&0xff) << 8
	dm_build_106++
	dm_build_107 |= int32(dm_build_105[dm_build_106]&0xff) << 16
	dm_build_106++
	dm_build_107 |= int32(dm_build_105[dm_build_106]&0xff) << 24
	return dm_build_107
}

func (Dm_build_109 *dm_build_0) Dm_build_108(dm_build_110 []byte, dm_build_111 int) int64 {
	var dm_build_112 int64
	dm_build_112 = int64(dm_build_110[dm_build_111] & 0xff)
	dm_build_111++
	dm_build_112 |= int64(dm_build_110[dm_build_111]&0xff) << 8
	dm_build_111++
	dm_build_112 |= int64(dm_build_110[dm_build_111]&0xff) << 16
	dm_build_111++
	dm_build_112 |= int64(dm_build_110[dm_build_111]&0xff) << 24
	dm_build_111++
	dm_build_112 |= int64(dm_build_110[dm_build_111]&0xff) << 32
	dm_build_111++
	dm_build_112 |= int64(dm_build_110[dm_build_111]&0xff) << 40
	dm_build_111++
	dm_build_112 |= int64(dm_build_110[dm_build_111]&0xff) << 48
	dm_build_111++
	dm_build_112 |= int64(dm_build_110[dm_build_111]&0xff) << 56
	return dm_build_112
}

func (Dm_build_114 *dm_build_0) Dm_build_113(dm_build_115 []byte, dm_build_116 int) float32 {
	return math.Float32frombits(Dm_build_114.Dm_build_130(dm_build_115, dm_build_116))
}

func (Dm_build_118 *dm_build_0) Dm_build_117(dm_build_119 []byte, dm_build_120 int) float64 {
	return math.Float64frombits(Dm_build_118.Dm_build_135(dm_build_119, dm_build_120))
}

func (Dm_build_122 *dm_build_0) Dm_build_121(dm_build_123 []byte, dm_build_124 int) uint8 {
	return dm_build_123[dm_build_124] & 0xff
}

func (Dm_build_126 *dm_build_0) Dm_build_125(dm_build_127 []byte, dm_build_128 int) uint16 {
	var dm_build_129 uint16
	dm_build_129 = uint16(dm_build_127[dm_build_128] & 0xff)
	dm_build_128++
	dm_build_129 |= uint16(dm_build_127[dm_build_128]&0xff) << 8
	return dm_build_129
}

func (Dm_build_131 *dm_build_0) Dm_build_130(dm_build_132 []byte, dm_build_133 int) uint32 {
	var dm_build_134 uint32
	dm_build_134 = uint32(dm_build_132[dm_build_133] & 0xff)
	dm_build_133++
	dm_build_134 |= uint32(dm_build_132[dm_build_133]&0xff) << 8
	dm_build_133++
	dm_build_134 |= uint32(dm_build_132[dm_build_133]&0xff) << 16
	dm_build_133++
	dm_build_134 |= uint32(dm_build_132[dm_build_133]&0xff) << 24
	return dm_build_134
}

func (Dm_build_136 *dm_build_0) Dm_build_135(dm_build_137 []byte, dm_build_138 int) uint64 {
	var dm_build_139 uint64
	dm_build_139 = uint64(dm_build_137[dm_build_138] & 0xff)
	dm_build_138++
	dm_build_139 |= uint64(dm_build_137[dm_build_138]&0xff) << 8
	dm_build_138++
	dm_build_139 |= uint64(dm_build_137[dm_build_138]&0xff) << 16
	dm_build_138++
	dm_build_139 |= uint64(dm_build_137[dm_build_138]&0xff) << 24
	dm_build_138++
	dm_build_139 |= uint64(dm_build_137[dm_build_138]&0xff) << 32
	dm_build_138++
	dm_build_139 |= uint64(dm_build_137[dm_build_138]&0xff) << 40
	dm_build_138++
	dm_build_139 |= uint64(dm_build_137[dm_build_138]&0xff) << 48
	dm_build_138++
	dm_build_139 |= uint64(dm_build_137[dm_build_138]&0xff) << 56
	return dm_build_139
}

func (Dm_build_141 *dm_build_0) Dm_build_140(dm_build_142 []byte, dm_build_143 int) []byte {
	dm_build_144 := Dm_build_141.Dm_build_130(dm_build_142, dm_build_143)

	dm_build_145 := make([]byte, dm_build_144)
	copy(dm_build_145[:int(dm_build_144)], dm_build_142[dm_build_143+4:dm_build_143+4+int(dm_build_144)])
	return dm_build_145
}

func (Dm_build_147 *dm_build_0) Dm_build_146(dm_build_148 []byte, dm_build_149 int) []byte {
	dm_build_150 := Dm_build_147.Dm_build_125(dm_build_148, dm_build_149)

	dm_build_151 := make([]byte, dm_build_150)
	copy(dm_build_151[:int(dm_build_150)], dm_build_148[dm_build_149+2:dm_build_149+2+int(dm_build_150)])
	return dm_build_151
}

func (Dm_build_153 *dm_build_0) Dm_build_152(dm_build_154 []byte, dm_build_155 int, dm_build_156 int) []byte {

	dm_build_157 := make([]byte, dm_build_156)
	copy(dm_build_157[:dm_build_156], dm_build_154[dm_build_155:dm_build_155+dm_build_156])
	return dm_build_157
}

func (Dm_build_159 *dm_build_0) Dm_build_158(dm_build_160 []byte, dm_build_161 int, dm_build_162 int, dm_build_163 string, dm_build_164 *DmConnection) string {
	return Dm_build_159.Dm_build_253(dm_build_160[dm_build_161:dm_build_161+dm_build_162], dm_build_163, dm_build_164)
}

func (Dm_build_166 *dm_build_0) Dm_build_165(dm_build_167 []byte, dm_build_168 int, dm_build_169 string, dm_build_170 *DmConnection) string {
	dm_build_171 := Dm_build_166.Dm_build_130(dm_build_167, dm_build_168)
	dm_build_168 += 4
	return Dm_build_166.Dm_build_158(dm_build_167, dm_build_168, int(dm_build_171), dm_build_169, dm_build_170)
}

func (Dm_build_173 *dm_build_0) Dm_build_172(dm_build_174 []byte, dm_build_175 int, dm_build_176 string, dm_build_177 *DmConnection) string {
	dm_build_178 := Dm_build_173.Dm_build_125(dm_build_174, dm_build_175)
	dm_build_175 += 2
	return Dm_build_173.Dm_build_158(dm_build_174, dm_build_175, int(dm_build_178), dm_build_176, dm_build_177)
}

func (Dm_build_180 *dm_build_0) Dm_build_179(dm_build_181 byte) []byte {
	return []byte{dm_build_181}
}

func (Dm_build_183 *dm_build_0) Dm_build_182(dm_build_184 int8) []byte {
	return []byte{byte(dm_build_184)}
}

func (Dm_build_186 *dm_build_0) Dm_build_185(dm_build_187 int16) []byte {
	return []byte{byte(dm_build_187), byte(dm_build_187 >> 8)}
}

func (Dm_build_189 *dm_build_0) Dm_build_188(dm_build_190 int32) []byte {
	return []byte{byte(dm_build_190), byte(dm_build_190 >> 8), byte(dm_build_190 >> 16), byte(dm_build_190 >> 24)}
}

func (Dm_build_192 *dm_build_0) Dm_build_191(dm_build_193 int64) []byte {
	return []byte{byte(dm_build_193), byte(dm_build_193 >> 8), byte(dm_build_193 >> 16), byte(dm_build_193 >> 24), byte(dm_build_193 >> 32),
		byte(dm_build_193 >> 40), byte(dm_build_193 >> 48), byte(dm_build_193 >> 56)}
}

func (Dm_build_195 *dm_build_0) Dm_build_194(dm_build_196 float32) []byte {
	return Dm_build_195.Dm_build_206(math.Float32bits(dm_build_196))
}

func (Dm_build_198 *dm_build_0) Dm_build_197(dm_build_199 float64) []byte {
	return Dm_build_198.Dm_build_209(math.Float64bits(dm_build_199))
}

func (Dm_build_201 *dm_build_0) Dm_build_200(dm_build_202 uint8) []byte {
	return []byte{dm_build_202}
}

func (Dm_build_204 *dm_build_0) Dm_build_203(dm_build_205 uint16) []byte {
	return []byte{byte(dm_build_205), byte(dm_build_205 >> 8)}
}

func (Dm_build_207 *dm_build_0) Dm_build_206(dm_build_208 uint32) []byte {
	return []byte{byte(dm_build_208), byte(dm_build_208 >> 8), byte(dm_build_208 >> 16), byte(dm_build_208 >> 24)}
}

func (Dm_build_210 *dm_build_0) Dm_build_209(dm_build_211 uint64) []byte {
	return []byte{byte(dm_build_211), byte(dm_build_211 >> 8), byte(dm_build_211 >> 16), byte(dm_build_211 >> 24), byte(dm_build_211 >> 32), byte(dm_build_211 >> 40), byte(dm_build_211 >> 48), byte(dm_build_211 >> 56)}
}

func (Dm_build_213 *dm_build_0) Dm_build_212(dm_build_214 []byte, dm_build_215 string, dm_build_216 *DmConnection) []byte {
	if dm_build_215 == "UTF-8" {
		return dm_build_214
	}

	if dm_build_216 == nil {
		if e := dm_build_258(dm_build_215); e != nil {
			tmp, err := io.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_214), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_216.encodeBuffer == nil {
		dm_build_216.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_216.encode = dm_build_258(dm_build_216.getServerEncoding())
		dm_build_216.transformReaderDst = make([]byte, 4096)
		dm_build_216.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_216.encode; e != nil {

		dm_build_216.encodeBuffer.Reset()

		n, err := dm_build_216.encodeBuffer.ReadFrom(
			Dm_build_272(bytes.NewReader(dm_build_214), e.NewEncoder(), dm_build_216.transformReaderDst, dm_build_216.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_216.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_218 *dm_build_0) Dm_build_217(dm_build_219 string, dm_build_220 string, dm_build_221 *DmConnection) []byte {
	return Dm_build_218.Dm_build_212([]byte(dm_build_219), dm_build_220, dm_build_221)
}

func (Dm_build_223 *dm_build_0) Dm_build_222(dm_build_224 []byte) byte {
	return Dm_build_223.Dm_build_94(dm_build_224, 0)
}

func (Dm_build_226 *dm_build_0) Dm_build_225(dm_build_227 []byte) int16 {
	return Dm_build_226.Dm_build_98(dm_build_227, 0)
}

func (Dm_build_229 *dm_build_0) Dm_build_228(dm_build_230 []byte) int32 {
	return Dm_build_229.Dm_build_103(dm_build_230, 0)
}

func (Dm_build_232 *dm_build_0) Dm_build_231(dm_build_233 []byte) int64 {
	return Dm_build_232.Dm_build_108(dm_build_233, 0)
}

func (Dm_build_235 *dm_build_0) Dm_build_234(dm_build_236 []byte) float32 {
	return Dm_build_235.Dm_build_113(dm_build_236, 0)
}

func (Dm_build_238 *dm_build_0) Dm_build_237(dm_build_239 []byte) float64 {
	return Dm_build_238.Dm_build_117(dm_build_239, 0)
}

func (Dm_build_241 *dm_build_0) Dm_build_240(dm_build_242 []byte) uint8 {
	return Dm_build_241.Dm_build_121(dm_build_242, 0)
}

func (Dm_build_244 *dm_build_0) Dm_build_243(dm_build_245 []byte) uint16 {
	return Dm_build_244.Dm_build_125(dm_build_245, 0)
}

func (Dm_build_247 *dm_build_0) Dm_build_246(dm_build_248 []byte) uint32 {
	return Dm_build_247.Dm_build_130(dm_build_248, 0)
}

func (Dm_build_250 *dm_build_0) Dm_build_249(dm_build_251 []byte, dm_build_252 string) []byte {
	if dm_build_252 == "UTF-8" {
		return dm_build_251
	}

	if e := dm_build_258(dm_build_252); e != nil {

		tmp, err := io.ReadAll(
			transform.NewReader(bytes.NewReader(dm_build_251), e.NewDecoder()),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return tmp
	}

	panic("Unsupported Charset!")

}

func (Dm_build_254 *dm_build_0) Dm_build_253(dm_build_255 []byte, dm_build_256 string, _ *DmConnection) string {
	return string(Dm_build_254.Dm_build_249(dm_build_255, dm_build_256))
}

func dm_build_258(dm_build_259 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_259); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_260 struct {
	dm_build_261 io.Reader
	dm_build_262 transform.Transformer
	dm_build_263 error

	dm_build_264               []byte
	dm_build_265, dm_build_266 int

	dm_build_267               []byte
	dm_build_268, dm_build_269 int

	dm_build_270 bool
}

const _ = 4096

func Dm_build_272(dm_build_273 io.Reader, dm_build_274 transform.Transformer, dm_build_275 []byte, dm_build_276 []byte) *Dm_build_260 {
	dm_build_274.Reset()
	return &Dm_build_260{
		dm_build_261: dm_build_273,
		dm_build_262: dm_build_274,
		dm_build_264: dm_build_275,
		dm_build_267: dm_build_276,
	}
}

func (dm_build_278 *Dm_build_260) Read(dm_build_279 []byte) (int, error) {
	dm_build_280, dm_build_281 := 0, error(nil)
	for {

		if dm_build_278.dm_build_265 != dm_build_278.dm_build_266 {
			dm_build_280 = copy(dm_build_279, dm_build_278.dm_build_264[dm_build_278.dm_build_265:dm_build_278.dm_build_266])
			dm_build_278.dm_build_265 += dm_build_280
			if dm_build_278.dm_build_265 == dm_build_278.dm_build_266 && dm_build_278.dm_build_270 {
				return dm_build_280, dm_build_278.dm_build_263
			}
			return dm_build_280, nil
		} else if dm_build_278.dm_build_270 {
			return 0, dm_build_278.dm_build_263
		}

		if dm_build_278.dm_build_268 != dm_build_278.dm_build_269 || dm_build_278.dm_build_263 != nil {
			dm_build_278.dm_build_265 = 0
			dm_build_278.dm_build_266, dm_build_280, dm_build_281 = dm_build_278.dm_build_262.Transform(dm_build_278.dm_build_264, dm_build_278.dm_build_267[dm_build_278.dm_build_268:dm_build_278.dm_build_269], dm_build_278.dm_build_263 == io.EOF)
			dm_build_278.dm_build_268 += dm_build_280

			switch {
			case dm_build_281 == nil:
				if dm_build_278.dm_build_268 != dm_build_278.dm_build_269 {
					dm_build_278.dm_build_263 = nil
				}

				dm_build_278.dm_build_270 = dm_build_278.dm_build_263 != nil
				continue
			case errors.Is(dm_build_281, transform.ErrShortDst) && (dm_build_278.dm_build_266 != 0 || dm_build_280 != 0):

				continue
			case errors.Is(dm_build_281, transform.ErrShortSrc) && dm_build_278.dm_build_269-dm_build_278.dm_build_268 != len(dm_build_278.dm_build_267) && dm_build_278.dm_build_263 == nil:

			default:
				dm_build_278.dm_build_270 = true

				if dm_build_278.dm_build_263 == nil || dm_build_278.dm_build_263 == io.EOF {
					dm_build_278.dm_build_263 = dm_build_281
				}
				continue
			}
		}

		if dm_build_278.dm_build_268 != 0 {
			dm_build_278.dm_build_268, dm_build_278.dm_build_269 = 0, copy(dm_build_278.dm_build_267, dm_build_278.dm_build_267[dm_build_278.dm_build_268:dm_build_278.dm_build_269])
		}
		dm_build_280, dm_build_278.dm_build_263 = dm_build_278.dm_build_261.Read(dm_build_278.dm_build_267[dm_build_278.dm_build_269:])
		dm_build_278.dm_build_269 += dm_build_280
	}
}
