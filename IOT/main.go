/*
 * @Descripttion:
 * @Author: zenghua.wang
 * @Date: 2020-06-13 11:06:50
 * @LastEditors: zenghua.wang
 * @LastEditTime: 2022-12-08 14:24:37
 */
package main

import (
	"golang-knowledge/IOT/mtk"
	"testing"
)

const FM24XX_I2C_ADDR = 0x50

func TestRead(t *testing.T) {

	fm := mtk.NewFM24cxxFram(0, FM24XX_I2C_ADDR, 3, 10)
	if err := fm.Open(); err != nil {
		t.Errorf("Open Fm failed %v", err)
		return
	}
	if _, err := fm.Write(0, []byte{1, 2, 3, 4, 5}); err != nil {
		t.Errorf("Write Fm failed %v", err)
		return
	}
	buf := make([]byte, 5)
	if _, err := fm.Read(0, buf); err != nil {
		t.Errorf("Read Fm failed %v", err)
		return
	}
	t.Logf("Read Data=%v", buf)

}
