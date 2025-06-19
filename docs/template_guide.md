# AdminLTE 模板快速开发指南

## 1. 基本模板结构
每个页面都应包含以下四个基本模板：
```html
{{ template "layout/header.html" . }}  <!-- 头部，包含基础CSS和meta信息 -->
{{ template "layout/nav.html" . }}     <!-- 导航栏和侧边栏 -->
{{ template "layout/scripts.html" . }} <!-- 基础JS脚本 -->
{{ template "layout/footer.html" . }}  <!-- 页面底部 -->
```

## 2. 标准页面结构
```html
{{ template "layout/header.html" . }}
{{ template "layout/nav.html" . }}

<!-- Content Wrapper -->
<div class="content-wrapper">
    <!-- Content Header -->
    <section class="content-header">
        <div class="container-fluid">
            <div class="row mb-2">
                <div class="col-sm-6">
                    <h1>页面标题</h1>
                </div>
                <div class="col-sm-6">
                    <ol class="breadcrumb float-sm-right">
                        <li class="breadcrumb-item"><a href="/">首页</a></li>
                        <li class="breadcrumb-item active">当前页面</li>
                    </ol>
                </div>
            </div>
        </div>
    </section>

    <!-- Main content -->
    <section class="content">
        <div class="container-fluid">
            <!-- 在这里添加你的内容 -->
        </div>
    </section>
</div>

{{ template "layout/scripts.html" . }}
<!-- 在这里添加页面特定的脚本 -->
{{ template "layout/footer.html" . }}
```

## 3. 常用组件示例

### 3.1 卡片组件
```html
<div class="card">
    <div class="card-header">
        <h3 class="card-title">标题</h3>
    </div>
    <div class="card-body">
        <!-- 内容 -->
    </div>
    <div class="card-footer">
        <!-- 底部 -->
    </div>
</div>
```

### 3.2 表单组件
```html
<form>
    <div class="card-body">
        <div class="form-group">
            <label for="fieldId">标签</label>
            <input type="text" class="form-control" id="fieldId" placeholder="提示文本">
        </div>
        <div class="custom-control custom-switch">
            <input type="checkbox" class="custom-control-input" id="switchId">
            <label class="custom-control-label" for="switchId">开关</label>
        </div>
    </div>
</form>
```

### 3.3 表格组件
```html
<div class="card">
    <div class="card-header">
        <h3 class="card-title">表格标题</h3>
    </div>
    <div class="card-body table-responsive p-0">
        <table class="table table-hover text-nowrap">
            <thead>
                <tr>
                    <th>列标题</th>
                </tr>
            </thead>
            <tbody>
                <tr>
                    <td>内容</td>
                </tr>
            </tbody>
        </table>
    </div>
</div>
```

## 4. 常用样式类

### 4.1 颜色类
- 背景色：`bg-primary`, `bg-success`, `bg-warning`, `bg-danger`, `bg-info`
- 文本色：`text-primary`, `text-success`, `text-warning`, `text-danger`, `text-info`
- 按钮：`btn-primary`, `btn-success`, `btn-warning`, `btn-danger`, `btn-info`

### 4.2 间距类
- 边距：`m-1` 到 `m-5` (margin)
- 内边距：`p-1` 到 `p-5` (padding)
- 方向性：`t`(top), `b`(bottom), `l`(left), `r`(right), `x`(左右), `y`(上下)

### 4.3 响应式类
- 列宽：`col-sm-6`, `col-md-4`, `col-lg-3`
- 显示：`d-none`, `d-block`, `d-flex`
- 响应式显示：`d-sm-none`, `d-md-block`

## 5. 最佳实践

### 5.1 页面组织
- 使用 `content-header` 提供清晰的页面标题和面包屑导航
- 将主要内容放在 `content` 部分
- 使用 `container-fluid` 和 `row` 进行布局

### 5.2 脚本管理
- 基础脚本通过 `scripts.html` 引入
- 页面特定的脚本放在 `scripts.html` 模板后面
- 考虑将复杂的页面脚本分离到单独的文件中

### 5.3 响应式设计
- 使用 Bootstrap 的栅格系统进行布局
- 合理使用响应式类适配不同屏幕尺寸
- 使用 `table-responsive` 处理表格在移动设备上的显示

### 5.4 组件复用
- 将常用的组件结构保存为代码片段
- 使用一致的命名和结构方式
- 保持组件的独立性和可复用性

## 6. 常用交互组件

### 6.1 提示框
```javascript
$(document).Toasts('create', {
    class: 'bg-success',
    title: '成功',
    body: '操作成功提示'
});
```

### 6.2 确认框
```html
<div class="modal fade" id="confirmModal">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <h4 class="modal-title">确认</h4>
                <button type="button" class="close" data-dismiss="modal">&times;</button>
            </div>
            <div class="modal-body">
                确认要执行此操作吗？
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">取消</button>
                <button type="button" class="btn btn-primary" id="confirmBtn">确定</button>
            </div>
        </div>
    </div>
</div>
```

### 6.3 加载动画
```html
<div class="overlay">
    <i class="fas fa-2x fa-sync-alt fa-spin"></i>
</div>
```

## 7. 开发技巧

### 7.1 页面模板创建步骤
1. 复制基本模板结构
2. 修改页面标题和面包屑
3. 规划页面布局和组件
4. 添加必要的交互功能
5. 测试响应式布局

### 7.2 调试技巧
- 使用浏览器开发者工具检查布局
- 在不同屏幕尺寸下测试响应式设计
- 检查 JavaScript 控制台的错误信息
- 使用 AdminLTE 提供的调试工具

### 7.3 性能优化
- 合理使用组件，避免过度嵌套
- 优化图片和资源加载
- 使用适当的缓存策略
- 延迟加载非关键资源

## 8. 常见问题解决

### 8.1 布局问题
- 确保正确使用 `content-wrapper` 和 `container-fluid`
- 检查响应式类的使用是否正确
- 验证 HTML 结构的完整性

### 8.2 样式冲突
- 使用特定的类名避免冲突
- 检查样式优先级
- 合理使用自定义样式

### 8.3 JavaScript 相关
- 确保脚本加载顺序正确
- 使用正确的选择器
- 处理好异步操作的回调 