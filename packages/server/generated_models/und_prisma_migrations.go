// Code generated by SQLBoiler 4.8.6 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// PrismaMigration is an object representing the database table.
type PrismaMigration struct {
	ID                string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	Checksum          string      `boil:"checksum" json:"checksum" toml:"checksum" yaml:"checksum"`
	FinishedAt        null.Time   `boil:"finished_at" json:"finished_at,omitempty" toml:"finished_at" yaml:"finished_at,omitempty"`
	MigrationName     string      `boil:"migration_name" json:"migration_name" toml:"migration_name" yaml:"migration_name"`
	Logs              null.String `boil:"logs" json:"logs,omitempty" toml:"logs" yaml:"logs,omitempty"`
	RolledBackAt      null.Time   `boil:"rolled_back_at" json:"rolled_back_at,omitempty" toml:"rolled_back_at" yaml:"rolled_back_at,omitempty"`
	StartedAt         time.Time   `boil:"started_at" json:"started_at" toml:"started_at" yaml:"started_at"`
	AppliedStepsCount string      `boil:"applied_steps_count" json:"applied_steps_count" toml:"applied_steps_count" yaml:"applied_steps_count"`

	R *prismaMigrationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L prismaMigrationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PrismaMigrationColumns = struct {
	ID                string
	Checksum          string
	FinishedAt        string
	MigrationName     string
	Logs              string
	RolledBackAt      string
	StartedAt         string
	AppliedStepsCount string
}{
	ID:                "id",
	Checksum:          "checksum",
	FinishedAt:        "finished_at",
	MigrationName:     "migration_name",
	Logs:              "logs",
	RolledBackAt:      "rolled_back_at",
	StartedAt:         "started_at",
	AppliedStepsCount: "applied_steps_count",
}

var PrismaMigrationTableColumns = struct {
	ID                string
	Checksum          string
	FinishedAt        string
	MigrationName     string
	Logs              string
	RolledBackAt      string
	StartedAt         string
	AppliedStepsCount string
}{
	ID:                "_prisma_migrations.id",
	Checksum:          "_prisma_migrations.checksum",
	FinishedAt:        "_prisma_migrations.finished_at",
	MigrationName:     "_prisma_migrations.migration_name",
	Logs:              "_prisma_migrations.logs",
	RolledBackAt:      "_prisma_migrations.rolled_back_at",
	StartedAt:         "_prisma_migrations.started_at",
	AppliedStepsCount: "_prisma_migrations.applied_steps_count",
}

// Generated where

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var PrismaMigrationWhere = struct {
	ID                whereHelperstring
	Checksum          whereHelperstring
	FinishedAt        whereHelpernull_Time
	MigrationName     whereHelperstring
	Logs              whereHelpernull_String
	RolledBackAt      whereHelpernull_Time
	StartedAt         whereHelpertime_Time
	AppliedStepsCount whereHelperstring
}{
	ID:                whereHelperstring{field: "\"_prisma_migrations\".\"id\""},
	Checksum:          whereHelperstring{field: "\"_prisma_migrations\".\"checksum\""},
	FinishedAt:        whereHelpernull_Time{field: "\"_prisma_migrations\".\"finished_at\""},
	MigrationName:     whereHelperstring{field: "\"_prisma_migrations\".\"migration_name\""},
	Logs:              whereHelpernull_String{field: "\"_prisma_migrations\".\"logs\""},
	RolledBackAt:      whereHelpernull_Time{field: "\"_prisma_migrations\".\"rolled_back_at\""},
	StartedAt:         whereHelpertime_Time{field: "\"_prisma_migrations\".\"started_at\""},
	AppliedStepsCount: whereHelperstring{field: "\"_prisma_migrations\".\"applied_steps_count\""},
}

// PrismaMigrationRels is where relationship names are stored.
var PrismaMigrationRels = struct {
}{}

// prismaMigrationR is where relationships are stored.
type prismaMigrationR struct {
}

// NewStruct creates a new relationship struct
func (*prismaMigrationR) NewStruct() *prismaMigrationR {
	return &prismaMigrationR{}
}

// prismaMigrationL is where Load methods for each relationship are stored.
type prismaMigrationL struct{}

var (
	prismaMigrationAllColumns            = []string{"id", "checksum", "finished_at", "migration_name", "logs", "rolled_back_at", "started_at", "applied_steps_count"}
	prismaMigrationColumnsWithoutDefault = []string{"id", "checksum", "migration_name"}
	prismaMigrationColumnsWithDefault    = []string{"finished_at", "logs", "rolled_back_at", "started_at", "applied_steps_count"}
	prismaMigrationPrimaryKeyColumns     = []string{"id"}
	prismaMigrationGeneratedColumns      = []string{}
)

type (
	// PrismaMigrationSlice is an alias for a slice of pointers to PrismaMigration.
	// This should almost always be used instead of []PrismaMigration.
	PrismaMigrationSlice []*PrismaMigration

	prismaMigrationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	prismaMigrationType                 = reflect.TypeOf(&PrismaMigration{})
	prismaMigrationMapping              = queries.MakeStructMapping(prismaMigrationType)
	prismaMigrationPrimaryKeyMapping, _ = queries.BindMapping(prismaMigrationType, prismaMigrationMapping, prismaMigrationPrimaryKeyColumns)
	prismaMigrationInsertCacheMut       sync.RWMutex
	prismaMigrationInsertCache          = make(map[string]insertCache)
	prismaMigrationUpdateCacheMut       sync.RWMutex
	prismaMigrationUpdateCache          = make(map[string]updateCache)
	prismaMigrationUpsertCacheMut       sync.RWMutex
	prismaMigrationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single prismaMigration record from the query.
func (q prismaMigrationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PrismaMigration, error) {
	o := &PrismaMigration{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for _prisma_migrations")
	}

	return o, nil
}

// All returns all PrismaMigration records from the query.
func (q prismaMigrationQuery) All(ctx context.Context, exec boil.ContextExecutor) (PrismaMigrationSlice, error) {
	var o []*PrismaMigration

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PrismaMigration slice")
	}

	return o, nil
}

// Count returns the count of all PrismaMigration records in the query.
func (q prismaMigrationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count _prisma_migrations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q prismaMigrationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if _prisma_migrations exists")
	}

	return count > 0, nil
}

// PrismaMigrations retrieves all the records using an executor.
func PrismaMigrations(mods ...qm.QueryMod) prismaMigrationQuery {
	mods = append(mods, qm.From("\"_prisma_migrations\""))
	return prismaMigrationQuery{NewQuery(mods...)}
}

// FindPrismaMigration retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPrismaMigration(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*PrismaMigration, error) {
	prismaMigrationObj := &PrismaMigration{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"_prisma_migrations\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, prismaMigrationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from _prisma_migrations")
	}

	return prismaMigrationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PrismaMigration) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no _prisma_migrations provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(prismaMigrationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	prismaMigrationInsertCacheMut.RLock()
	cache, cached := prismaMigrationInsertCache[key]
	prismaMigrationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			prismaMigrationAllColumns,
			prismaMigrationColumnsWithDefault,
			prismaMigrationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(prismaMigrationType, prismaMigrationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(prismaMigrationType, prismaMigrationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"_prisma_migrations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"_prisma_migrations\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into _prisma_migrations")
	}

	if !cached {
		prismaMigrationInsertCacheMut.Lock()
		prismaMigrationInsertCache[key] = cache
		prismaMigrationInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the PrismaMigration.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PrismaMigration) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	prismaMigrationUpdateCacheMut.RLock()
	cache, cached := prismaMigrationUpdateCache[key]
	prismaMigrationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			prismaMigrationAllColumns,
			prismaMigrationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update _prisma_migrations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"_prisma_migrations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, prismaMigrationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(prismaMigrationType, prismaMigrationMapping, append(wl, prismaMigrationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update _prisma_migrations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for _prisma_migrations")
	}

	if !cached {
		prismaMigrationUpdateCacheMut.Lock()
		prismaMigrationUpdateCache[key] = cache
		prismaMigrationUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q prismaMigrationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for _prisma_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for _prisma_migrations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PrismaMigrationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), prismaMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"_prisma_migrations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, prismaMigrationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in prismaMigration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all prismaMigration")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PrismaMigration) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no _prisma_migrations provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(prismaMigrationColumnsWithDefault, o)

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

	prismaMigrationUpsertCacheMut.RLock()
	cache, cached := prismaMigrationUpsertCache[key]
	prismaMigrationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			prismaMigrationAllColumns,
			prismaMigrationColumnsWithDefault,
			prismaMigrationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			prismaMigrationAllColumns,
			prismaMigrationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert _prisma_migrations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(prismaMigrationPrimaryKeyColumns))
			copy(conflict, prismaMigrationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"_prisma_migrations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(prismaMigrationType, prismaMigrationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(prismaMigrationType, prismaMigrationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert _prisma_migrations")
	}

	if !cached {
		prismaMigrationUpsertCacheMut.Lock()
		prismaMigrationUpsertCache[key] = cache
		prismaMigrationUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single PrismaMigration record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PrismaMigration) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PrismaMigration provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), prismaMigrationPrimaryKeyMapping)
	sql := "DELETE FROM \"_prisma_migrations\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from _prisma_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for _prisma_migrations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q prismaMigrationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no prismaMigrationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from _prisma_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for _prisma_migrations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PrismaMigrationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), prismaMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"_prisma_migrations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, prismaMigrationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from prismaMigration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for _prisma_migrations")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *PrismaMigration) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPrismaMigration(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PrismaMigrationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PrismaMigrationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), prismaMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"_prisma_migrations\".* FROM \"_prisma_migrations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, prismaMigrationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PrismaMigrationSlice")
	}

	*o = slice

	return nil
}

// PrismaMigrationExists checks if the PrismaMigration row exists.
func PrismaMigrationExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"_prisma_migrations\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if _prisma_migrations exists")
	}

	return exists, nil
}
