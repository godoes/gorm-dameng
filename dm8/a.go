/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */

package dm8

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/godoes/gorm-dameng/dm8/security"
)

const (
	Dm_build_412 = 8192
	Dm_build_413 = 2 * time.Second
)

type dm_build_414 struct {
	dm_build_415 *net.TCPConn
	dm_build_416 *tls.Conn
	dm_build_417 *Dm_build_78
	dm_build_418 *DmConnection
	dm_build_419 security.Cipher
	dm_build_420 bool
	dm_build_421 bool
	dm_build_422 *security.DhKey

	dm_build_423 bool
	dm_build_424 string
	dm_build_425 bool
}

func dm_build_426(dm_build_427 *DmConnection) (*dm_build_414, error) {
	dm_build_428, dm_build_429 := dm_build_431(dm_build_427.dmConnector.host+":"+strconv.Itoa(int(dm_build_427.dmConnector.port)), time.Duration(dm_build_427.dmConnector.socketTimeout)*time.Second)
	if dm_build_429 != nil {
		return nil, dm_build_429
	}

	dm_build_430 := dm_build_414{}
	dm_build_430.dm_build_415 = dm_build_428
	dm_build_430.dm_build_417 = Dm_build_81(Dm_build_696)
	dm_build_430.dm_build_418 = dm_build_427
	dm_build_430.dm_build_420 = false
	dm_build_430.dm_build_421 = false
	dm_build_430.dm_build_423 = false
	dm_build_430.dm_build_424 = ""
	dm_build_430.dm_build_425 = false
	dm_build_427.Access = &dm_build_430

	return &dm_build_430, nil
}

func dm_build_431(dm_build_432 string, dm_build_433 time.Duration) (*net.TCPConn, error) {
	dm_build_434, dm_build_435 := net.DialTimeout("tcp", dm_build_432, dm_build_433)
	if dm_build_435 != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_432).throw()
	}

	if tcpConn, ok := dm_build_434.(*net.TCPConn); ok {

		_ = tcpConn.SetKeepAlive(true)
		_ = tcpConn.SetKeepAlivePeriod(Dm_build_413)
		_ = tcpConn.SetNoDelay(true)

		return tcpConn, nil
	}

	return nil, nil
}

func (dm_build_437 *dm_build_414) dm_build_436(dm_build_438 dm_build_817) bool {
	var dm_build_439 = dm_build_437.dm_build_418.dmConnector.compress
	if dm_build_438.dm_build_832() == Dm_build_724 || dm_build_439 == Dm_build_773 {
		return false
	}

	if dm_build_439 == Dm_build_771 {
		return true
	} else if dm_build_439 == Dm_build_772 {
		return !dm_build_437.dm_build_418.Local && dm_build_438.dm_build_830() > Dm_build_770
	}

	return false
}

func (dm_build_441 *dm_build_414) dm_build_440(dm_build_442 dm_build_817) bool {
	var dm_build_443 = dm_build_441.dm_build_418.dmConnector.compress
	if dm_build_442.dm_build_832() == Dm_build_724 || dm_build_443 == Dm_build_773 {
		return false
	}

	if dm_build_443 == Dm_build_771 {
		return true
	} else if dm_build_443 == Dm_build_772 {
		return dm_build_441.dm_build_417.Dm_build_345(Dm_build_732) == 1
	}

	return false
}

func (dm_build_445 *dm_build_414) dm_build_444(dm_build_446 dm_build_817) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_448 := dm_build_446.dm_build_830()

	if dm_build_448 > 0 {

		if dm_build_445.dm_build_436(dm_build_446) {
			var retBytes, err = Compress(dm_build_445.dm_build_417, Dm_build_725, int(dm_build_448), int(dm_build_445.dm_build_418.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_445.dm_build_417.Dm_build_92(Dm_build_725)

			dm_build_445.dm_build_417.Dm_build_133(dm_build_448)

			dm_build_445.dm_build_417.Dm_build_161(retBytes)

			dm_build_446.dm_build_831(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_445.dm_build_417.Dm_build_265(Dm_build_732, 1)
		}

		if dm_build_445.dm_build_421 {
			dm_build_448 = dm_build_446.dm_build_830()
			var retBytes = dm_build_445.dm_build_419.Encrypt(dm_build_445.dm_build_417.Dm_build_372(Dm_build_725, int(dm_build_448)), true)

			dm_build_445.dm_build_417.Dm_build_92(Dm_build_725)

			dm_build_445.dm_build_417.Dm_build_161(retBytes)

			dm_build_446.dm_build_831(int32(len(retBytes)))
		}
	}

	if dm_build_445.dm_build_417.Dm_build_90() > Dm_build_697 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_446.dm_build_826()
	if dm_build_445.dm_build_679(dm_build_446) {
		if dm_build_445.dm_build_416 != nil {
			dm_build_445.dm_build_417.Dm_build_95(0)
			if _, err := dm_build_445.dm_build_417.Dm_build_114(dm_build_445.dm_build_416); err != nil {
				return err
			}
		}
	} else {
		dm_build_445.dm_build_417.Dm_build_95(0)
		if _, err := dm_build_445.dm_build_417.Dm_build_114(dm_build_445.dm_build_415); err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_450 *dm_build_414) dm_build_449(dm_build_451 dm_build_817) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_453 := int32(0)
	if dm_build_450.dm_build_679(dm_build_451) {
		if dm_build_450.dm_build_416 != nil {
			dm_build_450.dm_build_417.Dm_build_92(0)
			if _, err := dm_build_450.dm_build_417.Dm_build_108(dm_build_450.dm_build_416, Dm_build_725); err != nil {
				return err
			}

			dm_build_453 = dm_build_451.dm_build_830()
			if dm_build_453 > 0 {
				if _, err := dm_build_450.dm_build_417.Dm_build_108(dm_build_450.dm_build_416, int(dm_build_453)); err != nil {
					return err
				}
			}
		}
	} else {

		dm_build_450.dm_build_417.Dm_build_92(0)
		if _, err := dm_build_450.dm_build_417.Dm_build_108(dm_build_450.dm_build_415, Dm_build_725); err != nil {
			return err
		}
		dm_build_453 = dm_build_451.dm_build_830()

		if dm_build_453 > 0 {
			if _, err := dm_build_450.dm_build_417.Dm_build_108(dm_build_450.dm_build_415, int(dm_build_453)); err != nil {
				return err
			}
		}
	}

	_ = dm_build_451.dm_build_827()

	dm_build_453 = dm_build_451.dm_build_830()
	if dm_build_453 <= 0 {
		return nil
	}

	if dm_build_450.dm_build_421 {
		eBytes := dm_build_450.dm_build_417.Dm_build_372(Dm_build_725, int(dm_build_453))
		dBytes, err := dm_build_450.dm_build_419.Decrypt(eBytes, true)
		if err != nil {
			return err
		}
		dm_build_450.dm_build_417.Dm_build_92(Dm_build_725)
		dm_build_450.dm_build_417.Dm_build_161(dBytes)
		dm_build_451.dm_build_831(int32(len(dBytes)))
	}

	if dm_build_450.dm_build_440(dm_build_451) {

		dm_build_453 = dm_build_451.dm_build_830()
		cBytes := dm_build_450.dm_build_417.Dm_build_372(Dm_build_725+ULINT_SIZE, int(dm_build_453-ULINT_SIZE))
		uBytes, err := UnCompress(cBytes, int(dm_build_450.dm_build_418.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_450.dm_build_417.Dm_build_92(Dm_build_725)
		dm_build_450.dm_build_417.Dm_build_161(uBytes)
		dm_build_451.dm_build_831(int32(len(uBytes)))
	}
	return nil
}

func (dm_build_455 *dm_build_414) dm_build_454(dm_build_456 dm_build_817) (dm_build_457 interface{}, dm_build_458 error) {
	if dm_build_455.dm_build_425 {
		return nil, ECGO_CONNECTION_CLOSED.throw()
	}
	dm_build_459 := dm_build_455.dm_build_418
	dm_build_459.mu.Lock()
	defer dm_build_459.mu.Unlock()
	dm_build_458 = dm_build_456.dm_build_821(dm_build_456)
	if dm_build_458 != nil {
		return nil, dm_build_458
	}

	dm_build_458 = dm_build_455.dm_build_444(dm_build_456)
	if dm_build_458 != nil {
		return nil, dm_build_458
	}

	dm_build_458 = dm_build_455.dm_build_449(dm_build_456)
	if dm_build_458 != nil {
		return nil, dm_build_458
	}

	return dm_build_456.dm_build_825(dm_build_456)
}

func (dm_build_461 *dm_build_414) dm_build_460() (*dm_build_1273, error) {

	Dm_build_462 := dm_build_1279(dm_build_461)
	_, dm_build_463 := dm_build_461.dm_build_454(Dm_build_462)
	if dm_build_463 != nil {
		return nil, dm_build_463
	}

	return Dm_build_462, nil
}

func (dm_build_465 *dm_build_414) dm_build_464() error {

	dm_build_466 := dm_build_1140(dm_build_465)
	_, dm_build_467 := dm_build_465.dm_build_454(dm_build_466)
	if dm_build_467 != nil {
		return dm_build_467
	}

	return nil
}

func (dm_build_469 *dm_build_414) dm_build_468() error {

	var dm_build_470 *dm_build_1273
	var err error
	if dm_build_470, err = dm_build_469.dm_build_460(); err != nil {
		return err
	}

	if dm_build_469.dm_build_418.sslEncrypt == 2 {
		if err = dm_build_469.dm_build_675(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_469.dm_build_418.sslEncrypt == 1 {
		if err = dm_build_469.dm_build_675(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_469.dm_build_421 || dm_build_469.dm_build_420 {
		k, err := dm_build_469.dm_build_665()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_470.Dm_build_1277)
		encryptType := dm_build_470.dm_build_1275
		hashType := int(dm_build_470.Dm_build_1276)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_469.dm_build_668(encryptType, sessionKey, dm_build_469.dm_build_418.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_469.dm_build_464(); err != nil {
		return err
	}
	return nil
}

func (dm_build_473 *dm_build_414) Dm_build_472(dm_build_474 *DmStatement) error {
	dm_build_475 := dm_build_1302(dm_build_473, dm_build_474)
	_, dm_build_476 := dm_build_473.dm_build_454(dm_build_475)
	if dm_build_476 != nil {
		return dm_build_476
	}

	return nil
}

func (dm_build_478 *dm_build_414) Dm_build_477(dm_build_479 int32) error {
	dm_build_480 := dm_build_1312(dm_build_478, dm_build_479)
	_, dm_build_481 := dm_build_478.dm_build_454(dm_build_480)
	if dm_build_481 != nil {
		return dm_build_481
	}

	return nil
}

func (dm_build_483 *dm_build_414) Dm_build_482(dm_build_484 *DmStatement, dm_build_485 bool, dm_build_486 int16) (*execRetInfo, error) {
	dm_build_487 := dm_build_1179(dm_build_483, dm_build_484, dm_build_485, dm_build_486)
	dm_build_488, dm_build_489 := dm_build_483.dm_build_454(dm_build_487)
	if dm_build_489 != nil {
		return nil, dm_build_489
	}
	return dm_build_488.(*execRetInfo), nil
}

func (dm_build_491 *dm_build_414) Dm_build_490(dm_build_492 *DmStatement, dm_build_493 int16) (*execRetInfo, error) {
	return dm_build_491.Dm_build_482(dm_build_492, false, Dm_build_777)
}

func (dm_build_495 *dm_build_414) Dm_build_494(dm_build_496 *DmStatement, dm_build_497 []OptParameter) (*execRetInfo, error) {
	dm_build_498, dm_build_499 := dm_build_495.dm_build_454(dm_build_920(dm_build_495, dm_build_496, dm_build_497))
	if dm_build_499 != nil {
		return nil, dm_build_499
	}

	return dm_build_498.(*execRetInfo), nil
}

func (dm_build_501 *dm_build_414) Dm_build_500(dm_build_502 *DmStatement, dm_build_503 int16) (*execRetInfo, error) {
	return dm_build_501.Dm_build_482(dm_build_502, true, dm_build_503)
}

func (dm_build_505 *dm_build_414) Dm_build_504(dm_build_506 *DmStatement, dm_build_507 [][]interface{}) (*execRetInfo, error) {
	dm_build_508 := dm_build_952(dm_build_505, dm_build_506, dm_build_507)
	dm_build_509, dm_build_510 := dm_build_505.dm_build_454(dm_build_508)
	if dm_build_510 != nil {
		return nil, dm_build_510
	}
	return dm_build_509.(*execRetInfo), nil
}

func (dm_build_512 *dm_build_414) Dm_build_511(dm_build_513 *DmStatement, dm_build_514 [][]interface{}, dm_build_515 bool) (*execRetInfo, error) {
	var dm_build_516, dm_build_517 = 0, 0
	var dm_build_518 = len(dm_build_514)
	var dm_build_519 [][]interface{}
	var dm_build_520 = NewExceInfo()
	dm_build_520.updateCounts = make([]int64, dm_build_518)
	var dm_build_521 = false
	for dm_build_516 < dm_build_518 {
		for dm_build_517 = dm_build_516; dm_build_517 < dm_build_518; dm_build_517++ {
			paramData := dm_build_514[dm_build_517]
			bindData := make([]interface{}, dm_build_513.paramCount)
			dm_build_521 = false
			for iCol := 0; iCol < int(dm_build_513.paramCount); iCol++ {
				if dm_build_513.bindParams[iCol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_512.dm_build_648(bindData, paramData, iCol) {
					dm_build_521 = true
					break
				}
			}

			if dm_build_521 {
				break
			}
			dm_build_519 = append(dm_build_519, bindData)
		}

		if dm_build_517 != dm_build_516 {
			tmpExecInfo, err := dm_build_512.Dm_build_504(dm_build_513, dm_build_519)
			if err != nil {
				return nil, err
			}
			dm_build_519 = dm_build_519[0:0]
			dm_build_520.union(tmpExecInfo, dm_build_516, dm_build_517-dm_build_516)
		}

		if dm_build_517 < dm_build_518 {
			tmpExecInfo, err := dm_build_512.Dm_build_522(dm_build_513, dm_build_514[dm_build_517], dm_build_515)
			if err != nil {
				return nil, err
			}

			dm_build_515 = true
			dm_build_520.union(tmpExecInfo, dm_build_517, 1)
		}

		dm_build_516 = dm_build_517 + 1
	}
	for _, i := range dm_build_520.updateCounts {
		if i > 0 {
			dm_build_520.updateCount += i
		}
	}
	return dm_build_520, nil
}

func (dm_build_523 *dm_build_414) Dm_build_522(dm_build_524 *DmStatement, dm_build_525 []interface{}, dm_build_526 bool) (*execRetInfo, error) {

	var dm_build_527 = make([]interface{}, dm_build_524.paramCount)
	for iCol := 0; iCol < int(dm_build_524.paramCount); iCol++ {
		if dm_build_524.bindParams[iCol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_523.dm_build_648(dm_build_527, dm_build_525, iCol) {

			if !dm_build_526 {
				preExecute := dm_build_1168(dm_build_523, dm_build_524, dm_build_524.bindParams)
				_, _ = dm_build_523.dm_build_454(preExecute)
				dm_build_526 = true
			}

			_ = dm_build_523.dm_build_654(dm_build_524, dm_build_524.bindParams[iCol], iCol, dm_build_525[iCol].(iOffRowBinder))
			dm_build_527[iCol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_528 = make([][]interface{}, 1)
	dm_build_528[0] = dm_build_527

	dm_build_529 := dm_build_952(dm_build_523, dm_build_524, dm_build_528)
	dm_build_530, dm_build_531 := dm_build_523.dm_build_454(dm_build_529)
	if dm_build_531 != nil {
		return nil, dm_build_531
	}
	return dm_build_530.(*execRetInfo), nil
}

func (dm_build_533 *dm_build_414) Dm_build_532(dm_build_534 *DmStatement, dm_build_535 int16) (*execRetInfo, error) {
	dm_build_536 := dm_build_1155(dm_build_533, dm_build_534, dm_build_535)

	dm_build_537, dm_build_538 := dm_build_533.dm_build_454(dm_build_536)
	if dm_build_538 != nil {
		return nil, dm_build_538
	}
	return dm_build_537.(*execRetInfo), nil
}

func (dm_build_540 *dm_build_414) Dm_build_539(dm_build_541 *innerRows, dm_build_542 int64) (*execRetInfo, error) {
	dm_build_543 := dm_build_1058(dm_build_540, dm_build_541, dm_build_542, INT64_MAX)
	dm_build_544, dm_build_545 := dm_build_540.dm_build_454(dm_build_543)
	if dm_build_545 != nil {
		return nil, dm_build_545
	}
	return dm_build_544.(*execRetInfo), nil
}

func (dm_build_547 *dm_build_414) Commit() error {
	dm_build_548 := dm_build_905(dm_build_547)
	_, dm_build_549 := dm_build_547.dm_build_454(dm_build_548)
	if dm_build_549 != nil {
		return dm_build_549
	}

	return nil
}

func (dm_build_551 *dm_build_414) Rollback() error {
	dm_build_552 := dm_build_1217(dm_build_551)
	_, dm_build_553 := dm_build_551.dm_build_454(dm_build_552)
	if dm_build_553 != nil {
		return dm_build_553
	}

	return nil
}

func (dm_build_555 *dm_build_414) Dm_build_554(dm_build_556 *DmConnection) error {
	dm_build_557 := dm_build_1222(dm_build_555, dm_build_556.IsoLevel)
	_, dm_build_558 := dm_build_555.dm_build_454(dm_build_557)
	if dm_build_558 != nil {
		return dm_build_558
	}

	return nil
}

func (dm_build_560 *dm_build_414) Dm_build_559(dm_build_561 *DmStatement, dm_build_562 string) error {
	dm_build_563 := dm_build_910(dm_build_560, dm_build_561, dm_build_562)
	_, dm_build_564 := dm_build_560.dm_build_454(dm_build_563)
	if dm_build_564 != nil {
		return dm_build_564
	}

	return nil
}

func (dm_build_566 *dm_build_414) Dm_build_565(dm_build_567 []uint32) ([]int64, error) {
	dm_build_568 := dm_build_1320(dm_build_566, dm_build_567)
	dm_build_569, dm_build_570 := dm_build_566.dm_build_454(dm_build_568)
	if dm_build_570 != nil {
		return nil, dm_build_570
	}
	return dm_build_569.([]int64), nil
}

func (dm_build_572 *dm_build_414) Close() error {
	if dm_build_572.dm_build_425 {
		return nil
	}

	dm_build_573 := dm_build_572.dm_build_415.Close()
	if dm_build_573 != nil {
		return dm_build_573
	}

	dm_build_572.dm_build_418 = nil
	dm_build_572.dm_build_425 = true
	return nil
}

func (dm_build_575 *dm_build_414) dm_build_574(dm_build_576 *lob) (int64, error) {
	dm_build_577 := dm_build_1091(dm_build_575, dm_build_576)
	dm_build_578, dm_build_579 := dm_build_575.dm_build_454(dm_build_577)
	if dm_build_579 != nil {
		return 0, dm_build_579
	}
	return dm_build_578.(int64), nil
}

func (dm_build_581 *dm_build_414) dm_build_580(dm_build_582 *lob, dm_build_583 int32, dm_build_584 int32) (*lobRetInfo, error) {
	dm_build_585 := dm_build_1076(dm_build_581, dm_build_582, int(dm_build_583), int(dm_build_584))
	dm_build_586, dm_build_587 := dm_build_581.dm_build_454(dm_build_585)
	if dm_build_587 != nil {
		return nil, dm_build_587
	}
	return dm_build_586.(*lobRetInfo), nil
}

func (dm_build_589 *dm_build_414) dm_build_588(dm_build_590 *DmBlob, dm_build_591 int32, dm_build_592 int32) ([]byte, error) {
	var dm_build_593 = make([]byte, dm_build_592)
	var dm_build_594 int32 = 0
	var dm_build_595 int32 = 0
	var dm_build_596 *lobRetInfo
	var dm_build_597 []byte
	var dm_build_598 error
	for dm_build_594 < dm_build_592 {
		dm_build_595 = dm_build_592 - dm_build_594
		if dm_build_595 > Dm_build_810 {
			dm_build_595 = Dm_build_810
		}
		dm_build_596, dm_build_598 = dm_build_589.dm_build_580(&dm_build_590.lob, dm_build_591+dm_build_594, dm_build_595)
		if dm_build_598 != nil {
			return nil, dm_build_598
		}
		dm_build_597 = dm_build_596.data
		if dm_build_597 == nil || len(dm_build_597) == 0 {
			break
		}
		Dm_build_1331.Dm_build_1387(dm_build_593, int(dm_build_594), dm_build_597, 0, len(dm_build_597))
		dm_build_594 += int32(len(dm_build_597))
		if dm_build_590.readOver {
			break
		}
	}
	return dm_build_593, nil
}

func (dm_build_600 *dm_build_414) dm_build_599(dm_build_601 *DmClob, dm_build_602 int32, dm_build_603 int32) (string, error) {
	var dm_build_604 bytes.Buffer
	var dm_build_605 int32 = 0
	var dm_build_606 int32 = 0
	var dm_build_607 *lobRetInfo
	var dm_build_608 []byte
	var dm_build_609 string
	var dm_build_610 error
	for dm_build_605 < dm_build_603 {
		dm_build_606 = dm_build_603 - dm_build_605
		if dm_build_606 > Dm_build_810/2 {
			dm_build_606 = Dm_build_810 / 2
		}
		dm_build_607, dm_build_610 = dm_build_600.dm_build_580(&dm_build_601.lob, dm_build_602+dm_build_605, dm_build_606)
		if dm_build_610 != nil {
			return "", dm_build_610
		}
		dm_build_608 = dm_build_607.data
		if dm_build_608 == nil || len(dm_build_608) == 0 {
			break
		}
		dm_build_609 = Dm_build_1331.Dm_build_1488(dm_build_608, 0, len(dm_build_608), dm_build_601.serverEncoding, dm_build_600.dm_build_418)

		dm_build_604.WriteString(dm_build_609)
		var strLen = dm_build_607.charLen
		if strLen == -1 {
			strLen = int64(utf8.RuneCountInString(dm_build_609))
		}
		dm_build_605 += int32(strLen)
		if dm_build_601.readOver {
			break
		}
	}
	return dm_build_604.String(), nil
}

func (dm_build_612 *dm_build_414) dm_build_611(dm_build_613 *DmClob, dm_build_614 int, dm_build_615 string, dm_build_616 string) (int, error) {
	var dm_build_617 = Dm_build_1331.Dm_build_1547(dm_build_615, dm_build_616, dm_build_612.dm_build_418)
	var dm_build_618 = 0
	var dm_build_619 = len(dm_build_617)
	var dm_build_620 = 0
	var dm_build_621 = 0
	var dm_build_622 = 0
	var dm_build_623 = dm_build_619/Dm_build_809 + 1
	var dm_build_624 byte = 0
	var dm_build_625 byte = 0x01
	var dm_build_626 byte = 0x02
	for i := 0; i < dm_build_623; i++ {
		dm_build_624 = 0
		if i == 0 {
			dm_build_624 |= dm_build_625
		}
		if i == dm_build_623-1 {
			dm_build_624 |= dm_build_626
		}
		dm_build_622 = dm_build_619 - dm_build_621
		if dm_build_622 > Dm_build_809 {
			dm_build_622 = Dm_build_809
		}

		setLobData := dm_build_1236(dm_build_612, &dm_build_613.lob, dm_build_624, dm_build_614, dm_build_617, dm_build_618, dm_build_622)
		ret, err := dm_build_612.dm_build_454(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if err != nil {
			return -1, err
		}
		if tmp <= 0 {
			return dm_build_620, nil
		} else {
			dm_build_614 += int(tmp)
			dm_build_620 += int(tmp)
			dm_build_621 += dm_build_622
			dm_build_618 += dm_build_622
		}
	}
	return dm_build_620, nil
}

func (dm_build_628 *dm_build_414) dm_build_627(dm_build_629 *DmBlob, dm_build_630 int, dm_build_631 []byte) (int, error) {
	var dm_build_632 = 0
	var dm_build_633 = len(dm_build_631)
	var dm_build_634 = 0
	var dm_build_635 = 0
	var dm_build_636 = 0
	var dm_build_637 = dm_build_633/Dm_build_809 + 1
	var dm_build_638 byte = 0
	var dm_build_639 byte = 0x01
	var dm_build_640 byte = 0x02
	for i := 0; i < dm_build_637; i++ {
		dm_build_638 = 0
		if i == 0 {
			dm_build_638 |= dm_build_639
		}
		if i == dm_build_637-1 {
			dm_build_638 |= dm_build_640
		}
		dm_build_636 = dm_build_633 - dm_build_635
		if dm_build_636 > Dm_build_809 {
			dm_build_636 = Dm_build_809
		}

		setLobData := dm_build_1236(dm_build_628, &dm_build_629.lob, dm_build_638, dm_build_630, dm_build_631, dm_build_632, dm_build_636)
		ret, err := dm_build_628.dm_build_454(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_634, nil
		} else {
			dm_build_630 += int(tmp)
			dm_build_634 += int(tmp)
			dm_build_635 += dm_build_636
			dm_build_632 += dm_build_636
		}
	}
	return dm_build_634, nil
}

func (dm_build_642 *dm_build_414) dm_build_641(dm_build_643 *lob, dm_build_644 int) (int64, error) {
	dm_build_645 := dm_build_1102(dm_build_642, dm_build_643, dm_build_644)
	dm_build_646, dm_build_647 := dm_build_642.dm_build_454(dm_build_645)
	if dm_build_647 != nil {
		return dm_build_643.length, dm_build_647
	}
	return dm_build_646.(int64), nil
}

func (dm_build_649 *dm_build_414) dm_build_648(dm_build_650 []interface{}, dm_build_651 []interface{}, dm_build_652 int) bool {
	var dm_build_653 = false
	dm_build_650[dm_build_652] = dm_build_651[dm_build_652]

	if binder, ok := dm_build_651[dm_build_652].(iOffRowBinder); ok {
		dm_build_653 = true
		dm_build_650[dm_build_652] = make([]byte, 0)
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_649.dm_build_418) {
			dm_build_650[dm_build_652] = &lobCtl{lob.buildCtlData()}
			dm_build_653 = false
		}
	} else {
		dm_build_650[dm_build_652] = dm_build_651[dm_build_652]
	}
	return dm_build_653
}

func (dm_build_655 *dm_build_414) dm_build_654(dm_build_656 *DmStatement, dm_build_657 parameter, dm_build_658 int, dm_build_659 iOffRowBinder) error {
	var dm_build_660 = Dm_build_4()
	dm_build_659.read(dm_build_660)
	var dm_build_661 = 0
	for !dm_build_659.isReadOver() || dm_build_660.Dm_build_5() > 0 {
		if !dm_build_659.isReadOver() && dm_build_660.Dm_build_5() < Dm_build_809 {
			dm_build_659.read(dm_build_660)
		}
		if dm_build_660.Dm_build_5() > Dm_build_809 {
			dm_build_661 = Dm_build_809
		} else {
			dm_build_661 = dm_build_660.Dm_build_5()
		}

		putData := dm_build_1207(dm_build_655, dm_build_656, int16(dm_build_658), dm_build_660, int32(dm_build_661))
		_, err := dm_build_655.dm_build_454(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_663 *dm_build_414) dm_build_662() ([]byte, error) {
	var dm_build_664 error
	if dm_build_663.dm_build_422 == nil {
		if dm_build_663.dm_build_422, dm_build_664 = security.NewClientKeyPair(); dm_build_664 != nil {
			return nil, dm_build_664
		}
	}
	return security.Bn2Bytes(dm_build_663.dm_build_422.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_666 *dm_build_414) dm_build_665() (*security.DhKey, error) {
	var dm_build_667 error
	if dm_build_666.dm_build_422 == nil {
		if dm_build_666.dm_build_422, dm_build_667 = security.NewClientKeyPair(); dm_build_667 != nil {
			return nil, dm_build_667
		}
	}
	return dm_build_666.dm_build_422, nil
}

func (dm_build_669 *dm_build_414) dm_build_668(dm_build_670 int, dm_build_671 []byte, dm_build_672 string, dm_build_673 int) (dm_build_674 error) {
	if dm_build_670 > 0 && dm_build_670 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_671 != nil {
		dm_build_669.dm_build_419, dm_build_674 = security.NewSymmCipher(dm_build_670, dm_build_671)
	} else if dm_build_670 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_669.dm_build_419, dm_build_674 = security.NewThirdPartCipher(dm_build_670, dm_build_671, dm_build_672, dm_build_673); dm_build_674 != nil {
			dm_build_674 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_674.Error()).throw()
		}
	}
	return
}

func (dm_build_676 *dm_build_414) dm_build_675(dm_build_677 bool) (dm_build_678 error) {
	if dm_build_676.dm_build_416, dm_build_678 = security.NewTLSFromTCP(dm_build_676.dm_build_415, dm_build_676.dm_build_418.dmConnector.sslCertPath, dm_build_676.dm_build_418.dmConnector.sslKeyPath, dm_build_676.dm_build_418.dmConnector.user); dm_build_678 != nil {
		return
	}
	if !dm_build_677 {
		dm_build_676.dm_build_416 = nil
	}
	return
}

func (dm_build_680 *dm_build_414) dm_build_679(dm_build_681 dm_build_817) bool {
	return dm_build_681.dm_build_832() != Dm_build_724 && dm_build_680.dm_build_418.sslEncrypt == 1
}
