## chinese-segment-go


`chinese-segment-go` 仅对 `gojieba` 和 `gse` 的简单封装，解决 `gojieba` 依赖问题[1] ，用于进行性能测试和效果对比。

[1] `gojieba` 使用 `cgo`，编译时会遇到依赖问题

## 压力测试

```shell
go test ./segment -bench=Benchmark -benchtime=1000x
```

压测结果

```text
goos: darwin
goarch: amd64
pkg: github.com/forcemeter/chinese-segment-go/segment
cpu: Intel(R) Core(TM) i5-1038NG7 CPU @ 2.00GHz
BenchmarkGse-8             10000           6955941 ns/op
BenchmarkJieba-8           10000           2721562 ns/op
PASS
ok      github.com/forcemeter/chinese-segment-go/segment        103.699s
```

## 测试结果

1. `gojieba` 的性能是 `gse` 的 3 倍左右
2. `gse` 的准确率，词库有效率高于 `gojieba`

## 选型建议

从复杂度、准确率度上，推荐使用 `gse`。