/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */

package dm8

import (
	"io"
	"math"
)

type Dm_build_78 struct {
	dm_build_79 []byte
	dm_build_80 int
}

func Dm_build_81(dm_build_82 int) *Dm_build_78 {
	return &Dm_build_78{make([]byte, 0, dm_build_82), 0}
}

func Dm_build_83(dm_build_84 []byte) *Dm_build_78 {
	return &Dm_build_78{dm_build_84, 0}
}

func (dm_build_86 *Dm_build_78) dm_build_85(dm_build_87 int) *Dm_build_78 {

	dm_build_88 := len(dm_build_86.dm_build_79)
	dm_build_89 := cap(dm_build_86.dm_build_79)

	if dm_build_88+dm_build_87 <= dm_build_89 {
		dm_build_86.dm_build_79 = dm_build_86.dm_build_79[:dm_build_88+dm_build_87]
	} else {

		var calCap = int64(math.Max(float64(2*dm_build_89), float64(dm_build_87+dm_build_88)))

		nbuf := make([]byte, dm_build_87+dm_build_88, calCap)
		copy(nbuf, dm_build_86.dm_build_79)
		dm_build_86.dm_build_79 = nbuf
	}

	return dm_build_86
}

func (dm_build_91 *Dm_build_78) Dm_build_90() int {
	return len(dm_build_91.dm_build_79)
}

func (dm_build_93 *Dm_build_78) Dm_build_92(dm_build_94 int) *Dm_build_78 {
	for i := dm_build_94; i < len(dm_build_93.dm_build_79); i++ {
		dm_build_93.dm_build_79[i] = 0
	}
	dm_build_93.dm_build_79 = dm_build_93.dm_build_79[:dm_build_94]
	return dm_build_93
}

func (dm_build_96 *Dm_build_78) Dm_build_95(dm_build_97 int) *Dm_build_78 {
	dm_build_96.dm_build_80 = dm_build_97
	return dm_build_96
}

func (dm_build_99 *Dm_build_78) Dm_build_98() int {
	return dm_build_99.dm_build_80
}

func (dm_build_101 *Dm_build_78) Dm_build_100(dm_build_102 bool) int {
	return len(dm_build_101.dm_build_79) - dm_build_101.dm_build_80
}

func (dm_build_104 *Dm_build_78) Dm_build_103(dm_build_105 int, dm_build_106 bool, dm_build_107 bool) *Dm_build_78 {

	if dm_build_106 {
		if dm_build_107 {
			dm_build_104.dm_build_85(dm_build_105)
		} else {
			dm_build_104.dm_build_79 = dm_build_104.dm_build_79[:len(dm_build_104.dm_build_79)-dm_build_105]
		}
	} else {
		if dm_build_107 {
			dm_build_104.dm_build_80 += dm_build_105
		} else {
			dm_build_104.dm_build_80 -= dm_build_105
		}
	}

	return dm_build_104
}

func (dm_build_109 *Dm_build_78) Dm_build_108(dm_build_110 io.Reader, dm_build_111 int) (int, error) {
	dm_build_112 := len(dm_build_109.dm_build_79)
	dm_build_109.dm_build_85(dm_build_111)
	dm_build_113 := 0
	for dm_build_111 > 0 {
		n, err := dm_build_110.Read(dm_build_109.dm_build_79[dm_build_112+dm_build_113:])
		if n > 0 && err == io.EOF {
			dm_build_113 += n
			dm_build_109.dm_build_79 = dm_build_109.dm_build_79[:dm_build_112+dm_build_113]
			return dm_build_113, nil
		} else if n > 0 && err == nil {
			dm_build_111 -= n
			dm_build_113 += n
		} else if n == 0 && err != nil {
			return -1, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
		}
	}

	return dm_build_113, nil
}

func (dm_build_115 *Dm_build_78) Dm_build_114(dm_build_116 io.Writer) (*Dm_build_78, error) {
	if _, err := dm_build_116.Write(dm_build_115.dm_build_79); err != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
	}
	return dm_build_115, nil
}

func (dm_build_118 *Dm_build_78) Dm_build_117(dm_build_119 bool) int {
	dm_build_120 := len(dm_build_118.dm_build_79)
	dm_build_118.dm_build_85(1)

	if dm_build_119 {
		return copy(dm_build_118.dm_build_79[dm_build_120:], []byte{1})
	} else {
		return copy(dm_build_118.dm_build_79[dm_build_120:], []byte{0})
	}
}

func (dm_build_122 *Dm_build_78) Dm_build_121(dm_build_123 byte) int {
	dm_build_124 := len(dm_build_122.dm_build_79)
	dm_build_122.dm_build_85(1)

	return copy(dm_build_122.dm_build_79[dm_build_124:], Dm_build_1331.Dm_build_1509(dm_build_123))
}

func (dm_build_126 *Dm_build_78) Dm_build_125(dm_build_127 int8) int {
	dm_build_128 := len(dm_build_126.dm_build_79)
	dm_build_126.dm_build_85(1)

	return copy(dm_build_126.dm_build_79[dm_build_128:], Dm_build_1331.Dm_build_1512(dm_build_127))
}

func (dm_build_130 *Dm_build_78) Dm_build_129(dm_build_131 int16) int {
	dm_build_132 := len(dm_build_130.dm_build_79)
	dm_build_130.dm_build_85(2)

	return copy(dm_build_130.dm_build_79[dm_build_132:], Dm_build_1331.Dm_build_1515(dm_build_131))
}

func (dm_build_134 *Dm_build_78) Dm_build_133(dm_build_135 int32) int {
	dm_build_136 := len(dm_build_134.dm_build_79)
	dm_build_134.dm_build_85(4)

	return copy(dm_build_134.dm_build_79[dm_build_136:], Dm_build_1331.Dm_build_1518(dm_build_135))
}

func (dm_build_138 *Dm_build_78) Dm_build_137(dm_build_139 uint8) int {
	dm_build_140 := len(dm_build_138.dm_build_79)
	dm_build_138.dm_build_85(1)

	return copy(dm_build_138.dm_build_79[dm_build_140:], Dm_build_1331.Dm_build_1530(dm_build_139))
}

func (dm_build_142 *Dm_build_78) Dm_build_141(dm_build_143 uint16) int {
	dm_build_144 := len(dm_build_142.dm_build_79)
	dm_build_142.dm_build_85(2)

	return copy(dm_build_142.dm_build_79[dm_build_144:], Dm_build_1331.Dm_build_1533(dm_build_143))
}

func (dm_build_146 *Dm_build_78) Dm_build_145(dm_build_147 uint32) int {
	dm_build_148 := len(dm_build_146.dm_build_79)
	dm_build_146.dm_build_85(4)

	return copy(dm_build_146.dm_build_79[dm_build_148:], Dm_build_1331.Dm_build_1536(dm_build_147))
}

func (dm_build_150 *Dm_build_78) Dm_build_149(dm_build_151 uint64) int {
	dm_build_152 := len(dm_build_150.dm_build_79)
	dm_build_150.dm_build_85(8)

	return copy(dm_build_150.dm_build_79[dm_build_152:], Dm_build_1331.Dm_build_1539(dm_build_151))
}

func (dm_build_154 *Dm_build_78) Dm_build_153(dm_build_155 float32) int {
	dm_build_156 := len(dm_build_154.dm_build_79)
	dm_build_154.dm_build_85(4)

	return copy(dm_build_154.dm_build_79[dm_build_156:], Dm_build_1331.Dm_build_1536(math.Float32bits(dm_build_155)))
}

func (dm_build_158 *Dm_build_78) Dm_build_157(dm_build_159 float64) int {
	dm_build_160 := len(dm_build_158.dm_build_79)
	dm_build_158.dm_build_85(8)

	return copy(dm_build_158.dm_build_79[dm_build_160:], Dm_build_1331.Dm_build_1539(math.Float64bits(dm_build_159)))
}

func (dm_build_162 *Dm_build_78) Dm_build_161(dm_build_163 []byte) int {
	dm_build_164 := len(dm_build_162.dm_build_79)
	dm_build_162.dm_build_85(len(dm_build_163))
	return copy(dm_build_162.dm_build_79[dm_build_164:], dm_build_163)
}

func (dm_build_166 *Dm_build_78) Dm_build_165(dm_build_167 []byte) int {
	return dm_build_166.Dm_build_133(int32(len(dm_build_167))) + dm_build_166.Dm_build_161(dm_build_167)
}

func (dm_build_169 *Dm_build_78) Dm_build_168(dm_build_170 []byte) int {
	return dm_build_169.Dm_build_137(uint8(len(dm_build_170))) + dm_build_169.Dm_build_161(dm_build_170)
}

func (dm_build_172 *Dm_build_78) Dm_build_171(dm_build_173 []byte) int {
	return dm_build_172.Dm_build_141(uint16(len(dm_build_173))) + dm_build_172.Dm_build_161(dm_build_173)
}

func (dm_build_175 *Dm_build_78) Dm_build_174(dm_build_176 []byte) int {
	return dm_build_175.Dm_build_161(dm_build_176) + dm_build_175.Dm_build_121(0)
}

func (dm_build_178 *Dm_build_78) Dm_build_177(dm_build_179 string, dm_build_180 string, dm_build_181 *DmConnection) int {
	dm_build_182 := Dm_build_1331.Dm_build_1547(dm_build_179, dm_build_180, dm_build_181)
	return dm_build_178.Dm_build_165(dm_build_182)
}

func (dm_build_184 *Dm_build_78) Dm_build_183(dm_build_185 string, dm_build_186 string, dm_build_187 *DmConnection) int {
	dm_build_188 := Dm_build_1331.Dm_build_1547(dm_build_185, dm_build_186, dm_build_187)
	return dm_build_184.Dm_build_168(dm_build_188)
}

func (dm_build_190 *Dm_build_78) Dm_build_189(dm_build_191 string, dm_build_192 string, dm_build_193 *DmConnection) int {
	dm_build_194 := Dm_build_1331.Dm_build_1547(dm_build_191, dm_build_192, dm_build_193)
	return dm_build_190.Dm_build_171(dm_build_194)
}

func (dm_build_196 *Dm_build_78) Dm_build_195(dm_build_197 string, dm_build_198 string, dm_build_199 *DmConnection) int {
	dm_build_200 := Dm_build_1331.Dm_build_1547(dm_build_197, dm_build_198, dm_build_199)
	return dm_build_196.Dm_build_174(dm_build_200)
}

func (dm_build_202 *Dm_build_78) Dm_build_201() byte {
	dm_build_203 := Dm_build_1331.Dm_build_1424(dm_build_202.dm_build_79, dm_build_202.dm_build_80)
	dm_build_202.dm_build_80++
	return dm_build_203
}

func (dm_build_205 *Dm_build_78) Dm_build_204() int16 {
	dm_build_206 := Dm_build_1331.Dm_build_1428(dm_build_205.dm_build_79, dm_build_205.dm_build_80)
	dm_build_205.dm_build_80 += 2
	return dm_build_206
}

func (dm_build_208 *Dm_build_78) Dm_build_207() int32 {
	dm_build_209 := Dm_build_1331.Dm_build_1433(dm_build_208.dm_build_79, dm_build_208.dm_build_80)
	dm_build_208.dm_build_80 += 4
	return dm_build_209
}

func (dm_build_211 *Dm_build_78) Dm_build_210() int64 {
	dm_build_212 := Dm_build_1331.Dm_build_1438(dm_build_211.dm_build_79, dm_build_211.dm_build_80)
	dm_build_211.dm_build_80 += 8
	return dm_build_212
}

func (dm_build_214 *Dm_build_78) Dm_build_213() float32 {
	dm_build_215 := Dm_build_1331.Dm_build_1443(dm_build_214.dm_build_79, dm_build_214.dm_build_80)
	dm_build_214.dm_build_80 += 4
	return dm_build_215
}

func (dm_build_217 *Dm_build_78) Dm_build_216() float64 {
	dm_build_218 := Dm_build_1331.Dm_build_1447(dm_build_217.dm_build_79, dm_build_217.dm_build_80)
	dm_build_217.dm_build_80 += 8
	return dm_build_218
}

func (dm_build_220 *Dm_build_78) Dm_build_219() uint8 {
	dm_build_221 := Dm_build_1331.Dm_build_1451(dm_build_220.dm_build_79, dm_build_220.dm_build_80)
	dm_build_220.dm_build_80 += 1
	return dm_build_221
}

func (dm_build_223 *Dm_build_78) Dm_build_222() uint16 {
	dm_build_224 := Dm_build_1331.Dm_build_1455(dm_build_223.dm_build_79, dm_build_223.dm_build_80)
	dm_build_223.dm_build_80 += 2
	return dm_build_224
}

func (dm_build_226 *Dm_build_78) Dm_build_225() uint32 {
	dm_build_227 := Dm_build_1331.Dm_build_1460(dm_build_226.dm_build_79, dm_build_226.dm_build_80)
	dm_build_226.dm_build_80 += 4
	return dm_build_227
}

func (dm_build_229 *Dm_build_78) Dm_build_228(dm_build_230 int) []byte {
	dm_build_231 := Dm_build_1331.Dm_build_1482(dm_build_229.dm_build_79, dm_build_229.dm_build_80, dm_build_230)
	dm_build_229.dm_build_80 += dm_build_230
	return dm_build_231
}

func (dm_build_233 *Dm_build_78) Dm_build_232() []byte {
	return dm_build_233.Dm_build_228(int(dm_build_233.Dm_build_207()))
}

func (dm_build_235 *Dm_build_78) Dm_build_234() []byte {
	return dm_build_235.Dm_build_228(int(dm_build_235.Dm_build_201()))
}

func (dm_build_237 *Dm_build_78) Dm_build_236() []byte {
	return dm_build_237.Dm_build_228(int(dm_build_237.Dm_build_204()))
}

func (dm_build_239 *Dm_build_78) Dm_build_238(dm_build_240 int) []byte {
	return dm_build_239.Dm_build_228(dm_build_240)
}

func (dm_build_242 *Dm_build_78) Dm_build_241() []byte {
	dm_build_243 := 0
	for dm_build_242.Dm_build_201() != 0 {
		dm_build_243++
	}
	dm_build_242.Dm_build_103(dm_build_243, false, false)
	return dm_build_242.Dm_build_228(dm_build_243)
}

func (dm_build_245 *Dm_build_78) Dm_build_244(dm_build_246 int, dm_build_247 string, dm_build_248 *DmConnection) string {
	return Dm_build_1331.Dm_build_1584(dm_build_245.Dm_build_228(dm_build_246), dm_build_247, dm_build_248)
}

func (dm_build_250 *Dm_build_78) Dm_build_249(dm_build_251 string, dm_build_252 *DmConnection) string {
	return Dm_build_1331.Dm_build_1584(dm_build_250.Dm_build_232(), dm_build_251, dm_build_252)
}

func (dm_build_254 *Dm_build_78) Dm_build_253(dm_build_255 string, dm_build_256 *DmConnection) string {
	return Dm_build_1331.Dm_build_1584(dm_build_254.Dm_build_234(), dm_build_255, dm_build_256)
}

func (dm_build_258 *Dm_build_78) Dm_build_257(dm_build_259 string, dm_build_260 *DmConnection) string {
	return Dm_build_1331.Dm_build_1584(dm_build_258.Dm_build_236(), dm_build_259, dm_build_260)
}

func (dm_build_262 *Dm_build_78) Dm_build_261(dm_build_263 string, dm_build_264 *DmConnection) string {
	return Dm_build_1331.Dm_build_1584(dm_build_262.Dm_build_241(), dm_build_263, dm_build_264)
}

func (dm_build_266 *Dm_build_78) Dm_build_265(dm_build_267 int, dm_build_268 byte) int {
	return dm_build_266.Dm_build_301(dm_build_267, Dm_build_1331.Dm_build_1509(dm_build_268))
}

func (dm_build_270 *Dm_build_78) Dm_build_269(dm_build_271 int, dm_build_272 int16) int {
	return dm_build_270.Dm_build_301(dm_build_271, Dm_build_1331.Dm_build_1515(dm_build_272))
}

func (dm_build_274 *Dm_build_78) Dm_build_273(dm_build_275 int, dm_build_276 int32) int {
	return dm_build_274.Dm_build_301(dm_build_275, Dm_build_1331.Dm_build_1518(dm_build_276))
}

func (dm_build_278 *Dm_build_78) Dm_build_277(dm_build_279 int, dm_build_280 int64) int {
	return dm_build_278.Dm_build_301(dm_build_279, Dm_build_1331.Dm_build_1521(dm_build_280))
}

func (dm_build_282 *Dm_build_78) Dm_build_281(dm_build_283 int, dm_build_284 float32) int {
	return dm_build_282.Dm_build_301(dm_build_283, Dm_build_1331.Dm_build_1524(dm_build_284))
}

func (dm_build_286 *Dm_build_78) Dm_build_285(dm_build_287 int, dm_build_288 float64) int {
	return dm_build_286.Dm_build_301(dm_build_287, Dm_build_1331.Dm_build_1527(dm_build_288))
}

func (dm_build_290 *Dm_build_78) Dm_build_289(dm_build_291 int, dm_build_292 uint8) int {
	return dm_build_290.Dm_build_301(dm_build_291, Dm_build_1331.Dm_build_1530(dm_build_292))
}

func (dm_build_294 *Dm_build_78) Dm_build_293(dm_build_295 int, dm_build_296 uint16) int {
	return dm_build_294.Dm_build_301(dm_build_295, Dm_build_1331.Dm_build_1533(dm_build_296))
}

func (dm_build_298 *Dm_build_78) Dm_build_297(dm_build_299 int, dm_build_300 uint32) int {
	return dm_build_298.Dm_build_301(dm_build_299, Dm_build_1331.Dm_build_1536(dm_build_300))
}

func (dm_build_302 *Dm_build_78) Dm_build_301(dm_build_303 int, dm_build_304 []byte) int {
	return copy(dm_build_302.dm_build_79[dm_build_303:], dm_build_304)
}

func (dm_build_306 *Dm_build_78) Dm_build_305(dm_build_307 int, dm_build_308 []byte) int {
	return dm_build_306.Dm_build_273(dm_build_307, int32(len(dm_build_308))) + dm_build_306.Dm_build_301(dm_build_307+4, dm_build_308)
}

func (dm_build_310 *Dm_build_78) Dm_build_309(dm_build_311 int, dm_build_312 []byte) int {
	return dm_build_310.Dm_build_265(dm_build_311, byte(len(dm_build_312))) + dm_build_310.Dm_build_301(dm_build_311+1, dm_build_312)
}

func (dm_build_314 *Dm_build_78) Dm_build_313(dm_build_315 int, dm_build_316 []byte) int {
	return dm_build_314.Dm_build_269(dm_build_315, int16(len(dm_build_316))) + dm_build_314.Dm_build_301(dm_build_315+2, dm_build_316)
}

func (dm_build_318 *Dm_build_78) Dm_build_317(dm_build_319 int, dm_build_320 []byte) int {
	return dm_build_318.Dm_build_301(dm_build_319, dm_build_320) + dm_build_318.Dm_build_265(dm_build_319+len(dm_build_320), 0)
}

func (dm_build_322 *Dm_build_78) Dm_build_321(dm_build_323 int, dm_build_324 string, dm_build_325 string, dm_build_326 *DmConnection) int {
	return dm_build_322.Dm_build_305(dm_build_323, Dm_build_1331.Dm_build_1547(dm_build_324, dm_build_325, dm_build_326))
}

func (dm_build_328 *Dm_build_78) Dm_build_327(dm_build_329 int, dm_build_330 string, dm_build_331 string, dm_build_332 *DmConnection) int {
	return dm_build_328.Dm_build_309(dm_build_329, Dm_build_1331.Dm_build_1547(dm_build_330, dm_build_331, dm_build_332))
}

func (dm_build_334 *Dm_build_78) Dm_build_333(dm_build_335 int, dm_build_336 string, dm_build_337 string, dm_build_338 *DmConnection) int {
	return dm_build_334.Dm_build_313(dm_build_335, Dm_build_1331.Dm_build_1547(dm_build_336, dm_build_337, dm_build_338))
}

func (dm_build_340 *Dm_build_78) Dm_build_339(dm_build_341 int, dm_build_342 string, dm_build_343 string, dm_build_344 *DmConnection) int {
	return dm_build_340.Dm_build_317(dm_build_341, Dm_build_1331.Dm_build_1547(dm_build_342, dm_build_343, dm_build_344))
}

func (dm_build_346 *Dm_build_78) Dm_build_345(dm_build_347 int) byte {
	return Dm_build_1331.Dm_build_1552(dm_build_346.Dm_build_372(dm_build_347, 1))
}

func (dm_build_349 *Dm_build_78) Dm_build_348(dm_build_350 int) int16 {
	return Dm_build_1331.Dm_build_1555(dm_build_349.Dm_build_372(dm_build_350, 2))
}

func (dm_build_352 *Dm_build_78) Dm_build_351(dm_build_353 int) int32 {
	return Dm_build_1331.Dm_build_1558(dm_build_352.Dm_build_372(dm_build_353, 4))
}

func (dm_build_355 *Dm_build_78) Dm_build_354(dm_build_356 int) int64 {
	return Dm_build_1331.Dm_build_1561(dm_build_355.Dm_build_372(dm_build_356, 8))
}

func (dm_build_358 *Dm_build_78) Dm_build_357(dm_build_359 int) float32 {
	return Dm_build_1331.Dm_build_1564(dm_build_358.Dm_build_372(dm_build_359, 4))
}

func (dm_build_361 *Dm_build_78) Dm_build_360(dm_build_362 int) float64 {
	return Dm_build_1331.Dm_build_1567(dm_build_361.Dm_build_372(dm_build_362, 8))
}

func (dm_build_364 *Dm_build_78) Dm_build_363(dm_build_365 int) uint8 {
	return Dm_build_1331.Dm_build_1570(dm_build_364.Dm_build_372(dm_build_365, 1))
}

func (dm_build_367 *Dm_build_78) Dm_build_366(dm_build_368 int) uint16 {
	return Dm_build_1331.Dm_build_1573(dm_build_367.Dm_build_372(dm_build_368, 2))
}

func (dm_build_370 *Dm_build_78) Dm_build_369(dm_build_371 int) uint32 {
	return Dm_build_1331.Dm_build_1576(dm_build_370.Dm_build_372(dm_build_371, 4))
}

func (dm_build_373 *Dm_build_78) Dm_build_372(dm_build_374 int, dm_build_375 int) []byte {
	return dm_build_373.dm_build_79[dm_build_374 : dm_build_374+dm_build_375]
}

func (dm_build_377 *Dm_build_78) Dm_build_376(dm_build_378 int) []byte {
	dm_build_379 := dm_build_377.Dm_build_351(dm_build_378)
	return dm_build_377.Dm_build_372(dm_build_378+4, int(dm_build_379))
}

func (dm_build_381 *Dm_build_78) Dm_build_380(dm_build_382 int) []byte {
	dm_build_383 := dm_build_381.Dm_build_345(dm_build_382)
	return dm_build_381.Dm_build_372(dm_build_382+1, int(dm_build_383))
}

func (dm_build_385 *Dm_build_78) Dm_build_384(dm_build_386 int) []byte {
	dm_build_387 := dm_build_385.Dm_build_348(dm_build_386)
	return dm_build_385.Dm_build_372(dm_build_386+2, int(dm_build_387))
}

func (dm_build_389 *Dm_build_78) Dm_build_388(dm_build_390 int) []byte {
	dm_build_391 := 0
	for dm_build_389.Dm_build_345(dm_build_390) != 0 {
		dm_build_390++
		dm_build_391++
	}

	return dm_build_389.Dm_build_372(dm_build_390-dm_build_391, int(dm_build_391))
}

func (dm_build_393 *Dm_build_78) Dm_build_392(dm_build_394 int, dm_build_395 string, dm_build_396 *DmConnection) string {
	return Dm_build_1331.Dm_build_1584(dm_build_393.Dm_build_376(dm_build_394), dm_build_395, dm_build_396)
}

func (dm_build_398 *Dm_build_78) Dm_build_397(dm_build_399 int, dm_build_400 string, dm_build_401 *DmConnection) string {
	return Dm_build_1331.Dm_build_1584(dm_build_398.Dm_build_380(dm_build_399), dm_build_400, dm_build_401)
}

func (dm_build_403 *Dm_build_78) Dm_build_402(dm_build_404 int, dm_build_405 string, dm_build_406 *DmConnection) string {
	return Dm_build_1331.Dm_build_1584(dm_build_403.Dm_build_384(dm_build_404), dm_build_405, dm_build_406)
}

func (dm_build_408 *Dm_build_78) Dm_build_407(dm_build_409 int, dm_build_410 string, dm_build_411 *DmConnection) string {
	return Dm_build_1331.Dm_build_1584(dm_build_408.Dm_build_388(dm_build_409), dm_build_410, dm_build_411)
}
