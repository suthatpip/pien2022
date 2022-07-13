 getmenu();
 
$('#document_name').on('keyup', function (e) {
  if (isEmpty($('#document_name').val())) {
    $('#document_name').addClass("is-invalid");
    $('#document_name').removeClass("is-valid");
  } else {
    $('#document_name').removeClass("is-invalid");
    $('#document_name').addClass("is-valid");
  }
});

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
  

 
function submitdatac1() {   
  var data= editor.getData(); 

  if (isEmpty(data)) {  
    Swal.fire('เอกสารยังไม่สมบูรณ์')
     
  }else{
    stepper.next();
    
  }
}

function newDocument(){
  var document_no = $('#document_no').text();
  var document_name = $('#document_name').val();
  var document_data= editor.getData();  
  
  fetch(url + '/template/new', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        product_code: document_no,
        product_name: document_name,         
        product_detail: document_data 
      })
    })
    .then(function (response) {
      if (!response.ok) { 
        Swal.fire('เกิดข้อผิดพลาด'); 
        throw new Error("HTTP status " + response.status);
      }     
        return;
    })
    .then(data => { 
     
    }).catch((error) => {
      
    })
}


function submitdatac2() {
  var file_name = $('#document_name').val();
  var company_code = $('#companys').val();
  var document_no = $('#document_no').text(); 
 
  $('#document_name').removeClass("is-valid");
  if (isEmpty(file_name)) {  
    $('#document_name').addClass("is-invalid");
    Swal.fire('ชื่อเอกสารยังไม่สมบูรณ์')
    return
  }  
  
  $('#companys').removeClass("border border-danger");
  if (company_code===null){ 
    $('#companys').addClass("border border-danger");
    Swal.fire('การเลือกบริษัทยังไม่สมบูรณ์')
    return
  }     
    newDocument();  
    var file = {};
    var files = [];   
    file.code= document_no;    
    files.push({...file}); 
  
    var obj= JSON.stringify({
      start_date: publish_start,
      end_date: publish_end,
      company_code: company_code,
      products: files
    })
    lockobj();
    // $("#processing").modal('show');
    initpayment(obj);  
}

function previous(){
  var payment_code=$('#_payment_code').text();
  fetch(url + '/payment/delete/all', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ 
          payment_code: payment_code
      })    
    })
    .then(function (response) {  
      stepper.previous();
    }) 
    .catch(err =>{
    });
}
 
function getmenu(){
  var pathname = window.location.pathname;
  pathname=pathname.replaceAll ("/","-","/g").substring(1); 
    $('#' + pathname).addClass("active");
}

function lockobj(){
  $("#document_name").attr('disabled','disabled');
}

function unlockobj(){  
   $('#document_name').removeAttr('disabled');
 
}
 
 