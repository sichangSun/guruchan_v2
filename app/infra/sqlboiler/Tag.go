// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package mysql

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

// Tag is an object representing the database table.
type Tag struct {
	TagId       int       `boil:"tagId" json:"tagId" toml:"tagId" yaml:"tagId"`
	UserId      int       `boil:"userId" json:"userId" toml:"userId" yaml:"userId"`
	TagTittle   string    `boil:"tagTittle" json:"tagTittle" toml:"tagTittle" yaml:"tagTittle"`
	Typecode    int       `boil:"typecode" json:"typecode" toml:"typecode" yaml:"typecode"`
	TAddTime    time.Time `boil:"t_addTime" json:"t_addTime" toml:"t_addTime" yaml:"t_addTime"`
	TUpdataTime null.Time `boil:"t_updataTime" json:"t_updataTime,omitempty" toml:"t_updataTime" yaml:"t_updataTime,omitempty"`
	// 1为删除
	IsDel int `boil:"isDel" json:"isDel" toml:"isDel" yaml:"isDel"`

	R *tagR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tagL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TagColumns = struct {
	TagId       string
	UserId      string
	TagTittle   string
	Typecode    string
	TAddTime    string
	TUpdataTime string
	IsDel       string
}{
	TagId:       "tagId",
	UserId:      "userId",
	TagTittle:   "tagTittle",
	Typecode:    "typecode",
	TAddTime:    "t_addTime",
	TUpdataTime: "t_updataTime",
	IsDel:       "isDel",
}

var TagTableColumns = struct {
	TagId       string
	UserId      string
	TagTittle   string
	Typecode    string
	TAddTime    string
	TUpdataTime string
	IsDel       string
}{
	TagId:       "Tag.tagId",
	UserId:      "Tag.userId",
	TagTittle:   "Tag.tagTittle",
	Typecode:    "Tag.typecode",
	TAddTime:    "Tag.t_addTime",
	TUpdataTime: "Tag.t_updataTime",
	IsDel:       "Tag.isDel",
}

// Generated where

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var TagWhere = struct {
	TagId       whereHelperint
	UserId      whereHelperint
	TagTittle   whereHelperstring
	Typecode    whereHelperint
	TAddTime    whereHelpertime_Time
	TUpdataTime whereHelpernull_Time
	IsDel       whereHelperint
}{
	TagId:       whereHelperint{field: "`Tag`.`tagId`"},
	UserId:      whereHelperint{field: "`Tag`.`userId`"},
	TagTittle:   whereHelperstring{field: "`Tag`.`tagTittle`"},
	Typecode:    whereHelperint{field: "`Tag`.`typecode`"},
	TAddTime:    whereHelpertime_Time{field: "`Tag`.`t_addTime`"},
	TUpdataTime: whereHelpernull_Time{field: "`Tag`.`t_updataTime`"},
	IsDel:       whereHelperint{field: "`Tag`.`isDel`"},
}

// TagRels is where relationship names are stored.
var TagRels = struct {
}{}

// tagR is where relationships are stored.
type tagR struct {
}

// NewStruct creates a new relationship struct
func (*tagR) NewStruct() *tagR {
	return &tagR{}
}

// tagL is where Load methods for each relationship are stored.
type tagL struct{}

var (
	tagAllColumns            = []string{"tagId", "userId", "tagTittle", "typecode", "t_addTime", "t_updataTime", "isDel"}
	tagColumnsWithoutDefault = []string{"userId", "tagTittle", "t_addTime", "t_updataTime"}
	tagColumnsWithDefault    = []string{"tagId", "typecode", "isDel"}
	tagPrimaryKeyColumns     = []string{"tagId"}
	tagGeneratedColumns      = []string{}
)

type (
	// TagSlice is an alias for a slice of pointers to Tag.
	// This should almost always be used instead of []Tag.
	TagSlice []*Tag
	// TagHook is the signature for custom Tag hook methods
	TagHook func(context.Context, boil.ContextExecutor, *Tag) error

	tagQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tagType                 = reflect.TypeOf(&Tag{})
	tagMapping              = queries.MakeStructMapping(tagType)
	tagPrimaryKeyMapping, _ = queries.BindMapping(tagType, tagMapping, tagPrimaryKeyColumns)
	tagInsertCacheMut       sync.RWMutex
	tagInsertCache          = make(map[string]insertCache)
	tagUpdateCacheMut       sync.RWMutex
	tagUpdateCache          = make(map[string]updateCache)
	tagUpsertCacheMut       sync.RWMutex
	tagUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tagAfterSelectMu sync.Mutex
var tagAfterSelectHooks []TagHook

var tagBeforeInsertMu sync.Mutex
var tagBeforeInsertHooks []TagHook
var tagAfterInsertMu sync.Mutex
var tagAfterInsertHooks []TagHook

var tagBeforeUpdateMu sync.Mutex
var tagBeforeUpdateHooks []TagHook
var tagAfterUpdateMu sync.Mutex
var tagAfterUpdateHooks []TagHook

var tagBeforeDeleteMu sync.Mutex
var tagBeforeDeleteHooks []TagHook
var tagAfterDeleteMu sync.Mutex
var tagAfterDeleteHooks []TagHook

var tagBeforeUpsertMu sync.Mutex
var tagBeforeUpsertHooks []TagHook
var tagAfterUpsertMu sync.Mutex
var tagAfterUpsertHooks []TagHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Tag) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Tag) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Tag) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Tag) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Tag) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Tag) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Tag) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Tag) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Tag) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tagAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTagHook registers your hook function for all future operations.
func AddTagHook(hookPoint boil.HookPoint, tagHook TagHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		tagAfterSelectMu.Lock()
		tagAfterSelectHooks = append(tagAfterSelectHooks, tagHook)
		tagAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		tagBeforeInsertMu.Lock()
		tagBeforeInsertHooks = append(tagBeforeInsertHooks, tagHook)
		tagBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		tagAfterInsertMu.Lock()
		tagAfterInsertHooks = append(tagAfterInsertHooks, tagHook)
		tagAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		tagBeforeUpdateMu.Lock()
		tagBeforeUpdateHooks = append(tagBeforeUpdateHooks, tagHook)
		tagBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		tagAfterUpdateMu.Lock()
		tagAfterUpdateHooks = append(tagAfterUpdateHooks, tagHook)
		tagAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		tagBeforeDeleteMu.Lock()
		tagBeforeDeleteHooks = append(tagBeforeDeleteHooks, tagHook)
		tagBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		tagAfterDeleteMu.Lock()
		tagAfterDeleteHooks = append(tagAfterDeleteHooks, tagHook)
		tagAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		tagBeforeUpsertMu.Lock()
		tagBeforeUpsertHooks = append(tagBeforeUpsertHooks, tagHook)
		tagBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		tagAfterUpsertMu.Lock()
		tagAfterUpsertHooks = append(tagAfterUpsertHooks, tagHook)
		tagAfterUpsertMu.Unlock()
	}
}

// One returns a single tag record from the query.
func (q tagQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Tag, error) {
	o := &Tag{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mysql: failed to execute a one query for Tag")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Tag records from the query.
func (q tagQuery) All(ctx context.Context, exec boil.ContextExecutor) (TagSlice, error) {
	var o []*Tag

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "mysql: failed to assign all query results to Tag slice")
	}

	if len(tagAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Tag records in the query.
func (q tagQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to count Tag rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tagQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "mysql: failed to check if Tag exists")
	}

	return count > 0, nil
}

// Tags retrieves all the records using an executor.
func Tags(mods ...qm.QueryMod) tagQuery {
	mods = append(mods, qm.From("`Tag`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`Tag`.*"})
	}

	return tagQuery{q}
}

// FindTag retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTag(ctx context.Context, exec boil.ContextExecutor, tagId int, selectCols ...string) (*Tag, error) {
	tagObj := &Tag{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `Tag` where `tagId`=?", sel,
	)

	q := queries.Raw(query, tagId)

	err := q.Bind(ctx, exec, tagObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mysql: unable to select from Tag")
	}

	if err = tagObj.doAfterSelectHooks(ctx, exec); err != nil {
		return tagObj, err
	}

	return tagObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Tag) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("mysql: no Tag provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tagColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tagInsertCacheMut.RLock()
	cache, cached := tagInsertCache[key]
	tagInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tagAllColumns,
			tagColumnsWithDefault,
			tagColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tagType, tagMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `Tag` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `Tag` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `Tag` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tagPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "mysql: unable to insert into Tag")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.TagId = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == tagMapping["tagId"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.TagId,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to populate default values for Tag")
	}

CacheNoHooks:
	if !cached {
		tagInsertCacheMut.Lock()
		tagInsertCache[key] = cache
		tagInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Tag.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Tag) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tagUpdateCacheMut.RLock()
	cache, cached := tagUpdateCache[key]
	tagUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tagAllColumns,
			tagPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("mysql: unable to update Tag, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `Tag` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tagPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, append(wl, tagPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "mysql: unable to update Tag row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by update for Tag")
	}

	if !cached {
		tagUpdateCacheMut.Lock()
		tagUpdateCache[key] = cache
		tagUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tagQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to update all for Tag")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to retrieve rows affected for Tag")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TagSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("mysql: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `Tag` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tagPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to update all in tag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to retrieve rows affected all in update all tag")
	}
	return rowsAff, nil
}

var mySQLTagUniqueColumns = []string{
	"tagId",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Tag) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("mysql: no Tag provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tagColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTagUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	tagUpsertCacheMut.RLock()
	cache, cached := tagUpsertCache[key]
	tagUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			tagAllColumns,
			tagColumnsWithDefault,
			tagColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			tagAllColumns,
			tagPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("mysql: unable to upsert Tag, could not build update column list")
		}

		ret := strmangle.SetComplement(tagAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`Tag`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `Tag` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tagType, tagMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tagType, tagMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "mysql: unable to upsert for Tag")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.TagId = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == tagMapping["tagId"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tagType, tagMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to retrieve unique values for Tag")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to populate default values for Tag")
	}

CacheNoHooks:
	if !cached {
		tagUpsertCacheMut.Lock()
		tagUpsertCache[key] = cache
		tagUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Tag record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Tag) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("mysql: no Tag provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tagPrimaryKeyMapping)
	sql := "DELETE FROM `Tag` WHERE `tagId`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete from Tag")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by delete for Tag")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tagQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("mysql: no tagQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete all from Tag")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by deleteall for Tag")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TagSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(tagBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `Tag` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tagPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete all from tag slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by deleteall for Tag")
	}

	if len(tagAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Tag) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTag(ctx, exec, o.TagId)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TagSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TagSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tagPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `Tag`.* FROM `Tag` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tagPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to reload all in TagSlice")
	}

	*o = slice

	return nil
}

// TagExists checks if the Tag row exists.
func TagExists(ctx context.Context, exec boil.ContextExecutor, tagId int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `Tag` where `tagId`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, tagId)
	}
	row := exec.QueryRowContext(ctx, sql, tagId)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "mysql: unable to check if Tag exists")
	}

	return exists, nil
}

// Exists checks if the Tag row exists.
func (o *Tag) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return TagExists(ctx, exec, o.TagId)
}
