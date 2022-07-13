var systemFormatDate = "DD/MM/YYYY"
var formatDate = "DD/MM/YYYY"
var dateObj = new Date();
var month = String(dateObj.getMonth() + 1).padStart(2, '0');
var day = String(dateObj.getDate()).padStart(2, '0');
var year = dateObj.getFullYear() + 543;
var today = String(day + '/' + month + '/' + year);
var publish_start = moment(today, formatDate).format(systemFormatDate);
var publish_end = moment(today, formatDate).format(systemFormatDate);
moment.locale('th')

$('input[data-publish-date="publish-date"]').daterangepicker({
    "maxYear": 2023,
    "showCustomRangeLabel": false,
    "locale": {
      "format": "DD MMM YYYY",
      "separator": " - ",
      "applyLabel": "Apply",
      "cancelLabel": "Cancel",
      "fromLabel": "From",
      "toLabel": "To",
      "customRangeLabel": "Custom",
      "weekLabel": "W",
      "daysOfWeek": [
        "อา",
        "จ",
        "อ",
        "พ",
        "พฤ",
        "ศ",
        "ส"
      ],
      "monthNames": [
        "มกราคม ",
        "กุมภาพันธ์ ",
        "มีนาคม ",
        "เมษายน",
        "พฤษภาคม",
        "มิถุนายน",
        "กรกฎาคม",
        "สิงหาคม",
        "กันยายน",
        "ตุลาคม",
        "พฤศจิกายน",
        "ธันวาคม"
      ],
      "firstDay": 1
    },
    "startDate": moment(today, formatDate).format('LL'),
    "endDate": moment(today, formatDate).format('LL'),
    "minDate": moment(today, formatDate).format('LL')
});

$('input[data-publish-date="publish-date"]').on('apply.daterangepicker', function (ev, picker) {
    publish_start = picker.startDate.format(systemFormatDate);
    publish_end = picker.endDate.format(systemFormatDate);
});

$('input[data-publish-date="publish-date"]').on('hide.daterangepicker', function (ev, picker) {
    publish_start = picker.startDate.format(systemFormatDate);
    publish_end = picker.endDate.format(systemFormatDate);
});

$('.select2').select2()
  