/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */

package dm8

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/godoes/gorm-dameng/dm8/security"
)

const (
	Dm_build_694 = 8192
	Dm_build_695 = 2 * time.Second
)

type dm_build_696 struct {
	dm_build_697 net.Conn
	dm_build_698 *tls.Conn
	dm_build_699 *Dm_build_360
	dm_build_700 *DmConnection
	dm_build_701 security.Cipher
	dm_build_702 bool
	dm_build_703 bool
	dm_build_704 *security.DhKey

	dm_build_705 bool
	dm_build_706 string
	dm_build_707 bool
}

func dm_build_708(dm_build_709 context.Context, dm_build_710 *DmConnection) (*dm_build_696, error) {
	var dm_build_711 net.Conn
	var dm_build_712 error

	dialsLock.RLock()
	dm_build_713, dm_build_714 := dials[dm_build_710.dmConnector.dialName]
	dialsLock.RUnlock()
	if dm_build_714 {
		dm_build_711, dm_build_712 = dm_build_713(dm_build_709, dm_build_710.dmConnector.host+":"+strconv.Itoa(int(dm_build_710.dmConnector.port)))
	} else {
		dm_build_711, dm_build_712 = dm_build_716(dm_build_710.dmConnector.host+":"+strconv.Itoa(int(dm_build_710.dmConnector.port)), time.Duration(dm_build_710.dmConnector.socketTimeout)*time.Second)
	}
	if dm_build_712 != nil {
		return nil, dm_build_712
	}

	dm_build_715 := dm_build_696{}
	dm_build_715.dm_build_697 = dm_build_711
	dm_build_715.dm_build_699 = Dm_build_363(Dm_build_981)
	dm_build_715.dm_build_700 = dm_build_710
	dm_build_715.dm_build_702 = false
	dm_build_715.dm_build_703 = false
	dm_build_715.dm_build_705 = false
	dm_build_715.dm_build_706 = ""
	dm_build_715.dm_build_707 = false
	dm_build_710.Access = &dm_build_715

	return &dm_build_715, nil
}

func dm_build_716(dm_build_717 string, dm_build_718 time.Duration) (net.Conn, error) {
	dm_build_719, dm_build_720 := net.DialTimeout("tcp", dm_build_717, dm_build_718)
	if dm_build_720 != nil {
		return &net.TCPConn{}, ECGO_COMMUNITION_ERROR.addDetail("\tdial address: " + dm_build_717).throw()
	}

	if tcpConn, ok := dm_build_719.(*net.TCPConn); ok {

		_ = tcpConn.SetKeepAlive(true)
		_ = tcpConn.SetKeepAlivePeriod(Dm_build_695)
		_ = tcpConn.SetNoDelay(true)

	}
	return dm_build_719, nil
}

func (dm_build_722 *dm_build_696) dm_build_721(dm_build_723 dm_build_1102) bool {
	var dm_build_724 = dm_build_722.dm_build_700.dmConnector.compress
	if dm_build_723.dm_build_1117() == Dm_build_1009 || dm_build_724 == Dm_build_1058 {
		return false
	}

	if dm_build_724 == Dm_build_1056 {
		return true
	} else if dm_build_724 == Dm_build_1057 {
		return !dm_build_722.dm_build_700.Local && dm_build_723.dm_build_1115() > Dm_build_1055
	}

	return false
}

func (dm_build_726 *dm_build_696) dm_build_725(dm_build_727 dm_build_1102) bool {
	var dm_build_728 = dm_build_726.dm_build_700.dmConnector.compress
	if dm_build_727.dm_build_1117() == Dm_build_1009 || dm_build_728 == Dm_build_1058 {
		return false
	}

	if dm_build_728 == Dm_build_1056 {
		return true
	} else if dm_build_728 == Dm_build_1057 {
		return dm_build_726.dm_build_699.Dm_build_627(Dm_build_1017) == 1
	}

	return false
}

func (dm_build_730 *dm_build_696) dm_build_729(dm_build_731 dm_build_1102) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_733 := dm_build_731.dm_build_1115()

	if dm_build_733 > 0 {

		if dm_build_730.dm_build_721(dm_build_731) {
			var retBytes, err = Compress(dm_build_730.dm_build_699, Dm_build_1010, int(dm_build_733), int(dm_build_730.dm_build_700.dmConnector.compressID))
			if err != nil {
				return err
			}

			dm_build_730.dm_build_699.Dm_build_374(Dm_build_1010)

			dm_build_730.dm_build_699.Dm_build_415(dm_build_733)

			dm_build_730.dm_build_699.Dm_build_443(retBytes)

			dm_build_731.dm_build_1116(int32(len(retBytes)) + ULINT_SIZE)

			dm_build_730.dm_build_699.Dm_build_547(Dm_build_1017, 1)
		}

		if dm_build_730.dm_build_703 {
			dm_build_733 = dm_build_731.dm_build_1115()
			var retBytes = dm_build_730.dm_build_701.Encrypt(dm_build_730.dm_build_699.Dm_build_654(Dm_build_1010, int(dm_build_733)), true)

			dm_build_730.dm_build_699.Dm_build_374(Dm_build_1010)

			dm_build_730.dm_build_699.Dm_build_443(retBytes)

			dm_build_731.dm_build_1116(int32(len(retBytes)))
		}
	}

	if dm_build_730.dm_build_699.Dm_build_372() > Dm_build_982 {
		return ECGO_MSG_TOO_LONG.throw()
	}

	dm_build_731.dm_build_1111()
	if dm_build_730.dm_build_964(dm_build_731) {
		if dm_build_730.dm_build_698 != nil {
			dm_build_730.dm_build_699.Dm_build_377(0)
			if _, err := dm_build_730.dm_build_699.Dm_build_396(dm_build_730.dm_build_698); err != nil {
				return err
			}
		}
	} else {
		dm_build_730.dm_build_699.Dm_build_377(0)
		if _, err := dm_build_730.dm_build_699.Dm_build_396(dm_build_730.dm_build_697); err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_735 *dm_build_696) dm_build_734(dm_build_736 dm_build_1102) (err error) {
	defer func() {
		if p := recover(); p != nil {
			if _, ok := p.(string); ok {
				err = ECGO_COMMUNITION_ERROR.addDetail("\t" + p.(string)).throw()
			} else {
				err = fmt.Errorf("internal error: %v", p)
			}
		}
	}()

	dm_build_738 := int32(0)
	if dm_build_735.dm_build_964(dm_build_736) {
		if dm_build_735.dm_build_698 != nil {
			dm_build_735.dm_build_699.Dm_build_374(0)
			if _, err := dm_build_735.dm_build_699.Dm_build_390(dm_build_735.dm_build_698, Dm_build_1010); err != nil {
				return err
			}

			dm_build_738 = dm_build_736.dm_build_1115()
			if dm_build_738 > 0 {
				if _, err := dm_build_735.dm_build_699.Dm_build_390(dm_build_735.dm_build_698, int(dm_build_738)); err != nil {
					return err
				}
			}
		}
	} else {

		dm_build_735.dm_build_699.Dm_build_374(0)
		if _, err := dm_build_735.dm_build_699.Dm_build_390(dm_build_735.dm_build_697, Dm_build_1010); err != nil {
			return err
		}
		dm_build_738 = dm_build_736.dm_build_1115()

		if dm_build_738 > 0 {
			if _, err := dm_build_735.dm_build_699.Dm_build_390(dm_build_735.dm_build_697, int(dm_build_738)); err != nil {
				return err
			}
		}
	}

	_ = dm_build_736.dm_build_1112()

	dm_build_738 = dm_build_736.dm_build_1115()
	if dm_build_738 <= 0 {
		return nil
	}

	if dm_build_735.dm_build_703 {
		eBytes := dm_build_735.dm_build_699.Dm_build_654(Dm_build_1010, int(dm_build_738))
		dBytes, err := dm_build_735.dm_build_701.Decrypt(eBytes, true)
		if err != nil {
			return err
		}
		dm_build_735.dm_build_699.Dm_build_374(Dm_build_1010)
		dm_build_735.dm_build_699.Dm_build_443(dBytes)
		dm_build_736.dm_build_1116(int32(len(dBytes)))
	}

	if dm_build_735.dm_build_725(dm_build_736) {

		dm_build_738 = dm_build_736.dm_build_1115()
		cBytes := dm_build_735.dm_build_699.Dm_build_654(Dm_build_1010+ULINT_SIZE, int(dm_build_738-ULINT_SIZE))
		uBytes, err := UnCompress(cBytes, int(dm_build_735.dm_build_700.dmConnector.compressID))
		if err != nil {
			return err
		}
		dm_build_735.dm_build_699.Dm_build_374(Dm_build_1010)
		dm_build_735.dm_build_699.Dm_build_443(uBytes)
		dm_build_736.dm_build_1116(int32(len(uBytes)))
	}
	return nil
}

func (dm_build_740 *dm_build_696) dm_build_739(dm_build_741 dm_build_1102) (dm_build_742 interface{}, dm_build_743 error) {
	if dm_build_740.dm_build_707 {
		return nil, ECGO_CONNECTION_CLOSED.throw()
	}
	dm_build_744 := dm_build_740.dm_build_700
	dm_build_744.mu.Lock()
	defer dm_build_744.mu.Unlock()
	dm_build_743 = dm_build_741.dm_build_1106(dm_build_741)
	if dm_build_743 != nil {
		return nil, dm_build_743
	}

	dm_build_743 = dm_build_740.dm_build_729(dm_build_741)
	if dm_build_743 != nil {
		return nil, dm_build_743
	}

	dm_build_743 = dm_build_740.dm_build_734(dm_build_741)
	if dm_build_743 != nil {
		return nil, dm_build_743
	}

	return dm_build_741.dm_build_1110(dm_build_741)
}

func (dm_build_746 *dm_build_696) dm_build_745() (*dm_build_1559, error) {

	Dm_build_747 := dm_build_1565(dm_build_746)
	_, dm_build_748 := dm_build_746.dm_build_739(Dm_build_747)
	if dm_build_748 != nil {
		return nil, dm_build_748
	}

	return Dm_build_747, nil
}

func (dm_build_750 *dm_build_696) dm_build_749() error {

	dm_build_751 := dm_build_1426(dm_build_750)
	_, dm_build_752 := dm_build_750.dm_build_739(dm_build_751)
	if dm_build_752 != nil {
		return dm_build_752
	}

	return nil
}

func (dm_build_754 *dm_build_696) dm_build_753() error {

	var dm_build_755 *dm_build_1559
	var err error
	if dm_build_755, err = dm_build_754.dm_build_745(); err != nil {
		return err
	}

	if dm_build_754.dm_build_700.sslEncrypt == 2 {
		if err = dm_build_754.dm_build_960(false); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	} else if dm_build_754.dm_build_700.sslEncrypt == 1 {
		if err = dm_build_754.dm_build_960(true); err != nil {
			return ECGO_INIT_SSL_FAILED.addDetail("\n" + err.Error()).throw()
		}
	}

	if dm_build_754.dm_build_703 || dm_build_754.dm_build_702 {
		k, err := dm_build_754.dm_build_950()
		if err != nil {
			return err
		}
		sessionKey := security.ComputeSessionKey(k, dm_build_755.Dm_build_1563)
		encryptType := dm_build_755.dm_build_1561
		hashType := int(dm_build_755.Dm_build_1562)
		if encryptType == -1 {
			encryptType = security.DES_CFB
		}
		if hashType == -1 {
			hashType = security.MD5
		}
		err = dm_build_754.dm_build_953(encryptType, sessionKey, dm_build_754.dm_build_700.dmConnector.cipherPath, hashType)
		if err != nil {
			return err
		}
	}

	if err := dm_build_754.dm_build_749(); err != nil {
		return err
	}
	return nil
}

func (dm_build_758 *dm_build_696) Dm_build_757(dm_build_759 *DmStatement) error {
	dm_build_760 := dm_build_1588(dm_build_758, dm_build_759)
	_, dm_build_761 := dm_build_758.dm_build_739(dm_build_760)
	if dm_build_761 != nil {
		return dm_build_761
	}

	return nil
}

func (dm_build_763 *dm_build_696) Dm_build_762(dm_build_764 int32) error {
	dm_build_765 := dm_build_1598(dm_build_763, dm_build_764)
	_, dm_build_766 := dm_build_763.dm_build_739(dm_build_765)
	if dm_build_766 != nil {
		return dm_build_766
	}

	return nil
}

func (dm_build_768 *dm_build_696) Dm_build_767(dm_build_769 *DmStatement, dm_build_770 bool, dm_build_771 int16) (*execRetInfo, error) {
	dm_build_772 := dm_build_1465(dm_build_768, dm_build_769, dm_build_770, dm_build_771)
	dm_build_773, dm_build_774 := dm_build_768.dm_build_739(dm_build_772)
	if dm_build_774 != nil {
		return nil, dm_build_774
	}
	return dm_build_773.(*execRetInfo), nil
}

func (dm_build_776 *dm_build_696) Dm_build_775(dm_build_777 *DmStatement, _ int16) (*execRetInfo, error) {
	return dm_build_776.Dm_build_767(dm_build_777, false, Dm_build_1062)
}

func (dm_build_780 *dm_build_696) Dm_build_779(dm_build_781 *DmStatement, dm_build_782 []OptParameter) (*execRetInfo, error) {
	dm_build_783, dm_build_784 := dm_build_780.dm_build_739(dm_build_1205(dm_build_780, dm_build_781, dm_build_782))
	if dm_build_784 != nil {
		return nil, dm_build_784
	}

	return dm_build_783.(*execRetInfo), nil
}

func (dm_build_786 *dm_build_696) Dm_build_785(dm_build_787 *DmStatement, dm_build_788 int16) (*execRetInfo, error) {
	return dm_build_786.Dm_build_767(dm_build_787, true, dm_build_788)
}

func (dm_build_790 *dm_build_696) Dm_build_789(dm_build_791 *DmStatement, dm_build_792 [][]interface{}) (*execRetInfo, error) {
	dm_build_793 := dm_build_1237(dm_build_790, dm_build_791, dm_build_792)
	dm_build_794, dm_build_795 := dm_build_790.dm_build_739(dm_build_793)
	if dm_build_795 != nil {
		return nil, dm_build_795
	}
	return dm_build_794.(*execRetInfo), nil
}

func (dm_build_797 *dm_build_696) Dm_build_796(dm_build_798 *DmStatement, dm_build_799 [][]interface{}, dm_build_800 bool) (*execRetInfo, error) {
	var dm_build_801, dm_build_802 = 0, 0
	var dm_build_803 = len(dm_build_799)
	var dm_build_804 [][]interface{}
	var dm_build_805 = NewExceInfo()
	dm_build_805.updateCounts = make([]int64, dm_build_803)
	var dm_build_806 = false
	for dm_build_801 < dm_build_803 {
		for dm_build_802 = dm_build_801; dm_build_802 < dm_build_803; dm_build_802++ {
			paramData := dm_build_799[dm_build_802]
			bindData := make([]interface{}, dm_build_798.paramCount)
			dm_build_806 = false
			for icol := 0; icol < int(dm_build_798.paramCount); icol++ {
				if dm_build_798.bindParams[icol].ioType == IO_TYPE_OUT {
					continue
				}
				if dm_build_797.dm_build_933(bindData, paramData, icol) {
					dm_build_806 = true
					break
				}
			}

			if dm_build_806 {
				break
			}
			dm_build_804 = append(dm_build_804, bindData)
		}

		if dm_build_802 != dm_build_801 {
			tmpExecInfo, err := dm_build_797.Dm_build_789(dm_build_798, dm_build_804)
			if err != nil {
				return nil, err
			}
			dm_build_804 = dm_build_804[0:0]
			dm_build_805.union(tmpExecInfo, dm_build_801, dm_build_802-dm_build_801)
		}

		if dm_build_802 < dm_build_803 {
			tmpExecInfo, err := dm_build_797.Dm_build_807(dm_build_798, dm_build_799[dm_build_802], dm_build_800)
			if err != nil {
				return nil, err
			}

			dm_build_800 = true
			dm_build_805.union(tmpExecInfo, dm_build_802, 1)
		}

		dm_build_801 = dm_build_802 + 1
	}
	for _, i := range dm_build_805.updateCounts {
		if i > 0 {
			dm_build_805.updateCount += i
		}
	}
	return dm_build_805, nil
}

func (dm_build_808 *dm_build_696) Dm_build_807(dm_build_809 *DmStatement, dm_build_810 []interface{}, dm_build_811 bool) (*execRetInfo, error) {

	var dm_build_812 = make([]interface{}, dm_build_809.paramCount)
	for icol := 0; icol < int(dm_build_809.paramCount); icol++ {
		if dm_build_809.bindParams[icol].ioType == IO_TYPE_OUT {
			continue
		}
		if dm_build_808.dm_build_933(dm_build_812, dm_build_810, icol) {

			if !dm_build_811 {
				preExecute := dm_build_1454(dm_build_808, dm_build_809, dm_build_809.bindParams)
				_, _ = dm_build_808.dm_build_739(preExecute)
				dm_build_811 = true
			}

			_ = dm_build_808.dm_build_939(dm_build_809, dm_build_809.bindParams[icol], icol, dm_build_810[icol].(iOffRowBinder))
			dm_build_812[icol] = ParamDataEnum_OFF_ROW
		}
	}

	var dm_build_813 = make([][]interface{}, 1)
	dm_build_813[0] = dm_build_812

	dm_build_814 := dm_build_1237(dm_build_808, dm_build_809, dm_build_813)
	dm_build_815, dm_build_816 := dm_build_808.dm_build_739(dm_build_814)
	if dm_build_816 != nil {
		return nil, dm_build_816
	}
	return dm_build_815.(*execRetInfo), nil
}

func (dm_build_818 *dm_build_696) Dm_build_817(dm_build_819 *DmStatement, dm_build_820 int16) (*execRetInfo, error) {
	dm_build_821 := dm_build_1441(dm_build_818, dm_build_819, dm_build_820)

	dm_build_822, dm_build_823 := dm_build_818.dm_build_739(dm_build_821)
	if dm_build_823 != nil {
		return nil, dm_build_823
	}
	return dm_build_822.(*execRetInfo), nil
}

func (dm_build_825 *dm_build_696) Dm_build_824(dm_build_826 *innerRows, dm_build_827 int64) (*execRetInfo, error) {
	dm_build_828 := dm_build_1344(dm_build_825, dm_build_826, dm_build_827, INT64_MAX)
	dm_build_829, dm_build_830 := dm_build_825.dm_build_739(dm_build_828)
	if dm_build_830 != nil {
		return nil, dm_build_830
	}
	return dm_build_829.(*execRetInfo), nil
}

func (dm_build_832 *dm_build_696) Commit() error {
	dm_build_833 := dm_build_1190(dm_build_832)
	_, dm_build_834 := dm_build_832.dm_build_739(dm_build_833)
	if dm_build_834 != nil {
		return dm_build_834
	}

	return nil
}

func (dm_build_836 *dm_build_696) Rollback() error {
	dm_build_837 := dm_build_1503(dm_build_836)
	_, dm_build_838 := dm_build_836.dm_build_739(dm_build_837)
	if dm_build_838 != nil {
		return dm_build_838
	}

	return nil
}

func (dm_build_840 *dm_build_696) Dm_build_839(dm_build_841 *DmConnection) error {
	dm_build_842 := dm_build_1508(dm_build_840, dm_build_841.IsoLevel)
	_, dm_build_843 := dm_build_840.dm_build_739(dm_build_842)
	if dm_build_843 != nil {
		return dm_build_843
	}

	return nil
}

func (dm_build_845 *dm_build_696) Dm_build_844(dm_build_846 *DmStatement, dm_build_847 string) error {
	dm_build_848 := dm_build_1195(dm_build_845, dm_build_846, dm_build_847)
	_, dm_build_849 := dm_build_845.dm_build_739(dm_build_848)
	if dm_build_849 != nil {
		return dm_build_849
	}

	return nil
}

func (dm_build_851 *dm_build_696) Dm_build_850(dm_build_852 []uint32) ([]int64, error) {
	dm_build_853 := dm_build_1606(dm_build_851, dm_build_852)
	dm_build_854, dm_build_855 := dm_build_851.dm_build_739(dm_build_853)
	if dm_build_855 != nil {
		return nil, dm_build_855
	}
	return dm_build_854.([]int64), nil
}

func (dm_build_857 *dm_build_696) Close() error {
	if dm_build_857.dm_build_707 {
		return nil
	}

	dm_build_858 := dm_build_857.dm_build_697.Close()
	if dm_build_858 != nil {
		return dm_build_858
	}

	dm_build_857.dm_build_700 = nil
	dm_build_857.dm_build_707 = true
	return nil
}

func (dm_build_860 *dm_build_696) dm_build_859(dm_build_861 *lob) (int64, error) {
	dm_build_862 := dm_build_1377(dm_build_860, dm_build_861)
	dm_build_863, dm_build_864 := dm_build_860.dm_build_739(dm_build_862)
	if dm_build_864 != nil {
		return 0, dm_build_864
	}
	return dm_build_863.(int64), nil
}

func (dm_build_866 *dm_build_696) dm_build_865(dm_build_867 *lob, dm_build_868 int32, dm_build_869 int32) (*lobRetInfo, error) {
	dm_build_870 := dm_build_1362(dm_build_866, dm_build_867, int(dm_build_868), int(dm_build_869))
	dm_build_871, dm_build_872 := dm_build_866.dm_build_739(dm_build_870)
	if dm_build_872 != nil {
		return nil, dm_build_872
	}
	return dm_build_871.(*lobRetInfo), nil
}

func (dm_build_874 *dm_build_696) dm_build_873(dm_build_875 *DmBlob, dm_build_876 int32, dm_build_877 int32) ([]byte, error) {
	var dm_build_878 = make([]byte, dm_build_877)
	var dm_build_879 int32 = 0
	var dm_build_880 int32 = 0
	var dm_build_881 *lobRetInfo
	var dm_build_882 []byte
	var dm_build_883 error
	for dm_build_879 < dm_build_877 {
		dm_build_880 = dm_build_877 - dm_build_879
		if dm_build_880 > Dm_build_1095 {
			dm_build_880 = Dm_build_1095
		}
		dm_build_881, dm_build_883 = dm_build_874.dm_build_865(&dm_build_875.lob, dm_build_876+dm_build_879, dm_build_880)
		if dm_build_883 != nil {
			return nil, dm_build_883
		}
		dm_build_882 = dm_build_881.data
		if dm_build_882 == nil || len(dm_build_882) == 0 {
			break
		}
		Dm_build_1.Dm_build_57(dm_build_878, int(dm_build_879), dm_build_882, 0, len(dm_build_882))
		dm_build_879 += int32(len(dm_build_882))
		if dm_build_875.readOver {
			break
		}
	}
	return dm_build_878, nil
}

func (dm_build_885 *dm_build_696) dm_build_884(dm_build_886 *DmClob, dm_build_887 int32, dm_build_888 int32) (string, error) {
	var dm_build_889 bytes.Buffer
	var dm_build_890 int32 = 0
	var dm_build_891 int32 = 0
	var dm_build_892 *lobRetInfo
	var dm_build_893 []byte
	var dm_build_894 string
	var dm_build_895 error
	for dm_build_890 < dm_build_888 {
		dm_build_891 = dm_build_888 - dm_build_890
		if dm_build_891 > Dm_build_1095/2 {
			dm_build_891 = Dm_build_1095 / 2
		}
		dm_build_892, dm_build_895 = dm_build_885.dm_build_865(&dm_build_886.lob, dm_build_887+dm_build_890, dm_build_891)
		if dm_build_895 != nil {
			return "", dm_build_895
		}
		dm_build_893 = dm_build_892.data
		if dm_build_893 == nil || len(dm_build_893) == 0 {
			break
		}
		dm_build_894 = Dm_build_1.Dm_build_158(dm_build_893, 0, len(dm_build_893), dm_build_886.serverEncoding, dm_build_885.dm_build_700)

		dm_build_889.WriteString(dm_build_894)
		var strLen = dm_build_892.charLen
		if strLen == -1 {
			strLen = int64(utf8.RuneCountInString(dm_build_894))
		}
		dm_build_890 += int32(strLen)
		if dm_build_886.readOver {
			break
		}
	}
	return dm_build_889.String(), nil
}

func (dm_build_897 *dm_build_696) dm_build_896(dm_build_898 *DmClob, dm_build_899 int, dm_build_900 string, dm_build_901 string) (int, error) {
	var dm_build_902 = Dm_build_1.Dm_build_217(dm_build_900, dm_build_901, dm_build_897.dm_build_700)
	var dm_build_903 = 0
	var dm_build_904 = len(dm_build_902)
	var dm_build_905 = 0
	var dm_build_906 = 0
	var dm_build_907 = 0
	var dm_build_908 = dm_build_904/Dm_build_1094 + 1
	var dm_build_909 byte = 0
	var dm_build_910 byte = 0x01
	var dm_build_911 byte = 0x02
	for i := 0; i < dm_build_908; i++ {
		dm_build_909 = 0
		if i == 0 {
			dm_build_909 |= dm_build_910
		}
		if i == dm_build_908-1 {
			dm_build_909 |= dm_build_911
		}
		dm_build_907 = dm_build_904 - dm_build_906
		if dm_build_907 > Dm_build_1094 {
			dm_build_907 = Dm_build_1094
		}

		setLobData := dm_build_1522(dm_build_897, &dm_build_898.lob, dm_build_909, dm_build_899, dm_build_902, dm_build_903, dm_build_907)
		ret, err := dm_build_897.dm_build_739(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		//if err != nil {
		//	return -1, err
		//}
		if tmp <= 0 {
			return dm_build_905, nil
		} else {
			dm_build_899 += int(tmp)
			dm_build_905 += int(tmp)
			dm_build_906 += dm_build_907
			dm_build_903 += dm_build_907
		}
	}
	return dm_build_905, nil
}

func (dm_build_913 *dm_build_696) dm_build_912(dm_build_914 *DmBlob, dm_build_915 int, dm_build_916 []byte) (int, error) {
	var dm_build_917 = 0
	var dm_build_918 = len(dm_build_916)
	var dm_build_919 = 0
	var dm_build_920 = 0
	var dm_build_921 = 0
	var dm_build_922 = dm_build_918/Dm_build_1094 + 1
	var dm_build_923 byte = 0
	var dm_build_924 byte = 0x01
	var dm_build_925 byte = 0x02
	for i := 0; i < dm_build_922; i++ {
		dm_build_923 = 0
		if i == 0 {
			dm_build_923 |= dm_build_924
		}
		if i == dm_build_922-1 {
			dm_build_923 |= dm_build_925
		}
		dm_build_921 = dm_build_918 - dm_build_920
		if dm_build_921 > Dm_build_1094 {
			dm_build_921 = Dm_build_1094
		}

		setLobData := dm_build_1522(dm_build_913, &dm_build_914.lob, dm_build_923, dm_build_915, dm_build_916, dm_build_917, dm_build_921)
		ret, err := dm_build_913.dm_build_739(setLobData)
		if err != nil {
			return 0, err
		}
		tmp := ret.(int32)
		if tmp <= 0 {
			return dm_build_919, nil
		} else {
			dm_build_915 += int(tmp)
			dm_build_919 += int(tmp)
			dm_build_920 += dm_build_921
			dm_build_917 += dm_build_921
		}
	}
	return dm_build_919, nil
}

func (dm_build_927 *dm_build_696) dm_build_926(dm_build_928 *lob, dm_build_929 int) (int64, error) {
	dm_build_930 := dm_build_1388(dm_build_927, dm_build_928, dm_build_929)
	dm_build_931, dm_build_932 := dm_build_927.dm_build_739(dm_build_930)
	if dm_build_932 != nil {
		return dm_build_928.length, dm_build_932
	}
	return dm_build_931.(int64), nil
}

func (dm_build_934 *dm_build_696) dm_build_933(dm_build_935 []interface{}, dm_build_936 []interface{}, dm_build_937 int) bool {
	var dm_build_938 = false
	dm_build_935[dm_build_937] = dm_build_936[dm_build_937]

	if binder, ok := dm_build_936[dm_build_937].(iOffRowBinder); ok {
		dm_build_938 = true
		dm_build_935[dm_build_937] = make([]byte, 0)
		var lob lob
		if l, ok := binder.getObj().(DmBlob); ok {
			lob = l.lob
		} else if l, ok := binder.getObj().(DmClob); ok {
			lob = l.lob
		}
		if &lob != nil && lob.canOptimized(dm_build_934.dm_build_700) {
			dm_build_935[dm_build_937] = &lobCtl{lob.buildCtlData()}
			dm_build_938 = false
		}
	} else {
		dm_build_935[dm_build_937] = dm_build_936[dm_build_937]
	}
	return dm_build_938
}

func (dm_build_940 *dm_build_696) dm_build_939(dm_build_941 *DmStatement, _ parameter, dm_build_943 int, dm_build_944 iOffRowBinder) error {
	var dm_build_945 = Dm_build_286()
	dm_build_944.read(dm_build_945)
	var dm_build_946 = 0
	for !dm_build_944.isReadOver() || dm_build_945.Dm_build_287() > 0 {
		if !dm_build_944.isReadOver() && dm_build_945.Dm_build_287() < Dm_build_1094 {
			dm_build_944.read(dm_build_945)
		}
		if dm_build_945.Dm_build_287() > Dm_build_1094 {
			dm_build_946 = Dm_build_1094
		} else {
			dm_build_946 = dm_build_945.Dm_build_287()
		}

		putData := dm_build_1493(dm_build_940, dm_build_941, int16(dm_build_943), dm_build_945, int32(dm_build_946))
		_, err := dm_build_940.dm_build_739(putData)
		if err != nil {
			return err
		}
	}
	return nil
}

func (dm_build_948 *dm_build_696) dm_build_947() ([]byte, error) {
	var dm_build_949 error
	if dm_build_948.dm_build_704 == nil {
		if dm_build_948.dm_build_704, dm_build_949 = security.NewClientKeyPair(); dm_build_949 != nil {
			return nil, dm_build_949
		}
	}
	return security.Bn2Bytes(dm_build_948.dm_build_704.GetY(), security.DH_KEY_LENGTH), nil
}

func (dm_build_951 *dm_build_696) dm_build_950() (*security.DhKey, error) {
	var dm_build_952 error
	if dm_build_951.dm_build_704 == nil {
		if dm_build_951.dm_build_704, dm_build_952 = security.NewClientKeyPair(); dm_build_952 != nil {
			return nil, dm_build_952
		}
	}
	return dm_build_951.dm_build_704, nil
}

func (dm_build_954 *dm_build_696) dm_build_953(dm_build_955 int, dm_build_956 []byte, dm_build_957 string, dm_build_958 int) (dm_build_959 error) {
	if dm_build_955 > 0 && dm_build_955 < security.MIN_EXTERNAL_CIPHER_ID && dm_build_956 != nil {
		dm_build_954.dm_build_701, dm_build_959 = security.NewSymmCipher(dm_build_955, dm_build_956)
	} else if dm_build_955 >= security.MIN_EXTERNAL_CIPHER_ID {
		if dm_build_954.dm_build_701, dm_build_959 = security.NewThirdPartCipher(dm_build_955, dm_build_956, dm_build_957, dm_build_958); dm_build_959 != nil {
			dm_build_959 = THIRD_PART_CIPHER_INIT_FAILED.addDetailln(dm_build_959.Error()).throw()
		}
	}
	return
}

func (dm_build_961 *dm_build_696) dm_build_960(dm_build_962 bool) (dm_build_963 error) {
	if dm_build_961.dm_build_698, dm_build_963 = security.NewTLSFromTCP(dm_build_961.dm_build_697, dm_build_961.dm_build_700.dmConnector.sslCertPath, dm_build_961.dm_build_700.dmConnector.sslKeyPath, dm_build_961.dm_build_700.dmConnector.user); dm_build_963 != nil {
		return
	}
	if !dm_build_962 {
		dm_build_961.dm_build_698 = nil
	}
	return
}

func (dm_build_965 *dm_build_696) dm_build_964(dm_build_966 dm_build_1102) bool {
	return dm_build_966.dm_build_1117() != Dm_build_1009 && dm_build_965.dm_build_700.sslEncrypt == 1
}
