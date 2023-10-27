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

type dm_build_1330 struct{}

var Dm_build_1331 = &dm_build_1330{}

func (Dm_build_1333 *dm_build_1330) Dm_build_1332(dm_build_1334 []byte, dm_build_1335 int, dm_build_1336 byte) int {
	dm_build_1334[dm_build_1335] = dm_build_1336
	return 1
}

func (Dm_build_1338 *dm_build_1330) Dm_build_1337(dm_build_1339 []byte, dm_build_1340 int, dm_build_1341 int8) int {
	dm_build_1339[dm_build_1340] = byte(dm_build_1341)
	return 1
}

func (Dm_build_1343 *dm_build_1330) Dm_build_1342(dm_build_1344 []byte, dm_build_1345 int, dm_build_1346 int16) int {
	dm_build_1344[dm_build_1345] = byte(dm_build_1346)
	dm_build_1345++
	dm_build_1344[dm_build_1345] = byte(dm_build_1346 >> 8)
	return 2
}

func (Dm_build_1348 *dm_build_1330) Dm_build_1347(dm_build_1349 []byte, dm_build_1350 int, dm_build_1351 int32) int {
	dm_build_1349[dm_build_1350] = byte(dm_build_1351)
	dm_build_1350++
	dm_build_1349[dm_build_1350] = byte(dm_build_1351 >> 8)
	dm_build_1350++
	dm_build_1349[dm_build_1350] = byte(dm_build_1351 >> 16)
	dm_build_1350++
	dm_build_1349[dm_build_1350] = byte(dm_build_1351 >> 24)
	dm_build_1350++
	return 4
}

func (Dm_build_1353 *dm_build_1330) Dm_build_1352(dm_build_1354 []byte, dm_build_1355 int, dm_build_1356 int64) int {
	dm_build_1354[dm_build_1355] = byte(dm_build_1356)
	dm_build_1355++
	dm_build_1354[dm_build_1355] = byte(dm_build_1356 >> 8)
	dm_build_1355++
	dm_build_1354[dm_build_1355] = byte(dm_build_1356 >> 16)
	dm_build_1355++
	dm_build_1354[dm_build_1355] = byte(dm_build_1356 >> 24)
	dm_build_1355++
	dm_build_1354[dm_build_1355] = byte(dm_build_1356 >> 32)
	dm_build_1355++
	dm_build_1354[dm_build_1355] = byte(dm_build_1356 >> 40)
	dm_build_1355++
	dm_build_1354[dm_build_1355] = byte(dm_build_1356 >> 48)
	dm_build_1355++
	dm_build_1354[dm_build_1355] = byte(dm_build_1356 >> 56)
	return 8
}

func (Dm_build_1358 *dm_build_1330) Dm_build_1357(dm_build_1359 []byte, dm_build_1360 int, dm_build_1361 float32) int {
	return Dm_build_1358.Dm_build_1377(dm_build_1359, dm_build_1360, math.Float32bits(dm_build_1361))
}

func (Dm_build_1363 *dm_build_1330) Dm_build_1362(dm_build_1364 []byte, dm_build_1365 int, dm_build_1366 float64) int {
	return Dm_build_1363.Dm_build_1382(dm_build_1364, dm_build_1365, math.Float64bits(dm_build_1366))
}

func (Dm_build_1368 *dm_build_1330) Dm_build_1367(dm_build_1369 []byte, dm_build_1370 int, dm_build_1371 uint8) int {
	dm_build_1369[dm_build_1370] = byte(dm_build_1371)
	return 1
}

func (Dm_build_1373 *dm_build_1330) Dm_build_1372(dm_build_1374 []byte, dm_build_1375 int, dm_build_1376 uint16) int {
	dm_build_1374[dm_build_1375] = byte(dm_build_1376)
	dm_build_1375++
	dm_build_1374[dm_build_1375] = byte(dm_build_1376 >> 8)
	return 2
}

func (Dm_build_1378 *dm_build_1330) Dm_build_1377(dm_build_1379 []byte, dm_build_1380 int, dm_build_1381 uint32) int {
	dm_build_1379[dm_build_1380] = byte(dm_build_1381)
	dm_build_1380++
	dm_build_1379[dm_build_1380] = byte(dm_build_1381 >> 8)
	dm_build_1380++
	dm_build_1379[dm_build_1380] = byte(dm_build_1381 >> 16)
	dm_build_1380++
	dm_build_1379[dm_build_1380] = byte(dm_build_1381 >> 24)
	return 3
}

func (Dm_build_1383 *dm_build_1330) Dm_build_1382(dm_build_1384 []byte, dm_build_1385 int, dm_build_1386 uint64) int {
	dm_build_1384[dm_build_1385] = byte(dm_build_1386)
	dm_build_1385++
	dm_build_1384[dm_build_1385] = byte(dm_build_1386 >> 8)
	dm_build_1385++
	dm_build_1384[dm_build_1385] = byte(dm_build_1386 >> 16)
	dm_build_1385++
	dm_build_1384[dm_build_1385] = byte(dm_build_1386 >> 24)
	dm_build_1385++
	dm_build_1384[dm_build_1385] = byte(dm_build_1386 >> 32)
	dm_build_1385++
	dm_build_1384[dm_build_1385] = byte(dm_build_1386 >> 40)
	dm_build_1385++
	dm_build_1384[dm_build_1385] = byte(dm_build_1386 >> 48)
	dm_build_1385++
	dm_build_1384[dm_build_1385] = byte(dm_build_1386 >> 56)
	return 3
}

func (Dm_build_1388 *dm_build_1330) Dm_build_1387(dm_build_1389 []byte, dm_build_1390 int, dm_build_1391 []byte, dm_build_1392 int, dm_build_1393 int) int {
	copy(dm_build_1389[dm_build_1390:dm_build_1390+dm_build_1393], dm_build_1391[dm_build_1392:dm_build_1392+dm_build_1393])
	return dm_build_1393
}

func (Dm_build_1395 *dm_build_1330) Dm_build_1394(dm_build_1396 []byte, dm_build_1397 int, dm_build_1398 []byte, dm_build_1399 int, dm_build_1400 int) int {
	dm_build_1397 += Dm_build_1395.Dm_build_1377(dm_build_1396, dm_build_1397, uint32(dm_build_1400))
	return 4 + Dm_build_1395.Dm_build_1387(dm_build_1396, dm_build_1397, dm_build_1398, dm_build_1399, dm_build_1400)
}

func (Dm_build_1402 *dm_build_1330) Dm_build_1401(dm_build_1403 []byte, dm_build_1404 int, dm_build_1405 []byte, dm_build_1406 int, dm_build_1407 int) int {
	dm_build_1404 += Dm_build_1402.Dm_build_1372(dm_build_1403, dm_build_1404, uint16(dm_build_1407))
	return 2 + Dm_build_1402.Dm_build_1387(dm_build_1403, dm_build_1404, dm_build_1405, dm_build_1406, dm_build_1407)
}

func (Dm_build_1409 *dm_build_1330) Dm_build_1408(dm_build_1410 []byte, dm_build_1411 int, dm_build_1412 string, dm_build_1413 string, dm_build_1414 *DmConnection) int {
	dm_build_1415 := Dm_build_1409.Dm_build_1547(dm_build_1412, dm_build_1413, dm_build_1414)
	dm_build_1411 += Dm_build_1409.Dm_build_1377(dm_build_1410, dm_build_1411, uint32(len(dm_build_1415)))
	return 4 + Dm_build_1409.Dm_build_1387(dm_build_1410, dm_build_1411, dm_build_1415, 0, len(dm_build_1415))
}

func (Dm_build_1417 *dm_build_1330) Dm_build_1416(dm_build_1418 []byte, dm_build_1419 int, dm_build_1420 string, dm_build_1421 string, dm_build_1422 *DmConnection) int {
	dm_build_1423 := Dm_build_1417.Dm_build_1547(dm_build_1420, dm_build_1421, dm_build_1422)

	dm_build_1419 += Dm_build_1417.Dm_build_1372(dm_build_1418, dm_build_1419, uint16(len(dm_build_1423)))
	return 2 + Dm_build_1417.Dm_build_1387(dm_build_1418, dm_build_1419, dm_build_1423, 0, len(dm_build_1423))
}

func (Dm_build_1425 *dm_build_1330) Dm_build_1424(dm_build_1426 []byte, dm_build_1427 int) byte {
	return dm_build_1426[dm_build_1427]
}

func (Dm_build_1429 *dm_build_1330) Dm_build_1428(dm_build_1430 []byte, dm_build_1431 int) int16 {
	var dm_build_1432 int16
	dm_build_1432 = int16(dm_build_1430[dm_build_1431] & 0xff)
	dm_build_1431++
	dm_build_1432 |= int16(dm_build_1430[dm_build_1431]&0xff) << 8
	return dm_build_1432
}

func (Dm_build_1434 *dm_build_1330) Dm_build_1433(dm_build_1435 []byte, dm_build_1436 int) int32 {
	var dm_build_1437 int32
	dm_build_1437 = int32(dm_build_1435[dm_build_1436] & 0xff)
	dm_build_1436++
	dm_build_1437 |= int32(dm_build_1435[dm_build_1436]&0xff) << 8
	dm_build_1436++
	dm_build_1437 |= int32(dm_build_1435[dm_build_1436]&0xff) << 16
	dm_build_1436++
	dm_build_1437 |= int32(dm_build_1435[dm_build_1436]&0xff) << 24
	return dm_build_1437
}

func (Dm_build_1439 *dm_build_1330) Dm_build_1438(dm_build_1440 []byte, dm_build_1441 int) int64 {
	var dm_build_1442 int64
	dm_build_1442 = int64(dm_build_1440[dm_build_1441] & 0xff)
	dm_build_1441++
	dm_build_1442 |= int64(dm_build_1440[dm_build_1441]&0xff) << 8
	dm_build_1441++
	dm_build_1442 |= int64(dm_build_1440[dm_build_1441]&0xff) << 16
	dm_build_1441++
	dm_build_1442 |= int64(dm_build_1440[dm_build_1441]&0xff) << 24
	dm_build_1441++
	dm_build_1442 |= int64(dm_build_1440[dm_build_1441]&0xff) << 32
	dm_build_1441++
	dm_build_1442 |= int64(dm_build_1440[dm_build_1441]&0xff) << 40
	dm_build_1441++
	dm_build_1442 |= int64(dm_build_1440[dm_build_1441]&0xff) << 48
	dm_build_1441++
	dm_build_1442 |= int64(dm_build_1440[dm_build_1441]&0xff) << 56
	return dm_build_1442
}

func (Dm_build_1444 *dm_build_1330) Dm_build_1443(dm_build_1445 []byte, dm_build_1446 int) float32 {
	return math.Float32frombits(Dm_build_1444.Dm_build_1460(dm_build_1445, dm_build_1446))
}

func (Dm_build_1448 *dm_build_1330) Dm_build_1447(dm_build_1449 []byte, dm_build_1450 int) float64 {
	return math.Float64frombits(Dm_build_1448.Dm_build_1465(dm_build_1449, dm_build_1450))
}

func (Dm_build_1452 *dm_build_1330) Dm_build_1451(dm_build_1453 []byte, dm_build_1454 int) uint8 {
	return uint8(dm_build_1453[dm_build_1454] & 0xff)
}

func (Dm_build_1456 *dm_build_1330) Dm_build_1455(dm_build_1457 []byte, dm_build_1458 int) uint16 {
	var dm_build_1459 uint16
	dm_build_1459 = uint16(dm_build_1457[dm_build_1458] & 0xff)
	dm_build_1458++
	dm_build_1459 |= uint16(dm_build_1457[dm_build_1458]&0xff) << 8
	return dm_build_1459
}

func (Dm_build_1461 *dm_build_1330) Dm_build_1460(dm_build_1462 []byte, dm_build_1463 int) uint32 {
	var dm_build_1464 uint32
	dm_build_1464 = uint32(dm_build_1462[dm_build_1463] & 0xff)
	dm_build_1463++
	dm_build_1464 |= uint32(dm_build_1462[dm_build_1463]&0xff) << 8
	dm_build_1463++
	dm_build_1464 |= uint32(dm_build_1462[dm_build_1463]&0xff) << 16
	dm_build_1463++
	dm_build_1464 |= uint32(dm_build_1462[dm_build_1463]&0xff) << 24
	return dm_build_1464
}

func (Dm_build_1466 *dm_build_1330) Dm_build_1465(dm_build_1467 []byte, dm_build_1468 int) uint64 {
	var dm_build_1469 uint64
	dm_build_1469 = uint64(dm_build_1467[dm_build_1468] & 0xff)
	dm_build_1468++
	dm_build_1469 |= uint64(dm_build_1467[dm_build_1468]&0xff) << 8
	dm_build_1468++
	dm_build_1469 |= uint64(dm_build_1467[dm_build_1468]&0xff) << 16
	dm_build_1468++
	dm_build_1469 |= uint64(dm_build_1467[dm_build_1468]&0xff) << 24
	dm_build_1468++
	dm_build_1469 |= uint64(dm_build_1467[dm_build_1468]&0xff) << 32
	dm_build_1468++
	dm_build_1469 |= uint64(dm_build_1467[dm_build_1468]&0xff) << 40
	dm_build_1468++
	dm_build_1469 |= uint64(dm_build_1467[dm_build_1468]&0xff) << 48
	dm_build_1468++
	dm_build_1469 |= uint64(dm_build_1467[dm_build_1468]&0xff) << 56
	return dm_build_1469
}

func (Dm_build_1471 *dm_build_1330) Dm_build_1470(dm_build_1472 []byte, dm_build_1473 int) []byte {
	dm_build_1474 := Dm_build_1471.Dm_build_1460(dm_build_1472, dm_build_1473)

	dm_build_1475 := make([]byte, dm_build_1474)
	copy(dm_build_1475[:int(dm_build_1474)], dm_build_1472[dm_build_1473+4:dm_build_1473+4+int(dm_build_1474)])
	return dm_build_1475
}

func (Dm_build_1477 *dm_build_1330) Dm_build_1476(dm_build_1478 []byte, dm_build_1479 int) []byte {
	dm_build_1480 := Dm_build_1477.Dm_build_1455(dm_build_1478, dm_build_1479)

	dm_build_1481 := make([]byte, dm_build_1480)
	copy(dm_build_1481[:int(dm_build_1480)], dm_build_1478[dm_build_1479+2:dm_build_1479+2+int(dm_build_1480)])
	return dm_build_1481
}

func (Dm_build_1483 *dm_build_1330) Dm_build_1482(dm_build_1484 []byte, dm_build_1485 int, dm_build_1486 int) []byte {

	dm_build_1487 := make([]byte, dm_build_1486)
	copy(dm_build_1487[:dm_build_1486], dm_build_1484[dm_build_1485:dm_build_1485+dm_build_1486])
	return dm_build_1487
}

func (Dm_build_1489 *dm_build_1330) Dm_build_1488(dm_build_1490 []byte, dm_build_1491 int, dm_build_1492 int, dm_build_1493 string, dm_build_1494 *DmConnection) string {
	return Dm_build_1489.Dm_build_1584(dm_build_1490[dm_build_1491:dm_build_1491+dm_build_1492], dm_build_1493, dm_build_1494)
}

func (Dm_build_1496 *dm_build_1330) Dm_build_1495(dm_build_1497 []byte, dm_build_1498 int, dm_build_1499 string, dm_build_1500 *DmConnection) string {
	dm_build_1501 := Dm_build_1496.Dm_build_1460(dm_build_1497, dm_build_1498)
	dm_build_1498 += 4
	return Dm_build_1496.Dm_build_1488(dm_build_1497, dm_build_1498, int(dm_build_1501), dm_build_1499, dm_build_1500)
}

func (Dm_build_1503 *dm_build_1330) Dm_build_1502(dm_build_1504 []byte, dm_build_1505 int, dm_build_1506 string, dm_build_1507 *DmConnection) string {
	dm_build_1508 := Dm_build_1503.Dm_build_1455(dm_build_1504, dm_build_1505)
	dm_build_1505 += 2
	return Dm_build_1503.Dm_build_1488(dm_build_1504, dm_build_1505, int(dm_build_1508), dm_build_1506, dm_build_1507)
}

func (Dm_build_1510 *dm_build_1330) Dm_build_1509(dm_build_1511 byte) []byte {
	return []byte{dm_build_1511}
}

func (Dm_build_1513 *dm_build_1330) Dm_build_1512(dm_build_1514 int8) []byte {
	return []byte{byte(dm_build_1514)}
}

func (Dm_build_1516 *dm_build_1330) Dm_build_1515(dm_build_1517 int16) []byte {
	return []byte{byte(dm_build_1517), byte(dm_build_1517 >> 8)}
}

func (Dm_build_1519 *dm_build_1330) Dm_build_1518(dm_build_1520 int32) []byte {
	return []byte{byte(dm_build_1520), byte(dm_build_1520 >> 8), byte(dm_build_1520 >> 16), byte(dm_build_1520 >> 24)}
}

func (Dm_build_1522 *dm_build_1330) Dm_build_1521(dm_build_1523 int64) []byte {
	return []byte{byte(dm_build_1523), byte(dm_build_1523 >> 8), byte(dm_build_1523 >> 16), byte(dm_build_1523 >> 24), byte(dm_build_1523 >> 32),
		byte(dm_build_1523 >> 40), byte(dm_build_1523 >> 48), byte(dm_build_1523 >> 56)}
}

func (Dm_build_1525 *dm_build_1330) Dm_build_1524(dm_build_1526 float32) []byte {
	return Dm_build_1525.Dm_build_1536(math.Float32bits(dm_build_1526))
}

func (Dm_build_1528 *dm_build_1330) Dm_build_1527(dm_build_1529 float64) []byte {
	return Dm_build_1528.Dm_build_1539(math.Float64bits(dm_build_1529))
}

func (Dm_build_1531 *dm_build_1330) Dm_build_1530(dm_build_1532 uint8) []byte {
	return []byte{byte(dm_build_1532)}
}

func (Dm_build_1534 *dm_build_1330) Dm_build_1533(dm_build_1535 uint16) []byte {
	return []byte{byte(dm_build_1535), byte(dm_build_1535 >> 8)}
}

func (Dm_build_1537 *dm_build_1330) Dm_build_1536(dm_build_1538 uint32) []byte {
	return []byte{byte(dm_build_1538), byte(dm_build_1538 >> 8), byte(dm_build_1538 >> 16), byte(dm_build_1538 >> 24)}
}

func (Dm_build_1540 *dm_build_1330) Dm_build_1539(dm_build_1541 uint64) []byte {
	return []byte{byte(dm_build_1541), byte(dm_build_1541 >> 8), byte(dm_build_1541 >> 16), byte(dm_build_1541 >> 24), byte(dm_build_1541 >> 32), byte(dm_build_1541 >> 40), byte(dm_build_1541 >> 48), byte(dm_build_1541 >> 56)}
}

func (Dm_build_1543 *dm_build_1330) Dm_build_1542(dm_build_1544 []byte, dm_build_1545 string, dm_build_1546 *DmConnection) []byte {
	if dm_build_1545 == "UTF-8" {
		return dm_build_1544
	}

	if dm_build_1546 == nil {
		if e := dm_build_1589(dm_build_1545); e != nil {
			tmp, err := io.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1544), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1546.encodeBuffer == nil {
		dm_build_1546.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1546.encode = dm_build_1589(dm_build_1546.getServerEncoding())
		dm_build_1546.transformReaderDst = make([]byte, 4096)
		dm_build_1546.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1546.encode; e != nil {

		dm_build_1546.encodeBuffer.Reset()

		n, err := dm_build_1546.encodeBuffer.ReadFrom(
			Dm_build_1603(bytes.NewReader(dm_build_1544), e.NewEncoder(), dm_build_1546.transformReaderDst, dm_build_1546.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1546.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1548 *dm_build_1330) Dm_build_1547(dm_build_1549 string, dm_build_1550 string, dm_build_1551 *DmConnection) []byte {
	return Dm_build_1548.Dm_build_1542([]byte(dm_build_1549), dm_build_1550, dm_build_1551)
}

func (Dm_build_1553 *dm_build_1330) Dm_build_1552(dm_build_1554 []byte) byte {
	return Dm_build_1553.Dm_build_1424(dm_build_1554, 0)
}

func (Dm_build_1556 *dm_build_1330) Dm_build_1555(dm_build_1557 []byte) int16 {
	return Dm_build_1556.Dm_build_1428(dm_build_1557, 0)
}

func (Dm_build_1559 *dm_build_1330) Dm_build_1558(dm_build_1560 []byte) int32 {
	return Dm_build_1559.Dm_build_1433(dm_build_1560, 0)
}

func (Dm_build_1562 *dm_build_1330) Dm_build_1561(dm_build_1563 []byte) int64 {
	return Dm_build_1562.Dm_build_1438(dm_build_1563, 0)
}

func (Dm_build_1565 *dm_build_1330) Dm_build_1564(dm_build_1566 []byte) float32 {
	return Dm_build_1565.Dm_build_1443(dm_build_1566, 0)
}

func (Dm_build_1568 *dm_build_1330) Dm_build_1567(dm_build_1569 []byte) float64 {
	return Dm_build_1568.Dm_build_1447(dm_build_1569, 0)
}

func (Dm_build_1571 *dm_build_1330) Dm_build_1570(dm_build_1572 []byte) uint8 {
	return Dm_build_1571.Dm_build_1451(dm_build_1572, 0)
}

func (Dm_build_1574 *dm_build_1330) Dm_build_1573(dm_build_1575 []byte) uint16 {
	return Dm_build_1574.Dm_build_1455(dm_build_1575, 0)
}

func (Dm_build_1577 *dm_build_1330) Dm_build_1576(dm_build_1578 []byte) uint32 {
	return Dm_build_1577.Dm_build_1460(dm_build_1578, 0)
}

func (Dm_build_1580 *dm_build_1330) Dm_build_1579(dm_build_1581 []byte, dm_build_1582 string, dm_build_1583 *DmConnection) []byte {
	if dm_build_1582 == "UTF-8" {
		return dm_build_1581
	}

	if dm_build_1583 == nil {
		if e := dm_build_1589(dm_build_1582); e != nil {

			tmp, err := io.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1581), e.NewDecoder()),
			)
			if err != nil {

				panic("Charset To UTF8 error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1583.encodeBuffer == nil {
		dm_build_1583.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1583.encode = dm_build_1589(dm_build_1583.getServerEncoding())
		dm_build_1583.transformReaderDst = make([]byte, 4096)
		dm_build_1583.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1583.encode; e != nil {

		dm_build_1583.encodeBuffer.Reset()

		n, err := dm_build_1583.encodeBuffer.ReadFrom(
			Dm_build_1603(bytes.NewReader(dm_build_1581), e.NewDecoder(), dm_build_1583.transformReaderDst, dm_build_1583.transformReaderSrc),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return dm_build_1583.encodeBuffer.Next(int(n))
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1585 *dm_build_1330) Dm_build_1584(dm_build_1586 []byte, dm_build_1587 string, dm_build_1588 *DmConnection) string {
	return string(Dm_build_1585.Dm_build_1579(dm_build_1586, dm_build_1587, dm_build_1588))
}

func dm_build_1589(dm_build_1590 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1590); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1591 struct {
	dm_build_1592 io.Reader
	dm_build_1593 transform.Transformer
	dm_build_1594 error

	dm_build_1595                []byte
	dm_build_1596, dm_build_1597 int

	dm_build_1598                []byte
	dm_build_1599, dm_build_1600 int

	dm_build_1601 bool
}

const dm_build_1602 = 4096

func Dm_build_1603(dm_build_1604 io.Reader, dm_build_1605 transform.Transformer, dm_build_1606 []byte, dm_build_1607 []byte) *Dm_build_1591 {
	dm_build_1605.Reset()
	return &Dm_build_1591{
		dm_build_1592: dm_build_1604,
		dm_build_1593: dm_build_1605,
		dm_build_1595: dm_build_1606,
		dm_build_1598: dm_build_1607,
	}
}

func (dm_build_1609 *Dm_build_1591) Read(dm_build_1610 []byte) (int, error) {
	dm_build_1611, dm_build_1612 := 0, error(nil)
	for {

		if dm_build_1609.dm_build_1596 != dm_build_1609.dm_build_1597 {
			dm_build_1611 = copy(dm_build_1610, dm_build_1609.dm_build_1595[dm_build_1609.dm_build_1596:dm_build_1609.dm_build_1597])
			dm_build_1609.dm_build_1596 += dm_build_1611
			if dm_build_1609.dm_build_1596 == dm_build_1609.dm_build_1597 && dm_build_1609.dm_build_1601 {
				return dm_build_1611, dm_build_1609.dm_build_1594
			}
			return dm_build_1611, nil
		} else if dm_build_1609.dm_build_1601 {
			return 0, dm_build_1609.dm_build_1594
		}

		if dm_build_1609.dm_build_1599 != dm_build_1609.dm_build_1600 || dm_build_1609.dm_build_1594 != nil {
			dm_build_1609.dm_build_1596 = 0
			dm_build_1609.dm_build_1597, dm_build_1611, dm_build_1612 = dm_build_1609.dm_build_1593.Transform(dm_build_1609.dm_build_1595, dm_build_1609.dm_build_1598[dm_build_1609.dm_build_1599:dm_build_1609.dm_build_1600], dm_build_1609.dm_build_1594 == io.EOF)
			dm_build_1609.dm_build_1599 += dm_build_1611

			switch {
			case dm_build_1612 == nil:
				if dm_build_1609.dm_build_1599 != dm_build_1609.dm_build_1600 {
					dm_build_1609.dm_build_1594 = nil
				}

				dm_build_1609.dm_build_1601 = dm_build_1609.dm_build_1594 != nil
				continue
			case errors.Is(dm_build_1612, transform.ErrShortDst) && (dm_build_1609.dm_build_1597 != 0 || dm_build_1611 != 0):

				continue
			case errors.Is(dm_build_1612, transform.ErrShortSrc) && dm_build_1609.dm_build_1600-dm_build_1609.dm_build_1599 != len(dm_build_1609.dm_build_1598) && dm_build_1609.dm_build_1594 == nil:

			default:
				dm_build_1609.dm_build_1601 = true

				if dm_build_1609.dm_build_1594 == nil || dm_build_1609.dm_build_1594 == io.EOF {
					dm_build_1609.dm_build_1594 = dm_build_1612
				}
				continue
			}
		}

		if dm_build_1609.dm_build_1599 != 0 {
			dm_build_1609.dm_build_1599, dm_build_1609.dm_build_1600 = 0, copy(dm_build_1609.dm_build_1598, dm_build_1609.dm_build_1598[dm_build_1609.dm_build_1599:dm_build_1609.dm_build_1600])
		}
		dm_build_1611, dm_build_1609.dm_build_1594 = dm_build_1609.dm_build_1592.Read(dm_build_1609.dm_build_1598[dm_build_1609.dm_build_1600:])
		dm_build_1609.dm_build_1600 += dm_build_1611
	}
}
