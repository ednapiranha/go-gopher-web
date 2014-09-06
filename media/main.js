$(function () {
  var form = $('form');
  var page = $('#page');
  var url = $('#url');
  var results = [];

  page.on('click', 'a', function (ev) {
    ev.preventDefault();

    url.val($(this).prop('href'));
    form.submit();
  });

  form.on('submit', function (ev) {
    ev.preventDefault();

    page.empty();
    results = [];

    $.post('/get', form.serialize(), function (data) {
      data = JSON.parse(data);

      data.GopherMap.forEach(function (i) {
        switch (i.ItemType) {
          case '1':
          case 'h':
            results.push('<a href="gopher://' + i.HostName + ':70/1' + i.Selector + '">' + i.DisplayString + "</a>");
            break;
          default:
            results.push(i.DisplayString);
            break;
        }
      });

      page.html(results.join('\n'));
    });
  });
});
