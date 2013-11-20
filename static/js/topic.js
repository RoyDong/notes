$(document).ready(function() {
    var editor = KindEditor.create('#comment-content', {
        width: 938,
        height: 250,
        'min-height': 250,
        themeType: 'comment',
        items: [
            'undo', 'redo', '|', 
            'forecolor', 'bold', 'italic', 'underline', 'plainpaste', '|',
            'link', 'unlink', 'code', 'preview'
        ]
    });

    var msg = $('#error-message');
    var isSubmiting = false;
    var commentList = $('.comments-list');

    $('#comment-submit').click(function() {
        if (isSubmiting) {
            msg.text("正在提交，请耐心等待");
            msg.css("display", "block");
            return;
        }

        if (editor.html().length < 10) {
            msg.css("display", "block");
            msg.text("不能少于10个字");
            return;
        }

        isSubmiting = true;
        msg.text("");
        msg.css("display", "none");
        var me = $(this);
        $.ajax({
            url: me.data('url'),
            method: 'post',
            data: {tid: me.data('tid'), content: editor.html()},
            complete: function() {
                isSubmiting = false;
            },
            success: function(data) {
                editor.html("");
                commentList.append(data);
            },
            failure: function() {
            }
        });
    });
});
