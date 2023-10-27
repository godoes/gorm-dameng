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
	Access      *dm_build_414
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

func (conn *DmConnection) setTrxFinish(status int32) {
	switch status & Dm_build_814 {
	case Dm_build_811, Dm_build_812, Dm_build_813:
		conn.trxFinish = true
	default:
		conn.trxFinish = false
	}
}

func (conn *DmConnection) init() {

	conn.stmtMap = make(map[int32]*DmStatement)
	conn.DbTimezone = 0
	conn.GlobalServerSeries = 0
	conn.MaxRowSize = 0
	conn.LobEmptyCompOrcl = false
	conn.ReadOnly = false
	conn.DDLAutoCommit = false
	conn.ConstParaOpt = false
	conn.IsoLevel = -1
	conn.Malini2 = true
	conn.NewLobFlag = true
	conn.Execute2 = true
	conn.serverEncoding = ENCODING_GB18030
	conn.TrxStatus = Dm_build_762
	conn.setTrxFinish(conn.TrxStatus)
	conn.OracleDateLanguage = byte(Locale)
	conn.lastExecInfo = NewExceInfo()
	conn.MsgVersion = Dm_build_695

	conn.idGenerator = dmConnIDGenerator
}

func (conn *DmConnection) reset() {
	conn.DbTimezone = 0
	conn.GlobalServerSeries = 0
	conn.MaxRowSize = 0
	conn.LobEmptyCompOrcl = false
	conn.ReadOnly = false
	conn.DDLAutoCommit = false
	conn.ConstParaOpt = false
	conn.IsoLevel = -1
	conn.Malini2 = true
	conn.NewLobFlag = true
	conn.Execute2 = true
	conn.serverEncoding = ENCODING_GB18030
	conn.TrxStatus = Dm_build_762
	conn.setTrxFinish(conn.TrxStatus)
}

func (conn *DmConnection) checkClosed() error {
	if conn.closed.IsSet() {
		return driver.ErrBadConn
	}

	return nil
}

func (conn *DmConnection) executeInner(query string, execType int16) (interface{}, error) {

	stmt, err := NewDmStmt(conn, query)

	if err != nil {
		return nil, err
	}

	if execType == Dm_build_779 {
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

	if execType == Dm_build_778 && conn.dmConnector.enRsCache {
		rpv, err := rp.get(stmt, query)
		if err != nil {
			return nil, err
		}

		if rpv != nil {
			stmt.execInfo = rpv.execInfo
			conn.lastExecInfo = rpv.execInfo
			return newDmRows(rpv.getResultSet(stmt)), nil
		}
	}

	var info *execRetInfo

	if optParamList != nil && len(optParamList) > 0 {
		info, err = conn.Access.Dm_build_494(stmt, optParamList)
		if err != nil {
			stmt.nativeSql = query
			info, err = conn.Access.Dm_build_500(stmt, execType)
		}
	} else {
		info, err = conn.Access.Dm_build_500(stmt, execType)
	}

	if err != nil {
		_ = stmt.close()
		return nil, err
	}
	conn.lastExecInfo = info

	if execType == Dm_build_778 && info.hasResultSet {
		return newDmRows(newInnerRows(0, stmt, info)), nil
	} else {
		return newDmResult(stmt, info), nil
	}
}

func g2dbIsoLevel(isoLevel int32) int32 {
	switch isoLevel {
	case 1:
		return Dm_build_766
	case 2:
		return Dm_build_767
	case 4:
		return Dm_build_768
	case 6:
		return Dm_build_769
	default:
		return -1
	}
}

func (conn *DmConnection) Begin() (driver.Tx, error) {
	if len(conn.filterChain.filters) == 0 {
		return conn.begin()
	} else {
		return conn.filterChain.reset().DmConnectionBegin(conn)
	}
}

func (conn *DmConnection) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	if len(conn.filterChain.filters) == 0 {
		return conn.beginTx(ctx, opts)
	}
	return conn.filterChain.reset().DmConnectionBeginTx(conn, ctx, opts)
}

func (conn *DmConnection) Commit() error {
	if len(conn.filterChain.filters) == 0 {
		return conn.commit()
	} else {
		return conn.filterChain.reset().DmConnectionCommit(conn)
	}
}

func (conn *DmConnection) Rollback() error {
	if len(conn.filterChain.filters) == 0 {
		return conn.rollback()
	} else {
		return conn.filterChain.reset().DmConnectionRollback(conn)
	}
}

func (conn *DmConnection) Close() error {
	if len(conn.filterChain.filters) == 0 {
		return conn.close()
	} else {
		return conn.filterChain.reset().DmConnectionClose(conn)
	}
}

func (conn *DmConnection) Ping(ctx context.Context) error {
	if len(conn.filterChain.filters) == 0 {
		return conn.ping(ctx)
	} else {
		return conn.filterChain.reset().DmConnectionPing(conn, ctx)
	}
}

func (conn *DmConnection) Exec(query string, args []driver.Value) (driver.Result, error) {
	if len(conn.filterChain.filters) == 0 {
		return conn.exec(query, args)
	}
	return conn.filterChain.reset().DmConnectionExec(conn, query, args)
}

func (conn *DmConnection) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	if len(conn.filterChain.filters) == 0 {
		return conn.execContext(ctx, query, args)
	}
	return conn.filterChain.reset().DmConnectionExecContext(conn, ctx, query, args)
}

func (conn *DmConnection) Query(query string, args []driver.Value) (driver.Rows, error) {
	if len(conn.filterChain.filters) == 0 {
		return conn.query(query, args)
	}
	return conn.filterChain.reset().DmConnectionQuery(conn, query, args)
}

func (conn *DmConnection) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	if len(conn.filterChain.filters) == 0 {
		return conn.queryContext(ctx, query, args)
	}
	return conn.filterChain.reset().DmConnectionQueryContext(conn, ctx, query, args)
}

func (conn *DmConnection) Prepare(query string) (driver.Stmt, error) {
	if len(conn.filterChain.filters) == 0 {
		return conn.prepare(query)
	}
	return conn.filterChain.reset().DmConnectionPrepare(conn, query)
}

func (conn *DmConnection) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if len(conn.filterChain.filters) == 0 {
		return conn.prepareContext(ctx, query)
	}
	return conn.filterChain.reset().DmConnectionPrepareContext(conn, ctx, query)
}

func (conn *DmConnection) ResetSession(ctx context.Context) error {
	if len(conn.filterChain.filters) == 0 {
		return conn.resetSession(ctx)
	}
	if err := conn.filterChain.reset().DmConnectionResetSession(conn, ctx); err != nil {
		return driver.ErrBadConn
	} else {
		return nil
	}
}

func (conn *DmConnection) CheckNamedValue(nv *driver.NamedValue) error {
	if len(conn.filterChain.filters) == 0 {
		return conn.checkNamedValue(nv)
	}
	return conn.filterChain.reset().DmConnectionCheckNamedValue(conn, nv)
}

func (conn *DmConnection) begin() (*DmConnection, error) {
	return conn.beginTx(context.Background(), driver.TxOptions{Isolation: driver.IsolationLevel(sql.LevelDefault)})
}

func (conn *DmConnection) beginTx(ctx context.Context, opts driver.TxOptions) (*DmConnection, error) {
	if err := conn.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer conn.finish()

	err := conn.checkClosed()
	if err != nil {
		return nil, err
	}

	conn.autoCommit = false

	if sql.IsolationLevel(opts.Isolation) == sql.LevelDefault {
		opts.Isolation = driver.IsolationLevel(sql.LevelReadCommitted)
	}

	if conn.ReadOnly != opts.ReadOnly {
		conn.ReadOnly = opts.ReadOnly
		var readonly = 0
		if opts.ReadOnly {
			readonly = 1
		}
		_, _ = conn.exec(fmt.Sprintf("SP_SET_SESSION_READONLY(%d)", readonly), nil)
	}

	if conn.IsoLevel != int32(opts.Isolation) {
		switch sql.IsolationLevel(opts.Isolation) {
		case sql.LevelDefault, sql.LevelReadUncommitted:
			return conn, nil
		case sql.LevelReadCommitted, sql.LevelSerializable:
			conn.IsoLevel = int32(opts.Isolation)
		case sql.LevelRepeatableRead:
			if conn.CompatibleMysql() {
				conn.IsoLevel = int32(sql.LevelReadCommitted)
			} else {
				return nil, ECGO_INVALID_TRAN_ISOLATION.throw()
			}
		default:
			return nil, ECGO_INVALID_TRAN_ISOLATION.throw()
		}

		err = conn.Access.Dm_build_554(conn)
		if err != nil {
			return nil, err
		}
	}

	return conn, nil
}

func (conn *DmConnection) commit() error {
	err := conn.checkClosed()
	if err != nil {
		return err
	}

	defer func() {
		conn.autoCommit = conn.dmConnector.autoCommit
		if conn.ReadOnly {
			_, _ = conn.exec("SP_SET_SESSION_READONLY(0)", nil)
		}
	}()

	if !conn.autoCommit {
		err = conn.Access.Commit()
		if err != nil {
			return err
		}
		conn.trxFinish = true
		return nil
	} else if !conn.dmConnector.alwayseAllowCommit {
		return ECGO_COMMIT_IN_AUTOCOMMIT_MODE.throw()
	}

	return nil
}

func (conn *DmConnection) rollback() error {
	err := conn.checkClosed()
	if err != nil {
		return err
	}

	defer func() {
		conn.autoCommit = conn.dmConnector.autoCommit
		if conn.ReadOnly {
			_, _ = conn.exec("SP_SET_SESSION_READONLY(0)", nil)
		}
	}()

	if !conn.autoCommit {
		err = conn.Access.Rollback()
		if err != nil {
			return err
		}
		conn.trxFinish = true
		return nil
	} else if !conn.dmConnector.alwayseAllowCommit {
		return ECGO_ROLLBACK_IN_AUTOCOMMIT_MODE.throw()
	}

	return nil
}

func (conn *DmConnection) reconnect() error {
	err := conn.Access.Close()
	if err != nil {
		return err
	}

	for _, stmt := range conn.stmtMap {

		for id, rs := range stmt.rsMap {
			_ = rs.Close()
			delete(stmt.rsMap, id)
		}
	}

	var newConn *DmConnection
	if conn.dmConnector.group != nil {
		if newConn, err = conn.dmConnector.group.connect(conn.dmConnector); err != nil {
			return err
		}
	} else {
		newConn, err = conn.dmConnector.connect(context.Background())
	}

	oldMap := conn.stmtMap
	newConn.mu = conn.mu
	newConn.filterable = conn.filterable
	*conn = *newConn

	for _, stmt := range oldMap {
		if stmt.closed {
			continue
		}
		err = conn.Access.Dm_build_472(stmt)
		if err != nil {
			_ = stmt.free()
			continue
		}

		if stmt.prepared || stmt.paramCount > 0 {
			if err = stmt.prepare(); err != nil {
				continue
			}
		}

		conn.stmtMap[stmt.id] = stmt
	}

	return nil
}

func (conn *DmConnection) cleanup() {
	_ = conn.close()
}

func (conn *DmConnection) close() error {
	if !conn.closed.TrySet(true) {
		return nil
	}

	util.AbsorbPanic(func() {
		close(conn.closech)
	})
	if conn.Access == nil {
		return nil
	}

	_ = conn.rollback()

	for _, stmt := range conn.stmtMap {
		_ = stmt.free()
	}

	_ = conn.Access.Close()

	return nil
}

func (conn *DmConnection) ping(ctx context.Context) error {
	if err := conn.watchCancel(ctx); err != nil {
		return err
	}
	defer conn.finish()

	rows, err := conn.query("select 1", nil)
	if err != nil {
		return err
	}
	return rows.close()
}

func (conn *DmConnection) exec(query string, args []driver.Value) (*DmResult, error) {
	err := conn.checkClosed()
	if err != nil {
		return nil, err
	}

	if args != nil && len(args) > 0 {
		stmt, err := conn.prepare(query)
		defer func(stmt *DmStatement) {
			_ = stmt.close()
		}(stmt)
		if err != nil {
			return nil, err
		}
		conn.lastExecInfo = stmt.execInfo

		return stmt.exec(args)
	} else {
		r1, err := conn.executeInner(query, Dm_build_779)
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

func (conn *DmConnection) execContext(ctx context.Context, query string, args []driver.NamedValue) (*DmResult, error) {
	if err := conn.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer conn.finish()

	err := conn.checkClosed()
	if err != nil {
		return nil, err
	}

	if args != nil && len(args) > 0 {
		stmt, err := conn.prepare(query)
		defer func(stmt *DmStatement) {
			_ = stmt.close()
		}(stmt)
		if err != nil {
			return nil, err
		}
		conn.lastExecInfo = stmt.execInfo
		dargs, err := namedValueToValue(stmt, args)
		if err != nil {
			return nil, err
		}
		return stmt.exec(dargs)
	} else {
		r1, err := conn.executeInner(query, Dm_build_779)
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

func (conn *DmConnection) query(query string, args []driver.Value) (*DmRows, error) {

	err := conn.checkClosed()
	if err != nil {
		return nil, err
	}

	if args != nil && len(args) > 0 {
		stmt, err := conn.prepare(query)
		if err != nil {
			_ = stmt.close()
			return nil, err
		}
		conn.lastExecInfo = stmt.execInfo

		stmt.innerUsed = true
		return stmt.query(args)

	} else {
		r1, err := conn.executeInner(query, Dm_build_778)
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

func (conn *DmConnection) queryContext(ctx context.Context, query string, args []driver.NamedValue) (*DmRows, error) {
	if err := conn.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer conn.finish()

	err := conn.checkClosed()
	if err != nil {
		return nil, err
	}

	if args != nil && len(args) > 0 {
		stmt, err := conn.prepare(query)
		if err != nil {
			if stmt != nil {
				_ = stmt.close()
			}
			return nil, err
		}
		conn.lastExecInfo = stmt.execInfo

		stmt.innerUsed = true
		dargs, err := namedValueToValue(stmt, args)
		if err != nil {
			return nil, err
		}
		return stmt.query(dargs)

	} else {
		r1, err := conn.executeInner(query, Dm_build_778)
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

func (conn *DmConnection) prepare(query string) (stmt *DmStatement, err error) {
	if err = conn.checkClosed(); err != nil {
		return
	}
	if stmt, err = NewDmStmt(conn, query); err != nil {
		return
	}
	err = stmt.prepare()
	return
}

func (conn *DmConnection) prepareContext(ctx context.Context, query string) (*DmStatement, error) {
	if err := conn.watchCancel(ctx); err != nil {
		return nil, err
	}
	defer conn.finish()

	return conn.prepare(query)
}

func (conn *DmConnection) resetSession(ctx context.Context) error {
	if err := conn.watchCancel(ctx); err != nil {
		return err
	}
	defer conn.finish()

	err := conn.checkClosed()
	if err != nil {
		return err
	}

	return nil
}

func (conn *DmConnection) checkNamedValue(nv *driver.NamedValue) error {
	var err error
	var cvt = converter{conn, false}
	nv.Value, err = cvt.ConvertValue(nv.Value)
	conn.isBatch = cvt.isBatch
	return err
}

func (conn *DmConnection) driverQuery(query string) (*DmStatement, *DmRows, error) {
	stmt, err := NewDmStmt(conn, query)
	if err != nil {
		return nil, nil, err
	}
	stmt.innerUsed = true
	stmt.innerExec = true
	info, err := conn.Access.Dm_build_500(stmt, Dm_build_778)
	if err != nil {
		return nil, nil, err
	}
	conn.lastExecInfo = info
	stmt.innerExec = false
	return stmt, newDmRows(newInnerRows(0, stmt, info)), nil
}

func (conn *DmConnection) getIndexOnEPGroup() int32 {
	if conn.dmConnector.group == nil || conn.dmConnector.group.epList == nil {
		return -1
	}
	for i := 0; i < len(conn.dmConnector.group.epList); i++ {
		ep := conn.dmConnector.group.epList[i]
		if conn.dmConnector.host == ep.host && conn.dmConnector.port == ep.port {
			return int32(i)
		}
	}
	return -1
}

func (conn *DmConnection) getServerEncoding() string {
	if conn.dmConnector.charCode != "" {
		return conn.dmConnector.charCode
	}
	return conn.serverEncoding
}

func (conn *DmConnection) lobFetchAll() bool {
	return conn.dmConnector.lobMode == 2
}

func (conn *DmConnection) CompatibleOracle() bool {
	return conn.dmConnector.compatibleMode == COMPATIBLE_MODE_ORACLE
}

func (conn *DmConnection) CompatibleMysql() bool {
	return conn.dmConnector.compatibleMode == COMPATIBLE_MODE_MYSQL
}

func (conn *DmConnection) cancel(err error) {
	conn.canceled.Set(err)
	_ = conn.close()

}

func (conn *DmConnection) finish() {
	if !conn.watching || conn.finished == nil {
		return
	}
	select {
	case conn.finished <- struct{}{}:
		conn.watching = false
	case <-conn.closech:
	}
}

func (conn *DmConnection) startWatcher() {
	watcher := make(chan context.Context, 1)
	conn.watcher = watcher
	finished := make(chan struct{})
	conn.finished = finished
	go func() {
		for {
			var ctx context.Context
			select {
			case ctx = <-watcher:
			case <-conn.closech:
				return
			}

			select {
			case <-ctx.Done():
				conn.cancel(ctx.Err())
			case <-finished:
			case <-conn.closech:
				return
			}
		}
	}()
}

func (conn *DmConnection) watchCancel(ctx context.Context) error {
	if conn.watching {

		conn.cleanup()
		return nil
	}

	if err := ctx.Err(); err != nil {
		return err
	}

	if ctx.Done() == nil {
		return nil
	}

	if conn.watcher == nil {
		return nil
	}

	conn.watching = true
	conn.watcher <- ctx
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
