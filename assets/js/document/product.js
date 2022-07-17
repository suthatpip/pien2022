$(document).ready(function() { 
    filelist();

    
    
const id = "#custom_file";
const dropzone = document.querySelector(id); 
var previewNode = dropzone.querySelector(".dropzone-item");
previewNode.id = "";
var previewTemplate = previewNode.parentNode.innerHTML;
previewNode.parentNode.removeChild(previewNode);

var uploadDropzone = new Dropzone(id, { // Make the whole body a dropzone
    url: "/product/new", // Set the url for your upload script location
    parallelUploads: 20,
    previewTemplate: previewTemplate,
    maxFilesize: 1, // Max filesize in MB
    maxFiles:5,
    autoQueue: false, // Make sure the files aren't queued until manually added
    previewsContainer: id + " .dropzone-items", // Define the container to display the previews
    clickable: id + " .dropzone-select" // Define the element that should be used as click trigger to select files.
});

uploadDropzone.on("addedfile", function (file) {
    console.log("addedfile");
    // Hookup the start button
    file.previewElement.querySelector(id + " .dropzone-start").onclick = function () { uploadDropzone.enqueueFile(file); };
    const dropzoneItems = dropzone.querySelectorAll('.dropzone-item');
    dropzoneItems.forEach(dropzoneItem => {
        dropzoneItem.style.display = '';
    });
    dropzone.querySelector('.dropzone-upload').style.display = "inline-block";
    dropzone.querySelector('.dropzone-remove-all').style.display = "inline-block";
});

// Update the total progress bar
uploadDropzone.on("totaluploadprogress", function (progress) {
    console.log("totaluploadprogress");
    const progressBars = dropzone.querySelectorAll('.progress-bar');
    progressBars.forEach(progressBar => {
        progressBar.style.width = progress + "%";
    });
});

uploadDropzone.on("sending", function (file) {
    console.log("sending ");
    // Show the total progress bar when upload starts
    const progressBars = dropzone.querySelectorAll('.progress-bar');
    progressBars.forEach(progressBar => {
        progressBar.style.opacity = "1";
    });
    // And disable the start button
    file.previewElement.querySelector(id + " .dropzone-start").setAttribute("disabled", "disabled");
});

// Hide the total progress bar when nothing's uploading anymore
uploadDropzone.on("complete", function (progress) {
    console.log("complete " );
    const progressBars = dropzone.querySelectorAll('.dz-complete');

    setTimeout(function () {
        console.log("complete");
        progressBars.forEach(progressBar => {
            progressBar.querySelector('.progress-bar').style.opacity = "0";
            progressBar.querySelector('.progress').style.opacity = "0";
            progressBar.querySelector('.dropzone-start').style.opacity = "0";
        });
    }, 300);
});

uploadDropzone.on("successmultiple", function(files, response) {
    console.log("successmultiple "+ response );
});

uploadDropzone.on("success", function(files, response) {
    console.log("success "+ response.data );
});

// Setup the buttons for all transfers
dropzone.querySelector(".dropzone-upload").addEventListener('click', function () {
    console.log("dropzone-upload");
    uploadDropzone.enqueueFiles(uploadDropzone.getFilesWithStatus(Dropzone.ADDED));
});

// Setup the button for remove all files
dropzone.querySelector(".dropzone-remove-all").addEventListener('click', function () {
    console.log("dropzone-remove-all");
    dropzone.querySelector('.dropzone-upload').style.display = "none";
    dropzone.querySelector('.dropzone-remove-all').style.display = "none";
    uploadDropzone.removeAllFiles(true);
});

// On all files completed upload
uploadDropzone.on("queuecomplete", function (progress) {
    console.log("queuecomplete");
    const uploadIcons = dropzone.querySelectorAll('.dropzone-upload');
    uploadIcons.forEach(uploadIcon => {
        uploadIcon.style.display = "none";
    });
});

// On all files removed
uploadDropzone.on("removedfile", function (file) {
    console.log("removedfile");
    if (uploadDropzone.files.length < 1) {
        dropzone.querySelector('.dropzone-upload').style.display = "none";
        dropzone.querySelector('.dropzone-remove-all').style.display = "none";
    }
});


})
function filelist(){ 
    var html=``
    fetch(url + '/product/list', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        }
      })
      .then(response => response.json())
      .then(data => {
        html +=`<table class="table table-hover table-striped"><tbody>`
        data.forEach(obj => {
        html +=`           
            <tr>
            <td>
                <div class="icheck-primary">
                <input type="checkbox" value="${obj.product_code}" name="file_group[]">
                <label for="check1"></label>
                </div>
            </td>

            <td>` 
            if (obj.product_connect=="connect"){
                html +=`<i class="fas fa-solid fa-plug text-success"></i>` 
            }else{
                html +=`<i class="fas fa-solid fa-plug text-secondary" style="opacity: 0.3"></i>`
            }
                
            html +=`</td>
            <td><a href="${obj.product_detail}">${obj.product_name}</a></td>
            <td><b>${obj.product_size}</b></td>
            <td></td>
            <td>${obj.product_create_date}</td>
            </tr> `  
        });        
        html +=`</tbody></table>`
        $('#myfiles').html(html);

      }).catch((error) => {
        console.error('Error:', error)
      })
}


function submitdataf1() {   
  
    var files = new Array();
    $.each($("input[name='file_group[]']:checked"), function() {
        files.push($(this).val());      
    });

    if (files.length == 0 ){
        Swal.fire('เลือกไฟล์อย่างน้อย 1 ไฟล์')
       
    }else{
        stepper.next();
    }
}
 
function submitdataf2() {  
    var company_code = $('#companys').val();
    if (company_code===null){ 
        Swal.fire('การเลือกบริษัทยังไม่สมบูรณ์')
        return
    }else{       
        var file = {};
        var files = [];
        $.each($("input[name='file_group[]']:checked"), function() {     
            file.code= $(this).val();          
            files.push({...file}); 
        });
        var obj= JSON.stringify({
          start_date: publish_start,
          end_date: publish_end,
          company_code: company_code,
          products: files
        })
        // $("#processing").modal('show');
        initpayment(obj);
    } 
  } 

  function deletefile(){
    var file = {};
    var files = [];
     
    $.each($("input[name='file_group[]']:checked"), function() {     
          file.product_code= $(this).val();  
 
        files.push({...file});
    });
   
    fetch(url + '/product/delete', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ 
            products: files
        })    
      })
      .then(function (response) { 
        filelist();
      }) 
      .catch(err =>{
          
      });
  }

  function previous(){
    var payment_code=$('#_payment_code').text();
    fetch(url + '/payment/delete', {
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