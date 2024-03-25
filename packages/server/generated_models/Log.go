// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Log is an object representing the database table.
type Log struct {
	ID        string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt time.Time `boil:"createdAt" json:"createdAt" toml:"createdAt" yaml:"createdAt"`
	UpdatedAt time.Time `boil:"updatedAt" json:"updatedAt" toml:"updatedAt" yaml:"updatedAt"`
	UnitID    string    `boil:"unit_id" json:"unit_id" toml:"unit_id" yaml:"unit_id"`
	JobID     string    `boil:"job_id" json:"job_id" toml:"job_id" yaml:"job_id"`
	Message   string    `boil:"message" json:"message" toml:"message" yaml:"message"`

	R *logR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L logL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LogColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	UnitID    string
	JobID     string
	Message   string
}{
	ID:        "id",
	CreatedAt: "createdAt",
	UpdatedAt: "updatedAt",
	UnitID:    "unit_id",
	JobID:     "job_id",
	Message:   "message",
}

var LogTableColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	UnitID    string
	JobID     string
	Message   string
}{
	ID:        "Log.id",
	CreatedAt: "Log.createdAt",
	UpdatedAt: "Log.updatedAt",
	UnitID:    "Log.unit_id",
	JobID:     "Log.job_id",
	Message:   "Log.message",
}

// Generated where

var LogWhere = struct {
	ID        whereHelperstring
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
	UnitID    whereHelperstring
	JobID     whereHelperstring
	Message   whereHelperstring
}{
	ID:        whereHelperstring{field: "\"Log\".\"id\""},
	CreatedAt: whereHelpertime_Time{field: "\"Log\".\"createdAt\""},
	UpdatedAt: whereHelpertime_Time{field: "\"Log\".\"updatedAt\""},
	UnitID:    whereHelperstring{field: "\"Log\".\"unit_id\""},
	JobID:     whereHelperstring{field: "\"Log\".\"job_id\""},
	Message:   whereHelperstring{field: "\"Log\".\"message\""},
}

// LogRels is where relationship names are stored.
var LogRels = struct {
	Job  string
	Unit string
}{
	Job:  "Job",
	Unit: "Unit",
}

// logR is where relationships are stored.
type logR struct {
	Job  *Job  `boil:"Job" json:"Job" toml:"Job" yaml:"Job"`
	Unit *Unit `boil:"Unit" json:"Unit" toml:"Unit" yaml:"Unit"`
}

// NewStruct creates a new relationship struct
func (*logR) NewStruct() *logR {
	return &logR{}
}

func (r *logR) GetJob() *Job {
	if r == nil {
		return nil
	}
	return r.Job
}

func (r *logR) GetUnit() *Unit {
	if r == nil {
		return nil
	}
	return r.Unit
}

// logL is where Load methods for each relationship are stored.
type logL struct{}

var (
	logAllColumns            = []string{"id", "createdAt", "updatedAt", "unit_id", "job_id", "message"}
	logColumnsWithoutDefault = []string{"id", "updatedAt", "unit_id", "job_id", "message"}
	logColumnsWithDefault    = []string{"createdAt"}
	logPrimaryKeyColumns     = []string{"id"}
	logGeneratedColumns      = []string{}
)

type (
	// LogSlice is an alias for a slice of pointers to Log.
	// This should almost always be used instead of []Log.
	LogSlice []*Log

	logQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	logType                 = reflect.TypeOf(&Log{})
	logMapping              = queries.MakeStructMapping(logType)
	logPrimaryKeyMapping, _ = queries.BindMapping(logType, logMapping, logPrimaryKeyColumns)
	logInsertCacheMut       sync.RWMutex
	logInsertCache          = make(map[string]insertCache)
	logUpdateCacheMut       sync.RWMutex
	logUpdateCache          = make(map[string]updateCache)
	logUpsertCacheMut       sync.RWMutex
	logUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single log record from the query.
func (q logQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Log, error) {
	o := &Log{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for Log")
	}

	return o, nil
}

// All returns all Log records from the query.
func (q logQuery) All(ctx context.Context, exec boil.ContextExecutor) (LogSlice, error) {
	var o []*Log

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Log slice")
	}

	return o, nil
}

// Count returns the count of all Log records in the query.
func (q logQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count Log rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q logQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if Log exists")
	}

	return count > 0, nil
}

// Job pointed to by the foreign key.
func (o *Log) Job(mods ...qm.QueryMod) jobQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.JobID),
	}

	queryMods = append(queryMods, mods...)

	return Jobs(queryMods...)
}

// Unit pointed to by the foreign key.
func (o *Log) Unit(mods ...qm.QueryMod) unitQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UnitID),
	}

	queryMods = append(queryMods, mods...)

	return Units(queryMods...)
}

// LoadJob allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (logL) LoadJob(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLog interface{}, mods queries.Applicator) error {
	var slice []*Log
	var object *Log

	if singular {
		var ok bool
		object, ok = maybeLog.(*Log)
		if !ok {
			object = new(Log)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeLog))
			}
		}
	} else {
		s, ok := maybeLog.(*[]*Log)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeLog))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &logR{}
		}
		args[object.JobID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &logR{}
			}

			args[obj.JobID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`Job`),
		qm.WhereIn(`Job.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Job")
	}

	var resultSlice []*Job
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Job")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Job")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Job")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Job = foreign
		if foreign.R == nil {
			foreign.R = &jobR{}
		}
		foreign.R.JobLogs = append(foreign.R.JobLogs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.JobID == foreign.ID {
				local.R.Job = foreign
				if foreign.R == nil {
					foreign.R = &jobR{}
				}
				foreign.R.JobLogs = append(foreign.R.JobLogs, local)
				break
			}
		}
	}

	return nil
}

// LoadUnit allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (logL) LoadUnit(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLog interface{}, mods queries.Applicator) error {
	var slice []*Log
	var object *Log

	if singular {
		var ok bool
		object, ok = maybeLog.(*Log)
		if !ok {
			object = new(Log)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeLog))
			}
		}
	} else {
		s, ok := maybeLog.(*[]*Log)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeLog))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &logR{}
		}
		args[object.UnitID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &logR{}
			}

			args[obj.UnitID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`Unit`),
		qm.WhereIn(`Unit.id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Unit")
	}

	var resultSlice []*Unit
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Unit")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Unit")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Unit")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Unit = foreign
		if foreign.R == nil {
			foreign.R = &unitR{}
		}
		foreign.R.UnitLogs = append(foreign.R.UnitLogs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UnitID == foreign.ID {
				local.R.Unit = foreign
				if foreign.R == nil {
					foreign.R = &unitR{}
				}
				foreign.R.UnitLogs = append(foreign.R.UnitLogs, local)
				break
			}
		}
	}

	return nil
}

// SetJob of the log to the related item.
// Sets o.R.Job to related.
// Adds o to related.R.JobLogs.
func (o *Log) SetJob(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Job) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"Log\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"job_id"}),
		strmangle.WhereClause("\"", "\"", 0, logPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.JobID = related.ID
	if o.R == nil {
		o.R = &logR{
			Job: related,
		}
	} else {
		o.R.Job = related
	}

	if related.R == nil {
		related.R = &jobR{
			JobLogs: LogSlice{o},
		}
	} else {
		related.R.JobLogs = append(related.R.JobLogs, o)
	}

	return nil
}

// SetUnit of the log to the related item.
// Sets o.R.Unit to related.
// Adds o to related.R.UnitLogs.
func (o *Log) SetUnit(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Unit) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"Log\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"unit_id"}),
		strmangle.WhereClause("\"", "\"", 0, logPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UnitID = related.ID
	if o.R == nil {
		o.R = &logR{
			Unit: related,
		}
	} else {
		o.R.Unit = related
	}

	if related.R == nil {
		related.R = &unitR{
			UnitLogs: LogSlice{o},
		}
	} else {
		related.R.UnitLogs = append(related.R.UnitLogs, o)
	}

	return nil
}

// Logs retrieves all the records using an executor.
func Logs(mods ...qm.QueryMod) logQuery {
	mods = append(mods, qm.From("\"Log\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"Log\".*"})
	}

	return logQuery{q}
}

// FindLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLog(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Log, error) {
	logObj := &Log{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"Log\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, logObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from Log")
	}

	return logObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Log) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Log provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(logColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	logInsertCacheMut.RLock()
	cache, cached := logInsertCache[key]
	logInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			logAllColumns,
			logColumnsWithDefault,
			logColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(logType, logMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(logType, logMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"Log\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"Log\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into Log")
	}

	if !cached {
		logInsertCacheMut.Lock()
		logInsertCache[key] = cache
		logInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Log.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Log) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	logUpdateCacheMut.RLock()
	cache, cached := logUpdateCache[key]
	logUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			logAllColumns,
			logPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update Log, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"Log\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, logPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(logType, logMapping, append(wl, logPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update Log row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for Log")
	}

	if !cached {
		logUpdateCacheMut.Lock()
		logUpdateCache[key] = cache
		logUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q logQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for Log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for Log")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), logPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"Log\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, logPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in log slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all log")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Log) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Log provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(logColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	logUpsertCacheMut.RLock()
	cache, cached := logUpsertCache[key]
	logUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			logAllColumns,
			logColumnsWithDefault,
			logColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			logAllColumns,
			logPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert Log, could not build update column list")
		}

		ret := strmangle.SetComplement(logAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(logPrimaryKeyColumns))
			copy(conflict, logPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"Log\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(logType, logMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(logType, logMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert Log")
	}

	if !cached {
		logUpsertCacheMut.Lock()
		logUpsertCache[key] = cache
		logUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Log record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Log) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Log provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), logPrimaryKeyMapping)
	sql := "DELETE FROM \"Log\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from Log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for Log")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q logQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no logQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from Log")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Log")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), logPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"Log\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, logPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from log slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Log")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Log) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLog(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), logPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"Log\".* FROM \"Log\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, logPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LogSlice")
	}

	*o = slice

	return nil
}

// LogExists checks if the Log row exists.
func LogExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"Log\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if Log exists")
	}

	return exists, nil
}

// Exists checks if the Log row exists.
func (o *Log) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return LogExists(ctx, exec, o.ID)
}
