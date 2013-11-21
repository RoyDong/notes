$(document).ready(function() {
    var searchBtn = $('#header button');
    var query = $('#search');

    query.keyup(function(e) {
        if (e.which == 13 && query.val().length > 0) {
            location.href = "/?q=" + query.val();
        }
    });

    searchBtn.click(function() {
        if (query.val().length > 0) {
            location.href = "/?q=" + query.val();
        }
    });

    $(document).keyup(function(e) {
        if (e.which == 70) {
            query.focus();
        }
    });
});
