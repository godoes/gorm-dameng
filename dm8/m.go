/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */

package dm8

import (
	"bytes"
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/godoes/gorm-dameng/dm8/parser"
	"github.com/godoes/gorm-dameng/dm8/util"
	"golang.org/x/text/encoding"
)

type DmConnection struct {
	filterable
	mu sync.Mutex

	dmConnector *DmConnector
	Access      *dm_build_2
	stmtMap     map[int32]*DmStatement

	lastExecInfo       *execRetInfo
	lexer              *parser.Lexer
	encode             encoding.Encoding
	encodeBuffer       *bytes.Buffer
	transformReaderDst []byte
	transformReaderSrc []byte

	serverEncoding     string
	GlobalServerSeries int
	ServerVersion      string
	Malini2            bool
	Execute2           bool
	LobEmptyCompOrcl   bool
	IsoLevel           int32
	ReadOnly           bool
	NewLobFlag         bool
	sslEncrypt         int
	MaxRowSize         int32
	DDLAutoCommit      bool
	BackslashEscape    bool
	SvrStat            int32
	SvrMode            int32
	ConstParaOpt       bool
	DbTimezone         int16
	LifeTimeRemainder  int16
	InstanceName       string
	Schema             string
	LastLoginIP        string
	LastLoginTime      string
	FailedAttempts     int32
	LoginWarningID     int32
	GraceTimeRemainder int32
	Guid               string
	DbName             string
	StandbyHost        string
	StandbyPort        int32
	StandbyCount       int32
	SessionID          int64
	OracleDateLanguage byte
	FormatDate         string
	FormatTimestamp    string
	FormatTimestampTZ  string
	FormatTime         string
	FormatTimeTZ       string
	Local              bool
	MsgVersion         int32
	TrxStatus          int32
	dscControl         bool
	trxFinish          bool
	autoCommit         bool
	isBatch            bool

	watching bool
	watcher  chan<- context.Context
	closech  chan struct{}
	finished chan<- struct{}
	canceled atomicError
	closed   atomicBool
}

func (dc *DmConnection) setTrxFinish(status int32) {
	switch status & Dm_build_413 {
	case Dm_build_410, Dm_build_411, Dm_build_412:
		dc.trxFinish = true
	default:
		dc.trxFinish = false
	}
}

func (dc *DmConnection) init() {

	dc.stmtMap = make(map[int32]*DmStatement)
	dc.DbTimezone = 0
	dc.GlobalServerSeries = 0
	dc.MaxRowSize = 0
	dc.LobEmptyCompOrcl = false
	dc.ReadOnly = false
	dc.DDLAutoCommit = false
	dc.ConstParaOpt = false
	dc.IsoLevel = -1
	dc.Malini2 = true
	dc.NewLobFlag = true
	dc.Execute2 = true
	dc.serverEncoding = ENCODING_GB18030
	dc.TrxStatus = Dm_build_361
	dc.setTrxFinish(dc.TrxStatus)
	dc.OracleDateLanguage = byte(Locale)
	dc.lastExecInfo = NewExceInfo()
	dc.MsgVersion = Dm_build_294

	dc.idGenerator = dmConnIDGenerator
}

func (dc *DmConnection) reset() {
	dc.DbTimezone = 0
	dc.GlobalServerSeries = 0
	dc.MaxRowSize = 0
	dc.LobEmptyCompOrcl = false
	dc.ReadOnly = false
	dc.DDLAutoCommit = false
	dc.ConstParaOpt = false
	dc.IsoLevel = -1
	dc.Malini2 = true
	dc.NewLobFlag = true
	dc.Execute2 = true
	dc.serverEncoding = ENCODING_GB18030
	dc.TrxStatus = Dm_build_361
	dc.setTrxFinish(dc.TrxStatus)
}

func (dc *DmConnection) checkClosed() error {
	if dc.closed.IsSet() {
		return driver.ErrBadConn
	}

	return nil
}

func (dc *DmConnection) executeInner(query string, execType int16) (interface{}, error) {

	stmt, err := NewDmStmt(dc, query)

	if err != nil {
		return nil, err
	}

	if execType == Dm_build_378 {
		defer func(stmt *DmStatement) {
			_ = stmt.close()
		}(stmt)
	}

	stmt.innerUsed = true
	if stmt.dmConn.dmConnector.escapeProcess {
		stmt.nativeSql, err = stmt.dmConn.escape(stmt.nativeSql, stmt.dmConn.dmConnector.keyWords)
		if err != nil {
			_ = stmt.close()
			return nil, err
		}
	}

	var optParamList []OptParameter

	if stmt.dmConn.ConstParaOpt {
		optParamList = make([]OptParameter, 0)
		stmt.nativeSql, optParamList, err = stmt.dmConn.execOpt(stmt.nativeSql, optParamList, stmt.dmConn.getServerEncoding())
		if err != nil {
			_ = stmt.close()
			optParamList = nil
		}
	}

	if execType == Dm_build_377 && dc.dmConnector.enRsCache {
		rpv, err := rp.get(stmt, query)
		if err != nil {
			return nil, err
		}

		if rpv != nil {
			stmt.execInfo = rpv.execInfo
			dc.lastExecInfo = rpv.execInfo
			return newDmRows(rpv.getResultSet(stmt)), nil
		}
	}

	var info *execRetInfo

	if optParamList != nil && len(optParamList) > 0 {
		info, err = dc.Access.Dm_build_85(stmt, optParamList)
		if err != nil {
			stmt.nativeSql = query
			info, err = dc.Access.Dm_build_91(stmt, execType)
		}
	} else {
		info, err = dc.Access.Dm_build_91(stmt, execType)
	}

	if err != nil {
		_ = stmt.close()
		return nil, err
	}
	dc.lastExecInfo = info

	if execType == Dm_build_377 && info.hasResultSet {
		return newDmRows(newInnerRows(0, stmt, info)), nil
	} else {
		return newDmResult(stmt, info), nil
	}
}

func g2dbIsoLevel(isoLevel int32) int32 {
	switch isoLevel {
	case 1:
		return Dm_build_365
	case 2:
		return Dm_build_366
	case 4:
		return Dm_build_367
	case 6:
		return Dm_build_368
	default:
		return -1
	}
}

func (dc *DmConnection) Begin() (driver.Tx, error) {
	if len(dc.filterChain.filters) == 0 {
		return dc.begin()
	} else {
		return dc.filterChain.reset().DmConnectionBegin(dc)
	}
}

func (dc *DmConnection) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	if len(dc.filterChain.filters) == 0 {
		return dc.beginTx(ctx, opts)
	}
	return dc.filterChain.reset().DmConnectionBeginTx(dc, ctx, opts)
}

func (dc *DmConnection) Commit() error {
	if len(dc.filterChain.filters) == 0 {
		return dc.commit()
	} else {
		return dc.filterChain.reset().DmConnectionCommit(dc)
	}
}

func (dc *DmConnection) Rollback() error {
	if len(dc.filterChain.filters) == 0 {
		return dc.rollback()
	} else {
		return dc.filterChain.reset().DmConnectionRollback(dc)
	}
}

func (dc *DmConnection) Close() error {
	if len(dc.filterChain.filters) == 0 {
		return dc.close()
	} else {
		return dc.filterChain.reset().DmConnectionClose(dc)
	}
}

func (dc *DmConnection) Ping(ctx context.Context) error {
	if len(dc.filterChain.filters) == 0 {
		return dc.ping(ctx)
	} else {
		return dc.filterChain.reset().DmConnectionPing(dc, ctx)
	}
}

func (dc *DmConnection) Exec(query string, args []driver.Value) (driver.Result, error) {
	if len(dc.filterChain.filters) == 0 {
		return dc.exec(query, args)
	}
	return dc.filterChain.reset().DmConnectionExec(dc, query, args)
}

func (dc *DmConnection) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	if len(dc.filterChain.filters) == 0 {
		return dc.execContext(ctx, query, args)
	}
	return dc.filterChain.reset().DmConnectionExecContext(dc, ctx, query, args)
}

func (dc *DmConnection) Query(query string, args []driver.Value) (driver.Rows, error) {
	if len(dc.filterChain.filters) == 0 {
		return dc.query(query, args)
	}
	return dc.filterChain.reset().DmConnectionQuery(dc, query, args)
}

func (dc *DmConnection) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	if len(dc.filterChain.filters) == 0 {
		return dc.queryContext(ctx, query, args)
	}
	return dc.filterChain.reset().DmConnectionQueryContext(dc, ctx, query, args)
}

func (dc *DmConnection) Prepare(query string) (driver.Stmt, error) {
	if len(dc.filterChain.filters) == 0 {
		return dc.prepare(query)
	}
	return dc.filterChain.reset().DmConnectionPrepare(dc, query)
}

func (dc *DmConnection) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if len(dc.filterChain.filters) == 0 {
		return dc.prepareContext(ctx, query)
	}
	return dc.filterChain.reset().DmConnectionPrepareContext(dc, ctx, query)
}

func (dc *DmConnection) ResetSession(ctx context.Context) error {
	if len(dc.filterChain.filters) == 0 {
		return dc.resetSession(ctx)
	}
	if err := dc.filterChain.reset().DmConnectionResetSession(dc, ctx); err != nil {
		return driver.ErrBadConn
	} else {
		return nil
	}
}

func (dc *DmConnection) CheckNamedValue(nv *driver.NamedValue) error {
	if len(dc.filterChain.filters) == 0 {
		return dc.checkNamedValue(nv)
	}
	return dc.filterChain.reset().DmConnectionCheckNamedValue(dc, nv)
}

func (dc *DmConnection) begin() (*DmConnection, error) {
	return dc.beginTx(context.Background(), driver.TxOptions{Isolation: driver.IsolationLevel(sql.LevelDefault)})
}

func (dc *DmConnection) beginTx(ctx context.Context, opts driver.TxOptions) (*DmConnection, error) {
	if err := dc.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer dc.finish()

	err := dc.checkClosed()
	if err != nil {
		return nil, err
	}

	dc.autoCommit = false

	if dc.ReadOnly != opts.ReadOnly {
		dc.ReadOnly = opts.ReadOnly
		var readonly = 0
		if opts.ReadOnly {
			readonly = 1
		}
		_, _ = dc.exec(fmt.Sprintf("SP_SET_SESSION_READONLY(%d)", readonly), nil)
	}

	if dc.IsoLevel != int32(opts.Isolation) {
		switch sql.IsolationLevel(opts.Isolation) {
		case sql.LevelDefault:
			dc.IsoLevel = int32(sql.LevelReadCommitted)
		case sql.LevelReadUncommitted, sql.LevelReadCommitted, sql.LevelSerializable:
			dc.IsoLevel = int32(opts.Isolation)
		case sql.LevelRepeatableRead:
			if dc.CompatibleMysql() {
				dc.IsoLevel = int32(sql.LevelReadCommitted)
			} else {
				return nil, ECGO_INVALID_TRAN_ISOLATION.throw()
			}
		default:
			return nil, ECGO_INVALID_TRAN_ISOLATION.throw()
		}

		err = dc.Access.Dm_build_153(dc)
		if err != nil {
			return nil, err
		}
	}

	return dc, nil
}

func (dc *DmConnection) commit() error {
	err := dc.checkClosed()
	if err != nil {
		return err
	}

	defer func() {
		dc.autoCommit = dc.dmConnector.autoCommit
		if dc.ReadOnly {
			_, _ = dc.exec("SP_SET_SESSION_READONLY(0)", nil)
		}
	}()

	if !dc.autoCommit {
		err = dc.Access.Commit()
		if err != nil {
			return err
		}
		dc.trxFinish = true
		return nil
	} else if !dc.dmConnector.alwayseAllowCommit {
		return ECGO_COMMIT_IN_AUTOCOMMIT_MODE.throw()
	}

	return nil
}

func (dc *DmConnection) rollback() error {
	err := dc.checkClosed()
	if err != nil {
		return err
	}

	defer func() {
		dc.autoCommit = dc.dmConnector.autoCommit
		if dc.ReadOnly {
			_, _ = dc.exec("SP_SET_SESSION_READONLY(0)", nil)
		}
	}()

	if !dc.autoCommit {
		err = dc.Access.Rollback()
		if err != nil {
			return err
		}
		dc.trxFinish = true
		return nil
	} else if !dc.dmConnector.alwayseAllowCommit {
		return ECGO_ROLLBACK_IN_AUTOCOMMIT_MODE.throw()
	}

	return nil
}

func (dc *DmConnection) reconnect() error {
	err := dc.Access.Close()
	if err != nil {
		return err
	}

	for _, stmt := range dc.stmtMap {

		for id, rs := range stmt.rsMap {
			_ = rs.Close()
			delete(stmt.rsMap, id)
		}
	}

	var newConn *DmConnection
	if dc.dmConnector.group != nil {
		if newConn, err = dc.dmConnector.group.connect(dc.dmConnector); err != nil {
			return err
		}
	} else {
		newConn, err = dc.dmConnector.connect(context.Background())
	}

	oldMap := dc.stmtMap
	newConn.mu = dc.mu
	newConn.filterable = dc.filterable
	*dc = *newConn

	for _, stmt := range oldMap {
		if stmt.closed {
			continue
		}
		err = dc.Access.Dm_build_63(stmt)
		if err != nil {
			_ = stmt.free()
			continue
		}

		if stmt.prepared || stmt.paramCount > 0 {
			if err = stmt.prepare(); err != nil {
				continue
			}
		}

		dc.stmtMap[stmt.id] = stmt
	}

	return nil
}

func (dc *DmConnection) cleanup() {
	_ = dc.close()
}

func (dc *DmConnection) close() error {
	if !dc.closed.TrySet(true) {
		return nil
	}

	util.AbsorbPanic(func() {
		close(dc.closech)
	})
	if dc.Access == nil {
		return nil
	}

	_ = dc.rollback()

	for _, stmt := range dc.stmtMap {
		_ = stmt.free()
	}

	_ = dc.Access.Close()

	return nil
}

func (dc *DmConnection) ping(ctx context.Context) error {
	if err := dc.watchCancel(ctx); err != nil {
		return err
	}
	defer dc.finish()

	rows, err := dc.query("select 1", nil)
	if err != nil {
		return err
	}
	return rows.close()
}

func (dc *DmConnection) exec(query string, args []driver.Value) (*DmResult, error) {
	err := dc.checkClosed()
	if err != nil {
		return nil, err
	}

	if args != nil && len(args) > 0 {
		stmt, err := dc.prepare(query)
		if err != nil {
			return nil, err
		}
		defer func(stmt *DmStatement) {
			_ = stmt.close()
		}(stmt)
		dc.lastExecInfo = stmt.execInfo

		return stmt.exec(args)
	} else {
		r1, err := dc.executeInner(query, Dm_build_378)
		if err != nil {
			return nil, err
		}

		if r2, ok := r1.(*DmResult); ok {
			return r2, nil
		} else {
			return nil, ECGO_NOT_EXEC_SQL.throw()
		}
	}
}

func (dc *DmConnection) execContext(ctx context.Context, query string, args []driver.NamedValue) (*DmResult, error) {
	if err := dc.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer dc.finish()

	err := dc.checkClosed()
	if err != nil {
		return nil, err
	}

	if args != nil && len(args) > 0 {
		stmt, err := dc.prepare(query)
		if err != nil {
			return nil, err
		}
		defer func(stmt *DmStatement) {
			_ = stmt.close()
		}(stmt)
		dc.lastExecInfo = stmt.execInfo
		dargs, err := namedValueToValue(stmt, args)
		if err != nil {
			return nil, err
		}
		return stmt.exec(dargs)
	} else {
		r1, err := dc.executeInner(query, Dm_build_378)
		if err != nil {
			return nil, err
		}

		if r2, ok := r1.(*DmResult); ok {
			return r2, nil
		} else {
			return nil, ECGO_NOT_EXEC_SQL.throw()
		}
	}
}

func (dc *DmConnection) query(query string, args []driver.Value) (*DmRows, error) {

	err := dc.checkClosed()
	if err != nil {
		return nil, err
	}

	if args != nil && len(args) > 0 {
		stmt, err := dc.prepare(query)
		if err != nil {
			return nil, err
		}
		dc.lastExecInfo = stmt.execInfo

		stmt.innerUsed = true
		return stmt.query(args)

	} else {
		r1, err := dc.executeInner(query, Dm_build_377)
		if err != nil {
			return nil, err
		}

		if r2, ok := r1.(*DmRows); ok {
			return r2, nil
		} else {
			return nil, ECGO_NOT_QUERY_SQL.throw()
		}
	}
}

func (dc *DmConnection) queryContext(ctx context.Context, query string, args []driver.NamedValue) (*DmRows, error) {
	if err := dc.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer dc.finish()

	err := dc.checkClosed()
	if err != nil {
		return nil, err
	}

	if args != nil && len(args) > 0 {
		stmt, err := dc.prepare(query)
		if err != nil {
			return nil, err
		}
		dc.lastExecInfo = stmt.execInfo

		stmt.innerUsed = true
		dargs, err := namedValueToValue(stmt, args)
		if err != nil {
			return nil, err
		}
		return stmt.query(dargs)

	} else {
		r1, err := dc.executeInner(query, Dm_build_377)
		if err != nil {
			return nil, err
		}

		if r2, ok := r1.(*DmRows); ok {
			return r2, nil
		} else {
			return nil, ECGO_NOT_QUERY_SQL.throw()
		}
	}

}

func (dc *DmConnection) prepare(query string) (stmt *DmStatement, err error) {
	if err = dc.checkClosed(); err != nil {
		return
	}
	if stmt, err = NewDmStmt(dc, query); err != nil {
		return
	}
	if err = stmt.prepare(); err != nil {
		_ = stmt.close()
		stmt = nil
		return
	}
	return
}

func (dc *DmConnection) prepareContext(ctx context.Context, query string) (*DmStatement, error) {
	if err := dc.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer dc.finish()

	return dc.prepare(query)
}

func (dc *DmConnection) resetSession(ctx context.Context) error {
	if err := dc.watchCancel(ctx); err != nil {
		return err
	}
	defer dc.finish()

	err := dc.checkClosed()
	if err != nil {
		return err
	}

	return nil
}

func (dc *DmConnection) checkNamedValue(nv *driver.NamedValue) error {
	var err error
	var cvt = converter{dc, false}
	nv.Value, err = cvt.ConvertValue(nv.Value)
	dc.isBatch = cvt.isBatch
	return err
}

func (dc *DmConnection) driverQuery(query string) (*DmStatement, *DmRows, error) {
	stmt, err := NewDmStmt(dc, query)
	if err != nil {
		return nil, nil, err
	}
	stmt.innerUsed = true
	stmt.innerExec = true
	info, err := dc.Access.Dm_build_91(stmt, Dm_build_377)
	if err != nil {
		return nil, nil, err
	}
	dc.lastExecInfo = info
	stmt.innerExec = false
	return stmt, newDmRows(newInnerRows(0, stmt, info)), nil
}

func (dc *DmConnection) getIndexOnEPGroup() int32 {
	if dc.dmConnector.group == nil || dc.dmConnector.group.epList == nil {
		return -1
	}
	for i := 0; i < len(dc.dmConnector.group.epList); i++ {
		ep := dc.dmConnector.group.epList[i]
		if dc.dmConnector.host == ep.host && dc.dmConnector.port == ep.port {
			return int32(i)
		}
	}
	return -1
}

func (dc *DmConnection) getServerEncoding() string {
	if dc.dmConnector.charCode != "" {
		return dc.dmConnector.charCode
	}
	return dc.serverEncoding
}

func (dc *DmConnection) lobFetchAll() bool {
	return dc.dmConnector.lobMode == 2
}

func (dc *DmConnection) CompatibleOracle() bool {
	return dc.dmConnector.compatibleMode == COMPATIBLE_MODE_ORACLE
}

func (dc *DmConnection) CompatibleMysql() bool {
	return dc.dmConnector.compatibleMode == COMPATIBLE_MODE_MYSQL
}

func (dc *DmConnection) cancel(err error) {
	dc.canceled.Set(err)
	_ = dc.close()

}

func (dc *DmConnection) finish() {
	if !dc.watching || dc.finished == nil {
		return
	}
	select {
	case dc.finished <- struct{}{}:
		dc.watching = false
	case <-dc.closech:
	}
}

func (dc *DmConnection) startWatcher() {
	watcher := make(chan context.Context, 1)
	dc.watcher = watcher
	finished := make(chan struct{})
	dc.finished = finished
	go func() {
		for {
			var ctx context.Context
			select {
			case ctx = <-watcher:
			case <-dc.closech:
				return
			}

			select {
			case <-ctx.Done():
				dc.cancel(ctx.Err())
			case <-finished:
			case <-dc.closech:
				return
			}
		}
	}()
}

func (dc *DmConnection) watchCancel(ctx context.Context) error {
	if dc.watching {

		dc.cleanup()
		return nil
	}

	if err := ctx.Err(); err != nil {
		return err
	}

	if ctx.Done() == nil {
		return nil
	}

	if dc.watcher == nil {
		return nil
	}

	dc.watching = true
	dc.watcher <- ctx
	return nil
}

type noCopy struct{}

func (*noCopy) Lock() {}

type atomicBool struct {
	_noCopy noCopy
	value   uint32
}

func (ab *atomicBool) IsSet() bool {
	return atomic.LoadUint32(&ab.value) > 0
}

func (ab *atomicBool) Set(value bool) {
	if value {
		atomic.StoreUint32(&ab.value, 1)
	} else {
		atomic.StoreUint32(&ab.value, 0)
	}
}

func (ab *atomicBool) TrySet(value bool) bool {
	if value {
		return atomic.SwapUint32(&ab.value, 1) == 0
	}
	return atomic.SwapUint32(&ab.value, 0) > 0
}

type atomicError struct {
	_noCopy noCopy
	value   atomic.Value
}

func (ae *atomicError) Set(value error) {
	ae.value.Store(value)
}

func (ae *atomicError) Value() error {
	if v := ae.value.Load(); v != nil {

		return v.(error)
	}
	return nil
}
