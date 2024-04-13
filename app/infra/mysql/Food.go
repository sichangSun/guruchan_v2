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

// Food is an object representing the database table.
type Food struct {
	FoodId         int         `boil:"foodId" json:"foodId" toml:"foodId" yaml:"foodId"`
	UserId         int         `boil:"userId" json:"userId" toml:"userId" yaml:"userId"`
	RestaurantName string      `boil:"restaurantName" json:"restaurantName" toml:"restaurantName" yaml:"restaurantName"`
	FoodName       null.String `boil:"foodName" json:"foodName,omitempty" toml:"foodName" yaml:"foodName,omitempty"`
	Address        null.String `boil:"address" json:"address,omitempty" toml:"address" yaml:"address,omitempty"`
	FullAddress    null.String `boil:"fullAddress" json:"fullAddress,omitempty" toml:"fullAddress" yaml:"fullAddress,omitempty"`
	IsLikeflag     int         `boil:"isLikeflag" json:"isLikeflag" toml:"isLikeflag" yaml:"isLikeflag"`
	// 0未分类
	Typecode    int       `boil:"typecode" json:"typecode" toml:"typecode" yaml:"typecode"`
	FAddTime    null.Time `boil:"f_addTime" json:"f_addTime,omitempty" toml:"f_addTime" yaml:"f_addTime,omitempty"`
	FUpdataTime null.Time `boil:"f_updataTime" json:"f_updataTime,omitempty" toml:"f_updataTime" yaml:"f_updataTime,omitempty"`
	FirstTime   null.Time `boil:"firstTime" json:"firstTime,omitempty" toml:"firstTime" yaml:"firstTime,omitempty"`
	// 0未知,1无,2有
	TestedFlag int         `boil:"testedFlag" json:"testedFlag" toml:"testedFlag" yaml:"testedFlag"`
	FoodImg    null.String `boil:"foodImg" json:"foodImg,omitempty" toml:"foodImg" yaml:"foodImg,omitempty"`
	// 1为删除
	IsDel int `boil:"isDel" json:"isDel" toml:"isDel" yaml:"isDel"`

	R *foodR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L foodL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FoodColumns = struct {
	FoodId         string
	UserId         string
	RestaurantName string
	FoodName       string
	Address        string
	FullAddress    string
	IsLikeflag     string
	Typecode       string
	FAddTime       string
	FUpdataTime    string
	FirstTime      string
	TestedFlag     string
	FoodImg        string
	IsDel          string
}{
	FoodId:         "foodId",
	UserId:         "userId",
	RestaurantName: "restaurantName",
	FoodName:       "foodName",
	Address:        "address",
	FullAddress:    "fullAddress",
	IsLikeflag:     "isLikeflag",
	Typecode:       "typecode",
	FAddTime:       "f_addTime",
	FUpdataTime:    "f_updataTime",
	FirstTime:      "firstTime",
	TestedFlag:     "testedFlag",
	FoodImg:        "foodImg",
	IsDel:          "isDel",
}

var FoodTableColumns = struct {
	FoodId         string
	UserId         string
	RestaurantName string
	FoodName       string
	Address        string
	FullAddress    string
	IsLikeflag     string
	Typecode       string
	FAddTime       string
	FUpdataTime    string
	FirstTime      string
	TestedFlag     string
	FoodImg        string
	IsDel          string
}{
	FoodId:         "Food.foodId",
	UserId:         "Food.userId",
	RestaurantName: "Food.restaurantName",
	FoodName:       "Food.foodName",
	Address:        "Food.address",
	FullAddress:    "Food.fullAddress",
	IsLikeflag:     "Food.isLikeflag",
	Typecode:       "Food.typecode",
	FAddTime:       "Food.f_addTime",
	FUpdataTime:    "Food.f_updataTime",
	FirstTime:      "Food.firstTime",
	TestedFlag:     "Food.testedFlag",
	FoodImg:        "Food.foodImg",
	IsDel:          "Food.isDel",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod   { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod  { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

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
func (w whereHelpernull_String) LIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" LIKE ?", x)
}
func (w whereHelpernull_String) NLIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT LIKE ?", x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

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

var FoodWhere = struct {
	FoodId         whereHelperint
	UserId         whereHelperint
	RestaurantName whereHelperstring
	FoodName       whereHelpernull_String
	Address        whereHelpernull_String
	FullAddress    whereHelpernull_String
	IsLikeflag     whereHelperint
	Typecode       whereHelperint
	FAddTime       whereHelpernull_Time
	FUpdataTime    whereHelpernull_Time
	FirstTime      whereHelpernull_Time
	TestedFlag     whereHelperint
	FoodImg        whereHelpernull_String
	IsDel          whereHelperint
}{
	FoodId:         whereHelperint{field: "`Food`.`foodId`"},
	UserId:         whereHelperint{field: "`Food`.`userId`"},
	RestaurantName: whereHelperstring{field: "`Food`.`restaurantName`"},
	FoodName:       whereHelpernull_String{field: "`Food`.`foodName`"},
	Address:        whereHelpernull_String{field: "`Food`.`address`"},
	FullAddress:    whereHelpernull_String{field: "`Food`.`fullAddress`"},
	IsLikeflag:     whereHelperint{field: "`Food`.`isLikeflag`"},
	Typecode:       whereHelperint{field: "`Food`.`typecode`"},
	FAddTime:       whereHelpernull_Time{field: "`Food`.`f_addTime`"},
	FUpdataTime:    whereHelpernull_Time{field: "`Food`.`f_updataTime`"},
	FirstTime:      whereHelpernull_Time{field: "`Food`.`firstTime`"},
	TestedFlag:     whereHelperint{field: "`Food`.`testedFlag`"},
	FoodImg:        whereHelpernull_String{field: "`Food`.`foodImg`"},
	IsDel:          whereHelperint{field: "`Food`.`isDel`"},
}

// FoodRels is where relationship names are stored.
var FoodRels = struct {
}{}

// foodR is where relationships are stored.
type foodR struct {
}

// NewStruct creates a new relationship struct
func (*foodR) NewStruct() *foodR {
	return &foodR{}
}

// foodL is where Load methods for each relationship are stored.
type foodL struct{}

var (
	foodAllColumns            = []string{"foodId", "userId", "restaurantName", "foodName", "address", "fullAddress", "isLikeflag", "typecode", "f_addTime", "f_updataTime", "firstTime", "testedFlag", "foodImg", "isDel"}
	foodColumnsWithoutDefault = []string{"userId", "restaurantName", "foodName", "address", "fullAddress", "f_addTime", "f_updataTime", "firstTime", "foodImg"}
	foodColumnsWithDefault    = []string{"foodId", "isLikeflag", "typecode", "testedFlag", "isDel"}
	foodPrimaryKeyColumns     = []string{"foodId"}
	foodGeneratedColumns      = []string{}
)

type (
	// FoodSlice is an alias for a slice of pointers to Food.
	// This should almost always be used instead of []Food.
	FoodSlice []*Food
	// FoodHook is the signature for custom Food hook methods
	FoodHook func(context.Context, boil.ContextExecutor, *Food) error

	foodQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	foodType                 = reflect.TypeOf(&Food{})
	foodMapping              = queries.MakeStructMapping(foodType)
	foodPrimaryKeyMapping, _ = queries.BindMapping(foodType, foodMapping, foodPrimaryKeyColumns)
	foodInsertCacheMut       sync.RWMutex
	foodInsertCache          = make(map[string]insertCache)
	foodUpdateCacheMut       sync.RWMutex
	foodUpdateCache          = make(map[string]updateCache)
	foodUpsertCacheMut       sync.RWMutex
	foodUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var foodAfterSelectMu sync.Mutex
var foodAfterSelectHooks []FoodHook

var foodBeforeInsertMu sync.Mutex
var foodBeforeInsertHooks []FoodHook
var foodAfterInsertMu sync.Mutex
var foodAfterInsertHooks []FoodHook

var foodBeforeUpdateMu sync.Mutex
var foodBeforeUpdateHooks []FoodHook
var foodAfterUpdateMu sync.Mutex
var foodAfterUpdateHooks []FoodHook

var foodBeforeDeleteMu sync.Mutex
var foodBeforeDeleteHooks []FoodHook
var foodAfterDeleteMu sync.Mutex
var foodAfterDeleteHooks []FoodHook

var foodBeforeUpsertMu sync.Mutex
var foodBeforeUpsertHooks []FoodHook
var foodAfterUpsertMu sync.Mutex
var foodAfterUpsertHooks []FoodHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Food) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range foodAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Food) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range foodBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Food) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range foodAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Food) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range foodBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Food) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range foodAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Food) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range foodBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Food) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range foodAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Food) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range foodBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Food) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range foodAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFoodHook registers your hook function for all future operations.
func AddFoodHook(hookPoint boil.HookPoint, foodHook FoodHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		foodAfterSelectMu.Lock()
		foodAfterSelectHooks = append(foodAfterSelectHooks, foodHook)
		foodAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		foodBeforeInsertMu.Lock()
		foodBeforeInsertHooks = append(foodBeforeInsertHooks, foodHook)
		foodBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		foodAfterInsertMu.Lock()
		foodAfterInsertHooks = append(foodAfterInsertHooks, foodHook)
		foodAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		foodBeforeUpdateMu.Lock()
		foodBeforeUpdateHooks = append(foodBeforeUpdateHooks, foodHook)
		foodBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		foodAfterUpdateMu.Lock()
		foodAfterUpdateHooks = append(foodAfterUpdateHooks, foodHook)
		foodAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		foodBeforeDeleteMu.Lock()
		foodBeforeDeleteHooks = append(foodBeforeDeleteHooks, foodHook)
		foodBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		foodAfterDeleteMu.Lock()
		foodAfterDeleteHooks = append(foodAfterDeleteHooks, foodHook)
		foodAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		foodBeforeUpsertMu.Lock()
		foodBeforeUpsertHooks = append(foodBeforeUpsertHooks, foodHook)
		foodBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		foodAfterUpsertMu.Lock()
		foodAfterUpsertHooks = append(foodAfterUpsertHooks, foodHook)
		foodAfterUpsertMu.Unlock()
	}
}

// One returns a single food record from the query.
func (q foodQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Food, error) {
	o := &Food{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mysql: failed to execute a one query for Food")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Food records from the query.
func (q foodQuery) All(ctx context.Context, exec boil.ContextExecutor) (FoodSlice, error) {
	var o []*Food

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "mysql: failed to assign all query results to Food slice")
	}

	if len(foodAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Food records in the query.
func (q foodQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to count Food rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q foodQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "mysql: failed to check if Food exists")
	}

	return count > 0, nil
}

// Foods retrieves all the records using an executor.
func Foods(mods ...qm.QueryMod) foodQuery {
	mods = append(mods, qm.From("`Food`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`Food`.*"})
	}

	return foodQuery{q}
}

// FindFood retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFood(ctx context.Context, exec boil.ContextExecutor, foodId int, selectCols ...string) (*Food, error) {
	foodObj := &Food{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `Food` where `foodId`=?", sel,
	)

	q := queries.Raw(query, foodId)

	err := q.Bind(ctx, exec, foodObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mysql: unable to select from Food")
	}

	if err = foodObj.doAfterSelectHooks(ctx, exec); err != nil {
		return foodObj, err
	}

	return foodObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Food) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("mysql: no Food provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(foodColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	foodInsertCacheMut.RLock()
	cache, cached := foodInsertCache[key]
	foodInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			foodAllColumns,
			foodColumnsWithDefault,
			foodColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(foodType, foodMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(foodType, foodMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `Food` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `Food` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `Food` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, foodPrimaryKeyColumns))
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
		return errors.Wrap(err, "mysql: unable to insert into Food")
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

	o.FoodId = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == foodMapping["foodId"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.FoodId,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to populate default values for Food")
	}

CacheNoHooks:
	if !cached {
		foodInsertCacheMut.Lock()
		foodInsertCache[key] = cache
		foodInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Food.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Food) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	foodUpdateCacheMut.RLock()
	cache, cached := foodUpdateCache[key]
	foodUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			foodAllColumns,
			foodPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("mysql: unable to update Food, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `Food` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, foodPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(foodType, foodMapping, append(wl, foodPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "mysql: unable to update Food row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by update for Food")
	}

	if !cached {
		foodUpdateCacheMut.Lock()
		foodUpdateCache[key] = cache
		foodUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q foodQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to update all for Food")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to retrieve rows affected for Food")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FoodSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), foodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `Food` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, foodPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to update all in food slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to retrieve rows affected all in update all food")
	}
	return rowsAff, nil
}

var mySQLFoodUniqueColumns = []string{
	"foodId",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Food) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("mysql: no Food provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(foodColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLFoodUniqueColumns, o)

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

	foodUpsertCacheMut.RLock()
	cache, cached := foodUpsertCache[key]
	foodUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			foodAllColumns,
			foodColumnsWithDefault,
			foodColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			foodAllColumns,
			foodPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("mysql: unable to upsert Food, could not build update column list")
		}

		ret := strmangle.SetComplement(foodAllColumns, strmangle.SetIntersect(insert, update))

		cache.query = buildUpsertQueryMySQL(dialect, "`Food`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `Food` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(foodType, foodMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(foodType, foodMapping, ret)
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
		return errors.Wrap(err, "mysql: unable to upsert for Food")
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

	o.FoodId = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == foodMapping["foodId"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(foodType, foodMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to retrieve unique values for Food")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to populate default values for Food")
	}

CacheNoHooks:
	if !cached {
		foodUpsertCacheMut.Lock()
		foodUpsertCache[key] = cache
		foodUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Food record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Food) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("mysql: no Food provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), foodPrimaryKeyMapping)
	sql := "DELETE FROM `Food` WHERE `foodId`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete from Food")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by delete for Food")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q foodQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("mysql: no foodQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete all from Food")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by deleteall for Food")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FoodSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(foodBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), foodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `Food` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, foodPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete all from food slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by deleteall for Food")
	}

	if len(foodAfterDeleteHooks) != 0 {
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
func (o *Food) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFood(ctx, exec, o.FoodId)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FoodSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FoodSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), foodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `Food`.* FROM `Food` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, foodPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to reload all in FoodSlice")
	}

	*o = slice

	return nil
}

// FoodExists checks if the Food row exists.
func FoodExists(ctx context.Context, exec boil.ContextExecutor, foodId int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `Food` where `foodId`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, foodId)
	}
	row := exec.QueryRowContext(ctx, sql, foodId)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "mysql: unable to check if Food exists")
	}

	return exists, nil
}

// Exists checks if the Food row exists.
func (o *Food) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return FoodExists(ctx, exec, o.FoodId)
}
