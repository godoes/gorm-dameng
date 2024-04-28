/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */

package dm8

import (
	"container/list"
	"io"
)

type Dm_build_282 struct {
	dm_build_283 *list.List
	dm_build_284 *dm_build_336
	dm_build_285 int
}

func Dm_build_286() *Dm_build_282 {
	return &Dm_build_282{
		dm_build_283: list.New(),
		dm_build_285: 0,
	}
}

func (dm_build_288 *Dm_build_282) Dm_build_287() int {
	return dm_build_288.dm_build_285
}

func (dm_build_290 *Dm_build_282) Dm_build_289(dm_build_291 *Dm_build_360, dm_build_292 int) int {
	var dm_build_293 = 0
	var dm_build_294 = 0
	for dm_build_293 < dm_build_292 && dm_build_290.dm_build_284 != nil {
		dm_build_294 = dm_build_290.dm_build_284.dm_build_344(dm_build_291, dm_build_292-dm_build_293)
		if dm_build_290.dm_build_284.dm_build_339 == 0 {
			dm_build_290.dm_build_326()
		}
		dm_build_293 += dm_build_294
		dm_build_290.dm_build_285 -= dm_build_294
	}
	return dm_build_293
}

func (dm_build_296 *Dm_build_282) Dm_build_295(dm_build_297 []byte, dm_build_298 int, dm_build_299 int) int {
	var dm_build_300 = 0
	var dm_build_301 = 0
	for dm_build_300 < dm_build_299 && dm_build_296.dm_build_284 != nil {
		dm_build_301 = dm_build_296.dm_build_284.dm_build_348(dm_build_297, dm_build_298, dm_build_299-dm_build_300)
		if dm_build_296.dm_build_284.dm_build_339 == 0 {
			dm_build_296.dm_build_326()
		}
		dm_build_300 += dm_build_301
		dm_build_296.dm_build_285 -= dm_build_301
		dm_build_298 += dm_build_301
	}
	return dm_build_300
}

func (dm_build_303 *Dm_build_282) Dm_build_302(dm_build_304 io.Writer, dm_build_305 int) int {
	var dm_build_306 = 0
	var dm_build_307 = 0
	for dm_build_306 < dm_build_305 && dm_build_303.dm_build_284 != nil {
		dm_build_307 = dm_build_303.dm_build_284.dm_build_353(dm_build_304, dm_build_305-dm_build_306)
		if dm_build_303.dm_build_284.dm_build_339 == 0 {
			dm_build_303.dm_build_326()
		}
		dm_build_306 += dm_build_307
		dm_build_303.dm_build_285 -= dm_build_307
	}
	return dm_build_306
}

func (dm_build_309 *Dm_build_282) Dm_build_308(dm_build_310 []byte, dm_build_311 int, dm_build_312 int) {
	if dm_build_312 == 0 {
		return
	}
	var dm_build_313 = dm_build_340(dm_build_310, dm_build_311, dm_build_312)
	if dm_build_309.dm_build_284 == nil {
		dm_build_309.dm_build_284 = dm_build_313
	} else {
		dm_build_309.dm_build_283.PushBack(dm_build_313)
	}
	dm_build_309.dm_build_285 += dm_build_312
}

func (dm_build_315 *Dm_build_282) dm_build_314(dm_build_316 int) byte {
	var dm_build_317 = dm_build_316
	var dm_build_318 = dm_build_315.dm_build_284
	for dm_build_317 > 0 && dm_build_318 != nil {
		if dm_build_318.dm_build_339 == 0 {
			continue
		}
		if dm_build_317 > dm_build_318.dm_build_339-1 {
			dm_build_317 -= dm_build_318.dm_build_339
			dm_build_318 = dm_build_315.dm_build_283.Front().Value.(*dm_build_336)
		} else {
			break
		}
	}
	if dm_build_318 != nil {
		return dm_build_318.dm_build_357(dm_build_317)
	}
	return 0
}

func (dm_build_320 *Dm_build_282) Dm_build_319(dm_build_321 *Dm_build_282) {
	if dm_build_321.dm_build_285 == 0 {
		return
	}
	var dm_build_322 = dm_build_321.dm_build_284
	for dm_build_322 != nil {
		dm_build_320.dm_build_323(dm_build_322)
		dm_build_321.dm_build_326()
		dm_build_322 = dm_build_321.dm_build_284
	}
	dm_build_321.dm_build_285 = 0
}

func (dm_build_324 *Dm_build_282) dm_build_323(dm_build_325 *dm_build_336) {
	if dm_build_325.dm_build_339 == 0 {
		return
	}
	if dm_build_324.dm_build_284 == nil {
		dm_build_324.dm_build_284 = dm_build_325
	} else {
		dm_build_324.dm_build_283.PushBack(dm_build_325)
	}
	dm_build_324.dm_build_285 += dm_build_325.dm_build_339
}

func (dm_build_327 *Dm_build_282) dm_build_326() {
	var dm_build_328 = dm_build_327.dm_build_283.Front()
	if dm_build_328 == nil {
		dm_build_327.dm_build_284 = nil
	} else {
		dm_build_327.dm_build_284 = dm_build_328.Value.(*dm_build_336)
		dm_build_327.dm_build_283.Remove(dm_build_328)
	}
}

func (dm_build_330 *Dm_build_282) Dm_build_329() []byte {
	var dm_build_331 = make([]byte, dm_build_330.dm_build_285)
	var dm_build_332 = dm_build_330.dm_build_284
	var dm_build_333 = 0
	var dm_build_334 = len(dm_build_331)
	var dm_build_335 = 0
	for dm_build_332 != nil {
		if dm_build_332.dm_build_339 > 0 {
			if dm_build_334 > dm_build_332.dm_build_339 {
				dm_build_335 = dm_build_332.dm_build_339
			} else {
				dm_build_335 = dm_build_334
			}
			copy(dm_build_331[dm_build_333:dm_build_333+dm_build_335], dm_build_332.dm_build_337[dm_build_332.dm_build_338:dm_build_332.dm_build_338+dm_build_335])
			dm_build_333 += dm_build_335
			dm_build_334 -= dm_build_335
		}
		if dm_build_330.dm_build_283.Front() == nil {
			dm_build_332 = nil
		} else {
			dm_build_332 = dm_build_330.dm_build_283.Front().Value.(*dm_build_336)
		}
	}
	return dm_build_331
}

type dm_build_336 struct {
	dm_build_337 []byte
	dm_build_338 int
	dm_build_339 int
}

func dm_build_340(dm_build_341 []byte, dm_build_342 int, dm_build_343 int) *dm_build_336 {
	return &dm_build_336{
		dm_build_341,
		dm_build_342,
		dm_build_343,
	}
}

func (dm_build_345 *dm_build_336) dm_build_344(dm_build_346 *Dm_build_360, dm_build_347 int) int {
	if dm_build_345.dm_build_339 <= dm_build_347 {
		dm_build_347 = dm_build_345.dm_build_339
	}
	dm_build_346.Dm_build_443(dm_build_345.dm_build_337[dm_build_345.dm_build_338 : dm_build_345.dm_build_338+dm_build_347])
	dm_build_345.dm_build_338 += dm_build_347
	dm_build_345.dm_build_339 -= dm_build_347
	return dm_build_347
}

func (dm_build_349 *dm_build_336) dm_build_348(dm_build_350 []byte, dm_build_351 int, dm_build_352 int) int {
	if dm_build_349.dm_build_339 <= dm_build_352 {
		dm_build_352 = dm_build_349.dm_build_339
	}
	copy(dm_build_350[dm_build_351:dm_build_351+dm_build_352], dm_build_349.dm_build_337[dm_build_349.dm_build_338:dm_build_349.dm_build_338+dm_build_352])
	dm_build_349.dm_build_338 += dm_build_352
	dm_build_349.dm_build_339 -= dm_build_352
	return dm_build_352
}

func (dm_build_354 *dm_build_336) dm_build_353(dm_build_355 io.Writer, dm_build_356 int) int {
	if dm_build_354.dm_build_339 <= dm_build_356 {
		dm_build_356 = dm_build_354.dm_build_339
	}
	_, _ = dm_build_355.Write(dm_build_354.dm_build_337[dm_build_354.dm_build_338 : dm_build_354.dm_build_338+dm_build_356])
	dm_build_354.dm_build_338 += dm_build_356
	dm_build_354.dm_build_339 -= dm_build_356
	return dm_build_356
}

func (dm_build_358 *dm_build_336) dm_build_357(dm_build_359 int) byte {
	return dm_build_358.dm_build_337[dm_build_358.dm_build_338+dm_build_359]
}
