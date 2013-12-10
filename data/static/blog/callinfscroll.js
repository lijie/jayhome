/* @script.js文件 */

$(document).ready(function () {
	// layout
	$(window).resize(reAdjust);
	reAdjust();
	$(window).scroll(sdSize);  //for all but ie6&7

	if ($.browser.msie && $.browser.version === '10.0') {  // <ul>,<ol>
		$('html').addClass('ie10');
	}
	
	if ($('body').hasClass('p-homepage')) {
		// infinite scroll
		if (infTag) {
			$('.m-postlst').infinitescroll({
				navSelector : "#m-pageridx",
				nextSelector : "#next_page_link",
				itemSelector : ".m-postlst .m-post",
				loading : {
					msgText  : "",
					finishedMsg : "",
					img: "http://img.ph.126.net/5O_7n6mY2Xc1Zj4NSh9TGw==/6597088458353674794.gif"
				}
			}, 
			function(){
				sdSize();
			});
		} else {
			// pager
			$('#m-pager-idx .active').bind('click', function () {	
				return false;
			});
		}
	}
	// search
	$('#j-lnksch').bind('click', function () {
		$(this).css('visibility', 'hidden').parent().addClass('m-schshow');
		setTimeout(function () {
			$('#j-schform .txt').focus();
		}, 300);
		return false;
	});
	$('#j-schform .txt').bind('blur', function () {
		$('.m-sch').removeClass('m-schshow');
		setTimeout(function () {
			$('#j-lnksch').css('visibility', 'visible');
		}, 300);
		return false;
	});
});
function sdSize () {  // min-height
	var sdHeight = $('.g-sd').height();
	var mnHeight = $('.g-mn').height() + 80;        
	var winheight = $(window).height();
	var minHeight = mnHeight > winheight ? mnHeight : winheight;
	minHeight = minHeight - 140;
	if (minHeight > sdHeight) {
		sdHeight = minHeight;
	}
	$('.g-sd').css('height', sdHeight);        
}
function bodySize () {  // min-width
	var bWidth = ($(window).width() <= mWidth) ? mWidth : '100%';
	$('body').width(bWidth);
}
function reAdjust() {
	sdSize();
	if ($.browser.msie && $.browser.version === '6.0') {
		bodySize();
	}
}