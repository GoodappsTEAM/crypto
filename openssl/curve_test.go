// +build experimental

package openssl

import (
	"gopkg.in/dedis/crypto.v0/test"
	"testing"
)

var testSuite = NewAES128SHA256P256()
var groupBench = test.NewGroupBench(testSuite)

func TestP224(t *testing.T) { test.TestGroup(new(curve).InitP224()) }
func TestP256(t *testing.T) { test.TestGroup(new(curve).InitP256()) }
func TestP384(t *testing.T) { test.TestGroup(new(curve).InitP384()) }
func TestP521(t *testing.T) { test.TestGroup(new(curve).InitP521()) }

func TestSuite(t *testing.T) { test.TestSuite(testSuite) }

func BenchmarkSecretAdd(b *testing.B)    { groupBench.SecretAdd(b.N) }
func BenchmarkSecretSub(b *testing.B)    { groupBench.SecretSub(b.N) }
func BenchmarkSecretNeg(b *testing.B)    { groupBench.SecretNeg(b.N) }
func BenchmarkSecretMul(b *testing.B)    { groupBench.SecretMul(b.N) }
func BenchmarkSecretDiv(b *testing.B)    { groupBench.SecretDiv(b.N) }
func BenchmarkSecretInv(b *testing.B)    { groupBench.SecretInv(b.N) }
func BenchmarkSecretPick(b *testing.B)   { groupBench.SecretPick(b.N) }
func BenchmarkSecretEncode(b *testing.B) { groupBench.SecretEncode(b.N) }
func BenchmarkSecretDecode(b *testing.B) { groupBench.SecretDecode(b.N) }

func BenchmarkPointAdd(b *testing.B)     { groupBench.PointAdd(b.N) }
func BenchmarkPointSub(b *testing.B)     { groupBench.PointSub(b.N) }
func BenchmarkPointNeg(b *testing.B)     { groupBench.PointNeg(b.N) }
func BenchmarkPointMul(b *testing.B)     { groupBench.PointMul(b.N) }
func BenchmarkPointBaseMul(b *testing.B) { groupBench.PointBaseMul(b.N) }
func BenchmarkPointPick(b *testing.B)    { groupBench.PointPick(b.N) }
func BenchmarkPointEncode(b *testing.B)  { groupBench.PointEncode(b.N) }
func BenchmarkPointDecode(b *testing.B)  { groupBench.PointDecode(b.N) }
