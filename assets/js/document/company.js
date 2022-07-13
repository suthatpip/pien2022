Dropzone.autoDiscover = false
var previewNode = document.querySelector("#template")
previewNode.id = ""
var previewTemplate = previewNode.parentNode.innerHTML
previewNode.parentNode.removeChild(previewNode)
var companyDropzone = new Dropzone(document.body, {
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
    companyDropzone.enqueueFiles(companyDropzone.getFilesWithStatus(Dropzone.ADDED))
  }
}

document.querySelector(".cancel").onclick = function () {
  companyDropzone.removeAllFiles(true)
}
  companyDropzone.on("success", function (file, response) {
    companySave();
  })
  
  // this.prototype.events = ["drop", "dragstart", "dragend", "dragenter", "dragover", 
  // "dragleave", "addedfile", "addedfiles", "removedfile", 
  // "thumbnail", "error", "errormultiple", "processing", 
  // "processingmultiple", "uploadprogress", "totaluploadprogress", 
  // "sending", "sendingmultiple", "success", "successmultiple", 
  // "canceled", "canceledmultiple", "complete", "completemultiple", 
  // "reset", "maxfilesexceeded", "maxfilesreached", "queuecomplete"];
    
  companyDropzone.on("maxfilesexceeded", file => { 
    companyDropzone.removeAllFiles(true)
    companyDropzone.addFile(file);
    $('#companySave').prop( "disabled", false );
  });  
  
  companyDropzone.on("maxfilesreached", file => { 
    $('#companySave').prop( "disabled", false );
  }); 
  
  companyDropzone.on("error", file => {  
    companyDropzone.removeAllFiles(true);
    $('#companySave').prop( "disabled", true );
  }); 
  companyList();

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
    companyDropzone.removeAllFiles(true)
  
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