$(function () {
    'use strict'

    // 加载路由列表
    function loadRoutes() {
        $.get('/api/routes', function(routes) {
            const tbody = $('#routeList');
            tbody.empty();
            routes.forEach(route => {
                tbody.append(`
                    <tr>
                        <td>${route.path}</td>
                        <td>${route.template}</td>
                        <td>${route.title}</td>
                        <td>
                            <span class="badge badge-${route.enabled ? 'success' : 'danger'}">
                                ${route.enabled ? '启用' : '禁用'}
                            </span>
                        </td>
                        <td>
                            <button class="btn btn-info btn-xs" onclick="RouteManager.editRoute(${route.id})">
                                <i class="fas fa-edit"></i> 编辑
                            </button>
                            <button class="btn btn-danger btn-xs" onclick="RouteManager.deleteRoute(${route.id})">
                                <i class="fas fa-trash"></i> 删除
                            </button>
                        </td>
                    </tr>
                `);
            });
        });
    }

    // 路由管理器对象
    window.RouteManager = {
        // 编辑路由
        editRoute: function(id) {
            $.get(`/api/routes/${id}`, function(route) {
                $('#routeId').val(route.id);
                $('#path').val(route.path);
                $('#template').val(route.template);
                $('#title').val(route.title);
                $('#enabled').prop('checked', route.enabled);
            });
        },

        // 删除路由
        deleteRoute: function(id) {
            if (confirm('确定要删除这个路由吗？')) {
                $.ajax({
                    url: `/api/routes/${id}`,
                    method: 'DELETE',
                    success: function() {
                        loadRoutes();
                        // 显示成功提示
                        $(document).Toasts('create', {
                            class: 'bg-success',
                            title: '成功',
                            body: '路由已成功删除'
                        });
                    }
                });
            }
        },

        // 重置表单
        resetForm: function() {
            $('#routeForm')[0].reset();
            $('#routeId').val('');
        }
    };

    // 表单提交处理
    $('#routeForm').submit(function(e) {
        e.preventDefault();
        const id = $('#routeId').val();
        const data = {
            path: $('#path').val(),
            template: $('#template').val(),
            title: $('#title').val(),
            enabled: $('#enabled').is(':checked')
        };

        const method = id ? 'PUT' : 'POST';
        const url = id ? `/api/routes/${id}` : '/api/routes';

        $.ajax({
            url: url,
            method: method,
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function() {
                RouteManager.resetForm();
                loadRoutes();
                // 显示成功提示
                $(document).Toasts('create', {
                    class: 'bg-success',
                    title: '成功',
                    body: id ? '路由已更新' : '路由已添加'
                });
            }
        });
    });

    // 初始化
    loadRoutes();
}); 