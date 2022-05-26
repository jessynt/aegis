# Coding Guide

Aegis 开发指南

- [Style Guideline](#style-guideline)
- [Review Guideline](#review-guideline)

## Style Guideline

**项目代码风格是会随着项目发展而变化的，所以遇到任何问题，just ask**

### Golang Style Guideline

Golang 开发内容以以下几个文档作为基础：

- [Effective Go][effective_go]
- [CodeReviewComments][code_review_comments]
- [Style guideline for Go packages](rakyll_style_packages)

[effective_go]: https://golang.org/doc/effective_go.html
[code_review_comments]: https://github.com/golang/go/wiki/CodeReviewComments
[rakyll_style_packages]: https://rakyll.org/style-packages/

但是，有以下规则需要进行修改和补充：

#### 所有命名尽可能避免使用复数形式，使用对应单词的英语单数

<details>
避免因为英语复数不规则带来的困扰
</details>

#### 包名允许使用下划线

<details>
提高可读性
</details>

#### 包、模块、struct 结构文档不要使用对应对象开头

<details>
Aegis 中为了明确表达意图，请使用中文进行注释
</details>

#### 每行代码长度尽量不要超过 120 个字符

<details>
避免使用过长的名称
</details>

#### 当 id 表示 `identity` 时，camel case 请使用 `fooId`

<details>
保持代码风格兼容性
</details>

#### import 根据 标准库, 本地库, 第三方库进行分组, 最好使用 `goimports` 排序

<details>
统一的排序方便 review 看 diff
</details>

## Review Guideline

- 所有 golang 代码 review 之前请先使用 [gofmt][] 进行格式化，review 过程中应该
  将关注点放在功能实现上

[gofmt]: https://blog.golang.org/go-fmt-your-code

## 本地开发

### 开发惯例

#### git 提交说明

格式请参照 [angular git commit guidlines][angular_git_commit]，注意除 type/scope
外，提交内容请尽量使用中文，例：

```
docs(readme): 更新 README 中的 git 提交说明
```

[angular_git_commit]: https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#commits


