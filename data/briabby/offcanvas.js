$('#myModal').on('show.bs.modal', function (event) {
    var button = $(event.relatedTarget) // Button that triggered the modal
    var recipient = button.data('imagebig') // Extract info from data-* attributes
    // If necessary, you could initiate an AJAX request here (and then do the updating in a callback).
    // Update the modal's content. We'll use jQuery here, but you could use a data binding library or other methods instead.
    var modal = $(this)
    console.log(recipient)
    $('#preview-img').attr("src", recipient)
});

function showPageNav(currentPage, maxPage) {
    var p = $('#ba-page-nav')
    if (currentPage > 1) {
	p.append('<li><a href="' + (currentPage - 1) + '"><span aria-hidden="true">&laquo;</span><span class="sr-only">Previous</span></a></li>')
    }
    for (i = 1; i <= maxPage; i++) {
	p.append('<li><a href="hat?page=' + i + '">' + i + '</a></li>')
    }
    if (currentPage < maxPage) {
	p.append('<li><a href="' + (currentPage + 1) + '"><span aria-hidden="true">&raquo;</span><span class="sr-only">Next</span></a></li>')
    }
}
