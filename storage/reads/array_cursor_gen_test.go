// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: array_cursor_test.gen.go.tmpl

package reads

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/influxdata/influxdb/v2/storage/reads/datatypes"
	"github.com/influxdata/influxdb/v2/tsdb/cursors"
)

type MockFloatArrayCursor struct {
	CloseFunc func()
	ErrFunc   func() error
	StatsFunc func() cursors.CursorStats
	NextFunc  func() *cursors.FloatArray
}

func (c *MockFloatArrayCursor) Close()                     { c.CloseFunc() }
func (c *MockFloatArrayCursor) Err() error                 { return c.ErrFunc() }
func (c *MockFloatArrayCursor) Stats() cursors.CursorStats { return c.StatsFunc() }
func (c *MockFloatArrayCursor) Next() *cursors.FloatArray  { return c.NextFunc() }

func TestNewAggregateArrayCursor_Float(t *testing.T) {

	t.Run("Count", func(t *testing.T) {
		want := &floatWindowCountArrayCursor{
			FloatArrayCursor: &MockFloatArrayCursor{},
			res:              cursors.NewIntegerArrayLen(1),
			tmp:              &cursors.FloatArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeCount,
		}

		got := newAggregateArrayCursor(context.Background(), agg, &MockFloatArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(floatWindowCountArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

	t.Run("Sum", func(t *testing.T) {
		want := &floatWindowSumArrayCursor{
			FloatArrayCursor: &MockFloatArrayCursor{},
			res:              cursors.NewFloatArrayLen(1),
			tmp:              &cursors.FloatArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeSum,
		}

		got := newAggregateArrayCursor(context.Background(), agg, &MockFloatArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(floatWindowSumArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}

func TestNewWindowAggregateArrayCursor_Float(t *testing.T) {

	t.Run("Count", func(t *testing.T) {
		want := &floatWindowCountArrayCursor{
			FloatArrayCursor: &MockFloatArrayCursor{},
			every:            int64(time.Hour),
			res:              cursors.NewIntegerArrayLen(MaxPointsPerBlock),
			tmp:              &cursors.FloatArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeCount,
		}

		got := newWindowAggregateArrayCursor(context.Background(), agg, int64(time.Hour), 0, &MockFloatArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(floatWindowCountArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

	t.Run("Sum", func(t *testing.T) {
		want := &floatWindowSumArrayCursor{
			FloatArrayCursor: &MockFloatArrayCursor{},
			every:            int64(time.Hour),
			res:              cursors.NewFloatArrayLen(MaxPointsPerBlock),
			tmp:              &cursors.FloatArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeSum,
		}

		got := newWindowAggregateArrayCursor(context.Background(), agg, int64(time.Hour), 0, &MockFloatArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(floatWindowSumArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}

type MockIntegerArrayCursor struct {
	CloseFunc func()
	ErrFunc   func() error
	StatsFunc func() cursors.CursorStats
	NextFunc  func() *cursors.IntegerArray
}

func (c *MockIntegerArrayCursor) Close()                      { c.CloseFunc() }
func (c *MockIntegerArrayCursor) Err() error                  { return c.ErrFunc() }
func (c *MockIntegerArrayCursor) Stats() cursors.CursorStats  { return c.StatsFunc() }
func (c *MockIntegerArrayCursor) Next() *cursors.IntegerArray { return c.NextFunc() }

func TestNewAggregateArrayCursor_Integer(t *testing.T) {

	t.Run("Count", func(t *testing.T) {
		want := &integerWindowCountArrayCursor{
			IntegerArrayCursor: &MockIntegerArrayCursor{},
			res:                cursors.NewIntegerArrayLen(1),
			tmp:                &cursors.IntegerArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeCount,
		}

		got := newAggregateArrayCursor(context.Background(), agg, &MockIntegerArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(integerWindowCountArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

	t.Run("Sum", func(t *testing.T) {
		want := &integerWindowSumArrayCursor{
			IntegerArrayCursor: &MockIntegerArrayCursor{},
			res:                cursors.NewIntegerArrayLen(1),
			tmp:                &cursors.IntegerArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeSum,
		}

		got := newAggregateArrayCursor(context.Background(), agg, &MockIntegerArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(integerWindowSumArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}

func TestNewWindowAggregateArrayCursor_Integer(t *testing.T) {

	t.Run("Count", func(t *testing.T) {
		want := &integerWindowCountArrayCursor{
			IntegerArrayCursor: &MockIntegerArrayCursor{},
			every:              int64(time.Hour),
			res:                cursors.NewIntegerArrayLen(MaxPointsPerBlock),
			tmp:                &cursors.IntegerArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeCount,
		}

		got := newWindowAggregateArrayCursor(context.Background(), agg, int64(time.Hour), 0, &MockIntegerArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(integerWindowCountArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

	t.Run("Sum", func(t *testing.T) {
		want := &integerWindowSumArrayCursor{
			IntegerArrayCursor: &MockIntegerArrayCursor{},
			every:              int64(time.Hour),
			res:                cursors.NewIntegerArrayLen(MaxPointsPerBlock),
			tmp:                &cursors.IntegerArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeSum,
		}

		got := newWindowAggregateArrayCursor(context.Background(), agg, int64(time.Hour), 0, &MockIntegerArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(integerWindowSumArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}

type MockUnsignedArrayCursor struct {
	CloseFunc func()
	ErrFunc   func() error
	StatsFunc func() cursors.CursorStats
	NextFunc  func() *cursors.UnsignedArray
}

func (c *MockUnsignedArrayCursor) Close()                       { c.CloseFunc() }
func (c *MockUnsignedArrayCursor) Err() error                   { return c.ErrFunc() }
func (c *MockUnsignedArrayCursor) Stats() cursors.CursorStats   { return c.StatsFunc() }
func (c *MockUnsignedArrayCursor) Next() *cursors.UnsignedArray { return c.NextFunc() }

func TestNewAggregateArrayCursor_Unsigned(t *testing.T) {

	t.Run("Count", func(t *testing.T) {
		want := &unsignedWindowCountArrayCursor{
			UnsignedArrayCursor: &MockUnsignedArrayCursor{},
			res:                 cursors.NewIntegerArrayLen(1),
			tmp:                 &cursors.UnsignedArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeCount,
		}

		got := newAggregateArrayCursor(context.Background(), agg, &MockUnsignedArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(unsignedWindowCountArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

	t.Run("Sum", func(t *testing.T) {
		want := &unsignedWindowSumArrayCursor{
			UnsignedArrayCursor: &MockUnsignedArrayCursor{},
			res:                 cursors.NewUnsignedArrayLen(1),
			tmp:                 &cursors.UnsignedArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeSum,
		}

		got := newAggregateArrayCursor(context.Background(), agg, &MockUnsignedArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(unsignedWindowSumArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}

func TestNewWindowAggregateArrayCursor_Unsigned(t *testing.T) {

	t.Run("Count", func(t *testing.T) {
		want := &unsignedWindowCountArrayCursor{
			UnsignedArrayCursor: &MockUnsignedArrayCursor{},
			every:               int64(time.Hour),
			res:                 cursors.NewIntegerArrayLen(MaxPointsPerBlock),
			tmp:                 &cursors.UnsignedArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeCount,
		}

		got := newWindowAggregateArrayCursor(context.Background(), agg, int64(time.Hour), 0, &MockUnsignedArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(unsignedWindowCountArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

	t.Run("Sum", func(t *testing.T) {
		want := &unsignedWindowSumArrayCursor{
			UnsignedArrayCursor: &MockUnsignedArrayCursor{},
			every:               int64(time.Hour),
			res:                 cursors.NewUnsignedArrayLen(MaxPointsPerBlock),
			tmp:                 &cursors.UnsignedArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeSum,
		}

		got := newWindowAggregateArrayCursor(context.Background(), agg, int64(time.Hour), 0, &MockUnsignedArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(unsignedWindowSumArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}

type MockStringArrayCursor struct {
	CloseFunc func()
	ErrFunc   func() error
	StatsFunc func() cursors.CursorStats
	NextFunc  func() *cursors.StringArray
}

func (c *MockStringArrayCursor) Close()                     { c.CloseFunc() }
func (c *MockStringArrayCursor) Err() error                 { return c.ErrFunc() }
func (c *MockStringArrayCursor) Stats() cursors.CursorStats { return c.StatsFunc() }
func (c *MockStringArrayCursor) Next() *cursors.StringArray { return c.NextFunc() }

func TestNewAggregateArrayCursor_String(t *testing.T) {

	t.Run("Count", func(t *testing.T) {
		want := &stringWindowCountArrayCursor{
			StringArrayCursor: &MockStringArrayCursor{},
			res:               cursors.NewIntegerArrayLen(1),
			tmp:               &cursors.StringArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeCount,
		}

		got := newAggregateArrayCursor(context.Background(), agg, &MockStringArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(stringWindowCountArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}

func TestNewWindowAggregateArrayCursor_String(t *testing.T) {

	t.Run("Count", func(t *testing.T) {
		want := &stringWindowCountArrayCursor{
			StringArrayCursor: &MockStringArrayCursor{},
			every:             int64(time.Hour),
			res:               cursors.NewIntegerArrayLen(MaxPointsPerBlock),
			tmp:               &cursors.StringArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeCount,
		}

		got := newWindowAggregateArrayCursor(context.Background(), agg, int64(time.Hour), 0, &MockStringArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(stringWindowCountArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}

type MockBooleanArrayCursor struct {
	CloseFunc func()
	ErrFunc   func() error
	StatsFunc func() cursors.CursorStats
	NextFunc  func() *cursors.BooleanArray
}

func (c *MockBooleanArrayCursor) Close()                      { c.CloseFunc() }
func (c *MockBooleanArrayCursor) Err() error                  { return c.ErrFunc() }
func (c *MockBooleanArrayCursor) Stats() cursors.CursorStats  { return c.StatsFunc() }
func (c *MockBooleanArrayCursor) Next() *cursors.BooleanArray { return c.NextFunc() }

func TestNewAggregateArrayCursor_Boolean(t *testing.T) {

	t.Run("Count", func(t *testing.T) {
		want := &booleanWindowCountArrayCursor{
			BooleanArrayCursor: &MockBooleanArrayCursor{},
			res:                cursors.NewIntegerArrayLen(1),
			tmp:                &cursors.BooleanArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeCount,
		}

		got := newAggregateArrayCursor(context.Background(), agg, &MockBooleanArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(booleanWindowCountArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}

func TestNewWindowAggregateArrayCursor_Boolean(t *testing.T) {

	t.Run("Count", func(t *testing.T) {
		want := &booleanWindowCountArrayCursor{
			BooleanArrayCursor: &MockBooleanArrayCursor{},
			every:              int64(time.Hour),
			res:                cursors.NewIntegerArrayLen(MaxPointsPerBlock),
			tmp:                &cursors.BooleanArray{},
		}

		agg := &datatypes.Aggregate{
			Type: datatypes.AggregateTypeCount,
		}

		got := newWindowAggregateArrayCursor(context.Background(), agg, int64(time.Hour), 0, &MockBooleanArrayCursor{})

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(booleanWindowCountArrayCursor{})); diff != "" {
			t.Fatalf("did not get expected cursor; -got/+want:\n%v", diff)
		}
	})

}
