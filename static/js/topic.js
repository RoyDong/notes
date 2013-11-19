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



    $('#comment-submit').click(function() {
        var me = $(this);
        $.ajax({
            url: me.data('url'),
            method: 'post',
            data: {tid: me.data('tid'), content: editor.html()},
            success: function(data) {
                console.log(data);
            },
            failure: function(data) {
                console.log(data);
            }
        });
    });
});
