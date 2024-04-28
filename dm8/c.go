/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */

package dm8

import (
	"io"
	"math"
)

type Dm_build_360 struct {
	dm_build_361 []byte
	dm_build_362 int
}

func Dm_build_363(dm_build_364 int) *Dm_build_360 {
	return &Dm_build_360{make([]byte, 0, dm_build_364), 0}
}

func Dm_build_365(dm_build_366 []byte) *Dm_build_360 {
	return &Dm_build_360{dm_build_366, 0}
}

func (dm_build_368 *Dm_build_360) dm_build_367(dm_build_369 int) *Dm_build_360 {

	dm_build_370 := len(dm_build_368.dm_build_361)
	dm_build_371 := cap(dm_build_368.dm_build_361)

	if dm_build_370+dm_build_369 <= dm_build_371 {
		dm_build_368.dm_build_361 = dm_build_368.dm_build_361[:dm_build_370+dm_build_369]
	} else {

		var calCap = int64(math.Max(float64(2*dm_build_371), float64(dm_build_369+dm_build_370)))

		nbuf := make([]byte, dm_build_369+dm_build_370, calCap)
		copy(nbuf, dm_build_368.dm_build_361)
		dm_build_368.dm_build_361 = nbuf
	}

	return dm_build_368
}

func (dm_build_373 *Dm_build_360) Dm_build_372() int {
	return len(dm_build_373.dm_build_361)
}

func (dm_build_375 *Dm_build_360) Dm_build_374(dm_build_376 int) *Dm_build_360 {
	for i := dm_build_376; i < len(dm_build_375.dm_build_361); i++ {
		dm_build_375.dm_build_361[i] = 0
	}
	dm_build_375.dm_build_361 = dm_build_375.dm_build_361[:dm_build_376]
	return dm_build_375
}

func (dm_build_378 *Dm_build_360) Dm_build_377(dm_build_379 int) *Dm_build_360 {
	dm_build_378.dm_build_362 = dm_build_379
	return dm_build_378
}

func (dm_build_381 *Dm_build_360) Dm_build_380() int {
	return dm_build_381.dm_build_362
}

func (dm_build_383 *Dm_build_360) Dm_build_382(_ bool) int {
	return len(dm_build_383.dm_build_361) - dm_build_383.dm_build_362
}

func (dm_build_386 *Dm_build_360) Dm_build_385(dm_build_387 int, dm_build_388 bool, dm_build_389 bool) *Dm_build_360 {

	if dm_build_388 {
		if dm_build_389 {
			dm_build_386.dm_build_367(dm_build_387)
		} else {
			dm_build_386.dm_build_361 = dm_build_386.dm_build_361[:len(dm_build_386.dm_build_361)-dm_build_387]
		}
	} else {
		if dm_build_389 {
			dm_build_386.dm_build_362 += dm_build_387
		} else {
			dm_build_386.dm_build_362 -= dm_build_387
		}
	}

	return dm_build_386
}

func (dm_build_391 *Dm_build_360) Dm_build_390(dm_build_392 io.Reader, dm_build_393 int) (int, error) {
	dm_build_394 := len(dm_build_391.dm_build_361)
	dm_build_391.dm_build_367(dm_build_393)
	dm_build_395 := 0
	for dm_build_393 > 0 {
		n, err := dm_build_392.Read(dm_build_391.dm_build_361[dm_build_394+dm_build_395:])
		if n > 0 && err == io.EOF {
			dm_build_395 += n
			dm_build_391.dm_build_361 = dm_build_391.dm_build_361[:dm_build_394+dm_build_395]
			return dm_build_395, nil
		} else if n > 0 && err == nil {
			dm_build_393 -= n
			dm_build_395 += n
		} else if n == 0 && err != nil {
			return -1, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
		}
	}

	return dm_build_395, nil
}

func (dm_build_397 *Dm_build_360) Dm_build_396(dm_build_398 io.Writer) (*Dm_build_360, error) {
	if _, err := dm_build_398.Write(dm_build_397.dm_build_361); err != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
	}
	return dm_build_397, nil
}

func (dm_build_400 *Dm_build_360) Dm_build_399(dm_build_401 bool) int {
	dm_build_402 := len(dm_build_400.dm_build_361)
	dm_build_400.dm_build_367(1)

	if dm_build_401 {
		return copy(dm_build_400.dm_build_361[dm_build_402:], []byte{1})
	} else {
		return copy(dm_build_400.dm_build_361[dm_build_402:], []byte{0})
	}
}

func (dm_build_404 *Dm_build_360) Dm_build_403(dm_build_405 byte) int {
	dm_build_406 := len(dm_build_404.dm_build_361)
	dm_build_404.dm_build_367(1)

	return copy(dm_build_404.dm_build_361[dm_build_406:], Dm_build_1.Dm_build_179(dm_build_405))
}

func (dm_build_408 *Dm_build_360) Dm_build_407(dm_build_409 int8) int {
	dm_build_410 := len(dm_build_408.dm_build_361)
	dm_build_408.dm_build_367(1)

	return copy(dm_build_408.dm_build_361[dm_build_410:], Dm_build_1.Dm_build_182(dm_build_409))
}

func (dm_build_412 *Dm_build_360) Dm_build_411(dm_build_413 int16) int {
	dm_build_414 := len(dm_build_412.dm_build_361)
	dm_build_412.dm_build_367(2)

	return copy(dm_build_412.dm_build_361[dm_build_414:], Dm_build_1.Dm_build_185(dm_build_413))
}

func (dm_build_416 *Dm_build_360) Dm_build_415(dm_build_417 int32) int {
	dm_build_418 := len(dm_build_416.dm_build_361)
	dm_build_416.dm_build_367(4)

	return copy(dm_build_416.dm_build_361[dm_build_418:], Dm_build_1.Dm_build_188(dm_build_417))
}

func (dm_build_420 *Dm_build_360) Dm_build_419(dm_build_421 uint8) int {
	dm_build_422 := len(dm_build_420.dm_build_361)
	dm_build_420.dm_build_367(1)

	return copy(dm_build_420.dm_build_361[dm_build_422:], Dm_build_1.Dm_build_200(dm_build_421))
}

func (dm_build_424 *Dm_build_360) Dm_build_423(dm_build_425 uint16) int {
	dm_build_426 := len(dm_build_424.dm_build_361)
	dm_build_424.dm_build_367(2)

	return copy(dm_build_424.dm_build_361[dm_build_426:], Dm_build_1.Dm_build_203(dm_build_425))
}

func (dm_build_428 *Dm_build_360) Dm_build_427(dm_build_429 uint32) int {
	dm_build_430 := len(dm_build_428.dm_build_361)
	dm_build_428.dm_build_367(4)

	return copy(dm_build_428.dm_build_361[dm_build_430:], Dm_build_1.Dm_build_206(dm_build_429))
}

func (dm_build_432 *Dm_build_360) Dm_build_431(dm_build_433 uint64) int {
	dm_build_434 := len(dm_build_432.dm_build_361)
	dm_build_432.dm_build_367(8)

	return copy(dm_build_432.dm_build_361[dm_build_434:], Dm_build_1.Dm_build_209(dm_build_433))
}

func (dm_build_436 *Dm_build_360) Dm_build_435(dm_build_437 float32) int {
	dm_build_438 := len(dm_build_436.dm_build_361)
	dm_build_436.dm_build_367(4)

	return copy(dm_build_436.dm_build_361[dm_build_438:], Dm_build_1.Dm_build_206(math.Float32bits(dm_build_437)))
}

func (dm_build_440 *Dm_build_360) Dm_build_439(dm_build_441 float64) int {
	dm_build_442 := len(dm_build_440.dm_build_361)
	dm_build_440.dm_build_367(8)

	return copy(dm_build_440.dm_build_361[dm_build_442:], Dm_build_1.Dm_build_209(math.Float64bits(dm_build_441)))
}

func (dm_build_444 *Dm_build_360) Dm_build_443(dm_build_445 []byte) int {
	dm_build_446 := len(dm_build_444.dm_build_361)
	dm_build_444.dm_build_367(len(dm_build_445))
	return copy(dm_build_444.dm_build_361[dm_build_446:], dm_build_445)
}

func (dm_build_448 *Dm_build_360) Dm_build_447(dm_build_449 []byte) int {
	return dm_build_448.Dm_build_415(int32(len(dm_build_449))) + dm_build_448.Dm_build_443(dm_build_449)
}

func (dm_build_451 *Dm_build_360) Dm_build_450(dm_build_452 []byte) int {
	return dm_build_451.Dm_build_419(uint8(len(dm_build_452))) + dm_build_451.Dm_build_443(dm_build_452)
}

func (dm_build_454 *Dm_build_360) Dm_build_453(dm_build_455 []byte) int {
	return dm_build_454.Dm_build_423(uint16(len(dm_build_455))) + dm_build_454.Dm_build_443(dm_build_455)
}

func (dm_build_457 *Dm_build_360) Dm_build_456(dm_build_458 []byte) int {
	return dm_build_457.Dm_build_443(dm_build_458) + dm_build_457.Dm_build_403(0)
}

func (dm_build_460 *Dm_build_360) Dm_build_459(dm_build_461 string, dm_build_462 string, dm_build_463 *DmConnection) int {
	dm_build_464 := Dm_build_1.Dm_build_217(dm_build_461, dm_build_462, dm_build_463)
	return dm_build_460.Dm_build_447(dm_build_464)
}

func (dm_build_466 *Dm_build_360) Dm_build_465(dm_build_467 string, dm_build_468 string, dm_build_469 *DmConnection) int {
	dm_build_470 := Dm_build_1.Dm_build_217(dm_build_467, dm_build_468, dm_build_469)
	return dm_build_466.Dm_build_450(dm_build_470)
}

func (dm_build_472 *Dm_build_360) Dm_build_471(dm_build_473 string, dm_build_474 string, dm_build_475 *DmConnection) int {
	dm_build_476 := Dm_build_1.Dm_build_217(dm_build_473, dm_build_474, dm_build_475)
	return dm_build_472.Dm_build_453(dm_build_476)
}

func (dm_build_478 *Dm_build_360) Dm_build_477(dm_build_479 string, dm_build_480 string, dm_build_481 *DmConnection) int {
	dm_build_482 := Dm_build_1.Dm_build_217(dm_build_479, dm_build_480, dm_build_481)
	return dm_build_478.Dm_build_456(dm_build_482)
}

func (dm_build_484 *Dm_build_360) Dm_build_483() byte {
	dm_build_485 := Dm_build_1.Dm_build_94(dm_build_484.dm_build_361, dm_build_484.dm_build_362)
	dm_build_484.dm_build_362++
	return dm_build_485
}

func (dm_build_487 *Dm_build_360) Dm_build_486() int16 {
	dm_build_488 := Dm_build_1.Dm_build_98(dm_build_487.dm_build_361, dm_build_487.dm_build_362)
	dm_build_487.dm_build_362 += 2
	return dm_build_488
}

func (dm_build_490 *Dm_build_360) Dm_build_489() int32 {
	dm_build_491 := Dm_build_1.Dm_build_103(dm_build_490.dm_build_361, dm_build_490.dm_build_362)
	dm_build_490.dm_build_362 += 4
	return dm_build_491
}

func (dm_build_493 *Dm_build_360) Dm_build_492() int64 {
	dm_build_494 := Dm_build_1.Dm_build_108(dm_build_493.dm_build_361, dm_build_493.dm_build_362)
	dm_build_493.dm_build_362 += 8
	return dm_build_494
}

func (dm_build_496 *Dm_build_360) Dm_build_495() float32 {
	dm_build_497 := Dm_build_1.Dm_build_113(dm_build_496.dm_build_361, dm_build_496.dm_build_362)
	dm_build_496.dm_build_362 += 4
	return dm_build_497
}

func (dm_build_499 *Dm_build_360) Dm_build_498() float64 {
	dm_build_500 := Dm_build_1.Dm_build_117(dm_build_499.dm_build_361, dm_build_499.dm_build_362)
	dm_build_499.dm_build_362 += 8
	return dm_build_500
}

func (dm_build_502 *Dm_build_360) Dm_build_501() uint8 {
	dm_build_503 := Dm_build_1.Dm_build_121(dm_build_502.dm_build_361, dm_build_502.dm_build_362)
	dm_build_502.dm_build_362 += 1
	return dm_build_503
}

func (dm_build_505 *Dm_build_360) Dm_build_504() uint16 {
	dm_build_506 := Dm_build_1.Dm_build_125(dm_build_505.dm_build_361, dm_build_505.dm_build_362)
	dm_build_505.dm_build_362 += 2
	return dm_build_506
}

func (dm_build_508 *Dm_build_360) Dm_build_507() uint32 {
	dm_build_509 := Dm_build_1.Dm_build_130(dm_build_508.dm_build_361, dm_build_508.dm_build_362)
	dm_build_508.dm_build_362 += 4
	return dm_build_509
}

func (dm_build_511 *Dm_build_360) Dm_build_510(dm_build_512 int) []byte {
	dm_build_513 := Dm_build_1.Dm_build_152(dm_build_511.dm_build_361, dm_build_511.dm_build_362, dm_build_512)
	dm_build_511.dm_build_362 += dm_build_512
	return dm_build_513
}

func (dm_build_515 *Dm_build_360) Dm_build_514() []byte {
	return dm_build_515.Dm_build_510(int(dm_build_515.Dm_build_489()))
}

func (dm_build_517 *Dm_build_360) Dm_build_516() []byte {
	return dm_build_517.Dm_build_510(int(dm_build_517.Dm_build_483()))
}

func (dm_build_519 *Dm_build_360) Dm_build_518() []byte {
	return dm_build_519.Dm_build_510(int(dm_build_519.Dm_build_486()))
}

func (dm_build_521 *Dm_build_360) Dm_build_520(dm_build_522 int) []byte {
	return dm_build_521.Dm_build_510(dm_build_522)
}

func (dm_build_524 *Dm_build_360) Dm_build_523() []byte {
	dm_build_525 := 0
	for dm_build_524.Dm_build_483() != 0 {
		dm_build_525++
	}
	dm_build_524.Dm_build_385(dm_build_525, false, false)
	return dm_build_524.Dm_build_510(dm_build_525)
}

func (dm_build_527 *Dm_build_360) Dm_build_526(dm_build_528 int, dm_build_529 string, dm_build_530 *DmConnection) string {
	return Dm_build_1.Dm_build_253(dm_build_527.Dm_build_510(dm_build_528), dm_build_529, dm_build_530)
}

func (dm_build_532 *Dm_build_360) Dm_build_531(dm_build_533 string, dm_build_534 *DmConnection) string {
	return Dm_build_1.Dm_build_253(dm_build_532.Dm_build_514(), dm_build_533, dm_build_534)
}

func (dm_build_536 *Dm_build_360) Dm_build_535(dm_build_537 string, dm_build_538 *DmConnection) string {
	return Dm_build_1.Dm_build_253(dm_build_536.Dm_build_516(), dm_build_537, dm_build_538)
}

func (dm_build_540 *Dm_build_360) Dm_build_539(dm_build_541 string, dm_build_542 *DmConnection) string {
	return Dm_build_1.Dm_build_253(dm_build_540.Dm_build_518(), dm_build_541, dm_build_542)
}

func (dm_build_544 *Dm_build_360) Dm_build_543(dm_build_545 string, dm_build_546 *DmConnection) string {
	return Dm_build_1.Dm_build_253(dm_build_544.Dm_build_523(), dm_build_545, dm_build_546)
}

func (dm_build_548 *Dm_build_360) Dm_build_547(dm_build_549 int, dm_build_550 byte) int {
	return dm_build_548.Dm_build_583(dm_build_549, Dm_build_1.Dm_build_179(dm_build_550))
}

func (dm_build_552 *Dm_build_360) Dm_build_551(dm_build_553 int, dm_build_554 int16) int {
	return dm_build_552.Dm_build_583(dm_build_553, Dm_build_1.Dm_build_185(dm_build_554))
}

func (dm_build_556 *Dm_build_360) Dm_build_555(dm_build_557 int, dm_build_558 int32) int {
	return dm_build_556.Dm_build_583(dm_build_557, Dm_build_1.Dm_build_188(dm_build_558))
}

func (dm_build_560 *Dm_build_360) Dm_build_559(dm_build_561 int, dm_build_562 int64) int {
	return dm_build_560.Dm_build_583(dm_build_561, Dm_build_1.Dm_build_191(dm_build_562))
}

func (dm_build_564 *Dm_build_360) Dm_build_563(dm_build_565 int, dm_build_566 float32) int {
	return dm_build_564.Dm_build_583(dm_build_565, Dm_build_1.Dm_build_194(dm_build_566))
}

func (dm_build_568 *Dm_build_360) Dm_build_567(dm_build_569 int, dm_build_570 float64) int {
	return dm_build_568.Dm_build_583(dm_build_569, Dm_build_1.Dm_build_197(dm_build_570))
}

func (dm_build_572 *Dm_build_360) Dm_build_571(dm_build_573 int, dm_build_574 uint8) int {
	return dm_build_572.Dm_build_583(dm_build_573, Dm_build_1.Dm_build_200(dm_build_574))
}

func (dm_build_576 *Dm_build_360) Dm_build_575(dm_build_577 int, dm_build_578 uint16) int {
	return dm_build_576.Dm_build_583(dm_build_577, Dm_build_1.Dm_build_203(dm_build_578))
}

func (dm_build_580 *Dm_build_360) Dm_build_579(dm_build_581 int, dm_build_582 uint32) int {
	return dm_build_580.Dm_build_583(dm_build_581, Dm_build_1.Dm_build_206(dm_build_582))
}

func (dm_build_584 *Dm_build_360) Dm_build_583(dm_build_585 int, dm_build_586 []byte) int {
	return copy(dm_build_584.dm_build_361[dm_build_585:], dm_build_586)
}

func (dm_build_588 *Dm_build_360) Dm_build_587(dm_build_589 int, dm_build_590 []byte) int {
	return dm_build_588.Dm_build_555(dm_build_589, int32(len(dm_build_590))) + dm_build_588.Dm_build_583(dm_build_589+4, dm_build_590)
}

func (dm_build_592 *Dm_build_360) Dm_build_591(dm_build_593 int, dm_build_594 []byte) int {
	return dm_build_592.Dm_build_547(dm_build_593, byte(len(dm_build_594))) + dm_build_592.Dm_build_583(dm_build_593+1, dm_build_594)
}

func (dm_build_596 *Dm_build_360) Dm_build_595(dm_build_597 int, dm_build_598 []byte) int {
	return dm_build_596.Dm_build_551(dm_build_597, int16(len(dm_build_598))) + dm_build_596.Dm_build_583(dm_build_597+2, dm_build_598)
}

func (dm_build_600 *Dm_build_360) Dm_build_599(dm_build_601 int, dm_build_602 []byte) int {
	return dm_build_600.Dm_build_583(dm_build_601, dm_build_602) + dm_build_600.Dm_build_547(dm_build_601+len(dm_build_602), 0)
}

func (dm_build_604 *Dm_build_360) Dm_build_603(dm_build_605 int, dm_build_606 string, dm_build_607 string, dm_build_608 *DmConnection) int {
	return dm_build_604.Dm_build_587(dm_build_605, Dm_build_1.Dm_build_217(dm_build_606, dm_build_607, dm_build_608))
}

func (dm_build_610 *Dm_build_360) Dm_build_609(dm_build_611 int, dm_build_612 string, dm_build_613 string, dm_build_614 *DmConnection) int {
	return dm_build_610.Dm_build_591(dm_build_611, Dm_build_1.Dm_build_217(dm_build_612, dm_build_613, dm_build_614))
}

func (dm_build_616 *Dm_build_360) Dm_build_615(dm_build_617 int, dm_build_618 string, dm_build_619 string, dm_build_620 *DmConnection) int {
	return dm_build_616.Dm_build_595(dm_build_617, Dm_build_1.Dm_build_217(dm_build_618, dm_build_619, dm_build_620))
}

func (dm_build_622 *Dm_build_360) Dm_build_621(dm_build_623 int, dm_build_624 string, dm_build_625 string, dm_build_626 *DmConnection) int {
	return dm_build_622.Dm_build_599(dm_build_623, Dm_build_1.Dm_build_217(dm_build_624, dm_build_625, dm_build_626))
}

func (dm_build_628 *Dm_build_360) Dm_build_627(dm_build_629 int) byte {
	return Dm_build_1.Dm_build_222(dm_build_628.Dm_build_654(dm_build_629, 1))
}

func (dm_build_631 *Dm_build_360) Dm_build_630(dm_build_632 int) int16 {
	return Dm_build_1.Dm_build_225(dm_build_631.Dm_build_654(dm_build_632, 2))
}

func (dm_build_634 *Dm_build_360) Dm_build_633(dm_build_635 int) int32 {
	return Dm_build_1.Dm_build_228(dm_build_634.Dm_build_654(dm_build_635, 4))
}

func (dm_build_637 *Dm_build_360) Dm_build_636(dm_build_638 int) int64 {
	return Dm_build_1.Dm_build_231(dm_build_637.Dm_build_654(dm_build_638, 8))
}

func (dm_build_640 *Dm_build_360) Dm_build_639(dm_build_641 int) float32 {
	return Dm_build_1.Dm_build_234(dm_build_640.Dm_build_654(dm_build_641, 4))
}

func (dm_build_643 *Dm_build_360) Dm_build_642(dm_build_644 int) float64 {
	return Dm_build_1.Dm_build_237(dm_build_643.Dm_build_654(dm_build_644, 8))
}

func (dm_build_646 *Dm_build_360) Dm_build_645(dm_build_647 int) uint8 {
	return Dm_build_1.Dm_build_240(dm_build_646.Dm_build_654(dm_build_647, 1))
}

func (dm_build_649 *Dm_build_360) Dm_build_648(dm_build_650 int) uint16 {
	return Dm_build_1.Dm_build_243(dm_build_649.Dm_build_654(dm_build_650, 2))
}

func (dm_build_652 *Dm_build_360) Dm_build_651(dm_build_653 int) uint32 {
	return Dm_build_1.Dm_build_246(dm_build_652.Dm_build_654(dm_build_653, 4))
}

func (dm_build_655 *Dm_build_360) Dm_build_654(dm_build_656 int, dm_build_657 int) []byte {
	return dm_build_655.dm_build_361[dm_build_656 : dm_build_656+dm_build_657]
}

func (dm_build_659 *Dm_build_360) Dm_build_658(dm_build_660 int) []byte {
	dm_build_661 := dm_build_659.Dm_build_633(dm_build_660)
	return dm_build_659.Dm_build_654(dm_build_660+4, int(dm_build_661))
}

func (dm_build_663 *Dm_build_360) Dm_build_662(dm_build_664 int) []byte {
	dm_build_665 := dm_build_663.Dm_build_627(dm_build_664)
	return dm_build_663.Dm_build_654(dm_build_664+1, int(dm_build_665))
}

func (dm_build_667 *Dm_build_360) Dm_build_666(dm_build_668 int) []byte {
	dm_build_669 := dm_build_667.Dm_build_630(dm_build_668)
	return dm_build_667.Dm_build_654(dm_build_668+2, int(dm_build_669))
}

func (dm_build_671 *Dm_build_360) Dm_build_670(dm_build_672 int) []byte {
	dm_build_673 := 0
	for dm_build_671.Dm_build_627(dm_build_672) != 0 {
		dm_build_672++
		dm_build_673++
	}

	return dm_build_671.Dm_build_654(dm_build_672-dm_build_673, dm_build_673)
}

func (dm_build_675 *Dm_build_360) Dm_build_674(dm_build_676 int, dm_build_677 string, dm_build_678 *DmConnection) string {
	return Dm_build_1.Dm_build_253(dm_build_675.Dm_build_658(dm_build_676), dm_build_677, dm_build_678)
}

func (dm_build_680 *Dm_build_360) Dm_build_679(dm_build_681 int, dm_build_682 string, dm_build_683 *DmConnection) string {
	return Dm_build_1.Dm_build_253(dm_build_680.Dm_build_662(dm_build_681), dm_build_682, dm_build_683)
}

func (dm_build_685 *Dm_build_360) Dm_build_684(dm_build_686 int, dm_build_687 string, dm_build_688 *DmConnection) string {
	return Dm_build_1.Dm_build_253(dm_build_685.Dm_build_666(dm_build_686), dm_build_687, dm_build_688)
}

func (dm_build_690 *Dm_build_360) Dm_build_689(dm_build_691 int, dm_build_692 string, dm_build_693 *DmConnection) string {
	return Dm_build_1.Dm_build_253(dm_build_690.Dm_build_670(dm_build_691), dm_build_692, dm_build_693)
}
