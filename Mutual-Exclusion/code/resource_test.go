package main

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_resource_occupyAndRelease(t *testing.T) {
	// 避免 debugprint 输出
	temp := needDebug
	needDebug = false
	defer func() { needDebug = temp }()
	//
	ast := assert.New(t)
	//
	p := 0
	ts := newTimestamp(0, p)
	r := new(resource)
	r.Occupy(ts)
	//
	ast.Equal(ts, r.occupiedBy)
	ast.Equal(ts, r.timestamps[0])

}

func Test_resource_occupy_occupyInvalidResource(t *testing.T) {
	// 避免 debugprint 输出
	temp := needDebug
	needDebug = false
	defer func() { needDebug = temp }()
	//
	ast := assert.New(t)
	//
	p0 := 0
	p1 := 1
	ts0 := newTimestamp(0, p0)
	ts1 := newTimestamp(1, p1)
	r := new(resource)
	r.Occupy(ts0)
	//
	expected := fmt.Sprintf("资源正在被 %s 占据，%s 却想获取资源。", ts0, ts1)
	ast.PanicsWithValue(expected, func() { r.Occupy(ts1) })
}

func Test_resource_report(t *testing.T) {
	// 避免 debugprint 输出
	temp := needDebug
	needDebug = false
	defer func() { needDebug = temp }()
	//
	ast := assert.New(t)
	//
	p := 0
	ts0 := newTimestamp(0, p)
	ts1 := newTimestamp(1, p)
	r := new(resource)
	r.wg.Add(2)
	r.Occupy(ts0)
	r.Release(ts0)
	r.Occupy(ts1)
	r.Release(ts1)
	now := time.Now()
	r.times[0] = now
	r.times[1] = now.Add(100 * time.Second)
	r.times[2] = now.Add(200 * time.Second)
	r.times[3] = now.Add(400 * time.Second)
	//
	report := r.report()
	ast.True(strings.Contains(report, "75.00%"), report)
	//
	ast.Equal(4, len(r.times), "资源被占用了 2 次，但是 r.times 的长度不等于 4")
}

func Test_resource_timestamps(t *testing.T) {
	// 避免 debugprint 输出
	temp := needDebug
	needDebug = false
	defer func() { needDebug = temp }()
	//
	ast := assert.New(t)
	//
	time := 0
	p := 0
	times := 100
	r := new(resource)
	r.wg.Add(times)
	//
	for i := 0; i < times; i++ {
		if i%2 == 0 {
			time++
		} else {
			p++
		}
		ts := newTimestamp(time, p)
		r.Occupy(ts)
		r.Release(ts)
	}
	//
	expected := times * 2
	actual := len(r.times)
	ast.Equal(expected, actual)
}
