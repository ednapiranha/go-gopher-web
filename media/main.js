$(function () {
  var form = $('form');
  var page = $('#page');

  form.on('submit', function (ev) {
    ev.preventDefault();

    $.post('/get', form.serialize(), function (data) {
      data = JSON.parse(data);
      var results = [];

      data.GopherMap.forEach(function (i) {
        switch (i.ItemType) {
          case '1':
          case 'h':
            results.push('<a href="' + i.HostName + i.Selector + '">' + i.DisplayString + "</a>");
            break;
          default:
            results.push(i.DisplayString);
            break;
        }
      });
      console.log(results)
      page.text(results.join('\n'));
    });
  });
});
