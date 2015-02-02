
// submit form
$("#submit").on("click", function(event) {
    console.log("submit");
    var postdata = $("#form_item_add").serialize();
    console.log(postdata);
    $.ajax({
	type: 'POST',
	url: 'item?fn=add',
	data: postdata,
	cache: false,
	dataType: 'json',
	success: function(data) {
	    console.log(data);
	    if (data.result == "ok") {
		$('#alert_info').addClass("alert-success")
		$('#alert_info').text("Save success!")
	    }
	},
	error: function() {
	    console.log("error");
	},
    });
});

// file upload
/*jslint unparam: true */
/*global window, $ */
$(function () {
    'use strict';
    // Change this to the location of your server-side upload handler:
    var url = window.location.hostname === 'blueimp.github.io' ?
        '//jquery-file-upload.appspot.com/' : '/admin/upload';
    $('#item_small_image').fileupload({
        url: url,
        dataType: 'json',
        done: function (e, data) {
            // $('<p/>').text(data.name).appendTo('#files');
            // $("#userimg").attr("src", "/uploadfiles/" + data.result.name);
	    $("#item_small_image_url").attr("value", data.result.url)
	    $("#item_small_image_preview").attr("src", data.result.url)
	    console.log(data.result.url)
        },
        progressall: function (e, data) {
            var progress = parseInt(data.loaded / data.total * 100, 10);
            $('#progress .progress-bar').css(
                'width',
                progress + '%'
            );
        }
    }).prop('disabled', !$.support.fileInput)
        .parent().addClass($.support.fileInput ? undefined : 'disabled');
});

$(function () {
    'use strict';
    // Change this to the location of your server-side upload handler:
    var url = window.location.hostname === 'blueimp.github.io' ?
        '//jquery-file-upload.appspot.com/' : '/admin/upload';
    $('#item_big_image').fileupload({
        url: url,
        dataType: 'json',
        done: function (e, data) {
            // $('<p/>').text(data.name).appendTo('#files');
            // $("#userimg").attr("src", "/uploadfiles/" + data.result.name);
	    $("#item_big_image_url").attr("value", data.result.url)
	    $("#item_big_image_preview").attr("src", data.result.url)
	    console.log(data.result.url)
        },
        progressall: function (e, data) {
            var progress = parseInt(data.loaded / data.total * 100, 10);
            $('#progress .progress-bar').css(
                'width',
                progress + '%'
            );
        }
    }).prop('disabled', !$.support.fileInput)
        .parent().addClass($.support.fileInput ? undefined : 'disabled');
});
