/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */

package dm8

import (
	"io"
	"math"
)

type Dm_build_1290 struct {
	dm_build_1291 []byte
	dm_build_1292 int
}

func Dm_build_1293(dm_build_1294 int) *Dm_build_1290 {
	return &Dm_build_1290{make([]byte, 0, dm_build_1294), 0}
}

func Dm_build_1295(dm_build_1296 []byte) *Dm_build_1290 {
	return &Dm_build_1290{dm_build_1296, 0}
}

func (dm_build_1298 *Dm_build_1290) dm_build_1297(dm_build_1299 int) *Dm_build_1290 {

	dm_build_1300 := len(dm_build_1298.dm_build_1291)
	dm_build_1301 := cap(dm_build_1298.dm_build_1291)

	if dm_build_1300+dm_build_1299 <= dm_build_1301 {
		dm_build_1298.dm_build_1291 = dm_build_1298.dm_build_1291[:dm_build_1300+dm_build_1299]
	} else {

		var calCap = int64(math.Max(float64(2*dm_build_1301), float64(dm_build_1299+dm_build_1300)))

		nbuf := make([]byte, dm_build_1299+dm_build_1300, calCap)
		copy(nbuf, dm_build_1298.dm_build_1291)
		dm_build_1298.dm_build_1291 = nbuf
	}

	return dm_build_1298
}

func (dm_build_1303 *Dm_build_1290) Dm_build_1302() int {
	return len(dm_build_1303.dm_build_1291)
}

func (dm_build_1305 *Dm_build_1290) Dm_build_1304(dm_build_1306 int) *Dm_build_1290 {
	for i := dm_build_1306; i < len(dm_build_1305.dm_build_1291); i++ {
		dm_build_1305.dm_build_1291[i] = 0
	}
	dm_build_1305.dm_build_1291 = dm_build_1305.dm_build_1291[:dm_build_1306]
	return dm_build_1305
}

func (dm_build_1308 *Dm_build_1290) Dm_build_1307(dm_build_1309 int) *Dm_build_1290 {
	dm_build_1308.dm_build_1292 = dm_build_1309
	return dm_build_1308
}

func (dm_build_1311 *Dm_build_1290) Dm_build_1310() int {
	return dm_build_1311.dm_build_1292
}

func (dm_build_1313 *Dm_build_1290) Dm_build_1312(_ bool) int {
	return len(dm_build_1313.dm_build_1291) - dm_build_1313.dm_build_1292
}

func (dm_build_1316 *Dm_build_1290) Dm_build_1315(dm_build_1317 int, dm_build_1318 bool, dm_build_1319 bool) *Dm_build_1290 {

	if dm_build_1318 {
		if dm_build_1319 {
			dm_build_1316.dm_build_1297(dm_build_1317)
		} else {
			dm_build_1316.dm_build_1291 = dm_build_1316.dm_build_1291[:len(dm_build_1316.dm_build_1291)-dm_build_1317]
		}
	} else {
		if dm_build_1319 {
			dm_build_1316.dm_build_1292 += dm_build_1317
		} else {
			dm_build_1316.dm_build_1292 -= dm_build_1317
		}
	}

	return dm_build_1316
}

func (dm_build_1321 *Dm_build_1290) Dm_build_1320(dm_build_1322 io.Reader, dm_build_1323 int) (int, error) {
	dm_build_1324 := len(dm_build_1321.dm_build_1291)
	dm_build_1321.dm_build_1297(dm_build_1323)
	dm_build_1325 := 0
	for dm_build_1323 > 0 {
		n, err := dm_build_1322.Read(dm_build_1321.dm_build_1291[dm_build_1324+dm_build_1325:])
		if n > 0 && err == io.EOF {
			dm_build_1325 += n
			dm_build_1321.dm_build_1291 = dm_build_1321.dm_build_1291[:dm_build_1324+dm_build_1325]
			return dm_build_1325, nil
		} else if n > 0 && err == nil {
			dm_build_1323 -= n
			dm_build_1325 += n
		} else if n == 0 && err != nil {
			return -1, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
		}
	}

	return dm_build_1325, nil
}

func (dm_build_1327 *Dm_build_1290) Dm_build_1326(dm_build_1328 io.Writer) (*Dm_build_1290, error) {
	if _, err := dm_build_1328.Write(dm_build_1327.dm_build_1291); err != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
	}
	return dm_build_1327, nil
}

func (dm_build_1330 *Dm_build_1290) Dm_build_1329(dm_build_1331 bool) int {
	dm_build_1332 := len(dm_build_1330.dm_build_1291)
	dm_build_1330.dm_build_1297(1)

	if dm_build_1331 {
		return copy(dm_build_1330.dm_build_1291[dm_build_1332:], []byte{1})
	} else {
		return copy(dm_build_1330.dm_build_1291[dm_build_1332:], []byte{0})
	}
}

func (dm_build_1334 *Dm_build_1290) Dm_build_1333(dm_build_1335 byte) int {
	dm_build_1336 := len(dm_build_1334.dm_build_1291)
	dm_build_1334.dm_build_1297(1)

	return copy(dm_build_1334.dm_build_1291[dm_build_1336:], Dm_build_931.Dm_build_1109(dm_build_1335))
}

func (dm_build_1338 *Dm_build_1290) Dm_build_1337(dm_build_1339 int8) int {
	dm_build_1340 := len(dm_build_1338.dm_build_1291)
	dm_build_1338.dm_build_1297(1)

	return copy(dm_build_1338.dm_build_1291[dm_build_1340:], Dm_build_931.Dm_build_1112(dm_build_1339))
}

func (dm_build_1342 *Dm_build_1290) Dm_build_1341(dm_build_1343 int16) int {
	dm_build_1344 := len(dm_build_1342.dm_build_1291)
	dm_build_1342.dm_build_1297(2)

	return copy(dm_build_1342.dm_build_1291[dm_build_1344:], Dm_build_931.Dm_build_1115(dm_build_1343))
}

func (dm_build_1346 *Dm_build_1290) Dm_build_1345(dm_build_1347 int32) int {
	dm_build_1348 := len(dm_build_1346.dm_build_1291)
	dm_build_1346.dm_build_1297(4)

	return copy(dm_build_1346.dm_build_1291[dm_build_1348:], Dm_build_931.Dm_build_1118(dm_build_1347))
}

func (dm_build_1350 *Dm_build_1290) Dm_build_1349(dm_build_1351 uint8) int {
	dm_build_1352 := len(dm_build_1350.dm_build_1291)
	dm_build_1350.dm_build_1297(1)

	return copy(dm_build_1350.dm_build_1291[dm_build_1352:], Dm_build_931.Dm_build_1130(dm_build_1351))
}

func (dm_build_1354 *Dm_build_1290) Dm_build_1353(dm_build_1355 uint16) int {
	dm_build_1356 := len(dm_build_1354.dm_build_1291)
	dm_build_1354.dm_build_1297(2)

	return copy(dm_build_1354.dm_build_1291[dm_build_1356:], Dm_build_931.Dm_build_1133(dm_build_1355))
}

func (dm_build_1358 *Dm_build_1290) Dm_build_1357(dm_build_1359 uint32) int {
	dm_build_1360 := len(dm_build_1358.dm_build_1291)
	dm_build_1358.dm_build_1297(4)

	return copy(dm_build_1358.dm_build_1291[dm_build_1360:], Dm_build_931.Dm_build_1136(dm_build_1359))
}

func (dm_build_1362 *Dm_build_1290) Dm_build_1361(dm_build_1363 uint64) int {
	dm_build_1364 := len(dm_build_1362.dm_build_1291)
	dm_build_1362.dm_build_1297(8)

	return copy(dm_build_1362.dm_build_1291[dm_build_1364:], Dm_build_931.Dm_build_1139(dm_build_1363))
}

func (dm_build_1366 *Dm_build_1290) Dm_build_1365(dm_build_1367 float32) int {
	dm_build_1368 := len(dm_build_1366.dm_build_1291)
	dm_build_1366.dm_build_1297(4)

	return copy(dm_build_1366.dm_build_1291[dm_build_1368:], Dm_build_931.Dm_build_1136(math.Float32bits(dm_build_1367)))
}

func (dm_build_1370 *Dm_build_1290) Dm_build_1369(dm_build_1371 float64) int {
	dm_build_1372 := len(dm_build_1370.dm_build_1291)
	dm_build_1370.dm_build_1297(8)

	return copy(dm_build_1370.dm_build_1291[dm_build_1372:], Dm_build_931.Dm_build_1139(math.Float64bits(dm_build_1371)))
}

func (dm_build_1374 *Dm_build_1290) Dm_build_1373(dm_build_1375 []byte) int {
	dm_build_1376 := len(dm_build_1374.dm_build_1291)
	dm_build_1374.dm_build_1297(len(dm_build_1375))
	return copy(dm_build_1374.dm_build_1291[dm_build_1376:], dm_build_1375)
}

func (dm_build_1378 *Dm_build_1290) Dm_build_1377(dm_build_1379 []byte) int {
	return dm_build_1378.Dm_build_1345(int32(len(dm_build_1379))) + dm_build_1378.Dm_build_1373(dm_build_1379)
}

func (dm_build_1381 *Dm_build_1290) Dm_build_1380(dm_build_1382 []byte) int {
	return dm_build_1381.Dm_build_1349(uint8(len(dm_build_1382))) + dm_build_1381.Dm_build_1373(dm_build_1382)
}

func (dm_build_1384 *Dm_build_1290) Dm_build_1383(dm_build_1385 []byte) int {
	return dm_build_1384.Dm_build_1353(uint16(len(dm_build_1385))) + dm_build_1384.Dm_build_1373(dm_build_1385)
}

func (dm_build_1387 *Dm_build_1290) Dm_build_1386(dm_build_1388 []byte) int {
	return dm_build_1387.Dm_build_1373(dm_build_1388) + dm_build_1387.Dm_build_1333(0)
}

func (dm_build_1390 *Dm_build_1290) Dm_build_1389(dm_build_1391 string, dm_build_1392 string, dm_build_1393 *DmConnection) int {
	dm_build_1394 := Dm_build_931.Dm_build_1147(dm_build_1391, dm_build_1392, dm_build_1393)
	return dm_build_1390.Dm_build_1377(dm_build_1394)
}

func (dm_build_1396 *Dm_build_1290) Dm_build_1395(dm_build_1397 string, dm_build_1398 string, dm_build_1399 *DmConnection) int {
	dm_build_1400 := Dm_build_931.Dm_build_1147(dm_build_1397, dm_build_1398, dm_build_1399)
	return dm_build_1396.Dm_build_1380(dm_build_1400)
}

func (dm_build_1402 *Dm_build_1290) Dm_build_1401(dm_build_1403 string, dm_build_1404 string, dm_build_1405 *DmConnection) int {
	dm_build_1406 := Dm_build_931.Dm_build_1147(dm_build_1403, dm_build_1404, dm_build_1405)
	return dm_build_1402.Dm_build_1383(dm_build_1406)
}

func (dm_build_1408 *Dm_build_1290) Dm_build_1407(dm_build_1409 string, dm_build_1410 string, dm_build_1411 *DmConnection) int {
	dm_build_1412 := Dm_build_931.Dm_build_1147(dm_build_1409, dm_build_1410, dm_build_1411)
	return dm_build_1408.Dm_build_1386(dm_build_1412)
}

func (dm_build_1414 *Dm_build_1290) Dm_build_1413() byte {
	dm_build_1415 := Dm_build_931.Dm_build_1024(dm_build_1414.dm_build_1291, dm_build_1414.dm_build_1292)
	dm_build_1414.dm_build_1292++
	return dm_build_1415
}

func (dm_build_1417 *Dm_build_1290) Dm_build_1416() int16 {
	dm_build_1418 := Dm_build_931.Dm_build_1028(dm_build_1417.dm_build_1291, dm_build_1417.dm_build_1292)
	dm_build_1417.dm_build_1292 += 2
	return dm_build_1418
}

func (dm_build_1420 *Dm_build_1290) Dm_build_1419() int32 {
	dm_build_1421 := Dm_build_931.Dm_build_1033(dm_build_1420.dm_build_1291, dm_build_1420.dm_build_1292)
	dm_build_1420.dm_build_1292 += 4
	return dm_build_1421
}

func (dm_build_1423 *Dm_build_1290) Dm_build_1422() int64 {
	dm_build_1424 := Dm_build_931.Dm_build_1038(dm_build_1423.dm_build_1291, dm_build_1423.dm_build_1292)
	dm_build_1423.dm_build_1292 += 8
	return dm_build_1424
}

func (dm_build_1426 *Dm_build_1290) Dm_build_1425() float32 {
	dm_build_1427 := Dm_build_931.Dm_build_1043(dm_build_1426.dm_build_1291, dm_build_1426.dm_build_1292)
	dm_build_1426.dm_build_1292 += 4
	return dm_build_1427
}

func (dm_build_1429 *Dm_build_1290) Dm_build_1428() float64 {
	dm_build_1430 := Dm_build_931.Dm_build_1047(dm_build_1429.dm_build_1291, dm_build_1429.dm_build_1292)
	dm_build_1429.dm_build_1292 += 8
	return dm_build_1430
}

func (dm_build_1432 *Dm_build_1290) Dm_build_1431() uint8 {
	dm_build_1433 := Dm_build_931.Dm_build_1051(dm_build_1432.dm_build_1291, dm_build_1432.dm_build_1292)
	dm_build_1432.dm_build_1292 += 1
	return dm_build_1433
}

func (dm_build_1435 *Dm_build_1290) Dm_build_1434() uint16 {
	dm_build_1436 := Dm_build_931.Dm_build_1055(dm_build_1435.dm_build_1291, dm_build_1435.dm_build_1292)
	dm_build_1435.dm_build_1292 += 2
	return dm_build_1436
}

func (dm_build_1438 *Dm_build_1290) Dm_build_1437() uint32 {
	dm_build_1439 := Dm_build_931.Dm_build_1060(dm_build_1438.dm_build_1291, dm_build_1438.dm_build_1292)
	dm_build_1438.dm_build_1292 += 4
	return dm_build_1439
}

func (dm_build_1441 *Dm_build_1290) Dm_build_1440(dm_build_1442 int) []byte {
	dm_build_1443 := Dm_build_931.Dm_build_1082(dm_build_1441.dm_build_1291, dm_build_1441.dm_build_1292, dm_build_1442)
	dm_build_1441.dm_build_1292 += dm_build_1442
	return dm_build_1443
}

func (dm_build_1445 *Dm_build_1290) Dm_build_1444() []byte {
	return dm_build_1445.Dm_build_1440(int(dm_build_1445.Dm_build_1419()))
}

func (dm_build_1447 *Dm_build_1290) Dm_build_1446() []byte {
	return dm_build_1447.Dm_build_1440(int(dm_build_1447.Dm_build_1413()))
}

func (dm_build_1449 *Dm_build_1290) Dm_build_1448() []byte {
	return dm_build_1449.Dm_build_1440(int(dm_build_1449.Dm_build_1416()))
}

func (dm_build_1451 *Dm_build_1290) Dm_build_1450(dm_build_1452 int) []byte {
	return dm_build_1451.Dm_build_1440(dm_build_1452)
}

func (dm_build_1454 *Dm_build_1290) Dm_build_1453() []byte {
	dm_build_1455 := 0
	for dm_build_1454.Dm_build_1413() != 0 {
		dm_build_1455++
	}
	dm_build_1454.Dm_build_1315(dm_build_1455, false, false)
	return dm_build_1454.Dm_build_1440(dm_build_1455)
}

func (dm_build_1457 *Dm_build_1290) Dm_build_1456(dm_build_1458 int, dm_build_1459 string, dm_build_1460 *DmConnection) string {
	return Dm_build_931.Dm_build_1183(dm_build_1457.Dm_build_1440(dm_build_1458), dm_build_1459, dm_build_1460)
}

func (dm_build_1462 *Dm_build_1290) Dm_build_1461(dm_build_1463 string, dm_build_1464 *DmConnection) string {
	return Dm_build_931.Dm_build_1183(dm_build_1462.Dm_build_1444(), dm_build_1463, dm_build_1464)
}

func (dm_build_1466 *Dm_build_1290) Dm_build_1465(dm_build_1467 string, dm_build_1468 *DmConnection) string {
	return Dm_build_931.Dm_build_1183(dm_build_1466.Dm_build_1446(), dm_build_1467, dm_build_1468)
}

func (dm_build_1470 *Dm_build_1290) Dm_build_1469(dm_build_1471 string, dm_build_1472 *DmConnection) string {
	return Dm_build_931.Dm_build_1183(dm_build_1470.Dm_build_1448(), dm_build_1471, dm_build_1472)
}

func (dm_build_1474 *Dm_build_1290) Dm_build_1473(dm_build_1475 string, dm_build_1476 *DmConnection) string {
	return Dm_build_931.Dm_build_1183(dm_build_1474.Dm_build_1453(), dm_build_1475, dm_build_1476)
}

func (dm_build_1478 *Dm_build_1290) Dm_build_1477(dm_build_1479 int, dm_build_1480 byte) int {
	return dm_build_1478.Dm_build_1513(dm_build_1479, Dm_build_931.Dm_build_1109(dm_build_1480))
}

func (dm_build_1482 *Dm_build_1290) Dm_build_1481(dm_build_1483 int, dm_build_1484 int16) int {
	return dm_build_1482.Dm_build_1513(dm_build_1483, Dm_build_931.Dm_build_1115(dm_build_1484))
}

func (dm_build_1486 *Dm_build_1290) Dm_build_1485(dm_build_1487 int, dm_build_1488 int32) int {
	return dm_build_1486.Dm_build_1513(dm_build_1487, Dm_build_931.Dm_build_1118(dm_build_1488))
}

func (dm_build_1490 *Dm_build_1290) Dm_build_1489(dm_build_1491 int, dm_build_1492 int64) int {
	return dm_build_1490.Dm_build_1513(dm_build_1491, Dm_build_931.Dm_build_1121(dm_build_1492))
}

func (dm_build_1494 *Dm_build_1290) Dm_build_1493(dm_build_1495 int, dm_build_1496 float32) int {
	return dm_build_1494.Dm_build_1513(dm_build_1495, Dm_build_931.Dm_build_1124(dm_build_1496))
}

func (dm_build_1498 *Dm_build_1290) Dm_build_1497(dm_build_1499 int, dm_build_1500 float64) int {
	return dm_build_1498.Dm_build_1513(dm_build_1499, Dm_build_931.Dm_build_1127(dm_build_1500))
}

func (dm_build_1502 *Dm_build_1290) Dm_build_1501(dm_build_1503 int, dm_build_1504 uint8) int {
	return dm_build_1502.Dm_build_1513(dm_build_1503, Dm_build_931.Dm_build_1130(dm_build_1504))
}

func (dm_build_1506 *Dm_build_1290) Dm_build_1505(dm_build_1507 int, dm_build_1508 uint16) int {
	return dm_build_1506.Dm_build_1513(dm_build_1507, Dm_build_931.Dm_build_1133(dm_build_1508))
}

func (dm_build_1510 *Dm_build_1290) Dm_build_1509(dm_build_1511 int, dm_build_1512 uint32) int {
	return dm_build_1510.Dm_build_1513(dm_build_1511, Dm_build_931.Dm_build_1136(dm_build_1512))
}

func (dm_build_1514 *Dm_build_1290) Dm_build_1513(dm_build_1515 int, dm_build_1516 []byte) int {
	return copy(dm_build_1514.dm_build_1291[dm_build_1515:], dm_build_1516)
}

func (dm_build_1518 *Dm_build_1290) Dm_build_1517(dm_build_1519 int, dm_build_1520 []byte) int {
	return dm_build_1518.Dm_build_1485(dm_build_1519, int32(len(dm_build_1520))) + dm_build_1518.Dm_build_1513(dm_build_1519+4, dm_build_1520)
}

func (dm_build_1522 *Dm_build_1290) Dm_build_1521(dm_build_1523 int, dm_build_1524 []byte) int {
	return dm_build_1522.Dm_build_1477(dm_build_1523, byte(len(dm_build_1524))) + dm_build_1522.Dm_build_1513(dm_build_1523+1, dm_build_1524)
}

func (dm_build_1526 *Dm_build_1290) Dm_build_1525(dm_build_1527 int, dm_build_1528 []byte) int {
	return dm_build_1526.Dm_build_1481(dm_build_1527, int16(len(dm_build_1528))) + dm_build_1526.Dm_build_1513(dm_build_1527+2, dm_build_1528)
}

func (dm_build_1530 *Dm_build_1290) Dm_build_1529(dm_build_1531 int, dm_build_1532 []byte) int {
	return dm_build_1530.Dm_build_1513(dm_build_1531, dm_build_1532) + dm_build_1530.Dm_build_1477(dm_build_1531+len(dm_build_1532), 0)
}

func (dm_build_1534 *Dm_build_1290) Dm_build_1533(dm_build_1535 int, dm_build_1536 string, dm_build_1537 string, dm_build_1538 *DmConnection) int {
	return dm_build_1534.Dm_build_1517(dm_build_1535, Dm_build_931.Dm_build_1147(dm_build_1536, dm_build_1537, dm_build_1538))
}

func (dm_build_1540 *Dm_build_1290) Dm_build_1539(dm_build_1541 int, dm_build_1542 string, dm_build_1543 string, dm_build_1544 *DmConnection) int {
	return dm_build_1540.Dm_build_1521(dm_build_1541, Dm_build_931.Dm_build_1147(dm_build_1542, dm_build_1543, dm_build_1544))
}

func (dm_build_1546 *Dm_build_1290) Dm_build_1545(dm_build_1547 int, dm_build_1548 string, dm_build_1549 string, dm_build_1550 *DmConnection) int {
	return dm_build_1546.Dm_build_1525(dm_build_1547, Dm_build_931.Dm_build_1147(dm_build_1548, dm_build_1549, dm_build_1550))
}

func (dm_build_1552 *Dm_build_1290) Dm_build_1551(dm_build_1553 int, dm_build_1554 string, dm_build_1555 string, dm_build_1556 *DmConnection) int {
	return dm_build_1552.Dm_build_1529(dm_build_1553, Dm_build_931.Dm_build_1147(dm_build_1554, dm_build_1555, dm_build_1556))
}

func (dm_build_1558 *Dm_build_1290) Dm_build_1557(dm_build_1559 int) byte {
	return Dm_build_931.Dm_build_1152(dm_build_1558.Dm_build_1584(dm_build_1559, 1))
}

func (dm_build_1561 *Dm_build_1290) Dm_build_1560(dm_build_1562 int) int16 {
	return Dm_build_931.Dm_build_1155(dm_build_1561.Dm_build_1584(dm_build_1562, 2))
}

func (dm_build_1564 *Dm_build_1290) Dm_build_1563(dm_build_1565 int) int32 {
	return Dm_build_931.Dm_build_1158(dm_build_1564.Dm_build_1584(dm_build_1565, 4))
}

func (dm_build_1567 *Dm_build_1290) Dm_build_1566(dm_build_1568 int) int64 {
	return Dm_build_931.Dm_build_1161(dm_build_1567.Dm_build_1584(dm_build_1568, 8))
}

func (dm_build_1570 *Dm_build_1290) Dm_build_1569(dm_build_1571 int) float32 {
	return Dm_build_931.Dm_build_1164(dm_build_1570.Dm_build_1584(dm_build_1571, 4))
}

func (dm_build_1573 *Dm_build_1290) Dm_build_1572(dm_build_1574 int) float64 {
	return Dm_build_931.Dm_build_1167(dm_build_1573.Dm_build_1584(dm_build_1574, 8))
}

func (dm_build_1576 *Dm_build_1290) Dm_build_1575(dm_build_1577 int) uint8 {
	return Dm_build_931.Dm_build_1170(dm_build_1576.Dm_build_1584(dm_build_1577, 1))
}

func (dm_build_1579 *Dm_build_1290) Dm_build_1578(dm_build_1580 int) uint16 {
	return Dm_build_931.Dm_build_1173(dm_build_1579.Dm_build_1584(dm_build_1580, 2))
}

func (dm_build_1582 *Dm_build_1290) Dm_build_1581(dm_build_1583 int) uint32 {
	return Dm_build_931.Dm_build_1176(dm_build_1582.Dm_build_1584(dm_build_1583, 4))
}

func (dm_build_1585 *Dm_build_1290) Dm_build_1584(dm_build_1586 int, dm_build_1587 int) []byte {
	return dm_build_1585.dm_build_1291[dm_build_1586 : dm_build_1586+dm_build_1587]
}

func (dm_build_1589 *Dm_build_1290) Dm_build_1588(dm_build_1590 int) []byte {
	dm_build_1591 := dm_build_1589.Dm_build_1563(dm_build_1590)
	return dm_build_1589.Dm_build_1584(dm_build_1590+4, int(dm_build_1591))
}

func (dm_build_1593 *Dm_build_1290) Dm_build_1592(dm_build_1594 int) []byte {
	dm_build_1595 := dm_build_1593.Dm_build_1557(dm_build_1594)
	return dm_build_1593.Dm_build_1584(dm_build_1594+1, int(dm_build_1595))
}

func (dm_build_1597 *Dm_build_1290) Dm_build_1596(dm_build_1598 int) []byte {
	dm_build_1599 := dm_build_1597.Dm_build_1560(dm_build_1598)
	return dm_build_1597.Dm_build_1584(dm_build_1598+2, int(dm_build_1599))
}

func (dm_build_1601 *Dm_build_1290) Dm_build_1600(dm_build_1602 int) []byte {
	dm_build_1603 := 0
	for dm_build_1601.Dm_build_1557(dm_build_1602) != 0 {
		dm_build_1602++
		dm_build_1603++
	}

	return dm_build_1601.Dm_build_1584(dm_build_1602-dm_build_1603, dm_build_1603)
}

func (dm_build_1605 *Dm_build_1290) Dm_build_1604(dm_build_1606 int, dm_build_1607 string, dm_build_1608 *DmConnection) string {
	return Dm_build_931.Dm_build_1183(dm_build_1605.Dm_build_1588(dm_build_1606), dm_build_1607, dm_build_1608)
}

func (dm_build_1610 *Dm_build_1290) Dm_build_1609(dm_build_1611 int, dm_build_1612 string, dm_build_1613 *DmConnection) string {
	return Dm_build_931.Dm_build_1183(dm_build_1610.Dm_build_1592(dm_build_1611), dm_build_1612, dm_build_1613)
}

func (dm_build_1615 *Dm_build_1290) Dm_build_1614(dm_build_1616 int, dm_build_1617 string, dm_build_1618 *DmConnection) string {
	return Dm_build_931.Dm_build_1183(dm_build_1615.Dm_build_1596(dm_build_1616), dm_build_1617, dm_build_1618)
}

func (dm_build_1620 *Dm_build_1290) Dm_build_1619(dm_build_1621 int, dm_build_1622 string, dm_build_1623 *DmConnection) string {
	return Dm_build_931.Dm_build_1183(dm_build_1620.Dm_build_1600(dm_build_1621), dm_build_1622, dm_build_1623)
}
