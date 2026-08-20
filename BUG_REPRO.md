# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

服务站按时间导出事项时把 from 填成 8 月 21 日、to 填成 8 月 20 日，接口却像正常查询一样返回空文件，调用方无法知道范围写反了。请处理时间范围校验，拒绝 from 晚于 to，同时保留只填一侧边界的合法查询；生产代码修好即可，测试文件不要修改。

## 含 Bug 版本

- 仓库：11DingKing/rider-rights-task-29
- 仓库地址：https://github.com/11DingKing/rider-rights-task-29.git
- parent SHA：bb474e4ee7b98da537e7017452a94a0a3fcfb729

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/rider-rights-task-29.git bug-repro
cd bug-repro
git checkout --detach bb474e4ee7b98da537e7017452a94a0a3fcfb729
go test ./internal/domain -run "^TestExportRejectsReversedRange$" -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestExportRejectsReversedRange$" -count=1
--- FAIL: TestExportRejectsReversedRange (0.00s)
    task29_test.go:12: reversed export range was accepted
FAIL
FAIL	riderguard/internal/domain	0.040s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestExportRejectsReversedRange$" -count=1
--- FAIL: TestExportRejectsReversedRange (0.00s)
    task29_test.go:12: reversed export range was accepted
FAIL
FAIL	riderguard/internal/domain	0.001s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

修复后，from 和 to 都存在且 from 晚于 to 时必须在查询前返回校验错误；仅提供 from、仅提供 to 或合法闭区间仍应进入导出查询。定向测试、相关包测试及全量回归必须通过，不得删除、跳过或削弱测试。
