
  getmenu();
  var url = location.protocol + '//' + location.host;
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
  Dropzone.autoDiscover = false
  var previewNode = document.querySelector("#template")
  previewNode.id = ""
  var previewTemplate = previewNode.parentNode.innerHTML
  previewNode.parentNode.removeChild(previewNode)
  var myDropzone = new Dropzone(document.body, {
    url: "/company/logo",
    thumbnailWidth: 80,
    thumbnailHeight: 80,
    parallelUploads: 1,
    maxFilesize: 1,
    maxFiles:1,
    acceptedFiles: "image/jpeg,image/png",
    previewTemplate: previewTemplate,
    autoQueue: false,
    previewsContainer: "#previews",
    clickable: ".fileinput-button"
  })

  document.querySelector(".start").onclick = function () {
    if (companyValidate()) {
      myDropzone.enqueueFiles(myDropzone.getFilesWithStatus(Dropzone.ADDED))
    }
  }
  
  document.querySelector(".cancel").onclick = function () {
    myDropzone.removeAllFiles(true)
  }
  
  myDropzone.on("success", function (file, response) {
    companySave();
  })
  
  // this.prototype.events = ["drop", "dragstart", "dragend", "dragenter", "dragover", 
  // "dragleave", "addedfile", "addedfiles", "removedfile", 
  // "thumbnail", "error", "errormultiple", "processing", 
  // "processingmultiple", "uploadprogress", "totaluploadprogress", 
  // "sending", "sendingmultiple", "success", "successmultiple", 
  // "canceled", "canceledmultiple", "complete", "completemultiple", 
  // "reset", "maxfilesexceeded", "maxfilesreached", "queuecomplete"];
    
  myDropzone.on("maxfilesexceeded", file => { 
    myDropzone.removeAllFiles(true)
    myDropzone.addFile(file);
    $('#companySave').prop( "disabled", false );
  });  
  
  myDropzone.on("maxfilesreached", file => { 
    $('#companySave').prop( "disabled", false );
  }); 
  
  myDropzone.on("error", file => {  
    myDropzone.removeAllFiles(true);
    $('#companySave').prop( "disabled", true );
  }); 
  




  companyList();
 

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

$('#inputDocumentName').on('keyup', function (e) {
  if (isEmpty($('#inputDocumentName').val())) {
    $('#inputDocumentName').addClass("is-invalid");
    $('#inputDocumentName').removeClass("is-valid");
  } else {
    $('#inputDocumentName').removeClass("is-invalid");
    $('#inputDocumentName').addClass("is-valid");
  }
});

$('#template').change(function () {
  fetch(url + '/template/' + $(this).val(), {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
    })
    .then(response => response.json())
    .then(data => {
      editor.setData(data.detail)
    }).catch((error) => {
      console.error('Error:', error)
    })
});

document.addEventListener('DOMContentLoaded', function () {
  window.stepper = new Stepper(document.querySelector('.bs-stepper'))
})

DecoupledDocumentEditor
  .create(document.querySelector('.editor'), {
    licenseKey: '',
  })
  .then(editor => {
    window.editor = editor;
    document.querySelector('.document-editor__toolbar').appendChild(editor.ui.view.toolbar.element);
    document.querySelector('.ck-toolbar').classList.add('ck-reset_all');
  })
  .catch(error => {});
  

function isEmpty(str) {
  return (!str || str.length === 0);
} 

function companyInit() {
  $('#company_name').removeClass("is-invalid");
  $('#company_name').removeClass("is-valid");
  $('#company_address').removeClass("is-invalid");
  $('#company_address').removeClass("is-valid");
  $('#company_telephone').removeClass("is-invalid");
  $('#company_telephone').removeClass("is-valid");
  $('#company_name').val("");
  $('#company_address').val("");
  $('#company_telephone').val("");
  myDropzone.removeAllFiles(true)

  fetch(url + '/company/init', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      }
    })
    .then(response => response.json())
    .then(data => {
      $('#company_code').text(data.code);
      Cookies.set('code', data.code);  
    }).catch((error) => {
      console.error('Error:', error)
    })
}

function companyValidate() {
  if (isEmpty($('#company_name').val())) {
    $('#company_name').addClass("is-invalid");
    $('#company_name').removeClass("is-valid");
    return false;
  } else {
    $('#company_name').removeClass("is-invalid");
    $('#company_name').addClass("is-valid");
  }
  if (isEmpty($('#company_address').val())) {
    $('#company_address').addClass("is-invalid");
    $('#company_address').removeClass("is-valid");
    return false;
  } else {
    $('#company_address').removeClass("is-invalid");
    $('#company_address').addClass("is-valid");
  }
  if (isEmpty($('#company_telephone').val())) {
    $('#company_telephone').addClass("is-invalid");
    $('#company_telephone').removeClass("is-valid");
    return false;
  } else {
    $('#company_telephone').removeClass("is-invalid");
    $('#company_telephone').addClass("is-valid");
  }
  return true;
}

function companySave() {
  var code = $('#company_code').text();
  var name = $('#company_name').val();
  var address = $('#company_address').val();
  var telephone = $('#company_telephone').val();
  fetch(url + '/company/new', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: name,
        address: address,
        telephone: telephone,
        code: code
      })
    })
    // .then(response => response.json())
    .then(data => {
      companyList();
      $('#new-company').modal('hide');
    }).catch((error) => {
      console.error('Error:', error)
    })
}

function companyList() {
  fetch(url + '/company/list', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      }
    })
    .then(response => response.json())
    .then(data => {
      $('#companys').find('option:not(:first)').remove();
      data.forEach(obj => {
        $('#companys').append(`<option value="${obj.id}">${obj.name}</option>`);
      });
    }).catch((error) => {
      console.error('Error:', error)
    })
}

function summary() {
  var file_name = $('#inputDocumentName').val();
  var company_code = $('#companys').val();

  $('#inputDocumentName').removeClass("is-valid");
  if (isEmpty(file_name)) {  
    $('#inputDocumentName').addClass("is-invalid");
    return
  }  
  
  $('#companys').removeClass("border border-danger");
  if (company_code===null){ 
    $('#companys').addClass("border border-danger");
    return
  } 

  $('#processing').modal('show'); 

  fetch(url + '/payment/init', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        start_date: publish_start,
        end_date: publish_end,
        product_name: file_name,
        company_code: company_code
      })
    })
    .then(function (response) {
      if (!response.ok) { 
          $("#processing_message").html(processingError());  
          setTimeout(function() {hideProcessing();},3000);
        throw new Error("HTTP status " + response.status);
      }    
        return response.json();
    })
    .then(function (data) { 
      hideProcessing();  
      $('#_company_name').text(data.company.name);

      $('#_name').text(data.customer_name);
      $('#_address').text(data.company.address);
      $('#_telephone').text(data.company.telephone);
      $('#_customer_code').text(data.company.code);
      $('#_company_logo').attr("src", data.company.logo);

      $('#_order').text(data.order.order_no);
      $('#_payment_due_date').text(data.order.payment_due);
      $('#_tax_invoice_no').text(data.order.tax_invoice_no);
      $('#_payment_code').text(data.order.payment_code);

      $('#_no').text(data.product.no);
      $('#_product').text(data.product.product);
      $('#_start_date').text(data.product.start_date);
      $('#_end_date').text(data.product.end_date);
      $('#_days').text(data.product.days);
      $('#_product_baht').text(data.product.product_baht);
      $('#_sub_total').text(data.sub_total);
      $('#_vat').text(data.vat);
      $('#_total').text(data.total);

      stepper.next()
    })
    .catch(err =>{
        $('#processing').modal('hide'); 
    });
}
 

function hideProcessing(){
  console.log("hideProcessing");
    $("#processing").removeClass("in");
    $(".modal-backdrop").remove();
    $("#processing").hide();
}

function processingError() {
  console.log("processingError");
  return ` 
  <div class="card">
    <div class="card-body"> 
      <br>
      <p class="card-text text-danger"> พบข้อผิดพลาด กรุณาลองใหม่อีกครั้ง :(</p>             
    </div>
  </div> 
  `;
}

function getmenu(){
  var pathname = window.location.pathname;
  pathname=pathname.replaceAll ("/","-","/g").substring(1); 
    $('#' + pathname).addClass("active");
}
 