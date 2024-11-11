/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */

package dm8

import (
	"container/list"
	"io"
)

type Dm_build_1212 struct {
	dm_build_1213 *list.List
	dm_build_1214 *dm_build_1266
	dm_build_1215 int
}

func Dm_build_1216() *Dm_build_1212 {
	return &Dm_build_1212{
		dm_build_1213: list.New(),
		dm_build_1215: 0,
	}
}

func (dm_build_1218 *Dm_build_1212) Dm_build_1217() int {
	return dm_build_1218.dm_build_1215
}

func (dm_build_1220 *Dm_build_1212) Dm_build_1219(dm_build_1221 *Dm_build_1290, dm_build_1222 int) int {
	var dm_build_1223 = 0
	var dm_build_1224 = 0
	for dm_build_1223 < dm_build_1222 && dm_build_1220.dm_build_1214 != nil {
		dm_build_1224 = dm_build_1220.dm_build_1214.dm_build_1274(dm_build_1221, dm_build_1222-dm_build_1223)
		if dm_build_1220.dm_build_1214.dm_build_1269 == 0 {
			dm_build_1220.dm_build_1256()
		}
		dm_build_1223 += dm_build_1224
		dm_build_1220.dm_build_1215 -= dm_build_1224
	}
	return dm_build_1223
}

func (dm_build_1226 *Dm_build_1212) Dm_build_1225(dm_build_1227 []byte, dm_build_1228 int, dm_build_1229 int) int {
	var dm_build_1230 = 0
	var dm_build_1231 = 0
	for dm_build_1230 < dm_build_1229 && dm_build_1226.dm_build_1214 != nil {
		dm_build_1231 = dm_build_1226.dm_build_1214.dm_build_1278(dm_build_1227, dm_build_1228, dm_build_1229-dm_build_1230)
		if dm_build_1226.dm_build_1214.dm_build_1269 == 0 {
			dm_build_1226.dm_build_1256()
		}
		dm_build_1230 += dm_build_1231
		dm_build_1226.dm_build_1215 -= dm_build_1231
		dm_build_1228 += dm_build_1231
	}
	return dm_build_1230
}

func (dm_build_1233 *Dm_build_1212) Dm_build_1232(dm_build_1234 io.Writer, dm_build_1235 int) int {
	var dm_build_1236 = 0
	var dm_build_1237 = 0
	for dm_build_1236 < dm_build_1235 && dm_build_1233.dm_build_1214 != nil {
		dm_build_1237 = dm_build_1233.dm_build_1214.dm_build_1283(dm_build_1234, dm_build_1235-dm_build_1236)
		if dm_build_1233.dm_build_1214.dm_build_1269 == 0 {
			dm_build_1233.dm_build_1256()
		}
		dm_build_1236 += dm_build_1237
		dm_build_1233.dm_build_1215 -= dm_build_1237
	}
	return dm_build_1236
}

func (dm_build_1239 *Dm_build_1212) Dm_build_1238(dm_build_1240 []byte, dm_build_1241 int, dm_build_1242 int) {
	if dm_build_1242 == 0 {
		return
	}
	var dm_build_1243 = dm_build_1270(dm_build_1240, dm_build_1241, dm_build_1242)
	if dm_build_1239.dm_build_1214 == nil {
		dm_build_1239.dm_build_1214 = dm_build_1243
	} else {
		dm_build_1239.dm_build_1213.PushBack(dm_build_1243)
	}
	dm_build_1239.dm_build_1215 += dm_build_1242
}

func (dm_build_1245 *Dm_build_1212) dm_build_1244(dm_build_1246 int) byte {
	var dm_build_1247 = dm_build_1246
	var dm_build_1248 = dm_build_1245.dm_build_1214
	for dm_build_1247 > 0 && dm_build_1248 != nil {
		if dm_build_1248.dm_build_1269 == 0 {
			continue
		}
		if dm_build_1247 > dm_build_1248.dm_build_1269-1 {
			dm_build_1247 -= dm_build_1248.dm_build_1269
			dm_build_1248 = dm_build_1245.dm_build_1213.Front().Value.(*dm_build_1266)
		} else {
			break
		}
	}
	if dm_build_1248 != nil {
		return dm_build_1248.dm_build_1287(dm_build_1247)
	}
	return 0
}
func (dm_build_1250 *Dm_build_1212) Dm_build_1249(dm_build_1251 *Dm_build_1212) {
	if dm_build_1251.dm_build_1215 == 0 {
		return
	}
	var dm_build_1252 = dm_build_1251.dm_build_1214
	for dm_build_1252 != nil {
		dm_build_1250.dm_build_1253(dm_build_1252)
		dm_build_1251.dm_build_1256()
		dm_build_1252 = dm_build_1251.dm_build_1214
	}
	dm_build_1251.dm_build_1215 = 0
}
func (dm_build_1254 *Dm_build_1212) dm_build_1253(dm_build_1255 *dm_build_1266) {
	if dm_build_1255.dm_build_1269 == 0 {
		return
	}
	if dm_build_1254.dm_build_1214 == nil {
		dm_build_1254.dm_build_1214 = dm_build_1255
	} else {
		dm_build_1254.dm_build_1213.PushBack(dm_build_1255)
	}
	dm_build_1254.dm_build_1215 += dm_build_1255.dm_build_1269
}

func (dm_build_1257 *Dm_build_1212) dm_build_1256() {
	var dm_build_1258 = dm_build_1257.dm_build_1213.Front()
	if dm_build_1258 == nil {
		dm_build_1257.dm_build_1214 = nil
	} else {
		dm_build_1257.dm_build_1214 = dm_build_1258.Value.(*dm_build_1266)
		dm_build_1257.dm_build_1213.Remove(dm_build_1258)
	}
}

func (dm_build_1260 *Dm_build_1212) Dm_build_1259() []byte {
	var dm_build_1261 = make([]byte, dm_build_1260.dm_build_1215)
	var dm_build_1262 = dm_build_1260.dm_build_1214
	var dm_build_1263 = 0
	var dm_build_1264 = len(dm_build_1261)
	var dm_build_1265 = 0
	for dm_build_1262 != nil {
		if dm_build_1262.dm_build_1269 > 0 {
			if dm_build_1264 > dm_build_1262.dm_build_1269 {
				dm_build_1265 = dm_build_1262.dm_build_1269
			} else {
				dm_build_1265 = dm_build_1264
			}
			copy(dm_build_1261[dm_build_1263:dm_build_1263+dm_build_1265], dm_build_1262.dm_build_1267[dm_build_1262.dm_build_1268:dm_build_1262.dm_build_1268+dm_build_1265])
			dm_build_1263 += dm_build_1265
			dm_build_1264 -= dm_build_1265
		}
		if dm_build_1260.dm_build_1213.Front() == nil {
			dm_build_1262 = nil
		} else {
			dm_build_1262 = dm_build_1260.dm_build_1213.Front().Value.(*dm_build_1266)
		}
	}
	return dm_build_1261
}

type dm_build_1266 struct {
	dm_build_1267 []byte
	dm_build_1268 int
	dm_build_1269 int
}

func dm_build_1270(dm_build_1271 []byte, dm_build_1272 int, dm_build_1273 int) *dm_build_1266 {
	return &dm_build_1266{
		dm_build_1271,
		dm_build_1272,
		dm_build_1273,
	}
}

func (dm_build_1275 *dm_build_1266) dm_build_1274(dm_build_1276 *Dm_build_1290, dm_build_1277 int) int {
	if dm_build_1275.dm_build_1269 <= dm_build_1277 {
		dm_build_1277 = dm_build_1275.dm_build_1269
	}
	dm_build_1276.Dm_build_1373(dm_build_1275.dm_build_1267[dm_build_1275.dm_build_1268 : dm_build_1275.dm_build_1268+dm_build_1277])
	dm_build_1275.dm_build_1268 += dm_build_1277
	dm_build_1275.dm_build_1269 -= dm_build_1277
	return dm_build_1277
}

func (dm_build_1279 *dm_build_1266) dm_build_1278(dm_build_1280 []byte, dm_build_1281 int, dm_build_1282 int) int {
	if dm_build_1279.dm_build_1269 <= dm_build_1282 {
		dm_build_1282 = dm_build_1279.dm_build_1269
	}
	copy(dm_build_1280[dm_build_1281:dm_build_1281+dm_build_1282], dm_build_1279.dm_build_1267[dm_build_1279.dm_build_1268:dm_build_1279.dm_build_1268+dm_build_1282])
	dm_build_1279.dm_build_1268 += dm_build_1282
	dm_build_1279.dm_build_1269 -= dm_build_1282
	return dm_build_1282
}

func (dm_build_1284 *dm_build_1266) dm_build_1283(dm_build_1285 io.Writer, dm_build_1286 int) int {
	if dm_build_1284.dm_build_1269 <= dm_build_1286 {
		dm_build_1286 = dm_build_1284.dm_build_1269
	}
	_, _ = dm_build_1285.Write(dm_build_1284.dm_build_1267[dm_build_1284.dm_build_1268 : dm_build_1284.dm_build_1268+dm_build_1286])
	dm_build_1284.dm_build_1268 += dm_build_1286
	dm_build_1284.dm_build_1269 -= dm_build_1286
	return dm_build_1286
}
func (dm_build_1288 *dm_build_1266) dm_build_1287(dm_build_1289 int) byte {
	return dm_build_1288.dm_build_1267[dm_build_1288.dm_build_1268+dm_build_1289]
}
